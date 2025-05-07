// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// StarCatalogService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarCatalogService] method instead.
type StarCatalogService struct {
	Options []option.RequestOption
	History StarCatalogHistoryService
}

// NewStarCatalogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStarCatalogService(opts ...option.RequestOption) (r StarCatalogService) {
	r = StarCatalogService{}
	r.Options = opts
	r.History = NewStarCatalogHistoryService(opts...)
	return
}

// Service operation to take a single StarCatalog record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *StarCatalogService) New(ctx context.Context, body StarCatalogNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single starcatalog record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *StarCatalogService) Update(ctx context.Context, id string, body StarCatalogUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *StarCatalogService) List(ctx context.Context, query StarCatalogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StarCatalogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/starcatalog"
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
func (r *StarCatalogService) ListAutoPaging(ctx context.Context, query StarCatalogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StarCatalogListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a dataset specified by the passed ID path parameter.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *StarCatalogService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *StarCatalogService) Count(ctx context.Context, query StarCatalogCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/starcatalog/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// StarCatalog records as a POST body and ingest into the database. This operation
// is not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *StarCatalogService) NewBulk(ctx context.Context, body StarCatalogNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/starcatalog/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single StarCatalog record by its unique ID passed as
// a path parameter.
func (r *StarCatalogService) Get(ctx context.Context, id string, query StarCatalogGetParams, opts ...option.RequestOption) (res *StarCatalogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *StarCatalogService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/starcatalog/queryhelp"
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
func (r *StarCatalogService) Tuple(ctx context.Context, query StarCatalogTupleParams, opts ...option.RequestOption) (res *[]StarCatalogTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/starcatalog/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple StarCatalog records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StarCatalogService) UnvalidatedPublish(ctx context.Context, body StarCatalogUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogListResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarCatalogListResponseAstrometryOrigin `json:"astrometryOrigin,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	DataMode StarCatalogListResponseDataMode `json:"dataMode,required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag float64 `json:"bpmag"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag float64 `json:"gmag"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID int64 `json:"gncCatId"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID int64 `json:"hipCatId"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag float64 `json:"hmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc float64 `json:"hmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag float64 `json:"jmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc float64 `json:"jmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag float64 `json:"kmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc float64 `json:"kmagUnc"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag float64 `json:"rpmag"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc float64 `json:"rpmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Time the row was updated in the database.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin      respjson.Field
		ClassificationMarking respjson.Field
		CsID                  respjson.Field
		DataMode              respjson.Field
		Dec                   respjson.Field
		Ra                    respjson.Field
		Source                respjson.Field
		StarEpoch             respjson.Field
		ID                    respjson.Field
		Bpmag                 respjson.Field
		BpmagUnc              respjson.Field
		CatVersion            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecUnc                respjson.Field
		Gaiadr3CatID          respjson.Field
		Gmag                  respjson.Field
		GmagUnc               respjson.Field
		GncCatID              respjson.Field
		HipCatID              respjson.Field
		Hmag                  respjson.Field
		HmagUnc               respjson.Field
		Jmag                  respjson.Field
		JmagUnc               respjson.Field
		Kmag                  respjson.Field
		KmagUnc               respjson.Field
		MultFlag              respjson.Field
		NeighborDistance      respjson.Field
		NeighborFlag          respjson.Field
		NeighborID            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Parallax              respjson.Field
		ParallaxUnc           respjson.Field
		Pmdec                 respjson.Field
		PmdecUnc              respjson.Field
		Pmra                  respjson.Field
		PmraUnc               respjson.Field
		PmUncFlag             respjson.Field
		PosUncFlag            respjson.Field
		RaUnc                 respjson.Field
		Rpmag                 respjson.Field
		RpmagUnc              respjson.Field
		Shift                 respjson.Field
		ShiftFlag             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VarFlag               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogListResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarCatalogListResponseAstrometryOrigin string

const (
	StarCatalogListResponseAstrometryOriginGaiadr3   StarCatalogListResponseAstrometryOrigin = "GAIADR3"
	StarCatalogListResponseAstrometryOriginHipparcos StarCatalogListResponseAstrometryOrigin = "HIPPARCOS"
	StarCatalogListResponseAstrometryOriginUsnobsc   StarCatalogListResponseAstrometryOrigin = "USNOBSC"
)

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
type StarCatalogListResponseDataMode string

const (
	StarCatalogListResponseDataModeReal      StarCatalogListResponseDataMode = "REAL"
	StarCatalogListResponseDataModeTest      StarCatalogListResponseDataMode = "TEST"
	StarCatalogListResponseDataModeSimulated StarCatalogListResponseDataMode = "SIMULATED"
	StarCatalogListResponseDataModeExercise  StarCatalogListResponseDataMode = "EXERCISE"
)

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogGetResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarCatalogGetResponseAstrometryOrigin `json:"astrometryOrigin,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	DataMode StarCatalogGetResponseDataMode `json:"dataMode,required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag float64 `json:"bpmag"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag float64 `json:"gmag"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID int64 `json:"gncCatId"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID int64 `json:"hipCatId"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag float64 `json:"hmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc float64 `json:"hmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag float64 `json:"jmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc float64 `json:"jmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag float64 `json:"kmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc float64 `json:"kmagUnc"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag float64 `json:"rpmag"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc float64 `json:"rpmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Time the row was updated in the database.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin      respjson.Field
		ClassificationMarking respjson.Field
		CsID                  respjson.Field
		DataMode              respjson.Field
		Dec                   respjson.Field
		Ra                    respjson.Field
		Source                respjson.Field
		StarEpoch             respjson.Field
		ID                    respjson.Field
		Bpmag                 respjson.Field
		BpmagUnc              respjson.Field
		CatVersion            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecUnc                respjson.Field
		Gaiadr3CatID          respjson.Field
		Gmag                  respjson.Field
		GmagUnc               respjson.Field
		GncCatID              respjson.Field
		HipCatID              respjson.Field
		Hmag                  respjson.Field
		HmagUnc               respjson.Field
		Jmag                  respjson.Field
		JmagUnc               respjson.Field
		Kmag                  respjson.Field
		KmagUnc               respjson.Field
		MultFlag              respjson.Field
		NeighborDistance      respjson.Field
		NeighborFlag          respjson.Field
		NeighborID            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Parallax              respjson.Field
		ParallaxUnc           respjson.Field
		Pmdec                 respjson.Field
		PmdecUnc              respjson.Field
		Pmra                  respjson.Field
		PmraUnc               respjson.Field
		PmUncFlag             respjson.Field
		PosUncFlag            respjson.Field
		RaUnc                 respjson.Field
		Rpmag                 respjson.Field
		RpmagUnc              respjson.Field
		Shift                 respjson.Field
		ShiftFlag             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VarFlag               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarCatalogGetResponseAstrometryOrigin string

const (
	StarCatalogGetResponseAstrometryOriginGaiadr3   StarCatalogGetResponseAstrometryOrigin = "GAIADR3"
	StarCatalogGetResponseAstrometryOriginHipparcos StarCatalogGetResponseAstrometryOrigin = "HIPPARCOS"
	StarCatalogGetResponseAstrometryOriginUsnobsc   StarCatalogGetResponseAstrometryOrigin = "USNOBSC"
)

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
type StarCatalogGetResponseDataMode string

const (
	StarCatalogGetResponseDataModeReal      StarCatalogGetResponseDataMode = "REAL"
	StarCatalogGetResponseDataModeTest      StarCatalogGetResponseDataMode = "TEST"
	StarCatalogGetResponseDataModeSimulated StarCatalogGetResponseDataMode = "SIMULATED"
	StarCatalogGetResponseDataModeExercise  StarCatalogGetResponseDataMode = "EXERCISE"
)

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogTupleResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarCatalogTupleResponseAstrometryOrigin `json:"astrometryOrigin,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	DataMode StarCatalogTupleResponseDataMode `json:"dataMode,required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag float64 `json:"bpmag"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag float64 `json:"gmag"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID int64 `json:"gncCatId"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID int64 `json:"hipCatId"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag float64 `json:"hmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc float64 `json:"hmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag float64 `json:"jmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc float64 `json:"jmagUnc"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag float64 `json:"kmag"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc float64 `json:"kmagUnc"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag float64 `json:"rpmag"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc float64 `json:"rpmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Time the row was updated in the database.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin      respjson.Field
		ClassificationMarking respjson.Field
		CsID                  respjson.Field
		DataMode              respjson.Field
		Dec                   respjson.Field
		Ra                    respjson.Field
		Source                respjson.Field
		StarEpoch             respjson.Field
		ID                    respjson.Field
		Bpmag                 respjson.Field
		BpmagUnc              respjson.Field
		CatVersion            respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DecUnc                respjson.Field
		Gaiadr3CatID          respjson.Field
		Gmag                  respjson.Field
		GmagUnc               respjson.Field
		GncCatID              respjson.Field
		HipCatID              respjson.Field
		Hmag                  respjson.Field
		HmagUnc               respjson.Field
		Jmag                  respjson.Field
		JmagUnc               respjson.Field
		Kmag                  respjson.Field
		KmagUnc               respjson.Field
		MultFlag              respjson.Field
		NeighborDistance      respjson.Field
		NeighborFlag          respjson.Field
		NeighborID            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Parallax              respjson.Field
		ParallaxUnc           respjson.Field
		Pmdec                 respjson.Field
		PmdecUnc              respjson.Field
		Pmra                  respjson.Field
		PmraUnc               respjson.Field
		PmUncFlag             respjson.Field
		PosUncFlag            respjson.Field
		RaUnc                 respjson.Field
		Rpmag                 respjson.Field
		RpmagUnc              respjson.Field
		Shift                 respjson.Field
		ShiftFlag             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VarFlag               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarCatalogTupleResponseAstrometryOrigin string

const (
	StarCatalogTupleResponseAstrometryOriginGaiadr3   StarCatalogTupleResponseAstrometryOrigin = "GAIADR3"
	StarCatalogTupleResponseAstrometryOriginHipparcos StarCatalogTupleResponseAstrometryOrigin = "HIPPARCOS"
	StarCatalogTupleResponseAstrometryOriginUsnobsc   StarCatalogTupleResponseAstrometryOrigin = "USNOBSC"
)

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
type StarCatalogTupleResponseDataMode string

const (
	StarCatalogTupleResponseDataModeReal      StarCatalogTupleResponseDataMode = "REAL"
	StarCatalogTupleResponseDataModeTest      StarCatalogTupleResponseDataMode = "TEST"
	StarCatalogTupleResponseDataModeSimulated StarCatalogTupleResponseDataMode = "SIMULATED"
	StarCatalogTupleResponseDataModeExercise  StarCatalogTupleResponseDataMode = "EXERCISE"
)

type StarCatalogNewParams struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarCatalogNewParamsAstrometryOrigin `json:"astrometryOrigin,omitzero,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	DataMode StarCatalogNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	paramObj
}

func (r StarCatalogNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarCatalogNewParamsAstrometryOrigin string

const (
	StarCatalogNewParamsAstrometryOriginGaiadr3   StarCatalogNewParamsAstrometryOrigin = "GAIADR3"
	StarCatalogNewParamsAstrometryOriginHipparcos StarCatalogNewParamsAstrometryOrigin = "HIPPARCOS"
	StarCatalogNewParamsAstrometryOriginUsnobsc   StarCatalogNewParamsAstrometryOrigin = "USNOBSC"
)

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
type StarCatalogNewParamsDataMode string

const (
	StarCatalogNewParamsDataModeReal      StarCatalogNewParamsDataMode = "REAL"
	StarCatalogNewParamsDataModeTest      StarCatalogNewParamsDataMode = "TEST"
	StarCatalogNewParamsDataModeSimulated StarCatalogNewParamsDataMode = "SIMULATED"
	StarCatalogNewParamsDataModeExercise  StarCatalogNewParamsDataMode = "EXERCISE"
)

type StarCatalogUpdateParams struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarCatalogUpdateParamsAstrometryOrigin `json:"astrometryOrigin,omitzero,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	DataMode StarCatalogUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	paramObj
}

func (r StarCatalogUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarCatalogUpdateParamsAstrometryOrigin string

const (
	StarCatalogUpdateParamsAstrometryOriginGaiadr3   StarCatalogUpdateParamsAstrometryOrigin = "GAIADR3"
	StarCatalogUpdateParamsAstrometryOriginHipparcos StarCatalogUpdateParamsAstrometryOrigin = "HIPPARCOS"
	StarCatalogUpdateParamsAstrometryOriginUsnobsc   StarCatalogUpdateParamsAstrometryOrigin = "USNOBSC"
)

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
type StarCatalogUpdateParamsDataMode string

const (
	StarCatalogUpdateParamsDataModeReal      StarCatalogUpdateParamsDataMode = "REAL"
	StarCatalogUpdateParamsDataModeTest      StarCatalogUpdateParamsDataMode = "TEST"
	StarCatalogUpdateParamsDataModeSimulated StarCatalogUpdateParamsDataMode = "SIMULATED"
	StarCatalogUpdateParamsDataModeExercise  StarCatalogUpdateParamsDataMode = "EXERCISE"
)

type StarCatalogListParams struct {
	// (One or more of fields 'dec, ra' are required.) Barycentric declination of the
	// source in International Celestial Reference System (ICRS) at the reference
	// epoch, in degrees.
	Dec         param.Opt[float64] `query:"dec,omitzero" json:"-"`
	FirstResult param.Opt[int64]   `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]   `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'dec, ra' are required.) Barycentric right ascension of
	// the source in the International Celestial Reference System (ICRS) frame at the
	// reference epoch, in degrees.
	Ra param.Opt[float64] `query:"ra,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StarCatalogListParams]'s query parameters as `url.Values`.
func (r StarCatalogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarCatalogCountParams struct {
	// (One or more of fields 'dec, ra' are required.) Barycentric declination of the
	// source in International Celestial Reference System (ICRS) at the reference
	// epoch, in degrees.
	Dec         param.Opt[float64] `query:"dec,omitzero" json:"-"`
	FirstResult param.Opt[int64]   `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]   `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'dec, ra' are required.) Barycentric right ascension of
	// the source in the International Celestial Reference System (ICRS) frame at the
	// reference epoch, in degrees.
	Ra param.Opt[float64] `query:"ra,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StarCatalogCountParams]'s query parameters as `url.Values`.
func (r StarCatalogCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarCatalogNewBulkParams struct {
	Body []StarCatalogNewBulkParamsBody
	paramObj
}

func (r StarCatalogNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *StarCatalogNewBulkParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarCatalogNewBulkParamsBody struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin string `json:"astrometryOrigin,omitzero,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Time the row was updated in the database.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	paramObj
}

func (r StarCatalogNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StarCatalogNewBulkParamsBody](
		"astrometryOrigin", "GAIADR3", "HIPPARCOS", "USNOBSC",
	)
	apijson.RegisterFieldValidator[StarCatalogNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type StarCatalogGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StarCatalogGetParams]'s query parameters as `url.Values`.
func (r StarCatalogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarCatalogTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// (One or more of fields 'dec, ra' are required.) Barycentric declination of the
	// source in International Celestial Reference System (ICRS) at the reference
	// epoch, in degrees.
	Dec         param.Opt[float64] `query:"dec,omitzero" json:"-"`
	FirstResult param.Opt[int64]   `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]   `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'dec, ra' are required.) Barycentric right ascension of
	// the source in the International Celestial Reference System (ICRS) frame at the
	// reference epoch, in degrees.
	Ra param.Opt[float64] `query:"ra,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StarCatalogTupleParams]'s query parameters as `url.Values`.
func (r StarCatalogTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarCatalogUnvalidatedPublishParams struct {
	Body []StarCatalogUnvalidatedPublishParamsBody
	paramObj
}

func (r StarCatalogUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *StarCatalogUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarCatalogUnvalidatedPublishParamsBody struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin string `json:"astrometryOrigin,omitzero,required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// The ID of this object in the specific catalog associated with this record.
	CsID int64 `json:"csId,required"`
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
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec,required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Gaia DR3 optical photometric Bp-band magnitude in the Vega scale.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia DR3 optical Bp-band magnitude uncertainty in the Vega scale.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The ID of this object in the Gaia DR3 Catalog.
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude in the Vega scale.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia DR3 optical photometric G-band magnitude uncertainty in the Vega scale.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navagation Control (GNC) Catalog.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The ID of this object in the Hipparcos Catalog.
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude in the Vega scale.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric H-band magnitude uncertainty in the Vega scale.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude in the Vega scale.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric J-band magnitude uncertainty in the Vega scale.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude in the Vega scale.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Two Micron All Sky Survey (2MASS) Point Source Catalog (PSC) near-infrared
	// photometric K-band magnitude uncertainty in the Vega scale.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds/year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds/year, at the
	// reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds/year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds/year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Gaia DR3 optical Rp-band magnitude in the Vega scale.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia DR3 optical photometric Rp-band magnitude uncertainty in the Vega scale.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Time the row was updated in the database.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Application user who updated the row in the database.
	UpdatedBy param.Opt[string] `json:"updatedBy,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	paramObj
}

func (r StarCatalogUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StarCatalogUnvalidatedPublishParamsBody](
		"astrometryOrigin", "GAIADR3", "HIPPARCOS", "USNOBSC",
	)
	apijson.RegisterFieldValidator[StarCatalogUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
