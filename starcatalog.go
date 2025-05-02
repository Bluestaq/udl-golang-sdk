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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// StarcatalogService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarcatalogService] method instead.
type StarcatalogService struct {
	Options []option.RequestOption
	History StarcatalogHistoryService
}

// NewStarcatalogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStarcatalogService(opts ...option.RequestOption) (r StarcatalogService) {
	r = StarcatalogService{}
	r.Options = opts
	r.History = NewStarcatalogHistoryService(opts...)
	return
}

// Service operation to take a single StarCatalog record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *StarcatalogService) New(ctx context.Context, body StarcatalogNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single starcatalog record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *StarcatalogService) Update(ctx context.Context, id string, body StarcatalogUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *StarcatalogService) List(ctx context.Context, query StarcatalogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StarcatalogListResponse], err error) {
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
func (r *StarcatalogService) ListAutoPaging(ctx context.Context, query StarcatalogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StarcatalogListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a dataset specified by the passed ID path parameter.
// A specific role is required to perform this service operation. Please contact
// the UDL team for assistance.
func (r *StarcatalogService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *StarcatalogService) Count(ctx context.Context, query StarcatalogCountParams, opts ...option.RequestOption) (res *string, err error) {
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
func (r *StarcatalogService) NewBulk(ctx context.Context, body StarcatalogNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/starcatalog/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single StarCatalog record by its unique ID passed as
// a path parameter.
func (r *StarcatalogService) Get(ctx context.Context, id string, query StarcatalogGetParams, opts ...option.RequestOption) (res *StarcatalogGetResponse, err error) {
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
func (r *StarcatalogService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *StarcatalogService) Tuple(ctx context.Context, query StarcatalogTupleParams, opts ...option.RequestOption) (res *[]StarcatalogTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/starcatalog/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple StarCatalog records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StarcatalogService) UnvalidatedPublish(ctx context.Context, body StarcatalogUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarcatalogListResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarcatalogListResponseAstrometryOrigin `json:"astrometryOrigin,required"`
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
	DataMode StarcatalogListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AstrometryOrigin      resp.Field
		ClassificationMarking resp.Field
		CsID                  resp.Field
		DataMode              resp.Field
		Dec                   resp.Field
		Ra                    resp.Field
		Source                resp.Field
		StarEpoch             resp.Field
		ID                    resp.Field
		Bpmag                 resp.Field
		BpmagUnc              resp.Field
		CatVersion            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DecUnc                resp.Field
		Gaiadr3CatID          resp.Field
		Gmag                  resp.Field
		GmagUnc               resp.Field
		GncCatID              resp.Field
		HipCatID              resp.Field
		Hmag                  resp.Field
		HmagUnc               resp.Field
		Jmag                  resp.Field
		JmagUnc               resp.Field
		Kmag                  resp.Field
		KmagUnc               resp.Field
		MultFlag              resp.Field
		NeighborDistance      resp.Field
		NeighborFlag          resp.Field
		NeighborID            resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Parallax              resp.Field
		ParallaxUnc           resp.Field
		Pmdec                 resp.Field
		PmdecUnc              resp.Field
		Pmra                  resp.Field
		PmraUnc               resp.Field
		PmUncFlag             resp.Field
		PosUncFlag            resp.Field
		RaUnc                 resp.Field
		Rpmag                 resp.Field
		RpmagUnc              resp.Field
		Shift                 resp.Field
		ShiftFlag             resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		VarFlag               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarcatalogListResponse) RawJSON() string { return r.JSON.raw }
func (r *StarcatalogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarcatalogListResponseAstrometryOrigin string

const (
	StarcatalogListResponseAstrometryOriginGaiadr3   StarcatalogListResponseAstrometryOrigin = "GAIADR3"
	StarcatalogListResponseAstrometryOriginHipparcos StarcatalogListResponseAstrometryOrigin = "HIPPARCOS"
	StarcatalogListResponseAstrometryOriginUsnobsc   StarcatalogListResponseAstrometryOrigin = "USNOBSC"
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
type StarcatalogListResponseDataMode string

const (
	StarcatalogListResponseDataModeReal      StarcatalogListResponseDataMode = "REAL"
	StarcatalogListResponseDataModeTest      StarcatalogListResponseDataMode = "TEST"
	StarcatalogListResponseDataModeSimulated StarcatalogListResponseDataMode = "SIMULATED"
	StarcatalogListResponseDataModeExercise  StarcatalogListResponseDataMode = "EXERCISE"
)

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarcatalogGetResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarcatalogGetResponseAstrometryOrigin `json:"astrometryOrigin,required"`
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
	DataMode StarcatalogGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AstrometryOrigin      resp.Field
		ClassificationMarking resp.Field
		CsID                  resp.Field
		DataMode              resp.Field
		Dec                   resp.Field
		Ra                    resp.Field
		Source                resp.Field
		StarEpoch             resp.Field
		ID                    resp.Field
		Bpmag                 resp.Field
		BpmagUnc              resp.Field
		CatVersion            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DecUnc                resp.Field
		Gaiadr3CatID          resp.Field
		Gmag                  resp.Field
		GmagUnc               resp.Field
		GncCatID              resp.Field
		HipCatID              resp.Field
		Hmag                  resp.Field
		HmagUnc               resp.Field
		Jmag                  resp.Field
		JmagUnc               resp.Field
		Kmag                  resp.Field
		KmagUnc               resp.Field
		MultFlag              resp.Field
		NeighborDistance      resp.Field
		NeighborFlag          resp.Field
		NeighborID            resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Parallax              resp.Field
		ParallaxUnc           resp.Field
		Pmdec                 resp.Field
		PmdecUnc              resp.Field
		Pmra                  resp.Field
		PmraUnc               resp.Field
		PmUncFlag             resp.Field
		PosUncFlag            resp.Field
		RaUnc                 resp.Field
		Rpmag                 resp.Field
		RpmagUnc              resp.Field
		Shift                 resp.Field
		ShiftFlag             resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		VarFlag               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarcatalogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StarcatalogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarcatalogGetResponseAstrometryOrigin string

const (
	StarcatalogGetResponseAstrometryOriginGaiadr3   StarcatalogGetResponseAstrometryOrigin = "GAIADR3"
	StarcatalogGetResponseAstrometryOriginHipparcos StarcatalogGetResponseAstrometryOrigin = "HIPPARCOS"
	StarcatalogGetResponseAstrometryOriginUsnobsc   StarcatalogGetResponseAstrometryOrigin = "USNOBSC"
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
type StarcatalogGetResponseDataMode string

const (
	StarcatalogGetResponseDataModeReal      StarcatalogGetResponseDataMode = "REAL"
	StarcatalogGetResponseDataModeTest      StarcatalogGetResponseDataMode = "TEST"
	StarcatalogGetResponseDataModeSimulated StarcatalogGetResponseDataMode = "SIMULATED"
	StarcatalogGetResponseDataModeExercise  StarcatalogGetResponseDataMode = "EXERCISE"
)

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarcatalogTupleResponse struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarcatalogTupleResponseAstrometryOrigin `json:"astrometryOrigin,required"`
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
	DataMode StarcatalogTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AstrometryOrigin      resp.Field
		ClassificationMarking resp.Field
		CsID                  resp.Field
		DataMode              resp.Field
		Dec                   resp.Field
		Ra                    resp.Field
		Source                resp.Field
		StarEpoch             resp.Field
		ID                    resp.Field
		Bpmag                 resp.Field
		BpmagUnc              resp.Field
		CatVersion            resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		DecUnc                resp.Field
		Gaiadr3CatID          resp.Field
		Gmag                  resp.Field
		GmagUnc               resp.Field
		GncCatID              resp.Field
		HipCatID              resp.Field
		Hmag                  resp.Field
		HmagUnc               resp.Field
		Jmag                  resp.Field
		JmagUnc               resp.Field
		Kmag                  resp.Field
		KmagUnc               resp.Field
		MultFlag              resp.Field
		NeighborDistance      resp.Field
		NeighborFlag          resp.Field
		NeighborID            resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Parallax              resp.Field
		ParallaxUnc           resp.Field
		Pmdec                 resp.Field
		PmdecUnc              resp.Field
		Pmra                  resp.Field
		PmraUnc               resp.Field
		PmUncFlag             resp.Field
		PosUncFlag            resp.Field
		RaUnc                 resp.Field
		Rpmag                 resp.Field
		RpmagUnc              resp.Field
		Shift                 resp.Field
		ShiftFlag             resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		VarFlag               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarcatalogTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *StarcatalogTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarcatalogTupleResponseAstrometryOrigin string

const (
	StarcatalogTupleResponseAstrometryOriginGaiadr3   StarcatalogTupleResponseAstrometryOrigin = "GAIADR3"
	StarcatalogTupleResponseAstrometryOriginHipparcos StarcatalogTupleResponseAstrometryOrigin = "HIPPARCOS"
	StarcatalogTupleResponseAstrometryOriginUsnobsc   StarcatalogTupleResponseAstrometryOrigin = "USNOBSC"
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
type StarcatalogTupleResponseDataMode string

const (
	StarcatalogTupleResponseDataModeReal      StarcatalogTupleResponseDataMode = "REAL"
	StarcatalogTupleResponseDataModeTest      StarcatalogTupleResponseDataMode = "TEST"
	StarcatalogTupleResponseDataModeSimulated StarcatalogTupleResponseDataMode = "SIMULATED"
	StarcatalogTupleResponseDataModeExercise  StarcatalogTupleResponseDataMode = "EXERCISE"
)

type StarcatalogNewParams struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarcatalogNewParamsAstrometryOrigin `json:"astrometryOrigin,omitzero,required"`
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
	DataMode StarcatalogNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r StarcatalogNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StarcatalogNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarcatalogNewParamsAstrometryOrigin string

const (
	StarcatalogNewParamsAstrometryOriginGaiadr3   StarcatalogNewParamsAstrometryOrigin = "GAIADR3"
	StarcatalogNewParamsAstrometryOriginHipparcos StarcatalogNewParamsAstrometryOrigin = "HIPPARCOS"
	StarcatalogNewParamsAstrometryOriginUsnobsc   StarcatalogNewParamsAstrometryOrigin = "USNOBSC"
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
type StarcatalogNewParamsDataMode string

const (
	StarcatalogNewParamsDataModeReal      StarcatalogNewParamsDataMode = "REAL"
	StarcatalogNewParamsDataModeTest      StarcatalogNewParamsDataMode = "TEST"
	StarcatalogNewParamsDataModeSimulated StarcatalogNewParamsDataMode = "SIMULATED"
	StarcatalogNewParamsDataModeExercise  StarcatalogNewParamsDataMode = "EXERCISE"
)

type StarcatalogUpdateParams struct {
	// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
	// USNOBSC].
	//
	// Any of "GAIADR3", "HIPPARCOS", "USNOBSC".
	AstrometryOrigin StarcatalogUpdateParamsAstrometryOrigin `json:"astrometryOrigin,omitzero,required"`
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
	DataMode StarcatalogUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r StarcatalogUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StarcatalogUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Originating astrometric catalog for this object. Enum: [GAIADR3, HIPPARCOS,
// USNOBSC].
type StarcatalogUpdateParamsAstrometryOrigin string

const (
	StarcatalogUpdateParamsAstrometryOriginGaiadr3   StarcatalogUpdateParamsAstrometryOrigin = "GAIADR3"
	StarcatalogUpdateParamsAstrometryOriginHipparcos StarcatalogUpdateParamsAstrometryOrigin = "HIPPARCOS"
	StarcatalogUpdateParamsAstrometryOriginUsnobsc   StarcatalogUpdateParamsAstrometryOrigin = "USNOBSC"
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
type StarcatalogUpdateParamsDataMode string

const (
	StarcatalogUpdateParamsDataModeReal      StarcatalogUpdateParamsDataMode = "REAL"
	StarcatalogUpdateParamsDataModeTest      StarcatalogUpdateParamsDataMode = "TEST"
	StarcatalogUpdateParamsDataModeSimulated StarcatalogUpdateParamsDataMode = "SIMULATED"
	StarcatalogUpdateParamsDataModeExercise  StarcatalogUpdateParamsDataMode = "EXERCISE"
)

type StarcatalogListParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [StarcatalogListParams]'s query parameters as `url.Values`.
func (r StarcatalogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarcatalogCountParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [StarcatalogCountParams]'s query parameters as `url.Values`.
func (r StarcatalogCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarcatalogNewBulkParams struct {
	Body []StarcatalogNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r StarcatalogNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarcatalogNewBulkParamsBody struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r StarcatalogNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StarcatalogNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[StarcatalogNewBulkParamsBody](
		"AstrometryOrigin", false, "GAIADR3", "HIPPARCOS", "USNOBSC",
	)
	apijson.RegisterFieldValidator[StarcatalogNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type StarcatalogGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [StarcatalogGetParams]'s query parameters as `url.Values`.
func (r StarcatalogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarcatalogTupleParams struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [StarcatalogTupleParams]'s query parameters as `url.Values`.
func (r StarcatalogTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StarcatalogUnvalidatedPublishParams struct {
	Body []StarcatalogUnvalidatedPublishParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogUnvalidatedPublishParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r StarcatalogUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarcatalogUnvalidatedPublishParamsBody struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f StarcatalogUnvalidatedPublishParamsBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r StarcatalogUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StarcatalogUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[StarcatalogUnvalidatedPublishParamsBody](
		"AstrometryOrigin", false, "GAIADR3", "HIPPARCOS", "USNOBSC",
	)
	apijson.RegisterFieldValidator[StarcatalogUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
