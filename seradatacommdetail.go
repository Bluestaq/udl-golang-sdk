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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
)

// SeradatacommdetailService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeradatacommdetailService] method instead.
type SeradatacommdetailService struct {
	Options []option.RequestOption
}

// NewSeradatacommdetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSeradatacommdetailService(opts ...option.RequestOption) (r SeradatacommdetailService) {
	r = SeradatacommdetailService{}
	r.Options = opts
	return
}

// Service operation to take a single SeradataCommDetails as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *SeradatacommdetailService) New(ctx context.Context, body SeradatacommdetailNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/seradatacommdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update an SeradataCommDetails. A specific role is required
// to perform this service operation. Please contact the UDL team for assistance.
func (r *SeradatacommdetailService) Update(ctx context.Context, id string, body SeradatacommdetailUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *SeradatacommdetailService) List(ctx context.Context, query SeradatacommdetailListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SeradatacommdetailListResponse], err error) {
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
func (r *SeradatacommdetailService) ListAutoPaging(ctx context.Context, query SeradatacommdetailListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SeradatacommdetailListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete an SeradataCommDetails specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *SeradatacommdetailService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *SeradatacommdetailService) Count(ctx context.Context, query SeradatacommdetailCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/seradatacommdetails/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single SeradataCommDetails by its unique ID passed as
// a path parameter.
func (r *SeradatacommdetailService) Get(ctx context.Context, id string, query SeradatacommdetailGetParams, opts ...option.RequestOption) (res *SeradatacommdetailGetResponse, err error) {
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
func (r *SeradatacommdetailService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
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
func (r *SeradatacommdetailService) Tuple(ctx context.Context, query SeradatacommdetailTupleParams, opts ...option.RequestOption) (res *[]SeradatacommdetailTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/seradatacommdetails/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Seradata-compiled information on communications payloads.
type SeradatacommdetailListResponse struct {
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
	DataMode SeradatacommdetailListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                   resp.Field
		DataMode                                resp.Field
		Source                                  resp.Field
		ID                                      resp.Field
		Band                                    resp.Field
		Bandwidth                               resp.Field
		CreatedAt                               resp.Field
		CreatedBy                               resp.Field
		Eirp                                    resp.Field
		EstHtsTotalCapacity                     resp.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam resp.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   resp.Field
		GatewayDownlinkFrom                     resp.Field
		GatewayDownlinkTo                       resp.Field
		GatewayUplinkFrom                       resp.Field
		GatewayUplinkTo                         resp.Field
		HostedForCompanyOrgID                   resp.Field
		HtsNumUserSpotBeams                     resp.Field
		HtsUserDownlinkBandwidthPerBeam         resp.Field
		HtsUserUplinkBandwidthPerBeam           resp.Field
		IDComm                                  resp.Field
		ManufacturerOrgID                       resp.Field
		Num36MhzEquivalentTransponders          resp.Field
		NumOperationalTransponders              resp.Field
		NumSpareTransponders                    resp.Field
		Origin                                  resp.Field
		OrigNetwork                             resp.Field
		PayloadNotes                            resp.Field
		Polarization                            resp.Field
		SolidStatePowerAmp                      resp.Field
		SpacecraftID                            resp.Field
		TradeLeaseOrgID                         resp.Field
		TravelingWaveTubeAmplifier              resp.Field
		UserDownlinkFrom                        resp.Field
		UserDownlinkTo                          resp.Field
		UserUplinkFrom                          resp.Field
		UserUplinkTo                            resp.Field
		ExtraFields                             map[string]resp.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradatacommdetailListResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradatacommdetailListResponse) UnmarshalJSON(data []byte) error {
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
type SeradatacommdetailListResponseDataMode string

const (
	SeradatacommdetailListResponseDataModeReal      SeradatacommdetailListResponseDataMode = "REAL"
	SeradatacommdetailListResponseDataModeTest      SeradatacommdetailListResponseDataMode = "TEST"
	SeradatacommdetailListResponseDataModeSimulated SeradatacommdetailListResponseDataMode = "SIMULATED"
	SeradatacommdetailListResponseDataModeExercise  SeradatacommdetailListResponseDataMode = "EXERCISE"
)

// Seradata-compiled information on communications payloads.
type SeradatacommdetailGetResponse struct {
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
	DataMode SeradatacommdetailGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                   resp.Field
		DataMode                                resp.Field
		Source                                  resp.Field
		ID                                      resp.Field
		Band                                    resp.Field
		Bandwidth                               resp.Field
		CreatedAt                               resp.Field
		CreatedBy                               resp.Field
		Eirp                                    resp.Field
		EstHtsTotalCapacity                     resp.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam resp.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   resp.Field
		GatewayDownlinkFrom                     resp.Field
		GatewayDownlinkTo                       resp.Field
		GatewayUplinkFrom                       resp.Field
		GatewayUplinkTo                         resp.Field
		HostedForCompanyOrgID                   resp.Field
		HtsNumUserSpotBeams                     resp.Field
		HtsUserDownlinkBandwidthPerBeam         resp.Field
		HtsUserUplinkBandwidthPerBeam           resp.Field
		IDComm                                  resp.Field
		ManufacturerOrgID                       resp.Field
		Num36MhzEquivalentTransponders          resp.Field
		NumOperationalTransponders              resp.Field
		NumSpareTransponders                    resp.Field
		Origin                                  resp.Field
		OrigNetwork                             resp.Field
		PayloadNotes                            resp.Field
		Polarization                            resp.Field
		SolidStatePowerAmp                      resp.Field
		SpacecraftID                            resp.Field
		TradeLeaseOrgID                         resp.Field
		TravelingWaveTubeAmplifier              resp.Field
		UpdatedAt                               resp.Field
		UpdatedBy                               resp.Field
		UserDownlinkFrom                        resp.Field
		UserDownlinkTo                          resp.Field
		UserUplinkFrom                          resp.Field
		UserUplinkTo                            resp.Field
		ExtraFields                             map[string]resp.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradatacommdetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradatacommdetailGetResponse) UnmarshalJSON(data []byte) error {
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
type SeradatacommdetailGetResponseDataMode string

const (
	SeradatacommdetailGetResponseDataModeReal      SeradatacommdetailGetResponseDataMode = "REAL"
	SeradatacommdetailGetResponseDataModeTest      SeradatacommdetailGetResponseDataMode = "TEST"
	SeradatacommdetailGetResponseDataModeSimulated SeradatacommdetailGetResponseDataMode = "SIMULATED"
	SeradatacommdetailGetResponseDataModeExercise  SeradatacommdetailGetResponseDataMode = "EXERCISE"
)

// Seradata-compiled information on communications payloads.
type SeradatacommdetailTupleResponse struct {
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
	DataMode SeradatacommdetailTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking                   resp.Field
		DataMode                                resp.Field
		Source                                  resp.Field
		ID                                      resp.Field
		Band                                    resp.Field
		Bandwidth                               resp.Field
		CreatedAt                               resp.Field
		CreatedBy                               resp.Field
		Eirp                                    resp.Field
		EstHtsTotalCapacity                     resp.Field
		EstHtsTotalUserDownlinkBandwidthPerBeam resp.Field
		EstHtsTotalUserUplinkBandwidthPerBeam   resp.Field
		GatewayDownlinkFrom                     resp.Field
		GatewayDownlinkTo                       resp.Field
		GatewayUplinkFrom                       resp.Field
		GatewayUplinkTo                         resp.Field
		HostedForCompanyOrgID                   resp.Field
		HtsNumUserSpotBeams                     resp.Field
		HtsUserDownlinkBandwidthPerBeam         resp.Field
		HtsUserUplinkBandwidthPerBeam           resp.Field
		IDComm                                  resp.Field
		ManufacturerOrgID                       resp.Field
		Num36MhzEquivalentTransponders          resp.Field
		NumOperationalTransponders              resp.Field
		NumSpareTransponders                    resp.Field
		Origin                                  resp.Field
		OrigNetwork                             resp.Field
		PayloadNotes                            resp.Field
		Polarization                            resp.Field
		SolidStatePowerAmp                      resp.Field
		SpacecraftID                            resp.Field
		TradeLeaseOrgID                         resp.Field
		TravelingWaveTubeAmplifier              resp.Field
		UpdatedAt                               resp.Field
		UpdatedBy                               resp.Field
		UserDownlinkFrom                        resp.Field
		UserDownlinkTo                          resp.Field
		UserUplinkFrom                          resp.Field
		UserUplinkTo                            resp.Field
		ExtraFields                             map[string]resp.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeradatacommdetailTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SeradatacommdetailTupleResponse) UnmarshalJSON(data []byte) error {
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
type SeradatacommdetailTupleResponseDataMode string

const (
	SeradatacommdetailTupleResponseDataModeReal      SeradatacommdetailTupleResponseDataMode = "REAL"
	SeradatacommdetailTupleResponseDataModeTest      SeradatacommdetailTupleResponseDataMode = "TEST"
	SeradatacommdetailTupleResponseDataModeSimulated SeradatacommdetailTupleResponseDataMode = "SIMULATED"
	SeradatacommdetailTupleResponseDataModeExercise  SeradatacommdetailTupleResponseDataMode = "EXERCISE"
)

type SeradatacommdetailNewParams struct {
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
	DataMode SeradatacommdetailNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeradatacommdetailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradatacommdetailNewParams
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
type SeradatacommdetailNewParamsDataMode string

const (
	SeradatacommdetailNewParamsDataModeReal      SeradatacommdetailNewParamsDataMode = "REAL"
	SeradatacommdetailNewParamsDataModeTest      SeradatacommdetailNewParamsDataMode = "TEST"
	SeradatacommdetailNewParamsDataModeSimulated SeradatacommdetailNewParamsDataMode = "SIMULATED"
	SeradatacommdetailNewParamsDataModeExercise  SeradatacommdetailNewParamsDataMode = "EXERCISE"
)

type SeradatacommdetailUpdateParams struct {
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
	DataMode SeradatacommdetailUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SeradatacommdetailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SeradatacommdetailUpdateParams
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
type SeradatacommdetailUpdateParamsDataMode string

const (
	SeradatacommdetailUpdateParamsDataModeReal      SeradatacommdetailUpdateParamsDataMode = "REAL"
	SeradatacommdetailUpdateParamsDataModeTest      SeradatacommdetailUpdateParamsDataMode = "TEST"
	SeradatacommdetailUpdateParamsDataModeSimulated SeradatacommdetailUpdateParamsDataMode = "SIMULATED"
	SeradatacommdetailUpdateParamsDataModeExercise  SeradatacommdetailUpdateParamsDataMode = "EXERCISE"
)

type SeradatacommdetailListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradatacommdetailListParams]'s query parameters as
// `url.Values`.
func (r SeradatacommdetailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradatacommdetailCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradatacommdetailCountParams]'s query parameters as
// `url.Values`.
func (r SeradatacommdetailCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradatacommdetailGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradatacommdetailGetParams]'s query parameters as
// `url.Values`.
func (r SeradatacommdetailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SeradatacommdetailTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SeradatacommdetailTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SeradatacommdetailTupleParams]'s query parameters as
// `url.Values`.
func (r SeradatacommdetailTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
