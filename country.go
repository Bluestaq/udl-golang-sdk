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

// CountryService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCountryService] method instead.
type CountryService struct {
	Options []option.RequestOption
}

// NewCountryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCountryService(opts ...option.RequestOption) (r CountryService) {
	r = CountryService{}
	r.Options = opts
	return
}

// Service operation to take a single Country as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *CountryService) New(ctx context.Context, body CountryNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/country"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Country record by its unique code passed as a
// path parameter.
func (r *CountryService) Get(ctx context.Context, code string, query CountryGetParams, opts ...option.RequestOption) (res *CountryFull, err error) {
	opts = append(r.Options[:], opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("udl/country/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Country. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *CountryService) Update(ctx context.Context, code string, body CountryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("udl/country/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *CountryService) List(ctx context.Context, query CountryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CountryAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/country"
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
func (r *CountryService) ListAutoPaging(ctx context.Context, query CountryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CountryAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Country object specified by the passed code path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *CountryService) Delete(ctx context.Context, code string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("udl/country/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *CountryService) Count(ctx context.Context, query CountryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/country/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *CountryService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/country/queryhelp"
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
func (r *CountryService) Tuple(ctx context.Context, query CountryTupleParams, opts ...option.RequestOption) (res *[]CountryFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/country/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A Country may represent countries, multi-national consortiums, and international
// organizations.
type CountryAbridged struct {
	// The country code. Optimally, this value is the ISO 3166 Alpha-2-two-character
	// country code, however it can represent various consortiums that do not appear in
	// the ISO document.
	Code string `json:"code,required"`
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
	DataMode CountryAbridgedDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// 3 Digit or other alternate country code.
	CodeAlt string `json:"codeAlt"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Federal Information Processing Standard (FIPS) two-character country code. This
	// field is used when the country code for FIPS differs from the country code for
	// ISO-3166 value. For example, the ISO-3166 Alpha-2-country code for Vanuatu is
	// VU, whereas Vanuatu's FIPS equivalent country code is NH.
	FipsCode string `json:"fipsCode"`
	// ISO 3166 Alpha-3 country code. This is a three-character code that represents a
	// country name, which may be more closely related to the country name than its
	// corresponding Alpha-2 code.
	Iso3Code string `json:"iso3Code"`
	// Country name.
	Name string `json:"name"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Code        resp.Field
		DataMode    resp.Field
		Source      resp.Field
		CodeAlt     resp.Field
		CreatedAt   resp.Field
		CreatedBy   resp.Field
		FipsCode    resp.Field
		Iso3Code    resp.Field
		Name        resp.Field
		OrigNetwork resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryAbridged) RawJSON() string { return r.JSON.raw }
func (r *CountryAbridged) UnmarshalJSON(data []byte) error {
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
type CountryAbridgedDataMode string

const (
	CountryAbridgedDataModeReal      CountryAbridgedDataMode = "REAL"
	CountryAbridgedDataModeTest      CountryAbridgedDataMode = "TEST"
	CountryAbridgedDataModeSimulated CountryAbridgedDataMode = "SIMULATED"
	CountryAbridgedDataModeExercise  CountryAbridgedDataMode = "EXERCISE"
)

// A Country may represent countries, multi-national consortiums, and international
// organizations.
type CountryFull struct {
	// The country code. Optimally, this value is the ISO 3166 Alpha-2-two-character
	// country code, however it can represent various consortiums that do not appear in
	// the ISO document.
	Code string `json:"code,required"`
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
	DataMode CountryFullDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// 3 Digit or other alternate country code.
	CodeAlt string `json:"codeAlt"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Federal Information Processing Standard (FIPS) two-character country code. This
	// field is used when the country code for FIPS differs from the country code for
	// ISO-3166 value. For example, the ISO-3166 Alpha-2-country code for Vanuatu is
	// VU, whereas Vanuatu's FIPS equivalent country code is NH.
	FipsCode string `json:"fipsCode"`
	// ISO 3166 Alpha-3 country code. This is a three-character code that represents a
	// country name, which may be more closely related to the country name than its
	// corresponding Alpha-2 code.
	Iso3Code string `json:"iso3Code"`
	// Country name.
	Name string `json:"name"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Code        resp.Field
		DataMode    resp.Field
		Source      resp.Field
		CodeAlt     resp.Field
		CreatedAt   resp.Field
		CreatedBy   resp.Field
		FipsCode    resp.Field
		Iso3Code    resp.Field
		Name        resp.Field
		OrigNetwork resp.Field
		UpdatedAt   resp.Field
		UpdatedBy   resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryFull) RawJSON() string { return r.JSON.raw }
func (r *CountryFull) UnmarshalJSON(data []byte) error {
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
type CountryFullDataMode string

const (
	CountryFullDataModeReal      CountryFullDataMode = "REAL"
	CountryFullDataModeTest      CountryFullDataMode = "TEST"
	CountryFullDataModeSimulated CountryFullDataMode = "SIMULATED"
	CountryFullDataModeExercise  CountryFullDataMode = "EXERCISE"
)

type CountryNewParams struct {
	// The country code. Optimally, this value is the ISO 3166 Alpha-2-two-character
	// country code, however it can represent various consortiums that do not appear in
	// the ISO document.
	Code string `json:"code,required"`
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
	DataMode CountryNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// 3 Digit or other alternate country code.
	CodeAlt param.Opt[string] `json:"codeAlt,omitzero"`
	// Federal Information Processing Standard (FIPS) two-character country code. This
	// field is used when the country code for FIPS differs from the country code for
	// ISO-3166 value. For example, the ISO-3166 Alpha-2-country code for Vanuatu is
	// VU, whereas Vanuatu's FIPS equivalent country code is NH.
	FipsCode param.Opt[string] `json:"fipsCode,omitzero"`
	// ISO 3166 Alpha-3 country code. This is a three-character code that represents a
	// country name, which may be more closely related to the country name than its
	// corresponding Alpha-2 code.
	Iso3Code param.Opt[string] `json:"iso3Code,omitzero"`
	// Country name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CountryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CountryNewParams
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
type CountryNewParamsDataMode string

const (
	CountryNewParamsDataModeReal      CountryNewParamsDataMode = "REAL"
	CountryNewParamsDataModeTest      CountryNewParamsDataMode = "TEST"
	CountryNewParamsDataModeSimulated CountryNewParamsDataMode = "SIMULATED"
	CountryNewParamsDataModeExercise  CountryNewParamsDataMode = "EXERCISE"
)

type CountryGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CountryGetParams]'s query parameters as `url.Values`.
func (r CountryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CountryUpdateParams struct {
	// The country code. Optimally, this value is the ISO 3166 Alpha-2-two-character
	// country code, however it can represent various consortiums that do not appear in
	// the ISO document.
	Code string `json:"code,required"`
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
	DataMode CountryUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// 3 Digit or other alternate country code.
	CodeAlt param.Opt[string] `json:"codeAlt,omitzero"`
	// Federal Information Processing Standard (FIPS) two-character country code. This
	// field is used when the country code for FIPS differs from the country code for
	// ISO-3166 value. For example, the ISO-3166 Alpha-2-country code for Vanuatu is
	// VU, whereas Vanuatu's FIPS equivalent country code is NH.
	FipsCode param.Opt[string] `json:"fipsCode,omitzero"`
	// ISO 3166 Alpha-3 country code. This is a three-character code that represents a
	// country name, which may be more closely related to the country name than its
	// corresponding Alpha-2 code.
	Iso3Code param.Opt[string] `json:"iso3Code,omitzero"`
	// Country name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CountryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CountryUpdateParams
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
type CountryUpdateParamsDataMode string

const (
	CountryUpdateParamsDataModeReal      CountryUpdateParamsDataMode = "REAL"
	CountryUpdateParamsDataModeTest      CountryUpdateParamsDataMode = "TEST"
	CountryUpdateParamsDataModeSimulated CountryUpdateParamsDataMode = "SIMULATED"
	CountryUpdateParamsDataModeExercise  CountryUpdateParamsDataMode = "EXERCISE"
)

type CountryListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CountryListParams]'s query parameters as `url.Values`.
func (r CountryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CountryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CountryCountParams]'s query parameters as `url.Values`.
func (r CountryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CountryTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CountryTupleParams]'s query parameters as `url.Values`.
func (r CountryTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
