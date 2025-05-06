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

// ItemService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewItemService] method instead.
type ItemService struct {
	Options []option.RequestOption
}

// NewItemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewItemService(opts ...option.RequestOption) (r ItemService) {
	r = ItemService{}
	r.Options = opts
	return
}

// Service operation to take a single item record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *ItemService) New(ctx context.Context, body ItemNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/item"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Item. An Item can be cargo, equipment, or a
// passenger. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ItemService) Update(ctx context.Context, id string, body ItemUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/item/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *ItemService) List(ctx context.Context, query ItemListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/item"
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
func (r *ItemService) ListAutoPaging(ctx context.Context, query ItemListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ItemListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a item record specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ItemService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/item/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *ItemService) Count(ctx context.Context, query ItemCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/item/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single item record by its unique ID passed as a path
// parameter.
func (r *ItemService) Get(ctx context.Context, id string, query ItemGetParams, opts ...option.RequestOption) (res *ItemGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/item/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *ItemService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/item/queryhelp"
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
func (r *ItemService) Tuple(ctx context.Context, query ItemTupleParams, opts ...option.RequestOption) (res *[]ItemTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/item/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple item records as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *ItemService) UnvalidatedPublish(ctx context.Context, body ItemUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-item"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ItemListResponse struct {
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
	DataMode ItemListResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes string `json:"accSysNotes"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem string `json:"accSystem"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop bool `json:"airdrop"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat string `json:"altDataFormat"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType string `json:"cargoType"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset float64 `json:"centerlineOffset"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg float64 `json:"cg"`
	// The classification code of the commodity or group of commodities.
	CommodityCode string `json:"commodityCode"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys string `json:"commoditySys"`
	// Flag indicating this item acts as a container and contains additional items.
	Container bool `json:"container"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The departure code or location where this item has left or is leaving.
	Departure string `json:"departure"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination string `json:"destination"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode string `json:"dvCode"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs float64 `json:"fs"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height float64 `json:"height"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan string `json:"idAirLoadPlan"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate time.Time `json:"lastArrDate" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length float64 `json:"length"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment float64 `json:"moment"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name string `json:"name"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// Optional notes or comments about this item.
	Notes string `json:"notes"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos int64 `json:"numPalletPos"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The code denoting the type of material item.
	ProductCode string `json:"productCode"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys string `json:"productSys"`
	// The military branch receiving this item.
	ReceivingBranch string `json:"receivingBranch"`
	// The name of the unit receiving this item.
	ReceivingUnit string `json:"receivingUnit"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn string `json:"tcn"`
	// The unit line number of this item.
	Uln string `json:"uln"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume float64 `json:"volume"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight float64 `json:"weight"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs time.Time `json:"weightTS" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AccSysKeys            respjson.Field
		AccSysNotes           respjson.Field
		AccSystem             respjson.Field
		AccSysValues          respjson.Field
		Airdrop               respjson.Field
		AltDataFormat         respjson.Field
		CargoType             respjson.Field
		CenterlineOffset      respjson.Field
		Cg                    respjson.Field
		CommodityCode         respjson.Field
		CommoditySys          respjson.Field
		Container             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Departure             respjson.Field
		Destination           respjson.Field
		DvCode                respjson.Field
		Fs                    respjson.Field
		HazCodes              respjson.Field
		Height                respjson.Field
		IDAirLoadPlan         respjson.Field
		ItemContains          respjson.Field
		Keys                  respjson.Field
		LastArrDate           respjson.Field
		Length                respjson.Field
		Moment                respjson.Field
		Name                  respjson.Field
		NetExpWt              respjson.Field
		Notes                 respjson.Field
		NumPalletPos          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProductCode           respjson.Field
		ProductSys            respjson.Field
		ReceivingBranch       respjson.Field
		ReceivingUnit         respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Tcn                   respjson.Field
		Uln                   respjson.Field
		Values                respjson.Field
		Volume                respjson.Field
		Weight                respjson.Field
		WeightTs              respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemListResponse) UnmarshalJSON(data []byte) error {
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
type ItemListResponseDataMode string

const (
	ItemListResponseDataModeReal      ItemListResponseDataMode = "REAL"
	ItemListResponseDataModeTest      ItemListResponseDataMode = "TEST"
	ItemListResponseDataModeSimulated ItemListResponseDataMode = "SIMULATED"
	ItemListResponseDataModeExercise  ItemListResponseDataMode = "EXERCISE"
)

type ItemGetResponse struct {
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
	DataMode ItemGetResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes string `json:"accSysNotes"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem string `json:"accSystem"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop bool `json:"airdrop"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat string `json:"altDataFormat"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType string `json:"cargoType"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset float64 `json:"centerlineOffset"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg float64 `json:"cg"`
	// The classification code of the commodity or group of commodities.
	CommodityCode string `json:"commodityCode"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys string `json:"commoditySys"`
	// Flag indicating this item acts as a container and contains additional items.
	Container bool `json:"container"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The departure code or location where this item has left or is leaving.
	Departure string `json:"departure"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination string `json:"destination"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode string `json:"dvCode"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs float64 `json:"fs"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height float64 `json:"height"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan string `json:"idAirLoadPlan"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate time.Time `json:"lastArrDate" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length float64 `json:"length"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment float64 `json:"moment"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name string `json:"name"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// Optional notes or comments about this item.
	Notes string `json:"notes"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos int64 `json:"numPalletPos"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The code denoting the type of material item.
	ProductCode string `json:"productCode"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys string `json:"productSys"`
	// The military branch receiving this item.
	ReceivingBranch string `json:"receivingBranch"`
	// The name of the unit receiving this item.
	ReceivingUnit string `json:"receivingUnit"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn string `json:"tcn"`
	// The unit line number of this item.
	Uln string `json:"uln"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume float64 `json:"volume"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight float64 `json:"weight"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs time.Time `json:"weightTS" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AccSysKeys            respjson.Field
		AccSysNotes           respjson.Field
		AccSystem             respjson.Field
		AccSysValues          respjson.Field
		Airdrop               respjson.Field
		AltDataFormat         respjson.Field
		CargoType             respjson.Field
		CenterlineOffset      respjson.Field
		Cg                    respjson.Field
		CommodityCode         respjson.Field
		CommoditySys          respjson.Field
		Container             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Departure             respjson.Field
		Destination           respjson.Field
		DvCode                respjson.Field
		Fs                    respjson.Field
		HazCodes              respjson.Field
		Height                respjson.Field
		IDAirLoadPlan         respjson.Field
		ItemContains          respjson.Field
		Keys                  respjson.Field
		LastArrDate           respjson.Field
		Length                respjson.Field
		Moment                respjson.Field
		Name                  respjson.Field
		NetExpWt              respjson.Field
		Notes                 respjson.Field
		NumPalletPos          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProductCode           respjson.Field
		ProductSys            respjson.Field
		ReceivingBranch       respjson.Field
		ReceivingUnit         respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Tcn                   respjson.Field
		Uln                   respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Values                respjson.Field
		Volume                respjson.Field
		Weight                respjson.Field
		WeightTs              respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemGetResponse) UnmarshalJSON(data []byte) error {
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
type ItemGetResponseDataMode string

const (
	ItemGetResponseDataModeReal      ItemGetResponseDataMode = "REAL"
	ItemGetResponseDataModeTest      ItemGetResponseDataMode = "TEST"
	ItemGetResponseDataModeSimulated ItemGetResponseDataMode = "SIMULATED"
	ItemGetResponseDataModeExercise  ItemGetResponseDataMode = "EXERCISE"
)

type ItemTupleResponse struct {
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
	DataMode ItemTupleResponseDataMode `json:"dataMode,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID string `json:"id"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes string `json:"accSysNotes"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem string `json:"accSystem"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop bool `json:"airdrop"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat string `json:"altDataFormat"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType string `json:"cargoType"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset float64 `json:"centerlineOffset"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg float64 `json:"cg"`
	// The classification code of the commodity or group of commodities.
	CommodityCode string `json:"commodityCode"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys string `json:"commoditySys"`
	// Flag indicating this item acts as a container and contains additional items.
	Container bool `json:"container"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The departure code or location where this item has left or is leaving.
	Departure string `json:"departure"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination string `json:"destination"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode string `json:"dvCode"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs float64 `json:"fs"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height float64 `json:"height"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan string `json:"idAirLoadPlan"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate time.Time `json:"lastArrDate" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length float64 `json:"length"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment float64 `json:"moment"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name string `json:"name"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt float64 `json:"netExpWt"`
	// Optional notes or comments about this item.
	Notes string `json:"notes"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos int64 `json:"numPalletPos"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The code denoting the type of material item.
	ProductCode string `json:"productCode"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys string `json:"productSys"`
	// The military branch receiving this item.
	ReceivingBranch string `json:"receivingBranch"`
	// The name of the unit receiving this item.
	ReceivingUnit string `json:"receivingUnit"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool string `json:"scGenTool"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn string `json:"tcn"`
	// The unit line number of this item.
	Uln string `json:"uln"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume float64 `json:"volume"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight float64 `json:"weight"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs time.Time `json:"weightTS" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ScanCode              respjson.Field
		Source                respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		AccSysKeys            respjson.Field
		AccSysNotes           respjson.Field
		AccSystem             respjson.Field
		AccSysValues          respjson.Field
		Airdrop               respjson.Field
		AltDataFormat         respjson.Field
		CargoType             respjson.Field
		CenterlineOffset      respjson.Field
		Cg                    respjson.Field
		CommodityCode         respjson.Field
		CommoditySys          respjson.Field
		Container             respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Departure             respjson.Field
		Destination           respjson.Field
		DvCode                respjson.Field
		Fs                    respjson.Field
		HazCodes              respjson.Field
		Height                respjson.Field
		IDAirLoadPlan         respjson.Field
		ItemContains          respjson.Field
		Keys                  respjson.Field
		LastArrDate           respjson.Field
		Length                respjson.Field
		Moment                respjson.Field
		Name                  respjson.Field
		NetExpWt              respjson.Field
		Notes                 respjson.Field
		NumPalletPos          respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		ProductCode           respjson.Field
		ProductSys            respjson.Field
		ReceivingBranch       respjson.Field
		ReceivingUnit         respjson.Field
		ScGenTool             respjson.Field
		SourceDl              respjson.Field
		Tcn                   respjson.Field
		Uln                   respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		Values                respjson.Field
		Volume                respjson.Field
		Weight                respjson.Field
		WeightTs              respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemTupleResponse) UnmarshalJSON(data []byte) error {
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
type ItemTupleResponseDataMode string

const (
	ItemTupleResponseDataModeReal      ItemTupleResponseDataMode = "REAL"
	ItemTupleResponseDataModeTest      ItemTupleResponseDataMode = "TEST"
	ItemTupleResponseDataModeSimulated ItemTupleResponseDataMode = "SIMULATED"
	ItemTupleResponseDataModeExercise  ItemTupleResponseDataMode = "EXERCISE"
)

type ItemNewParams struct {
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
	DataMode ItemNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes param.Opt[string] `json:"accSysNotes,omitzero"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem param.Opt[string] `json:"accSystem,omitzero"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop param.Opt[bool] `json:"airdrop,omitzero"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat param.Opt[string] `json:"altDataFormat,omitzero"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType param.Opt[string] `json:"cargoType,omitzero"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset param.Opt[float64] `json:"centerlineOffset,omitzero"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg param.Opt[float64] `json:"cg,omitzero"`
	// The classification code of the commodity or group of commodities.
	CommodityCode param.Opt[string] `json:"commodityCode,omitzero"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys param.Opt[string] `json:"commoditySys,omitzero"`
	// Flag indicating this item acts as a container and contains additional items.
	Container param.Opt[bool] `json:"container,omitzero"`
	// The departure code or location where this item has left or is leaving.
	Departure param.Opt[string] `json:"departure,omitzero"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode param.Opt[string] `json:"dvCode,omitzero"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs param.Opt[float64] `json:"fs,omitzero"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height param.Opt[float64] `json:"height,omitzero"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan param.Opt[string] `json:"idAirLoadPlan,omitzero"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate param.Opt[time.Time] `json:"lastArrDate,omitzero" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment param.Opt[float64] `json:"moment,omitzero"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name param.Opt[string] `json:"name,omitzero"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	// Optional notes or comments about this item.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos param.Opt[int64] `json:"numPalletPos,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The code denoting the type of material item.
	ProductCode param.Opt[string] `json:"productCode,omitzero"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys param.Opt[string] `json:"productSys,omitzero"`
	// The military branch receiving this item.
	ReceivingBranch param.Opt[string] `json:"receivingBranch,omitzero"`
	// The name of the unit receiving this item.
	ReceivingUnit param.Opt[string] `json:"receivingUnit,omitzero"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool param.Opt[string] `json:"scGenTool,omitzero"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn param.Opt[string] `json:"tcn,omitzero"`
	// The unit line number of this item.
	Uln param.Opt[string] `json:"uln,omitzero"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume param.Opt[float64] `json:"volume,omitzero"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs param.Opt[time.Time] `json:"weightTS,omitzero" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys,omitzero"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues,omitzero"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes,omitzero"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains,omitzero"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys,omitzero"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ItemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemNewParams
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
type ItemNewParamsDataMode string

const (
	ItemNewParamsDataModeReal      ItemNewParamsDataMode = "REAL"
	ItemNewParamsDataModeTest      ItemNewParamsDataMode = "TEST"
	ItemNewParamsDataModeSimulated ItemNewParamsDataMode = "SIMULATED"
	ItemNewParamsDataModeExercise  ItemNewParamsDataMode = "EXERCISE"
)

type ItemUpdateParams struct {
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
	DataMode ItemUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The tracking identifier of an item or person. May be similar in representation
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes param.Opt[string] `json:"accSysNotes,omitzero"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem param.Opt[string] `json:"accSystem,omitzero"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop param.Opt[bool] `json:"airdrop,omitzero"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat param.Opt[string] `json:"altDataFormat,omitzero"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType param.Opt[string] `json:"cargoType,omitzero"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset param.Opt[float64] `json:"centerlineOffset,omitzero"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg param.Opt[float64] `json:"cg,omitzero"`
	// The classification code of the commodity or group of commodities.
	CommodityCode param.Opt[string] `json:"commodityCode,omitzero"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys param.Opt[string] `json:"commoditySys,omitzero"`
	// Flag indicating this item acts as a container and contains additional items.
	Container param.Opt[bool] `json:"container,omitzero"`
	// The departure code or location where this item has left or is leaving.
	Departure param.Opt[string] `json:"departure,omitzero"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode param.Opt[string] `json:"dvCode,omitzero"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs param.Opt[float64] `json:"fs,omitzero"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height param.Opt[float64] `json:"height,omitzero"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan param.Opt[string] `json:"idAirLoadPlan,omitzero"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate param.Opt[time.Time] `json:"lastArrDate,omitzero" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment param.Opt[float64] `json:"moment,omitzero"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name param.Opt[string] `json:"name,omitzero"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	// Optional notes or comments about this item.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos param.Opt[int64] `json:"numPalletPos,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The code denoting the type of material item.
	ProductCode param.Opt[string] `json:"productCode,omitzero"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys param.Opt[string] `json:"productSys,omitzero"`
	// The military branch receiving this item.
	ReceivingBranch param.Opt[string] `json:"receivingBranch,omitzero"`
	// The name of the unit receiving this item.
	ReceivingUnit param.Opt[string] `json:"receivingUnit,omitzero"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool param.Opt[string] `json:"scGenTool,omitzero"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn param.Opt[string] `json:"tcn,omitzero"`
	// The unit line number of this item.
	Uln param.Opt[string] `json:"uln,omitzero"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume param.Opt[float64] `json:"volume,omitzero"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs param.Opt[time.Time] `json:"weightTS,omitzero" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys,omitzero"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues,omitzero"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes,omitzero"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains,omitzero"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys,omitzero"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ItemUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemUpdateParams
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
type ItemUpdateParamsDataMode string

const (
	ItemUpdateParamsDataModeReal      ItemUpdateParamsDataMode = "REAL"
	ItemUpdateParamsDataModeTest      ItemUpdateParamsDataMode = "TEST"
	ItemUpdateParamsDataModeSimulated ItemUpdateParamsDataMode = "SIMULATED"
	ItemUpdateParamsDataModeExercise  ItemUpdateParamsDataMode = "EXERCISE"
)

type ItemListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemListParams]'s query parameters as `url.Values`.
func (r ItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemCountParams]'s query parameters as `url.Values`.
func (r ItemCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemGetParams]'s query parameters as `url.Values`.
func (r ItemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemTupleParams]'s query parameters as `url.Values`.
func (r ItemTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemUnvalidatedPublishParams struct {
	Body []ItemUnvalidatedPublishParamsBody
	paramObj
}

func (r ItemUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// The properties ClassificationMarking, DataMode, ScanCode, Source, Type are
// required.
type ItemUnvalidatedPublishParamsBody struct {
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
	// of a barcode or UPC. If no scanCode or tracking number equivalent is available,
	// 'NONE' should be used.
	ScanCode string `json:"scanCode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The item type of this record (e.g. EQUIPMENT, CARGO, PASSENGER).
	Type string `json:"type,required"`
	// Unique identifier of the record, auto-generated by the system if not provided on
	// create operations.
	ID param.Opt[string] `json:"id,omitzero"`
	// Additional data required to find this item in the accepting system.
	AccSysNotes param.Opt[string] `json:"accSysNotes,omitzero"`
	// Name of the system that accepted this item from a customer. Where a user or
	// application could go look for additional information.
	AccSystem param.Opt[string] `json:"accSystem,omitzero"`
	// Flag indicating this item is planned to be airdropped. Applicable for cargo and
	// passenger item types only.
	Airdrop param.Opt[bool] `json:"airdrop,omitzero"`
	// Name of the additional data format so downstream consuming applications can know
	// how to parse it. Typically includes the source system name and the format name.
	AltDataFormat param.Opt[string] `json:"altDataFormat,omitzero"`
	// The type of cargo (e.g. PALLET, ROLLING STOCK, LOOSE, OTHER). Applicable for
	// cargo item types only.
	CargoType param.Opt[string] `json:"cargoType,omitzero"`
	// How far left or right of centerline is the item in meters. Applicable for cargo
	// and passenger item types only.
	CenterlineOffset param.Opt[float64] `json:"centerlineOffset,omitzero"`
	// Center of gravity position of the item, measured from the item's front datum, in
	// centimeters.
	Cg param.Opt[float64] `json:"cg,omitzero"`
	// The classification code of the commodity or group of commodities.
	CommodityCode param.Opt[string] `json:"commodityCode,omitzero"`
	// The classification system denoting the commodity code, commodityCode (e.g. AIR,
	// WATER, NMFC, UFC, STCC, DODUNQ, etc.).
	CommoditySys param.Opt[string] `json:"commoditySys,omitzero"`
	// Flag indicating this item acts as a container and contains additional items.
	Container param.Opt[bool] `json:"container,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The departure code or location where this item has left or is leaving.
	Departure param.Opt[string] `json:"departure,omitzero"`
	// The destination of the item, typically an ICAO or port code. Applicable for
	// cargo and passenger item types only.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// United States Distinguished Visitor Code, only applicable to people.
	DvCode param.Opt[string] `json:"dvCode,omitzero"`
	// The fuselage station of the item measured from the reference datum, in
	// centimeters. Applicable for cargo and passenger item types only.
	Fs param.Opt[float64] `json:"fs,omitzero"`
	// Height of the cargo in meters. Applicable for cargo item types only.
	Height param.Opt[float64] `json:"height,omitzero"`
	// The UDL ID of the air load plan this item is associated with.
	IDAirLoadPlan param.Opt[string] `json:"idAirLoadPlan,omitzero"`
	// The latest acceptable arrival date of the item at the destination, in ISO 8601
	// date-only format (e.g. YYYY-MM-DD).
	LastArrDate param.Opt[time.Time] `json:"lastArrDate,omitzero" format:"date"`
	// Length of the cargo in meters. Applicable for cargo item types only.
	Length param.Opt[float64] `json:"length,omitzero"`
	// Moment of the item in Newton-meters. Applicable for equipment and cargo item
	// types only.
	Moment param.Opt[float64] `json:"moment,omitzero"`
	// Name of the item. Applicable for equipment and cargo item types only.
	Name param.Opt[string] `json:"name,omitzero"`
	// Net explosive weight of the item, in kilograms.
	NetExpWt param.Opt[float64] `json:"netExpWt,omitzero"`
	// Optional notes or comments about this item.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Number of pallet positions or equivalent on the aircraft, ship, or conveyance
	// equipment that this item occupies.
	NumPalletPos param.Opt[int64] `json:"numPalletPos,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The code denoting the type of material item.
	ProductCode param.Opt[string] `json:"productCode,omitzero"`
	// The assigning system that denotes the type of material item, productCode (e.g.
	// NSN-national stock number, NDC-national drug code, MPN-manufacturer part number,
	// etc.).
	ProductSys param.Opt[string] `json:"productSys,omitzero"`
	// The military branch receiving this item.
	ReceivingBranch param.Opt[string] `json:"receivingBranch,omitzero"`
	// The name of the unit receiving this item.
	ReceivingUnit param.Opt[string] `json:"receivingUnit,omitzero"`
	// The algorithm name or standard that generated the scanCode (e.g. UPC-A, EAN-13,
	// GTIN, SSCC, bID, JAN, etc.).
	ScGenTool param.Opt[string] `json:"scGenTool,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// Transportation Control Number of the cargo. Applicable for cargo item types
	// only.
	Tcn param.Opt[string] `json:"tcn,omitzero"`
	// The unit line number of this item.
	Uln param.Opt[string] `json:"uln,omitzero"`
	// The volume of the item, in cubic meters. Applicable for cargo item types only.
	Volume param.Opt[float64] `json:"volume,omitzero"`
	// Weight of the item in kilograms (if item is a passenger, include on-person
	// bags).
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Timestamp when the weight was taken, in ISO 8601 UTC format with millisecond
	// precision.
	WeightTs param.Opt[time.Time] `json:"weightTS,omitzero" format:"date-time"`
	// Width of the cargo in meters. Applicable for cargo item types only.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Array of keys that may be associated to the accepting system data. The entries
	// in this array must correspond to the position index in accSysValues array. This
	// array must be the same length as accSysValues.
	AccSysKeys []string `json:"accSysKeys,omitzero"`
	// Array of values for the keys that may be associated to the accepting system
	// data. The entries in this array must correspond to the position index in
	// accSysKeys array. This array must be the same length as accSysKeys.
	AccSysValues []string `json:"accSysValues,omitzero"`
	// Array of UN hazard classes or division codes that apply to this item.
	HazCodes []float64 `json:"hazCodes,omitzero"`
	// Array of tracking identifiers that are contained within this item.
	ItemContains []string `json:"itemContains,omitzero"`
	// Array of keys that may be associated to this item. The entries in this array
	// must correspond to the position index in the values array. This array must be
	// the same length as values..
	Keys []string `json:"keys,omitzero"`
	// Array of values for the keys that may be associated to this tracked item. The
	// entries in this array must correspond to the position index in the keys array.
	// This array must be the same length as keys.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ItemUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ItemUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ItemUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
