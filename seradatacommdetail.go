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
)

// SeraDataCommDetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeraDataCommDetailService] method instead.
type SeraDataCommDetailService struct {
	Options []option.RequestOption
}

// NewSeraDataCommDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeraDataCommDetailService(opts ...option.RequestOption) (r SeraDataCommDetailService) {
	r = SeraDataCommDetailService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataCommDetails as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeraDataCommDetailService) New(ctx context.Context, body SeraDataCommDetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatacommdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataCommDetails. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeraDataCommDetailService) Update(ctx context.Context, id string, body SeraDataCommDetailUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatacommdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SeraDataCommDetailService) List(ctx context.Context, query SeraDataCommDetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeraDataCommDetailListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/seradatacommdetails"
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
func (r *SeraDataCommDetailService) ListAutoPaging(ctx context.Context, query SeraDataCommDetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeraDataCommDetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataCommDetails specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeraDataCommDetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatacommdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SeraDataCommDetailService) Count(ctx context.Context, query SeraDataCommDetailCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradatacommdetails/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataCommDetails by its unique ID passed as
// a path parameter.
func (r *SeraDataCommDetailService) Get(ctx context.Context, id string, query SeraDataCommDetailGetParams, opts ...option.RequestOption) (res *SeraDataCommDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/seradatacommdetails/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SeraDataCommDetailService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatacommdetails/queryhelp"
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
func (r *SeraDataCommDetailService) Tuple(ctx context.Context, query SeraDataCommDetailTupleParams, opts ...option.RequestOption) (res *[]SeraDataCommDetailTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradatacommdetails/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Seradata-compiled information on communications payloads.
type SeraDataCommDetailListResponse struct {
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
	DataMode SeraDataCommDetailListResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// Comm bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Effective isotropic radiated power in dB.
	Eirp float64 `json:"eirp"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity float64 `json:"estHtsTotalCapacity"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam float64 `json:"estHtsTotalUserDownlinkBandwidthPerBeam"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam float64 `json:"estHtsTotalUserUplinkBandwidthPerBeam"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom float64 `json:"gatewayDownlinkFrom"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo float64 `json:"gatewayDownlinkTo"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom float64 `json:"gatewayUplinkFrom"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo float64 `json:"gatewayUplinkTo"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams int64 `json:"htsNumUserSpotBeams"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam float64 `json:"htsUserDownlinkBandwidthPerBeam"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam float64 `json:"htsUserUplinkBandwidthPerBeam"`
	// UUID of the parent Comm record.
	IDComm string `json:"idComm"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders int64 `json:"num36MhzEquivalentTransponders"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders int64 `json:"numOperationalTransponders"`
	// Comm numSpareTransponders.
	NumSpareTransponders int64 `json:"numSpareTransponders"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Comm polarization.
	Polarization string `json:"polarization"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp float64 `json:"solidStatePowerAmp"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID string `json:"tradeLeaseOrgId"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier float64 `json:"travelingWaveTubeAmplifier"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom float64 `json:"userDownlinkFrom"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo float64 `json:"userDownlinkTo"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom float64 `json:"userUplinkFrom"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo float64 `json:"userUplinkTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		Source                                  respjson.Field
		ID                                      respjson.Field
		Band                                    respjson.Field
		Bandwidth                               respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		Eirp                                    respjson.Field
		EstHtsTotalCapacity                     respjson.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam respjson.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   respjson.Field
		GatewayDownlinkFrom                     respjson.Field
		GatewayDownlinkTo                       respjson.Field
		GatewayUplinkFrom                       respjson.Field
		GatewayUplinkTo                         respjson.Field
		HostedForCompanyOrgID                   respjson.Field
		HtsNumUserSpotBeams                     respjson.Field
		HtsUserDownlinkBandwidthPerBeam         respjson.Field
		HtsUserUplinkBandwidthPerBeam           respjson.Field
		IDComm                                  respjson.Field
		ManufacturerOrgID                       respjson.Field
		Num36MhzEquivalentTransponders          respjson.Field
		NumOperationalTransponders              respjson.Field
		NumSpareTransponders                    respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		PayloadNotes                            respjson.Field
		Polarization                            respjson.Field
		SolidStatePowerAmp                      respjson.Field
		SpacecraftID                            respjson.Field
		TradeLeaseOrgID                         respjson.Field
		TravelingWaveTubeAmplifier              respjson.Field
		UserDownlinkFrom                        respjson.Field
		UserDownlinkTo                          respjson.Field
		UserUplinkFrom                          respjson.Field
		UserUplinkTo                            respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataCommDetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataCommDetailListResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataCommDetailListResponseDataMode string

const (
	SeraDataCommDetailListResponseDataModeReal      SeraDataCommDetailListResponseDataMode = "REAL"
	SeraDataCommDetailListResponseDataModeTest      SeraDataCommDetailListResponseDataMode = "TEST"
	SeraDataCommDetailListResponseDataModeSimulated SeraDataCommDetailListResponseDataMode = "SIMULATED"
	SeraDataCommDetailListResponseDataModeExercise  SeraDataCommDetailListResponseDataMode = "EXERCISE"
)

// Seradata-compiled information on communications payloads.
type SeraDataCommDetailGetResponse struct {
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
	DataMode SeraDataCommDetailGetResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// Comm bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Effective isotropic radiated power in dB.
	Eirp float64 `json:"eirp"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity float64 `json:"estHtsTotalCapacity"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam float64 `json:"estHtsTotalUserDownlinkBandwidthPerBeam"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam float64 `json:"estHtsTotalUserUplinkBandwidthPerBeam"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom float64 `json:"gatewayDownlinkFrom"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo float64 `json:"gatewayDownlinkTo"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom float64 `json:"gatewayUplinkFrom"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo float64 `json:"gatewayUplinkTo"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams int64 `json:"htsNumUserSpotBeams"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam float64 `json:"htsUserDownlinkBandwidthPerBeam"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam float64 `json:"htsUserUplinkBandwidthPerBeam"`
	// UUID of the parent Comm record.
	IDComm string `json:"idComm"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders int64 `json:"num36MhzEquivalentTransponders"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders int64 `json:"numOperationalTransponders"`
	// Comm numSpareTransponders.
	NumSpareTransponders int64 `json:"numSpareTransponders"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Comm polarization.
	Polarization string `json:"polarization"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp float64 `json:"solidStatePowerAmp"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID string `json:"tradeLeaseOrgId"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier float64 `json:"travelingWaveTubeAmplifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom float64 `json:"userDownlinkFrom"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo float64 `json:"userDownlinkTo"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom float64 `json:"userUplinkFrom"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo float64 `json:"userUplinkTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		Source                                  respjson.Field
		ID                                      respjson.Field
		Band                                    respjson.Field
		Bandwidth                               respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		Eirp                                    respjson.Field
		EstHtsTotalCapacity                     respjson.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam respjson.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   respjson.Field
		GatewayDownlinkFrom                     respjson.Field
		GatewayDownlinkTo                       respjson.Field
		GatewayUplinkFrom                       respjson.Field
		GatewayUplinkTo                         respjson.Field
		HostedForCompanyOrgID                   respjson.Field
		HtsNumUserSpotBeams                     respjson.Field
		HtsUserDownlinkBandwidthPerBeam         respjson.Field
		HtsUserUplinkBandwidthPerBeam           respjson.Field
		IDComm                                  respjson.Field
		ManufacturerOrgID                       respjson.Field
		Num36MhzEquivalentTransponders          respjson.Field
		NumOperationalTransponders              respjson.Field
		NumSpareTransponders                    respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		PayloadNotes                            respjson.Field
		Polarization                            respjson.Field
		SolidStatePowerAmp                      respjson.Field
		SpacecraftID                            respjson.Field
		TradeLeaseOrgID                         respjson.Field
		TravelingWaveTubeAmplifier              respjson.Field
		UpdatedAt                               respjson.Field
		UpdatedBy                               respjson.Field
		UserDownlinkFrom                        respjson.Field
		UserDownlinkTo                          respjson.Field
		UserUplinkFrom                          respjson.Field
		UserUplinkTo                            respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataCommDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataCommDetailGetResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataCommDetailGetResponseDataMode string

const (
	SeraDataCommDetailGetResponseDataModeReal      SeraDataCommDetailGetResponseDataMode = "REAL"
	SeraDataCommDetailGetResponseDataModeTest      SeraDataCommDetailGetResponseDataMode = "TEST"
	SeraDataCommDetailGetResponseDataModeSimulated SeraDataCommDetailGetResponseDataMode = "SIMULATED"
	SeraDataCommDetailGetResponseDataModeExercise  SeraDataCommDetailGetResponseDataMode = "EXERCISE"
)

// Seradata-compiled information on communications payloads.
type SeraDataCommDetailTupleResponse struct {
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
	DataMode SeraDataCommDetailTupleResponseDataMode `json:"dataMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band string `json:"band"`
	// Comm bandwidth in Mhz.
	Bandwidth float64 `json:"bandwidth"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Effective isotropic radiated power in dB.
	Eirp float64 `json:"eirp"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity float64 `json:"estHtsTotalCapacity"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam float64 `json:"estHtsTotalUserDownlinkBandwidthPerBeam"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam float64 `json:"estHtsTotalUserUplinkBandwidthPerBeam"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom float64 `json:"gatewayDownlinkFrom"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo float64 `json:"gatewayDownlinkTo"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom float64 `json:"gatewayUplinkFrom"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo float64 `json:"gatewayUplinkTo"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID string `json:"hostedForCompanyOrgId"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams int64 `json:"htsNumUserSpotBeams"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam float64 `json:"htsUserDownlinkBandwidthPerBeam"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam float64 `json:"htsUserUplinkBandwidthPerBeam"`
	// UUID of the parent Comm record.
	IDComm string `json:"idComm"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID string `json:"manufacturerOrgId"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders int64 `json:"num36MhzEquivalentTransponders"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders int64 `json:"numOperationalTransponders"`
	// Comm numSpareTransponders.
	NumSpareTransponders int64 `json:"numSpareTransponders"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Payload notes.
	PayloadNotes string `json:"payloadNotes"`
	// Comm polarization.
	Polarization string `json:"polarization"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp float64 `json:"solidStatePowerAmp"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID string `json:"spacecraftId"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID string `json:"tradeLeaseOrgId"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier float64 `json:"travelingWaveTubeAmplifier"`
	// Time the row was last updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom float64 `json:"userDownlinkFrom"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo float64 `json:"userDownlinkTo"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom float64 `json:"userUplinkFrom"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo float64 `json:"userUplinkTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking                   respjson.Field
		DataMode                                respjson.Field
		Source                                  respjson.Field
		ID                                      respjson.Field
		Band                                    respjson.Field
		Bandwidth                               respjson.Field
		CreatedAt                               respjson.Field
		CreatedBy                               respjson.Field
		Eirp                                    respjson.Field
		EstHtsTotalCapacity                     respjson.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam respjson.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   respjson.Field
		GatewayDownlinkFrom                     respjson.Field
		GatewayDownlinkTo                       respjson.Field
		GatewayUplinkFrom                       respjson.Field
		GatewayUplinkTo                         respjson.Field
		HostedForCompanyOrgID                   respjson.Field
		HtsNumUserSpotBeams                     respjson.Field
		HtsUserDownlinkBandwidthPerBeam         respjson.Field
		HtsUserUplinkBandwidthPerBeam           respjson.Field
		IDComm                                  respjson.Field
		ManufacturerOrgID                       respjson.Field
		Num36MhzEquivalentTransponders          respjson.Field
		NumOperationalTransponders              respjson.Field
		NumSpareTransponders                    respjson.Field
		Origin                                  respjson.Field
		OrigNetwork                             respjson.Field
		PayloadNotes                            respjson.Field
		Polarization                            respjson.Field
		SolidStatePowerAmp                      respjson.Field
		SpacecraftID                            respjson.Field
		TradeLeaseOrgID                         respjson.Field
		TravelingWaveTubeAmplifier              respjson.Field
		UpdatedAt                               respjson.Field
		UpdatedBy                               respjson.Field
		UserDownlinkFrom                        respjson.Field
		UserDownlinkTo                          respjson.Field
		UserUplinkFrom                          respjson.Field
		UserUplinkTo                            respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeraDataCommDetailTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeraDataCommDetailTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeraDataCommDetailTupleResponseDataMode string

const (
	SeraDataCommDetailTupleResponseDataModeReal      SeraDataCommDetailTupleResponseDataMode = "REAL"
	SeraDataCommDetailTupleResponseDataModeTest      SeraDataCommDetailTupleResponseDataMode = "TEST"
	SeraDataCommDetailTupleResponseDataModeSimulated SeraDataCommDetailTupleResponseDataMode = "SIMULATED"
	SeraDataCommDetailTupleResponseDataModeExercise  SeraDataCommDetailTupleResponseDataMode = "EXERCISE"
)

type SeraDataCommDetailNewParams struct {
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
	DataMode SeraDataCommDetailNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band param.Opt[string] `json:"band,omitzero"`
	// Comm bandwidth in Mhz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Effective isotropic radiated power in dB.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity param.Opt[float64] `json:"estHtsTotalCapacity,omitzero"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam param.Opt[float64] `json:"estHtsTotalUserDownlinkBandwidthPerBeam,omitzero"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam param.Opt[float64] `json:"estHtsTotalUserUplinkBandwidthPerBeam,omitzero"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom param.Opt[float64] `json:"gatewayDownlinkFrom,omitzero"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo param.Opt[float64] `json:"gatewayDownlinkTo,omitzero"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom param.Opt[float64] `json:"gatewayUplinkFrom,omitzero"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo param.Opt[float64] `json:"gatewayUplinkTo,omitzero"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams param.Opt[int64] `json:"htsNumUserSpotBeams,omitzero"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam param.Opt[float64] `json:"htsUserDownlinkBandwidthPerBeam,omitzero"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam param.Opt[float64] `json:"htsUserUplinkBandwidthPerBeam,omitzero"`
	// UUID of the parent Comm record.
	IDComm param.Opt[string] `json:"idComm,omitzero"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders param.Opt[int64] `json:"num36MhzEquivalentTransponders,omitzero"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders param.Opt[int64] `json:"numOperationalTransponders,omitzero"`
	// Comm numSpareTransponders.
	NumSpareTransponders param.Opt[int64] `json:"numSpareTransponders,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Payload notes.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Comm polarization.
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp param.Opt[float64] `json:"solidStatePowerAmp,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID param.Opt[string] `json:"spacecraftId,omitzero"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID param.Opt[string] `json:"tradeLeaseOrgId,omitzero"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier param.Opt[float64] `json:"travelingWaveTubeAmplifier,omitzero"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom param.Opt[float64] `json:"userDownlinkFrom,omitzero"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo param.Opt[float64] `json:"userDownlinkTo,omitzero"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom param.Opt[float64] `json:"userUplinkFrom,omitzero"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo param.Opt[float64] `json:"userUplinkTo,omitzero"`
	paramObj
}

func (r SeraDataCommDetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataCommDetailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeraDataCommDetailNewParams) UnmarshalJSON(data []byte) error {
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
type SeraDataCommDetailNewParamsDataMode string

const (
	SeraDataCommDetailNewParamsDataModeReal      SeraDataCommDetailNewParamsDataMode = "REAL"
	SeraDataCommDetailNewParamsDataModeTest      SeraDataCommDetailNewParamsDataMode = "TEST"
	SeraDataCommDetailNewParamsDataModeSimulated SeraDataCommDetailNewParamsDataMode = "SIMULATED"
	SeraDataCommDetailNewParamsDataModeExercise  SeraDataCommDetailNewParamsDataMode = "EXERCISE"
)

type SeraDataCommDetailUpdateParams struct {
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
	DataMode SeraDataCommDetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of the band of this RF range (e.g.
	// X,K,Ku,Ka,L,S,C,UHF,VHF,EHF,SHF,UNK,VLF,HF,E,Q,V,W). See RFBandType for more
	// details and descriptions of each band name.
	Band param.Opt[string] `json:"band,omitzero"`
	// Comm bandwidth in Mhz.
	Bandwidth param.Opt[float64] `json:"bandwidth,omitzero"`
	// Effective isotropic radiated power in dB.
	Eirp param.Opt[float64] `json:"eirp,omitzero"`
	// Comm estimated HtsTotalCapacity in Gbps.
	EstHtsTotalCapacity param.Opt[float64] `json:"estHtsTotalCapacity,omitzero"`
	// Comm estimated HtsTotalUserDownlinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserDownlinkBandwidthPerBeam param.Opt[float64] `json:"estHtsTotalUserDownlinkBandwidthPerBeam,omitzero"`
	// Comm estimated HtsTotalUserUplinkBandwidthPerBeam in Mhz.
	EstHtsTotalUserUplinkBandwidthPerBeam param.Opt[float64] `json:"estHtsTotalUserUplinkBandwidthPerBeam,omitzero"`
	// Comm gatewayDownlinkFrom in Ghz.
	GatewayDownlinkFrom param.Opt[float64] `json:"gatewayDownlinkFrom,omitzero"`
	// Comm gatewayDownlinkTo in Ghz.
	GatewayDownlinkTo param.Opt[float64] `json:"gatewayDownlinkTo,omitzero"`
	// Comm gatewayUplinkFrom in Ghz.
	GatewayUplinkFrom param.Opt[float64] `json:"gatewayUplinkFrom,omitzero"`
	// Comm gatewayUplinkTo in Ghz.
	GatewayUplinkTo param.Opt[float64] `json:"gatewayUplinkTo,omitzero"`
	// Comm hostedForCompanyOrgId.
	HostedForCompanyOrgID param.Opt[string] `json:"hostedForCompanyOrgId,omitzero"`
	// Comm htsNumUserSpotBeams.
	HtsNumUserSpotBeams param.Opt[int64] `json:"htsNumUserSpotBeams,omitzero"`
	// Comm htsUserDownlinkBandwidthPerBeam in Mhz.
	HtsUserDownlinkBandwidthPerBeam param.Opt[float64] `json:"htsUserDownlinkBandwidthPerBeam,omitzero"`
	// Comm htsUserUplinkBandwidthPerBeam in Mhz.
	HtsUserUplinkBandwidthPerBeam param.Opt[float64] `json:"htsUserUplinkBandwidthPerBeam,omitzero"`
	// UUID of the parent Comm record.
	IDComm param.Opt[string] `json:"idComm,omitzero"`
	// Comm manufacturerOrgId.
	ManufacturerOrgID param.Opt[string] `json:"manufacturerOrgId,omitzero"`
	// Comm num36MhzEquivalentTransponders.
	Num36MhzEquivalentTransponders param.Opt[int64] `json:"num36MhzEquivalentTransponders,omitzero"`
	// Comm numOperationalTransponders.
	NumOperationalTransponders param.Opt[int64] `json:"numOperationalTransponders,omitzero"`
	// Comm numSpareTransponders.
	NumSpareTransponders param.Opt[int64] `json:"numSpareTransponders,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Payload notes.
	PayloadNotes param.Opt[string] `json:"payloadNotes,omitzero"`
	// Comm polarization.
	Polarization param.Opt[string] `json:"polarization,omitzero"`
	// Solid state power amplifier, in Watts.
	SolidStatePowerAmp param.Opt[float64] `json:"solidStatePowerAmp,omitzero"`
	// Seradata ID of the spacecraft (SeradataSpacecraftDetails ID).
	SpacecraftID param.Opt[string] `json:"spacecraftId,omitzero"`
	// Comm tradeLeaseOrgId.
	TradeLeaseOrgID param.Opt[string] `json:"tradeLeaseOrgId,omitzero"`
	// Comm travelingWaveTubeAmplifier in Watts.
	TravelingWaveTubeAmplifier param.Opt[float64] `json:"travelingWaveTubeAmplifier,omitzero"`
	// Comm userDownlinkFrom in Ghz.
	UserDownlinkFrom param.Opt[float64] `json:"userDownlinkFrom,omitzero"`
	// Comm userDownlinkTo in Ghz.
	UserDownlinkTo param.Opt[float64] `json:"userDownlinkTo,omitzero"`
	// Comm userUplinkFrom in Ghz.
	UserUplinkFrom param.Opt[float64] `json:"userUplinkFrom,omitzero"`
	// Comm userUplinkTo in Ghz.
	UserUplinkTo param.Opt[float64] `json:"userUplinkTo,omitzero"`
	paramObj
}

func (r SeraDataCommDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeraDataCommDetailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeraDataCommDetailUpdateParams) UnmarshalJSON(data []byte) error {
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
type SeraDataCommDetailUpdateParamsDataMode string

const (
	SeraDataCommDetailUpdateParamsDataModeReal      SeraDataCommDetailUpdateParamsDataMode = "REAL"
	SeraDataCommDetailUpdateParamsDataModeTest      SeraDataCommDetailUpdateParamsDataMode = "TEST"
	SeraDataCommDetailUpdateParamsDataModeSimulated SeraDataCommDetailUpdateParamsDataMode = "SIMULATED"
	SeraDataCommDetailUpdateParamsDataModeExercise  SeraDataCommDetailUpdateParamsDataMode = "EXERCISE"
)

type SeraDataCommDetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataCommDetailListParams]'s query parameters as
// `url.Values`.
func (r SeraDataCommDetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataCommDetailCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataCommDetailCountParams]'s query parameters as
// `url.Values`.
func (r SeraDataCommDetailCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataCommDetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataCommDetailGetParams]'s query parameters as
// `url.Values`.
func (r SeraDataCommDetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeraDataCommDetailTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SeraDataCommDetailTupleParams]'s query parameters as
// `url.Values`.
func (r SeraDataCommDetailTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
