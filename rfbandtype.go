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

// RfBandTypeService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRfBandTypeService] method instead.
type RfBandTypeService struct {
	Options []option.RequestOption
}

// NewRfBandTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRfBandTypeService(opts ...option.RequestOption) (r RfBandTypeService) {
	r = RfBandTypeService{}
	r.Options = opts
	return
}

// Service operation to take a single RFBandType as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfBandTypeService) New(ctx context.Context, body RfBandTypeNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/rfbandtype"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an RFBandType. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *RfBandTypeService) Update(ctx context.Context, id string, body RfBandTypeUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfbandtype/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *RfBandTypeService) List(ctx context.Context, query RfBandTypeListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RfBandTypeListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/rfbandtype"
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
func (r *RfBandTypeService) ListAutoPaging(ctx context.Context, query RfBandTypeListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RfBandTypeListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an RFBandType specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *RfBandTypeService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfbandtype/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *RfBandTypeService) Count(ctx context.Context, query RfBandTypeCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/rfbandtype/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single RFBandType by its unique ID passed as a path
// parameter.
func (r *RfBandTypeService) Get(ctx context.Context, id string, query RfBandTypeGetParams, opts ...option.RequestOption) (res *RfBandTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/rfbandtype/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *RfBandTypeService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *RfBandTypeQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/rfbandtype/queryhelp"
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
func (r *RfBandTypeService) Tuple(ctx context.Context, query RfBandTypeTupleParams, opts ...option.RequestOption) (res *[]RfBandTypeTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/rfbandtype/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This table contains descriptions for common satellite RF bands.
type RfBandTypeListResponse struct {
	// Unique identifier for the RF band (e.g. X, K, Ku, etc).
	ID string `json:"id,required"`
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
	DataMode RfBandTypeListResponseDataMode `json:"dataMode,required"`
	// Description of the band and common uses.
	Description string `json:"description,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Approximate end of the band frequency range, in Ghz.
	EndFreq float64 `json:"endFreq"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Approximate start of the band frequency range, in Ghz.
	StartFreq float64 `json:"startFreq"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Source                respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndFreq               respjson.Field
		Origin                respjson.Field
		StartFreq             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandTypeListResponse) UnmarshalJSON(data []byte) error {
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
type RfBandTypeListResponseDataMode string

const (
	RfBandTypeListResponseDataModeReal      RfBandTypeListResponseDataMode = "REAL"
	RfBandTypeListResponseDataModeTest      RfBandTypeListResponseDataMode = "TEST"
	RfBandTypeListResponseDataModeSimulated RfBandTypeListResponseDataMode = "SIMULATED"
	RfBandTypeListResponseDataModeExercise  RfBandTypeListResponseDataMode = "EXERCISE"
)

// This table contains descriptions for common satellite RF bands.
type RfBandTypeGetResponse struct {
	// Unique identifier for the RF band (e.g. X, K, Ku, etc).
	ID string `json:"id,required"`
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
	DataMode RfBandTypeGetResponseDataMode `json:"dataMode,required"`
	// Description of the band and common uses.
	Description string `json:"description,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Approximate end of the band frequency range, in Ghz.
	EndFreq float64 `json:"endFreq"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Approximate start of the band frequency range, in Ghz.
	StartFreq float64 `json:"startFreq"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Source                respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndFreq               respjson.Field
		Origin                respjson.Field
		StartFreq             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandTypeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandTypeGetResponse) UnmarshalJSON(data []byte) error {
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
type RfBandTypeGetResponseDataMode string

const (
	RfBandTypeGetResponseDataModeReal      RfBandTypeGetResponseDataMode = "REAL"
	RfBandTypeGetResponseDataModeTest      RfBandTypeGetResponseDataMode = "TEST"
	RfBandTypeGetResponseDataModeSimulated RfBandTypeGetResponseDataMode = "SIMULATED"
	RfBandTypeGetResponseDataModeExercise  RfBandTypeGetResponseDataMode = "EXERCISE"
)

type RfBandTypeQueryhelpResponse struct {
	AodrSupported         bool                                   `json:"aodrSupported"`
	ClassificationMarking string                                 `json:"classificationMarking"`
	Description           string                                 `json:"description"`
	HistorySupported      bool                                   `json:"historySupported"`
	Name                  string                                 `json:"name"`
	Parameters            []RfBandTypeQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                               `json:"requiredRoles"`
	RestSupported         bool                                   `json:"restSupported"`
	SortSupported         bool                                   `json:"sortSupported"`
	TypeName              string                                 `json:"typeName"`
	Uri                   string                                 `json:"uri"`
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
func (r RfBandTypeQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandTypeQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RfBandTypeQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandTypeQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *RfBandTypeQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This table contains descriptions for common satellite RF bands.
type RfBandTypeTupleResponse struct {
	// Unique identifier for the RF band (e.g. X, K, Ku, etc).
	ID string `json:"id,required"`
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
	DataMode RfBandTypeTupleResponseDataMode `json:"dataMode,required"`
	// Description of the band and common uses.
	Description string `json:"description,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Approximate end of the band frequency range, in Ghz.
	EndFreq float64 `json:"endFreq"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// Approximate start of the band frequency range, in Ghz.
	StartFreq float64 `json:"startFreq"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Description           respjson.Field
		Source                respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		EndFreq               respjson.Field
		Origin                respjson.Field
		StartFreq             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RfBandTypeTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *RfBandTypeTupleResponse) UnmarshalJSON(data []byte) error {
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
type RfBandTypeTupleResponseDataMode string

const (
	RfBandTypeTupleResponseDataModeReal      RfBandTypeTupleResponseDataMode = "REAL"
	RfBandTypeTupleResponseDataModeTest      RfBandTypeTupleResponseDataMode = "TEST"
	RfBandTypeTupleResponseDataModeSimulated RfBandTypeTupleResponseDataMode = "SIMULATED"
	RfBandTypeTupleResponseDataModeExercise  RfBandTypeTupleResponseDataMode = "EXERCISE"
)

type RfBandTypeNewParams struct {
	// Unique identifier for the RF band (e.g. X, K, Ku, etc).
	ID string `json:"id,required"`
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
	DataMode RfBandTypeNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Description of the band and common uses.
	Description string `json:"description,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Approximate end of the band frequency range, in Ghz.
	EndFreq param.Opt[float64] `json:"endFreq,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Approximate start of the band frequency range, in Ghz.
	StartFreq param.Opt[float64] `json:"startFreq,omitzero"`
	paramObj
}

func (r RfBandTypeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandTypeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfBandTypeNewParams) UnmarshalJSON(data []byte) error {
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
type RfBandTypeNewParamsDataMode string

const (
	RfBandTypeNewParamsDataModeReal      RfBandTypeNewParamsDataMode = "REAL"
	RfBandTypeNewParamsDataModeTest      RfBandTypeNewParamsDataMode = "TEST"
	RfBandTypeNewParamsDataModeSimulated RfBandTypeNewParamsDataMode = "SIMULATED"
	RfBandTypeNewParamsDataModeExercise  RfBandTypeNewParamsDataMode = "EXERCISE"
)

type RfBandTypeUpdateParams struct {
	// Unique identifier for the RF band (e.g. X, K, Ku, etc).
	ID string `json:"id,required"`
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
	DataMode RfBandTypeUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Description of the band and common uses.
	Description string `json:"description,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Approximate end of the band frequency range, in Ghz.
	EndFreq param.Opt[float64] `json:"endFreq,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Approximate start of the band frequency range, in Ghz.
	StartFreq param.Opt[float64] `json:"startFreq,omitzero"`
	paramObj
}

func (r RfBandTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RfBandTypeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RfBandTypeUpdateParams) UnmarshalJSON(data []byte) error {
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
type RfBandTypeUpdateParamsDataMode string

const (
	RfBandTypeUpdateParamsDataModeReal      RfBandTypeUpdateParamsDataMode = "REAL"
	RfBandTypeUpdateParamsDataModeTest      RfBandTypeUpdateParamsDataMode = "TEST"
	RfBandTypeUpdateParamsDataModeSimulated RfBandTypeUpdateParamsDataMode = "SIMULATED"
	RfBandTypeUpdateParamsDataModeExercise  RfBandTypeUpdateParamsDataMode = "EXERCISE"
)

type RfBandTypeListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandTypeListParams]'s query parameters as `url.Values`.
func (r RfBandTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandTypeCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandTypeCountParams]'s query parameters as `url.Values`.
func (r RfBandTypeCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandTypeGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandTypeGetParams]'s query parameters as `url.Values`.
func (r RfBandTypeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RfBandTypeTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RfBandTypeTupleParams]'s query parameters as `url.Values`.
func (r RfBandTypeTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
