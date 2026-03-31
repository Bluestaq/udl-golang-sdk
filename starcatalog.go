// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	shimjson "github.com/Bluestaq/udl-golang-sdk/internal/encoding/json"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// These services provide operations for posting and querying Star Catalog data.
// The Star Catalog model is a representation of astronomical data and photometric
// data for stars. Astronomical data includes positional information, proper
// motions, parallaxes and their respective uncertainties. Photometric data
// contains optical and near-infrared magnitudes, and their uncertainties across
// multiple bandpasses. Note: Multiple source catalogs may contribute to a single
// record.
//
// StarCatalogService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarCatalogService] method instead.
type StarCatalogService struct {
	Options []option.RequestOption
	// These services provide operations for posting and querying Star Catalog data.
	// The Star Catalog model is a representation of astronomical data and photometric
	// data for stars. Astronomical data includes positional information, proper
	// motions, parallaxes and their respective uncertainties. Photometric data
	// contains optical and near-infrared magnitudes, and their uncertainties across
	// multiple bandpasses. Note: Multiple source catalogs may contribute to a single
	// record.
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Service operation to update a single starcatalog record. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *StarCatalogService) Update(ctx context.Context, id string, body StarCatalogUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *StarCatalogService) List(ctx context.Context, query StarCatalogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StarCatalogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *StarCatalogService) Count(ctx context.Context, query StarCatalogCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/starcatalog/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Service operation intended for initial integration only, to take a list of
// StarCatalog records as a POST body and ingest into the database. This operation
// is not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *StarCatalogService) NewBulk(ctx context.Context, body StarCatalogNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/starcatalog/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Service operation to get a single StarCatalog record by its unique ID passed as
// a path parameter.
func (r *StarCatalogService) Get(ctx context.Context, id string, query StarCatalogGetParams, opts ...option.RequestOption) (res *StarCatalogGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("udl/starcatalog/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *StarCatalogService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *StarCatalogQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/starcatalog/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	opts = slices.Concat(r.Options, opts)
	path := "udl/starcatalog/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Service operation to take multiple StarCatalog records as a POST body and ingest
// into the database. This operation is intended to be used for automated feeds
// into UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StarCatalogService) UnvalidatedPublish(ctx context.Context, body StarCatalogUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "filedrop/udl-starcatalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogListResponse struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode StarCatalogListResponseDataMode `json:"dataMode" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID int64 `json:"aavsoVsxId"`
	// Optical AB g magnitude.
	Abgmag float64 `json:"abgmag"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin string `json:"abgmagOrigin"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc float64 `json:"abgmagUnc"`
	// Optical AB i magnitude.
	Abimag float64 `json:"abimag"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin string `json:"abimagOrigin"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc float64 `json:"abimagUnc"`
	// Optical AB r magnitude.
	Abrmag float64 `json:"abrmag"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin string `json:"abrmagOrigin"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc float64 `json:"abrmagUnc"`
	// Optical AB y magnitude.
	Abymag float64 `json:"abymag"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin string `json:"abymagOrigin"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc float64 `json:"abymagUnc"`
	// Optical AB z magnitude.
	Abzmag float64 `json:"abzmag"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin string `json:"abzmagOrigin"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc float64 `json:"abzmagUnc"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd string `json:"allWISEccInd"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID string `json:"allWISEId"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd int64 `json:"allWISEnaInd"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd string `json:"allWISEphQualInd"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID string `json:"apassId"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise float64 `json:"astrometricExcessNoise"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig float64 `json:"astrometricExcessNoiseSig"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag float64 `json:"bmag"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin string `json:"bmagOrigin"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc float64 `json:"bmagUnc"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag float64 `json:"bpmag"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID int64 `json:"carrascoCatId"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID string `json:"catWise2020Id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID string `json:"ducatiCatId"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag float64 `json:"gmag"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID int64 `json:"gncCatId"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex int64 `json:"healpixIndex"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID int64 `json:"hipCatId"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag float64 `json:"hmag"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin string `json:"hmagOrigin"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc float64 `json:"hmagUnc"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag float64 `json:"imag"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin string `json:"imagOrigin"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc float64 `json:"imagUnc"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag float64 `json:"jmag"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin string `json:"jmagOrigin"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc float64 `json:"jmagUnc"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag float64 `json:"kmag"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin string `json:"kmagOrigin"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc float64 `json:"kmagUnc"`
	// Morphology indicator.
	MorphologyInd int64 `json:"morphologyInd"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity string `json:"multiplicity"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec float64 `json:"neighborDec"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa float64 `json:"neighborRa"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar string `json:"nonSingleStar"`
	// Number of neighbors.
	NumNeighbors int64 `json:"numNeighbors"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID int64 `json:"panStarrsId"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag int64 `json:"ps1astrometryCorrectionFlag"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag int64 `json:"ps1ObjInfoFlag"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag int64 `json:"ps1QualityFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag float64 `json:"rmag"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin string `json:"rmagOrigin"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc float64 `json:"rmagUnc"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag float64 `json:"rpmag"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc float64 `json:"rpmagUnc"`
	// RUWE in Gaia.
	Ruwe float64 `json:"ruwe"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID int64 `json:"sdaCatId"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag float64 `json:"sgmag"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc float64 `json:"sgmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 float64 `json:"shiftFWHM1"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 float64 `json:"shiftFWHM6"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID int64 `json:"skyMapperId"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID string `json:"twoMASSId"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd string `json:"twoMassPHQualInd"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag string `json:"twoMassReadFlag"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID string `json:"twoMassXscId"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID int64 `json:"tychoDscId"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID int64 `json:"uhsId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID int64 `json:"ukidssGCSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID int64 `json:"ukidssGPSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID int64 `json:"ukidssLASId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability string `json:"variability"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID int64 `json:"vhsId"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag float64 `json:"vmag"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin string `json:"vmagOrigin"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc float64 `json:"vmagUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag float64 `json:"w1mag"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin string `json:"w1magOrigin"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc float64 `json:"w1magUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat float64 `json:"w1sat"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag float64 `json:"w2mag"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin string `json:"w2magOrigin"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc float64 `json:"w2magUnc"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat float64 `json:"w2sat"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag float64 `json:"w3mag"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin string `json:"w3magOrigin"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc float64 `json:"w3magUnc"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat float64 `json:"w3sat"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag float64 `json:"w4mag"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin string `json:"w4magOrigin"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc float64 `json:"w4magUnc"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat float64 `json:"w4sat"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID string `json:"wdsCatId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin            respjson.Field
		ClassificationMarking       respjson.Field
		CsID                        respjson.Field
		DataMode                    respjson.Field
		Dec                         respjson.Field
		Ra                          respjson.Field
		Source                      respjson.Field
		StarEpoch                   respjson.Field
		ID                          respjson.Field
		AavsoVsxID                  respjson.Field
		Abgmag                      respjson.Field
		AbgmagOrigin                respjson.Field
		AbgmagUnc                   respjson.Field
		Abimag                      respjson.Field
		AbimagOrigin                respjson.Field
		AbimagUnc                   respjson.Field
		Abrmag                      respjson.Field
		AbrmagOrigin                respjson.Field
		AbrmagUnc                   respjson.Field
		Abymag                      respjson.Field
		AbymagOrigin                respjson.Field
		AbymagUnc                   respjson.Field
		Abzmag                      respjson.Field
		AbzmagOrigin                respjson.Field
		AbzmagUnc                   respjson.Field
		AllWisEccInd                respjson.Field
		AllWiseID                   respjson.Field
		AllWisEnaInd                respjson.Field
		AllWisEphQualInd            respjson.Field
		ApassID                     respjson.Field
		AstrometricExcessNoise      respjson.Field
		AstrometricExcessNoiseSig   respjson.Field
		Bmag                        respjson.Field
		BmagOrigin                  respjson.Field
		BmagUnc                     respjson.Field
		Bpmag                       respjson.Field
		BpmagUnc                    respjson.Field
		CarrascoCatID               respjson.Field
		CatVersion                  respjson.Field
		CatWise2020ID               respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		DecUnc                      respjson.Field
		DucatiCatID                 respjson.Field
		Gaiadr3CatID                respjson.Field
		Gmag                        respjson.Field
		GmagUnc                     respjson.Field
		GncCatID                    respjson.Field
		HealpixIndex                respjson.Field
		HipCatID                    respjson.Field
		Hmag                        respjson.Field
		HmagOrigin                  respjson.Field
		HmagUnc                     respjson.Field
		Imag                        respjson.Field
		ImagOrigin                  respjson.Field
		ImagUnc                     respjson.Field
		Jmag                        respjson.Field
		JmagOrigin                  respjson.Field
		JmagUnc                     respjson.Field
		Kmag                        respjson.Field
		KmagOrigin                  respjson.Field
		KmagUnc                     respjson.Field
		MorphologyInd               respjson.Field
		MultFlag                    respjson.Field
		Multiplicity                respjson.Field
		NeighborDec                 respjson.Field
		NeighborDistance            respjson.Field
		NeighborFlag                respjson.Field
		NeighborID                  respjson.Field
		NeighborRa                  respjson.Field
		NonSingleStar               respjson.Field
		NumNeighbors                respjson.Field
		Origin                      respjson.Field
		OrigNetwork                 respjson.Field
		PanStarrsID                 respjson.Field
		Parallax                    respjson.Field
		ParallaxUnc                 respjson.Field
		Pmdec                       respjson.Field
		PmdecUnc                    respjson.Field
		Pmra                        respjson.Field
		PmraUnc                     respjson.Field
		PmUncFlag                   respjson.Field
		PosUncFlag                  respjson.Field
		Ps1astrometryCorrectionFlag respjson.Field
		Ps1ObjInfoFlag              respjson.Field
		Ps1QualityFlag              respjson.Field
		RaUnc                       respjson.Field
		Rmag                        respjson.Field
		RmagOrigin                  respjson.Field
		RmagUnc                     respjson.Field
		Rpmag                       respjson.Field
		RpmagUnc                    respjson.Field
		Ruwe                        respjson.Field
		SdaCatID                    respjson.Field
		Sgmag                       respjson.Field
		SgmagUnc                    respjson.Field
		Shift                       respjson.Field
		ShiftFlag                   respjson.Field
		ShiftFwhm1                  respjson.Field
		ShiftFwhm6                  respjson.Field
		SkyMapperID                 respjson.Field
		TwoMassID                   respjson.Field
		TwoMassPhQualInd            respjson.Field
		TwoMassReadFlag             respjson.Field
		TwoMassXscID                respjson.Field
		TychoDscID                  respjson.Field
		UhsID                       respjson.Field
		UkidssGcsID                 respjson.Field
		UkidssGpsID                 respjson.Field
		UkidssLasID                 respjson.Field
		UpdatedAt                   respjson.Field
		UpdatedBy                   respjson.Field
		VarFlag                     respjson.Field
		Variability                 respjson.Field
		VhsID                       respjson.Field
		Vmag                        respjson.Field
		VmagOrigin                  respjson.Field
		VmagUnc                     respjson.Field
		W1mag                       respjson.Field
		W1magOrigin                 respjson.Field
		W1magUnc                    respjson.Field
		W1sat                       respjson.Field
		W2mag                       respjson.Field
		W2magOrigin                 respjson.Field
		W2magUnc                    respjson.Field
		W2sat                       respjson.Field
		W3mag                       respjson.Field
		W3magOrigin                 respjson.Field
		W3magUnc                    respjson.Field
		W3sat                       respjson.Field
		W4mag                       respjson.Field
		W4magOrigin                 respjson.Field
		W4magUnc                    respjson.Field
		W4sat                       respjson.Field
		WdsCatID                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogListResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type StarCatalogListResponseDataMode string

const (
	StarCatalogListResponseDataModeReal      StarCatalogListResponseDataMode = "REAL"
	StarCatalogListResponseDataModeTest      StarCatalogListResponseDataMode = "TEST"
	StarCatalogListResponseDataModeExercise  StarCatalogListResponseDataMode = "EXERCISE"
	StarCatalogListResponseDataModeSimulated StarCatalogListResponseDataMode = "SIMULATED"
)

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogGetResponse struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode StarCatalogGetResponseDataMode `json:"dataMode" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID int64 `json:"aavsoVsxId"`
	// Optical AB g magnitude.
	Abgmag float64 `json:"abgmag"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin string `json:"abgmagOrigin"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc float64 `json:"abgmagUnc"`
	// Optical AB i magnitude.
	Abimag float64 `json:"abimag"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin string `json:"abimagOrigin"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc float64 `json:"abimagUnc"`
	// Optical AB r magnitude.
	Abrmag float64 `json:"abrmag"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin string `json:"abrmagOrigin"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc float64 `json:"abrmagUnc"`
	// Optical AB y magnitude.
	Abymag float64 `json:"abymag"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin string `json:"abymagOrigin"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc float64 `json:"abymagUnc"`
	// Optical AB z magnitude.
	Abzmag float64 `json:"abzmag"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin string `json:"abzmagOrigin"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc float64 `json:"abzmagUnc"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd string `json:"allWISEccInd"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID string `json:"allWISEId"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd int64 `json:"allWISEnaInd"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd string `json:"allWISEphQualInd"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID string `json:"apassId"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise float64 `json:"astrometricExcessNoise"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig float64 `json:"astrometricExcessNoiseSig"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag float64 `json:"bmag"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin string `json:"bmagOrigin"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc float64 `json:"bmagUnc"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag float64 `json:"bpmag"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID int64 `json:"carrascoCatId"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID string `json:"catWise2020Id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID string `json:"ducatiCatId"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag float64 `json:"gmag"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID int64 `json:"gncCatId"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex int64 `json:"healpixIndex"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID int64 `json:"hipCatId"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag float64 `json:"hmag"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin string `json:"hmagOrigin"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc float64 `json:"hmagUnc"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag float64 `json:"imag"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin string `json:"imagOrigin"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc float64 `json:"imagUnc"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag float64 `json:"jmag"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin string `json:"jmagOrigin"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc float64 `json:"jmagUnc"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag float64 `json:"kmag"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin string `json:"kmagOrigin"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc float64 `json:"kmagUnc"`
	// Morphology indicator.
	MorphologyInd int64 `json:"morphologyInd"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity string `json:"multiplicity"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec float64 `json:"neighborDec"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa float64 `json:"neighborRa"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar string `json:"nonSingleStar"`
	// Number of neighbors.
	NumNeighbors int64 `json:"numNeighbors"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID int64 `json:"panStarrsId"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag int64 `json:"ps1astrometryCorrectionFlag"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag int64 `json:"ps1ObjInfoFlag"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag int64 `json:"ps1QualityFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag float64 `json:"rmag"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin string `json:"rmagOrigin"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc float64 `json:"rmagUnc"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag float64 `json:"rpmag"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc float64 `json:"rpmagUnc"`
	// RUWE in Gaia.
	Ruwe float64 `json:"ruwe"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID int64 `json:"sdaCatId"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag float64 `json:"sgmag"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc float64 `json:"sgmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 float64 `json:"shiftFWHM1"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 float64 `json:"shiftFWHM6"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID int64 `json:"skyMapperId"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID string `json:"twoMASSId"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd string `json:"twoMassPHQualInd"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag string `json:"twoMassReadFlag"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID string `json:"twoMassXscId"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID int64 `json:"tychoDscId"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID int64 `json:"uhsId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID int64 `json:"ukidssGCSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID int64 `json:"ukidssGPSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID int64 `json:"ukidssLASId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability string `json:"variability"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID int64 `json:"vhsId"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag float64 `json:"vmag"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin string `json:"vmagOrigin"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc float64 `json:"vmagUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag float64 `json:"w1mag"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin string `json:"w1magOrigin"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc float64 `json:"w1magUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat float64 `json:"w1sat"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag float64 `json:"w2mag"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin string `json:"w2magOrigin"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc float64 `json:"w2magUnc"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat float64 `json:"w2sat"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag float64 `json:"w3mag"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin string `json:"w3magOrigin"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc float64 `json:"w3magUnc"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat float64 `json:"w3sat"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag float64 `json:"w4mag"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin string `json:"w4magOrigin"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc float64 `json:"w4magUnc"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat float64 `json:"w4sat"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID string `json:"wdsCatId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin            respjson.Field
		ClassificationMarking       respjson.Field
		CsID                        respjson.Field
		DataMode                    respjson.Field
		Dec                         respjson.Field
		Ra                          respjson.Field
		Source                      respjson.Field
		StarEpoch                   respjson.Field
		ID                          respjson.Field
		AavsoVsxID                  respjson.Field
		Abgmag                      respjson.Field
		AbgmagOrigin                respjson.Field
		AbgmagUnc                   respjson.Field
		Abimag                      respjson.Field
		AbimagOrigin                respjson.Field
		AbimagUnc                   respjson.Field
		Abrmag                      respjson.Field
		AbrmagOrigin                respjson.Field
		AbrmagUnc                   respjson.Field
		Abymag                      respjson.Field
		AbymagOrigin                respjson.Field
		AbymagUnc                   respjson.Field
		Abzmag                      respjson.Field
		AbzmagOrigin                respjson.Field
		AbzmagUnc                   respjson.Field
		AllWisEccInd                respjson.Field
		AllWiseID                   respjson.Field
		AllWisEnaInd                respjson.Field
		AllWisEphQualInd            respjson.Field
		ApassID                     respjson.Field
		AstrometricExcessNoise      respjson.Field
		AstrometricExcessNoiseSig   respjson.Field
		Bmag                        respjson.Field
		BmagOrigin                  respjson.Field
		BmagUnc                     respjson.Field
		Bpmag                       respjson.Field
		BpmagUnc                    respjson.Field
		CarrascoCatID               respjson.Field
		CatVersion                  respjson.Field
		CatWise2020ID               respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		DecUnc                      respjson.Field
		DucatiCatID                 respjson.Field
		Gaiadr3CatID                respjson.Field
		Gmag                        respjson.Field
		GmagUnc                     respjson.Field
		GncCatID                    respjson.Field
		HealpixIndex                respjson.Field
		HipCatID                    respjson.Field
		Hmag                        respjson.Field
		HmagOrigin                  respjson.Field
		HmagUnc                     respjson.Field
		Imag                        respjson.Field
		ImagOrigin                  respjson.Field
		ImagUnc                     respjson.Field
		Jmag                        respjson.Field
		JmagOrigin                  respjson.Field
		JmagUnc                     respjson.Field
		Kmag                        respjson.Field
		KmagOrigin                  respjson.Field
		KmagUnc                     respjson.Field
		MorphologyInd               respjson.Field
		MultFlag                    respjson.Field
		Multiplicity                respjson.Field
		NeighborDec                 respjson.Field
		NeighborDistance            respjson.Field
		NeighborFlag                respjson.Field
		NeighborID                  respjson.Field
		NeighborRa                  respjson.Field
		NonSingleStar               respjson.Field
		NumNeighbors                respjson.Field
		Origin                      respjson.Field
		OrigNetwork                 respjson.Field
		PanStarrsID                 respjson.Field
		Parallax                    respjson.Field
		ParallaxUnc                 respjson.Field
		Pmdec                       respjson.Field
		PmdecUnc                    respjson.Field
		Pmra                        respjson.Field
		PmraUnc                     respjson.Field
		PmUncFlag                   respjson.Field
		PosUncFlag                  respjson.Field
		Ps1astrometryCorrectionFlag respjson.Field
		Ps1ObjInfoFlag              respjson.Field
		Ps1QualityFlag              respjson.Field
		RaUnc                       respjson.Field
		Rmag                        respjson.Field
		RmagOrigin                  respjson.Field
		RmagUnc                     respjson.Field
		Rpmag                       respjson.Field
		RpmagUnc                    respjson.Field
		Ruwe                        respjson.Field
		SdaCatID                    respjson.Field
		Sgmag                       respjson.Field
		SgmagUnc                    respjson.Field
		Shift                       respjson.Field
		ShiftFlag                   respjson.Field
		ShiftFwhm1                  respjson.Field
		ShiftFwhm6                  respjson.Field
		SkyMapperID                 respjson.Field
		TwoMassID                   respjson.Field
		TwoMassPhQualInd            respjson.Field
		TwoMassReadFlag             respjson.Field
		TwoMassXscID                respjson.Field
		TychoDscID                  respjson.Field
		UhsID                       respjson.Field
		UkidssGcsID                 respjson.Field
		UkidssGpsID                 respjson.Field
		UkidssLasID                 respjson.Field
		UpdatedAt                   respjson.Field
		UpdatedBy                   respjson.Field
		VarFlag                     respjson.Field
		Variability                 respjson.Field
		VhsID                       respjson.Field
		Vmag                        respjson.Field
		VmagOrigin                  respjson.Field
		VmagUnc                     respjson.Field
		W1mag                       respjson.Field
		W1magOrigin                 respjson.Field
		W1magUnc                    respjson.Field
		W1sat                       respjson.Field
		W2mag                       respjson.Field
		W2magOrigin                 respjson.Field
		W2magUnc                    respjson.Field
		W2sat                       respjson.Field
		W3mag                       respjson.Field
		W3magOrigin                 respjson.Field
		W3magUnc                    respjson.Field
		W3sat                       respjson.Field
		W4mag                       respjson.Field
		W4magOrigin                 respjson.Field
		W4magUnc                    respjson.Field
		W4sat                       respjson.Field
		WdsCatID                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type StarCatalogGetResponseDataMode string

const (
	StarCatalogGetResponseDataModeReal      StarCatalogGetResponseDataMode = "REAL"
	StarCatalogGetResponseDataModeTest      StarCatalogGetResponseDataMode = "TEST"
	StarCatalogGetResponseDataModeExercise  StarCatalogGetResponseDataMode = "EXERCISE"
	StarCatalogGetResponseDataModeSimulated StarCatalogGetResponseDataMode = "SIMULATED"
)

type StarCatalogQueryhelpResponse struct {
	AodrSupported         bool                         `json:"aodrSupported"`
	ClassificationMarking string                       `json:"classificationMarking"`
	Description           string                       `json:"description"`
	HistorySupported      bool                         `json:"historySupported"`
	Name                  string                       `json:"name"`
	Parameters            []shared.ParamDescriptorResp `json:"parameters"`
	RequiredRoles         []string                     `json:"requiredRoles"`
	RestSupported         bool                         `json:"restSupported"`
	SortSupported         bool                         `json:"sortSupported"`
	TypeName              string                       `json:"typeName"`
	Uri                   string                       `json:"uri"`
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
func (r StarCatalogQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
type StarCatalogTupleResponse struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode StarCatalogTupleResponseDataMode `json:"dataMode" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID int64 `json:"aavsoVsxId"`
	// Optical AB g magnitude.
	Abgmag float64 `json:"abgmag"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin string `json:"abgmagOrigin"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc float64 `json:"abgmagUnc"`
	// Optical AB i magnitude.
	Abimag float64 `json:"abimag"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin string `json:"abimagOrigin"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc float64 `json:"abimagUnc"`
	// Optical AB r magnitude.
	Abrmag float64 `json:"abrmag"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin string `json:"abrmagOrigin"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc float64 `json:"abrmagUnc"`
	// Optical AB y magnitude.
	Abymag float64 `json:"abymag"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin string `json:"abymagOrigin"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc float64 `json:"abymagUnc"`
	// Optical AB z magnitude.
	Abzmag float64 `json:"abzmag"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin string `json:"abzmagOrigin"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc float64 `json:"abzmagUnc"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd string `json:"allWISEccInd"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID string `json:"allWISEId"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd int64 `json:"allWISEnaInd"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd string `json:"allWISEphQualInd"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID string `json:"apassId"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise float64 `json:"astrometricExcessNoise"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig float64 `json:"astrometricExcessNoiseSig"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag float64 `json:"bmag"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin string `json:"bmagOrigin"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc float64 `json:"bmagUnc"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag float64 `json:"bpmag"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc float64 `json:"bpmagUnc"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID int64 `json:"carrascoCatId"`
	// The version of the catalog associated with this object.
	CatVersion string `json:"catVersion"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID string `json:"catWise2020Id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc float64 `json:"decUnc"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID string `json:"ducatiCatId"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID int64 `json:"gaiadr3CatId"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag float64 `json:"gmag"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc float64 `json:"gmagUnc"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID int64 `json:"gncCatId"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex int64 `json:"healpixIndex"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID int64 `json:"hipCatId"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag float64 `json:"hmag"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin string `json:"hmagOrigin"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc float64 `json:"hmagUnc"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag float64 `json:"imag"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin string `json:"imagOrigin"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc float64 `json:"imagUnc"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag float64 `json:"jmag"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin string `json:"jmagOrigin"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc float64 `json:"jmagUnc"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag float64 `json:"kmag"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin string `json:"kmagOrigin"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc float64 `json:"kmagUnc"`
	// Morphology indicator.
	MorphologyInd int64 `json:"morphologyInd"`
	// Flag indicating that this is a multiple object source.
	MultFlag bool `json:"multFlag"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity string `json:"multiplicity"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec float64 `json:"neighborDec"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance float64 `json:"neighborDistance"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag bool `json:"neighborFlag"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID int64 `json:"neighborId"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa float64 `json:"neighborRa"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar string `json:"nonSingleStar"`
	// Number of neighbors.
	NumNeighbors int64 `json:"numNeighbors"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID int64 `json:"panStarrsId"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax float64 `json:"parallax"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc float64 `json:"parallaxUnc"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec float64 `json:"pmdec"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc float64 `json:"pmdecUnc"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra float64 `json:"pmra"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc float64 `json:"pmraUnc"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag bool `json:"pmUncFlag"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag bool `json:"posUncFlag"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag int64 `json:"ps1astrometryCorrectionFlag"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag int64 `json:"ps1ObjInfoFlag"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag int64 `json:"ps1QualityFlag"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc float64 `json:"raUnc"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag float64 `json:"rmag"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin string `json:"rmagOrigin"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc float64 `json:"rmagUnc"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag float64 `json:"rpmag"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc float64 `json:"rpmagUnc"`
	// RUWE in Gaia.
	Ruwe float64 `json:"ruwe"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID int64 `json:"sdaCatId"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag float64 `json:"sgmag"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc float64 `json:"sgmagUnc"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift float64 `json:"shift"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag bool `json:"shiftFlag"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 float64 `json:"shiftFWHM1"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 float64 `json:"shiftFWHM6"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID int64 `json:"skyMapperId"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID string `json:"twoMASSId"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd string `json:"twoMassPHQualInd"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag string `json:"twoMassReadFlag"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID string `json:"twoMassXscId"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID int64 `json:"tychoDscId"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID int64 `json:"uhsId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID int64 `json:"ukidssGCSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID int64 `json:"ukidssGPSId"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID int64 `json:"ukidssLASId"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag bool `json:"varFlag"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability string `json:"variability"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID int64 `json:"vhsId"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag float64 `json:"vmag"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin string `json:"vmagOrigin"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc float64 `json:"vmagUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag float64 `json:"w1mag"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin string `json:"w1magOrigin"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc float64 `json:"w1magUnc"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat float64 `json:"w1sat"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag float64 `json:"w2mag"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin string `json:"w2magOrigin"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc float64 `json:"w2magUnc"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat float64 `json:"w2sat"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag float64 `json:"w3mag"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin string `json:"w3magOrigin"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc float64 `json:"w3magUnc"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat float64 `json:"w3sat"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag float64 `json:"w4mag"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin string `json:"w4magOrigin"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc float64 `json:"w4magUnc"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat float64 `json:"w4sat"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID string `json:"wdsCatId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AstrometryOrigin            respjson.Field
		ClassificationMarking       respjson.Field
		CsID                        respjson.Field
		DataMode                    respjson.Field
		Dec                         respjson.Field
		Ra                          respjson.Field
		Source                      respjson.Field
		StarEpoch                   respjson.Field
		ID                          respjson.Field
		AavsoVsxID                  respjson.Field
		Abgmag                      respjson.Field
		AbgmagOrigin                respjson.Field
		AbgmagUnc                   respjson.Field
		Abimag                      respjson.Field
		AbimagOrigin                respjson.Field
		AbimagUnc                   respjson.Field
		Abrmag                      respjson.Field
		AbrmagOrigin                respjson.Field
		AbrmagUnc                   respjson.Field
		Abymag                      respjson.Field
		AbymagOrigin                respjson.Field
		AbymagUnc                   respjson.Field
		Abzmag                      respjson.Field
		AbzmagOrigin                respjson.Field
		AbzmagUnc                   respjson.Field
		AllWisEccInd                respjson.Field
		AllWiseID                   respjson.Field
		AllWisEnaInd                respjson.Field
		AllWisEphQualInd            respjson.Field
		ApassID                     respjson.Field
		AstrometricExcessNoise      respjson.Field
		AstrometricExcessNoiseSig   respjson.Field
		Bmag                        respjson.Field
		BmagOrigin                  respjson.Field
		BmagUnc                     respjson.Field
		Bpmag                       respjson.Field
		BpmagUnc                    respjson.Field
		CarrascoCatID               respjson.Field
		CatVersion                  respjson.Field
		CatWise2020ID               respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		DecUnc                      respjson.Field
		DucatiCatID                 respjson.Field
		Gaiadr3CatID                respjson.Field
		Gmag                        respjson.Field
		GmagUnc                     respjson.Field
		GncCatID                    respjson.Field
		HealpixIndex                respjson.Field
		HipCatID                    respjson.Field
		Hmag                        respjson.Field
		HmagOrigin                  respjson.Field
		HmagUnc                     respjson.Field
		Imag                        respjson.Field
		ImagOrigin                  respjson.Field
		ImagUnc                     respjson.Field
		Jmag                        respjson.Field
		JmagOrigin                  respjson.Field
		JmagUnc                     respjson.Field
		Kmag                        respjson.Field
		KmagOrigin                  respjson.Field
		KmagUnc                     respjson.Field
		MorphologyInd               respjson.Field
		MultFlag                    respjson.Field
		Multiplicity                respjson.Field
		NeighborDec                 respjson.Field
		NeighborDistance            respjson.Field
		NeighborFlag                respjson.Field
		NeighborID                  respjson.Field
		NeighborRa                  respjson.Field
		NonSingleStar               respjson.Field
		NumNeighbors                respjson.Field
		Origin                      respjson.Field
		OrigNetwork                 respjson.Field
		PanStarrsID                 respjson.Field
		Parallax                    respjson.Field
		ParallaxUnc                 respjson.Field
		Pmdec                       respjson.Field
		PmdecUnc                    respjson.Field
		Pmra                        respjson.Field
		PmraUnc                     respjson.Field
		PmUncFlag                   respjson.Field
		PosUncFlag                  respjson.Field
		Ps1astrometryCorrectionFlag respjson.Field
		Ps1ObjInfoFlag              respjson.Field
		Ps1QualityFlag              respjson.Field
		RaUnc                       respjson.Field
		Rmag                        respjson.Field
		RmagOrigin                  respjson.Field
		RmagUnc                     respjson.Field
		Rpmag                       respjson.Field
		RpmagUnc                    respjson.Field
		Ruwe                        respjson.Field
		SdaCatID                    respjson.Field
		Sgmag                       respjson.Field
		SgmagUnc                    respjson.Field
		Shift                       respjson.Field
		ShiftFlag                   respjson.Field
		ShiftFwhm1                  respjson.Field
		ShiftFwhm6                  respjson.Field
		SkyMapperID                 respjson.Field
		TwoMassID                   respjson.Field
		TwoMassPhQualInd            respjson.Field
		TwoMassReadFlag             respjson.Field
		TwoMassXscID                respjson.Field
		TychoDscID                  respjson.Field
		UhsID                       respjson.Field
		UkidssGcsID                 respjson.Field
		UkidssGpsID                 respjson.Field
		UkidssLasID                 respjson.Field
		UpdatedAt                   respjson.Field
		UpdatedBy                   respjson.Field
		VarFlag                     respjson.Field
		Variability                 respjson.Field
		VhsID                       respjson.Field
		Vmag                        respjson.Field
		VmagOrigin                  respjson.Field
		VmagUnc                     respjson.Field
		W1mag                       respjson.Field
		W1magOrigin                 respjson.Field
		W1magUnc                    respjson.Field
		W1sat                       respjson.Field
		W2mag                       respjson.Field
		W2magOrigin                 respjson.Field
		W2magUnc                    respjson.Field
		W2sat                       respjson.Field
		W3mag                       respjson.Field
		W3magOrigin                 respjson.Field
		W3magUnc                    respjson.Field
		W3sat                       respjson.Field
		W4mag                       respjson.Field
		W4magOrigin                 respjson.Field
		W4magUnc                    respjson.Field
		W4sat                       respjson.Field
		WdsCatID                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StarCatalogTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *StarCatalogTupleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type StarCatalogTupleResponseDataMode string

const (
	StarCatalogTupleResponseDataModeReal      StarCatalogTupleResponseDataMode = "REAL"
	StarCatalogTupleResponseDataModeTest      StarCatalogTupleResponseDataMode = "TEST"
	StarCatalogTupleResponseDataModeExercise  StarCatalogTupleResponseDataMode = "EXERCISE"
	StarCatalogTupleResponseDataModeSimulated StarCatalogTupleResponseDataMode = "SIMULATED"
)

type StarCatalogNewParams struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode StarCatalogNewParamsDataMode `json:"dataMode,omitzero" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID param.Opt[int64] `json:"aavsoVsxId,omitzero"`
	// Optical AB g magnitude.
	Abgmag param.Opt[float64] `json:"abgmag,omitzero"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin param.Opt[string] `json:"abgmagOrigin,omitzero"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc param.Opt[float64] `json:"abgmagUnc,omitzero"`
	// Optical AB i magnitude.
	Abimag param.Opt[float64] `json:"abimag,omitzero"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin param.Opt[string] `json:"abimagOrigin,omitzero"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc param.Opt[float64] `json:"abimagUnc,omitzero"`
	// Optical AB r magnitude.
	Abrmag param.Opt[float64] `json:"abrmag,omitzero"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin param.Opt[string] `json:"abrmagOrigin,omitzero"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc param.Opt[float64] `json:"abrmagUnc,omitzero"`
	// Optical AB y magnitude.
	Abymag param.Opt[float64] `json:"abymag,omitzero"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin param.Opt[string] `json:"abymagOrigin,omitzero"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc param.Opt[float64] `json:"abymagUnc,omitzero"`
	// Optical AB z magnitude.
	Abzmag param.Opt[float64] `json:"abzmag,omitzero"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin param.Opt[string] `json:"abzmagOrigin,omitzero"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc param.Opt[float64] `json:"abzmagUnc,omitzero"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd param.Opt[string] `json:"allWISEccInd,omitzero"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID param.Opt[string] `json:"allWISEId,omitzero"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd param.Opt[int64] `json:"allWISEnaInd,omitzero"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd param.Opt[string] `json:"allWISEphQualInd,omitzero"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID param.Opt[string] `json:"apassId,omitzero"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise param.Opt[float64] `json:"astrometricExcessNoise,omitzero"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig param.Opt[float64] `json:"astrometricExcessNoiseSig,omitzero"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag param.Opt[float64] `json:"bmag,omitzero"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin param.Opt[string] `json:"bmagOrigin,omitzero"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc param.Opt[float64] `json:"bmagUnc,omitzero"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID param.Opt[int64] `json:"carrascoCatId,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID param.Opt[string] `json:"catWise2020Id,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID param.Opt[string] `json:"ducatiCatId,omitzero"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex param.Opt[int64] `json:"healpixIndex,omitzero"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin param.Opt[string] `json:"hmagOrigin,omitzero"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag param.Opt[float64] `json:"imag,omitzero"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin param.Opt[string] `json:"imagOrigin,omitzero"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc param.Opt[float64] `json:"imagUnc,omitzero"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin param.Opt[string] `json:"jmagOrigin,omitzero"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin param.Opt[string] `json:"kmagOrigin,omitzero"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Morphology indicator.
	MorphologyInd param.Opt[int64] `json:"morphologyInd,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity param.Opt[string] `json:"multiplicity,omitzero"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec param.Opt[float64] `json:"neighborDec,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa param.Opt[float64] `json:"neighborRa,omitzero"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar param.Opt[string] `json:"nonSingleStar,omitzero"`
	// Number of neighbors.
	NumNeighbors param.Opt[int64] `json:"numNeighbors,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID param.Opt[int64] `json:"panStarrsId,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag param.Opt[int64] `json:"ps1astrometryCorrectionFlag,omitzero"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag param.Opt[int64] `json:"ps1ObjInfoFlag,omitzero"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag param.Opt[int64] `json:"ps1QualityFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag param.Opt[float64] `json:"rmag,omitzero"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin param.Opt[string] `json:"rmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc param.Opt[float64] `json:"rmagUnc,omitzero"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// RUWE in Gaia.
	Ruwe param.Opt[float64] `json:"ruwe,omitzero"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID param.Opt[int64] `json:"sdaCatId,omitzero"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag param.Opt[float64] `json:"sgmag,omitzero"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc param.Opt[float64] `json:"sgmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 param.Opt[float64] `json:"shiftFWHM1,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 param.Opt[float64] `json:"shiftFWHM6,omitzero"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID param.Opt[int64] `json:"skyMapperId,omitzero"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID param.Opt[string] `json:"twoMASSId,omitzero"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd param.Opt[string] `json:"twoMassPHQualInd,omitzero"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag param.Opt[string] `json:"twoMassReadFlag,omitzero"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID param.Opt[string] `json:"twoMassXscId,omitzero"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID param.Opt[int64] `json:"tychoDscId,omitzero"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID param.Opt[int64] `json:"uhsId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID param.Opt[int64] `json:"ukidssGCSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID param.Opt[int64] `json:"ukidssGPSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID param.Opt[int64] `json:"ukidssLASId,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability param.Opt[string] `json:"variability,omitzero"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID param.Opt[int64] `json:"vhsId,omitzero"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag param.Opt[float64] `json:"vmag,omitzero"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin param.Opt[string] `json:"vmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc param.Opt[float64] `json:"vmagUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag param.Opt[float64] `json:"w1mag,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin param.Opt[string] `json:"w1magOrigin,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc param.Opt[float64] `json:"w1magUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat param.Opt[float64] `json:"w1sat,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag param.Opt[float64] `json:"w2mag,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin param.Opt[string] `json:"w2magOrigin,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc param.Opt[float64] `json:"w2magUnc,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat param.Opt[float64] `json:"w2sat,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag param.Opt[float64] `json:"w3mag,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin param.Opt[string] `json:"w3magOrigin,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc param.Opt[float64] `json:"w3magUnc,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat param.Opt[float64] `json:"w3sat,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag param.Opt[float64] `json:"w4mag,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin param.Opt[string] `json:"w4magOrigin,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc param.Opt[float64] `json:"w4magUnc,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat param.Opt[float64] `json:"w4sat,omitzero"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID param.Opt[string] `json:"wdsCatId,omitzero"`
	paramObj
}

func (r StarCatalogNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type StarCatalogNewParamsDataMode string

const (
	StarCatalogNewParamsDataModeReal      StarCatalogNewParamsDataMode = "REAL"
	StarCatalogNewParamsDataModeTest      StarCatalogNewParamsDataMode = "TEST"
	StarCatalogNewParamsDataModeExercise  StarCatalogNewParamsDataMode = "EXERCISE"
	StarCatalogNewParamsDataModeSimulated StarCatalogNewParamsDataMode = "SIMULATED"
)

type StarCatalogUpdateParams struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode StarCatalogUpdateParamsDataMode `json:"dataMode,omitzero" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID param.Opt[int64] `json:"aavsoVsxId,omitzero"`
	// Optical AB g magnitude.
	Abgmag param.Opt[float64] `json:"abgmag,omitzero"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin param.Opt[string] `json:"abgmagOrigin,omitzero"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc param.Opt[float64] `json:"abgmagUnc,omitzero"`
	// Optical AB i magnitude.
	Abimag param.Opt[float64] `json:"abimag,omitzero"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin param.Opt[string] `json:"abimagOrigin,omitzero"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc param.Opt[float64] `json:"abimagUnc,omitzero"`
	// Optical AB r magnitude.
	Abrmag param.Opt[float64] `json:"abrmag,omitzero"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin param.Opt[string] `json:"abrmagOrigin,omitzero"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc param.Opt[float64] `json:"abrmagUnc,omitzero"`
	// Optical AB y magnitude.
	Abymag param.Opt[float64] `json:"abymag,omitzero"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin param.Opt[string] `json:"abymagOrigin,omitzero"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc param.Opt[float64] `json:"abymagUnc,omitzero"`
	// Optical AB z magnitude.
	Abzmag param.Opt[float64] `json:"abzmag,omitzero"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin param.Opt[string] `json:"abzmagOrigin,omitzero"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc param.Opt[float64] `json:"abzmagUnc,omitzero"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd param.Opt[string] `json:"allWISEccInd,omitzero"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID param.Opt[string] `json:"allWISEId,omitzero"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd param.Opt[int64] `json:"allWISEnaInd,omitzero"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd param.Opt[string] `json:"allWISEphQualInd,omitzero"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID param.Opt[string] `json:"apassId,omitzero"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise param.Opt[float64] `json:"astrometricExcessNoise,omitzero"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig param.Opt[float64] `json:"astrometricExcessNoiseSig,omitzero"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag param.Opt[float64] `json:"bmag,omitzero"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin param.Opt[string] `json:"bmagOrigin,omitzero"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc param.Opt[float64] `json:"bmagUnc,omitzero"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID param.Opt[int64] `json:"carrascoCatId,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID param.Opt[string] `json:"catWise2020Id,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID param.Opt[string] `json:"ducatiCatId,omitzero"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex param.Opt[int64] `json:"healpixIndex,omitzero"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin param.Opt[string] `json:"hmagOrigin,omitzero"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag param.Opt[float64] `json:"imag,omitzero"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin param.Opt[string] `json:"imagOrigin,omitzero"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc param.Opt[float64] `json:"imagUnc,omitzero"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin param.Opt[string] `json:"jmagOrigin,omitzero"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin param.Opt[string] `json:"kmagOrigin,omitzero"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Morphology indicator.
	MorphologyInd param.Opt[int64] `json:"morphologyInd,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity param.Opt[string] `json:"multiplicity,omitzero"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec param.Opt[float64] `json:"neighborDec,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa param.Opt[float64] `json:"neighborRa,omitzero"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar param.Opt[string] `json:"nonSingleStar,omitzero"`
	// Number of neighbors.
	NumNeighbors param.Opt[int64] `json:"numNeighbors,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID param.Opt[int64] `json:"panStarrsId,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag param.Opt[int64] `json:"ps1astrometryCorrectionFlag,omitzero"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag param.Opt[int64] `json:"ps1ObjInfoFlag,omitzero"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag param.Opt[int64] `json:"ps1QualityFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag param.Opt[float64] `json:"rmag,omitzero"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin param.Opt[string] `json:"rmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc param.Opt[float64] `json:"rmagUnc,omitzero"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// RUWE in Gaia.
	Ruwe param.Opt[float64] `json:"ruwe,omitzero"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID param.Opt[int64] `json:"sdaCatId,omitzero"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag param.Opt[float64] `json:"sgmag,omitzero"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc param.Opt[float64] `json:"sgmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 param.Opt[float64] `json:"shiftFWHM1,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 param.Opt[float64] `json:"shiftFWHM6,omitzero"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID param.Opt[int64] `json:"skyMapperId,omitzero"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID param.Opt[string] `json:"twoMASSId,omitzero"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd param.Opt[string] `json:"twoMassPHQualInd,omitzero"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag param.Opt[string] `json:"twoMassReadFlag,omitzero"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID param.Opt[string] `json:"twoMassXscId,omitzero"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID param.Opt[int64] `json:"tychoDscId,omitzero"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID param.Opt[int64] `json:"uhsId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID param.Opt[int64] `json:"ukidssGCSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID param.Opt[int64] `json:"ukidssGPSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID param.Opt[int64] `json:"ukidssLASId,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability param.Opt[string] `json:"variability,omitzero"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID param.Opt[int64] `json:"vhsId,omitzero"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag param.Opt[float64] `json:"vmag,omitzero"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin param.Opt[string] `json:"vmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc param.Opt[float64] `json:"vmagUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag param.Opt[float64] `json:"w1mag,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin param.Opt[string] `json:"w1magOrigin,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc param.Opt[float64] `json:"w1magUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat param.Opt[float64] `json:"w1sat,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag param.Opt[float64] `json:"w2mag,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin param.Opt[string] `json:"w2magOrigin,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc param.Opt[float64] `json:"w2magUnc,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat param.Opt[float64] `json:"w2sat,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag param.Opt[float64] `json:"w3mag,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin param.Opt[string] `json:"w3magOrigin,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc param.Opt[float64] `json:"w3magUnc,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat param.Opt[float64] `json:"w3sat,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag param.Opt[float64] `json:"w4mag,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin param.Opt[string] `json:"w4magOrigin,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc param.Opt[float64] `json:"w4magUnc,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat param.Opt[float64] `json:"w4sat,omitzero"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID param.Opt[string] `json:"wdsCatId,omitzero"`
	paramObj
}

func (r StarCatalogUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StarCatalogUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StarCatalogUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
type StarCatalogUpdateParamsDataMode string

const (
	StarCatalogUpdateParamsDataModeReal      StarCatalogUpdateParamsDataMode = "REAL"
	StarCatalogUpdateParamsDataModeTest      StarCatalogUpdateParamsDataMode = "TEST"
	StarCatalogUpdateParamsDataModeExercise  StarCatalogUpdateParamsDataMode = "EXERCISE"
	StarCatalogUpdateParamsDataModeSimulated StarCatalogUpdateParamsDataMode = "SIMULATED"
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
	return shimjson.Marshal(r.Body)
}
func (r *StarCatalogNewBulkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarCatalogNewBulkParamsBody struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode string `json:"dataMode,omitzero" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID param.Opt[int64] `json:"aavsoVsxId,omitzero"`
	// Optical AB g magnitude.
	Abgmag param.Opt[float64] `json:"abgmag,omitzero"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin param.Opt[string] `json:"abgmagOrigin,omitzero"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc param.Opt[float64] `json:"abgmagUnc,omitzero"`
	// Optical AB i magnitude.
	Abimag param.Opt[float64] `json:"abimag,omitzero"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin param.Opt[string] `json:"abimagOrigin,omitzero"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc param.Opt[float64] `json:"abimagUnc,omitzero"`
	// Optical AB r magnitude.
	Abrmag param.Opt[float64] `json:"abrmag,omitzero"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin param.Opt[string] `json:"abrmagOrigin,omitzero"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc param.Opt[float64] `json:"abrmagUnc,omitzero"`
	// Optical AB y magnitude.
	Abymag param.Opt[float64] `json:"abymag,omitzero"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin param.Opt[string] `json:"abymagOrigin,omitzero"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc param.Opt[float64] `json:"abymagUnc,omitzero"`
	// Optical AB z magnitude.
	Abzmag param.Opt[float64] `json:"abzmag,omitzero"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin param.Opt[string] `json:"abzmagOrigin,omitzero"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc param.Opt[float64] `json:"abzmagUnc,omitzero"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd param.Opt[string] `json:"allWISEccInd,omitzero"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID param.Opt[string] `json:"allWISEId,omitzero"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd param.Opt[int64] `json:"allWISEnaInd,omitzero"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd param.Opt[string] `json:"allWISEphQualInd,omitzero"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID param.Opt[string] `json:"apassId,omitzero"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise param.Opt[float64] `json:"astrometricExcessNoise,omitzero"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig param.Opt[float64] `json:"astrometricExcessNoiseSig,omitzero"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag param.Opt[float64] `json:"bmag,omitzero"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin param.Opt[string] `json:"bmagOrigin,omitzero"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc param.Opt[float64] `json:"bmagUnc,omitzero"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID param.Opt[int64] `json:"carrascoCatId,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID param.Opt[string] `json:"catWise2020Id,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID param.Opt[string] `json:"ducatiCatId,omitzero"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex param.Opt[int64] `json:"healpixIndex,omitzero"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin param.Opt[string] `json:"hmagOrigin,omitzero"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag param.Opt[float64] `json:"imag,omitzero"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin param.Opt[string] `json:"imagOrigin,omitzero"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc param.Opt[float64] `json:"imagUnc,omitzero"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin param.Opt[string] `json:"jmagOrigin,omitzero"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin param.Opt[string] `json:"kmagOrigin,omitzero"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Morphology indicator.
	MorphologyInd param.Opt[int64] `json:"morphologyInd,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity param.Opt[string] `json:"multiplicity,omitzero"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec param.Opt[float64] `json:"neighborDec,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa param.Opt[float64] `json:"neighborRa,omitzero"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar param.Opt[string] `json:"nonSingleStar,omitzero"`
	// Number of neighbors.
	NumNeighbors param.Opt[int64] `json:"numNeighbors,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID param.Opt[int64] `json:"panStarrsId,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag param.Opt[int64] `json:"ps1astrometryCorrectionFlag,omitzero"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag param.Opt[int64] `json:"ps1ObjInfoFlag,omitzero"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag param.Opt[int64] `json:"ps1QualityFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag param.Opt[float64] `json:"rmag,omitzero"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin param.Opt[string] `json:"rmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc param.Opt[float64] `json:"rmagUnc,omitzero"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// RUWE in Gaia.
	Ruwe param.Opt[float64] `json:"ruwe,omitzero"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID param.Opt[int64] `json:"sdaCatId,omitzero"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag param.Opt[float64] `json:"sgmag,omitzero"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc param.Opt[float64] `json:"sgmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 param.Opt[float64] `json:"shiftFWHM1,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 param.Opt[float64] `json:"shiftFWHM6,omitzero"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID param.Opt[int64] `json:"skyMapperId,omitzero"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID param.Opt[string] `json:"twoMASSId,omitzero"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd param.Opt[string] `json:"twoMassPHQualInd,omitzero"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag param.Opt[string] `json:"twoMassReadFlag,omitzero"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID param.Opt[string] `json:"twoMassXscId,omitzero"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID param.Opt[int64] `json:"tychoDscId,omitzero"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID param.Opt[int64] `json:"uhsId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID param.Opt[int64] `json:"ukidssGCSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID param.Opt[int64] `json:"ukidssGPSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID param.Opt[int64] `json:"ukidssLASId,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability param.Opt[string] `json:"variability,omitzero"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID param.Opt[int64] `json:"vhsId,omitzero"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag param.Opt[float64] `json:"vmag,omitzero"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin param.Opt[string] `json:"vmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc param.Opt[float64] `json:"vmagUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag param.Opt[float64] `json:"w1mag,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin param.Opt[string] `json:"w1magOrigin,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc param.Opt[float64] `json:"w1magUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat param.Opt[float64] `json:"w1sat,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag param.Opt[float64] `json:"w2mag,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin param.Opt[string] `json:"w2magOrigin,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc param.Opt[float64] `json:"w2magUnc,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat param.Opt[float64] `json:"w2sat,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag param.Opt[float64] `json:"w3mag,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin param.Opt[string] `json:"w3magOrigin,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc param.Opt[float64] `json:"w3magUnc,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat param.Opt[float64] `json:"w3sat,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag param.Opt[float64] `json:"w4mag,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin param.Opt[string] `json:"w4magOrigin,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc param.Opt[float64] `json:"w4magUnc,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat param.Opt[float64] `json:"w4sat,omitzero"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID param.Opt[string] `json:"wdsCatId,omitzero"`
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
		"dataMode", "REAL", "TEST", "EXERCISE", "SIMULATED",
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
	Columns string `query:"columns" api:"required" json:"-"`
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
	return shimjson.Marshal(r.Body)
}
func (r *StarCatalogUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The star catalog provides the position, proper motion, parallax, and photometric
// magnitudes at various bandpasses of a star.
//
// The properties AstrometryOrigin, ClassificationMarking, CsID, DataMode, Dec, Ra,
// Source, StarEpoch are required.
type StarCatalogUnvalidatedPublishParamsBody struct {
	// Originating astrometric catalog for this object (GA (GAIA), HI (HIPPARCOS), UB
	// (USNOBSC), AL, AP, CA, CR, DU, FK6_I, FK6_III, PS, SK, TD, TP, TX, UC, UL, UH,
	// UP, VH, VS, WD).
	AstrometryOrigin string `json:"astrometryOrigin" api:"required"`
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking" api:"required"`
	// The ID of this object in the specific catalog associated with this record. This
	// field will either contain the value in the gncCatId or sdaCatId field.
	CsID int64 `json:"csId" api:"required"`
	// Indicator of whether the data is REAL, TEST, EXERCISE, or SIMULATED data:
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// Any of "REAL", "TEST", "EXERCISE", "SIMULATED".
	DataMode string `json:"dataMode,omitzero" api:"required"`
	// Barycentric declination of the source in International Celestial Reference
	// System (ICRS) at the reference epoch, in degrees.
	Dec float64 `json:"dec" api:"required"`
	// Barycentric right ascension of the source in the International Celestial
	// Reference System (ICRS) frame at the reference epoch, in degrees.
	Ra float64 `json:"ra" api:"required"`
	// Source of the data.
	Source string `json:"source" api:"required"`
	// Reference epoch to which the astrometric source parameters are referred,
	// expressed as Julian Year in Barycentric Coordinate Time (TCB).
	StarEpoch float64 `json:"starEpoch" api:"required"`
	// The American Association of Variable Star Observers (AAVSO) Variable Star Index
	// (VSX) (VX) object ID of this object.
	AavsoVsxID param.Opt[int64] `json:"aavsoVsxId,omitzero"`
	// Optical AB g magnitude.
	Abgmag param.Opt[float64] `json:"abgmag,omitzero"`
	// Catalog of origin of optical AB g magnitude.
	AbgmagOrigin param.Opt[string] `json:"abgmagOrigin,omitzero"`
	// Uncertainty of optical AB g magnitude.
	AbgmagUnc param.Opt[float64] `json:"abgmagUnc,omitzero"`
	// Optical AB i magnitude.
	Abimag param.Opt[float64] `json:"abimag,omitzero"`
	// Catalog of origin of optical AB i magnitude.
	AbimagOrigin param.Opt[string] `json:"abimagOrigin,omitzero"`
	// Uncertainty of optical AB i magnitude.
	AbimagUnc param.Opt[float64] `json:"abimagUnc,omitzero"`
	// Optical AB r magnitude.
	Abrmag param.Opt[float64] `json:"abrmag,omitzero"`
	// Catalog of origin of optical AB r magnitude.
	AbrmagOrigin param.Opt[string] `json:"abrmagOrigin,omitzero"`
	// Uncertainty of optical AB r magnitude.
	AbrmagUnc param.Opt[float64] `json:"abrmagUnc,omitzero"`
	// Optical AB y magnitude.
	Abymag param.Opt[float64] `json:"abymag,omitzero"`
	// Catalog of origin of optical AB y magnitude.
	AbymagOrigin param.Opt[string] `json:"abymagOrigin,omitzero"`
	// Uncertainty of optical AB y magnitude.
	AbymagUnc param.Opt[float64] `json:"abymagUnc,omitzero"`
	// Optical AB z magnitude.
	Abzmag param.Opt[float64] `json:"abzmag,omitzero"`
	// Catalog of origin of optical AB z magnitude.
	AbzmagOrigin param.Opt[string] `json:"abzmagOrigin,omitzero"`
	// Uncertainty of optical AB z magnitude.
	AbzmagUnc param.Opt[float64] `json:"abzmagUnc,omitzero"`
	// Contamination and confusion indicator in AllWISE.
	AllWisEccInd param.Opt[string] `json:"allWISEccInd,omitzero"`
	// The designation of this object in the All Wide-field Infrared Survey Explorer
	// (AllWISE) catalog (AL).
	AllWiseID param.Opt[string] `json:"allWISEId,omitzero"`
	// Active deblending indicator in AllWISE.
	AllWisEnaInd param.Opt[int64] `json:"allWISEnaInd,omitzero"`
	// Photometric quality indicator in AllWISE.
	AllWisEphQualInd param.Opt[string] `json:"allWISEphQualInd,omitzero"`
	// The American Association of Variable Star Observers (AAVSO) Photometric All-Sky
	// Survey (APASS) (AP) name of this object.
	ApassID param.Opt[string] `json:"apassId,omitzero"`
	// Astrometric excess noise in the Gaia catalog measured in milliarcseconds.
	AstrometricExcessNoise param.Opt[float64] `json:"astrometricExcessNoise,omitzero"`
	// Astrometric excess noise sigma in Gaia.
	AstrometricExcessNoiseSig param.Opt[float64] `json:"astrometricExcessNoiseSig,omitzero"`
	// Optical Johnson B magnitude measured in magnitudes.
	Bmag param.Opt[float64] `json:"bmag,omitzero"`
	// Catalog of origin of optical Johnson B magnitude (AP, CR, HI).
	BmagOrigin param.Opt[string] `json:"bmagOrigin,omitzero"`
	// Uncertainty of optical Johnson B magnitude measured in magnitudes.
	BmagUnc param.Opt[float64] `json:"bmagUnc,omitzero"`
	// Gaia optical photometric Bp-band in the Vega scale measured in magnitudes.
	Bpmag param.Opt[float64] `json:"bpmag,omitzero"`
	// Gaia optical Bp-band uncertainty in the Vega scale measured in magnitudes.
	BpmagUnc param.Opt[float64] `json:"bpmagUnc,omitzero"`
	// The Carrasco catalog (CR) identifier of this object.
	CarrascoCatID param.Opt[int64] `json:"carrascoCatId,omitzero"`
	// The version of the catalog associated with this object.
	CatVersion param.Opt[string] `json:"catVersion,omitzero"`
	// The CatWISE2020 (CA) catalog source ID of this object.
	CatWise2020ID param.Opt[string] `json:"catWise2020Id,omitzero"`
	// Uncertainty of the declination of the source, in milliarcseconds, at the
	// reference epoch.
	DecUnc param.Opt[float64] `json:"decUnc,omitzero"`
	// The Ducati catalog (DU) name of this object.
	DucatiCatID param.Opt[string] `json:"ducatiCatId,omitzero"`
	// The source ID of this object in the Gaia DR3 Catalog (GA).
	Gaiadr3CatID param.Opt[int64] `json:"gaiadr3CatId,omitzero"`
	// Gaia optical photometric G-band in the Vega scale measured in magnitudes.
	Gmag param.Opt[float64] `json:"gmag,omitzero"`
	// Gaia optical photometric G-band uncertainty in the Vega scale measured in
	// magnitudes.
	GmagUnc param.Opt[float64] `json:"gmagUnc,omitzero"`
	// The ID of this object in the Guidance and Navigation Control (GNC) Catalog. If
	// this field is populated it shall match the csId field.
	GncCatID param.Opt[int64] `json:"gncCatId,omitzero"`
	// The Healpix index. Consumers should contact the provider for details on the
	// indexing scheme.
	HealpixIndex param.Opt[int64] `json:"healpixIndex,omitzero"`
	// The HIP ID of this object in the Hipparcos Catalog (HI).
	HipCatID param.Opt[int64] `json:"hipCatId,omitzero"`
	// Near-infrared photometric H-band magnitude in the Vega scale measured in
	// magnitudes.
	Hmag param.Opt[float64] `json:"hmag,omitzero"`
	// Near-infrared photometric H-band catalog of origin in the Vega scale (TP, UC,
	// UL, UP, VH).
	HmagOrigin param.Opt[string] `json:"hmagOrigin,omitzero"`
	// Near-infrared photometric H-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	HmagUnc param.Opt[float64] `json:"hmagUnc,omitzero"`
	// Optical Johnson I magnitude measured in magnitudes.
	Imag param.Opt[float64] `json:"imag,omitzero"`
	// Catalog of origin of optical Johnson I magnitude (CR, GA, HI).
	ImagOrigin param.Opt[string] `json:"imagOrigin,omitzero"`
	// Uncertainty of optical Johnson I magnitude measured in magnitudes.
	ImagUnc param.Opt[float64] `json:"imagUnc,omitzero"`
	// Near-infrared photometric J-band magnitude in the Vega scale measured in
	// magnitudes.
	Jmag param.Opt[float64] `json:"jmag,omitzero"`
	// Near-infrared photometric J-band catalog of origin in the Vega scale (TP, UH,
	// UL, UP, VH).
	JmagOrigin param.Opt[string] `json:"jmagOrigin,omitzero"`
	// Near-infrared photometric J-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	JmagUnc param.Opt[float64] `json:"jmagUnc,omitzero"`
	// Near-infrared photometric K-band magnitude in the Vega scale measured in
	// magnitudes.
	Kmag param.Opt[float64] `json:"kmag,omitzero"`
	// Near-infrared photometric K-band catalog of origin in the Vega scale (TP, UC,
	// UH, UL, UP, VH).
	KmagOrigin param.Opt[string] `json:"kmagOrigin,omitzero"`
	// Near-infrared photometric K-band magnitude uncertainty in the Vega scale
	// measured in magnitudes.
	KmagUnc param.Opt[float64] `json:"kmagUnc,omitzero"`
	// Morphology indicator.
	MorphologyInd param.Opt[int64] `json:"morphologyInd,omitzero"`
	// Flag indicating that this is a multiple object source.
	MultFlag param.Opt[bool] `json:"multFlag,omitzero"`
	// Identifier indicating multiplicity is detected. Consumers should contact the
	// provider for details on the specifications.
	Multiplicity param.Opt[string] `json:"multiplicity,omitzero"`
	// Dec of nearest neighbor measured in degrees.
	NeighborDec param.Opt[float64] `json:"neighborDec,omitzero"`
	// Distance between source and nearest neighbor, in arcseconds.
	NeighborDistance param.Opt[float64] `json:"neighborDistance,omitzero"`
	// Flag indicating that the nearest catalog neighbor is closer than 4.6 arcseconds.
	NeighborFlag param.Opt[bool] `json:"neighborFlag,omitzero"`
	// The catalog ID of the nearest neighbor to this source.
	NeighborID param.Opt[int64] `json:"neighborId,omitzero"`
	// RA of nearest neighbor measured in degrees.
	NeighborRa param.Opt[float64] `json:"neighborRa,omitzero"`
	// Identifier indicating the source is a non-single star and additional information
	// is available in non-single star tables. Consumers should contact the provider
	// for details on the specifications.
	NonSingleStar param.Opt[string] `json:"nonSingleStar,omitzero"`
	// Number of neighbors.
	NumNeighbors param.Opt[int64] `json:"numNeighbors,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The Panoramic Survey Telescope and Rapid Response System (Pan-STARRS) (PS)
	// object ID.
	PanStarrsID param.Opt[int64] `json:"panStarrsId,omitzero"`
	// Absolute stellar parallax of the source, in milliarcseconds.
	Parallax param.Opt[float64] `json:"parallax,omitzero"`
	// Uncertainty of the stellar parallax, in milliarcseconds.
	ParallaxUnc param.Opt[float64] `json:"parallaxUnc,omitzero"`
	// Proper motion in declination of the source, in milliarcseconds per year, at the
	// reference epoch.
	Pmdec param.Opt[float64] `json:"pmdec,omitzero"`
	// Uncertainty of proper motion in declination, in milliarcseconds per year.
	PmdecUnc param.Opt[float64] `json:"pmdecUnc,omitzero"`
	// Proper motion in right ascension of the source, in milliarcseconds per year, at
	// the reference epoch.
	Pmra param.Opt[float64] `json:"pmra,omitzero"`
	// Uncertainty of proper motion in right ascension, in milliarcseconds per year.
	PmraUnc param.Opt[float64] `json:"pmraUnc,omitzero"`
	// Flag indicating that the proper motion uncertainty in either ra or dec is
	// greater than 10 milliarcseconds per year.
	PmUncFlag param.Opt[bool] `json:"pmUncFlag,omitzero"`
	// Flag indicating that the position uncertainty in either ra or dec is greater
	// than 100 milliarcseconds.
	PosUncFlag param.Opt[bool] `json:"posUncFlag,omitzero"`
	// Astrometry correction flag in Pan-STARRS.
	Ps1astrometryCorrectionFlag param.Opt[int64] `json:"ps1astrometryCorrectionFlag,omitzero"`
	// Object information flag in Pan-STARRS.
	Ps1ObjInfoFlag param.Opt[int64] `json:"ps1ObjInfoFlag,omitzero"`
	// Quality flag in Pan-STARRS.
	Ps1QualityFlag param.Opt[int64] `json:"ps1QualityFlag,omitzero"`
	// Uncertainty of the right ascension of the source, in milliarcseconds, at the
	// reference epoch.
	RaUnc param.Opt[float64] `json:"raUnc,omitzero"`
	// Optical Johnson R magnitude measured in magnitudes.
	Rmag param.Opt[float64] `json:"rmag,omitzero"`
	// Catalog of origin of the Optical Johnson R magnitude (CR, GA).
	RmagOrigin param.Opt[string] `json:"rmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson R magnitude measured in magnitudes.
	RmagUnc param.Opt[float64] `json:"rmagUnc,omitzero"`
	// Gaia optical Rp-band in the Vega scale measured in magnitudes.
	Rpmag param.Opt[float64] `json:"rpmag,omitzero"`
	// Gaia optical photometric Rp-band uncertainty in the Vega scale measured in
	// magnitudes.
	RpmagUnc param.Opt[float64] `json:"rpmagUnc,omitzero"`
	// RUWE in Gaia.
	Ruwe param.Opt[float64] `json:"ruwe,omitzero"`
	// The ID of this object in the Space Domain Awareness (SDA) Catalog. If this field
	// is populated it shall match the csId field.
	SdaCatID param.Opt[int64] `json:"sdaCatId,omitzero"`
	// Original G magnitude if the source is in Gaia, otherwise the magnitude is
	// converted from other photometric passbands, when possible, measured in
	// magnitudes.
	Sgmag param.Opt[float64] `json:"sgmag,omitzero"`
	// Uncertainty of sgmag measured in magnitudes.
	SgmagUnc param.Opt[float64] `json:"sgmagUnc,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds.
	Shift param.Opt[float64] `json:"shift,omitzero"`
	// Flag indicating that the photocentric shift is greater than 50 milliarcseconds.
	ShiftFlag param.Opt[bool] `json:"shiftFlag,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of one
	// arcsecond.
	ShiftFwhm1 param.Opt[float64] `json:"shiftFWHM1,omitzero"`
	// Photocentric shift caused by neighbors, in arcseconds. This value is constrained
	// to a Point Spread Function (PSF) with Full Width at Half Maximum (FWHM) of six
	// arcseconds.
	ShiftFwhm6 param.Opt[float64] `json:"shiftFWHM6,omitzero"`
	// The SkyMapper (SK) catalog object ID.
	SkyMapperID param.Opt[int64] `json:"skyMapperId,omitzero"`
	// The designation of this object in the Two Micron All Sky Survey (2MASS) Point
	// Source Catalog (TP).
	TwoMassID param.Opt[string] `json:"twoMASSId,omitzero"`
	// Photometric (PH) quality indicator in 2MASS PSC.
	TwoMassPhQualInd param.Opt[string] `json:"twoMassPHQualInd,omitzero"`
	// Read flag in 2MASS PSC.
	TwoMassReadFlag param.Opt[string] `json:"twoMassReadFlag,omitzero"`
	// The Two Micron All Sky Survey (2MASS) Extended Source Catalog (XSC) (TX)
	// designation of this object.
	TwoMassXscID param.Opt[string] `json:"twoMassXscId,omitzero"`
	// The Tycho Double Star Catalog (TD) identifier (specified as Tycho-2 ID) of this
	// object.
	TychoDscID param.Opt[int64] `json:"tychoDscId,omitzero"`
	// The United Kingdom Infrared Telescope (UKIRT) Hemispheric Survey (UHS) (UH)
	// source ID of this object.
	UhsID param.Opt[int64] `json:"uhsId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Clusters Survey
	// (GCS) (UC) source ID of this object.
	UkidssGcsID param.Opt[int64] `json:"ukidssGCSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Galactic Plane Survey (GPS)
	// (UP) source ID of this object.
	UkidssGpsID param.Opt[int64] `json:"ukidssGPSId,omitzero"`
	// The United Kingdom Infrared Deep Sky Survey (UKIDSS) Large Area Survey (LAS)
	// (UL) source ID of this object.
	UkidssLasID param.Opt[int64] `json:"ukidssLASId,omitzero"`
	// Flag indicating that the source exhibits variable magnitude.
	VarFlag param.Opt[bool] `json:"varFlag,omitzero"`
	// Identifier indicating variability is present in the photometric data. Consumers
	// should contact the provider for details on the specifications.
	Variability param.Opt[string] `json:"variability,omitzero"`
	// The Visible and Infrared Survey Telescope for Astronomy (VISTA) Hemisphere
	// Survey (VHS) (VS) source ID of this object.
	VhsID param.Opt[int64] `json:"vhsId,omitzero"`
	// Optical Johnson V magnitude measured in magnitudes.
	Vmag param.Opt[float64] `json:"vmag,omitzero"`
	// Catalog of origin of Optical Johnson V magnitude (AP, CR, DU, GA, HI).
	VmagOrigin param.Opt[string] `json:"vmagOrigin,omitzero"`
	// Uncertainty of the Optical Johnson V magnitude measured in magnitudes.
	VmagUnc param.Opt[float64] `json:"vmagUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude in the Vega system
	// measured in magnitudes.
	W1mag param.Opt[float64] `json:"w1mag,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) catalog of origin in the Vega
	// system (AL, CA).
	W1magOrigin param.Opt[string] `json:"w1magOrigin,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W1magUnc param.Opt[float64] `json:"w1magUnc,omitzero"`
	// Mid-infrared photometric W1-band (3.4 microns) saturated pixel fraction in the
	// Vega system measured in magnitudes.
	W1sat param.Opt[float64] `json:"w1sat,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude in the Vega system
	// measured in magnitudes.
	W2mag param.Opt[float64] `json:"w2mag,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) catalog of origin in the Vega
	// system (AL, CA).
	W2magOrigin param.Opt[string] `json:"w2magOrigin,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W2magUnc param.Opt[float64] `json:"w2magUnc,omitzero"`
	// Mid-infrared photometric W2-band (4.6 microns) saturated pixel fraction in the
	// Vega system.
	W2sat param.Opt[float64] `json:"w2sat,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude in the Vega system
	// measured in magnitudes.
	W3mag param.Opt[float64] `json:"w3mag,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) catalog of origin in the Vega
	// system (AL).
	W3magOrigin param.Opt[string] `json:"w3magOrigin,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W3magUnc param.Opt[float64] `json:"w3magUnc,omitzero"`
	// Mid-infrared photometric W3-band (12 microns) saturated pixel fraction in the
	// Vega system.
	W3sat param.Opt[float64] `json:"w3sat,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude in the Vega system
	// measured in magnitudes.
	W4mag param.Opt[float64] `json:"w4mag,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) catalog of origin in the Vega
	// system (AL).
	W4magOrigin param.Opt[string] `json:"w4magOrigin,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) magnitude uncertainty in the Vega
	// system measured in magnitudes.
	W4magUnc param.Opt[float64] `json:"w4magUnc,omitzero"`
	// Mid-infrared photometric W4-band (22 microns) saturated pixel fraction in the
	// Vega system.
	W4sat param.Opt[float64] `json:"w4sat,omitzero"`
	// The Washington Double Star Catalog (WD) identifier of this object.
	WdsCatID param.Opt[string] `json:"wdsCatId,omitzero"`
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
		"dataMode", "REAL", "TEST", "EXERCISE", "SIMULATED",
	)
}
