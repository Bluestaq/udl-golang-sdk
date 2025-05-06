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

// BatterydetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatterydetailService] method instead.
type BatterydetailService struct {
	Options []option.RequestOption
}

// NewBatterydetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBatterydetailService(opts ...option.RequestOption) (r BatterydetailService) {
	r = BatterydetailService{}
	r.Options = opts
	return
}

// Service operation to take a single BatteryDetails as a POST body and ingest into
// the database. A Battery record may have multiple details records from several
// sources. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *BatterydetailService) New(ctx context.Context, body BatterydetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/batterydetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single BatteryDetails record by its unique ID passed
// as a path parameter. A Battery record may have multiple details records from
// several sources.
func (r *BatterydetailService) Get(ctx context.Context, id string, query BatterydetailGetParams, opts ...option.RequestOption) (res *BatterydetailsFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/batterydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single BatteryDetails. A Battery record may have
// multiple details records from several sources. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *BatterydetailService) Update(ctx context.Context, id string, body BatterydetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/batterydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *BatterydetailService) List(ctx context.Context, query BatterydetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BatterydetailsAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/batterydetails"
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
func (r *BatterydetailService) ListAutoPaging(ctx context.Context, query BatterydetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BatterydetailsAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a BatteryDetails object specified by the passed ID
// path parameter. A Battery record may have multiple details records from several
// sources. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *BatterydetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/batterydetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Detailed information on a spacecraft battery type compiled by a particular
// source. A Battery record may have multiple details records from several sources.
type BatterydetailsAbridged struct {
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
	DataMode BatterydetailsAbridgedDataMode `json:"dataMode,required"`
	// Identifier of the parent battery type record.
	IDBattery string `json:"idBattery,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Battery capacity in Ah.
	Capacity float64 `json:"capacity"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Battery description/notes.
	Description string `json:"description"`
	// Depth of discharge as a percentage/fraction.
	DischargeDepth float64 `json:"dischargeDepth"`
	// ID of the organization that manufactures the battery.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Battery model number or name.
	Model string `json:"model"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Type of battery technology (e.g. Ni-Cd, Ni-H2, Li-ion, etc.).
	Technology string `json:"technology"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDBattery             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Capacity              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		DischargeDepth        respjson.Field
		ManufacturerOrgID     respjson.Field
		Model                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Technology            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatterydetailsAbridged) RawJSON() string { return r.JSON.raw }
func (r *BatterydetailsAbridged) UnmarshalJSON(data []byte) error {
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
type BatterydetailsAbridgedDataMode string

const (
	BatterydetailsAbridgedDataModeReal      BatterydetailsAbridgedDataMode = "REAL"
	BatterydetailsAbridgedDataModeTest      BatterydetailsAbridgedDataMode = "TEST"
	BatterydetailsAbridgedDataModeSimulated BatterydetailsAbridgedDataMode = "SIMULATED"
	BatterydetailsAbridgedDataModeExercise  BatterydetailsAbridgedDataMode = "EXERCISE"
)

// Detailed information on a spacecraft battery type compiled by a particular
// source. A Battery record may have multiple details records from several sources.
type BatterydetailsFull struct {
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
	DataMode BatterydetailsFullDataMode `json:"dataMode,required"`
	// Identifier of the parent battery type record.
	IDBattery string `json:"idBattery,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Battery capacity in Ah.
	Capacity float64 `json:"capacity"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Battery description/notes.
	Description string `json:"description"`
	// Depth of discharge as a percentage/fraction.
	DischargeDepth float64 `json:"dischargeDepth"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg OrganizationFull `json:"manufacturerOrg"`
	// ID of the organization that manufactures the battery.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Battery model number or name.
	Model string `json:"model"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Type of battery technology (e.g. Ni-Cd, Ni-H2, Li-ion, etc.).
	Technology string `json:"technology"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDBattery             respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Capacity              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Description           respjson.Field
		DischargeDepth        respjson.Field
		ManufacturerOrg       respjson.Field
		ManufacturerOrgID     respjson.Field
		Model                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Tags                  respjson.Field
		Technology            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatterydetailsFull) RawJSON() string { return r.JSON.raw }
func (r *BatterydetailsFull) UnmarshalJSON(data []byte) error {
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
type BatterydetailsFullDataMode string

const (
	BatterydetailsFullDataModeReal      BatterydetailsFullDataMode = "REAL"
	BatterydetailsFullDataModeTest      BatterydetailsFullDataMode = "TEST"
	BatterydetailsFullDataModeSimulated BatterydetailsFullDataMode = "SIMULATED"
	BatterydetailsFullDataModeExercise  BatterydetailsFullDataMode = "EXERCISE"
)

type BatterydetailNewParams struct {
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
	DataMode BatterydetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent battery type record.
	IDBattery string `json:"idBattery,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Battery capacity in Ah.
	Capacity param.Opt[float64] `json:"capacity,omitzero"`
	// Battery description/notes.
	Description param.Opt[string] `json:"description,omitzero"`
	// Depth of discharge as a percentage/fraction.
	DischargeDepth param.Opt[float64] `json:"dischargeDepth,omitzero"`
	// ID of the organization that manufactures the battery.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Battery model number or name.
	Model param.Opt[string] `json:"model,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Type of battery technology (e.g. Ni-Cd, Ni-H2, Li-ion, etc.).
	Technology param.Opt[string] `json:"technology,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BatterydetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatterydetailNewParams
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
type BatterydetailNewParamsDataMode string

const (
	BatterydetailNewParamsDataModeReal      BatterydetailNewParamsDataMode = "REAL"
	BatterydetailNewParamsDataModeTest      BatterydetailNewParamsDataMode = "TEST"
	BatterydetailNewParamsDataModeSimulated BatterydetailNewParamsDataMode = "SIMULATED"
	BatterydetailNewParamsDataModeExercise  BatterydetailNewParamsDataMode = "EXERCISE"
)

type BatterydetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatterydetailGetParams]'s query parameters as `url.Values`.
func (r BatterydetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatterydetailUpdateParams struct {
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
	DataMode BatterydetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Identifier of the parent battery type record.
	IDBattery string `json:"idBattery,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Battery capacity in Ah.
	Capacity param.Opt[float64] `json:"capacity,omitzero"`
	// Battery description/notes.
	Description param.Opt[string] `json:"description,omitzero"`
	// Depth of discharge as a percentage/fraction.
	DischargeDepth param.Opt[float64] `json:"dischargeDepth,omitzero"`
	// ID of the organization that manufactures the battery.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Battery model number or name.
	Model param.Opt[string] `json:"model,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Type of battery technology (e.g. Ni-Cd, Ni-H2, Li-ion, etc.).
	Technology param.Opt[string] `json:"technology,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BatterydetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BatterydetailUpdateParams
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
type BatterydetailUpdateParamsDataMode string

const (
	BatterydetailUpdateParamsDataModeReal      BatterydetailUpdateParamsDataMode = "REAL"
	BatterydetailUpdateParamsDataModeTest      BatterydetailUpdateParamsDataMode = "TEST"
	BatterydetailUpdateParamsDataModeSimulated BatterydetailUpdateParamsDataMode = "SIMULATED"
	BatterydetailUpdateParamsDataModeExercise  BatterydetailUpdateParamsDataMode = "EXERCISE"
)

type BatterydetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatterydetailListParams]'s query parameters as
// `url.Values`.
func (r BatterydetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
