// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// UdlService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUdlService] method instead.
type UdlService struct {
	Options                 []option.RequestOption
	Geostatus               UdlGeostatusService
	Gnssobservationset      UdlGnssobservationsetService
	Mti                     UdlMtiService
	Notification            UdlNotificationService
	Onboardnavigation       UdlOnboardnavigationService
	Onorbitthrusterstatus   UdlOnorbitthrusterstatusService
	Gnssrawif               UdlGnssrawifService
	Groundimagery           UdlGroundimageryService
	Hazard                  UdlHazardService
	Ionoobservation         UdlIonoobservationService
	Isrcollection           UdlIsrcollectionService
	Itemtracking            UdlItemtrackingService
	Linkstatus              UdlLinkstatusService
	Maneuver                UdlManeuverService
	Missiletrack            UdlMissiletrackService
	Missionassignment       UdlMissionassignmentService
	Monoradar               UdlMonoradarService
	Orbitdetermination      UdlOrbitdeterminationService
	Orbittrack              UdlOrbittrackService
	Passiveradarobservation UdlPassiveradarobservationService
	Poi                     UdlPoiService
	Radarobservation        UdlRadarobservationService
	Rfobservation           UdlRfobservationService
	Sarobservation          UdlSarobservationService
	Sensormaintenance       UdlSensormaintenanceService
	Sensorplan              UdlSensorplanService
	Sgi                     UdlSgiService
	Sigact                  UdlSigactService
	Sitestatus              UdlSitestatusService
	Skyimagery              UdlSkyimageryService
}

// NewUdlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUdlService(opts ...option.RequestOption) (r UdlService) {
	r = UdlService{}
	r.Options = opts
	r.Geostatus = NewUdlGeostatusService(opts...)
	r.Gnssobservationset = NewUdlGnssobservationsetService(opts...)
	r.Mti = NewUdlMtiService(opts...)
	r.Notification = NewUdlNotificationService(opts...)
	r.Onboardnavigation = NewUdlOnboardnavigationService(opts...)
	r.Onorbitthrusterstatus = NewUdlOnorbitthrusterstatusService(opts...)
	r.Gnssrawif = NewUdlGnssrawifService(opts...)
	r.Groundimagery = NewUdlGroundimageryService(opts...)
	r.Hazard = NewUdlHazardService(opts...)
	r.Ionoobservation = NewUdlIonoobservationService(opts...)
	r.Isrcollection = NewUdlIsrcollectionService(opts...)
	r.Itemtracking = NewUdlItemtrackingService(opts...)
	r.Linkstatus = NewUdlLinkstatusService(opts...)
	r.Maneuver = NewUdlManeuverService(opts...)
	r.Missiletrack = NewUdlMissiletrackService(opts...)
	r.Missionassignment = NewUdlMissionassignmentService(opts...)
	r.Monoradar = NewUdlMonoradarService(opts...)
	r.Orbitdetermination = NewUdlOrbitdeterminationService(opts...)
	r.Orbittrack = NewUdlOrbittrackService(opts...)
	r.Passiveradarobservation = NewUdlPassiveradarobservationService(opts...)
	r.Poi = NewUdlPoiService(opts...)
	r.Radarobservation = NewUdlRadarobservationService(opts...)
	r.Rfobservation = NewUdlRfobservationService(opts...)
	r.Sarobservation = NewUdlSarobservationService(opts...)
	r.Sensormaintenance = NewUdlSensormaintenanceService(opts...)
	r.Sensorplan = NewUdlSensorplanService(opts...)
	r.Sgi = NewUdlSgiService(opts...)
	r.Sigact = NewUdlSigactService(opts...)
	r.Sitestatus = NewUdlSitestatusService(opts...)
	r.Skyimagery = NewUdlSkyimageryService(opts...)
	return
}
