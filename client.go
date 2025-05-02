// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"os"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the unifieddatalibrary API. You should not instantiate this
// client directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                   []option.RequestOption
	AirEvents                 AirEventService
	AirLoadPlans              AirLoadPlanService
	AirOperations             AirOperationService
	AirTaskingOrders          AirTaskingOrderService
	AirTransportMissions      AirTransportMissionService
	Aircraft                  AircraftService
	AircraftSorties           AircraftSortyService
	AircraftStatusRemarks     AircraftStatusRemarkService
	AircraftStatuses          AircraftStatusService
	Aircraftstatusremark      AircraftstatusremarkService
	AirfieldSlots             AirfieldSlotService
	AirfieldStatus            AirfieldStatusService
	Airfields                 AirfieldService
	Airfieldslotconsumptions  AirfieldslotconsumptionService
	Airfieldslots             AirfieldslotService
	Airfieldstatus            AirfieldstatusService
	Airloadplans              AirloadplanService
	AirspaceControlOrders     AirspaceControlOrderService
	Airtaskingorders          AirtaskingorderService
	AIs                       AIService
	AIsObjects                AIsObjectService
	AnalyticImagery           AnalyticImageryService
	Antennas                  AntennaService
	AttitudeData              AttitudeDataService
	AttitudeSets              AttitudeSetService
	Attitudesets              AttitudesetService
	Batteries                 BatteryService
	Batterydetails            BatterydetailService
	Beam                      BeamService
	BeamContours              BeamContourService
	Buses                     BusService
	Channels                  ChannelService
	CollectRequests           CollectRequestService
	CollectResponses          CollectResponseService
	Comm                      CommService
	Conjunctions              ConjunctionService
	Cots                      CotService
	Aviationriskmanagement    AviationriskmanagementService
	Dropzone                  DropzoneService
	Emittergeolocation        EmittergeolocationService
	Featureassessment         FeatureassessmentService
	Globalatmosphericmodel    GlobalatmosphericmodelService
	Routestats                RoutestatService
	Countries                 CountryService
	Crew                      CrewService
	Diffofarrival             DiffofarrivalService
	DiplomaticClearance       DiplomaticClearanceService
	DriftHistory              DriftHistoryService
	Ecpsdr                    EcpsdrService
	EffectRequests            EffectRequestService
	EffectResponses           EffectResponseService
	Elsets                    ElsetService
	EngineDetails             EngineDetailService
	Enginedetails             EnginedetailService
	Engines                   EngineService
	Entities                  EntityService
	EoObservations            EoObservationService
	Eoobservations            EoobservationService
	Eop                       EopService
	Ephemeris                 EphemerisService
	EphemerisSets             EphemerisSetService
	Equipment                 EquipmentService
	Equipmentremarks          EquipmentremarkService
	Evac                      EvacService
	EventEvolution            EventEvolutionService
	Flightplan                FlightplanService
	Geostatus                 GeostatusService
	Gnssobservationset        GnssobservationsetService
	Gnssrawif                 GnssrawifService
	GroundImagery             GroundImageryService
	Groundimagery             GroundimageryService
	H3geo                     H3geoService
	H3geohexcell              H3geohexcellService
	Hazard                    HazardService
	Ionoobservation           IonoobservationService
	Ir                        IrService
	IsrCollections            IsrCollectionService
	Item                      ItemService
	ItemTrackings             ItemTrackingService
	Launchdetection           LaunchdetectionService
	Launchevent               LauncheventService
	Launchsite                LaunchsiteService
	Launchsitedetails         LaunchsitedetailService
	Launchvehicle             LaunchvehicleService
	Launchvehicledetails      LaunchvehicledetailService
	LinkStatus                LinkStatusService
	Location                  LocationService
	Logisticssupport          LogisticssupportService
	Maneuvers                 ManeuverService
	Manifold                  ManifoldService
	Manifoldelset             ManifoldelsetService
	MissileTracks             MissileTrackService
	Missionassignment         MissionassignmentService
	Monoradar                 MonoradarService
	Mti                       MtiService
	Navigation                NavigationService
	Navigationalobstruction   NavigationalobstructionService
	Notification              NotificationService
	Objectofinterest          ObjectofinterestService
	Observations              ObservationService
	Onboardnavigation         OnboardnavigationService
	Onorbit                   OnorbitService
	Onorbitantenna            OnorbitantennaService
	Onorbitbattery            OnorbitbatteryService
	Onorbitdetails            OnorbitdetailService
	Onorbitevent              OnorbiteventService
	Onorbitlist               OnorbitlistService
	Onorbitsolararray         OnorbitsolararrayService
	Onorbitthruster           OnorbitthrusterService
	Onorbitthrusterstatus     OnorbitthrusterstatusService
	Operatingunit             OperatingunitService
	Operatingunitremark       OperatingunitremarkService
	Orbitdetermination        OrbitdeterminationService
	Orbittrack                OrbittrackService
	Organization              OrganizationService
	Organizationdetails       OrganizationdetailService
	Passiveradarobservation   PassiveradarobservationService
	Personnelrecovery         PersonnelrecoveryService
	Poi                       PoiService
	Port                      PortService
	Radarobservation          RadarobservationService
	Rfband                    RfbandService
	Rfbandtype                RfbandtypeService
	Rfemitter                 RfemitterService
	Rfemitterdetails          RfemitterdetailService
	Rfobservation             RfobservationService
	Sarobservation            SarobservationService
	Scientific                ScientificService
	Sensor                    SensorService
	Sensormaintenance         SensormaintenanceService
	Sensorobservationtype     SensorobservationtypeService
	Sensorplan                SensorplanService
	Sensortype                SensortypeService
	Seradatacommdetails       SeradatacommdetailService
	Seradataearlywarning      SeradataearlywarningService
	Seradatanavigation        SeradatanavigationService
	Seradataopticalpayload    SeradataopticalpayloadService
	Seradataradarpayload      SeradataradarpayloadService
	Seradatasigintpayload     SeradatasigintpayloadService
	Seradataspacecraftdetails SeradataspacecraftdetailService
	Sgi                       SgiService
	Sigact                    SigactService
	Site                      SiteService
	Siteremark                SiteremarkService
	Sitestatus                SitestatusService
	Skyimagery                SkyimageryService
	Soiobservationset         SoiobservationsetService
	Solararray                SolararrayService
	Solararraydetails         SolararraydetailService
	Sortieppr                 SortiepprService
	Spaceenvobservation       SpaceenvobservationService
	Stage                     StageService
	Starcatalog               StarcatalogService
	Statevector               StatevectorService
	Status                    StatusService
	Substatus                 SubstatusService
	SupportingData            SupportingDataService
	Surface                   SurfaceService
	Surfaceobstruction        SurfaceobstructionService
	Swir                      SwirService
	Taiutc                    TaiutcService
	TdoaFdoa                  TdoaFdoaService
	Track                     TrackService
	Trackdetails              TrackdetailService
	Trackroute                TrackrouteService
	Transponder               TransponderService
	Vessel                    VesselService
	Video                     VideoService
	Weatherdata               WeatherdataService
	Weatherreport             WeatherreportService
	Udl                       UdlService
	GnssObservations          GnssObservationService
	GnssRawIf                 GnssRawIfService
	IonoObservation           IonoObservationService
	LaunchEvent               LaunchEventService
	ReportAndActivity         ReportAndActivityService
	SecureMessaging           SecureMessagingService
	Scs                       ScService
	ScsViews                  ScsViewService
}

