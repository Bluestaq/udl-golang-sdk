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

// SubstatusService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubstatusService] method instead.
type SubstatusService struct {
	Options []option.RequestOption
}

// NewSubstatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubstatusService(opts ...option.RequestOption) (r SubstatusService) {
	r = SubstatusService{}
	r.Options = opts
	return
}

// Service operation to take a single Sub Status record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SubstatusService) New(ctx context.Context, body SubstatusNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/substatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Sub Status record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SubstatusService) Update(ctx context.Context, id string, body SubstatusUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/substatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SubstatusService) List(ctx context.Context, query SubstatusListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SubstatusListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/substatus"
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
func (r *SubstatusService) ListAutoPaging(ctx context.Context, query SubstatusListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SubstatusListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Sub Status record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *SubstatusService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/substatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SubstatusService) Count(ctx context.Context, query SubstatusCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/substatus/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Sub Status record by its unique ID passed as a
// path parameter.
func (r *SubstatusService) Get(ctx context.Context, id string, query SubstatusGetParams, opts ...option.RequestOption) (res *SubstatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/substatus/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SubstatusService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/substatus/queryhelp"
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
func (r *SubstatusService) Tuple(ctx context.Context, query SubstatusTupleParams, opts ...option.RequestOption) (res *[]SubstatusTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/substatus/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Additional sub-system or capability status for the parent entity.
type SubstatusListResponse struct {
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
	DataMode SubstatusListResponseDataMode `json:"dataMode,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status SubstatusListResponseStatus `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type SubstatusListResponseType `json:"type,required"`
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
func (r SubstatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubstatusListResponse) UnmarshalJSON(data []byte) error {
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
type SubstatusListResponseDataMode string

const (
	SubstatusListResponseDataModeReal      SubstatusListResponseDataMode = "REAL"
	SubstatusListResponseDataModeTest      SubstatusListResponseDataMode = "TEST"
	SubstatusListResponseDataModeSimulated SubstatusListResponseDataMode = "SIMULATED"
	SubstatusListResponseDataModeExercise  SubstatusListResponseDataMode = "EXERCISE"
)

// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
type SubstatusListResponseStatus string

const (
	SubstatusListResponseStatusFmc SubstatusListResponseStatus = "FMC"
	SubstatusListResponseStatusNmc SubstatusListResponseStatus = "NMC"
	SubstatusListResponseStatusPmc SubstatusListResponseStatus = "PMC"
	SubstatusListResponseStatusUnk SubstatusListResponseStatus = "UNK"
)

// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
type SubstatusListResponseType string

const (
	SubstatusListResponseTypeMwCap SubstatusListResponseType = "mwCap"
	SubstatusListResponseTypeSSCap SubstatusListResponseType = "ssCap"
	SubstatusListResponseTypeMdCap SubstatusListResponseType = "mdCap"
)

// Additional sub-system or capability status for the parent entity.
type SubstatusGetResponse struct {
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
	DataMode SubstatusGetResponseDataMode `json:"dataMode,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status SubstatusGetResponseStatus `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type SubstatusGetResponseType `json:"type,required"`
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
func (r SubstatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SubstatusGetResponse) UnmarshalJSON(data []byte) error {
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
type SubstatusGetResponseDataMode string

const (
	SubstatusGetResponseDataModeReal      SubstatusGetResponseDataMode = "REAL"
	SubstatusGetResponseDataModeTest      SubstatusGetResponseDataMode = "TEST"
	SubstatusGetResponseDataModeSimulated SubstatusGetResponseDataMode = "SIMULATED"
	SubstatusGetResponseDataModeExercise  SubstatusGetResponseDataMode = "EXERCISE"
)

// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
type SubstatusGetResponseStatus string

const (
	SubstatusGetResponseStatusFmc SubstatusGetResponseStatus = "FMC"
	SubstatusGetResponseStatusNmc SubstatusGetResponseStatus = "NMC"
	SubstatusGetResponseStatusPmc SubstatusGetResponseStatus = "PMC"
	SubstatusGetResponseStatusUnk SubstatusGetResponseStatus = "UNK"
)

// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
type SubstatusGetResponseType string

const (
	SubstatusGetResponseTypeMwCap SubstatusGetResponseType = "mwCap"
	SubstatusGetResponseTypeSSCap SubstatusGetResponseType = "ssCap"
	SubstatusGetResponseTypeMdCap SubstatusGetResponseType = "mdCap"
)

// Additional sub-system or capability status for the parent entity.
type SubstatusTupleResponse struct {
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
	DataMode SubstatusTupleResponseDataMode `json:"dataMode,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status SubstatusTupleResponseStatus `json:"status,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type SubstatusTupleResponseType `json:"type,required"`
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
func (r SubstatusTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SubstatusTupleResponse) UnmarshalJSON(data []byte) error {
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
type SubstatusTupleResponseDataMode string

const (
	SubstatusTupleResponseDataModeReal      SubstatusTupleResponseDataMode = "REAL"
	SubstatusTupleResponseDataModeTest      SubstatusTupleResponseDataMode = "TEST"
	SubstatusTupleResponseDataModeSimulated SubstatusTupleResponseDataMode = "SIMULATED"
	SubstatusTupleResponseDataModeExercise  SubstatusTupleResponseDataMode = "EXERCISE"
)

// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
type SubstatusTupleResponseStatus string

const (
	SubstatusTupleResponseStatusFmc SubstatusTupleResponseStatus = "FMC"
	SubstatusTupleResponseStatusNmc SubstatusTupleResponseStatus = "NMC"
	SubstatusTupleResponseStatusPmc SubstatusTupleResponseStatus = "PMC"
	SubstatusTupleResponseStatusUnk SubstatusTupleResponseStatus = "UNK"
)

// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
type SubstatusTupleResponseType string

const (
	SubstatusTupleResponseTypeMwCap SubstatusTupleResponseType = "mwCap"
	SubstatusTupleResponseTypeSSCap SubstatusTupleResponseType = "ssCap"
	SubstatusTupleResponseTypeMdCap SubstatusTupleResponseType = "mdCap"
)

type SubstatusNewParams struct {
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
	DataMode SubstatusNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status SubstatusNewParamsStatus `json:"status,omitzero,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type SubstatusNewParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	paramObj
}

func (r SubstatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubstatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubstatusNewParams) UnmarshalJSON(data []byte) error {
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
type SubstatusNewParamsDataMode string

const (
	SubstatusNewParamsDataModeReal      SubstatusNewParamsDataMode = "REAL"
	SubstatusNewParamsDataModeTest      SubstatusNewParamsDataMode = "TEST"
	SubstatusNewParamsDataModeSimulated SubstatusNewParamsDataMode = "SIMULATED"
	SubstatusNewParamsDataModeExercise  SubstatusNewParamsDataMode = "EXERCISE"
)

// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
type SubstatusNewParamsStatus string

const (
	SubstatusNewParamsStatusFmc SubstatusNewParamsStatus = "FMC"
	SubstatusNewParamsStatusNmc SubstatusNewParamsStatus = "NMC"
	SubstatusNewParamsStatusPmc SubstatusNewParamsStatus = "PMC"
	SubstatusNewParamsStatusUnk SubstatusNewParamsStatus = "UNK"
)

// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
type SubstatusNewParamsType string

const (
	SubstatusNewParamsTypeMwCap SubstatusNewParamsType = "mwCap"
	SubstatusNewParamsTypeSSCap SubstatusNewParamsType = "ssCap"
	SubstatusNewParamsTypeMdCap SubstatusNewParamsType = "mdCap"
)

type SubstatusUpdateParams struct {
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
	DataMode SubstatusUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Descriptions and/or comments associated with the sub-status.
	Notes string `json:"notes,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
	//
	// Any of "FMC", "NMC", "PMC", "UNK".
	Status SubstatusUpdateParamsStatus `json:"status,omitzero,required"`
	// Id of the parent status.
	StatusID string `json:"statusId,required"`
	// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
	//
	// Any of "mwCap", "ssCap", "mdCap".
	Type SubstatusUpdateParamsType `json:"type,omitzero,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	paramObj
}

func (r SubstatusUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubstatusUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubstatusUpdateParams) UnmarshalJSON(data []byte) error {
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
type SubstatusUpdateParamsDataMode string

const (
	SubstatusUpdateParamsDataModeReal      SubstatusUpdateParamsDataMode = "REAL"
	SubstatusUpdateParamsDataModeTest      SubstatusUpdateParamsDataMode = "TEST"
	SubstatusUpdateParamsDataModeSimulated SubstatusUpdateParamsDataMode = "SIMULATED"
	SubstatusUpdateParamsDataModeExercise  SubstatusUpdateParamsDataMode = "EXERCISE"
)

// Status of the sub-system/capability, e.g. FMC, NMC, PMC, UNK.
type SubstatusUpdateParamsStatus string

const (
	SubstatusUpdateParamsStatusFmc SubstatusUpdateParamsStatus = "FMC"
	SubstatusUpdateParamsStatusNmc SubstatusUpdateParamsStatus = "NMC"
	SubstatusUpdateParamsStatusPmc SubstatusUpdateParamsStatus = "PMC"
	SubstatusUpdateParamsStatusUnk SubstatusUpdateParamsStatus = "UNK"
)

// Parent entity's sub-system or capability status: mwCap, mdCap, ssCap, etc.
type SubstatusUpdateParamsType string

const (
	SubstatusUpdateParamsTypeMwCap SubstatusUpdateParamsType = "mwCap"
	SubstatusUpdateParamsTypeSSCap SubstatusUpdateParamsType = "ssCap"
	SubstatusUpdateParamsTypeMdCap SubstatusUpdateParamsType = "mdCap"
)

type SubstatusListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubstatusListParams]'s query parameters as `url.Values`.
func (r SubstatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubstatusCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubstatusCountParams]'s query parameters as `url.Values`.
func (r SubstatusCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubstatusGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubstatusGetParams]'s query parameters as `url.Values`.
func (r SubstatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubstatusTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubstatusTupleParams]'s query parameters as `url.Values`.
func (r SubstatusTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
