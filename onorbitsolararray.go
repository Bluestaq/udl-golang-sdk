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

// OnorbitsolararrayService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitsolararrayService] method instead.
type OnorbitsolararrayService struct {
	Options []option.RequestOption
}

// NewOnorbitsolararrayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnorbitsolararrayService(opts ...option.RequestOption) (r OnorbitsolararrayService) {
	r = OnorbitsolararrayService{}
	r.Options = opts
	return
}

// Service operation to take a single OnorbitSolarArray as a POST body and ingest
// into the database. An OnorbitSolarArray is the association between on-orbit
// spacecraft SolarArrays and a particular on-orbit spacecraft. A SolarArray type
// may be associated with many different on-orbit spacecraft. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *OnorbitsolararrayService) New(ctx context.Context, body OnorbitsolararrayNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/onorbitsolararray"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single OnorbitSolarArray. An OnorbitSolarArray is
// the association between on-orbit spacecraft SolarArrays and a particular
// on-orbit spacecraft. A SolarArray type may be associated with many different
// on-orbit spacecraft. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *OnorbitsolararrayService) Update(ctx context.Context, id string, body OnorbitsolararrayUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitsolararray/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitsolararrayService) List(ctx context.Context, query OnorbitsolararrayListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OnorbitsolararrayListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/onorbitsolararray"
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
func (r *OnorbitsolararrayService) ListAutoPaging(ctx context.Context, query OnorbitsolararrayListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OnorbitsolararrayListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a OnorbitSolarArray object specified by the passed
// ID path parameter. An OnorbitSolarArray is the association between on-orbit
// spacecraft SolarArrays and a particular on-orbit spacecraft. A SolarArray type
// may be associated with many different on-orbit spacecraft. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *OnorbitsolararrayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitsolararray/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to get a single OnorbitSolarArray record by its unique ID
// passed as a path parameter. An OnorbitSolarArray is the association between
// on-orbit spacecraft SolarArrays and a particular on-orbit spacecraft. A
// SolarArray type may be associated with many different on-orbit spacecraft.
func (r *OnorbitsolararrayService) Get(ctx context.Context, id string, query OnorbitsolararrayGetParams, opts ...option.RequestOption) (res *OnorbitsolararrayGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/onorbitsolararray/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OnorbitsolararrayListResponse struct {
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
	DataMode OnorbitsolararrayListResponseDataMode `json:"dataMode,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// ID of the SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
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
	// The number of solar arrays on the spacecraft of the type identified by
	// idSolarArray.
	Quantity int64 `json:"quantity"`
	// Model representation of information on on-orbit/spacecraft solar arrays. A
	// spacecraft may have multiple solar arrays and each solar array can have multiple
	// 'details' records compiled by different sources.
	SolarArray OnorbitsolararrayListResponseSolarArray `json:"solarArray"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOnOrbit             respjson.Field
		IDSolarArray          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		SolarArray            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitsolararrayListResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitsolararrayListResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitsolararrayListResponseDataMode string

const (
	OnorbitsolararrayListResponseDataModeReal      OnorbitsolararrayListResponseDataMode = "REAL"
	OnorbitsolararrayListResponseDataModeTest      OnorbitsolararrayListResponseDataMode = "TEST"
	OnorbitsolararrayListResponseDataModeSimulated OnorbitsolararrayListResponseDataMode = "SIMULATED"
	OnorbitsolararrayListResponseDataModeExercise  OnorbitsolararrayListResponseDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
type OnorbitsolararrayListResponseSolarArray struct {
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
	// Solar Array name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
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
		DataMode    respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Origin      respjson.Field
		OrigNetwork respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitsolararrayListResponseSolarArray) RawJSON() string { return r.JSON.raw }
func (r *OnorbitsolararrayListResponseSolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitsolararrayGetResponse struct {
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
	DataMode OnorbitsolararrayGetResponseDataMode `json:"dataMode,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// ID of the SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
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
	// The number of solar arrays on the spacecraft of the type identified by
	// idSolarArray.
	Quantity int64 `json:"quantity"`
	// Model representation of information on on-orbit/spacecraft solar arrays. A
	// spacecraft may have multiple solar arrays and each solar array can have multiple
	// 'details' records compiled by different sources.
	SolarArray OnorbitsolararrayGetResponseSolarArray `json:"solarArray"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDOnOrbit             respjson.Field
		IDSolarArray          respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Quantity              respjson.Field
		SolarArray            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitsolararrayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OnorbitsolararrayGetResponse) UnmarshalJSON(data []byte) error {
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
type OnorbitsolararrayGetResponseDataMode string

const (
	OnorbitsolararrayGetResponseDataModeReal      OnorbitsolararrayGetResponseDataMode = "REAL"
	OnorbitsolararrayGetResponseDataModeTest      OnorbitsolararrayGetResponseDataMode = "TEST"
	OnorbitsolararrayGetResponseDataModeSimulated OnorbitsolararrayGetResponseDataMode = "SIMULATED"
	OnorbitsolararrayGetResponseDataModeExercise  OnorbitsolararrayGetResponseDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
type OnorbitsolararrayGetResponseSolarArray struct {
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
	// Solar Array name.
	Name string `json:"name,required"`
	// Source of the data.
	Source string `json:"source,required"`
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
	// Read-only collection of additional SolarArrayDetails by various sources for this
	// organization, ignored on create/update. These details must be created separately
	// via the /udl/solararraydetails operations.
	SolarArrayDetails []SolarArrayDetailsFull `json:"solarArrayDetails"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataMode          respjson.Field
		Name              respjson.Field
		Source            respjson.Field
		ID                respjson.Field
		CreatedAt         respjson.Field
		CreatedBy         respjson.Field
		Origin            respjson.Field
		OrigNetwork       respjson.Field
		SolarArrayDetails respjson.Field
		UpdatedAt         respjson.Field
		UpdatedBy         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnorbitsolararrayGetResponseSolarArray) RawJSON() string { return r.JSON.raw }
func (r *OnorbitsolararrayGetResponseSolarArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnorbitsolararrayNewParams struct {
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
	DataMode OnorbitsolararrayNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// ID of the SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The number of solar arrays on the spacecraft of the type identified by
	// idSolarArray.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Model representation of information on on-orbit/spacecraft solar arrays. A
	// spacecraft may have multiple solar arrays and each solar array can have multiple
	// 'details' records compiled by different sources.
	SolarArray OnorbitsolararrayNewParamsSolarArray `json:"solarArray,omitzero"`
	paramObj
}

func (r OnorbitsolararrayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitsolararrayNewParams
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
type OnorbitsolararrayNewParamsDataMode string

const (
	OnorbitsolararrayNewParamsDataModeReal      OnorbitsolararrayNewParamsDataMode = "REAL"
	OnorbitsolararrayNewParamsDataModeTest      OnorbitsolararrayNewParamsDataMode = "TEST"
	OnorbitsolararrayNewParamsDataModeSimulated OnorbitsolararrayNewParamsDataMode = "SIMULATED"
	OnorbitsolararrayNewParamsDataModeExercise  OnorbitsolararrayNewParamsDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
//
// The properties DataMode, Name, Source are required.
type OnorbitsolararrayNewParamsSolarArray struct {
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
	// Solar Array name.
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

func (r OnorbitsolararrayNewParamsSolarArray) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitsolararrayNewParamsSolarArray
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnorbitsolararrayNewParamsSolarArray](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitsolararrayUpdateParams struct {
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
	DataMode OnorbitsolararrayUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the on-orbit object.
	IDOnOrbit string `json:"idOnOrbit,required"`
	// ID of the SolarArray.
	IDSolarArray string `json:"idSolarArray,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The number of solar arrays on the spacecraft of the type identified by
	// idSolarArray.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Model representation of information on on-orbit/spacecraft solar arrays. A
	// spacecraft may have multiple solar arrays and each solar array can have multiple
	// 'details' records compiled by different sources.
	SolarArray OnorbitsolararrayUpdateParamsSolarArray `json:"solarArray,omitzero"`
	paramObj
}

func (r OnorbitsolararrayUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitsolararrayUpdateParams
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
type OnorbitsolararrayUpdateParamsDataMode string

const (
	OnorbitsolararrayUpdateParamsDataModeReal      OnorbitsolararrayUpdateParamsDataMode = "REAL"
	OnorbitsolararrayUpdateParamsDataModeTest      OnorbitsolararrayUpdateParamsDataMode = "TEST"
	OnorbitsolararrayUpdateParamsDataModeSimulated OnorbitsolararrayUpdateParamsDataMode = "SIMULATED"
	OnorbitsolararrayUpdateParamsDataModeExercise  OnorbitsolararrayUpdateParamsDataMode = "EXERCISE"
)

// Model representation of information on on-orbit/spacecraft solar arrays. A
// spacecraft may have multiple solar arrays and each solar array can have multiple
// 'details' records compiled by different sources.
//
// The properties DataMode, Name, Source are required.
type OnorbitsolararrayUpdateParamsSolarArray struct {
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
	// Solar Array name.
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

func (r OnorbitsolararrayUpdateParamsSolarArray) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitsolararrayUpdateParamsSolarArray
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnorbitsolararrayUpdateParamsSolarArray](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type OnorbitsolararrayListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitsolararrayListParams]'s query parameters as
// `url.Values`.
func (r OnorbitsolararrayListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitsolararrayGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitsolararrayGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitsolararrayGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
