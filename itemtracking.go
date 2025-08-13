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

// ItemTrackingService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewItemTrackingService] method instead.
type ItemTrackingService struct {
	Options []option.RequestOption
	History ItemTrackingHistoryService
}

// NewItemTrackingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewItemTrackingService(opts ...option.RequestOption) (r ItemTrackingService) {
	r = ItemTrackingService{}
	r.Options = opts
	r.History = NewItemTrackingHistoryService(opts...)
	return
}

// Service operation to take a single itemtracking record as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *ItemTrackingService) New(ctx context.Context, body ItemTrackingNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/itemtracking"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ItemTrackingService) List(ctx context.Context, query ItemTrackingListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ItemTrackingListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/itemtracking"
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
func (r *ItemTrackingService) ListAutoPaging(ctx context.Context, query ItemTrackingListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ItemTrackingListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an item tracking object specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ItemTrackingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/itemtracking/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ItemTrackingService) Count(ctx context.Context, query ItemTrackingCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/itemtracking/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single item tracking record by its unique ID passed
// as a path parameter.
func (r *ItemTrackingService) Get(ctx context.Context, id string, query ItemTrackingGetParams, opts ...option.RequestOption) (res *ItemTrackingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/itemtracking/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ItemTrackingService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *ItemTrackingQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/itemtracking/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *ItemTrackingService) Tuple(ctx context.Context, query ItemTrackingTupleParams, opts ...option.RequestOption) (res *[]ItemTrackingTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/itemtracking/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple itemtracking records as a POST body and
// ingest into the database. This operation is intended to be used for automated
// feeds into UDL. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ItemTrackingService) UnvalidatedPublish(ctx context.Context, body ItemTrackingUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-itemtracking"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ItemTrackingListResponse struct {
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
	DataMode ItemTrackingListResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC.
	ScanCode string `json:"scanCode,required"`
	// The ID of the scanner or input device.
	ScannerID string `json:"scannerId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The United States distinguished visitor code of the person scanned, only
	// applicable to people.
	DvCode string `json:"dvCode"`
	// The UDL ID of the item this record is associated with.
	IDItem string `json:"idItem"`
	// Array of keys that may be associated with this tracked item.
	Keys []string `json:"keys"`
	// WGS84 latitude where the item was scanned, in degrees.
	Lat float64 `json:"lat"`
	// WGS84 longitude where the item was scanned, in degrees.
	Lon float64 `json:"lon"`
	// Optional notes or comments about the tracking data.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of scan for tracking the item within it's journey (e.g. IN, OUT,
	// RECEIVED, DELIVERED, TRANSIT, ABANDONDED, REFUSED, UNABLE, RETURNED, HELD,
	// OTHER). For example, received and delivered are for when an item is received
	// from or delivered to the end customer. In and out are for stops in between such
	// as being loaded on an airplane or received at a warehouse.
	ScanType string `json:"scanType"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The type of item that is being scanned (e.g. CARGO, PERSON, MAIL, MICAP, OTHER).
	Type string `json:"type"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		ScannerID             respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DvCode                respjson.Field
		IDItem                respjson.Field
		Keys                  respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ScanType              respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Type                  respjson.Field
		Values                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemTrackingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemTrackingListResponse) UnmarshalJSON(data []byte) error {
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
type ItemTrackingListResponseDataMode string

const (
	ItemTrackingListResponseDataModeReal      ItemTrackingListResponseDataMode = "REAL"
	ItemTrackingListResponseDataModeTest      ItemTrackingListResponseDataMode = "TEST"
	ItemTrackingListResponseDataModeSimulated ItemTrackingListResponseDataMode = "SIMULATED"
	ItemTrackingListResponseDataModeExercise  ItemTrackingListResponseDataMode = "EXERCISE"
)

type ItemTrackingGetResponse struct {
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
	DataMode ItemTrackingGetResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC.
	ScanCode string `json:"scanCode,required"`
	// The ID of the scanner or input device.
	ScannerID string `json:"scannerId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The United States distinguished visitor code of the person scanned, only
	// applicable to people.
	DvCode string `json:"dvCode"`
	// The UDL ID of the item this record is associated with.
	IDItem string `json:"idItem"`
	// Array of keys that may be associated with this tracked item.
	Keys []string `json:"keys"`
	// WGS84 latitude where the item was scanned, in degrees.
	Lat float64 `json:"lat"`
	// WGS84 longitude where the item was scanned, in degrees.
	Lon float64 `json:"lon"`
	// Optional notes or comments about the tracking data.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of scan for tracking the item within it's journey (e.g. IN, OUT,
	// RECEIVED, DELIVERED, TRANSIT, ABANDONDED, REFUSED, UNABLE, RETURNED, HELD,
	// OTHER). For example, received and delivered are for when an item is received
	// from or delivered to the end customer. In and out are for stops in between such
	// as being loaded on an airplane or received at a warehouse.
	ScanType string `json:"scanType"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The type of item that is being scanned (e.g. CARGO, PERSON, MAIL, MICAP, OTHER).
	Type string `json:"type"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		ScannerID             respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DvCode                respjson.Field
		IDItem                respjson.Field
		Keys                  respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ScanType              respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Type                  respjson.Field
		Values                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemTrackingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemTrackingGetResponse) UnmarshalJSON(data []byte) error {
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
type ItemTrackingGetResponseDataMode string

const (
	ItemTrackingGetResponseDataModeReal      ItemTrackingGetResponseDataMode = "REAL"
	ItemTrackingGetResponseDataModeTest      ItemTrackingGetResponseDataMode = "TEST"
	ItemTrackingGetResponseDataModeSimulated ItemTrackingGetResponseDataMode = "SIMULATED"
	ItemTrackingGetResponseDataModeExercise  ItemTrackingGetResponseDataMode = "EXERCISE"
)

type ItemTrackingQueryhelpResponse struct {
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
func (r ItemTrackingQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemTrackingQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemTrackingTupleResponse struct {
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
	DataMode ItemTrackingTupleResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC.
	ScanCode string `json:"scanCode,required"`
	// The ID of the scanner or input device.
	ScannerID string `json:"scannerId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The United States distinguished visitor code of the person scanned, only
	// applicable to people.
	DvCode string `json:"dvCode"`
	// The UDL ID of the item this record is associated with.
	IDItem string `json:"idItem"`
	// Array of keys that may be associated with this tracked item.
	Keys []string `json:"keys"`
	// WGS84 latitude where the item was scanned, in degrees.
	Lat float64 `json:"lat"`
	// WGS84 longitude where the item was scanned, in degrees.
	Lon float64 `json:"lon"`
	// Optional notes or comments about the tracking data.
	Notes string `json:"notes"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of scan for tracking the item within it's journey (e.g. IN, OUT,
	// RECEIVED, DELIVERED, TRANSIT, ABANDONDED, REFUSED, UNABLE, RETURNED, HELD,
	// OTHER). For example, received and delivered are for when an item is received
	// from or delivered to the end customer. In and out are for stops in between such
	// as being loaded on an airplane or received at a warehouse.
	ScanType string `json:"scanType"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The type of item that is being scanned (e.g. CARGO, PERSON, MAIL, MICAP, OTHER).
	Type string `json:"type"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		ScannerID             respjson.Field
		Source                respjson.Field
		Ts                    respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DvCode                respjson.Field
		IDItem                respjson.Field
		Keys                  respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Notes                 respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ScanType              respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Type                  respjson.Field
		Values                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemTrackingTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemTrackingTupleResponse) UnmarshalJSON(data []byte) error {
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
type ItemTrackingTupleResponseDataMode string

const (
	ItemTrackingTupleResponseDataModeReal      ItemTrackingTupleResponseDataMode = "REAL"
	ItemTrackingTupleResponseDataModeTest      ItemTrackingTupleResponseDataMode = "TEST"
	ItemTrackingTupleResponseDataModeSimulated ItemTrackingTupleResponseDataMode = "SIMULATED"
	ItemTrackingTupleResponseDataModeExercise  ItemTrackingTupleResponseDataMode = "EXERCISE"
)

type ItemTrackingNewParams struct {
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
	DataMode ItemTrackingNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC.
	ScanCode string `json:"scanCode,required"`
	// The ID of the scanner or input device.
	ScannerID string `json:"scannerId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The United States distinguished visitor code of the person scanned, only
	// applicable to people.
	DvCode param.Opt[string] `json:"dvCode,omitzero"`
	// The UDL ID of the item this record is associated with.
	IDItem param.Opt[string] `json:"idItem,omitzero"`
	// WGS84 latitude where the item was scanned, in degrees.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude where the item was scanned, in degrees.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Optional notes or comments about the tracking data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The type of scan for tracking the item within it's journey (e.g. IN, OUT,
	// RECEIVED, DELIVERED, TRANSIT, ABANDONDED, REFUSED, UNABLE, RETURNED, HELD,
	// OTHER). For example, received and delivered are for when an item is received
	// from or delivered to the end customer. In and out are for stops in between such
	// as being loaded on an airplane or received at a warehouse.
	ScanType param.Opt[string] `json:"scanType,omitzero"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool param.Opt[string] `json:"scGenTool,omitzero"`
	// The type of item that is being scanned (e.g. CARGO, PERSON, MAIL, MICAP, OTHER).
	Type param.Opt[string] `json:"type,omitzero"`
	// Array of keys that may be associated with this tracked item.
	Keys []string `json:"keys,omitzero"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ItemTrackingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemTrackingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemTrackingNewParams) UnmarshalJSON(data []byte) error {
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
type ItemTrackingNewParamsDataMode string

const (
	ItemTrackingNewParamsDataModeReal      ItemTrackingNewParamsDataMode = "REAL"
	ItemTrackingNewParamsDataModeTest      ItemTrackingNewParamsDataMode = "TEST"
	ItemTrackingNewParamsDataModeSimulated ItemTrackingNewParamsDataMode = "SIMULATED"
	ItemTrackingNewParamsDataModeExercise  ItemTrackingNewParamsDataMode = "EXERCISE"
)

type ItemTrackingListParams struct {
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemTrackingListParams]'s query parameters as `url.Values`.
func (r ItemTrackingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemTrackingCountParams struct {
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemTrackingCountParams]'s query parameters as
// `url.Values`.
func (r ItemTrackingCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemTrackingGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemTrackingGetParams]'s query parameters as `url.Values`.
func (r ItemTrackingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemTrackingTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemTrackingTupleParams]'s query parameters as
// `url.Values`.
func (r ItemTrackingTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemTrackingUnvalidatedPublishParams struct {
	Body []ItemTrackingUnvalidatedPublishParamsBody
	paramObj
}

func (r ItemTrackingUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ItemTrackingUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties ClassificationMarking, DataMode, ScanCode, ScannerID, Source, Ts
// are required.
type ItemTrackingUnvalidatedPublishParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC.
	ScanCode string `json:"scanCode,required"`
	// The ID of the scanner or input device.
	ScannerID string `json:"scannerId,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the scan, in ISO 8601 UTC format with millisecond precision.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The United States distinguished visitor code of the person scanned, only
	// applicable to people.
	DvCode param.Opt[string] `json:"dvCode,omitzero"`
	// The UDL ID of the item this record is associated with.
	IDItem param.Opt[string] `json:"idItem,omitzero"`
	// WGS84 latitude where the item was scanned, in degrees.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude where the item was scanned, in degrees.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Optional notes or comments about the tracking data.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The type of scan for tracking the item within it's journey (e.g. IN, OUT,
	// RECEIVED, DELIVERED, TRANSIT, ABANDONDED, REFUSED, UNABLE, RETURNED, HELD,
	// OTHER). For example, received and delivered are for when an item is received
	// from or delivered to the end customer. In and out are for stops in between such
	// as being loaded on an airplane or received at a warehouse.
	ScanType param.Opt[string] `json:"scanType,omitzero"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool param.Opt[string] `json:"scGenTool,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The type of item that is being scanned (e.g. CARGO, PERSON, MAIL, MICAP, OTHER).
	Type param.Opt[string] `json:"type,omitzero"`
	// Array of keys that may be associated with this tracked item.
	Keys []string `json:"keys,omitzero"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ItemTrackingUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ItemTrackingUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemTrackingUnvalidatedPublishParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ItemTrackingUnvalidatedPublishParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
