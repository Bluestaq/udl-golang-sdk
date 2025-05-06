// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apijson"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/apiquery"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/pagination"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/resp"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// SarObservationService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSarObservationService] method instead.
type SarObservationService struct {
	Options []option.RequestOption
	History SarObservationHistoryService
}

// NewSarObservationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSarObservationService(opts ...option.RequestOption) (r SarObservationService) {
	r = SarObservationService{}
	r.Options = opts
	r.History = NewSarObservationHistoryService(opts...)
	return
}

// Service operation to take a single SAR observation as a POST body and ingest
// into the database. This operation is not intended to be used for automated feeds
// into UDL. Data providers should contact the UDL team for specific role
// assignments and for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *SarObservationService) New(ctx context.Context, body SarObservationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sarobservation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SarObservationService) List(ctx context.Context, query SarObservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SarObservationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sarobservation"
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

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SarObservationService) ListAutoPaging(ctx context.Context, query SarObservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SarObservationListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SarObservationService) Count(ctx context.Context, query SarObservationCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sarobservation/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of SAR
// observations as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *SarObservationService) NewBulk(ctx context.Context, body SarObservationNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sarobservation/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single SAR observations by its unique ID passed as a
// path parameter.
func (r *SarObservationService) Get(ctx context.Context, id string, query SarObservationGetParams, opts ...option.RequestOption) (res *SarObservationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/sarobservation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SarObservationService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sarobservation/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Service operation to dynamically query data and only return specified
// columns/fields. Requested columns are specified by the 'columns' query parameter
// and should be a comma separated list of valid fields for the specified data
// type. classificationMarking is always returned. See the queryhelp operation
// (/udl/<datatype>/queryhelp) for more details on valid/required query parameter
// information. An example URI: /udl/elset/tuple?columns=satNo,period&epoch=>now-5
// hours would return the satNo and period of elsets with an epoch greater than 5
// hours ago.
func (r *SarObservationService) Tuple(ctx context.Context, query SarObservationTupleParams, opts ...option.RequestOption) (res *[]SarObservationTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sarobservation/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take SAR observations as a POST body and ingest into the
// database. This operation is intended to be used for automated feeds into UDL. A
// specific role is required to perform this service operation. Please contact the
// UDL team for assistance.
func (r *SarObservationService) UnvalidatedPublish(ctx context.Context, body SarObservationUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Model representation of observation data for SAR based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type SarObservationListResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode SarObservationListResponseDataMode `json:"dataMode,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype string `json:"atype"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle float64 `json:"azimuthAngle"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime time.Time `json:"centerTime" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID string `json:"collectionId"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle float64 `json:"continuousSpotAngle"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys string `json:"coordSys"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd time.Time `json:"detectionEnd" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID string `json:"detectionId"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart time.Time `json:"detectionStart" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime float64 `json:"dwellTime"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange float64 `json:"farRange"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle float64 `json:"grazeAngle"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection float64 `json:"groundResolutionProjection"`
	// Unique identifier of the spacecraft hosting the sensor associated with this
	// collection.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The center incidence angle in degrees.
	IncidenceAngle float64 `json:"incidenceAngle"`
	// The number of looks in the azimuth direction.
	LooksAzimuth int64 `json:"looksAzimuth"`
	// The number of looks in the range direction.
	LooksRange int64 `json:"looksRange"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber float64 `json:"multilookNumber"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange float64 `json:"nearRange"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection string `json:"obDirection"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand string `json:"operatingBand"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq float64 `json:"operatingFreq"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState string `json:"orbitState"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth float64 `json:"pulseBandwidth"`
	// The duration of a pulse in seconds.
	PulseDuration float64 `json:"pulseDuration"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth float64 `json:"resolutionAzimuth"`
	// The resolution in the range direction measured in meters.
	ResolutionRange float64 `json:"resolutionRange"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization string `json:"rxPolarization"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo int64 `json:"satNo"`
	// Sensor altitude during collection in kilometers.
	Senalt float64 `json:"senalt"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd float64 `json:"senlatEnd"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart float64 `json:"senlatStart"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd float64 `json:"senlonEnd"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart float64 `json:"senlonStart"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx float64 `json:"senvelx"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely float64 `json:"senvely"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz float64 `json:"senvelz"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange float64 `json:"slantRange"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth float64 `json:"spacingAzimuth"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange float64 `json:"spacingRange"`
	// The squint angle for the collection in degrees.
	SquintAngle float64 `json:"squintAngle"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength float64 `json:"swathLength"`
	// The collection target X position in kilometers.
	Targetposx float64 `json:"targetposx"`
	// The collection target Y position in kilometers.
	Targetposy float64 `json:"targetposy"`
	// The collection target Z position in kilometers.
	Targetposz float64 `json:"targetposz"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization string `json:"txPolarization"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		CollectionEnd              resp.Field
		CollectionStart            resp.Field
		DataMode                   resp.Field
		SarMode                    resp.Field
		Source                     resp.Field
		ID                         resp.Field
		Agjson                     resp.Field
		Andims                     resp.Field
		Asrid                      resp.Field
		Atext                      resp.Field
		Atype                      resp.Field
		AzimuthAngle               resp.Field
		CenterTime                 resp.Field
		CollectionID               resp.Field
		ContinuousSpotAngle        resp.Field
		CoordSys                   resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		DetectionEnd               resp.Field
		DetectionID                resp.Field
		DetectionStart             resp.Field
		DwellTime                  resp.Field
		ExternalID                 resp.Field
		FarRange                   resp.Field
		GrazeAngle                 resp.Field
		GroundResolutionProjection resp.Field
		IDOnOrbit                  resp.Field
		IDSensor                   resp.Field
		IncidenceAngle             resp.Field
		LooksAzimuth               resp.Field
		LooksRange                 resp.Field
		MultilookNumber            resp.Field
		NearRange                  resp.Field
		ObDirection                resp.Field
		OperatingBand              resp.Field
		OperatingFreq              resp.Field
		OrbitState                 resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		OrigObjectID               resp.Field
		OrigSensorID               resp.Field
		PulseBandwidth             resp.Field
		PulseDuration              resp.Field
		ResolutionAzimuth          resp.Field
		ResolutionRange            resp.Field
		RxPolarization             resp.Field
		SatNo                      resp.Field
		Senalt                     resp.Field
		SenlatEnd                  resp.Field
		SenlatStart                resp.Field
		SenlonEnd                  resp.Field
		SenlonStart                resp.Field
		Senvelx                    resp.Field
		Senvely                    resp.Field
		Senvelz                    resp.Field
		SlantRange                 resp.Field
		Snr                        resp.Field
		SourceDl                   resp.Field
		SpacingAzimuth             resp.Field
		SpacingRange               resp.Field
		SquintAngle                resp.Field
		SrcIDs                     resp.Field
		SrcTyps                    resp.Field
		SwathLength                resp.Field
		Targetposx                 resp.Field
		Targetposy                 resp.Field
		Targetposz                 resp.Field
		TransactionID              resp.Field
		TxPolarization             resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SarObservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SarObservationListResponse) UnmarshalJSON(data []byte) error {
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
type SarObservationListResponseDataMode string

const (
	SarObservationListResponseDataModeReal      SarObservationListResponseDataMode = "REAL"
	SarObservationListResponseDataModeTest      SarObservationListResponseDataMode = "TEST"
	SarObservationListResponseDataModeSimulated SarObservationListResponseDataMode = "SIMULATED"
	SarObservationListResponseDataModeExercise  SarObservationListResponseDataMode = "EXERCISE"
)

// Model representation of observation data for SAR based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type SarObservationGetResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode SarObservationGetResponseDataMode `json:"dataMode,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the image event as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype string `json:"atype"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle float64 `json:"azimuthAngle"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime time.Time `json:"centerTime" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID string `json:"collectionId"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle float64 `json:"continuousSpotAngle"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys string `json:"coordSys"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd time.Time `json:"detectionEnd" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID string `json:"detectionId"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart time.Time `json:"detectionStart" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime float64 `json:"dwellTime"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange float64 `json:"farRange"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle float64 `json:"grazeAngle"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection float64 `json:"groundResolutionProjection"`
	// Unique identifier of the spacecraft hosting the sensor associated with this
	// collection.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The center incidence angle in degrees.
	IncidenceAngle float64 `json:"incidenceAngle"`
	// The number of looks in the azimuth direction.
	LooksAzimuth int64 `json:"looksAzimuth"`
	// The number of looks in the range direction.
	LooksRange int64 `json:"looksRange"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber float64 `json:"multilookNumber"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange float64 `json:"nearRange"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection string `json:"obDirection"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand string `json:"operatingBand"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq float64 `json:"operatingFreq"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState string `json:"orbitState"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth float64 `json:"pulseBandwidth"`
	// The duration of a pulse in seconds.
	PulseDuration float64 `json:"pulseDuration"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth float64 `json:"resolutionAzimuth"`
	// The resolution in the range direction measured in meters.
	ResolutionRange float64 `json:"resolutionRange"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization string `json:"rxPolarization"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo int64 `json:"satNo"`
	// Sensor altitude during collection in kilometers.
	Senalt float64 `json:"senalt"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd float64 `json:"senlatEnd"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart float64 `json:"senlatStart"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd float64 `json:"senlonEnd"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart float64 `json:"senlonStart"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx float64 `json:"senvelx"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely float64 `json:"senvely"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz float64 `json:"senvelz"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange float64 `json:"slantRange"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth float64 `json:"spacingAzimuth"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange float64 `json:"spacingRange"`
	// The squint angle for the collection in degrees.
	SquintAngle float64 `json:"squintAngle"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength float64 `json:"swathLength"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The collection target X position in kilometers.
	Targetposx float64 `json:"targetposx"`
	// The collection target Y position in kilometers.
	Targetposy float64 `json:"targetposy"`
	// The collection target Z position in kilometers.
	Targetposz float64 `json:"targetposz"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization string `json:"txPolarization"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		CollectionEnd              resp.Field
		CollectionStart            resp.Field
		DataMode                   resp.Field
		SarMode                    resp.Field
		Source                     resp.Field
		ID                         resp.Field
		Agjson                     resp.Field
		Andims                     resp.Field
		Area                       resp.Field
		Asrid                      resp.Field
		Atext                      resp.Field
		Atype                      resp.Field
		AzimuthAngle               resp.Field
		CenterTime                 resp.Field
		CollectionID               resp.Field
		ContinuousSpotAngle        resp.Field
		CoordSys                   resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		DetectionEnd               resp.Field
		DetectionID                resp.Field
		DetectionStart             resp.Field
		DwellTime                  resp.Field
		ExternalID                 resp.Field
		FarRange                   resp.Field
		GrazeAngle                 resp.Field
		GroundResolutionProjection resp.Field
		IDOnOrbit                  resp.Field
		IDSensor                   resp.Field
		IncidenceAngle             resp.Field
		LooksAzimuth               resp.Field
		LooksRange                 resp.Field
		MultilookNumber            resp.Field
		NearRange                  resp.Field
		ObDirection                resp.Field
		OnOrbit                    resp.Field
		OperatingBand              resp.Field
		OperatingFreq              resp.Field
		OrbitState                 resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		OrigObjectID               resp.Field
		OrigSensorID               resp.Field
		PulseBandwidth             resp.Field
		PulseDuration              resp.Field
		ResolutionAzimuth          resp.Field
		ResolutionRange            resp.Field
		RxPolarization             resp.Field
		SatNo                      resp.Field
		Senalt                     resp.Field
		SenlatEnd                  resp.Field
		SenlatStart                resp.Field
		SenlonEnd                  resp.Field
		SenlonStart                resp.Field
		Senvelx                    resp.Field
		Senvely                    resp.Field
		Senvelz                    resp.Field
		SlantRange                 resp.Field
		Snr                        resp.Field
		SourceDl                   resp.Field
		SpacingAzimuth             resp.Field
		SpacingRange               resp.Field
		SquintAngle                resp.Field
		SrcIDs                     resp.Field
		SrcTyps                    resp.Field
		SwathLength                resp.Field
		Tags                       resp.Field
		Targetposx                 resp.Field
		Targetposy                 resp.Field
		Targetposz                 resp.Field
		TransactionID              resp.Field
		TxPolarization             resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SarObservationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SarObservationGetResponse) UnmarshalJSON(data []byte) error {
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
type SarObservationGetResponseDataMode string

const (
	SarObservationGetResponseDataModeReal      SarObservationGetResponseDataMode = "REAL"
	SarObservationGetResponseDataModeTest      SarObservationGetResponseDataMode = "TEST"
	SarObservationGetResponseDataModeSimulated SarObservationGetResponseDataMode = "SIMULATED"
	SarObservationGetResponseDataModeExercise  SarObservationGetResponseDataMode = "EXERCISE"
)

// Model representation of observation data for SAR based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
type SarObservationTupleResponse struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode SarObservationTupleResponseDataMode `json:"dataMode,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
	// Number of dimensions of the geometry depicted by region.
	Andims int64 `json:"andims"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the image event as projected on the ground.
	Area string `json:"area"`
	// Geographical spatial_ref_sys for region.
	Asrid int64 `json:"asrid"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext string `json:"atext"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype string `json:"atype"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle float64 `json:"azimuthAngle"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime time.Time `json:"centerTime" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID string `json:"collectionId"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle float64 `json:"continuousSpotAngle"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys string `json:"coordSys"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd time.Time `json:"detectionEnd" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID string `json:"detectionId"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart time.Time `json:"detectionStart" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime float64 `json:"dwellTime"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID string `json:"externalId"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange float64 `json:"farRange"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle float64 `json:"grazeAngle"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection float64 `json:"groundResolutionProjection"`
	// Unique identifier of the spacecraft hosting the sensor associated with this
	// collection.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the reporting sensor.
	IDSensor string `json:"idSensor"`
	// The center incidence angle in degrees.
	IncidenceAngle float64 `json:"incidenceAngle"`
	// The number of looks in the azimuth direction.
	LooksAzimuth int64 `json:"looksAzimuth"`
	// The number of looks in the range direction.
	LooksRange int64 `json:"looksRange"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber float64 `json:"multilookNumber"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange float64 `json:"nearRange"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection string `json:"obDirection"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand string `json:"operatingBand"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq float64 `json:"operatingFreq"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState string `json:"orbitState"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID string `json:"origSensorId"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth float64 `json:"pulseBandwidth"`
	// The duration of a pulse in seconds.
	PulseDuration float64 `json:"pulseDuration"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth float64 `json:"resolutionAzimuth"`
	// The resolution in the range direction measured in meters.
	ResolutionRange float64 `json:"resolutionRange"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization string `json:"rxPolarization"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo int64 `json:"satNo"`
	// Sensor altitude during collection in kilometers.
	Senalt float64 `json:"senalt"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd float64 `json:"senlatEnd"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart float64 `json:"senlatStart"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd float64 `json:"senlonEnd"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart float64 `json:"senlonStart"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx float64 `json:"senvelx"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely float64 `json:"senvely"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz float64 `json:"senvelz"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange float64 `json:"slantRange"`
	// Signal to noise ratio, in dB.
	Snr float64 `json:"snr"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth float64 `json:"spacingAzimuth"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange float64 `json:"spacingRange"`
	// The squint angle for the collection in degrees.
	SquintAngle float64 `json:"squintAngle"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength float64 `json:"swathLength"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// The collection target X position in kilometers.
	Targetposx float64 `json:"targetposx"`
	// The collection target Y position in kilometers.
	Targetposy float64 `json:"targetposy"`
	// The collection target Z position in kilometers.
	Targetposz float64 `json:"targetposz"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization string `json:"txPolarization"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		CollectionEnd              resp.Field
		CollectionStart            resp.Field
		DataMode                   resp.Field
		SarMode                    resp.Field
		Source                     resp.Field
		ID                         resp.Field
		Agjson                     resp.Field
		Andims                     resp.Field
		Area                       resp.Field
		Asrid                      resp.Field
		Atext                      resp.Field
		Atype                      resp.Field
		AzimuthAngle               resp.Field
		CenterTime                 resp.Field
		CollectionID               resp.Field
		ContinuousSpotAngle        resp.Field
		CoordSys                   resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		DetectionEnd               resp.Field
		DetectionID                resp.Field
		DetectionStart             resp.Field
		DwellTime                  resp.Field
		ExternalID                 resp.Field
		FarRange                   resp.Field
		GrazeAngle                 resp.Field
		GroundResolutionProjection resp.Field
		IDOnOrbit                  resp.Field
		IDSensor                   resp.Field
		IncidenceAngle             resp.Field
		LooksAzimuth               resp.Field
		LooksRange                 resp.Field
		MultilookNumber            resp.Field
		NearRange                  resp.Field
		ObDirection                resp.Field
		OnOrbit                    resp.Field
		OperatingBand              resp.Field
		OperatingFreq              resp.Field
		OrbitState                 resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		OrigObjectID               resp.Field
		OrigSensorID               resp.Field
		PulseBandwidth             resp.Field
		PulseDuration              resp.Field
		ResolutionAzimuth          resp.Field
		ResolutionRange            resp.Field
		RxPolarization             resp.Field
		SatNo                      resp.Field
		Senalt                     resp.Field
		SenlatEnd                  resp.Field
		SenlatStart                resp.Field
		SenlonEnd                  resp.Field
		SenlonStart                resp.Field
		Senvelx                    resp.Field
		Senvely                    resp.Field
		Senvelz                    resp.Field
		SlantRange                 resp.Field
		Snr                        resp.Field
		SourceDl                   resp.Field
		SpacingAzimuth             resp.Field
		SpacingRange               resp.Field
		SquintAngle                resp.Field
		SrcIDs                     resp.Field
		SrcTyps                    resp.Field
		SwathLength                resp.Field
		Tags                       resp.Field
		Targetposx                 resp.Field
		Targetposy                 resp.Field
		Targetposz                 resp.Field
		TransactionID              resp.Field
		TxPolarization             resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SarObservationTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SarObservationTupleResponse) UnmarshalJSON(data []byte) error {
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
type SarObservationTupleResponseDataMode string

const (
	SarObservationTupleResponseDataModeReal      SarObservationTupleResponseDataMode = "REAL"
	SarObservationTupleResponseDataModeTest      SarObservationTupleResponseDataMode = "TEST"
	SarObservationTupleResponseDataModeSimulated SarObservationTupleResponseDataMode = "SIMULATED"
	SarObservationTupleResponseDataModeExercise  SarObservationTupleResponseDataMode = "EXERCISE"
)

type SarObservationNewParams struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode SarObservationNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the image event as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle param.Opt[float64] `json:"azimuthAngle,omitzero"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime param.Opt[time.Time] `json:"centerTime,omitzero" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle param.Opt[float64] `json:"continuousSpotAngle,omitzero"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys param.Opt[string] `json:"coordSys,omitzero"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd param.Opt[time.Time] `json:"detectionEnd,omitzero" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID param.Opt[string] `json:"detectionId,omitzero"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart param.Opt[time.Time] `json:"detectionStart,omitzero" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime param.Opt[float64] `json:"dwellTime,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange param.Opt[float64] `json:"farRange,omitzero"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle param.Opt[float64] `json:"grazeAngle,omitzero"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection param.Opt[float64] `json:"groundResolutionProjection,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The center incidence angle in degrees.
	IncidenceAngle param.Opt[float64] `json:"incidenceAngle,omitzero"`
	// The number of looks in the azimuth direction.
	LooksAzimuth param.Opt[int64] `json:"looksAzimuth,omitzero"`
	// The number of looks in the range direction.
	LooksRange param.Opt[int64] `json:"looksRange,omitzero"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber param.Opt[float64] `json:"multilookNumber,omitzero"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange param.Opt[float64] `json:"nearRange,omitzero"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection param.Opt[string] `json:"obDirection,omitzero"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand param.Opt[string] `json:"operatingBand,omitzero"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq param.Opt[float64] `json:"operatingFreq,omitzero"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState param.Opt[string] `json:"orbitState,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// The duration of a pulse in seconds.
	PulseDuration param.Opt[float64] `json:"pulseDuration,omitzero"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth param.Opt[float64] `json:"resolutionAzimuth,omitzero"`
	// The resolution in the range direction measured in meters.
	ResolutionRange param.Opt[float64] `json:"resolutionRange,omitzero"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization param.Opt[string] `json:"rxPolarization,omitzero"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude during collection in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd param.Opt[float64] `json:"senlatEnd,omitzero"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart param.Opt[float64] `json:"senlatStart,omitzero"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd param.Opt[float64] `json:"senlonEnd,omitzero"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart param.Opt[float64] `json:"senlonStart,omitzero"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange param.Opt[float64] `json:"slantRange,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth param.Opt[float64] `json:"spacingAzimuth,omitzero"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange param.Opt[float64] `json:"spacingRange,omitzero"`
	// The squint angle for the collection in degrees.
	SquintAngle param.Opt[float64] `json:"squintAngle,omitzero"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength param.Opt[float64] `json:"swathLength,omitzero"`
	// The collection target X position in kilometers.
	Targetposx param.Opt[float64] `json:"targetposx,omitzero"`
	// The collection target Y position in kilometers.
	Targetposy param.Opt[float64] `json:"targetposy,omitzero"`
	// The collection target Z position in kilometers.
	Targetposz param.Opt[float64] `json:"targetposz,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization param.Opt[string] `json:"txPolarization,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SarObservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SarObservationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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
type SarObservationNewParamsDataMode string

const (
	SarObservationNewParamsDataModeReal      SarObservationNewParamsDataMode = "REAL"
	SarObservationNewParamsDataModeTest      SarObservationNewParamsDataMode = "TEST"
	SarObservationNewParamsDataModeSimulated SarObservationNewParamsDataMode = "SIMULATED"
	SarObservationNewParamsDataModeExercise  SarObservationNewParamsDataMode = "EXERCISE"
)

type SarObservationListParams struct {
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	CollectionStart time.Time        `query:"collectionStart,required" format:"date-time" json:"-"`
	FirstResult     param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults      param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SarObservationListParams]'s query parameters as
// `url.Values`.
func (r SarObservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SarObservationCountParams struct {
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	CollectionStart time.Time        `query:"collectionStart,required" format:"date-time" json:"-"`
	FirstResult     param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults      param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SarObservationCountParams]'s query parameters as
// `url.Values`.
func (r SarObservationCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SarObservationNewBulkParams struct {
	Body []SarObservationNewBulkParamsBody
	paramObj
}

func (r SarObservationNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of observation data for SAR based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
//
// The properties ClassificationMarking, CollectionEnd, CollectionStart, DataMode,
// SarMode, Source are required.
type SarObservationNewBulkParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the image event as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle param.Opt[float64] `json:"azimuthAngle,omitzero"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime param.Opt[time.Time] `json:"centerTime,omitzero" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle param.Opt[float64] `json:"continuousSpotAngle,omitzero"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys param.Opt[string] `json:"coordSys,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd param.Opt[time.Time] `json:"detectionEnd,omitzero" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID param.Opt[string] `json:"detectionId,omitzero"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart param.Opt[time.Time] `json:"detectionStart,omitzero" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime param.Opt[float64] `json:"dwellTime,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange param.Opt[float64] `json:"farRange,omitzero"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle param.Opt[float64] `json:"grazeAngle,omitzero"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection param.Opt[float64] `json:"groundResolutionProjection,omitzero"`
	// Unique identifier of the spacecraft hosting the sensor associated with this
	// collection.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The center incidence angle in degrees.
	IncidenceAngle param.Opt[float64] `json:"incidenceAngle,omitzero"`
	// The number of looks in the azimuth direction.
	LooksAzimuth param.Opt[int64] `json:"looksAzimuth,omitzero"`
	// The number of looks in the range direction.
	LooksRange param.Opt[int64] `json:"looksRange,omitzero"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber param.Opt[float64] `json:"multilookNumber,omitzero"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange param.Opt[float64] `json:"nearRange,omitzero"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection param.Opt[string] `json:"obDirection,omitzero"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand param.Opt[string] `json:"operatingBand,omitzero"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq param.Opt[float64] `json:"operatingFreq,omitzero"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState param.Opt[string] `json:"orbitState,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// The duration of a pulse in seconds.
	PulseDuration param.Opt[float64] `json:"pulseDuration,omitzero"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth param.Opt[float64] `json:"resolutionAzimuth,omitzero"`
	// The resolution in the range direction measured in meters.
	ResolutionRange param.Opt[float64] `json:"resolutionRange,omitzero"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization param.Opt[string] `json:"rxPolarization,omitzero"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude during collection in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd param.Opt[float64] `json:"senlatEnd,omitzero"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart param.Opt[float64] `json:"senlatStart,omitzero"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd param.Opt[float64] `json:"senlonEnd,omitzero"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart param.Opt[float64] `json:"senlonStart,omitzero"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange param.Opt[float64] `json:"slantRange,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth param.Opt[float64] `json:"spacingAzimuth,omitzero"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange param.Opt[float64] `json:"spacingRange,omitzero"`
	// The squint angle for the collection in degrees.
	SquintAngle param.Opt[float64] `json:"squintAngle,omitzero"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength param.Opt[float64] `json:"swathLength,omitzero"`
	// The collection target X position in kilometers.
	Targetposx param.Opt[float64] `json:"targetposx,omitzero"`
	// The collection target Y position in kilometers.
	Targetposy param.Opt[float64] `json:"targetposy,omitzero"`
	// The collection target Z position in kilometers.
	Targetposz param.Opt[float64] `json:"targetposz,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization param.Opt[string] `json:"txPolarization,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SarObservationNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SarObservationNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SarObservationNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SarObservationGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SarObservationGetParams]'s query parameters as
// `url.Values`.
func (r SarObservationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SarObservationTupleParams struct {
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	CollectionStart time.Time `query:"collectionStart,required" format:"date-time" json:"-"`
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SarObservationTupleParams]'s query parameters as
// `url.Values`.
func (r SarObservationTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SarObservationUnvalidatedPublishParams struct {
	Body []SarObservationUnvalidatedPublishParamsBody
	paramObj
}

func (r SarObservationUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Model representation of observation data for SAR based sensor phenomenologies.
// J2000 is the preferred coordinate frame for all observations, but in some cases
// observations may be in another frame depending on the provider. Please see the
// 'Discover' tab in the storefront to confirm coordinate frames by data provider.
//
// The properties ClassificationMarking, CollectionEnd, CollectionStart, DataMode,
// SarMode, Source are required.
type SarObservationUnvalidatedPublishParamsBody struct {
	// Classification marking of the data in IC/CAPCO Portion-marked format.
	ClassificationMarking string `json:"classificationMarking,required"`
	// Collection end time in ISO 8601 UTC format with microsecond precision.
	CollectionEnd time.Time `json:"collectionEnd,required" format:"date-time"`
	// Collection start time in ISO 8601 UTC format with microsecond precision.
	CollectionStart time.Time `json:"collectionStart,required" format:"date-time"`
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
	DataMode string `json:"dataMode,omitzero,required"`
	// Collection mode setting for this collection (e.g. AREA, SPOTLIGHT, STRIP, etc.).
	SarMode string `json:"sarMode,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the image event as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground (POLYGON, POINT, LINE).
	Atype param.Opt[string] `json:"atype,omitzero"`
	// The azimuth angle, in degrees, of the SAR satellite nadir subpoint measured
	// clockwise from true north at the subpoint.
	AzimuthAngle param.Opt[float64] `json:"azimuthAngle,omitzero"`
	// The datetime at the center point of the collection in ISO 8601 UTC format with
	// microsecond precision.
	CenterTime param.Opt[time.Time] `json:"centerTime,omitzero" format:"date-time"`
	// Optional identifier to indicate the specific collection tasking which produced
	// this observation.
	CollectionID param.Opt[string] `json:"collectionId,omitzero"`
	// Required sweep angle for the continuous spot scene in degrees.
	ContinuousSpotAngle param.Opt[float64] `json:"continuousSpotAngle,omitzero"`
	// The coordinate system used for the sensor velocity and target position vectors
	// for the collection.
	CoordSys param.Opt[string] `json:"coordSys,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The detection end time in ISO 8601 UTC format, with microsecond precision.
	DetectionEnd param.Opt[time.Time] `json:"detectionEnd,omitzero" format:"date-time"`
	// Identifier of the specific detection within a collection which produced this
	// observation.
	DetectionID param.Opt[string] `json:"detectionId,omitzero"`
	// The detection start time in ISO 8601 UTC format, with microsecond precision.
	DetectionStart param.Opt[time.Time] `json:"detectionStart,omitzero" format:"date-time"`
	// The duration, in seconds, of this detection.
	DwellTime param.Opt[float64] `json:"dwellTime,omitzero"`
	// Optional ID from external systems. This field has no meaning within UDL and is
	// provided as a convenience for systems that require tracking of an internal
	// system generated ID.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Specifies the farthest range, in kilometers, from the flight path to target
	// during the collection.
	FarRange param.Opt[float64] `json:"farRange,omitzero"`
	// The graze angle (also referred to as look angle) for the collection in degrees.
	GrazeAngle param.Opt[float64] `json:"grazeAngle,omitzero"`
	// Distance between independent measurements, representing the physical dimension
	// that represents a pixel of the image.
	GroundResolutionProjection param.Opt[float64] `json:"groundResolutionProjection,omitzero"`
	// Unique identifier of the spacecraft hosting the sensor associated with this
	// collection.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the reporting sensor.
	IDSensor param.Opt[string] `json:"idSensor,omitzero"`
	// The center incidence angle in degrees.
	IncidenceAngle param.Opt[float64] `json:"incidenceAngle,omitzero"`
	// The number of looks in the azimuth direction.
	LooksAzimuth param.Opt[int64] `json:"looksAzimuth,omitzero"`
	// The number of looks in the range direction.
	LooksRange param.Opt[int64] `json:"looksRange,omitzero"`
	// Averages the input synthetic aperture radar (SAR) data by looks in range and
	// azimuth to approximate square pixels, mitigates speckle, and reduces SAR tool
	// processing time.
	MultilookNumber param.Opt[float64] `json:"multilookNumber,omitzero"`
	// Specifies the closest range, in kilometers, from the flight path to target
	// during the collection.
	NearRange param.Opt[float64] `json:"nearRange,omitzero"`
	// The antenna pointing direction (LEFT, RIGHT).
	ObDirection param.Opt[string] `json:"obDirection,omitzero"`
	// Name of the band containing operating frequency for the collection (e.g. C, E,
	// EHF, HF, K, Ka, Ku, L, Q, S, SHF, UNK, UHF, V, VHF, VLF, W, X). See RFBandType
	// for more details and descriptions of each band name.
	OperatingBand param.Opt[string] `json:"operatingBand,omitzero"`
	// The operating frequency, in Mhz, for the collection.
	OperatingFreq param.Opt[float64] `json:"operatingFreq,omitzero"`
	// The orbital direction (ASCENDING, DESCENDING) of the platform during the
	// collection.
	OrbitState param.Opt[string] `json:"orbitState,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by the source to indicate the onorbit object
	// hosting the sensor associated with this collection. This may be an internal
	// identifier and not necessarily a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Optional identifier provided by the source to indicate the sensor for this
	// collection. This may be an internal identifier and not necessarily a valid
	// sensor ID.
	OrigSensorID param.Opt[string] `json:"origSensorId,omitzero"`
	// The bandwidth frequency of the pulse in Mhz.
	PulseBandwidth param.Opt[float64] `json:"pulseBandwidth,omitzero"`
	// The duration of a pulse in seconds.
	PulseDuration param.Opt[float64] `json:"pulseDuration,omitzero"`
	// The resolution in the azimuth direction measured in meters.
	ResolutionAzimuth param.Opt[float64] `json:"resolutionAzimuth,omitzero"`
	// The resolution in the range direction measured in meters.
	ResolutionRange param.Opt[float64] `json:"resolutionRange,omitzero"`
	// Receive polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	RxPolarization param.Opt[string] `json:"rxPolarization,omitzero"`
	// Satellite/Catalog number of the spacecraft hosting the sensor associated with
	// this collection.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Sensor altitude during collection in kilometers.
	Senalt param.Opt[float64] `json:"senalt,omitzero"`
	// WGS-84 sensor latitude sub-point at collect end time (collectionEnd),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatEnd param.Opt[float64] `json:"senlatEnd,omitzero"`
	// WGS-84 sensor latitude sub-point at collect start time (collectionStart),
	// represented as -90 to 90 degrees (negative values south of equator).
	SenlatStart param.Opt[float64] `json:"senlatStart,omitzero"`
	// WGS-84 sensor longitude sub-point at collect end time (collectionEnd),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonEnd param.Opt[float64] `json:"senlonEnd,omitzero"`
	// WGS-84 sensor longitude sub-point at collect start time (collectionStart),
	// represented as -180 to 180 degrees (negative values west of Prime Meridian).
	SenlonStart param.Opt[float64] `json:"senlonStart,omitzero"`
	// Sensor platform X-velocity during collection in kilometers/second.
	Senvelx param.Opt[float64] `json:"senvelx,omitzero"`
	// Sensor platform Y-velocity during collection in kilometers/second.
	Senvely param.Opt[float64] `json:"senvely,omitzero"`
	// Sensor platform Z-velocity during collection in kilometers/second.
	Senvelz param.Opt[float64] `json:"senvelz,omitzero"`
	// Slant distance from sensor to center point of imaging event in kilometers.
	SlantRange param.Opt[float64] `json:"slantRange,omitzero"`
	// Signal to noise ratio, in dB.
	Snr param.Opt[float64] `json:"snr,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The pixel spacing in the azimuth direction measured in meters.
	SpacingAzimuth param.Opt[float64] `json:"spacingAzimuth,omitzero"`
	// The pixel spacing in the range direction measured in meters.
	SpacingRange param.Opt[float64] `json:"spacingRange,omitzero"`
	// The squint angle for the collection in degrees.
	SquintAngle param.Opt[float64] `json:"squintAngle,omitzero"`
	// The length of the collection as projected on the ground in kilometers.
	SwathLength param.Opt[float64] `json:"swathLength,omitzero"`
	// The collection target X position in kilometers.
	Targetposx param.Opt[float64] `json:"targetposx,omitzero"`
	// The collection target Y position in kilometers.
	Targetposy param.Opt[float64] `json:"targetposy,omitzero"`
	// The collection target Z position in kilometers.
	Targetposz param.Opt[float64] `json:"targetposz,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Transmit polarization e.g. H - (Horizontally Polarized) Perpendicular to Earth's
	// surface, V - (Vertically Polarized) Parallel to Earth's surface, L - (Left Hand
	// Circularly Polarized) Rotating left relative to the earth's surface, R - (Right
	// Hand Circularly Polarized) Rotating right relative to the earth's surface.
	TxPolarization param.Opt[string] `json:"txPolarization,omitzero"`
	// Array of UUIDs of the UDL data records that are related to the SAR Observation.
	// See the associated 'srcTyps' array for the specific types of data, positionally
	// corresponding to the UUIDs in this array. The 'srcTyps' and 'srcIds' arrays must
	// match in size. See the corresponding srcTyps array element for the data type of
	// the UUID and use the appropriate API operation to retrieve that object (e.g.
	// /udl/sarobservation/{uuid}).
	SrcIDs []string `json:"srcIds,omitzero"`
	// Array of UDL record types (e.g. ANALYTICMAGERY, ESID, GROUNDIMAGE, NOTIFICATION,
	// POI, SV, TRACK) that are related to the SAR Observation. See the associated
	// 'srcIds' array for the record UUIDs, positionally corresponding to the record
	// types in this array. The 'srcTyps' and 'srcIds' arrays must match in size.
	SrcTyps []string `json:"srcTyps,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SarObservationUnvalidatedPublishParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SarObservationUnvalidatedPublishParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SarObservationUnvalidatedPublishParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}
