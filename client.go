// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the unifieddatalibrary API. You should not instantiate this
// client directly, and instead use the [NewClient] method instead.
type Client struct {
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
	AirEvents     AirEventService
	AirOperations AirOperationService
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
	AirTransportMissions AirTransportMissionService
	// This service provides operations for manipulation and querying of Aircraft and
	// Aircraft Status data. Aircraft contains the static data of the specific
	// aircraft: tail number, cruise speed, max speed, minimum required runway length,
	// etc. The Aircraft Status contains the dynamic data associated with the specific
	// aircraft: remaining fuel, mission readiness, and inventory for example.
	Aircraft AircraftService
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
	AircraftSorties AircraftSortyService
	// This service provides operations for manipulation and querying of Aircraft and
	// Aircraft Status data. Aircraft contains the static data of the specific
	// aircraft: tail number, cruise speed, max speed, minimum required runway length,
	// etc. The Aircraft Status contains the dynamic data associated with the specific
	// aircraft: remaining fuel, mission readiness, and inventory for example.
	AircraftStatusRemarks AircraftStatusRemarkService
	// This service provides operations for manipulation and querying of Aircraft and
	// Aircraft Status data. Aircraft contains the static data of the specific
	// aircraft: tail number, cruise speed, max speed, minimum required runway length,
	// etc. The Aircraft Status contains the dynamic data associated with the specific
	// aircraft: remaining fuel, mission readiness, and inventory for example.
	AircraftStatuses AircraftStatusService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	AirfieldSlotConsumptions AirfieldSlotConsumptionService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	AirfieldSlots AirfieldSlotService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	AirfieldStatus AirfieldStatusService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	Airfields AirfieldService
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
	AirloadPlans AirloadPlanService
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
	AirspaceControlOrders AirspaceControlOrderService
	// These services provide for posting and querying of self-reported information
	// obtained from the Automatic Identification System (AIS) equipment. This contains
	// information such as unique identification, status, position, course, and speed.
	// The AIS is an automatic tracking system that uses transceivers on ships and is
	// used by vessel traffic services. Although technically and operationally
	// distinct, the AIS system is analogous to ADS-B which performs a similar function
	// for aircraft. AIS is intended to assist a vessel's watchstanding officers and
	// allow maritime authorities to track and monitor vessel movements. AIS integrates
	// a standardized VHF transceiver with a positioning system, such as Global
	// Positioning System receiver, with other electronic navigation sensors, such as
	// gyrocompass or rate of turn indicator. Vessels fitted with AIS transceivers can
	// be tracked by AIS base stations located along coastlines or, when out of range
	// of terrestrial networks, through a growing number of satellites that are fitted
	// with special AIS receivers that are capable of deconflicting a large number of
	// signatures.
	AIs AIService
	// These services provide for posting and querying of self-reported information
	// obtained from the Automatic Identification System (AIS) equipment. This contains
	// information such as unique identification, status, position, course, and speed.
	// The AIS is an automatic tracking system that uses transceivers on ships and is
	// used by vessel traffic services. Although technically and operationally
	// distinct, the AIS system is analogous to ADS-B which performs a similar function
	// for aircraft. AIS is intended to assist a vessel's watchstanding officers and
	// allow maritime authorities to track and monitor vessel movements. AIS integrates
	// a standardized VHF transceiver with a positioning system, such as Global
	// Positioning System receiver, with other electronic navigation sensors, such as
	// gyrocompass or rate of turn indicator. Vessels fitted with AIS transceivers can
	// be tracked by AIS base stations located along coastlines or, when out of range
	// of terrestrial networks, through a growing number of satellites that are fitted
	// with special AIS receivers that are capable of deconflicting a large number of
	// signatures.
	AIsObjects AIsObjectService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	AnalyticImagery AnalyticImageryService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Antennas AntennaService
	// These services provide operations for the posting and querying of satellite
	// Ephemeris Point data. Each point contains a position and velocity vector and
	// optionally, an acceleration vector and/or covariance matrix at a specified time.
	// ECI J2K is the preferred reference frame for ephemeris and covariance, however,
	// several user specified reference frames are accommodated. The EphemerisSet ID
	// (esId) identifies the 'EphemerisSet' record which contains details of the
	// underlying data and models used in the generation of the ephemeris as well as a
	// collection of ephemeris points. Points must be retrieved by first identifying a
	// desired EphemerisSet and pulling its points by that EphemerisSet 'esId'.
	AttitudeData AttitudeDataService
	// These services provide operations for the posting and querying of satellite
	// Ephemeris Point data. Each point contains a position and velocity vector and
	// optionally, an acceleration vector and/or covariance matrix at a specified time.
	// ECI J2K is the preferred reference frame for ephemeris and covariance, however,
	// several user specified reference frames are accommodated. The EphemerisSet ID
	// (esId) identifies the 'EphemerisSet' record which contains details of the
	// underlying data and models used in the generation of the ephemeris as well as a
	// collection of ephemeris points. Points must be retrieved by first identifying a
	// desired EphemerisSet and pulling its points by that EphemerisSet 'esId'.
	AttitudeSets AttitudeSetService
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
	AviationRiskManagement AviationRiskManagementService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Batteries BatteryService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Batterydetails BatterydetailService
	// This collection of services provides operations for querying and manipulation of
	// satellite antenna beams, and querying of beam contours and service areas. Beam
	// contours are the geographic representation of the relative gain levels of beam
	// power off of the maximum gain boresight points. Similarly, service areas are the
	// geographic footprints of the areas served by a particular beam, and may be made
	// up of multiple service regions. Well-Known Text (WKT) and GeoJSON formats are
	// used for GIS representation and query support (see
	// https://www.opengeospatial.org/standards/wkt-crs and https://geojson.org/ for
	// more information on these formats).
	Beam BeamService
	// This collection of services provides operations for querying and manipulation of
	// satellite antenna beams, and querying of beam contours and service areas. Beam
	// contours are the geographic representation of the relative gain levels of beam
	// power off of the maximum gain boresight points. Similarly, service areas are the
	// geographic footprints of the areas served by a particular beam, and may be made
	// up of multiple service regions. Well-Known Text (WKT) and GeoJSON formats are
	// used for GIS representation and query support (see
	// https://www.opengeospatial.org/standards/wkt-crs and https://geojson.org/ for
	// more information on these formats).
	BeamContours BeamContourService
	// Services for querying and manipulation of satellite buses. A bus is the physical
	// and software infrastructure backbone to which on-orbit satellite payloads are
	// attached for power, control, and other support functions.
	Buses BusService
	// These services provide operations for manipulation and querying of on-orbit
	// communications payloads (Comm), including supporting data such as transponders
	// and channels, etc.
	Channels ChannelService
	// This collection of services provides operations for manipulating and querying of
	// closely spaced objects (on orbit) operations including docking, rendezvous,
	// proximity and reporting of payload zone engagements observed and characterized
	// over a period of time.
	Closelyspacedobjects CloselyspacedobjectService
	// These services provide operations for posting and querying Sensor Tasking data.
	CollectRequests CollectRequestService
	// These services provide operations for posting and querying Sensor Tasking data.
	CollectResponses CollectResponseService
	// These services provide operations for manipulation and querying of on-orbit
	// communications payloads (Comm), including supporting data such as transponders
	// and channels, etc.
	Comm CommService
	// These services provide operations for manipulation and querying of conjunctions.
	Conjunctions ConjunctionService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Cots CotService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Countries CountryService
	// These services provide operations for posting and querying crew data. Crew data
	// contains information about its members and their assignments.
	Crew CrewService
	// These services provide operations for manipulation and querying of Mission Ops
	// information.
	Deconflictset DeconflictsetService
	// These services provide operations for querying and manipulation of Signal time
	// and frequency difference of arrival (TDOA/FDOA) information obtained by using
	// passive RF based sensor phenomenologies and sensor triangulation. The J2000
	// coordinate frame is the preferred frame for all observations, but in some cases
	// observations may be in another frame depending on the provider.
	DiffOfArrival DiffOfArrivalService
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
	DiplomaticClearance DiplomaticClearanceService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	DriftHistory DriftHistoryService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	Dropzone DropzoneService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Ecpedr EcpedrService
	// These services provide operations for manipulation and querying of Mission Ops
	// information.
	EffectRequests EffectRequestService
	// These services provide operations for manipulation and querying of Mission Ops
	// information.
	EffectResponses EffectResponseService
	// These services provide operations for querying and manipulation of element set
	// data describing orbital characteristics of on-orbit objects. An element set is a
	// collection of parameters that are used, along with an orbit propagator, to
	// predict the motion of a satellite. The element set, or elset for short, consists
	// of identification data, the classical elements and drag parameters.
	Elsets ElsetService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Emireport EmireportService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	EmitterGeolocation EmitterGeolocationService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	EngineDetails EngineDetailService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	Engines EngineService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Entities EntityService
	// This service provides operations for manipulation and querying of earth
	// orientation parameter (EOP) data. Earth Orientation Parameters (EOP) are
	// produced by the IERS (International Earth Rotation and Reference Systems
	// Service). Earth Orientation Parameters describe the irregularities of the
	// earth's rotation. Technically, they are the parameters which provide the
	// rotation of the ITRS (International Terrestrial Reference System) to the ICRS
	// (International Celestial Reference System) as a function of time. Universal time
	// -- Universal time (UT1) is the time of the earth clock, which performs one
	// revolution in about 24h. It is practically proportional to the sidereal time.
	// The excess revolution time is called length of day (LOD). Coordinates of the
	// pole -- x and y are the coordinates of the Celestial Ephemeris Pole (CEP)
	// relative to the IRP, the IERS Reference Pole. The CEP differs from the
	// instantaneous rotation axis by quasi-diurnal terms with amplitudes under 0.01"
	// (see Seidelmann, 1982). The x-axis is in the direction of the ITRF
	// zero-meridian; the y-axis is in the direction 90 degrees West longitude.
	// Celestial pole offsets -- Celestial pole offsets are described in the IAU
	// Precession and Nutation models. The observed differences with respect to the
	// conventional celestial pole position defined by the models are monitored and
	// reported by the IERS. IERS Bulletins A and B provide current information on the
	// Earth's orientation in the IERS Reference System. This includes Universal Time,
	// coordinates of the terrestrial pole, and celestial pole offsets. Bulletin A
	// gives an advanced solution updated weekly; the standard solution is given
	// monthly in Bulletin B. Fields suffixed with ”B” are Bulletin B values. All
	// solutions are continuous within their respective uncertainties. Bulletin A is
	// issued by the IERS Rapid Service/Prediction Centre at the U.S. Naval
	// Observatory, Washington, DC and Bulletin B is issued by the IERS Earth
	// Orientation Centre at the Paris Observatory. IERS Bulletin A reports the latest
	// determinations for polar motion, UT1-UTC, and nutation offsets at daily
	// intervals based on a combination of contributed analysis results using data from
	// Very Long Baseline Interferometry (VLBI), Satellite Laser Ranging (SLR), Global
	// Positioning System (GPS) satellites, and Lunar Laser Ranging (LLR). Predictions
	// for variations a year into the future are also provided. Meteorological
	// predictions of variations in Atmospheric Angular Momentum (AAM) are used to aid
	// in the prediction of near-term UT1-UTC changes. This publication is prepared by
	// the IERS Rapid Service/Prediction Center.
	Eop EopService
	// These services provide operations for the posting and querying of satellite
	// Ephemeris Point data. Each point contains a position and velocity vector and
	// optionally, an acceleration vector and/or covariance matrix at a specified time.
	// ECI J2K is the preferred reference frame for ephemeris and covariance, however,
	// several user specified reference frames are accommodated. The EphemerisSet ID
	// (esId) identifies the 'EphemerisSet' record which contains details of the
	// underlying data and models used in the generation of the ephemeris as well as a
	// collection of ephemeris points. Points must be retrieved by first identifying a
	// desired EphemerisSet and pulling its points by that EphemerisSet 'esId'.
	Ephemeris EphemerisService
	// These services provide operations for the posting and querying of satellite
	// Ephemeris Point data. Each point contains a position and velocity vector and
	// optionally, an acceleration vector and/or covariance matrix at a specified time.
	// ECI J2K is the preferred reference frame for ephemeris and covariance, however,
	// several user specified reference frames are accommodated. The EphemerisSet ID
	// (esId) identifies the 'EphemerisSet' record which contains details of the
	// underlying data and models used in the generation of the ephemeris as well as a
	// collection of ephemeris points. Points must be retrieved by first identifying a
	// desired EphemerisSet and pulling its points by that EphemerisSet 'esId'.
	EphemerisSets EphemerisSetService
	// This collection of services provide operations for manipulating and querying of
	// equipment related data.
	Equipment EquipmentService
	// This collection of services provide operations for manipulating and querying of
	// equipment related data.
	EquipmentRemarks EquipmentRemarkService
	// These services provide operations for manipulation and querying of Mission Ops
	// information.
	Evac EvacService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	EventEvolution EventEvolutionService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	FeatureAssessment FeatureAssessmentService
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
	Flightplan FlightplanService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	GeoStatus GeoStatusService
	// Models and Simulations is a collection of services that allow consumers to
	// interact with data products representing independent models of various
	// phenomenon, artificial intelligence models and predictions, or of mathematical
	// parameters meant to feed mod and sim tools to produce estimates of environmental
	// entities such as atmospheric models and heat maps.
	GlobalAtmosphericModel GlobalAtmosphericModelService
	GnssObservations       GnssObservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	GnssObservationset GnssObservationsetService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	GnssRawIf GnssRawIfService
	// This collection of services provides operations for querying and manipulation of
	// ground imagery of terrestrial regions from on-orbit, air, and other sensors.
	// Includes metadata on the image (time, region, source, etc) as well as binary
	// content (typically GeoTIFF). Binary content must be downloaded individually by
	// ID using the 'getFile' operation. Well-Known Text (WKT) and GeoJSON formats are
	// used for GIS representation and query support (see
	// https://www.opengeospatial.org/standards/wkt-crs and https://geojson.org/ for
	// more information on these formats).
	GroundImagery GroundImageryService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	H3Geo H3GeoService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	H3GeoHexCell H3GeoHexCellService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Hazard HazardService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	IonoObservations IonoObservationService
	// These services provide operations for manipulation and querying of on-orbit
	// payloads.
	Ir IrService
	// These services provide operations for posting and querying Sensor Tasking data.
	IsrCollections IsrCollectionService
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
	Item ItemService
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
	ItemTrackings ItemTrackingService
	// This collection of services provides operations for querying and manipulation of
	// laser related information to include the laser emitters, the laser deconflict
	// requests, and laser deconflict responses.
	Laserdeconflictrequest LaserdeconflictrequestService
	// This collection of services provides operations for querying and manipulation of
	// laser related information to include the laser emitters, the laser deconflict
	// requests, and laser deconflict responses.
	Laseremitter LaseremitterService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	LaunchDetection LaunchDetectionService
	// These services provide operations for manipulation and querying of LaunchEvent
	// data. Launch Event data are known space launches, either future or historic
	// records containing items such as the launch site, launch epoch, and object.
	LaunchEvent LaunchEventService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	LaunchSite LaunchSiteService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	LaunchSiteDetails LaunchSiteDetailService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	LaunchVehicle LaunchVehicleService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	LaunchVehicleDetails LaunchVehicleDetailService
	// These services provide operations for manipulation and querying tactical data
	// links and link statuses of beams or a satellite constellation. Communication
	// link statuses provide definitions and status such as, positional endpoints,
	// where each endpoint may be associated with a specific beam or with a satellite
	// constellation. Data links provide detailed instructions regarding the
	// operational use of a tactical data link and interface coordination through
	// various message formats.
	LinkStatus LinkStatusService
	// These services provide operations for manipulation and querying tactical data
	// links and link statuses of beams or a satellite constellation. Communication
	// link statuses provide definitions and status such as, positional endpoints,
	// where each endpoint may be associated with a specific beam or with a satellite
	// constellation. Data links provide detailed instructions regarding the
	// operational use of a tactical data link and interface coordination through
	// various message formats.
	Linkstatus LinkstatusService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Location LocationService
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
	LogisticsSupport LogisticsSupportService
	// This service provides operations for querying and manipulation of
	// detected/possible/confirmed on-orbit maneuvers. The J2000 coordinate frame is
	// the preferred frame for all maneuver data, but in some cases data may be in
	// another frame depending on the provider. Check the Storefront 'Data Products'
	// section under the 'Discover' tab for maneuver data provider information.
	Maneuvers ManeuverService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Manifold ManifoldService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Manifoldelset ManifoldelsetService
	// These services provide operations for posting and querying of air, space, and
	// ground 'tracks'. A track is a position and optionally a heading/velocity of an
	// object at a particular timestamp.
	MissileTracks MissileTrackService
	// These services provide operations for manipulation and querying of mission
	// assignment objects. MissionAssignment is used by C2 JUs and, optionally, non-C2
	// JUs to assign missions, designate targets, and provide target information to
	// non-C2 JU platforms. Provision is made for the non-C2 JU platforms to
	// acknowledge the message through receipt/compliance action.
	MissionAssignment MissionAssignmentService
	// These services provide operations for posting and querying Moving Target
	// Indicator (MTI) STANAG 4607 data. Detailed MTI data supports activities such as
	// targeting or less detailed data for applications such as situational awareness
	// used/derived by exploitation systems.
	Mti MtiService
	// These services provide operations for manipulation and querying of on-orbit
	// payloads.
	Navigation NavigationService
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
	NavigationalObstruction NavigationalObstructionService
	// A Notification Service allowing the broadcast of generic messages to the
	// community. Users can send free-form messages, publish lists, and notify the
	// community about events or alerts across various domains. Notifications and
	// alerts are categorized by a 'msgType' field and are accessible via the UDL
	// Secure Messaging API and REST API services.
	Notification NotificationService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	ObjectOfInterest ObjectOfInterestService
	Observations     ObservationService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Onboardnavigation OnboardnavigationService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbit OnorbitService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitantenna OnorbitantennaService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitbattery OnorbitbatteryService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitdetails OnorbitdetailService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitevent OnorbiteventService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitlist OnorbitlistService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitsolararray OnorbitsolararrayService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitthruster OnorbitthrusterService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	Onorbitthrusterstatus OnorbitthrusterstatusService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Onorbitassessment OnorbitassessmentService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Operatingunit OperatingunitService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Operatingunitremark OperatingunitremarkService
	// These services provide operations for manipulating and querying Orbit
	// Determination (OD) data. Orbit Determination data contains algorithm results
	// that describe General Perturbations or Special Perturbations orbital updates.
	Orbitdetermination OrbitdeterminationService
	// These services provide operations for posting and querying of air, space, and
	// ground 'tracks'. A track is a position and optionally a heading/velocity of an
	// object at a particular timestamp.
	Orbittrack OrbittrackService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Organization OrganizationService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Organizationdetails OrganizationdetailService
	// These services provide operations for manipulation and querying of Mission Ops
	// information.
	Personnelrecovery PersonnelrecoveryService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Poi PoiService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	Port                PortService
	ReportAndActivities ReportAndActivityService
	// This collection of services provides operations for querying and manipulation of
	// RF related information to include RFEmitters which could potentially interfere
	// with communications/operations of space related entities, and RFBands commonly
	// used by various space related entities.
	RfBand RfBandService
	// This collection of services provides operations for querying and manipulation of
	// RF related information to include RFEmitters which could potentially interfere
	// with communications/operations of space related entities, and RFBands commonly
	// used by various space related entities.
	RfBandType RfBandTypeService
	// This collection of services provides operations for querying and manipulation of
	// RF related information to include RFEmitters which could potentially interfere
	// with communications/operations of space related entities, and RFBands commonly
	// used by various space related entities.
	RfEmitter RfEmitterService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	RouteStats RouteStatService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	SarObservation SarObservationService
	// These services provide operations for manipulation and querying of on-orbit
	// payloads.
	Scientific ScientificService
	Scs        ScService
	// Secure Messaging is based on Apache Kafka which is an open-source
	// stream-processing software platform developed by the Apache Software Foundation,
	// written in Scala and Java. Kafka provides a unified, high-throughput,
	// low-latency platform for handling real-time data feeds. All messaging is
	// secured; consumers will not receive messages unless authorized to do so. J2000
	// is the preferred coordinate frame for all observations, but in some cases
	// observations may be in another frame depending on the provider. Please see the
	// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
	SecureMessaging SecureMessagingService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	Sensor SensorService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SensorStating SensorStatingService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SensorMaintenance SensorMaintenanceService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SensorObservationType SensorObservationTypeService
	// These services provide operations for posting and querying Sensor Tasking data.
	SensorPlan SensorPlanService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SensorType SensorTypeService
	// These services provide operations for manipulation and querying of on-orbit
	// communications payloads (Comm), including supporting data such as transponders
	// and channels, etc.
	SeraDataCommDetails SeraDataCommDetailService
	// These services provide operations for manipulation and querying of on-orbit
	// payloads.
	SeraDataEarlyWarning SeraDataEarlyWarningService
	// These services provide operations for manipulation and querying of on-orbit
	// payloads.
	SeraDataNavigation SeraDataNavigationService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SeradataOpticalPayload SeradataOpticalPayloadService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SeradataRadarPayload SeradataRadarPayloadService
	// This service provides operations for querying and manipulation of sensor data.
	// Sensors are terrestrial or on-orbit equipment capable of taking measurements or
	// 'observations' of on-orbit objects via several phenomenologies such as
	// Electro-Optical (EO), Radar, and Radio Frequency (RF). This collection of
	// operations includes 'SensorMaintenance' schedules which define known/planned
	// future maintenance and associated operational impact of sensors as well as
	// 'SensorCalibration' records which contains data about a sensor's overall
	// accuracy and is used to adjust sensor settings.
	SeradataSigintPayload SeradataSigintPayloadService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	SeradataSpacecraftDetails SeradataSpacecraftDetailService
	// This service provides operations for manipulation and querying of space
	// weather/solar, geomagnetic, and radiation belt index data.
	Sgi SgiService
	// These services provide operations for manipulation and querying of Report and
	// Activity information. This information includes analytic reports, significant
	// events, route statistics, EMI Reports, and other georeferenced reports and
	// activities.
	Sigact SigactService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	Site SiteService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	SiteRemark SiteRemarkService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	SiteStatus SiteStatusService
	// This collection of services provides operations for querying and manipulation of
	// sky imagery data. Sky imagery is ground or space based telescope imagery of
	// RSO's and includes metadata on the image (time, source, etc) as well as binary
	// image content (e.g. FITS, EOSSA, EOCHIP, MP4). Binary content must be downloaded
	// individually by ID using the 'getFile' operation.
	SkyImagery SkyImageryService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	SoiObservationSet SoiObservationSetService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	SolarArray SolarArrayService
	// These services provide operations for manipulation and querying of on-orbit
	// objects of interest, their components, and various lists and status of those
	// objects.
	SolarArrayDetails SolarArrayDetailService
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
	SortiePpr SortiePprService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	SpaceEnvObservation SpaceEnvObservationService
	// Collection of launch related services which provide operations for querying and
	// manipulation of launch site data and detailed information on launch vehicles
	// including engines, stages, and manufacturers. Sites, engines, and stages can
	// each have multiple 'detail' records which may be compiled by different sources.
	Stage StageService
	// These services provide operations for posting and querying Star Catalog data.
	// The Star Catalog model is a representation of astronomical data and photometric
	// data for stars. Astronomical data includes positional information, proper
	// motions, parallaxes and their respective uncertainties. Photometric data
	// contains optical and near-infrared magnitudes, and their uncertainties across
	// multiple bandpasses. Note: Multiple source catalogs may contribute to a single
	// record.
	StarCatalog StarCatalogService
	// This service provides operations for querying and manipulation of state vectors
	// for On-orbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that together with their time (epoch) (t) uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVector StateVectorService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Status StatusService
	// Service operations for querying and manipulation of miscellaneous supporting
	// data such as countries (which can represent countries, multi-national
	// consortiums, and international organizations), data owners, locations, entities,
	// organizations, etc.
	Substatus      SubstatusService
	SupportingData SupportingDataService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	Surface SurfaceService
	// This collection of services provide operations for manipulating and querying of
	// various site related data, including site status, site operations, and site
	// type-specific records.
	SurfaceObstruction SurfaceObstructionService
	// This collection of services provides operations for querying and manipulation of
	// electro-optical (EO), radar, radio frequency (RF), Global Navigation Satellite
	// Systems (GNSS), Ionospheric (IONO), Infrared (SWIR), and Space Environment
	// observation data. The J2000 coordinate frame is the preferred frame for all
	// observations, as applicable, but in some cases observations may be in an
	// alternate frame depending on the provider and/or datatype.
	Swir SwirService
	// This service provides operations for manipulation and querying of earth
	// orientation parameter (EOP) data. Earth Orientation Parameters (EOP) are
	// produced by the IERS (International Earth Rotation and Reference Systems
	// Service). Earth Orientation Parameters describe the irregularities of the
	// earth's rotation. Technically, they are the parameters which provide the
	// rotation of the ITRS (International Terrestrial Reference System) to the ICRS
	// (International Celestial Reference System) as a function of time. Universal time
	// -- Universal time (UT1) is the time of the earth clock, which performs one
	// revolution in about 24h. It is practically proportional to the sidereal time.
	// The excess revolution time is called length of day (LOD). Coordinates of the
	// pole -- x and y are the coordinates of the Celestial Ephemeris Pole (CEP)
	// relative to the IRP, the IERS Reference Pole. The CEP differs from the
	// instantaneous rotation axis by quasi-diurnal terms with amplitudes under 0.01"
	// (see Seidelmann, 1982). The x-axis is in the direction of the ITRF
	// zero-meridian; the y-axis is in the direction 90 degrees West longitude.
	// Celestial pole offsets -- Celestial pole offsets are described in the IAU
	// Precession and Nutation models. The observed differences with respect to the
	// conventional celestial pole position defined by the models are monitored and
	// reported by the IERS. IERS Bulletins A and B provide current information on the
	// Earth's orientation in the IERS Reference System. This includes Universal Time,
	// coordinates of the terrestrial pole, and celestial pole offsets. Bulletin A
	// gives an advanced solution updated weekly; the standard solution is given
	// monthly in Bulletin B. Fields suffixed with ”B” are Bulletin B values. All
	// solutions are continuous within their respective uncertainties. Bulletin A is
	// issued by the IERS Rapid Service/Prediction Centre at the U.S. Naval
	// Observatory, Washington, DC and Bulletin B is issued by the IERS Earth
	// Orientation Centre at the Paris Observatory. IERS Bulletin A reports the latest
	// determinations for polar motion, UT1-UTC, and nutation offsets at daily
	// intervals based on a combination of contributed analysis results using data from
	// Very Long Baseline Interferometry (VLBI), Satellite Laser Ranging (SLR), Global
	// Positioning System (GPS) satellites, and Lunar Laser Ranging (LLR). Predictions
	// for variations a year into the future are also provided. Meteorological
	// predictions of variations in Atmospheric Angular Momentum (AAM) are used to aid
	// in the prediction of near-term UT1-UTC changes. This publication is prepared by
	// the IERS Rapid Service/Prediction Center.
	TaiUtc   TaiUtcService
	TdoaFdoa TdoaFdoaService
	// These services provide operations for posting and querying of air, space, and
	// ground 'tracks'. A track is a position and optionally a heading/velocity of an
	// object at a particular timestamp.
	Track TrackService
	// These services provide operations for posting and querying of air, space, and
	// ground 'tracks'. A track is a position and optionally a heading/velocity of an
	// object at a particular timestamp.
	TrackDetails TrackDetailService
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
	TrackRoute TrackRouteService
	// These services provide operations for manipulation and querying of on-orbit
	// communications payloads (Comm), including supporting data such as transponders
	// and channels, etc.
	Transponder TransponderService
	User        UserService
	// This service provides operations for manipulation and querying of maritime
	// Vessel and Vessel Status data. Vessel contains the static data of the specific
	// vessel: mmsi, cruise speed, max speed, etc.
	Vessel VesselService
	// This collection of services provides operations for video streaming.
	Video VideoService
	// These services provide for posting and querying terrestrial weather conditions
	// over a target area or region and raw sensor data used to produce condition
	// reports. Weather Reports describe current weather conditions over a target point
	// or region to include general temperatures, pressures, and moisture accumulation,
	// as well as navigational considerations such as altimeter settings, visibility,
	// wind speeds, and cloud heights etc. Weather Data contains algorithmic parameters
	// and dynamic, raw measurements collected by individual sensors such as signal
	// power, noise level, etc., which are generally processed across multiple sensors
	// to produce weather reports.
	WeatherData WeatherDataService
	// These services provide for posting and querying terrestrial weather conditions
	// over a target area or region and raw sensor data used to produce condition
	// reports. Weather Reports describe current weather conditions over a target point
	// or region to include general temperatures, pressures, and moisture accumulation,
	// as well as navigational considerations such as altimeter settings, visibility,
	// wind speeds, and cloud heights etc. Weather Data contains algorithmic parameters
	// and dynamic, raw measurements collected by individual sensors such as signal
	// power, noise level, etc., which are generally processed across multiple sensors
	// to produce weather reports.
	WeatherReport WeatherReportService
}

