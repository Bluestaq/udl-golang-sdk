// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/v2/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/v2/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/v2/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
)

// These services provide operations for manipulating and querying Aircraft Sortie,
// Aircraft Mission, Item Tracking, Flight Plan, Air Event, Sortie Prior Permission
// Required (PPR), Diplomatic Clearance, Diplomatic Clearance Country, Airspace
// Control Order, Air Tasking Order, Navigational Obstruction, Logistics Support,
// Track Route, Air Load Plan, and Aviation Risk Management data. Aircraft Sortie
// information contains static and dynamic aircraft assignments, departure and
// arrival times, and remarks. Aircraft Mission information contains static data
// for mission planning to include assigned aircraft and crews, cargo pickup and
// dropoff locations, unique identifiers, and prioritization. Item Tracking
// information contains data for tracking an item from its origin to destination
// and how it may be configured during transport. Flight Plan information contains
// schedule and route details. Air Event provides information concerning various
// aerial events such as fuel transfer and air drops, as well as the associated
// aircraft involved. Sortie PPR information contains details on operational access
// to a runway, taxiway, or airport service. Diplomatic Clearance information
// contains details on the issuance and coordination of aircraft clearance
// requests. Diplomatic Clearance Country provides information such as entry/exit
// points, requirements, and points of contact for countries diplomatic clearances
// are being created for. Airspace Control Order provides information concerning
// the allocation, restriction, and deconfliction of airspace. Air Tasking Order
// information contains details on the coordination of air missions and their
// tasks, resources, and timelines. Navigational Obstruction provides the
// locations, characteristics, and boundaries of obstacles and structures that can
// restrict or interfere with navigation. Logistics Support contains information
// regarding the transport and maintenance of resources and equipment to sustain
// air operations. Track Route information defines specific flight paths used by
// aircraft during the transport of fuel and other resources. Air Load Plan
// information provides mission actuals concerning the loading and air transport of
// cargo and passengers. Aviation Risk Management information help aid in mission
// planning by accounting for factors such as mission complexity and crew fatigue.
//
// AirOperationCrewpaperService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirOperationCrewpaperService] method instead.
type AirOperationCrewpaperService struct {
	Options []option.RequestOption
}

// NewAirOperationCrewpaperService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAirOperationCrewpaperService(opts ...option.RequestOption) (r AirOperationCrewpaperService) {
	r = AirOperationCrewpaperService{}
	r.Options = opts
	return
}

// Service operation to remove supporting PDF from an aircraft sortie or sorties. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *AirOperationCrewpaperService) Unpublish(ctx context.Context, body AirOperationCrewpaperUnpublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/crewpapers/unpublish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Service operation to upload a supporting PDF for the aircraft sortie. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *AirOperationCrewpaperService) UploadPdf(ctx context.Context, fileContent io.Reader, params AirOperationCrewpaperUploadPdfParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithRequestBody("application/pdf", fileContent)}, opts...)
	path := "filedrop/crewpapers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type AirOperationCrewpaperUnpublishParams struct {
	// Comma-separated list of AircraftSortie IDs where Crew Papers are unpublished.
	IDs string `query:"ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AirOperationCrewpaperUnpublishParams]'s query parameters as
// `url.Values`.
func (r AirOperationCrewpaperUnpublishParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AirOperationCrewpaperUploadPdfParams struct {
	// Comma-separated list of AircraftSortie IDs the Crew Papers are being added to.
	AircraftSortieIDs string `query:"aircraftSortieIds" api:"required" json:"-"`
	// classificationMarking of the Crew Papers.
	ClassificationMarking string `query:"classificationMarking" api:"required" json:"-"`
	// The status of the supporting document.
	//
	// Any of "PUBLISHED", "DELETED", "UPDATED", "READ".
	PaperStatus AirOperationCrewpaperUploadPdfParamsPaperStatus `query:"paperStatus,omitzero" api:"required" json:"-"`
	// The version number of the crew paper.
	PapersVersion string `query:"papersVersion" api:"required" json:"-"`
	paramObj
}

func (r AirOperationCrewpaperUploadPdfParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [AirOperationCrewpaperUploadPdfParams]'s query parameters as
// `url.Values`.
func (r AirOperationCrewpaperUploadPdfParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The status of the supporting document.
type AirOperationCrewpaperUploadPdfParamsPaperStatus string

const (
	AirOperationCrewpaperUploadPdfParamsPaperStatusPublished AirOperationCrewpaperUploadPdfParamsPaperStatus = "PUBLISHED"
	AirOperationCrewpaperUploadPdfParamsPaperStatusDeleted   AirOperationCrewpaperUploadPdfParamsPaperStatus = "DELETED"
	AirOperationCrewpaperUploadPdfParamsPaperStatusUpdated   AirOperationCrewpaperUploadPdfParamsPaperStatus = "UPDATED"
	AirOperationCrewpaperUploadPdfParamsPaperStatusRead      AirOperationCrewpaperUploadPdfParamsPaperStatus = "READ"
)
