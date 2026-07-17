// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
)

// AirOperationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAirOperationService] method instead.
type AirOperationService struct {
	Options []option.RequestOption
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
	AirTaskingOrders AirOperationAirTaskingOrderService
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
	AircraftSorties AirOperationAircraftSortyService
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
	AirspaceControlOrders AirOperationAirspaceControlOrderService
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
	Crewpapers AirOperationCrewpaperService
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
	DiplomaticClearance AirOperationDiplomaticClearanceService
}

// NewAirOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAirOperationService(opts ...option.RequestOption) (r AirOperationService) {
	r = AirOperationService{}
	r.Options = opts
	r.AirTaskingOrders = NewAirOperationAirTaskingOrderService(opts...)
	r.AircraftSorties = NewAirOperationAircraftSortyService(opts...)
	r.AirspaceControlOrders = NewAirOperationAirspaceControlOrderService(opts...)
	r.Crewpapers = NewAirOperationCrewpaperService(opts...)
	r.DiplomaticClearance = NewAirOperationDiplomaticClearanceService(opts...)
	return
}
