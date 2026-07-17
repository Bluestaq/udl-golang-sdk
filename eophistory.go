// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/v2/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/v2/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
	"github.com/Bluestaq/udl-golang-sdk/v2/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/v2/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/v2/shared"
)

// This service provides operations for manipulation and querying of earth
// orientation parameter (EOP) data. Earth Orientation Parameters (EOP) are
// produced by the IERS (International Earth Rotation and Reference Systems
// Service). Earth Orientation Parameters describe the irregularities of the
// earth's rotation. Technically, they are the parameters which provide the
// rotation of the ITRS (International Terrestrial Reference System) to the ICRS
// (International Celestial Reference System) as a function of time. Universal time
// -- Universal time (UT1) is the time of the earth clock, which performs one
// revolution in about 24h. It is practically proportional to the sidereal time.
// The excess revolution time is called length of day (LOD). Coordinates of the
// pole -- x and y are the coordinates of the Celestial Ephemeris Pole (CEP)
// relative to the IRP, the IERS Reference Pole. The CEP differs from the
// instantaneous rotation axis by quasi-diurnal terms with amplitudes under 0.01"
// (see Seidelmann, 1982). The x-axis is in the direction of the ITRF
// zero-meridian; the y-axis is in the direction 90 degrees West longitude.
// Celestial pole offsets -- Celestial pole offsets are described in the IAU
// Precession and Nutation models. The observed differences with respect to the
// conventional celestial pole position defined by the models are monitored and
// reported by the IERS. IERS Bulletins A and B provide current information on the
// Earth's orientation in the IERS Reference System. This includes Universal Time,
// coordinates of the terrestrial pole, and celestial pole offsets. Bulletin A
// gives an advanced solution updated weekly; the standard solution is given
// monthly in Bulletin B. Fields suffixed with ”B” are Bulletin B values. All
// solutions are continuous within their respective uncertainties. Bulletin A is
// issued by the IERS Rapid Service/Prediction Centre at the U.S. Naval
// Observatory, Washington, DC and Bulletin B is issued by the IERS Earth
// Orientation Centre at the Paris Observatory. IERS Bulletin A reports the latest
// determinations for polar motion, UT1-UTC, and nutation offsets at daily
// intervals based on a combination of contributed analysis results using data from
// Very Long Baseline Interferometry (VLBI), Satellite Laser Ranging (SLR), Global
// Positioning System (GPS) satellites, and Lunar Laser Ranging (LLR). Predictions
// for variations a year into the future are also provided. Meteorological
// predictions of variations in Atmospheric Angular Momentum (AAM) are used to aid
// in the prediction of near-term UT1-UTC changes. This publication is prepared by
// the IERS Rapid Service/Prediction Center.
//
// EopHistoryService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEopHistoryService] method instead.
type EopHistoryService struct {
	Options []option.RequestOption
}

// NewEopHistoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEopHistoryService(opts ...option.RequestOption) (r EopHistoryService) {
	r = EopHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (`/udl/<datatype>/queryhelp`) for more details on valid/required query parameter
// information.
func (r *EopHistoryService) List(ctx context.Context, query EopHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[shared.EopFull], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/eop/history"
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

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (`/udl/<datatype>/queryhelp`) for more details on valid/required query parameter
// information.
func (r *EopHistoryService) ListAutoPaging(ctx context.Context, query EopHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[shared.EopFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation (`/udl/<datatype>/queryhelp`)
// for more details on valid/required query parameter information.
func (r *EopHistoryService) Aodr(ctx context.Context, query EopHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/eop/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (`/udl/<datatype>/queryhelp`) for more details on
// valid/required query parameter information.
func (r *EopHistoryService) Count(ctx context.Context, query EopHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/eop/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EopHistoryListParams struct {
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate time.Time `query:"eopDate" api:"required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (`/udl/<datatype>/queryhelp`) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopHistoryListParams]'s query parameters as `url.Values`.
func (r EopHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EopHistoryAodrParams struct {
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate time.Time `query:"eopDate" api:"required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (`/udl/<datatype>/queryhelp`) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopHistoryAodrParams]'s query parameters as `url.Values`.
func (r EopHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EopHistoryCountParams struct {
	// Effective date/time for the EOP values in ISO8601 UTC format. The values could
	// be current or predicted. (YYYY-MM-DDTHH:MM:SS.sssZ)
	EopDate     time.Time        `query:"eopDate" api:"required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EopHistoryCountParams]'s query parameters as `url.Values`.
func (r EopHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
