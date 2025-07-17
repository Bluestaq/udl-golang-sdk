// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// SupportingDataService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataService] method instead.
type SupportingDataService struct {
	Options          []option.RequestOption
	DataTypes        SupportingDataDataTypeService
	Dataowner        SupportingDataDataownerService
	DataownerTypes   SupportingDataDataownerTypeService
	ProviderMetadata SupportingDataProviderMetadataService
	QueryHelp        SupportingDataQueryHelpService
}

// NewSupportingDataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSupportingDataService(opts ...option.RequestOption) (r SupportingDataService) {
	r = SupportingDataService{}
	r.Options = opts
	r.DataTypes = NewSupportingDataDataTypeService(opts...)
	r.Dataowner = NewSupportingDataDataownerService(opts...)
	r.DataownerTypes = NewSupportingDataDataownerTypeService(opts...)
	r.ProviderMetadata = NewSupportingDataProviderMetadataService(opts...)
	r.QueryHelp = NewSupportingDataQueryHelpService(opts...)
	return
}
