// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// WeatherReportHistoryService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeatherReportHistoryService] method instead.
type WeatherReportHistoryService struct {
	Options []option.RequestOption
}

// NewWeatherReportHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWeatherReportHistoryService(opts ...option.RequestOption) (r WeatherReportHistoryService) {
	r = WeatherReportHistoryService{}
	r.Options = opts
	return
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherReportHistoryService) List(ctx context.Context, query WeatherReportHistoryListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WeatherReportFull], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/weatherreport/history"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherReportHistoryService) ListAutoPaging(ctx context.Context, query WeatherReportHistoryListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WeatherReportFull] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to dynamically query historical data by a variety of query
// parameters not specified in this API documentation, then write that data to the
// Secure Content Store. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *WeatherReportHistoryService) Aodr(ctx context.Context, query WeatherReportHistoryAodrParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "udl/weatherreport/history/aodr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *WeatherReportHistoryService) Count(ctx context.Context, query WeatherReportHistoryCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/weatherreport/history/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// These services provide for posting and querying Weather Over Target information.
// The information contained within describes the current weather conditions over a
// target area or region to include navigational considerations such as altimeter
// settings, visibility, cloud heights etc.
type WeatherReportFull struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
	//
	// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
	// may include both real and simulated data.
	//
	// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
	// events, and analysis.
	//
	// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
	// datasets.
	//
	// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
	// requirements, and for validating technical, functional, and performance
	// characteristics.
	//
	// Any of "REAL", "TEST", "SIMULATED", "EXERCISE".
	DataMode WeatherReportFullDataMode `json:"dataMode,required"`
	// The central WGS-84 latitude of the weather report, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat,required"`
	// The central WGS-84 longitude of the weather report, in degrees. -180 to 180
	// degrees (negative values west of Prime Meridian).
	Lon float64 `json:"lon,required"`
	// Datetime when a weather observation was made or forecast was issued in ISO 8601
	// UTC datetime format with microsecond precision.
	ObTime time.Time `json:"obTime,required" format:"date-time"`
	// Identifies the type of weather report (e.g. OBSERVATION, FORECAST, etc.).
	ReportType string `json:"reportType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Describes the actual weather at position. Intended as, but not constrained to,
	// MIL-STD-6016 actual weather (e.g. NO STATEMENT, NO SIGNIFICANT WEATHER, DRIZZLE,
	// RAIN, SNOW, SNOW GRAINS, DIAMOND DUST, ICE PELLETS, HAIL, SMALL HAIL, MIST, FOG,
	// SMOKE, VOLCANIC ASH, WIDESPREAD DUST, SAND, HAZE, WELL DEVELOPED DUST, SQUALLS,
	// FUNNEL CLOUDS, SANDSTORM, DUSTSTORM, LOW CLOUDS, CLOUDY, GROUND FOG, DUST, HEAVY
	// RAIN, THUNDERSTORMS AWT, HEAVY THUNDERSTORMS, HURRICANE TYPHOON CYCLONE,
	// TROPICAL STORM, TORNADO, HIGH WINDS, LIGHTNING, FREEZING DRIZZLE, FREEZING RAIN,
	// HEAVY SNOW, ICING, SNOW OR RAIN AND SNOW MIXED, SHOWERS, CLEAR).
	ActWeather string `json:"actWeather"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt float64 `json:"alt"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// The measurement of air pressure in the atmosphere in kilopascals.
	BarPress float64 `json:"barPress"`
	// Flag indicating detection of a cloud-to-cloud lightning event.
	CcEvent bool `json:"ccEvent"`
	// Array of cloud cover descriptions - each element can be maximum of 16 characters
	// long. Intended as, but not constrained to, MIL-STD-6016 cloud cover designations
	// (e.g. SKY CLEAR, SCATTERED, BROKEN, OVERCAST, SKY OBSCURED). Each element of the
	// array corresponds to the elements in the cloudHght array specified respectively.
	CloudCover []string `json:"cloudCover"`
	// Array of cloud base heights in meters described by the cloudHght array. Each
	// element of the array corresponds to the elements in the cloudCover array
	// specified respectively.
	CloudHght []float64 `json:"cloudHght"`
	// Reports the lowest altitude at which contrails are occurring, in meters.
	ContrailHghtLower float64 `json:"contrailHghtLower"`
	// Reports the highest altitude at which contrails are occurring, in meters.
	ContrailHghtUpper float64 `json:"contrailHghtUpper"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Specific pressures or heights where measurements are taken, labeled as either
	// MANDATORY or SIGNIFICANT levels. Mandatory levels are at particular pressures at
	// geopotential heights. Significant levels are at particular geometric heights.
	DataLevel string `json:"dataLevel"`
	// The temperature at which air is saturated with water vapor, in degrees C.
	DewPoint float64 `json:"dewPoint"`
	// The amount of radiation that reaches earth's surface after being scattered by
	// the atmosphere, in Watts per square meter.
	DifRad float64 `json:"difRad"`
	// The difference in wind direction recorded over a period of time, in degrees.
	DirDev float64 `json:"dirDev"`
	// Describes the flight conditions in route to the target (NO STATEMENT, MAINLY
	// IFR, MAINLY VFR, THUNDERSTORMS).
	//
	// MAINLY IFR:&nbsp;&nbsp;Predominantly Instrument Flight Rules.
	//
	// MAINLY VFR:&nbsp;&nbsp;Predominantly Visual Flight Rules.
	//
	// THUNDERSTORMS:&nbsp;&nbsp;Thunderstorms expected in route.
	EnRouteWeather string `json:"enRouteWeather"`
	// Optional observation or forecast ID from external systems. This field has no
	// meaning within UDL and is provided as a convenience for systems that require
	// tracking of an internal system generated ID.
	ExternalID string `json:"externalId"`
	// Optional location ID from external systems. This field has no meaning within UDL
	// and is provided as a convenience for systems that require tracking of an
	// internal system generated ID.
	ExternalLocationID string `json:"externalLocationId"`
	// Valid end time of a weather forecast in ISO 8601 UTC datetime format with
	// millisecond precision.
	ForecastEndTime time.Time `json:"forecastEndTime" format:"date-time"`
	// Valid start time of a weather forecast in ISO 8601 UTC datetime format with
	// millisecond precision.
	ForecastStartTime time.Time `json:"forecastStartTime" format:"date-time"`
	// Altitude of a pressure surface in the atmosphere above mean sea level, in
	// meters.
	GeoPotentialAlt float64 `json:"geoPotentialAlt"`
	// The change in wind speed between two different lateral positions at a given
	// altitude divided by the horizontal distance between them, in units of 1/sec.
	Hshear float64 `json:"hshear"`
	// The International Civil Aviation Organization (ICAO) code of the airfield
	// associated with this weather report.
	Icao string `json:"icao"`
	// Reports the lowest altitude at which icing or freezing rain is occurring, in
	// meters.
	IcingLowerLimit float64 `json:"icingLowerLimit"`
	// Reports the highest altitude at which icing or freezing rain is occurring, in
	// meters.
	IcingUpperLimit float64 `json:"icingUpperLimit"`
	// Identifier of the Airfield associated with this weather report.
	IDAirfield string `json:"idAirfield"`
	// Identifier of the ground imagery associated for this weather over target report.
	IDGroundImagery string `json:"idGroundImagery"`
	// Unique identifier of the sensor making the weather measurement.
	IDSensor string `json:"idSensor"`
	// Identifier of the Site that is associated with this weather report.
	IDSite string `json:"idSite"`
	// An indication of how much the atmosphere refracts light.
	IndexRefraction float64 `json:"indexRefraction"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the record source. This may be an internal
	// identifier and not necessarily a valid sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The speed at which water is being applied to a specific area in millimeters per
	// hour.
	PrecipRate float64 `json:"precipRate"`
	// Altimeter set to read zero at mean sea level in kilopascals.
	Qnh float64 `json:"qnh"`
	// Average radial velocity of wind as measured by radar with multi-beam
	// configurations. Radial velocity is the component of wind velocity moving
	// directly toward or away from a sensor's radar beam, in meters per second. Values
	// can either be positive (wind is moving away from the radar) or negative (wind is
	// moving toward the radar).
	RadVel float64 `json:"radVel"`
	// Component of wind velocity moving directly toward or away from radar beam 1, in
	// meters per second. Radial velocity values can either be positive (wind is moving
	// away from the radar) or negative (wind is moving toward the radar). The beam
	// number designation is defined by the data source.
	RadVelBeam1 float64 `json:"radVelBeam1"`
	// Component of wind velocity moving directly toward or away from radar beam 2, in
	// meters per second. Radial velocity values can either be positive (wind is moving
	// away from the radar) or negative (wind is moving toward the radar). The beam
	// number designation is defined by the data source.
	RadVelBeam2 float64 `json:"radVelBeam2"`
	// Component of wind velocity moving directly toward or away from radar beam 3, in
	// meters per second. Radial velocity values can either be positive (wind is moving
	// away from the radar) or negative (wind is moving toward the radar). The beam
	// number designation is defined by the data source.
	RadVelBeam3 float64 `json:"radVelBeam3"`
	// Component of wind velocity moving directly toward or away from radar beam 4, in
	// meters per second. Radial velocity values can either be positive (wind is moving
	// away from the radar) or negative (wind is moving toward the radar). The beam
	// number designation is defined by the data source.
	RadVelBeam4 float64 `json:"radVelBeam4"`
	// Component of wind velocity moving directly toward or away from radar beam 5, in
	// meters per second. Radial velocity values can either be positive (wind is moving
	// away from the radar) or negative (wind is moving toward the radar). The beam
	// number designation is defined by the data source.
	RadVelBeam5 float64 `json:"radVelBeam5"`
	// The amount of rain that has fallen in the past hour, in centimeters.
	RainHour float64 `json:"rainHour"`
	// The Raw Meteorological Aerodrome Report (METAR) string.
	RawMetar string `json:"rawMETAR"`
	// Terminal Aerodrome Forecast (TAF) containing detailed weather predictions for a
	// specific airport or aerodrome.
	RawTaf string `json:"rawTAF"`
	// The amount of radiation that changes direction as a function of atmospheric
	// density, in Watts per square meter.
	RefRad float64 `json:"refRad"`
	// The percentage of water vapor in the atmosphere.
	RelHumidity float64 `json:"relHumidity"`
	// Sensor altitude at obTime in km. This includes pilot reports or other means of
	// weather observation.
	Senalt float64 `json:"senalt"`
	// Sensor WGS84 latitude at obTime in degrees. -90 to 90 degrees (negative values
	// south of equator). This includes pilot reports or other means of weather
	// observation.
	Senlat float64 `json:"senlat"`
	// Sensor WGS84 longitude at obTime in degrees. -180 to 180 degrees (negative
	// values west of Prime Meridian). This includes pilot reports or other means of
	// weather observation.
	Senlon float64 `json:"senlon"`
	// The volumetric percentage of soil water contained in a given volume of soil.
	SoilMoisture float64 `json:"soilMoisture"`
	// The measurement of soil temperature in degrees C.
	SoilTemp float64 `json:"soilTemp"`
	// The power per unit area received from the Sun in the form of electromagnetic
	// radiation as measured in the wavelength range of the measuring instrument. The
	// solar irradiance is measured in watt per square meter (W/m2).
	SolarRad float64 `json:"solarRad"`
	// Array of UUID(s) of the UDL data record(s) that are related to this
	// WeatherReport record. See the associated 'srcTyps' array for the specific types
	// of data, positionally corresponding to the UUIDs in this array. The 'srcTyps'
	// and 'srcIds' arrays must match in size. See the corresponding srcTyps array
	// element for the data type of the UUID and use the appropriate API operation to
	// retrieve that object.
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (SENSOR, WEATHERDATA) that are related to this
	// WeatherReport record. See the associated 'srcIds' array for the record UUIDs,
	// positionally corresponding to the record types in this array. The 'srcTyps' and
	// 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// Describes in which direction (if any) that better weather conditions exist.
	// Intended as, but not constrained to, MIL-STD-6016 surrounding weather
	// designations (e.g. NO STATEMENT, BETTER TO NORTH, BETTER TO EAST, BETTER TO
	// SOUTH, BETTER TO WEST).
	SurroundingWeather string `json:"surroundingWeather"`
	// The measurement of air temperature in degrees C.
	Temperature float64 `json:"temperature"`
	// Visual distance in meters.
	Visibility float64 `json:"visibility"`
	// The change in wind speed between two different altitudes divided by the vertical
	// distance between them, in units of 1/sec.
	Vshear float64 `json:"vshear"`
	// Amplifies the actual weather being reported. Intended as, but not constrained
	// to, MIL-STD-6016 weather amplification designations (e.g. NO STATEMENT, NO
	// SCATTERED BROKEN MEDIUM CLOUD, SCATTERED BROKEN MEDIUM CLOUDS, GUSTY WINDS AT
	// SERVICE, FOG IN VALLEYS, HIGHER TERRAIN OBSCURED, SURFACE CONDITIONS VARIABLE,
	// SURFACE WIND NE, SURFACE WIND SE, SURFACE WIND SW, SURFACE WIND NW, PRESENCE OF
	// CUMULONIMBUS).
	WeatherAmp string `json:"weatherAmp"`
	// Used in conjunction with actWeather and weatherInt. Intended as, but not
	// constrained to, MIL-STD-6016 actual weather descriptor (e.g. NO STATEMENT,
	// SHALLOW, PATCHES, LOW DRIFTING, BLOWING, SHOWERS, THUNDERSTORMS, SUPERCOOLED).
	WeatherDesc string `json:"weatherDesc"`
	// Identifier of the weather over target, which should remain the same on
	// subsequent Weather Over Target records.
	WeatherID string `json:"weatherId"`
	// Weather Intensity. Used in conjunction with actWeather and weatherDesc. Intended
	// as, but not constrained to, MIL-STD-6016 weather intensity (e.g. NO STATEMENT,
	// LIGHT, MODERATE, HEAVY, IN VICINITY).
	WeatherInt string `json:"weatherInt"`
	// The perceived temperature in degrees C.
	WindChill float64 `json:"windChill"`
	// Covariance matrix, in knots and second based units. The array values represent
	// the lower triangular half of the covariance matrix. The size of the covariance
	// matrix is 2x2. The covariance elements are position dependent within the array
	// with values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;3
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	WindCov []float64 `json:"windCov"`
	// Direction the wind is blowing, in degrees clockwise from true north.
	WindDir float64 `json:"windDir"`
	// Average wind direction over a 1 minute period, in degrees clockwise from true
	// north.
	WindDirAvg float64 `json:"windDirAvg"`
	// Wind direction corresponding to the peak wind speed during a 1 minute period, in
	// degrees clockwise from true north.
	WindDirPeak float64 `json:"windDirPeak"`
	// Wind direction corresponding to the peak wind speed during a 10 minute period,
	// in degrees clockwise from true north.
	WindDirPeak10 float64 `json:"windDirPeak10"`
	// Expresses the max gust speed of the wind, in meters/second.
	WindGust float64 `json:"windGust"`
	// Expresses the max gust speed of the wind recorded in a 10 minute period, in
	// meters/second.
	WindGust10 float64 `json:"windGust10"`
	// Expresses the speed of the wind in meters/second.
	WindSpd float64 `json:"windSpd"`
	// Average wind speed over a 1 minute period, in meters/second.
	WindSpdAvg float64 `json:"windSpdAvg"`
	// Boolean describing whether or not the wind direction and/or speed is variable.
	WindVar bool `json:"windVar"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		ObTime                respjson.Field
		ReportType            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		ActWeather            respjson.Field
		Agjson                respjson.Field
		Alt                   respjson.Field
		Andims                respjson.Field
		Area                  respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		BarPress              respjson.Field
		CcEvent               respjson.Field
		CloudCover            respjson.Field
		CloudHght             respjson.Field
		ContrailHghtLower     respjson.Field
		ContrailHghtUpper     respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		DataLevel             respjson.Field
		DewPoint              respjson.Field
		DifRad                respjson.Field
		DirDev                respjson.Field
		EnRouteWeather        respjson.Field
		ExternalID            respjson.Field
		ExternalLocationID    respjson.Field
		ForecastEndTime       respjson.Field
		ForecastStartTime     respjson.Field
		GeoPotentialAlt       respjson.Field
		Hshear                respjson.Field
		Icao                  respjson.Field
		IcingLowerLimit       respjson.Field
		IcingUpperLimit       respjson.Field
		IDAirfield            respjson.Field
		IDGroundImagery       respjson.Field
		IDSensor              respjson.Field
		IDSite                respjson.Field
		IndexRefraction       respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigSensorID          respjson.Field
		PrecipRate            respjson.Field
		Qnh                   respjson.Field
		RadVel                respjson.Field
		RadVelBeam1           respjson.Field
		RadVelBeam2           respjson.Field
		RadVelBeam3           respjson.Field
		RadVelBeam4           respjson.Field
		RadVelBeam5           respjson.Field
		RainHour              respjson.Field
		RawMetar              respjson.Field
		RawTaf                respjson.Field
		RefRad                respjson.Field
		RelHumidity           respjson.Field
		Senalt                respjson.Field
		Senlat                respjson.Field
		Senlon                respjson.Field
		SoilMoisture          respjson.Field
		SoilTemp              respjson.Field
		SolarRad              respjson.Field
		SrcIDs                respjson.Field
		SrcTyps               respjson.Field
		SurroundingWeather    respjson.Field
		Temperature           respjson.Field
		Visibility            respjson.Field
		Vshear                respjson.Field
		WeatherAmp            respjson.Field
		WeatherDesc           respjson.Field
		WeatherID             respjson.Field
		WeatherInt            respjson.Field
		WindChill             respjson.Field
		WindCov               respjson.Field
		WindDir               respjson.Field
		WindDirAvg            respjson.Field
		WindDirPeak           respjson.Field
		WindDirPeak10         respjson.Field
		WindGust              respjson.Field
		WindGust10            respjson.Field
		WindSpd               respjson.Field
		WindSpdAvg            respjson.Field
		WindVar               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeatherReportFull) RawJSON() string { return r.JSON.raw }