// DefaultClientOptions read from the environment (UDL_AUTH_PASSWORD,
// UDL_AUTH_USERNAME, UNIFIEDDATALIBRARY_BASE_URL). This should be used to
// initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("UNIFIEDDATALIBRARY_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
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
// environment (UDL_AUTH_PASSWORD, UDL_AUTH_USERNAME, UNIFIEDDATALIBRARY_BASE_URL).
// The option passed in as arguments are applied after these default arguments, and
// all option will be passed down to the services and requests that this client
// makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.AirEvents = NewAirEventService(opts...)
	r.AirLoadPlans = NewAirLoadPlanService(opts...)
	r.AirOperations = NewAirOperationService(opts...)
	r.AirTaskingOrders = NewAirTaskingOrderService(opts...)
	r.AirTransportMissions = NewAirTransportMissionService(opts...)
	r.Aircraft = NewAircraftService(opts...)
	r.AircraftSorties = NewAircraftSortyService(opts...)
	r.AircraftStatusRemarks = NewAircraftStatusRemarkService(opts...)
	r.AircraftStatuses = NewAircraftStatusService(opts...)
	r.Aircraftstatusremark = NewAircraftstatusremarkService(opts...)
	r.AirfieldSlots = NewAirfieldSlotService(opts...)
	r.AirfieldStatus = NewAirfieldStatusService(opts...)
	r.Airfields = NewAirfieldService(opts...)
	r.Airfieldslotconsumptions = NewAirfieldslotconsumptionService(opts...)
	r.Airfieldslots = NewAirfieldslotService(opts...)
	r.Airfieldstatus = NewAirfieldstatusService(opts...)
	r.Airloadplans = NewAirloadplanService(opts...)
	r.AirspaceControlOrders = NewAirspaceControlOrderService(opts...)
	r.Airtaskingorders = NewAirtaskingorderService(opts...)
	r.AIs = NewAIService(opts...)
	r.AIsObjects = NewAIsObjectService(opts...)
	r.AnalyticImagery = NewAnalyticImageryService(opts...)
	r.Antennas = NewAntennaService(opts...)
	r.AttitudeData = NewAttitudeDataService(opts...)
	r.AttitudeSets = NewAttitudeSetService(opts...)
	r.Attitudesets = NewAttitudesetService(opts...)
	r.Batteries = NewBatteryService(opts...)
	r.Batterydetails = NewBatterydetailService(opts...)
	r.Beam = NewBeamService(opts...)
	r.BeamContours = NewBeamContourService(opts...)
	r.Buses = NewBusService(opts...)
	r.Channels = NewChannelService(opts...)
	r.CollectRequests = NewCollectRequestService(opts...)
	r.CollectResponses = NewCollectResponseService(opts...)
	r.Comm = NewCommService(opts...)
	r.Conjunctions = NewConjunctionService(opts...)
	r.Cots = NewCotService(opts...)
	r.Aviationriskmanagement = NewAviationriskmanagementService(opts...)
	r.Dropzone = NewDropzoneService(opts...)
	r.Emittergeolocation = NewEmittergeolocationService(opts...)
	r.Featureassessment = NewFeatureassessmentService(opts...)
	r.Globalatmosphericmodel = NewGlobalatmosphericmodelService(opts...)
	r.Routestats = NewRoutestatService(opts...)
	r.Countries = NewCountryService(opts...)
	r.Crew = NewCrewService(opts...)
	r.Diffofarrival = NewDiffofarrivalService(opts...)
	r.DiplomaticClearance = NewDiplomaticClearanceService(opts...)
	r.DriftHistory = NewDriftHistoryService(opts...)
	r.Ecpsdr = NewEcpsdrService(opts...)
	r.EffectRequests = NewEffectRequestService(opts...)
	r.EffectResponses = NewEffectResponseService(opts...)
	r.Elsets = NewElsetService(opts...)
	r.EngineDetails = NewEngineDetailService(opts...)
	r.Enginedetails = NewEnginedetailService(opts...)
	r.Engines = NewEngineService(opts...)
	r.Entities = NewEntityService(opts...)
	r.EoObservations = NewEoObservationService(opts...)
	r.Eoobservations = NewEoobservationService(opts...)
	r.Eop = NewEopService(opts...)
	r.Ephemeris = NewEphemerisService(opts...)
	r.EphemerisSets = NewEphemerisSetService(opts...)
	r.Equipment = NewEquipmentService(opts...)
	r.Equipmentremarks = NewEquipmentremarkService(opts...)
	r.Evac = NewEvacService(opts...)
	r.EventEvolution = NewEventEvolutionService(opts...)
	r.Flightplan = NewFlightplanService(opts...)
	r.Geostatus = NewGeostatusService(opts...)
	r.Gnssobservationset = NewGnssobservationsetService(opts...)
	r.Gnssrawif = NewGnssrawifService(opts...)
	r.GroundImagery = NewGroundImageryService(opts...)
	r.Groundimagery = NewGroundimageryService(opts...)
	r.H3geo = NewH3geoService(opts...)
	r.H3geohexcell = NewH3geohexcellService(opts...)
	r.Hazard = NewHazardService(opts...)
	r.Ionoobservation = NewIonoobservationService(opts...)
	r.Ir = NewIrService(opts...)
	r.IsrCollections = NewIsrCollectionService(opts...)
	r.Item = NewItemService(opts...)
	r.ItemTrackings = NewItemTrackingService(opts...)
	r.Launchdetection = NewLaunchdetectionService(opts...)
	r.Launchevent = NewLauncheventService(opts...)
	r.Launchsite = NewLaunchsiteService(opts...)
	r.Launchsitedetails = NewLaunchsitedetailService(opts...)
	r.Launchvehicle = NewLaunchvehicleService(opts...)
	r.Launchvehicledetails = NewLaunchvehicledetailService(opts...)
	r.LinkStatus = NewLinkStatusService(opts...)
	r.Location = NewLocationService(opts...)
	r.Logisticssupport = NewLogisticssupportService(opts...)
	r.Maneuvers = NewManeuverService(opts...)
	r.Manifold = NewManifoldService(opts...)
	r.Manifoldelset = NewManifoldelsetService(opts...)
	r.MissileTracks = NewMissileTrackService(opts...)
	r.Missionassignment = NewMissionassignmentService(opts...)
	r.Monoradar = NewMonoradarService(opts...)
	r.Mti = NewMtiService(opts...)
	r.Navigation = NewNavigationService(opts...)
	r.Navigationalobstruction = NewNavigationalobstructionService(opts...)
	r.Notification = NewNotificationService(opts...)
	r.Objectofinterest = NewObjectofinterestService(opts...)
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
	r.Operatingunit = NewOperatingunitService(opts...)
	r.Operatingunitremark = NewOperatingunitremarkService(opts...)
	r.Orbitdetermination = NewOrbitdeterminationService(opts...)
	r.Orbittrack = NewOrbittrackService(opts...)
	r.Organization = NewOrganizationService(opts...)
	r.Organizationdetails = NewOrganizationdetailService(opts...)
	r.Passiveradarobservation = NewPassiveradarobservationService(opts...)
	r.Personnelrecovery = NewPersonnelrecoveryService(opts...)
	r.Poi = NewPoiService(opts...)
	r.Port = NewPortService(opts...)
	r.Radarobservation = NewRadarobservationService(opts...)
	r.Rfband = NewRfbandService(opts...)
	r.Rfbandtype = NewRfbandtypeService(opts...)
	r.Rfemitter = NewRfemitterService(opts...)
	r.Rfemitterdetails = NewRfemitterdetailService(opts...)
	r.Rfobservation = NewRfobservationService(opts...)
	r.Sarobservation = NewSarobservationService(opts...)
	r.Scientific = NewScientificService(opts...)
	r.Sensor = NewSensorService(opts...)
	r.Sensormaintenance = NewSensormaintenanceService(opts...)
	r.Sensorobservationtype = NewSensorobservationtypeService(opts...)
	r.Sensorplan = NewSensorplanService(opts...)
	r.Sensortype = NewSensortypeService(opts...)
	r.Seradatacommdetails = NewSeradatacommdetailService(opts...)
	r.Seradataearlywarning = NewSeradataearlywarningService(opts...)
	r.Seradatanavigation = NewSeradatanavigationService(opts...)
	r.Seradataopticalpayload = NewSeradataopticalpayloadService(opts...)
	r.Seradataradarpayload = NewSeradataradarpayloadService(opts...)
	r.Seradatasigintpayload = NewSeradatasigintpayloadService(opts...)
	r.Seradataspacecraftdetails = NewSeradataspacecraftdetailService(opts...)
	r.Sgi = NewSgiService(opts...)
	r.Sigact = NewSigactService(opts...)
	r.Site = NewSiteService(opts...)
	r.Siteremark = NewSiteremarkService(opts...)
	r.Sitestatus = NewSitestatusService(opts...)
	r.Skyimagery = NewSkyimageryService(opts...)
	r.Soiobservationset = NewSoiobservationsetService(opts...)
	r.Solararray = NewSolararrayService(opts...)
	r.Solararraydetails = NewSolararraydetailService(opts...)
	r.Sortieppr = NewSortiepprService(opts...)
	r.Spaceenvobservation = NewSpaceenvobservationService(opts...)
	r.Stage = NewStageService(opts...)
	r.Starcatalog = NewStarcatalogService(opts...)
	r.Statevector = NewStatevectorService(opts...)
	r.Status = NewStatusService(opts...)
	r.Substatus = NewSubstatusService(opts...)
	r.SupportingData = NewSupportingDataService(opts...)
	r.Surface = NewSurfaceService(opts...)
	r.Surfaceobstruction = NewSurfaceobstructionService(opts...)
	r.Swir = NewSwirService(opts...)
	r.Taiutc = NewTaiutcService(opts...)
	r.TdoaFdoa = NewTdoaFdoaService(opts...)
	r.Track = NewTrackService(opts...)
	r.Trackdetails = NewTrackdetailService(opts...)
	r.Trackroute = NewTrackrouteService(opts...)
	r.Transponder = NewTransponderService(opts...)
	r.Vessel = NewVesselService(opts...)
	r.Video = NewVideoService(opts...)
	r.Weatherdata = NewWeatherdataService(opts...)
	r.Weatherreport = NewWeatherreportService(opts...)
	r.Udl = NewUdlService(opts...)
	r.GnssObservations = NewGnssObservationService(opts...)
	r.GnssRawIf = NewGnssRawIfService(opts...)
	r.IonoObservation = NewIonoObservationService(opts...)
	r.LaunchEvent = NewLaunchEventService(opts...)
	r.ReportAndActivity = NewReportAndActivityService(opts...)
	r.SecureMessaging = NewSecureMessagingService(opts...)
	r.Scs = NewScService(opts...)
	r.ScsViews = NewScsViewService(opts...)

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
	opts = append(r.Options, opts...)
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
