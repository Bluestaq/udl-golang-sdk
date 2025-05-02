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

// OnorbitantennaService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitantennaService] method instead.
type OnorbitantennaService struct {
	Options []option.RequestOption
}

// NewOnorbitantennaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnorbitantennaService(opts ...option.RequestOption) (r OnorbitantennaService) {
	r = OnorbitantennaService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitAntenna as a POST body and ingest into
// the database. An OnorbitAntenna is the association between on-orbit spacecraft
// antennas and a particular on-orbit spacecraft. An antenna type may be associated
// with many different on-orbit spacecraft. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *OnorbitantennaService) New(ctx context.Context, body OnorbitantennaNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitantenna"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitAntenna. An OnorbitAntenna is the
// association between on-orbit spacecraft antennas and a particular on-orbit
// spacecraft. An antenna type may be associated with many different on-orbit
// spacecraft. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *OnorbitantennaService) Update(ctx context.Context, id string, body OnorbitantennaUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitantenna/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitantennaService) List(ctx context.Context, query OnorbitantennaListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitantennaListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitantenna"
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
func (r *OnorbitantennaService) ListAutoPaging(ctx context.Context, query OnorbitantennaListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitantennaListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitAntenna object specified by the passed ID
// path parameter. An OnorbitAntenna is the association between on-orbit spacecraft
// antennas and a particular on-orbit spacecraft. An antenna type may be associated
// with many different on-orbit spacecraft. A specific role is required to perform
// this service operation. Please contact the UDL team for assistance.
func (r *OnorbitantennaService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitantenna/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OnorbitAntenna record by its unique ID passed
// as a path parameter. An OnorbitAntenna is the association between on-orbit
// spacecraft antennas and a particular on-orbit spacecraft. An antenna type may be
// associated with many different on-orbit spacecraft.
func (r *OnorbitantennaService) Get(ctx context.Context, id string, query OnorbitantennaGetParams, opts ...option.RequestOption) (res *OnorbitantennaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitantenna/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitantennaListResponse struct {
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
	DataMode OnorbitantennaListResponseDataMode `json:"dataMode,required"`
	// ID of the antenna.
	IDAntenna string `json:"idAntenna,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model representation of information on on-orbit/spacecraft communication
	// antennas. A spacecraft may have multiple antennas and each antenna can have
	// multiple 'details' records compiled by different sources.
	Antenna AntennaAbridged `json:"antenna"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDAntenna             resp.Field
		IDOnOrbit             resp.Field
		Source                resp.Field
		ID                    resp.Field
		Antenna               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitantennaListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitantennaListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitantennaListResponseDataMode string

const (
	OnorbitantennaListResponseDataModeReal      OnorbitantennaListResponseDataMode = "REAL"
	OnorbitantennaListResponseDataModeTest      OnorbitantennaListResponseDataMode = "TEST"
	OnorbitantennaListResponseDataModeSimulated OnorbitantennaListResponseDataMode = "SIMULATED"
	OnorbitantennaListResponseDataModeExercise  OnorbitantennaListResponseDataMode = "EXERCISE"
)

type OnorbitantennaGetResponse struct {
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
	DataMode OnorbitantennaGetResponseDataMode `json:"dataMode,required"`
	// ID of the antenna.
	IDAntenna string `json:"idAntenna,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Model representation of information on on-orbit/spacecraft communication
	// antennas. A spacecraft may have multiple antennas and each antenna can have
	// multiple 'details' records compiled by different sources.
	Antenna AntennaFull `json:"antenna"`
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
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDAntenna             resp.Field
		IDOnOrbit             resp.Field
		Source                resp.Field
		ID                    resp.Field
		Antenna               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitantennaGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitantennaGetResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitantennaGetResponseDataMode string

const (
	OnorbitantennaGetResponseDataModeReal      OnorbitantennaGetResponseDataMode = "REAL"
	OnorbitantennaGetResponseDataModeTest      OnorbitantennaGetResponseDataMode = "TEST"
	OnorbitantennaGetResponseDataModeSimulated OnorbitantennaGetResponseDataMode = "SIMULATED"
	OnorbitantennaGetResponseDataModeExercise  OnorbitantennaGetResponseDataMode = "EXERCISE"
)

type OnorbitantennaNewParams struct {
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
	DataMode OnorbitantennaNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the antenna.
	IDAntenna string `json:"idAntenna,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Model representation of information on on-orbit/spacecraft communication
	// antennas. A spacecraft may have multiple antennas and each antenna can have
	// multiple 'details' records compiled by different sources.
	Antenna OnorbitantennaNewParamsAntenna `json:"antenna,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r OnorbitantennaNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitantennaNewParams
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
type OnorbitantennaNewParamsDataMode string

const (
	OnorbitantennaNewParamsDataModeReal      OnorbitantennaNewParamsDataMode = "REAL"
	OnorbitantennaNewParamsDataModeTest      OnorbitantennaNewParamsDataMode = "TEST"
	OnorbitantennaNewParamsDataModeSimulated OnorbitantennaNewParamsDataMode = "SIMULATED"
	OnorbitantennaNewParamsDataModeExercise  OnorbitantennaNewParamsDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft communication
// antennas. A spacecraft may have multiple antennas and each antenna can have
// multiple 'details' records compiled by different sources.
//
// The properties DataMode, Name, Source are required.
type OnorbitantennaNewParamsAntenna struct {
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
	// Antenna name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaNewParamsAntenna) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r OnorbitantennaNewParamsAntenna) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitantennaNewParamsAntenna
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnorbitantennaNewParamsAntenna](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitantennaUpdateParams struct {
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
	DataMode OnorbitantennaUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the antenna.
	IDAntenna string `json:"idAntenna,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Model representation of information on on-orbit/spacecraft communication
	// antennas. A spacecraft may have multiple antennas and each antenna can have
	// multiple 'details' records compiled by different sources.
	Antenna OnorbitantennaUpdateParamsAntenna `json:"antenna,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r OnorbitantennaUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitantennaUpdateParams
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
type OnorbitantennaUpdateParamsDataMode string

const (
	OnorbitantennaUpdateParamsDataModeReal      OnorbitantennaUpdateParamsDataMode = "REAL"
	OnorbitantennaUpdateParamsDataModeTest      OnorbitantennaUpdateParamsDataMode = "TEST"
	OnorbitantennaUpdateParamsDataModeSimulated OnorbitantennaUpdateParamsDataMode = "SIMULATED"
	OnorbitantennaUpdateParamsDataModeExercise  OnorbitantennaUpdateParamsDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft communication
// antennas. A spacecraft may have multiple antennas and each antenna can have
// multiple 'details' records compiled by different sources.
//
// The properties DataMode, Name, Source are required.
type OnorbitantennaUpdateParamsAntenna struct {
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
	// Antenna name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaUpdateParamsAntenna) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r OnorbitantennaUpdateParamsAntenna) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitantennaUpdateParamsAntenna
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnorbitantennaUpdateParamsAntenna](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitantennaListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [OnorbitantennaListParams]'s query parameters as
// `url.Values`.
func (r OnorbitantennaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitantennaGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnorbitantennaGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [OnorbitantennaGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitantennaGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
