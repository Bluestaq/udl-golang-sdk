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
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// EopService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEopService] method instead.
type EopService struct {
	Options []option.RequestOption
	History EopHistoryService
}

// NewEopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEopService(opts ...option.RequestOption) (r EopService) {
	r = EopService{}
	r.Options = opts
	r.History = NewEopHistoryService(opts...)
	return
}

// Service operation to take a single EOP Record as a POST body and ingest into the
// database. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *EopService) New(ctx context.Context, body EopNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/eop"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single EOP record by its unique ID passed as a path
// parameter.
func (r *EopService) Get(ctx context.Context, id string, query EopGetParams, opts ...option.RequestOption) (res *shared.EopFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/eop/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to update a single EOP Record. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance.
func (r *EopService) Update(ctx context.Context, id string, body EopUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/eop/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *EopService) List(ctx context.Context, query EopListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[EopAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eop"
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
func (r *EopService) ListAutoPaging(ctx context.Context, query EopListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[EopAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an EOP Record specified by the passed ID path
// parameter. Note, delete operations do not remove data from historical or
// publish/subscribe stores. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *EopService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/eop/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *EopService) Count(ctx context.Context, query EopCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eop/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *EopService) ListTuple(ctx context.Context, query EopListTupleParams, opts ...option.RequestOption) (res *[]shared.EopFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/eop/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *EopService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *EopQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/eop/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model representation of Earth Orientation Parameters (EOP) produced by the IERS
// (International Earth Rotation and Reference Systems Service). Earth Orientation
// Parameters describe the irregularities of the earth's rotation. Technically,
// they are the parameters which provide the rotation of the ITRS (International
// Terrestrial Reference System) to the ICRS (International Celestial Reference
// System) as a function of time. Universal time -- Universal time (UT1) is the
// time of the earth clock, which performs one revolution in about 24h. It is
// practically proportional to the sidereal time. The excess revolution time is
// called length of day (LOD). Coordinates of the pole -- x and y are the
// coordinates of the Celestial Ephemeris Pole (CEP) relative to the IRP, the IERS
// Reference Pole. The CEP differs from the instantaneous rotation axis by
// quasi-diurnal terms with amplitudes under 0.01" (see Seidelmann, 1982). The
// x-axis is in the direction of the ITRF zero-meridian; the y-axis is in the
// direction 90 degrees West longitude. Celestial pole offsets -- Celestial pole
// offsets are described in the IAU Precession and Nutation models. The observed
// differences with respect to the conventional celestial pole position defined by
// the models are monitored and reported by the IERS. IERS Bulletins A and B
// provide current information on the Earth's orientation in the IERS Reference
// System. This includes Universal Time, coordinates of the terrestrial pole, and
// celestial pole offsets. Bulletin A gives an advanced solution updated weekly;
// the standard solution is given monthly in Bulletin B. Fields suffixed with 'B'
// are Bulletin B values. All solutions are continuous within their respective
// uncertainties. Bulletin A is issued by the IERS Rapid Service/Prediction Centre
// at the U.S. Naval Observatory, Washington, DC and Bulletin B is issued by the
// IERS Earth Orientation Centre at the Paris Observatory. IERS Bulletin A reports
// the latest determinations for polar motion, UT1-UTC, and nutation offsets at
// daily intervals based on a combination of contributed analysis results using
// data from Very Long Baseline Interferometry (VLBI), Satellite Laser Ranging
// (SLR), Global Positioning System (GPS) satellites, and Lunar Laser Ranging
// (LLR). Predictions for variations a year into the future are also provided.
// Meteorological predictions of variations in Atmospheric Angular Momentum (AAM)
// are used to aid in the prediction of near-term UT1-UTC changes. This publication
// is prepared by the IERS Rapid Service/Prediction Center.
type EopAbridged struct {
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
	DataMode EopAbridgedDataMode `json:"dataMode,required"`
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted.
	EopDate time.Time `json:"eopDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The Bulletin A offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilon is not used when this record represents IAU 2000 data.
	DEpsilon float64 `json:"dEpsilon"`
	// The Bulletin B offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilonB is not used when this record represents IAU 2000 data.
	DEpsilonB float64 `json:"dEpsilonB"`
	// The estimated uncertainty/error in the dEpsilon value in milliseconds of arc.
	DEpsilonUnc float64 `json:"dEpsilonUnc"`
	// The Bulletin A offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSI is not used when this record represents IAU 2000 data.
	DPsi float64 `json:"dPSI"`
	// The Bulletin B offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSIB is not used when this record represents IAU 2000 data.
	DPsib float64 `json:"dPSIB"`
	// The estimated uncertainty/error in the dPSI value in milliseconds of arc.
	DPsiUnc float64 `json:"dPSIUnc"`
	// The Bulletin A celestial pole offset along x-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dX is
	// not used when this record represents IAU 1980 data.
	DX float64 `json:"dX"`
	// The Bulletin B celestial pole offset along the X-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dXB is not used when this record represents IAU 1980 data.
	DXb float64 `json:"dXB"`
	// The estimated uncertainty/error in the Bulletin A dX value, in milliseconds of
	// arc.
	DXUnc float64 `json:"dXUnc"`
	// The Bulletin A celestial pole offset along y-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dY is
	// not used when this record represents IAU 1980 data.
	DY float64 `json:"dY"`
	// The Bulletin B celestial pole offset along the Y-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dYB is not used when this record represents IAU 1980 data.
	DYb float64 `json:"dYB"`
	// The estimated uncertainty/error in the Bulletin A dY value, in milliseconds of
	// arc.
	DYUnc float64 `json:"dYUnc"`
	// Bulletin A length of day or LOD in milliseconds. Universal time (UT1) is the
	// time of the earth clock, which performs one revolution in about 24h. It is
	// practically proportional to the sidereal time. The excess revolution time is
	// called length of day (LOD).
	Lod float64 `json:"lod"`
	// The estimated uncertainty/error in the lod value in seconds.
	LodUnc float64 `json:"lodUnc"`
	// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
	// (dPSI and dEpsilon).
	//
	// Any of "I", "P".
	NutationState EopAbridgedNutationState `json:"nutationState"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
	// values.
	//
	// Any of "I", "P".
	PolarMotionState EopAbridgedPolarMotionState `json:"polarMotionState"`
	// The Bulletin A X coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionX float64 `json:"polarMotionX"`
	// Bulletin B X coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionXb float64 `json:"polarMotionXB"`
	// Estimated uncertainty/error in polarMotionX value in arc seconds.
	PolarMotionXUnc float64 `json:"polarMotionXUnc"`
	// The Bulletin A Y coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionY float64 `json:"polarMotionY"`
	// Bulletin B Y coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionYb float64 `json:"polarMotionYB"`
	// Estimated uncertainty/error in polarMotionY value in arc seconds.
	PolarMotionYUnc float64 `json:"polarMotionYUnc"`
	// The IAU Theory of Precession and Theory of Nutation applied to the data in this
	// record. IAU1980 records employ the IAU 1976 Theory of Precession and IAU 1980
	// Theory of Nutation, and IAU2000 records employ the IAU 2000A Theory of
	// Precession and Nutation.
	PrecessionNutationStd string `json:"precessionNutationStd"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The difference between the Bulletin A UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utc float64 `json:"ut1UTC"`
	// The Bulletin B difference between the UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utcb float64 `json:"ut1UTCB"`
	// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
	// UT1-UTC values.
	//
	// Any of "I", "P".
	Ut1UtcState EopAbridgedUt1UtcState `json:"ut1UTCState"`
	// The estimated uncertainty/error in the ut1UTC value in seconds.
	Ut1UtcUnc float64 `json:"ut1UTCUnc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		EopDate               respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DEpsilon              respjson.Field
		DEpsilonB             respjson.Field
		DEpsilonUnc           respjson.Field
		DPsi                  respjson.Field
		DPsib                 respjson.Field
		DPsiUnc               respjson.Field
		DX                    respjson.Field
		DXb                   respjson.Field
		DXUnc                 respjson.Field
		DY                    respjson.Field
		DYb                   respjson.Field
		DYUnc                 respjson.Field
		Lod                   respjson.Field
		LodUnc                respjson.Field
		NutationState         respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PolarMotionState      respjson.Field
		PolarMotionX          respjson.Field
		PolarMotionXb         respjson.Field
		PolarMotionXUnc       respjson.Field
		PolarMotionY          respjson.Field
		PolarMotionYb         respjson.Field
		PolarMotionYUnc       respjson.Field
		PrecessionNutationStd respjson.Field
		RawFileUri            respjson.Field
		Ut1Utc                respjson.Field
		Ut1Utcb               respjson.Field
		Ut1UtcState           respjson.Field
		Ut1UtcUnc             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EopAbridged) RawJSON() string { return r.JSON.raw }
func (r *EopAbridged) UnmarshalJSON(data []byte) error {
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
type EopAbridgedDataMode string

const (
	EopAbridgedDataModeReal      EopAbridgedDataMode = "REAL"
	EopAbridgedDataModeTest      EopAbridgedDataMode = "TEST"
	EopAbridgedDataModeSimulated EopAbridgedDataMode = "SIMULATED"
	EopAbridgedDataModeExercise  EopAbridgedDataMode = "EXERCISE"
)

// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
// (dPSI and dEpsilon).
type EopAbridgedNutationState string

const (
	EopAbridgedNutationStateI EopAbridgedNutationState = "I"
	EopAbridgedNutationStateP EopAbridgedNutationState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
// values.
type EopAbridgedPolarMotionState string

const (
	EopAbridgedPolarMotionStateI EopAbridgedPolarMotionState = "I"
	EopAbridgedPolarMotionStateP EopAbridgedPolarMotionState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
// UT1-UTC values.
type EopAbridgedUt1UtcState string

const (
	EopAbridgedUt1UtcStateI EopAbridgedUt1UtcState = "I"
	EopAbridgedUt1UtcStateP EopAbridgedUt1UtcState = "P"
)

type EopQueryhelpResponse struct {
	AodrSupported         bool                            `json:"aodrSupported"`
	ClassificationMarking string                          `json:"classificationMarking"`
	Description           string                          `json:"description"`
	HistorySupported      bool                            `json:"historySupported"`
	Name                  string                          `json:"name"`
	Parameters            []EopQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                        `json:"requiredRoles"`
	RestSupported         bool                            `json:"restSupported"`
	SortSupported         bool                            `json:"sortSupported"`
	TypeName              string                          `json:"typeName"`
	Uri                   string                          `json:"uri"`
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
func (r EopQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *EopQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EopQueryhelpResponseParameter struct {
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
func (r EopQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *EopQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EopNewParams struct {
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
	DataMode EopNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted.
	EopDate time.Time `json:"eopDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Bulletin A offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilon is not used when this record represents IAU 2000 data.
	DEpsilon param.Opt[float64] `json:"dEpsilon,omitzero"`
	// The Bulletin B offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilonB is not used when this record represents IAU 2000 data.
	DEpsilonB param.Opt[float64] `json:"dEpsilonB,omitzero"`
	// The estimated uncertainty/error in the dEpsilon value in milliseconds of arc.
	DEpsilonUnc param.Opt[float64] `json:"dEpsilonUnc,omitzero"`
	// The Bulletin A offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSI is not used when this record represents IAU 2000 data.
	DPsi param.Opt[float64] `json:"dPSI,omitzero"`
	// The Bulletin B offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSIB is not used when this record represents IAU 2000 data.
	DPsib param.Opt[float64] `json:"dPSIB,omitzero"`
	// The estimated uncertainty/error in the dPSI value in milliseconds of arc.
	DPsiUnc param.Opt[float64] `json:"dPSIUnc,omitzero"`
	// The Bulletin A celestial pole offset along x-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dX is
	// not used when this record represents IAU 1980 data.
	DX param.Opt[float64] `json:"dX,omitzero"`
	// The Bulletin B celestial pole offset along the X-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dXB is not used when this record represents IAU 1980 data.
	DXb param.Opt[float64] `json:"dXB,omitzero"`
	// The estimated uncertainty/error in the Bulletin A dX value, in milliseconds of
	// arc.
	DXUnc param.Opt[float64] `json:"dXUnc,omitzero"`
	// The Bulletin A celestial pole offset along y-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dY is
	// not used when this record represents IAU 1980 data.
	DY param.Opt[float64] `json:"dY,omitzero"`
	// The Bulletin B celestial pole offset along the Y-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dYB is not used when this record represents IAU 1980 data.
	DYb param.Opt[float64] `json:"dYB,omitzero"`
	// The estimated uncertainty/error in the Bulletin A dY value, in milliseconds of
	// arc.
	DYUnc param.Opt[float64] `json:"dYUnc,omitzero"`
	// Bulletin A length of day or LOD in milliseconds. Universal time (UT1) is the
	// time of the earth clock, which performs one revolution in about 24h. It is
	// practically proportional to the sidereal time. The excess revolution time is
	// called length of day (LOD).
	Lod param.Opt[float64] `json:"lod,omitzero"`
	// The estimated uncertainty/error in the lod value in seconds.
	LodUnc param.Opt[float64] `json:"lodUnc,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Bulletin A X coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Bulletin B X coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionXb param.Opt[float64] `json:"polarMotionXB,omitzero"`
	// Estimated uncertainty/error in polarMotionX value in arc seconds.
	PolarMotionXUnc param.Opt[float64] `json:"polarMotionXUnc,omitzero"`
	// The Bulletin A Y coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// Bulletin B Y coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionYb param.Opt[float64] `json:"polarMotionYB,omitzero"`
	// Estimated uncertainty/error in polarMotionY value in arc seconds.
	PolarMotionYUnc param.Opt[float64] `json:"polarMotionYUnc,omitzero"`
	// The IAU Theory of Precession and Theory of Nutation applied to the data in this
	// record. IAU1980 records employ the IAU 1976 Theory of Precession and IAU 1980
	// Theory of Nutation, and IAU2000 records employ the IAU 2000A Theory of
	// Precession and Nutation.
	PrecessionNutationStd param.Opt[string] `json:"precessionNutationStd,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The difference between the Bulletin A UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utc param.Opt[float64] `json:"ut1UTC,omitzero"`
	// The Bulletin B difference between the UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utcb param.Opt[float64] `json:"ut1UTCB,omitzero"`
	// The estimated uncertainty/error in the ut1UTC value in seconds.
	Ut1UtcUnc param.Opt[float64] `json:"ut1UTCUnc,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
	// (dPSI and dEpsilon).
	//
	// Any of "I", "P".
	NutationState EopNewParamsNutationState `json:"nutationState,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
	// values.
	//
	// Any of "I", "P".
	PolarMotionState EopNewParamsPolarMotionState `json:"polarMotionState,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
	// UT1-UTC values.
	//
	// Any of "I", "P".
	Ut1UtcState EopNewParamsUt1UtcState `json:"ut1UTCState,omitzero"`
	paramObj
}

func (r EopNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EopNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EopNewParams) UnmarshalJSON(data []byte) error {
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
type EopNewParamsDataMode string

const (
	EopNewParamsDataModeReal      EopNewParamsDataMode = "REAL"
	EopNewParamsDataModeTest      EopNewParamsDataMode = "TEST"
	EopNewParamsDataModeSimulated EopNewParamsDataMode = "SIMULATED"
	EopNewParamsDataModeExercise  EopNewParamsDataMode = "EXERCISE"
)

// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
// (dPSI and dEpsilon).
type EopNewParamsNutationState string

const (
	EopNewParamsNutationStateI EopNewParamsNutationState = "I"
	EopNewParamsNutationStateP EopNewParamsNutationState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
// values.
type EopNewParamsPolarMotionState string

const (
	EopNewParamsPolarMotionStateI EopNewParamsPolarMotionState = "I"
	EopNewParamsPolarMotionStateP EopNewParamsPolarMotionState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
// UT1-UTC values.
type EopNewParamsUt1UtcState string

const (
	EopNewParamsUt1UtcStateI EopNewParamsUt1UtcState = "I"
	EopNewParamsUt1UtcStateP EopNewParamsUt1UtcState = "P"
)

type EopGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopGetParams]'s query parameters as `url.Values`.
func (r EopGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EopUpdateParams struct {
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
	DataMode EopUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted.
	EopDate time.Time `json:"eopDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The Bulletin A offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilon is not used when this record represents IAU 2000 data.
	DEpsilon param.Opt[float64] `json:"dEpsilon,omitzero"`
	// The Bulletin B offset in obliquity dDe1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dEpsilonB is not used when this record represents IAU 2000 data.
	DEpsilonB param.Opt[float64] `json:"dEpsilonB,omitzero"`
	// The estimated uncertainty/error in the dEpsilon value in milliseconds of arc.
	DEpsilonUnc param.Opt[float64] `json:"dEpsilonUnc,omitzero"`
	// The Bulletin A offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSI is not used when this record represents IAU 2000 data.
	DPsi param.Opt[float64] `json:"dPSI,omitzero"`
	// The Bulletin B offset in longitude dDy1980 with respect to the IAU 1976 Theory
	// of Precession and the IAU 1980 Theory of Nutation, measured in milliseconds of
	// arc. Note: dPSIB is not used when this record represents IAU 2000 data.
	DPsib param.Opt[float64] `json:"dPSIB,omitzero"`
	// The estimated uncertainty/error in the dPSI value in milliseconds of arc.
	DPsiUnc param.Opt[float64] `json:"dPSIUnc,omitzero"`
	// The Bulletin A celestial pole offset along x-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dX is
	// not used when this record represents IAU 1980 data.
	DX param.Opt[float64] `json:"dX,omitzero"`
	// The Bulletin B celestial pole offset along the X-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dXB is not used when this record represents IAU 1980 data.
	DXb param.Opt[float64] `json:"dXB,omitzero"`
	// The estimated uncertainty/error in the Bulletin A dX value, in milliseconds of
	// arc.
	DXUnc param.Opt[float64] `json:"dXUnc,omitzero"`
	// The Bulletin A celestial pole offset along y-axis with respect to the IAU 2000A
	// Theory of Precession and Nutation, measured in milliseconds of arc. Note: dY is
	// not used when this record represents IAU 1980 data.
	DY param.Opt[float64] `json:"dY,omitzero"`
	// The Bulletin B celestial pole offset along the Y-axis with respect to the IAU
	// 2000A Theory of Precession and Nutation, measured in milliseconds of arc. Note:
	// dYB is not used when this record represents IAU 1980 data.
	DYb param.Opt[float64] `json:"dYB,omitzero"`
	// The estimated uncertainty/error in the Bulletin A dY value, in milliseconds of
	// arc.
	DYUnc param.Opt[float64] `json:"dYUnc,omitzero"`
	// Bulletin A length of day or LOD in milliseconds. Universal time (UT1) is the
	// time of the earth clock, which performs one revolution in about 24h. It is
	// practically proportional to the sidereal time. The excess revolution time is
	// called length of day (LOD).
	Lod param.Opt[float64] `json:"lod,omitzero"`
	// The estimated uncertainty/error in the lod value in seconds.
	LodUnc param.Opt[float64] `json:"lodUnc,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Bulletin A X coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Bulletin B X coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionXb param.Opt[float64] `json:"polarMotionXB,omitzero"`
	// Estimated uncertainty/error in polarMotionX value in arc seconds.
	PolarMotionXUnc param.Opt[float64] `json:"polarMotionXUnc,omitzero"`
	// The Bulletin A Y coordinate value of earth polar motion at eopDate. Polar motion
	// of the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// Bulletin B Y coordinate value of earth polar motion at eopDate. Polar motion of
	// the Earth is the motion of the Earth's rotational axis relative to its crust.
	// This is measured with respect to a reference frame in which the solid Earth is
	// fixed (a so-called Earth-centered, Earth-fixed or ECEF reference frame).
	// Measured in arc seconds.
	PolarMotionYb param.Opt[float64] `json:"polarMotionYB,omitzero"`
	// Estimated uncertainty/error in polarMotionY value in arc seconds.
	PolarMotionYUnc param.Opt[float64] `json:"polarMotionYUnc,omitzero"`
	// The IAU Theory of Precession and Theory of Nutation applied to the data in this
	// record. IAU1980 records employ the IAU 1976 Theory of Precession and IAU 1980
	// Theory of Nutation, and IAU2000 records employ the IAU 2000A Theory of
	// Precession and Nutation.
	PrecessionNutationStd param.Opt[string] `json:"precessionNutationStd,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The difference between the Bulletin A UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utc param.Opt[float64] `json:"ut1UTC,omitzero"`
	// The Bulletin B difference between the UT1 and UTC time scales as of eopDate in
	// seconds.
	Ut1Utcb param.Opt[float64] `json:"ut1UTCB,omitzero"`
	// The estimated uncertainty/error in the ut1UTC value in seconds.
	Ut1UtcUnc param.Opt[float64] `json:"ut1UTCUnc,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
	// (dPSI and dEpsilon).
	//
	// Any of "I", "P".
	NutationState EopUpdateParamsNutationState `json:"nutationState,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
	// values.
	//
	// Any of "I", "P".
	PolarMotionState EopUpdateParamsPolarMotionState `json:"polarMotionState,omitzero"`
	// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
	// UT1-UTC values.
	//
	// Any of "I", "P".
	Ut1UtcState EopUpdateParamsUt1UtcState `json:"ut1UTCState,omitzero"`
	paramObj
}

func (r EopUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EopUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EopUpdateParams) UnmarshalJSON(data []byte) error {
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
type EopUpdateParamsDataMode string

const (
	EopUpdateParamsDataModeReal      EopUpdateParamsDataMode = "REAL"
	EopUpdateParamsDataModeTest      EopUpdateParamsDataMode = "TEST"
	EopUpdateParamsDataModeSimulated EopUpdateParamsDataMode = "SIMULATED"
	EopUpdateParamsDataModeExercise  EopUpdateParamsDataMode = "EXERCISE"
)

// Flag indicating Issued (I), or Predicted (P) for this record's nutation values
// (dPSI and dEpsilon).
type EopUpdateParamsNutationState string

const (
	EopUpdateParamsNutationStateI EopUpdateParamsNutationState = "I"
	EopUpdateParamsNutationStateP EopUpdateParamsNutationState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record's polar motion
// values.
type EopUpdateParamsPolarMotionState string

const (
	EopUpdateParamsPolarMotionStateI EopUpdateParamsPolarMotionState = "I"
	EopUpdateParamsPolarMotionStateP EopUpdateParamsPolarMotionState = "P"
)

// Flag indicating Issued (I), or Predicted (P) for this record”s Bulletin A
// UT1-UTC values.
type EopUpdateParamsUt1UtcState string

const (
	EopUpdateParamsUt1UtcStateI EopUpdateParamsUt1UtcState = "I"
	EopUpdateParamsUt1UtcStateP EopUpdateParamsUt1UtcState = "P"
)

type EopListParams struct {
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate     time.Time        `query:"eopDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopListParams]'s query parameters as `url.Values`.
func (r EopListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EopCountParams struct {
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate     time.Time        `query:"eopDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopCountParams]'s query parameters as `url.Values`.
func (r EopCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EopListTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate     time.Time        `query:"eopDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopListTupleParams]'s query parameters as `url.Values`.
func (r EopListTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
