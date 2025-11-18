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

	"github.com/Bluestaq/udl-golang-sdk/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

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
	return
}

// Service operation to upload a supporting PDF for the aircraft sortie. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *AirOperationCrewpaperService) UploadPdf(ctx context.Context, fileContent io.Reader, body AirOperationCrewpaperUploadPdfParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithRequestBody("application/pdf", fileContent)}, opts...)
	path := "filedrop/crewpapers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type AirOperationCrewpaperUnpublishParams struct {
	// Comma-separated list of AircraftSortie IDs where Crew Papers are unpublished.
	IDs string `query:"ids,required" json:"-"`
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
	AircraftSortieIDs string `query:"aircraftSortieIds,required" json:"-"`
	// classificationMarking of the Crew Papers.
	ClassificationMarking string `query:"classificationMarking,required" json:"-"`
	// The status of the supporting document.
	//
	// Any of "PUBLISHED", "DELETED", "UPDATED", "READ".
	PaperStatus AirOperationCrewpaperUploadPdfParamsPaperStatus `query:"paperStatus,omitzero,required" json:"-"`
	// The version number of the crew paper.
	PapersVersion string `query:"papersVersion,required" json:"-"`
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