// DefaultClientOptions read from the environment (UDL_ACCESS_TOKEN,
// UDL_AUTH_PASSWORD, UDL_AUTH_USERNAME, UNIFIEDDATALIBRARY_BASE_URL). This should
// be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("UNIFIEDDATALIBRARY_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("UDL_ACCESS_TOKEN"); ok {
		defaults = append(defaults, option.WithAccessToken(o))
	}
	if o, ok := os.LookupEnv("UDL_AUTH_PASSWORD"); ok {
		defaults = append(defaults, option.WithPassword(o))
	}
	if o, ok := os.LookupEnv("UDL_AUTH_USERNAME"); ok {
		defaults = append(defaults, option.WithUsername(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (UDL_ACCESS_TOKEN, UDL_AUTH_PASSWORD, UDL_AUTH_USERNAME,
// UNIFIEDDATALIBRARY_BASE_URL). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.AirEvents = NewAirEventService(opts...)
	r.AirOperations = NewAirOperationService(opts...)
	r.AirTransportMissions = NewAirTransportMissionService(opts...)
	r.Aircraft = NewAircraftService(opts...)
	r.AircraftSorties = NewAircraftSortyService(opts...)
	r.AircraftStatusRemarks = NewAircraftStatusRemarkService(opts...)
	r.AircraftStatuses = NewAircraftStatusService(opts...)
	r.AirfieldSlotConsumptions = NewAirfieldSlotConsumptionService(opts...)
	r.AirfieldSlots = NewAirfieldSlotService(opts...)
	r.AirfieldStatus = NewAirfieldStatusService(opts...)
	r.Airfields = NewAirfieldService(opts...)
	r.AirloadPlans = NewAirloadPlanService(opts...)
	r.AirspaceControlOrders = NewAirspaceControlOrderService(opts...)
	r.AIs = NewAIService(opts...)
	r.AIsObjects = NewAIsObjectService(opts...)
	r.AnalyticImagery = NewAnalyticImageryService(opts...)
	r.Antennas = NewAntennaService(opts...)
	r.AttitudeData = NewAttitudeDataService(opts...)
	r.AttitudeSets = NewAttitudeSetService(opts...)
	r.AviationRiskManagement = NewAviationRiskManagementService(opts...)
	r.Batteries = NewBatteryService(opts...)
	r.Batterydetails = NewBatterydetailService(opts...)
	r.Beam = NewBeamService(opts...)
	r.BeamContours = NewBeamContourService(opts...)
	r.Buses = NewBusService(opts...)
	r.Channels = NewChannelService(opts...)
	r.Closelyspacedobjects = NewCloselyspacedobjectService(opts...)
	r.CollectRequests = NewCollectRequestService(opts...)
	r.CollectResponses = NewCollectResponseService(opts...)
	r.Comm = NewCommService(opts...)
	r.Conjunctions = NewConjunctionService(opts...)
	r.Cots = NewCotService(opts...)
	r.Countries = NewCountryService(opts...)
	r.Crew = NewCrewService(opts...)
	r.Deconflictset = NewDeconflictsetService(opts...)
	r.DiffOfArrival = NewDiffOfArrivalService(opts...)
	r.DiplomaticClearance = NewDiplomaticClearanceService(opts...)
	r.DriftHistory = NewDriftHistoryService(opts...)
	r.Dropzone = NewDropzoneService(opts...)
	r.Ecpedr = NewEcpedrService(opts...)
	r.EffectRequests = NewEffectRequestService(opts...)
	r.EffectResponses = NewEffectResponseService(opts...)
	r.Elsets = NewElsetService(opts...)
	r.Emireport = NewEmireportService(opts...)
	r.EmitterGeolocation = NewEmitterGeolocationService(opts...)
	r.EngineDetails = NewEngineDetailService(opts...)
	r.Engines = NewEngineService(opts...)
	r.Entities = NewEntityService(opts...)
	r.Eop = NewEopService(opts...)
	r.Ephemeris = NewEphemerisService(opts...)
	r.EphemerisSets = NewEphemerisSetService(opts...)
	r.Equipment = NewEquipmentService(opts...)
	r.EquipmentRemarks = NewEquipmentRemarkService(opts...)
	r.Evac = NewEvacService(opts...)
	r.EventEvolution = NewEventEvolutionService(opts...)
	r.FeatureAssessment = NewFeatureAssessmentService(opts...)
	r.Flightplan = NewFlightplanService(opts...)
	r.GeoStatus = NewGeoStatusService(opts...)
	r.GlobalAtmosphericModel = NewGlobalAtmosphericModelService(opts...)
	r.GnssObservations = NewGnssObservationService(opts...)
	r.GnssObservationset = NewGnssObservationsetService(opts...)
	r.GnssRawIf = NewGnssRawIfService(opts...)
	r.GroundImagery = NewGroundImageryService(opts...)
	r.H3Geo = NewH3GeoService(opts...)
	r.H3GeoHexCell = NewH3GeoHexCellService(opts...)
	r.Hazard = NewHazardService(opts...)
	r.IonoObservations = NewIonoObservationService(opts...)
	r.Ir = NewIrService(opts...)
	r.IsrCollections = NewIsrCollectionService(opts...)
	r.Item = NewItemService(opts...)
	r.ItemTrackings = NewItemTrackingService(opts...)
	r.Laserdeconflictrequest = NewLaserdeconflictrequestService(opts...)
	r.Laseremitter = NewLaseremitterService(opts...)
	r.LaunchDetection = NewLaunchDetectionService(opts...)
	r.LaunchEvent = NewLaunchEventService(opts...)
	r.LaunchSite = NewLaunchSiteService(opts...)
	r.LaunchSiteDetails = NewLaunchSiteDetailService(opts...)
	r.LaunchVehicle = NewLaunchVehicleService(opts...)
	r.LaunchVehicleDetails = NewLaunchVehicleDetailService(opts...)
	r.LinkStatus = NewLinkStatusService(opts...)
	r.Linkstatus = NewLinkstatusService(opts...)
	r.Location = NewLocationService(opts...)
	r.LogisticsSupport = NewLogisticsSupportService(opts...)
	r.Maneuvers = NewManeuverService(opts...)
	r.Manifold = NewManifoldService(opts...)
	r.Manifoldelset = NewManifoldelsetService(opts...)
	r.MissileTracks = NewMissileTrackService(opts...)
	r.MissionAssignment = NewMissionAssignmentService(opts...)
	r.Mti = NewMtiService(opts...)
	r.Navigation = NewNavigationService(opts...)
	r.NavigationalObstruction = NewNavigationalObstructionService(opts...)
	r.Notification = NewNotificationService(opts...)
	r.ObjectOfInterest = NewObjectOfInterestService(opts...)
	r.Observations = NewObservationService(opts...)
	r.Onboardnavigation = NewOnboardnavigationService(opts...)
	r.Onorbit = NewOnorbitService(opts...)
	r.Onorbitantenna = NewOnorbitantennaService(opts...)
	r.Onorbitbattery = NewOnorbitbatteryService(opts...)
	r.Onorbitdetails = NewOnorbitdetailService(opts...)
	r.Onorbitevent = NewOnorbiteventService(opts...)
	r.Onorbitlist = NewOnorbitlistService(opts...)
	r.Onorbitsolararray = NewOnorbitsolararrayService(opts...)
	r.Onorbitthruster = NewOnorbitthrusterService(opts...)
	r.Onorbitthrusterstatus = NewOnorbitthrusterstatusService(opts...)
	r.Onorbitassessment = NewOnorbitassessmentService(opts...)
	r.Operatingunit = NewOperatingunitService(opts...)
	r.Operatingunitremark = NewOperatingunitremarkService(opts...)
	r.Orbitdetermination = NewOrbitdeterminationService(opts...)
	r.Orbittrack = NewOrbittrackService(opts...)
	r.Organization = NewOrganizationService(opts...)
	r.Organizationdetails = NewOrganizationdetailService(opts...)
	r.Personnelrecovery = NewPersonnelrecoveryService(opts...)
	r.Poi = NewPoiService(opts...)
	r.Port = NewPortService(opts...)
	r.ReportAndActivities = NewReportAndActivityService(opts...)
	r.RfBand = NewRfBandService(opts...)
	r.RfBandType = NewRfBandTypeService(opts...)
	r.RfEmitter = NewRfEmitterService(opts...)
	r.RouteStats = NewRouteStatService(opts...)
	r.SarObservation = NewSarObservationService(opts...)
	r.Scientific = NewScientificService(opts...)
	r.Scs = NewScService(opts...)
	r.SecureMessaging = NewSecureMessagingService(opts...)
	r.Sensor = NewSensorService(opts...)
	r.SensorStating = NewSensorStatingService(opts...)
	r.SensorMaintenance = NewSensorMaintenanceService(opts...)
	r.SensorObservationType = NewSensorObservationTypeService(opts...)
	r.SensorPlan = NewSensorPlanService(opts...)
	r.SensorType = NewSensorTypeService(opts...)
	r.SeraDataCommDetails = NewSeraDataCommDetailService(opts...)
	r.SeraDataEarlyWarning = NewSeraDataEarlyWarningService(opts...)
	r.SeraDataNavigation = NewSeraDataNavigationService(opts...)
	r.SeradataOpticalPayload = NewSeradataOpticalPayloadService(opts...)
	r.SeradataRadarPayload = NewSeradataRadarPayloadService(opts...)
	r.SeradataSigintPayload = NewSeradataSigintPayloadService(opts...)
	r.SeradataSpacecraftDetails = NewSeradataSpacecraftDetailService(opts...)
	r.Sgi = NewSgiService(opts...)
	r.Sigact = NewSigactService(opts...)
	r.Site = NewSiteService(opts...)
	r.SiteRemark = NewSiteRemarkService(opts...)
	r.SiteStatus = NewSiteStatusService(opts...)
	r.SkyImagery = NewSkyImageryService(opts...)
	r.SoiObservationSet = NewSoiObservationSetService(opts...)
	r.SolarArray = NewSolarArrayService(opts...)
	r.SolarArrayDetails = NewSolarArrayDetailService(opts...)
	r.SortiePpr = NewSortiePprService(opts...)
	r.SpaceEnvObservation = NewSpaceEnvObservationService(opts...)
	r.Stage = NewStageService(opts...)
	r.StarCatalog = NewStarCatalogService(opts...)
	r.StateVector = NewStateVectorService(opts...)
	r.Status = NewStatusService(opts...)
	r.Substatus = NewSubstatusService(opts...)
	r.SupportingData = NewSupportingDataService(opts...)
	r.Surface = NewSurfaceService(opts...)
	r.SurfaceObstruction = NewSurfaceObstructionService(opts...)
	r.Swir = NewSwirService(opts...)
	r.TaiUtc = NewTaiUtcService(opts...)
	r.TdoaFdoa = NewTdoaFdoaService(opts...)
	r.Track = NewTrackService(opts...)
	r.TrackDetails = NewTrackDetailService(opts...)
	r.TrackRoute = NewTrackRouteService(opts...)
	r.Transponder = NewTransponderService(opts...)
	r.User = NewUserService(opts...)
	r.Vessel = NewVesselService(opts...)
	r.Video = NewVideoService(opts...)
	r.WeatherData = NewWeatherDataService(opts...)
	r.WeatherReport = NewWeatherReportService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
