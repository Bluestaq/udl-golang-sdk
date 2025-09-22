// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
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

// VesselService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVesselService] method instead.
type VesselService struct {
	Options []option.RequestOption
}

// NewVesselService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVesselService(opts ...option.RequestOption) (r VesselService) {
	r = VesselService{}
	r.Options = opts
	return
}

// Service operation to take a single vessel record as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *VesselService) New(ctx context.Context, body VesselNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/vessel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single vessel record. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *VesselService) Update(ctx context.Context, id string, body VesselUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/vessel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *VesselService) List(ctx context.Context, query VesselListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[VesselListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/vessel"
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
func (r *VesselService) ListAutoPaging(ctx context.Context, query VesselListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[VesselListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *VesselService) Count(ctx context.Context, query VesselCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/vessel/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// vessel records as a POST body and ingest into the database. This operation is
// not intended to be used for automated feeds into UDL. Data providers should
// contact the UDL team for specific role assignments and for instructions on
// setting up a permanent feed through an alternate mechanism.
func (r *VesselService) NewBulk(ctx context.Context, body VesselNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/vessel/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single vessel record by its unique ID passed as a
// path parameter.
func (r *VesselService) Get(ctx context.Context, id string, query VesselGetParams, opts ...option.RequestOption) (res *VesselGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/vessel/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *VesselService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *VesselQueryhelpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/vessel/queryhelp"
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
func (r *VesselService) Tuple(ctx context.Context, query VesselTupleParams, opts ...option.RequestOption) (res *[]VesselTupleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "udl/vessel/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This service provides operations for manipulation and querying of maritime
// Vessel data. Vessel contains the static data of the specific vessel: mmsi,
// cruise speed, max speed, etc.
type VesselListResponse struct {
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
	DataMode VesselListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The original source Id for this vessel.
	AltVesselID string `json:"altVesselId"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign string `json:"callsign"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The date this vessel was first seen.
	FirstSeen time.Time `json:"firstSeen" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum string `json:"hullNum"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization string `json:"idOrganization"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon int64 `json:"imon"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length float64 `json:"length"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught float64 `json:"maxDraught"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed float64 `json:"maxSpeed"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi string `json:"mmsi"`
	// The number of blades per shaft for this vessel.
	NumBlades int64 `json:"numBlades"`
	// The number of shafts on this vessel.
	NumShafts int64 `json:"numShafts"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of propulsion employed by this vessel.
	PropType string `json:"propType"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum string `json:"sconum"`
	// The status of this vessel.
	Status string `json:"status"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType string `json:"sternType"`
	// The shipbuilder who built this vessel.
	VesselBuilder string `json:"vesselBuilder"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass string `json:"vesselClass"`
	// Further description or explanation of the vessel or type.
	VesselDescription string `json:"vesselDescription"`
	// The flag of the subject vessel.
	VesselFlag string `json:"vesselFlag"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName string `json:"vesselName"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType string `json:"vesselType"`
	// The weight in tons, of this vessel.
	VslWt float64 `json:"vslWt"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width float64 `json:"width"`
	// Year the vessel went into service.
	YearBuilt string `json:"yearBuilt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltVesselID           respjson.Field
		Callsign              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		FirstSeen             respjson.Field
		HullNum               respjson.Field
		IDEntity              respjson.Field
		IDOrganization        respjson.Field
		Imon                  respjson.Field
		Length                respjson.Field
		MaxDraught            respjson.Field
		MaxSpeed              respjson.Field
		Mmsi                  respjson.Field
		NumBlades             respjson.Field
		NumShafts             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PropType              respjson.Field
		Sconum                respjson.Field
		Status                respjson.Field
		SternType             respjson.Field
		VesselBuilder         respjson.Field
		VesselClass           respjson.Field
		VesselDescription     respjson.Field
		VesselFlag            respjson.Field
		VesselName            respjson.Field
		VesselType            respjson.Field
		VslWt                 respjson.Field
		Width                 respjson.Field
		YearBuilt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VesselListResponse) RawJSON() string { return r.JSON.raw }
func (r *VesselListResponse) UnmarshalJSON(data []byte) error {
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
type VesselListResponseDataMode string

const (
	VesselListResponseDataModeReal      VesselListResponseDataMode = "REAL"
	VesselListResponseDataModeTest      VesselListResponseDataMode = "TEST"
	VesselListResponseDataModeSimulated VesselListResponseDataMode = "SIMULATED"
	VesselListResponseDataModeExercise  VesselListResponseDataMode = "EXERCISE"
)

// This service provides operations for manipulation and querying of maritime
// Vessel data. Vessel contains the static data of the specific vessel: mmsi,
// cruise speed, max speed, etc.
type VesselGetResponse struct {
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
	DataMode VesselGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The original source Id for this vessel.
	AltVesselID string `json:"altVesselId"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign string `json:"callsign"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// The date this vessel was first seen.
	FirstSeen time.Time `json:"firstSeen" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum string `json:"hullNum"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization string `json:"idOrganization"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon int64 `json:"imon"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length float64 `json:"length"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught float64 `json:"maxDraught"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed float64 `json:"maxSpeed"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi string `json:"mmsi"`
	// The number of blades per shaft for this vessel.
	NumBlades int64 `json:"numBlades"`
	// The number of shafts on this vessel.
	NumShafts int64 `json:"numShafts"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of propulsion employed by this vessel.
	PropType string `json:"propType"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum string `json:"sconum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The status of this vessel.
	Status string `json:"status"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType string `json:"sternType"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The shipbuilder who built this vessel.
	VesselBuilder string `json:"vesselBuilder"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass string `json:"vesselClass"`
	// Further description or explanation of the vessel or type.
	VesselDescription string `json:"vesselDescription"`
	// The flag of the subject vessel.
	VesselFlag string `json:"vesselFlag"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName string `json:"vesselName"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType string `json:"vesselType"`
	// The weight in tons, of this vessel.
	VslWt float64 `json:"vslWt"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width float64 `json:"width"`
	// Year the vessel went into service.
	YearBuilt string `json:"yearBuilt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltVesselID           respjson.Field
		Callsign              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Entity                respjson.Field
		FirstSeen             respjson.Field
		HullNum               respjson.Field
		IDEntity              respjson.Field
		IDOrganization        respjson.Field
		Imon                  respjson.Field
		Length                respjson.Field
		MaxDraught            respjson.Field
		MaxSpeed              respjson.Field
		Mmsi                  respjson.Field
		NumBlades             respjson.Field
		NumShafts             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PropType              respjson.Field
		Sconum                respjson.Field
		SourceDl              respjson.Field
		Status                respjson.Field
		SternType             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VesselBuilder         respjson.Field
		VesselClass           respjson.Field
		VesselDescription     respjson.Field
		VesselFlag            respjson.Field
		VesselName            respjson.Field
		VesselType            respjson.Field
		VslWt                 respjson.Field
		Width                 respjson.Field
		YearBuilt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VesselGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VesselGetResponse) UnmarshalJSON(data []byte) error {
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
type VesselGetResponseDataMode string

const (
	VesselGetResponseDataModeReal      VesselGetResponseDataMode = "REAL"
	VesselGetResponseDataModeTest      VesselGetResponseDataMode = "TEST"
	VesselGetResponseDataModeSimulated VesselGetResponseDataMode = "SIMULATED"
	VesselGetResponseDataModeExercise  VesselGetResponseDataMode = "EXERCISE"
)

type VesselQueryhelpResponse struct {
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
func (r VesselQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *VesselQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This service provides operations for manipulation and querying of maritime
// Vessel data. Vessel contains the static data of the specific vessel: mmsi,
// cruise speed, max speed, etc.
type VesselTupleResponse struct {
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
	DataMode VesselTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The original source Id for this vessel.
	AltVesselID string `json:"altVesselId"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign string `json:"callsign"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity shared.EntityFull `json:"entity"`
	// The date this vessel was first seen.
	FirstSeen time.Time `json:"firstSeen" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum string `json:"hullNum"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity string `json:"idEntity"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization string `json:"idOrganization"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon int64 `json:"imon"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length float64 `json:"length"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught float64 `json:"maxDraught"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed float64 `json:"maxSpeed"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi string `json:"mmsi"`
	// The number of blades per shaft for this vessel.
	NumBlades int64 `json:"numBlades"`
	// The number of shafts on this vessel.
	NumShafts int64 `json:"numShafts"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The type of propulsion employed by this vessel.
	PropType string `json:"propType"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum string `json:"sconum"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The status of this vessel.
	Status string `json:"status"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType string `json:"sternType"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// The shipbuilder who built this vessel.
	VesselBuilder string `json:"vesselBuilder"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass string `json:"vesselClass"`
	// Further description or explanation of the vessel or type.
	VesselDescription string `json:"vesselDescription"`
	// The flag of the subject vessel.
	VesselFlag string `json:"vesselFlag"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName string `json:"vesselName"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType string `json:"vesselType"`
	// The weight in tons, of this vessel.
	VslWt float64 `json:"vslWt"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width float64 `json:"width"`
	// Year the vessel went into service.
	YearBuilt string `json:"yearBuilt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		AltVesselID           respjson.Field
		Callsign              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Entity                respjson.Field
		FirstSeen             respjson.Field
		HullNum               respjson.Field
		IDEntity              respjson.Field
		IDOrganization        respjson.Field
		Imon                  respjson.Field
		Length                respjson.Field
		MaxDraught            respjson.Field
		MaxSpeed              respjson.Field
		Mmsi                  respjson.Field
		NumBlades             respjson.Field
		NumShafts             respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		PropType              respjson.Field
		Sconum                respjson.Field
		SourceDl              respjson.Field
		Status                respjson.Field
		SternType             respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		VesselBuilder         respjson.Field
		VesselClass           respjson.Field
		VesselDescription     respjson.Field
		VesselFlag            respjson.Field
		VesselName            respjson.Field
		VesselType            respjson.Field
		VslWt                 respjson.Field
		Width                 respjson.Field
		YearBuilt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VesselTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *VesselTupleResponse) UnmarshalJSON(data []byte) error {
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
type VesselTupleResponseDataMode string

const (
	VesselTupleResponseDataModeReal      VesselTupleResponseDataMode = "REAL"
	VesselTupleResponseDataModeTest      VesselTupleResponseDataMode = "TEST"
	VesselTupleResponseDataModeSimulated VesselTupleResponseDataMode = "SIMULATED"
	VesselTupleResponseDataModeExercise  VesselTupleResponseDataMode = "EXERCISE"
)

type VesselNewParams struct {
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
	DataMode VesselNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The original source Id for this vessel.
	AltVesselID param.Opt[string] `json:"altVesselId,omitzero"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// The date this vessel was first seen.
	FirstSeen param.Opt[time.Time] `json:"firstSeen,omitzero" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum param.Opt[string] `json:"hullNum,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization param.Opt[string] `json:"idOrganization,omitzero"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon param.Opt[int64] `json:"imon,omitzero"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi param.Opt[string] `json:"mmsi,omitzero"`
	// The number of blades per shaft for this vessel.
	NumBlades param.Opt[int64] `json:"numBlades,omitzero"`
	// The number of shafts on this vessel.
	NumShafts param.Opt[int64] `json:"numShafts,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The type of propulsion employed by this vessel.
	PropType param.Opt[string] `json:"propType,omitzero"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum param.Opt[string] `json:"sconum,omitzero"`
	// The status of this vessel.
	Status param.Opt[string] `json:"status,omitzero"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType param.Opt[string] `json:"sternType,omitzero"`
	// The shipbuilder who built this vessel.
	VesselBuilder param.Opt[string] `json:"vesselBuilder,omitzero"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass param.Opt[string] `json:"vesselClass,omitzero"`
	// Further description or explanation of the vessel or type.
	VesselDescription param.Opt[string] `json:"vesselDescription,omitzero"`
	// The flag of the subject vessel.
	VesselFlag param.Opt[string] `json:"vesselFlag,omitzero"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName param.Opt[string] `json:"vesselName,omitzero"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType param.Opt[string] `json:"vesselType,omitzero"`
	// The weight in tons, of this vessel.
	VslWt param.Opt[float64] `json:"vslWt,omitzero"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Year the vessel went into service.
	YearBuilt param.Opt[string] `json:"yearBuilt,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r VesselNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VesselNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VesselNewParams) UnmarshalJSON(data []byte) error {
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
type VesselNewParamsDataMode string

const (
	VesselNewParamsDataModeReal      VesselNewParamsDataMode = "REAL"
	VesselNewParamsDataModeTest      VesselNewParamsDataMode = "TEST"
	VesselNewParamsDataModeSimulated VesselNewParamsDataMode = "SIMULATED"
	VesselNewParamsDataModeExercise  VesselNewParamsDataMode = "EXERCISE"
)

type VesselUpdateParams struct {
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
	DataMode VesselUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The original source Id for this vessel.
	AltVesselID param.Opt[string] `json:"altVesselId,omitzero"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// The date this vessel was first seen.
	FirstSeen param.Opt[time.Time] `json:"firstSeen,omitzero" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum param.Opt[string] `json:"hullNum,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization param.Opt[string] `json:"idOrganization,omitzero"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon param.Opt[int64] `json:"imon,omitzero"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi param.Opt[string] `json:"mmsi,omitzero"`
	// The number of blades per shaft for this vessel.
	NumBlades param.Opt[int64] `json:"numBlades,omitzero"`
	// The number of shafts on this vessel.
	NumShafts param.Opt[int64] `json:"numShafts,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The type of propulsion employed by this vessel.
	PropType param.Opt[string] `json:"propType,omitzero"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum param.Opt[string] `json:"sconum,omitzero"`
	// The status of this vessel.
	Status param.Opt[string] `json:"status,omitzero"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType param.Opt[string] `json:"sternType,omitzero"`
	// The shipbuilder who built this vessel.
	VesselBuilder param.Opt[string] `json:"vesselBuilder,omitzero"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass param.Opt[string] `json:"vesselClass,omitzero"`
	// Further description or explanation of the vessel or type.
	VesselDescription param.Opt[string] `json:"vesselDescription,omitzero"`
	// The flag of the subject vessel.
	VesselFlag param.Opt[string] `json:"vesselFlag,omitzero"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName param.Opt[string] `json:"vesselName,omitzero"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType param.Opt[string] `json:"vesselType,omitzero"`
	// The weight in tons, of this vessel.
	VslWt param.Opt[float64] `json:"vslWt,omitzero"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Year the vessel went into service.
	YearBuilt param.Opt[string] `json:"yearBuilt,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r VesselUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VesselUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VesselUpdateParams) UnmarshalJSON(data []byte) error {
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
type VesselUpdateParamsDataMode string

const (
	VesselUpdateParamsDataModeReal      VesselUpdateParamsDataMode = "REAL"
	VesselUpdateParamsDataModeTest      VesselUpdateParamsDataMode = "TEST"
	VesselUpdateParamsDataModeSimulated VesselUpdateParamsDataMode = "SIMULATED"
	VesselUpdateParamsDataModeExercise  VesselUpdateParamsDataMode = "EXERCISE"
)

type VesselListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VesselListParams]'s query parameters as `url.Values`.
func (r VesselListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VesselCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VesselCountParams]'s query parameters as `url.Values`.
func (r VesselCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VesselNewBulkParams struct {
	Body []VesselNewBulkParamsBody
	paramObj
}

func (r VesselNewBulkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *VesselNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// This service provides operations for manipulation and querying of maritime
// Vessel data. Vessel contains the static data of the specific vessel: mmsi,
// cruise speed, max speed, etc.
//
// The properties ClassificationMarking, DataMode, Source are required.
type VesselNewBulkParamsBody struct {
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
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The original source Id for this vessel.
	AltVesselID param.Opt[string] `json:"altVesselId,omitzero"`
	// A uniquely designated identifier for the vessel's transmitter station. All radio
	// transmissions must be individually identified by the call sign. Merchant and
	// naval vessels are assigned call signs by their national licensing authorities.
	Callsign param.Opt[string] `json:"callsign,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The date this vessel was first seen.
	FirstSeen param.Opt[time.Time] `json:"firstSeen,omitzero" format:"date-time"`
	// The vessel hull number designation of this maritime vessel. The hull number is a
	// 1-6 character alphanumeric entry assigned to a ship and painted on the hull.
	HullNum param.Opt[string] `json:"hullNum,omitzero"`
	// Unique identifier of the parent entity. idEntity is required for Put.
	IDEntity param.Opt[string] `json:"idEntity,omitzero"`
	// The UDL ID of the organization that owns the vessel.
	IDOrganization param.Opt[string] `json:"idOrganization,omitzero"`
	// The International Maritime Organization Number of the vessel. IMON is a
	// seven-digit number that uniquely identifies the vessel.
	Imon param.Opt[int64] `json:"imon,omitzero"`
	// The overall length of the vessel, in meters. A value of 511 indicates a vessel
	// length of 511 meters or greater.
	Length param.Opt[float64] `json:"length,omitzero"`
	// The maximum static draught, in meters, of the vessel defined as the distance
	// between the ship’s keel and the waterline of the vessel.
	MaxDraught param.Opt[float64] `json:"maxDraught,omitzero"`
	// The maximum possible speed of this vessel in meters per second.
	MaxSpeed param.Opt[float64] `json:"maxSpeed,omitzero"`
	// The Maritime Mobile Service Identity of the vessel. MMSI is a nine-digit number
	// that identifies the transmitter station of the vessel.
	Mmsi param.Opt[string] `json:"mmsi,omitzero"`
	// The number of blades per shaft for this vessel.
	NumBlades param.Opt[int64] `json:"numBlades,omitzero"`
	// The number of shafts on this vessel.
	NumShafts param.Opt[int64] `json:"numShafts,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The type of propulsion employed by this vessel.
	PropType param.Opt[string] `json:"propType,omitzero"`
	// The Ship Control Number (SCONUM) is a naval vessel identification number
	// (alphanumeric code) assigned by the Office of Naval Intelligence. SCONUM is
	// sometimes referred to as NOIC ID. SCONUMs are typically of the form A#####,
	// where A is an alpha character and # is numerical.
	Sconum param.Opt[string] `json:"sconum,omitzero"`
	// The status of this vessel.
	Status param.Opt[string] `json:"status,omitzero"`
	// The stern type code (Counter, Cruiser) associated with this vessel.
	SternType param.Opt[string] `json:"sternType,omitzero"`
	// The shipbuilder who built this vessel.
	VesselBuilder param.Opt[string] `json:"vesselBuilder,omitzero"`
	// The common name for a group of ships with similar design, usually named for the
	// first vessel of the class.
	VesselClass param.Opt[string] `json:"vesselClass,omitzero"`
	// Further description or explanation of the vessel or type.
	VesselDescription param.Opt[string] `json:"vesselDescription,omitzero"`
	// The flag of the subject vessel.
	VesselFlag param.Opt[string] `json:"vesselFlag,omitzero"`
	// The name of this vessel. Vessel names that exceed the AIS 20 character are
	// shortened (not truncated) to 15 character-spaces, followed by an underscore and
	// the last 4 characters-spaces of the vessel full name.
	VesselName param.Opt[string] `json:"vesselName,omitzero"`
	// The reported ship type (e.g. Passenger, Tanker, Cargo, Other, etc.).
	VesselType param.Opt[string] `json:"vesselType,omitzero"`
	// The weight in tons, of this vessel.
	VslWt param.Opt[float64] `json:"vslWt,omitzero"`
	// The breadth of the vessel, in meters. A value of 63 indicates a vessel breadth
	// of 63 meters or greater.
	Width param.Opt[float64] `json:"width,omitzero"`
	// Year the vessel went into service.
	YearBuilt param.Opt[string] `json:"yearBuilt,omitzero"`
	// An entity is a generic representation of any object within a space/SSA system
	// such as sensors, on-orbit objects, RF Emitters, space craft buses, etc. An
	// entity can have an operating unit, a location (if terrestrial), and statuses.
	Entity EntityIngestParam `json:"entity,omitzero"`
	paramObj
}

func (r VesselNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow VesselNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VesselNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VesselNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type VesselGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VesselGetParams]'s query parameters as `url.Values`.
func (r VesselGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VesselTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VesselTupleParams]'s query parameters as `url.Values`.
func (r VesselTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
