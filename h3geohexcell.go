// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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

// H3GeoHexCellService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewH3GeoHexCellService] method instead.
type H3GeoHexCellService struct {
	Options []option.RequestOption
}

// NewH3GeoHexCellService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewH3GeoHexCellService(opts ...option.RequestOption) (r H3GeoHexCellService) {
	r = H3GeoHexCellService{}
	r.Options = opts
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *H3GeoHexCellService) List(ctx context.Context, query H3GeoHexCellListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[H3GeoHexCellListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/h3geohexcell"
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
func (r *H3GeoHexCellService) ListAutoPaging(ctx context.Context, query H3GeoHexCellListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[H3GeoHexCellListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *H3GeoHexCellService) Count(ctx context.Context, query H3GeoHexCellCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/h3geohexcell/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *H3GeoHexCellService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/h3geohexcell/queryhelp"
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
func (r *H3GeoHexCellService) Tuple(ctx context.Context, query H3GeoHexCellTupleParams, opts ...option.RequestOption) (res *[]H3GeoHexCellTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/h3geohexcell/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model representation of a hex cell array containing data for a set of
// observations.
type H3GeoHexCellListResponse struct {
	// The H3 index represented as a 16 character hexadecimal string.
	CellID string `json:"cellId,required"`
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
	DataMode H3GeoHexCellListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The mean altitude of the set of observations within this cell, measured in
	// kilometers.
	AltMean float64 `json:"altMean"`
	// The standard deviation of alttitude in the set of observations within this cell,
	// measured in kilometers.
	AltSigma float64 `json:"altSigma"`
	// The anomaly score for probable manufactured interference or RF interference;
	// calculated as a ratio of #anomalous obs / #total obs or coverage.
	AnomScoreInterference float64 `json:"anomScoreInterference"`
	// The anomaly score for probable spoofing; calculated as a ratio of #anomalous obs
	// / #total obs or coverage.
	AnomScoreSpoofing float64 `json:"anomScoreSpoofing"`
	// The percentage degree of change in the aggregated observables for a particular
	// H3 bin.
	ChangeScore float64 `json:"changeScore"`
	// The total number of available observations in the H3 cell during the start/end
	// times.
	Coverage int64 `json:"coverage"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent H3 Geo record containing this hex cell.
	IDH3Geo string `json:"idH3Geo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The max received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMax float64 `json:"rpmMax"`
	// The mean received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMean float64 `json:"rpmMean"`
	// The median received power monitor (RPM) output value for the set of data
	// contained within this cell.
	RpmMedian float64 `json:"rpmMedian"`
	// The min received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMin float64 `json:"rpmMin"`
	// The standard deviation of the received power monitor (RPM) output value for the
	// set of data contained within this cell.
	RpmSigma float64 `json:"rpmSigma"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CellID                resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		ID                    resp.Field
		AltMean               resp.Field
		AltSigma              resp.Field
		AnomScoreInterference resp.Field
		AnomScoreSpoofing     resp.Field
		ChangeScore           resp.Field
		Coverage              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		IDH3Geo               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RpmMax                resp.Field
		RpmMean               resp.Field
		RpmMedian             resp.Field
		RpmMin                resp.Field
		RpmSigma              resp.Field
		SourceDl              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoHexCellListResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoHexCellListResponse) UnmarshalJSON(data []byte) error {
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
type H3GeoHexCellListResponseDataMode string

const (
	H3GeoHexCellListResponseDataModeReal      H3GeoHexCellListResponseDataMode = "REAL"
	H3GeoHexCellListResponseDataModeTest      H3GeoHexCellListResponseDataMode = "TEST"
	H3GeoHexCellListResponseDataModeSimulated H3GeoHexCellListResponseDataMode = "SIMULATED"
	H3GeoHexCellListResponseDataModeExercise  H3GeoHexCellListResponseDataMode = "EXERCISE"
)

// Model representation of a hex cell array containing data for a set of
// observations.
type H3GeoHexCellTupleResponse struct {
	// The H3 index represented as a 16 character hexadecimal string.
	CellID string `json:"cellId,required"`
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
	DataMode H3GeoHexCellTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The mean altitude of the set of observations within this cell, measured in
	// kilometers.
	AltMean float64 `json:"altMean"`
	// The standard deviation of alttitude in the set of observations within this cell,
	// measured in kilometers.
	AltSigma float64 `json:"altSigma"`
	// The anomaly score for probable manufactured interference or RF interference;
	// calculated as a ratio of #anomalous obs / #total obs or coverage.
	AnomScoreInterference float64 `json:"anomScoreInterference"`
	// The anomaly score for probable spoofing; calculated as a ratio of #anomalous obs
	// / #total obs or coverage.
	AnomScoreSpoofing float64 `json:"anomScoreSpoofing"`
	// The percentage degree of change in the aggregated observables for a particular
	// H3 bin.
	ChangeScore float64 `json:"changeScore"`
	// The total number of available observations in the H3 cell during the start/end
	// times.
	Coverage int64 `json:"coverage"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Unique identifier of the parent H3 Geo record containing this hex cell.
	IDH3Geo string `json:"idH3Geo"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The max received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMax float64 `json:"rpmMax"`
	// The mean received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMean float64 `json:"rpmMean"`
	// The median received power monitor (RPM) output value for the set of data
	// contained within this cell.
	RpmMedian float64 `json:"rpmMedian"`
	// The min received power monitor (RPM) output value for the set of data contained
	// within this cell.
	RpmMin float64 `json:"rpmMin"`
	// The standard deviation of the received power monitor (RPM) output value for the
	// set of data contained within this cell.
	RpmSigma float64 `json:"rpmSigma"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CellID                resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		Source                resp.Field
		ID                    resp.Field
		AltMean               resp.Field
		AltSigma              resp.Field
		AnomScoreInterference resp.Field
		AnomScoreSpoofing     resp.Field
		ChangeScore           resp.Field
		Coverage              resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		IDH3Geo               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		RpmMax                resp.Field
		RpmMean               resp.Field
		RpmMedian             resp.Field
		RpmMin                resp.Field
		RpmSigma              resp.Field
		SourceDl              resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r H3GeoHexCellTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *H3GeoHexCellTupleResponse) UnmarshalJSON(data []byte) error {
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
type H3GeoHexCellTupleResponseDataMode string

const (
	H3GeoHexCellTupleResponseDataModeReal      H3GeoHexCellTupleResponseDataMode = "REAL"
	H3GeoHexCellTupleResponseDataModeTest      H3GeoHexCellTupleResponseDataMode = "TEST"
	H3GeoHexCellTupleResponseDataModeSimulated H3GeoHexCellTupleResponseDataMode = "SIMULATED"
	H3GeoHexCellTupleResponseDataModeExercise  H3GeoHexCellTupleResponseDataMode = "EXERCISE"
)

type H3GeoHexCellListParams struct {
	// Unique identifier of the parent H3 Geo record containing this hex cell. (uuid)
	IDH3Geo     string           `query:"idH3Geo,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3GeoHexCellListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3GeoHexCellListParams]'s query parameters as `url.Values`.
func (r H3GeoHexCellListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3GeoHexCellCountParams struct {
	// Unique identifier of the parent H3 Geo record containing this hex cell. (uuid)
	IDH3Geo     string           `query:"idH3Geo,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3GeoHexCellCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3GeoHexCellCountParams]'s query parameters as
// `url.Values`.
func (r H3GeoHexCellCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type H3GeoHexCellTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Unique identifier of the parent H3 Geo record containing this hex cell. (uuid)
	IDH3Geo     string           `query:"idH3Geo,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f H3GeoHexCellTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [H3GeoHexCellTupleParams]'s query parameters as
// `url.Values`.
func (r H3GeoHexCellTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
