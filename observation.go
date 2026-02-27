// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// ObservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservationService] method instead.
type ObservationService struct {
	Options []option.RequestOption
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Ecpsdr ObservationEcpsdrService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	EoObservations ObservationEoObservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Monoradar ObservationMonoradarService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Obscorrelation ObservationObscorrelationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	PassiveRadarObservation ObservationPassiveRadarObservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Radarobservation ObservationRadarobservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	RfObservation ObservationRfObservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Swir ObservationSwirService
}

// NewObservationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObservationService(opts ...option.RequestOption) (r ObservationService) {
	r = ObservationService{}
	r.Options = opts
	r.Ecpsdr = NewObservationEcpsdrService(opts...)
	r.EoObservations = NewObservationEoObservationService(opts...)
	r.Monoradar = NewObservationMonoradarService(opts...)
	r.Obscorrelation = NewObservationObscorrelationService(opts...)
	r.PassiveRadarObservation = NewObservationPassiveRadarObservationService(opts...)
	r.Radarobservation = NewObservationRadarobservationService(opts...)
	r.RfObservation = NewObservationRfObservationService(opts...)
	r.Swir = NewObservationSwirService(opts...)
	return
}
