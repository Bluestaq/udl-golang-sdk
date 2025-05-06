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

// StatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	Options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r StatusService) {
	r = StatusService{}
	r.Options = opts
	return
}

// Service operation to take a single Status as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StatusService) New(ctx context.Context, body StatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Status. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *StatusService) Update(ctx context.Context, id string, body StatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/status/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *StatusService) List(ctx context.Context, query StatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/status"
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
func (r *StatusService) ListAutoPaging(ctx context.Context, query StatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Status object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/status/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *StatusService) Count(ctx context.Context, query StatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/status/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Status record by its unique ID passed as a
// path parameter.
func (r *StatusService) Get(ctx context.Context, id string, query StatusGetParams, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/status/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get all statuses related to a particular entity.
func (r *StatusService) GetByEntityID(ctx context.Context, idEntity string, query StatusGetByEntityIDParams, opts ...option.RequestOption) (res *[]StatusGetByEntityIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if idEntity == "" {
		err = errors.New("missing required idEntity parameter")
		return
	}
	path := fmt.Sprintf("udl/status/byIdEntity/%s", idEntity)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get all statuses related to a particular entity type.
func (r *StatusService) GetByEntityType(ctx context.Context, entityType string, query StatusGetByEntityTypeParams, opts ...option.RequestOption) (res *[]StatusGetByEntityTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityType == "" {
		err = errors.New("missing required entityType parameter")
		return
	}
	path := fmt.Sprintf("udl/status/byEntityType/%s", entityType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *StatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/status/queryhelp"
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
func (r *StatusService) Tuple(ctx context.Context, query StatusTupleParams, opts ...option.RequestOption) (res *[]StatusTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/status/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type StatusListResponse struct {
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
	DataMode StatusListResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusListResponseOpsCap `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusListResponseState                 `json:"state"`
	SubStatusCollection []StatusListResponseSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusListResponseSysCap `json:"sysCap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusListResponse) UnmarshalJSON(data []byte) error {
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
type StatusListResponseDataMode string

const (
	StatusListResponseDataModeReal      StatusListResponseDataMode = "REAL"
	StatusListResponseDataModeTest      StatusListResponseDataMode = "TEST"
	StatusListResponseDataModeSimulated StatusListResponseDataMode = "SIMULATED"
	StatusListResponseDataModeExercise  StatusListResponseDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusListResponseOpsCap string

const (
	StatusListResponseOpsCapFmc StatusListResponseOpsCap = "FMC"
	StatusListResponseOpsCapNmc StatusListResponseOpsCap = "NMC"
	StatusListResponseOpsCapPmc StatusListResponseOpsCap = "PMC"
	StatusListResponseOpsCapUnk StatusListResponseOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusListResponseState string

const (
	StatusListResponseStateUnknown  StatusListResponseState = "UNKNOWN"
	StatusListResponseStateDead     StatusListResponseState = "DEAD"
	StatusListResponseStateActive   StatusListResponseState = "ACTIVE"
	StatusListResponseStateRfActive StatusListResponseState = "RF ACTIVE"
	StatusListResponseStateStandby  StatusListResponseState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
type StatusListResponseSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusListResponseSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *StatusListResponseSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusListResponseSysCap string

const (
	StatusListResponseSysCapFmc StatusListResponseSysCap = "FMC"
	StatusListResponseSysCapNmc StatusListResponseSysCap = "NMC"
	StatusListResponseSysCapPmc StatusListResponseSysCap = "PMC"
	StatusListResponseSysCapUnk StatusListResponseSysCap = "UNK"
)

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type StatusGetResponse struct {
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
	DataMode StatusGetResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusGetResponseOpsCap `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusGetResponseState                 `json:"state"`
	SubStatusCollection []StatusGetResponseSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusGetResponseSysCap `json:"sysCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponse) UnmarshalJSON(data []byte) error {
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
type StatusGetResponseDataMode string

const (
	StatusGetResponseDataModeReal      StatusGetResponseDataMode = "REAL"
	StatusGetResponseDataModeTest      StatusGetResponseDataMode = "TEST"
	StatusGetResponseDataModeSimulated StatusGetResponseDataMode = "SIMULATED"
	StatusGetResponseDataModeExercise  StatusGetResponseDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetResponseOpsCap string

const (
	StatusGetResponseOpsCapFmc StatusGetResponseOpsCap = "FMC"
	StatusGetResponseOpsCapNmc StatusGetResponseOpsCap = "NMC"
	StatusGetResponseOpsCapPmc StatusGetResponseOpsCap = "PMC"
	StatusGetResponseOpsCapUnk StatusGetResponseOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusGetResponseState string

const (
	StatusGetResponseStateUnknown  StatusGetResponseState = "UNKNOWN"
	StatusGetResponseStateDead     StatusGetResponseState = "DEAD"
	StatusGetResponseStateActive   StatusGetResponseState = "ACTIVE"
	StatusGetResponseStateRfActive StatusGetResponseState = "RF ACTIVE"
	StatusGetResponseStateStandby  StatusGetResponseState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
type StatusGetResponseSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetResponseSysCap string

const (
	StatusGetResponseSysCapFmc StatusGetResponseSysCap = "FMC"
	StatusGetResponseSysCapNmc StatusGetResponseSysCap = "NMC"
	StatusGetResponseSysCapPmc StatusGetResponseSysCap = "PMC"
	StatusGetResponseSysCapUnk StatusGetResponseSysCap = "UNK"
)

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type StatusGetByEntityIDResponse struct {
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
	DataMode StatusGetByEntityIDResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusGetByEntityIDResponseOpsCap `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusGetByEntityIDResponseState                 `json:"state"`
	SubStatusCollection []StatusGetByEntityIDResponseSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusGetByEntityIDResponseSysCap `json:"sysCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetByEntityIDResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusGetByEntityIDResponse) UnmarshalJSON(data []byte) error {
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
type StatusGetByEntityIDResponseDataMode string

const (
	StatusGetByEntityIDResponseDataModeReal      StatusGetByEntityIDResponseDataMode = "REAL"
	StatusGetByEntityIDResponseDataModeTest      StatusGetByEntityIDResponseDataMode = "TEST"
	StatusGetByEntityIDResponseDataModeSimulated StatusGetByEntityIDResponseDataMode = "SIMULATED"
	StatusGetByEntityIDResponseDataModeExercise  StatusGetByEntityIDResponseDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetByEntityIDResponseOpsCap string

const (
	StatusGetByEntityIDResponseOpsCapFmc StatusGetByEntityIDResponseOpsCap = "FMC"
	StatusGetByEntityIDResponseOpsCapNmc StatusGetByEntityIDResponseOpsCap = "NMC"
	StatusGetByEntityIDResponseOpsCapPmc StatusGetByEntityIDResponseOpsCap = "PMC"
	StatusGetByEntityIDResponseOpsCapUnk StatusGetByEntityIDResponseOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusGetByEntityIDResponseState string

const (
	StatusGetByEntityIDResponseStateUnknown  StatusGetByEntityIDResponseState = "UNKNOWN"
	StatusGetByEntityIDResponseStateDead     StatusGetByEntityIDResponseState = "DEAD"
	StatusGetByEntityIDResponseStateActive   StatusGetByEntityIDResponseState = "ACTIVE"
	StatusGetByEntityIDResponseStateRfActive StatusGetByEntityIDResponseState = "RF ACTIVE"
	StatusGetByEntityIDResponseStateStandby  StatusGetByEntityIDResponseState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
type StatusGetByEntityIDResponseSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetByEntityIDResponseSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *StatusGetByEntityIDResponseSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetByEntityIDResponseSysCap string

const (
	StatusGetByEntityIDResponseSysCapFmc StatusGetByEntityIDResponseSysCap = "FMC"
	StatusGetByEntityIDResponseSysCapNmc StatusGetByEntityIDResponseSysCap = "NMC"
	StatusGetByEntityIDResponseSysCapPmc StatusGetByEntityIDResponseSysCap = "PMC"
	StatusGetByEntityIDResponseSysCapUnk StatusGetByEntityIDResponseSysCap = "UNK"
)

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type StatusGetByEntityTypeResponse struct {
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
	DataMode StatusGetByEntityTypeResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusGetByEntityTypeResponseOpsCap `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusGetByEntityTypeResponseState                 `json:"state"`
	SubStatusCollection []StatusGetByEntityTypeResponseSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusGetByEntityTypeResponseSysCap `json:"sysCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetByEntityTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusGetByEntityTypeResponse) UnmarshalJSON(data []byte) error {
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
type StatusGetByEntityTypeResponseDataMode string

const (
	StatusGetByEntityTypeResponseDataModeReal      StatusGetByEntityTypeResponseDataMode = "REAL"
	StatusGetByEntityTypeResponseDataModeTest      StatusGetByEntityTypeResponseDataMode = "TEST"
	StatusGetByEntityTypeResponseDataModeSimulated StatusGetByEntityTypeResponseDataMode = "SIMULATED"
	StatusGetByEntityTypeResponseDataModeExercise  StatusGetByEntityTypeResponseDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetByEntityTypeResponseOpsCap string

const (
	StatusGetByEntityTypeResponseOpsCapFmc StatusGetByEntityTypeResponseOpsCap = "FMC"
	StatusGetByEntityTypeResponseOpsCapNmc StatusGetByEntityTypeResponseOpsCap = "NMC"
	StatusGetByEntityTypeResponseOpsCapPmc StatusGetByEntityTypeResponseOpsCap = "PMC"
	StatusGetByEntityTypeResponseOpsCapUnk StatusGetByEntityTypeResponseOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusGetByEntityTypeResponseState string

const (
	StatusGetByEntityTypeResponseStateUnknown  StatusGetByEntityTypeResponseState = "UNKNOWN"
	StatusGetByEntityTypeResponseStateDead     StatusGetByEntityTypeResponseState = "DEAD"
	StatusGetByEntityTypeResponseStateActive   StatusGetByEntityTypeResponseState = "ACTIVE"
	StatusGetByEntityTypeResponseStateRfActive StatusGetByEntityTypeResponseState = "RF ACTIVE"
	StatusGetByEntityTypeResponseStateStandby  StatusGetByEntityTypeResponseState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
type StatusGetByEntityTypeResponseSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetByEntityTypeResponseSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *StatusGetByEntityTypeResponseSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusGetByEntityTypeResponseSysCap string

const (
	StatusGetByEntityTypeResponseSysCapFmc StatusGetByEntityTypeResponseSysCap = "FMC"
	StatusGetByEntityTypeResponseSysCapNmc StatusGetByEntityTypeResponseSysCap = "NMC"
	StatusGetByEntityTypeResponseSysCapPmc StatusGetByEntityTypeResponseSysCap = "PMC"
	StatusGetByEntityTypeResponseSysCapUnk StatusGetByEntityTypeResponseSysCap = "UNK"
)

// Status for a particular Entity. An entity may have multiple status records
// collected by various sources.
type StatusTupleResponse struct {
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
	DataMode StatusTupleResponseDataMode `json:"dataMode,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate time.Time `json:"declassificationDate" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString string `json:"declassificationString"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom string `json:"derivedFrom"`
	// Comments describing the status creation and or updates to an entity.
	Notes string `json:"notes"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusTupleResponseOpsCap `json:"opsCap"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusTupleResponseState                 `json:"state"`
	SubStatusCollection []StatusTupleResponseSubStatusCollection `json:"subStatusCollection"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusTupleResponseSysCap `json:"sysCap"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		IDEntity               respjson.Field
		Source                 respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		DeclassificationDate   respjson.Field
		DeclassificationString respjson.Field
		DerivedFrom            respjson.Field
		Notes                  respjson.Field
		OpsCap                 respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		State                  respjson.Field
		SubStatusCollection    respjson.Field
		SysCap                 respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusTupleResponse) UnmarshalJSON(data []byte) error {
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
type StatusTupleResponseDataMode string

const (
	StatusTupleResponseDataModeReal      StatusTupleResponseDataMode = "REAL"
	StatusTupleResponseDataModeTest      StatusTupleResponseDataMode = "TEST"
	StatusTupleResponseDataModeSimulated StatusTupleResponseDataMode = "SIMULATED"
	StatusTupleResponseDataModeExercise  StatusTupleResponseDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusTupleResponseOpsCap string

const (
	StatusTupleResponseOpsCapFmc StatusTupleResponseOpsCap = "FMC"
	StatusTupleResponseOpsCapNmc StatusTupleResponseOpsCap = "NMC"
	StatusTupleResponseOpsCapPmc StatusTupleResponseOpsCap = "PMC"
	StatusTupleResponseOpsCapUnk StatusTupleResponseOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusTupleResponseState string

const (
	StatusTupleResponseStateUnknown  StatusTupleResponseState = "UNKNOWN"
	StatusTupleResponseStateDead     StatusTupleResponseState = "DEAD"
	StatusTupleResponseStateActive   StatusTupleResponseState = "ACTIVE"
	StatusTupleResponseStateRfActive StatusTupleResponseState = "RF ACTIVE"
	StatusTupleResponseStateStandby  StatusTupleResponseState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
type StatusTupleResponseSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Notes                 respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		StatusID              respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusTupleResponseSubStatusCollection) RawJSON() string { return r.JSON.raw }
func (r *StatusTupleResponseSubStatusCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusTupleResponseSysCap string

const (
	StatusTupleResponseSysCapFmc StatusTupleResponseSysCap = "FMC"
	StatusTupleResponseSysCapNmc StatusTupleResponseSysCap = "NMC"
	StatusTupleResponseSysCapPmc StatusTupleResponseSysCap = "PMC"
	StatusTupleResponseSysCapUnk StatusTupleResponseSysCap = "UNK"
)

type StatusNewParams struct {
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
	DataMode StatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Comments describing the status creation and or updates to an entity.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusNewParamsOpsCap `json:"opsCap,omitzero"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusNewParamsState                 `json:"state,omitzero"`
	SubStatusCollection []StatusNewParamsSubStatusCollection `json:"subStatusCollection,omitzero"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusNewParamsSysCap `json:"sysCap,omitzero"`
	paramObj
}

func (r StatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StatusNewParams
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
type StatusNewParamsDataMode string

const (
	StatusNewParamsDataModeReal      StatusNewParamsDataMode = "REAL"
	StatusNewParamsDataModeTest      StatusNewParamsDataMode = "TEST"
	StatusNewParamsDataModeSimulated StatusNewParamsDataMode = "SIMULATED"
	StatusNewParamsDataModeExercise  StatusNewParamsDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusNewParamsOpsCap string

const (
	StatusNewParamsOpsCapFmc StatusNewParamsOpsCap = "FMC"
	StatusNewParamsOpsCapNmc StatusNewParamsOpsCap = "NMC"
	StatusNewParamsOpsCapPmc StatusNewParamsOpsCap = "PMC"
	StatusNewParamsOpsCapUnk StatusNewParamsOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusNewParamsState string

const (
	StatusNewParamsStateUnknown  StatusNewParamsState = "UNKNOWN"
	StatusNewParamsStateDead     StatusNewParamsState = "DEAD"
	StatusNewParamsStateActive   StatusNewParamsState = "ACTIVE"
	StatusNewParamsStateRfActive StatusNewParamsState = "RF ACTIVE"
	StatusNewParamsStateStandby  StatusNewParamsState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
//
// The properties ClassificationMarking, DataMode, Notes, Source, Status, StatusID,
// Type are required.
type StatusNewParamsSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,omitzero,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	paramObj
}

func (r StatusNewParamsSubStatusCollection) MarshalJSON() (data []byte, err error) {
	type shadow StatusNewParamsSubStatusCollection
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[StatusNewParamsSubStatusCollection](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[StatusNewParamsSubStatusCollection](
		"Status", false, "FMC", "NMC", "PMC", "UNK",
	)
	apijson.RegisterFieldValidator[StatusNewParamsSubStatusCollection](
		"Type", false, "mwCap", "ssCap", "mdCap",
	)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusNewParamsSysCap string

const (
	StatusNewParamsSysCapFmc StatusNewParamsSysCap = "FMC"
	StatusNewParamsSysCapNmc StatusNewParamsSysCap = "NMC"
	StatusNewParamsSysCapPmc StatusNewParamsSysCap = "PMC"
	StatusNewParamsSysCapUnk StatusNewParamsSysCap = "UNK"
)

type StatusUpdateParams struct {
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
	DataMode StatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent entity.
	IDEntity string `json:"idEntity,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The declassification date of this data, in ISO 8601 UTC format.
	DeclassificationDate param.Opt[time.Time] `json:"declassificationDate,omitzero" format:"date-time"`
	// Declassification string of this data.
	DeclassificationString param.Opt[string] `json:"declassificationString,omitzero"`
	// The sources or SCG references from which the classification of this data is
	// derived.
	DerivedFrom param.Opt[string] `json:"derivedFrom,omitzero"`
	// Comments describing the status creation and or updates to an entity.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	OpsCap StatusUpdateParamsOpsCap `json:"opsCap,omitzero"`
	// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
	// ACTIVE, STANDBY).
	//
	// Any of "UNKNOWN", "DEAD", "ACTIVE", "RF ACTIVE", "STANDBY".
	State               StatusUpdateParamsState                 `json:"state,omitzero"`
	SubStatusCollection []StatusUpdateParamsSubStatusCollection `json:"subStatusCollection,omitzero"`
	// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	SysCap StatusUpdateParamsSysCap `json:"sysCap,omitzero"`
	paramObj
}

func (r StatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StatusUpdateParams
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
type StatusUpdateParamsDataMode string

const (
	StatusUpdateParamsDataModeReal      StatusUpdateParamsDataMode = "REAL"
	StatusUpdateParamsDataModeTest      StatusUpdateParamsDataMode = "TEST"
	StatusUpdateParamsDataModeSimulated StatusUpdateParamsDataMode = "SIMULATED"
	StatusUpdateParamsDataModeExercise  StatusUpdateParamsDataMode = "EXERCISE"
)

// Operation capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusUpdateParamsOpsCap string

const (
	StatusUpdateParamsOpsCapFmc StatusUpdateParamsOpsCap = "FMC"
	StatusUpdateParamsOpsCapNmc StatusUpdateParamsOpsCap = "NMC"
	StatusUpdateParamsOpsCapPmc StatusUpdateParamsOpsCap = "PMC"
	StatusUpdateParamsOpsCapUnk StatusUpdateParamsOpsCap = "UNK"
)

// Overall state of the entity, if applicable (e.g. UNKNOWN, DEAD, ACTIVE, RF
// ACTIVE, STANDBY).
type StatusUpdateParamsState string

const (
	StatusUpdateParamsStateUnknown  StatusUpdateParamsState = "UNKNOWN"
	StatusUpdateParamsStateDead     StatusUpdateParamsState = "DEAD"
	StatusUpdateParamsStateActive   StatusUpdateParamsState = "ACTIVE"
	StatusUpdateParamsStateRfActive StatusUpdateParamsState = "RF ACTIVE"
	StatusUpdateParamsStateStandby  StatusUpdateParamsState = "STANDBY"
)

// Additional sub-system or capability status for the parent entity.
//
// The properties ClassificationMarking, DataMode, Notes, Source, Status, StatusID,
// Type are required.
type StatusUpdateParamsSubStatusCollection struct {
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
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status string `json:"status,omitzero,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type string `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	paramObj
}

func (r StatusUpdateParamsSubStatusCollection) MarshalJSON() (data []byte, err error) {
	type shadow StatusUpdateParamsSubStatusCollection
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[StatusUpdateParamsSubStatusCollection](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
	apijson.RegisterFieldValidator[StatusUpdateParamsSubStatusCollection](
		"Status", false, "FMC", "NMC", "PMC", "UNK",
	)
	apijson.RegisterFieldValidator[StatusUpdateParamsSubStatusCollection](
		"Type", false, "mwCap", "ssCap", "mdCap",
	)
}

// System capability of the entity, if applicable (e.g. FMC, NMC, PMC, UNK).
type StatusUpdateParamsSysCap string

const (
	StatusUpdateParamsSysCapFmc StatusUpdateParamsSysCap = "FMC"
	StatusUpdateParamsSysCapNmc StatusUpdateParamsSysCap = "NMC"
	StatusUpdateParamsSysCapPmc StatusUpdateParamsSysCap = "PMC"
	StatusUpdateParamsSysCapUnk StatusUpdateParamsSysCap = "UNK"
)

type StatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusListParams]'s query parameters as `url.Values`.
func (r StatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusCountParams]'s query parameters as `url.Values`.
func (r StatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusGetParams]'s query parameters as `url.Values`.
func (r StatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatusGetByEntityIDParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusGetByEntityIDParams]'s query parameters as
// `url.Values`.
func (r StatusGetByEntityIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatusGetByEntityTypeParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusGetByEntityTypeParams]'s query parameters as
// `url.Values`.
func (r StatusGetByEntityTypeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatusTupleParams]'s query parameters as `url.Values`.
func (r StatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
