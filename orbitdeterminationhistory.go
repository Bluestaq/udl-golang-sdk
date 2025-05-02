// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// OrbitdeterminationHistoryService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrbitdeterminationHistoryService] method instead.
type OrbitdeterminationHistoryService struct {
	Options []option.RequestOption
}

// NewOrbitdeterminationHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrbitdeterminationHistoryService(opts ...option.RequestOption) (r OrbitdeterminationHistoryService) {
	r = OrbitdeterminationHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbitdeterminationHistoryService) List(ctx context.Context, query OrbitdeterminationHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[OrbitdeterminationFull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/orbitdetermination/history"
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
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbitdeterminationHistoryService) ListAutoPaging(ctx context.Context, query OrbitdeterminationHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[OrbitdeterminationFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *OrbitdeterminationHistoryService) Aodr(ctx context.Context, query OrbitdeterminationHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/orbitdetermination/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *OrbitdeterminationHistoryService) Count(ctx context.Context, query OrbitdeterminationHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/orbitdetermination/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrbitdeterminationHistoryListParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnOrbit, startTime' are required.) Unique identifier
	// of the target satellite on-orbit object. This ID can be used to obtain
	// additional information on an OnOrbit object using the 'get by ID' operation
	// (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit = 25544 would
	// be queried as /udl/onorbit/25544.
	IDOnOrbit  param.Opt[string] `query:"idOnOrbit,omitzero" json:"-"`
	MaxResults param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnOrbit, startTime' are required.) Start time for OD
	// solution in ISO 8601 UTC datetime format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OrbitdeterminationHistoryListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [OrbitdeterminationHistoryListParams]'s query parameters as
// `url.Values`.
func (r OrbitdeterminationHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbitdeterminationHistoryAodrParams struct {
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnOrbit, startTime' are required.) Unique identifier
	// of the target satellite on-orbit object. This ID can be used to obtain
	// additional information on an OnOrbit object using the 'get by ID' operation
	// (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit = 25544 would
	// be queried as /udl/onorbit/25544.
	IDOnOrbit  param.Opt[string] `query:"idOnOrbit,omitzero" json:"-"`
	MaxResults param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
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
	// (One or more of fields 'idOnOrbit, startTime' are required.) Start time for OD
	// solution in ISO 8601 UTC datetime format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OrbitdeterminationHistoryAodrParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [OrbitdeterminationHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r OrbitdeterminationHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrbitdeterminationHistoryCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	// (One or more of fields 'idOnOrbit, startTime' are required.) Unique identifier
	// of the target satellite on-orbit object. This ID can be used to obtain
	// additional information on an OnOrbit object using the 'get by ID' operation
	// (e.g. /udl/onorbit/{id}). For example, the OnOrbit with idOnOrbit = 25544 would
	// be queried as /udl/onorbit/25544.
	IDOnOrbit  param.Opt[string] `query:"idOnOrbit,omitzero" json:"-"`
	MaxResults param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// (One or more of fields 'idOnOrbit, startTime' are required.) Start time for OD
	// solution in ISO 8601 UTC datetime format, with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OrbitdeterminationHistoryCountParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [OrbitdeterminationHistoryCountParams]'s query parameters as
// `url.Values`.
func (r OrbitdeterminationHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