func (r *WeatherReportFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicator of whether the data is EXERCISE, REAL, SIMULATED, or TEST data:
//
// EXERCISE:&nbsp;Data pertaining to a government or military exercise. The data
// may include both real and simulated data.
//
// REAL:&nbsp;Data collected or produced that pertains to real-world objects,
// events, and analysis.
//
// SIMULATED:&nbsp;Synthetic data generated by a model to mimic real-world
// datasets.
//
// TEST:&nbsp;Specific datasets used to evaluate compliance with specifications and
// requirements, and for validating technical, functional, and performance
// characteristics.
type WeatherReportFullDataMode string

const (
	WeatherReportFullDataModeReal      WeatherReportFullDataMode = "REAL"
	WeatherReportFullDataModeTest      WeatherReportFullDataMode = "TEST"
	WeatherReportFullDataModeSimulated WeatherReportFullDataMode = "SIMULATED"
	WeatherReportFullDataModeExercise  WeatherReportFullDataMode = "EXERCISE"
)

type WeatherReportHistoryListParams struct {
	// Datetime when a weather observation was made or forecast was issued in ISO 8601
	// UTC datetime format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherReportHistoryListParams]'s query parameters as
// `url.Values`.
func (r WeatherReportHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherReportHistoryAodrParams struct {
	// Datetime when a weather observation was made or forecast was issued in ISO 8601
	// UTC datetime format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime time.Time `query:"obTime,required" format:"date-time" json:"-"`
	// optional, fields for retrieval. When omitted, ALL fields are assumed. See the
	// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on valid
	// query fields that can be selected.
	Columns     param.Opt[string] `query:"columns,omitzero" json:"-"`
	FirstResult param.Opt[int64]  `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64]  `query:"maxResults,omitzero" json:"-"`
	// optional, notification method for the created file link. When omitted, EMAIL is
	// assumed. Current valid values are: EMAIL, SMS.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// optional, field delimiter when the created file is not JSON. Must be a single
	// character chosen from this set: (',', ';', ':', '|'). When omitted, "," is used.
	// It is strongly encouraged that your field delimiter be a character unlikely to
	// occur within the data.
	OutputDelimiter param.Opt[string] `query:"outputDelimiter,omitzero" json:"-"`
	// optional, output format for the file. When omitted, JSON is assumed. Current
	// valid values are: JSON and CSV.
	OutputFormat param.Opt[string] `query:"outputFormat,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherReportHistoryAodrParams]'s query parameters as
// `url.Values`.
func (r WeatherReportHistoryAodrParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WeatherReportHistoryCountParams struct {
	// Datetime when a weather observation was made or forecast was issued in ISO 8601
	// UTC datetime format with microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	ObTime      time.Time        `query:"obTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WeatherReportHistoryCountParams]'s query parameters as
// `url.Values`.
func (r WeatherReportHistoryCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
