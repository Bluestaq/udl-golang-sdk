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
	AirTransportMissions      AirTransportMissionService
	Aircraft                  AircraftService
	AircraftSorties           AircraftSortyService
	AircraftStatusRemarks     AircraftStatusRemarkService
	AircraftStatuses          AircraftStatusService
	AirfieldSlots             AirfieldSlotService
	AirfieldStatus            AirfieldStatusService
	Airfields                 AirfieldService
	AirfieldSlotConsumptions  AirfieldSlotConsumptionService
	AirloadPlans              AirloadPlanService
	AirspaceControlOrders     AirspaceControlOrderService
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
	AviationRiskManagement    AviationRiskManagementService
	Dropzone                  DropzoneService
	EmitterGeolocation        EmitterGeolocationService
	FeatureAssessment         FeatureAssessmentService
	GlobalAtmosphericModel    GlobalAtmosphericModelService
	RouteStats                RouteStatService
	Countries                 CountryService
	Crew                      CrewService
	DiffOfArrival             DiffOfArrivalService
	DiplomaticClearance       DiplomaticClearanceService
	DriftHistory              DriftHistoryService
	EcpSdr                    EcpSdrService
	EffectRequests            EffectRequestService
	EffectResponses           EffectResponseService
	Elsets                    ElsetService
	EngineDetails             EngineDetailService
	Engines                   EngineService
	Entities                  EntityService
	Eop                       EopService
	Ephemeris                 EphemerisService
	EphemerisSets             EphemerisSetService
	Equipment                 EquipmentService
	EquipmentRemarks          EquipmentRemarkService
	Evac                      EvacService
	EventEvolution            EventEvolutionService
	Flightplan                FlightplanService
	GeoStatus                 GeoStatusService
	GnssObservationset        GnssObservationsetService
	GnssRawif                 GnssRawifService
	GroundImagery             GroundImageryService
	H3Geo                     H3GeoService
	H3GeoHexCell              H3GeoHexCellService
	Hazard                    HazardService
	IonOobservation           IonOobservationService
	Ir                        IrService
	IsrCollections            IsrCollectionService
	Item                      ItemService
	ItemTrackings             ItemTrackingService
	LaunchDetection           LaunchDetectionService
	LaunchEvent               LaunchEventService
	LaunchSite                LaunchSiteService
	LaunchSiteDetails         LaunchSiteDetailService
	LaunchVehicle             LaunchVehicleService
	LaunchVehicleDetails      LaunchVehicleDetailService
	LinkStatus                LinkStatusService
	Location                  LocationService
	LogisticsSupport          LogisticsSupportService
	Maneuvers                 ManeuverService
	Manifold                  ManifoldService
	Manifoldelset             ManifoldelsetService
	MissileTracks             MissileTrackService
	MissionAssignment         MissionAssignmentService
	Mti                       MtiService
	Navigation                NavigationService
	NavigationalObstruction   NavigationalObstructionService
	Notification              NotificationService
	ObjectOfInterest          ObjectOfInterestService
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
	Personnelrecovery         PersonnelrecoveryService
	Poi                       PoiService
	Port                      PortService
	RfBand                    RfBandService
	RfBandType                RfBandTypeService
	RfEmitter                 RfEmitterService
	RfEmitterDetails          RfEmitterDetailService
	SarObservation            SarObservationService
	Scientific                ScientificService
	Sensor                    SensorService
	SensorMaintenance         SensorMaintenanceService
	SensorObservationType     SensorObservationTypeService
	SensorPlan                SensorPlanService
	SensorType                SensorTypeService
	SeraDataCommDetails       SeraDataCommDetailService
	SeraDataEarlyWarning      SeraDataEarlyWarningService
	SeraDataNavigation        SeraDataNavigationService
	SeradataOpticalPayload    SeradataOpticalPayloadService
	SeradataRadarPayload      SeradataRadarPayloadService
	SeradataSigintPayload     SeradataSigintPayloadService
	SeradataSpacecraftDetails SeradataSpacecraftDetailService
	Sgi                       SgiService
	Sigact                    SigactService
	Site                      SiteService
	SiteRemark                SiteRemarkService
	SiteStatus                SiteStatusService
	SkyImagery                SkyImageryService
	SoiObservationSet         SoiObservationSetService
	SolarArray                SolarArrayService
	SolarArrayDetails         SolarArrayDetailService
	SortiePpr                 SortiePprService
	SpaceEnvObservation       SpaceEnvObservationService
	Stage                     StageService
	StarCatalog               StarCatalogService
	StateVector               StateVectorService
	Status                    StatusService
	Substatus                 SubstatusService
	SupportingData            SupportingDataService
	Surface                   SurfaceService
	SurfaceObstruction        SurfaceObstructionService
	Swir                      SwirService
	TaiUtc                    TaiUtcService
	TdoaFdoa                  TdoaFdoaService
	Track                     TrackService
	TrackDetails              TrackDetailService
	TrackRoute                TrackRouteService
	Transponder               TransponderService
	Vessel                    VesselService
	Video                     VideoService
	WeatherData               WeatherDataService
	WeatherReport             WeatherReportService
	GnssObservations          GnssObservationService
	GnssRawIf                 GnssRawIfService
	IonoObservation           IonoObservationService
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
	r.AirTransportMissions = NewAirTransportMissionService(opts...)
	r.Aircraft = NewAircraftService(opts...)
	r.AircraftSorties = NewAircraftSortyService(opts...)
	r.AircraftStatusRemarks = NewAircraftStatusRemarkService(opts...)
	r.AircraftStatuses = NewAircraftStatusService(opts...)
	r.AirfieldSlots = NewAirfieldSlotService(opts...)
	r.AirfieldStatus = NewAirfieldStatusService(opts...)
	r.Airfields = NewAirfieldService(opts...)
	r.AirfieldSlotConsumptions = NewAirfieldSlotConsumptionService(opts...)
	r.AirloadPlans = NewAirloadPlanService(opts...)
	r.AirspaceControlOrders = NewAirspaceControlOrderService(opts...)
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
	r.AviationRiskManagement = NewAviationRiskManagementService(opts...)
	r.Dropzone = NewDropzoneService(opts...)
	r.EmitterGeolocation = NewEmitterGeolocationService(opts...)
	r.FeatureAssessment = NewFeatureAssessmentService(opts...)
	r.GlobalAtmosphericModel = NewGlobalAtmosphericModelService(opts...)
	r.RouteStats = NewRouteStatService(opts...)
	r.Countries = NewCountryService(opts...)
	r.Crew = NewCrewService(opts...)
	r.DiffOfArrival = NewDiffOfArrivalService(opts...)
	r.DiplomaticClearance = NewDiplomaticClearanceService(opts...)
	r.DriftHistory = NewDriftHistoryService(opts...)
	r.EcpSdr = NewEcpSdrService(opts...)
	r.EffectRequests = NewEffectRequestService(opts...)
	r.EffectResponses = NewEffectResponseService(opts...)
	r.Elsets = NewElsetService(opts...)
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
	r.Flightplan = NewFlightplanService(opts...)
	r.GeoStatus = NewGeoStatusService(opts...)
	r.GnssObservationset = NewGnssObservationsetService(opts...)
	r.GnssRawif = NewGnssRawifService(opts...)
	r.GroundImagery = NewGroundImageryService(opts...)
	r.H3Geo = NewH3GeoService(opts...)
	r.H3GeoHexCell = NewH3GeoHexCellService(opts...)
	r.Hazard = NewHazardService(opts...)
	r.IonOobservation = NewIonOobservationService(opts...)
	r.Ir = NewIrService(opts...)
	r.IsrCollections = NewIsrCollectionService(opts...)
	r.Item = NewItemService(opts...)
	r.ItemTrackings = NewItemTrackingService(opts...)
	r.LaunchDetection = NewLaunchDetectionService(opts...)
	r.LaunchEvent = NewLaunchEventService(opts...)
	r.LaunchSite = NewLaunchSiteService(opts...)
	r.LaunchSiteDetails = NewLaunchSiteDetailService(opts...)
	r.LaunchVehicle = NewLaunchVehicleService(opts...)
	r.LaunchVehicleDetails = NewLaunchVehicleDetailService(opts...)
	r.LinkStatus = NewLinkStatusService(opts...)
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
	r.Operatingunit = NewOperatingunitService(opts...)
	r.Operatingunitremark = NewOperatingunitremarkService(opts...)
	r.Orbitdetermination = NewOrbitdeterminationService(opts...)
	r.Orbittrack = NewOrbittrackService(opts...)
	r.Organization = NewOrganizationService(opts...)
	r.Organizationdetails = NewOrganizationdetailService(opts...)
	r.Personnelrecovery = NewPersonnelrecoveryService(opts...)
	r.Poi = NewPoiService(opts...)
	r.Port = NewPortService(opts...)
	r.RfBand = NewRfBandService(opts...)
	r.RfBandType = NewRfBandTypeService(opts...)
	r.RfEmitter = NewRfEmitterService(opts...)
	r.RfEmitterDetails = NewRfEmitterDetailService(opts...)
	r.SarObservation = NewSarObservationService(opts...)
	r.Scientific = NewScientificService(opts...)
	r.Sensor = NewSensorService(opts...)
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
	r.Vessel = NewVesselService(opts...)
	r.Video = NewVideoService(opts...)
	r.WeatherData = NewWeatherDataService(opts...)
	r.WeatherReport = NewWeatherReportService(opts...)
	r.GnssObservations = NewGnssObservationService(opts...)
	r.GnssRawIf = NewGnssRawIfService(opts...)
	r.IonoObservation = NewIonoObservationService(opts...)
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
