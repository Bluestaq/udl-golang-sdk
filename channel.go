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

// ChannelService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	return
}

// Service operation to take a single Channel as a POST body and ingest into the
// database. A Comm payload may have multiple transponders and a transponder may
// have many channels. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ChannelService) New(ctx context.Context, body ChannelNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/channel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Channel record by its unique ID passed as a
// path parameter. A Comm payload may have multiple transponders and a transponder
// may have many channels.
func (r *ChannelService) Get(ctx context.Context, id string, query ChannelGetParams, opts ...option.RequestOption) (res *ChannelFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/channel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single Channel. A Comm payload may have multiple
// transponders and a transponder may have many channels. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *ChannelService) Update(ctx context.Context, id string, body ChannelUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/channel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ChannelService) List(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ChannelAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/channel"
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
func (r *ChannelService) ListAutoPaging(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ChannelAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Channel object specified by the passed ID path
// parameter. A Comm payload may have multiple transponders and a transponder may
// have many channels. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ChannelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/channel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ChannelService) Count(ctx context.Context, query ChannelCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/channel/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ChannelService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/channel/queryhelp"
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
func (r *ChannelService) Tuple(ctx context.Context, query ChannelTupleParams, opts ...option.RequestOption) (res *[]ChannelFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/channel/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Channel information on a particular transponder.
type ChannelAbridged struct {
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
	DataMode ChannelAbridgedDataMode `json:"dataMode,required"`
	// ID of the parent transponder object for this Channel.
	IDTransponder string `json:"idTransponder,required"`
	// Channel name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Channel aPid.
	Apid string `json:"apid"`
	// The antenna beam ID of the particular beam for this channel. beamName is not
	// unique across payloads.
	BeamName string `json:"beamName"`
	// Channel compression.
	Compression string `json:"compression"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Channel encryption.
	Encryption string `json:"encryption"`
	// Identifier of the particular beam for this channel.
	IDBeam string `json:"idBeam"`
	// ID of the RF Band object for this channel.
	IDRfBand string `json:"idRFBand"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Owner.
	Owner string `json:"owner"`
	// Pkg.
	Pkg string `json:"pkg"`
	// Res.
	Res string `json:"res"`
	// SID.
	Sid string `json:"sid"`
	// Channel type.
	Type string `json:"type"`
	// Channel vPid.
	Vpid string `json:"vpid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDTransponder         respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Apid                  respjson.Field
		BeamName              respjson.Field
		Compression           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Encryption            respjson.Field
		IDBeam                respjson.Field
		IDRfBand              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Owner                 respjson.Field
		Pkg                   respjson.Field
		Res                   respjson.Field
		Sid                   respjson.Field
		Type                  respjson.Field
		Vpid                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelAbridged) RawJSON() string { return r.JSON.raw }
func (r *ChannelAbridged) UnmarshalJSON(data []byte) error {
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
type ChannelAbridgedDataMode string

const (
	ChannelAbridgedDataModeReal      ChannelAbridgedDataMode = "REAL"
	ChannelAbridgedDataModeTest      ChannelAbridgedDataMode = "TEST"
	ChannelAbridgedDataModeSimulated ChannelAbridgedDataMode = "SIMULATED"
	ChannelAbridgedDataModeExercise  ChannelAbridgedDataMode = "EXERCISE"
)

// Channel information on a particular transponder.
type ChannelFull struct {
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
	DataMode ChannelFullDataMode `json:"dataMode,required"`
	// ID of the parent transponder object for this Channel.
	IDTransponder string `json:"idTransponder,required"`
	// Channel name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Channel aPid.
	Apid string `json:"apid"`
	// The antenna beam ID of the particular beam for this channel. beamName is not
	// unique across payloads.
	BeamName string `json:"beamName"`
	// Channel compression.
	Compression string `json:"compression"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Channel encryption.
	Encryption string `json:"encryption"`
	// Identifier of the particular beam for this channel.
	IDBeam string `json:"idBeam"`
	// ID of the RF Band object for this channel.
	IDRfBand string `json:"idRFBand"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Owner.
	Owner string `json:"owner"`
	// Pkg.
	Pkg string `json:"pkg"`
	// Res.
	Res string `json:"res"`
	// SID.
	Sid string `json:"sid"`
	// Channel type.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Channel vPid.
	Vpid string `json:"vpid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDTransponder         respjson.Field
		Name                  respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Apid                  respjson.Field
		BeamName              respjson.Field
		Compression           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Encryption            respjson.Field
		IDBeam                respjson.Field
		IDRfBand              respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Owner                 respjson.Field
		Pkg                   respjson.Field
		Res                   respjson.Field
		Sid                   respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Vpid                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelFull) RawJSON() string { return r.JSON.raw }
func (r *ChannelFull) UnmarshalJSON(data []byte) error {
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
type ChannelFullDataMode string

const (
	ChannelFullDataModeReal      ChannelFullDataMode = "REAL"
	ChannelFullDataModeTest      ChannelFullDataMode = "TEST"
	ChannelFullDataModeSimulated ChannelFullDataMode = "SIMULATED"
	ChannelFullDataModeExercise  ChannelFullDataMode = "EXERCISE"
)

type ChannelNewParams struct {
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
	DataMode ChannelNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent transponder object for this Channel.
	IDTransponder string `json:"idTransponder,required"`
	// Channel name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Channel aPid.
	Apid param.Opt[string] `json:"apid,omitzero"`
	// The antenna beam ID of the particular beam for this channel. beamName is not
	// unique across payloads.
	BeamName param.Opt[string] `json:"beamName,omitzero"`
	// Channel compression.
	Compression param.Opt[string] `json:"compression,omitzero"`
	// Channel encryption.
	Encryption param.Opt[string] `json:"encryption,omitzero"`
	// Identifier of the particular beam for this channel.
	IDBeam param.Opt[string] `json:"idBeam,omitzero"`
	// ID of the RF Band object for this channel.
	IDRfBand param.Opt[string] `json:"idRFBand,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Owner.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Pkg.
	Pkg param.Opt[string] `json:"pkg,omitzero"`
	// Res.
	Res param.Opt[string] `json:"res,omitzero"`
	// SID.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Channel type.
	Type param.Opt[string] `json:"type,omitzero"`
	// Channel vPid.
	Vpid param.Opt[string] `json:"vpid,omitzero"`
	paramObj
}

func (r ChannelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelNewParams) UnmarshalJSON(data []byte) error {
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
type ChannelNewParamsDataMode string

const (
	ChannelNewParamsDataModeReal      ChannelNewParamsDataMode = "REAL"
	ChannelNewParamsDataModeTest      ChannelNewParamsDataMode = "TEST"
	ChannelNewParamsDataModeSimulated ChannelNewParamsDataMode = "SIMULATED"
	ChannelNewParamsDataModeExercise  ChannelNewParamsDataMode = "EXERCISE"
)

type ChannelGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelGetParams]'s query parameters as `url.Values`.
func (r ChannelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChannelUpdateParams struct {
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
	DataMode ChannelUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent transponder object for this Channel.
	IDTransponder string `json:"idTransponder,required"`
	// Channel name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Channel aPid.
	Apid param.Opt[string] `json:"apid,omitzero"`
	// The antenna beam ID of the particular beam for this channel. beamName is not
	// unique across payloads.
	BeamName param.Opt[string] `json:"beamName,omitzero"`
	// Channel compression.
	Compression param.Opt[string] `json:"compression,omitzero"`
	// Channel encryption.
	Encryption param.Opt[string] `json:"encryption,omitzero"`
	// Identifier of the particular beam for this channel.
	IDBeam param.Opt[string] `json:"idBeam,omitzero"`
	// ID of the RF Band object for this channel.
	IDRfBand param.Opt[string] `json:"idRFBand,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Owner.
	Owner param.Opt[string] `json:"owner,omitzero"`
	// Pkg.
	Pkg param.Opt[string] `json:"pkg,omitzero"`
	// Res.
	Res param.Opt[string] `json:"res,omitzero"`
	// SID.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Channel type.
	Type param.Opt[string] `json:"type,omitzero"`
	// Channel vPid.
	Vpid param.Opt[string] `json:"vpid,omitzero"`
	paramObj
}

func (r ChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelUpdateParams) UnmarshalJSON(data []byte) error {
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
type ChannelUpdateParamsDataMode string

const (
	ChannelUpdateParamsDataModeReal      ChannelUpdateParamsDataMode = "REAL"
	ChannelUpdateParamsDataModeTest      ChannelUpdateParamsDataMode = "TEST"
	ChannelUpdateParamsDataModeSimulated ChannelUpdateParamsDataMode = "SIMULATED"
	ChannelUpdateParamsDataModeExercise  ChannelUpdateParamsDataMode = "EXERCISE"
)

type ChannelListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelListParams]'s query parameters as `url.Values`.
func (r ChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChannelCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelCountParams]'s query parameters as `url.Values`.
func (r ChannelCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChannelTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelTupleParams]'s query parameters as `url.Values`.
func (r ChannelTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
