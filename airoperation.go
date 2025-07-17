// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// AirOperationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirOperationService] method instead.
type AirOperationService struct {
	Options               []option.RequestOption
	AirTaskingOrders      AirOperationAirTaskingOrderService
	AircraftSortie        AirOperationAircraftSortieService
	AircraftSorties       AirOperationAircraftSortyService
	AirspaceControlOrders AirOperationAirspaceControlOrderService
	Crewpapers            AirOperationCrewpaperService
	DiplomaticClearance   AirOperationDiplomaticClearanceService
}

// NewAirOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirOperationService(opts ...option.RequestOption) (r AirOperationService) {
	r = AirOperationService{}
	r.Options = opts
	r.AirTaskingOrders = NewAirOperationAirTaskingOrderService(opts...)
	r.AircraftSortie = NewAirOperationAircraftSortieService(opts...)
	r.AircraftSorties = NewAirOperationAircraftSortyService(opts...)
	r.AirspaceControlOrders = NewAirOperationAirspaceControlOrderService(opts...)
	r.Crewpapers = NewAirOperationCrewpaperService(opts...)
	r.DiplomaticClearance = NewAirOperationDiplomaticClearanceService(opts...)
	return
}
