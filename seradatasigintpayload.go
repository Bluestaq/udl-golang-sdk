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

// SeradataSigintPayloadService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradataSigintPayloadService] method instead.
type SeradataSigintPayloadService struct {
	Options []option.RequestOption
}

// NewSeradataSigintPayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradataSigintPayloadService(opts ...option.RequestOption) (r SeradataSigintPayloadService) {
	r = SeradataSigintPayloadService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataSigIntPayload as a POST body and
// ingest into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradataSigintPayloadService) New(ctx context.Context, body SeradataSigintPayloadNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatasigintpayload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataSigIntPayload. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *SeradataSigintPayloadService) Update(ctx context.Context, id string, body SeradataSigintPayloadUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatasigintpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeradataSigintPayloadService) List(ctx context.Context, query SeradataSigintPayloadListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradataSigintPayloadListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradatasigintpayload"
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
func (r *SeradataSigintPayloadService) ListAutoPaging(ctx context.Context, query SeradataSigintPayloadListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradataSigintPayloadListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataSigIntPayload specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradataSigintPayloadService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatasigintpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeradataSigintPayloadService) Count(ctx context.Context, query SeradataSigintPayloadCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradatasigintpayload/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataSigIntPayload by its unique ID passed
// as a path parameter.
func (r *SeradataSigintPayloadService) Get(ctx context.Context, id string, query SeradataSigintPayloadGetParams, opts ...option.RequestOption) (res *SeradataSigintPayloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatasigintpayload/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeradataSigintPayloadService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatasigintpayload/queryhelp"
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
func (r *SeradataSigintPayloadService) Tuple(ctx context.Context, query SeradataSigintPayloadTupleParams, opts ...option.RequestOption) (res *[]SeradataSigintPayloadTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradatasigintpayload/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details for an sigint payload from Seradata.
type SeradataSigintPayloadListResponse struct {
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
	DataMode SeradataSigintPayloadListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Frequency coverage for this payload.
	FrequencyCoverage string `json:"frequencyCoverage"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Intercept parameters.
	InterceptParameters string `json:"interceptParameters"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Positional Accuracy for this payload.
	PositionalAccuracy string `json:"positionalAccuracy"`
	// Swath Width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Source                 respjson.Field
		SpacecraftID           respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		FrequencyCoverage      respjson.Field
		GroundStationLocations respjson.Field
		GroundStations         respjson.Field
		HostedForCompanyOrgID  respjson.Field
		IDSensor               respjson.Field
		InterceptParameters    respjson.Field
		ManufacturerOrgID      respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		PositionalAccuracy     respjson.Field
		SwathWidth             respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSigintPayloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSigintPayloadListResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSigintPayloadListResponseDataMode string

const (
	SeradataSigintPayloadListResponseDataModeReal      SeradataSigintPayloadListResponseDataMode = "REAL"
	SeradataSigintPayloadListResponseDataModeTest      SeradataSigintPayloadListResponseDataMode = "TEST"
	SeradataSigintPayloadListResponseDataModeSimulated SeradataSigintPayloadListResponseDataMode = "SIMULATED"
	SeradataSigintPayloadListResponseDataModeExercise  SeradataSigintPayloadListResponseDataMode = "EXERCISE"
)

// Details for an sigint payload from Seradata.
type SeradataSigintPayloadGetResponse struct {
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
	DataMode SeradataSigintPayloadGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Frequency coverage for this payload.
	FrequencyCoverage string `json:"frequencyCoverage"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Intercept parameters.
	InterceptParameters string `json:"interceptParameters"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Positional Accuracy for this payload.
	PositionalAccuracy string `json:"positionalAccuracy"`
	// Swath Width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Source                 respjson.Field
		SpacecraftID           respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		FrequencyCoverage      respjson.Field
		GroundStationLocations respjson.Field
		GroundStations         respjson.Field
		HostedForCompanyOrgID  respjson.Field
		IDSensor               respjson.Field
		InterceptParameters    respjson.Field
		ManufacturerOrgID      respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		PositionalAccuracy     respjson.Field
		SwathWidth             respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSigintPayloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSigintPayloadGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSigintPayloadGetResponseDataMode string

const (
	SeradataSigintPayloadGetResponseDataModeReal      SeradataSigintPayloadGetResponseDataMode = "REAL"
	SeradataSigintPayloadGetResponseDataModeTest      SeradataSigintPayloadGetResponseDataMode = "TEST"
	SeradataSigintPayloadGetResponseDataModeSimulated SeradataSigintPayloadGetResponseDataMode = "SIMULATED"
	SeradataSigintPayloadGetResponseDataModeExercise  SeradataSigintPayloadGetResponseDataMode = "EXERCISE"
)

// Details for an sigint payload from Seradata.
type SeradataSigintPayloadTupleResponse struct {
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
	DataMode SeradataSigintPayloadTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Frequency coverage for this payload.
	FrequencyCoverage string `json:"frequencyCoverage"`
	// Ground Station Locations for this payload.
	GroundStationLocations string `json:"groundStationLocations"`
	// Ground Station info for this payload.
	GroundStations string `json:"groundStations"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// UUID of the Sensor record.
	IDSensor string `json:"idSensor"`
	// Intercept parameters.
	InterceptParameters string `json:"interceptParameters"`
	// Manufacturer Organization Id.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Sensor name from Seradata.
	Name string `json:"name"`
	// Payload notes.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Positional Accuracy for this payload.
	PositionalAccuracy string `json:"positionalAccuracy"`
	// Swath Width in kilometers.
	SwathWidth float64 `json:"swathWidth"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking  respjson.Field
		DataMode               respjson.Field
		Source                 respjson.Field
		SpacecraftID           respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		FrequencyCoverage      respjson.Field
		GroundStationLocations respjson.Field
		GroundStations         respjson.Field
		HostedForCompanyOrgID  respjson.Field
		IDSensor               respjson.Field
		InterceptParameters    respjson.Field
		ManufacturerOrgID      respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		Origin                 respjson.Field
		OrigNetwork            respjson.Field
		PositionalAccuracy     respjson.Field
		SwathWidth             respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		UpdatedBy              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradataSigintPayloadTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradataSigintPayloadTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradataSigintPayloadTupleResponseDataMode string

const (
	SeradataSigintPayloadTupleResponseDataModeReal      SeradataSigintPayloadTupleResponseDataMode = "REAL"
	SeradataSigintPayloadTupleResponseDataModeTest      SeradataSigintPayloadTupleResponseDataMode = "TEST"
	SeradataSigintPayloadTupleResponseDataModeSimulated SeradataSigintPayloadTupleResponseDataMode = "SIMULATED"
	SeradataSigintPayloadTupleResponseDataModeExercise  SeradataSigintPayloadTupleResponseDataMode = "EXERCISE"
)

type SeradataSigintPayloadNewParams struct {
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
	DataMode SeradataSigintPayloadNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Frequency coverage for this payload.
	FrequencyCoverage param.Opt[string] `json:"frequencyCoverage,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Intercept parameters.
	InterceptParameters param.Opt[string] `json:"interceptParameters,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Positional Accuracy for this payload.
	PositionalAccuracy param.Opt[string] `json:"positionalAccuracy,omitzero"`
	// Swath Width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SeradataSigintPayloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataSigintPayloadNewParams
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
type SeradataSigintPayloadNewParamsDataMode string

const (
	SeradataSigintPayloadNewParamsDataModeReal      SeradataSigintPayloadNewParamsDataMode = "REAL"
	SeradataSigintPayloadNewParamsDataModeTest      SeradataSigintPayloadNewParamsDataMode = "TEST"
	SeradataSigintPayloadNewParamsDataModeSimulated SeradataSigintPayloadNewParamsDataMode = "SIMULATED"
	SeradataSigintPayloadNewParamsDataModeExercise  SeradataSigintPayloadNewParamsDataMode = "EXERCISE"
)

type SeradataSigintPayloadUpdateParams struct {
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
	DataMode SeradataSigintPayloadUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Frequency coverage for this payload.
	FrequencyCoverage param.Opt[string] `json:"frequencyCoverage,omitzero"`
	// Ground Station Locations for this payload.
	GroundStationLocations param.Opt[string] `json:"groundStationLocations,omitzero"`
	// Ground Station info for this payload.
	GroundStations param.Opt[string] `json:"groundStations,omitzero"`
	// Hosted for company/Organization Id.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// UUID of the Sensor record.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// Intercept parameters.
	InterceptParameters param.Opt[string] `json:"interceptParameters,omitzero"`
	// Manufacturer Organization Id.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Sensor name from Seradata.
	Name param.Opt[string] `json:"name,omitzero"`
	// Payload notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Positional Accuracy for this payload.
	PositionalAccuracy param.Opt[string] `json:"positionalAccuracy,omitzero"`
	// Swath Width in kilometers.
	SwathWidth param.Opt[float64] `json:"swathWidth,omitzero"`
	// SIGINT Payload type, e.g. Comint, Elint, etc.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SeradataSigintPayloadUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradataSigintPayloadUpdateParams
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
type SeradataSigintPayloadUpdateParamsDataMode string

const (
	SeradataSigintPayloadUpdateParamsDataModeReal      SeradataSigintPayloadUpdateParamsDataMode = "REAL"
	SeradataSigintPayloadUpdateParamsDataModeTest      SeradataSigintPayloadUpdateParamsDataMode = "TEST"
	SeradataSigintPayloadUpdateParamsDataModeSimulated SeradataSigintPayloadUpdateParamsDataMode = "SIMULATED"
	SeradataSigintPayloadUpdateParamsDataModeExercise  SeradataSigintPayloadUpdateParamsDataMode = "EXERCISE"
)

type SeradataSigintPayloadListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSigintPayloadListParams]'s query parameters as
// `url.Values`.
func (r SeradataSigintPayloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSigintPayloadCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSigintPayloadCountParams]'s query parameters as
// `url.Values`.
func (r SeradataSigintPayloadCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSigintPayloadGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSigintPayloadGetParams]'s query parameters as
// `url.Values`.
func (r SeradataSigintPayloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradataSigintPayloadTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeradataSigintPayloadTupleParams]'s query parameters as
// `url.Values`.
func (r SeradataSigintPayloadTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
