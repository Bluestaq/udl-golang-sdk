// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// TransponderService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransponderService] method instead.
type TransponderService struct {
	Options []option.RequestOption
}

// NewTransponderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransponderService(opts ...option.RequestOption) (r TransponderService) {
	r = TransponderService{}
	r.Options = opts
	return
}

// Service operation to take a single Transponder as a POST body and ingest into
// the database. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance. A Comm payload may have multiple
// transponders and a transponder may have many channels.
func (r *TransponderService) New(ctx context.Context, body TransponderNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/transponder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single Transponder. A specific role is required to
// perform this service operation. Please contact the UDL team for assistance. A
// Comm payload may have multiple transponders and a transponder may have many
// channels.
func (r *TransponderService) Update(ctx context.Context, id string, body TransponderUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/transponder/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *TransponderService) List(ctx context.Context, query TransponderListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[TransponderListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/transponder"
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
func (r *TransponderService) ListAutoPaging(ctx context.Context, query TransponderListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[TransponderListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a Transponder object specified by the passed ID path
// parameter. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance. A Comm payload may have multiple
// transponders and a transponder may have many channels.
func (r *TransponderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/transponder/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *TransponderService) Count(ctx context.Context, query TransponderCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/transponder/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single Transponder record by its unique ID passed as
// a path parameter. A Comm payload may have multiple transponders and a
// transponder may have many channels.
func (r *TransponderService) Get(ctx context.Context, id string, query TransponderGetParams, opts ...option.RequestOption) (res *TransponderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/transponder/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *TransponderService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *TransponderQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/transponder/queryhelp"
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
func (r *TransponderService) Tuple(ctx context.Context, query TransponderTupleParams, opts ...option.RequestOption) (res *[]TransponderTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/transponder/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A transponder receives and transmits radio signals at a prescribed frequency
// range. A communications satellite's transponder is the series of interconnected
// units that form a communications channel between the receiving and the
// transmitting antennas. It is mainly used in satellite communication to transfer
// the received signals.
type TransponderListResponse struct {
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
	DataMode TransponderListResponseDataMode `json:"dataMode,required"`
	// ID of the parent Comm object for this transponder.
	IDComm string `json:"idComm,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Forward error correction, e.g. 0=Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 = 5/6, 5 =
	// 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	Fec int64 `json:"fec"`
	// Format.
	Format string `json:"format"`
	// Transponder modulation, e.g. Auto, QPSK, 8PSK.
	Modulation string `json:"modulation"`
	// Optional name of the transponder.
	Name string `json:"name"`
	// Optional external network id as provided data source.
	Nid string `json:"nid"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Symbol rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per time unit using a digitally modulated
	// signal or a line code. Also measured in Hz.
	SymbolRate float64 `json:"symbolRate"`
	// Transponder system, e.g. DVB-S, DVB-S2.
	System string `json:"system"`
	// Optional external transponder id as provided data source.
	Tid string `json:"tid"`
	// Transponder Translation Factor. This is the frequency difference between the
	// uplink received by a satellite, and the downlink transmitted back. It varies
	// satellite to satellite dependent on the mission.
	Ttf float64 `json:"ttf"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDComm                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Fec                   respjson.Field
		Format                respjson.Field
		Modulation            respjson.Field
		Name                  respjson.Field
		Nid                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SymbolRate            respjson.Field
		System                respjson.Field
		Tid                   respjson.Field
		Ttf                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransponderListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransponderListResponse) UnmarshalJSON(data []byte) error {
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
type TransponderListResponseDataMode string

const (
	TransponderListResponseDataModeReal      TransponderListResponseDataMode = "REAL"
	TransponderListResponseDataModeTest      TransponderListResponseDataMode = "TEST"
	TransponderListResponseDataModeSimulated TransponderListResponseDataMode = "SIMULATED"
	TransponderListResponseDataModeExercise  TransponderListResponseDataMode = "EXERCISE"
)

// A transponder receives and transmits radio signals at a prescribed frequency
// range. A communications satellite's transponder is the series of interconnected
// units that form a communications channel between the receiving and the
// transmitting antennas. It is mainly used in satellite communication to transfer
// the received signals.
type TransponderGetResponse struct {
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
	DataMode TransponderGetResponseDataMode `json:"dataMode,required"`
	// ID of the parent Comm object for this transponder.
	IDComm string `json:"idComm,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Collection of Channels for this Transponder.
	Channels []ChannelFull `json:"channels"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Forward error correction, e.g. 0=Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 = 5/6, 5 =
	// 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	Fec int64 `json:"fec"`
	// Format.
	Format string `json:"format"`
	// Transponder modulation, e.g. Auto, QPSK, 8PSK.
	Modulation string `json:"modulation"`
	// Optional name of the transponder.
	Name string `json:"name"`
	// Optional external network id as provided data source.
	Nid string `json:"nid"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Symbol rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per time unit using a digitally modulated
	// signal or a line code. Also measured in Hz.
	SymbolRate float64 `json:"symbolRate"`
	// Transponder system, e.g. DVB-S, DVB-S2.
	System string `json:"system"`
	// Optional external transponder id as provided data source.
	Tid string `json:"tid"`
	// Transponder Translation Factor. This is the frequency difference between the
	// uplink received by a satellite, and the downlink transmitted back. It varies
	// satellite to satellite dependent on the mission.
	Ttf float64 `json:"ttf"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDComm                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Channels              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Fec                   respjson.Field
		Format                respjson.Field
		Modulation            respjson.Field
		Name                  respjson.Field
		Nid                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SymbolRate            respjson.Field
		System                respjson.Field
		Tid                   respjson.Field
		Ttf                   respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransponderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransponderGetResponse) UnmarshalJSON(data []byte) error {
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
type TransponderGetResponseDataMode string

const (
	TransponderGetResponseDataModeReal      TransponderGetResponseDataMode = "REAL"
	TransponderGetResponseDataModeTest      TransponderGetResponseDataMode = "TEST"
	TransponderGetResponseDataModeSimulated TransponderGetResponseDataMode = "SIMULATED"
	TransponderGetResponseDataModeExercise  TransponderGetResponseDataMode = "EXERCISE"
)

type TransponderQueryhelpResponse struct {
	AodrSupported         bool                                    `json:"aodrSupported"`
	ClassificationMarking string                                  `json:"classificationMarking"`
	Description           string                                  `json:"description"`
	HistorySupported      bool                                    `json:"historySupported"`
	Name                  string                                  `json:"name"`
	Parameters            []TransponderQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                `json:"requiredRoles"`
	RestSupported         bool                                    `json:"restSupported"`
	SortSupported         bool                                    `json:"sortSupported"`
	TypeName              string                                  `json:"typeName"`
	Uri                   string                                  `json:"uri"`
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
func (r TransponderQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *TransponderQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransponderQueryhelpResponseParameter struct {
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
func (r TransponderQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *TransponderQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A transponder receives and transmits radio signals at a prescribed frequency
// range. A communications satellite's transponder is the series of interconnected
// units that form a communications channel between the receiving and the
// transmitting antennas. It is mainly used in satellite communication to transfer
// the received signals.
type TransponderTupleResponse struct {
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
	DataMode TransponderTupleResponseDataMode `json:"dataMode,required"`
	// ID of the parent Comm object for this transponder.
	IDComm string `json:"idComm,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Collection of Channels for this Transponder.
	Channels []ChannelFull `json:"channels"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Forward error correction, e.g. 0=Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 = 5/6, 5 =
	// 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	Fec int64 `json:"fec"`
	// Format.
	Format string `json:"format"`
	// Transponder modulation, e.g. Auto, QPSK, 8PSK.
	Modulation string `json:"modulation"`
	// Optional name of the transponder.
	Name string `json:"name"`
	// Optional external network id as provided data source.
	Nid string `json:"nid"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Symbol rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per time unit using a digitally modulated
	// signal or a line code. Also measured in Hz.
	SymbolRate float64 `json:"symbolRate"`
	// Transponder system, e.g. DVB-S, DVB-S2.
	System string `json:"system"`
	// Optional external transponder id as provided data source.
	Tid string `json:"tid"`
	// Transponder Translation Factor. This is the frequency difference between the
	// uplink received by a satellite, and the downlink transmitted back. It varies
	// satellite to satellite dependent on the mission.
	Ttf float64 `json:"ttf"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		IDComm                respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Channels              respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Fec                   respjson.Field
		Format                respjson.Field
		Modulation            respjson.Field
		Name                  respjson.Field
		Nid                   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		SymbolRate            respjson.Field
		System                respjson.Field
		Tid                   respjson.Field
		Ttf                   respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransponderTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *TransponderTupleResponse) UnmarshalJSON(data []byte) error {
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
type TransponderTupleResponseDataMode string

const (
	TransponderTupleResponseDataModeReal      TransponderTupleResponseDataMode = "REAL"
	TransponderTupleResponseDataModeTest      TransponderTupleResponseDataMode = "TEST"
	TransponderTupleResponseDataModeSimulated TransponderTupleResponseDataMode = "SIMULATED"
	TransponderTupleResponseDataModeExercise  TransponderTupleResponseDataMode = "EXERCISE"
)

type TransponderNewParams struct {
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
	DataMode TransponderNewParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent Comm object for this transponder.
	IDComm string `json:"idComm,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Forward error correction, e.g. 0=Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 = 5/6, 5 =
	// 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	Fec param.Opt[int64] `json:"fec,omitzero"`
	// Format.
	Format param.Opt[string] `json:"format,omitzero"`
	// Transponder modulation, e.g. Auto, QPSK, 8PSK.
	Modulation param.Opt[string] `json:"modulation,omitzero"`
	// Optional name of the transponder.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional external network id as provided data source.
	Nid param.Opt[string] `json:"nid,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Symbol rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per time unit using a digitally modulated
	// signal or a line code. Also measured in Hz.
	SymbolRate param.Opt[float64] `json:"symbolRate,omitzero"`
	// Transponder system, e.g. DVB-S, DVB-S2.
	System param.Opt[string] `json:"system,omitzero"`
	// Optional external transponder id as provided data source.
	Tid param.Opt[string] `json:"tid,omitzero"`
	// Transponder Translation Factor. This is the frequency difference between the
	// uplink received by a satellite, and the downlink transmitted back. It varies
	// satellite to satellite dependent on the mission.
	Ttf param.Opt[float64] `json:"ttf,omitzero"`
	paramObj
}

func (r TransponderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransponderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransponderNewParams) UnmarshalJSON(data []byte) error {
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
type TransponderNewParamsDataMode string

const (
	TransponderNewParamsDataModeReal      TransponderNewParamsDataMode = "REAL"
	TransponderNewParamsDataModeTest      TransponderNewParamsDataMode = "TEST"
	TransponderNewParamsDataModeSimulated TransponderNewParamsDataMode = "SIMULATED"
	TransponderNewParamsDataModeExercise  TransponderNewParamsDataMode = "EXERCISE"
)

type TransponderUpdateParams struct {
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
	DataMode TransponderUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// ID of the parent Comm object for this transponder.
	IDComm string `json:"idComm,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Forward error correction, e.g. 0=Auto, 1 = 1/2, 2 = 2/3, 3 = 3/4, 4 = 5/6, 5 =
	// 7/8, 6 = 8/9, 7 = 3/5, 8 = 4/5, 9 = 9/10, 15 = None.
	Fec param.Opt[int64] `json:"fec,omitzero"`
	// Format.
	Format param.Opt[string] `json:"format,omitzero"`
	// Transponder modulation, e.g. Auto, QPSK, 8PSK.
	Modulation param.Opt[string] `json:"modulation,omitzero"`
	// Optional name of the transponder.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional external network id as provided data source.
	Nid param.Opt[string] `json:"nid,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Symbol rate is the number of symbol changes, waveform changes, or signaling
	// events, across the transmission medium per time unit using a digitally modulated
	// signal or a line code. Also measured in Hz.
	SymbolRate param.Opt[float64] `json:"symbolRate,omitzero"`
	// Transponder system, e.g. DVB-S, DVB-S2.
	System param.Opt[string] `json:"system,omitzero"`
	// Optional external transponder id as provided data source.
	Tid param.Opt[string] `json:"tid,omitzero"`
	// Transponder Translation Factor. This is the frequency difference between the
	// uplink received by a satellite, and the downlink transmitted back. It varies
	// satellite to satellite dependent on the mission.
	Ttf param.Opt[float64] `json:"ttf,omitzero"`
	paramObj
}

func (r TransponderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransponderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransponderUpdateParams) UnmarshalJSON(data []byte) error {
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
type TransponderUpdateParamsDataMode string

const (
	TransponderUpdateParamsDataModeReal      TransponderUpdateParamsDataMode = "REAL"
	TransponderUpdateParamsDataModeTest      TransponderUpdateParamsDataMode = "TEST"
	TransponderUpdateParamsDataModeSimulated TransponderUpdateParamsDataMode = "SIMULATED"
	TransponderUpdateParamsDataModeExercise  TransponderUpdateParamsDataMode = "EXERCISE"
)

type TransponderListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransponderListParams]'s query parameters as `url.Values`.
func (r TransponderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransponderCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransponderCountParams]'s query parameters as `url.Values`.
func (r TransponderCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransponderGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransponderGetParams]'s query parameters as `url.Values`.
func (r TransponderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransponderTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransponderTupleParams]'s query parameters as `url.Values`.
func (r TransponderTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
