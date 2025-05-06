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

// OnorbitAntennaDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnorbitAntennaDetailService] method instead.
type OnorbitAntennaDetailService struct {
	Options []option.RequestOption
}

// NewOnorbitAntennaDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnorbitAntennaDetailService(opts ...option.RequestOption) (r OnorbitAntennaDetailService) {
	r = OnorbitAntennaDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single AntennaDetails as a POST body and ingest into
// the database. An antenna may have multiple details records compiled by various
// sources. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *OnorbitAntennaDetailService) New(ctx context.Context, body OnorbitAntennaDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/antennadetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single AntennaDetails record by its unique ID passed
// as a path parameter. An antenna may have multiple details records compiled by
// various sources.
func (r *OnorbitAntennaDetailService) Get(ctx context.Context, id string, query OnorbitAntennaDetailGetParams, opts ...option.RequestOption) (res *AntennaDetailsFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/antennadetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single AntennaDetails. An antenna may have
// multiple details records compiled by various sources. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *OnorbitAntennaDetailService) Update(ctx context.Context, id string, body OnorbitAntennaDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/antennadetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OnorbitAntennaDetailService) List(ctx context.Context, query OnorbitAntennaDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AntennaDetailsAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/antennadetails"
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
func (r *OnorbitAntennaDetailService) ListAutoPaging(ctx context.Context, query OnorbitAntennaDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AntennaDetailsAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a AntennaDetails object specified by the passed ID
// path parameter. An antenna may have multiple details records compiled by various
// sources. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *OnorbitAntennaDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/antennadetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Detailed information for a spacecraft communication antenna. One antenna may
// have multiple AntennaDetails records, compiled by various sources.
type AntennaDetailsAbridged struct {
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
	DataMode AntennaDetailsAbridgedDataMode `json:"dataMode,required"`
	// Unique identifier of the parent Antenna.
	IDAntenna string `json:"idAntenna,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Boolean indicating if this is a beam forming antenna.
	BeamForming bool `json:"beamForming"`
	// Array of angles between the half-power (-3 dB) points of the main lobe of the
	// antenna, in degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Antenna description.
	Description string `json:"description"`
	// Antenna diameter in meters.
	Diameter float64 `json:"diameter"`
	// Antenna end of frequency range in Mhz.
	EndFrequency float64 `json:"endFrequency"`
	// Antenna maximum gain in dBi.
	Gain float64 `json:"gain"`
	// Antenna gain tolerance in dB.
	GainTolerance float64 `json:"gainTolerance"`
	// ID of the organization that manufactures the antenna.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Antenna mode (e.g. TX,RX).
	//
	// Any of "TX", "RX".
	Mode AntennaDetailsAbridgedMode `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Antenna polarization in degrees.
	Polarization float64 `json:"polarization"`
	// Antenna position (e.g. Top, Nadir, Side).
	Position string `json:"position"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	Size []float64 `json:"size"`
	// Antenna start of frequency range in Mhz.
	StartFrequency float64 `json:"startFrequency"`
	// Boolean indicating if this antenna is steerable.
	Steerable bool `json:"steerable"`
	// Type of antenna (e.g. Reflector, Double Reflector, Shaped Reflector, Horn,
	// Parabolic, etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDAntenna             resp.Field
		Source                resp.Field
		ID                    resp.Field
		BeamForming           resp.Field
		Beamwidth             resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Description           resp.Field
		Diameter              resp.Field
		EndFrequency          resp.Field
		Gain                  resp.Field
		GainTolerance         resp.Field
		ManufacturerOrgID     resp.Field
		Mode                  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Polarization          resp.Field
		Position              resp.Field
		Size                  resp.Field
		StartFrequency        resp.Field
		Steerable             resp.Field
		Type                  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AntennaDetailsAbridged) RawJSON() string { return r.JSON.raw }
func (r *AntennaDetailsAbridged) UnmarshalJSON(data []byte) error {
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
type AntennaDetailsAbridgedDataMode string

const (
	AntennaDetailsAbridgedDataModeReal      AntennaDetailsAbridgedDataMode = "REAL"
	AntennaDetailsAbridgedDataModeTest      AntennaDetailsAbridgedDataMode = "TEST"
	AntennaDetailsAbridgedDataModeSimulated AntennaDetailsAbridgedDataMode = "SIMULATED"
	AntennaDetailsAbridgedDataModeExercise  AntennaDetailsAbridgedDataMode = "EXERCISE"
)

// Antenna mode (e.g. TX,RX).
type AntennaDetailsAbridgedMode string

const (
	AntennaDetailsAbridgedModeTx AntennaDetailsAbridgedMode = "TX"
	AntennaDetailsAbridgedModeRx AntennaDetailsAbridgedMode = "RX"
)

// Detailed information for a spacecraft communication antenna. One antenna may
// have multiple AntennaDetails records, compiled by various sources.
type AntennaDetailsFull struct {
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
	DataMode AntennaDetailsFullDataMode `json:"dataMode,required"`
	// Unique identifier of the parent Antenna.
	IDAntenna string `json:"idAntenna,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Boolean indicating if this is a beam forming antenna.
	BeamForming bool `json:"beamForming"`
	// Array of angles between the half-power (-3 dB) points of the main lobe of the
	// antenna, in degrees.
	Beamwidth float64 `json:"beamwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Antenna description.
	Description string `json:"description"`
	// Antenna diameter in meters.
	Diameter float64 `json:"diameter"`
	// Antenna end of frequency range in Mhz.
	EndFrequency float64 `json:"endFrequency"`
	// Antenna maximum gain in dBi.
	Gain float64 `json:"gain"`
	// Antenna gain tolerance in dB.
	GainTolerance float64 `json:"gainTolerance"`
	// An organization such as a corporation, manufacturer, consortium, government,
	// etc. An organization may have parent and child organizations as well as link to
	// a former organization if this org previously existed as another organization.
	ManufacturerOrg OrganizationFull `json:"manufacturerOrg"`
	// ID of the organization that manufactures the antenna.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Antenna mode (e.g. TX,RX).
	//
	// Any of "TX", "RX".
	Mode AntennaDetailsFullMode `json:"mode"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Antenna polarization in degrees.
	Polarization float64 `json:"polarization"`
	// Antenna position (e.g. Top, Nadir, Side).
	Position string `json:"position"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	Size []float64 `json:"size"`
	// Antenna start of frequency range in Mhz.
	StartFrequency float64 `json:"startFrequency"`
	// Boolean indicating if this antenna is steerable.
	Steerable bool `json:"steerable"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Type of antenna (e.g. Reflector, Double Reflector, Shaped Reflector, Horn,
	// Parabolic, etc.).
	Type string `json:"type"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		IDAntenna             resp.Field
		Source                resp.Field
		ID                    resp.Field
		BeamForming           resp.Field
		Beamwidth             resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Description           resp.Field
		Diameter              resp.Field
		EndFrequency          resp.Field
		Gain                  resp.Field
		GainTolerance         resp.Field
		ManufacturerOrg       resp.Field
		ManufacturerOrgID     resp.Field
		Mode                  resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Polarization          resp.Field
		Position              resp.Field
		Size                  resp.Field
		StartFrequency        resp.Field
		Steerable             resp.Field
		Tags                  resp.Field
		Type                  resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AntennaDetailsFull) RawJSON() string { return r.JSON.raw }
func (r *AntennaDetailsFull) UnmarshalJSON(data []byte) error {
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
type AntennaDetailsFullDataMode string

const (
	AntennaDetailsFullDataModeReal      AntennaDetailsFullDataMode = "REAL"
	AntennaDetailsFullDataModeTest      AntennaDetailsFullDataMode = "TEST"
	AntennaDetailsFullDataModeSimulated AntennaDetailsFullDataMode = "SIMULATED"
	AntennaDetailsFullDataModeExercise  AntennaDetailsFullDataMode = "EXERCISE"
)

// Antenna mode (e.g. TX,RX).
type AntennaDetailsFullMode string

const (
	AntennaDetailsFullModeTx AntennaDetailsFullMode = "TX"
	AntennaDetailsFullModeRx AntennaDetailsFullMode = "RX"
)

type OnorbitAntennaDetailNewParams struct {
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
	DataMode OnorbitAntennaDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent Antenna.
	IDAntenna string `json:"idAntenna,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Boolean indicating if this is a beam forming antenna.
	BeamForming param.Opt[bool] `json:"beamForming,omitzero"`
	// Array of angles between the half-power (-3 dB) points of the main lobe of the
	// antenna, in degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Antenna description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Antenna diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Antenna end of frequency range in Mhz.
	EndFrequency param.Opt[float64] `json:"endFrequency,omitzero"`
	// Antenna maximum gain in dBi.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// Antenna gain tolerance in dB.
	GainTolerance param.Opt[float64] `json:"gainTolerance,omitzero"`
	// ID of the organization that manufactures the antenna.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Antenna polarization in degrees.
	Polarization param.Opt[float64] `json:"polarization,omitzero"`
	// Antenna position (e.g. Top, Nadir, Side).
	Position param.Opt[string] `json:"position,omitzero"`
	// Antenna start of frequency range in Mhz.
	StartFrequency param.Opt[float64] `json:"startFrequency,omitzero"`
	// Boolean indicating if this antenna is steerable.
	Steerable param.Opt[bool] `json:"steerable,omitzero"`
	// Type of antenna (e.g. Reflector, Double Reflector, Shaped Reflector, Horn,
	// Parabolic, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Antenna mode (e.g. TX,RX).
	//
	// Any of "TX", "RX".
	Mode OnorbitAntennaDetailNewParamsMode `json:"mode,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	Size []float64 `json:"size,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OnorbitAntennaDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitAntennaDetailNewParams
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
type OnorbitAntennaDetailNewParamsDataMode string

const (
	OnorbitAntennaDetailNewParamsDataModeReal      OnorbitAntennaDetailNewParamsDataMode = "REAL"
	OnorbitAntennaDetailNewParamsDataModeTest      OnorbitAntennaDetailNewParamsDataMode = "TEST"
	OnorbitAntennaDetailNewParamsDataModeSimulated OnorbitAntennaDetailNewParamsDataMode = "SIMULATED"
	OnorbitAntennaDetailNewParamsDataModeExercise  OnorbitAntennaDetailNewParamsDataMode = "EXERCISE"
)

// Antenna mode (e.g. TX,RX).
type OnorbitAntennaDetailNewParamsMode string

const (
	OnorbitAntennaDetailNewParamsModeTx OnorbitAntennaDetailNewParamsMode = "TX"
	OnorbitAntennaDetailNewParamsModeRx OnorbitAntennaDetailNewParamsMode = "RX"
)

type OnorbitAntennaDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitAntennaDetailGetParams]'s query parameters as
// `url.Values`.
func (r OnorbitAntennaDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OnorbitAntennaDetailUpdateParams struct {
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
	DataMode OnorbitAntennaDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Unique identifier of the parent Antenna.
	IDAntenna string `json:"idAntenna,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Boolean indicating if this is a beam forming antenna.
	BeamForming param.Opt[bool] `json:"beamForming,omitzero"`
	// Array of angles between the half-power (-3 dB) points of the main lobe of the
	// antenna, in degrees.
	Beamwidth param.Opt[float64] `json:"beamwidth,omitzero"`
	// Antenna description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Antenna diameter in meters.
	Diameter param.Opt[float64] `json:"diameter,omitzero"`
	// Antenna end of frequency range in Mhz.
	EndFrequency param.Opt[float64] `json:"endFrequency,omitzero"`
	// Antenna maximum gain in dBi.
	Gain param.Opt[float64] `json:"gain,omitzero"`
	// Antenna gain tolerance in dB.
	GainTolerance param.Opt[float64] `json:"gainTolerance,omitzero"`
	// ID of the organization that manufactures the antenna.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Antenna polarization in degrees.
	Polarization param.Opt[float64] `json:"polarization,omitzero"`
	// Antenna position (e.g. Top, Nadir, Side).
	Position param.Opt[string] `json:"position,omitzero"`
	// Antenna start of frequency range in Mhz.
	StartFrequency param.Opt[float64] `json:"startFrequency,omitzero"`
	// Boolean indicating if this antenna is steerable.
	Steerable param.Opt[bool] `json:"steerable,omitzero"`
	// Type of antenna (e.g. Reflector, Double Reflector, Shaped Reflector, Horn,
	// Parabolic, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	// Antenna mode (e.g. TX,RX).
	//
	// Any of "TX", "RX".
	Mode OnorbitAntennaDetailUpdateParamsMode `json:"mode,omitzero"`
	// Array with 1-2 values specifying the length and width (for rectangular) and just
	// length for dipole antennas in meters.
	Size []float64 `json:"size,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OnorbitAntennaDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OnorbitAntennaDetailUpdateParams
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
type OnorbitAntennaDetailUpdateParamsDataMode string

const (
	OnorbitAntennaDetailUpdateParamsDataModeReal      OnorbitAntennaDetailUpdateParamsDataMode = "REAL"
	OnorbitAntennaDetailUpdateParamsDataModeTest      OnorbitAntennaDetailUpdateParamsDataMode = "TEST"
	OnorbitAntennaDetailUpdateParamsDataModeSimulated OnorbitAntennaDetailUpdateParamsDataMode = "SIMULATED"
	OnorbitAntennaDetailUpdateParamsDataModeExercise  OnorbitAntennaDetailUpdateParamsDataMode = "EXERCISE"
)

// Antenna mode (e.g. TX,RX).
type OnorbitAntennaDetailUpdateParamsMode string

const (
	OnorbitAntennaDetailUpdateParamsModeTx OnorbitAntennaDetailUpdateParamsMode = "TX"
	OnorbitAntennaDetailUpdateParamsModeRx OnorbitAntennaDetailUpdateParamsMode = "RX"
)

type OnorbitAntennaDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OnorbitAntennaDetailListParams]'s query parameters as
// `url.Values`.
func (r OnorbitAntennaDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
