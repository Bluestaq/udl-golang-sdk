// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

// LaunchDetectionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchDetectionService] method instead.
type LaunchDetectionService struct {
	Options []option.RequestOption
}

// NewLaunchDetectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaunchDetectionService(opts ...option.RequestOption) (r LaunchDetectionService) {
	r = LaunchDetectionService{}
	r.Options = opts
	return
}

// Service operation to take a single launch detection as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LaunchDetectionService) New(ctx context.Context, body LaunchDetectionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchdetection"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single launch detection. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LaunchDetectionService) Update(ctx context.Context, id string, body LaunchDetectionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchdetection/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *LaunchDetectionService) List(ctx context.Context, query LaunchDetectionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchDetectionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/launchdetection"
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
func (r *LaunchDetectionService) ListAutoPaging(ctx context.Context, query LaunchDetectionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchDetectionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a launch detection object specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchDetectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchdetection/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *LaunchDetectionService) Count(ctx context.Context, query LaunchDetectionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchdetection/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single launch detection record by its unique ID
// passed as a path parameter.
func (r *LaunchDetectionService) Get(ctx context.Context, id string, query LaunchDetectionGetParams, opts ...option.RequestOption) (res *LaunchDetectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/launchdetection/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *LaunchDetectionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *LaunchDetectionQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchdetection/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *LaunchDetectionService) Tuple(ctx context.Context, query LaunchDetectionTupleParams, opts ...option.RequestOption) (res *[]LaunchDetectionTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchdetection/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Data to analyze launch detections.
type LaunchDetectionListResponse struct {
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
	DataMode LaunchDetectionListResponseDataMode `json:"dataMode,required"`
	// Type of message for the data.
	MessageType string `json:"messageType,required"`
	// Latitude of launch vehicle at observation time (in Degrees). -90 to 90 degrees
	// (negative values south of equator).
	ObservationLatitude float64 `json:"observationLatitude,required"`
	// Longitude of launch vehicle at observation time (in Degrees).
	ObservationLongitude float64 `json:"observationLongitude,required"`
	// Time of observation.
	ObservationTime time.Time `json:"observationTime,required" format:"date-time"`
	// Integer indicating how messages should be sequenced for a specific event.
	SequenceNumber int64 `json:"sequenceNumber,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Id to be able to correlate different messages to a specific event.
	EventID string `json:"eventId"`
	// Flag indicating that the Launch azimuth is uncertain due to near vertical flight
	// path.
	HighZenithAzimuth bool `json:"highZenithAzimuth"`
	// Orbit Inclination (in Degrees).
	Inclination float64 `json:"inclination"`
	// Angle measured clockwise from North for the launch heading (in Degrees).
	LaunchAzimuth float64 `json:"launchAzimuth"`
	// Geodetic Latitude of launch origin (in Degrees). -90 to 90 degrees (negative
	// values south of equator).
	LaunchLatitude float64 `json:"launchLatitude"`
	// Geodetic Longitude of launch origin (in Degrees). -180 to 180 degrees (negative
	// values west of Prime Meridian).
	LaunchLongitude float64 `json:"launchLongitude"`
	// Time of Launch.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// Altitude of launch vehicle at observation time (in KM).
	ObservationAltitude float64 `json:"observationAltitude"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Orbit Right Ascension of Ascending Node (in Degrees).
	Raan float64 `json:"raan"`
	// Flag indicating multiple observers were used.
	StereoFlag bool `json:"stereoFlag"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MessageType           respjson.Field
		ObservationLatitude   respjson.Field
		ObservationLongitude  respjson.Field
		ObservationTime       respjson.Field
		SequenceNumber        respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		EventID               respjson.Field
		HighZenithAzimuth     respjson.Field
		Inclination           respjson.Field
		LaunchAzimuth         respjson.Field
		LaunchLatitude        respjson.Field
		LaunchLongitude       respjson.Field
		LaunchTime            respjson.Field
		ObservationAltitude   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Raan                  respjson.Field
		StereoFlag            respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchDetectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchDetectionListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchDetectionListResponseDataMode string

const (
	LaunchDetectionListResponseDataModeReal      LaunchDetectionListResponseDataMode = "REAL"
	LaunchDetectionListResponseDataModeTest      LaunchDetectionListResponseDataMode = "TEST"
	LaunchDetectionListResponseDataModeSimulated LaunchDetectionListResponseDataMode = "SIMULATED"
	LaunchDetectionListResponseDataModeExercise  LaunchDetectionListResponseDataMode = "EXERCISE"
)

// Data to analyze launch detections.
type LaunchDetectionGetResponse struct {
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
	DataMode LaunchDetectionGetResponseDataMode `json:"dataMode,required"`
	// Type of message for the data.
	MessageType string `json:"messageType,required"`
	// Latitude of launch vehicle at observation time (in Degrees). -90 to 90 degrees
	// (negative values south of equator).
	ObservationLatitude float64 `json:"observationLatitude,required"`
	// Longitude of launch vehicle at observation time (in Degrees).
	ObservationLongitude float64 `json:"observationLongitude,required"`
	// Time of observation.
	ObservationTime time.Time `json:"observationTime,required" format:"date-time"`
	// Integer indicating how messages should be sequenced for a specific event.
	SequenceNumber int64 `json:"sequenceNumber,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Id to be able to correlate different messages to a specific event.
	EventID string `json:"eventId"`
	// Flag indicating that the Launch azimuth is uncertain due to near vertical flight
	// path.
	HighZenithAzimuth bool `json:"highZenithAzimuth"`
	// Orbit Inclination (in Degrees).
	Inclination float64 `json:"inclination"`
	// Angle measured clockwise from North for the launch heading (in Degrees).
	LaunchAzimuth float64 `json:"launchAzimuth"`
	// Geodetic Latitude of launch origin (in Degrees). -90 to 90 degrees (negative
	// values south of equator).
	LaunchLatitude float64 `json:"launchLatitude"`
	// Geodetic Longitude of launch origin (in Degrees). -180 to 180 degrees (negative
	// values west of Prime Meridian).
	LaunchLongitude float64 `json:"launchLongitude"`
	// Time of Launch.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// Altitude of launch vehicle at observation time (in KM).
	ObservationAltitude float64 `json:"observationAltitude"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Orbit Right Ascension of Ascending Node (in Degrees).
	Raan float64 `json:"raan"`
	// Flag indicating multiple observers were used.
	StereoFlag bool `json:"stereoFlag"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MessageType           respjson.Field
		ObservationLatitude   respjson.Field
		ObservationLongitude  respjson.Field
		ObservationTime       respjson.Field
		SequenceNumber        respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		EventID               respjson.Field
		HighZenithAzimuth     respjson.Field
		Inclination           respjson.Field
		LaunchAzimuth         respjson.Field
		LaunchLatitude        respjson.Field
		LaunchLongitude       respjson.Field
		LaunchTime            respjson.Field
		ObservationAltitude   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Raan                  respjson.Field
		StereoFlag            respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchDetectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchDetectionGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchDetectionGetResponseDataMode string

const (
	LaunchDetectionGetResponseDataModeReal      LaunchDetectionGetResponseDataMode = "REAL"
	LaunchDetectionGetResponseDataModeTest      LaunchDetectionGetResponseDataMode = "TEST"
	LaunchDetectionGetResponseDataModeSimulated LaunchDetectionGetResponseDataMode = "SIMULATED"
	LaunchDetectionGetResponseDataModeExercise  LaunchDetectionGetResponseDataMode = "EXERCISE"
)

type LaunchDetectionQueryhelpResponse struct {
	AodrSupported         bool                         `json:"aodrSupported"`
	ClassificationMarking string                       `json:"classificationMarking"`
	Description           string                       `json:"description"`
	HistorySupported      bool                         `json:"historySupported"`
	Name                  string                       `json:"name"`
	Parameters            []shared.ParamDescriptorResp `json:"parameters"`
	RequiredRoles         []string                     `json:"requiredRoles"`
	RestSupported         bool                         `json:"restSupported"`
	SortSupported         bool                         `json:"sortSupported"`
	TypeName              string                       `json:"typeName"`
	Uri                   string                       `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AodrSupported         respjson.Field
		ClassificationMarking respjson.Field
		Description           respjson.Field
		HistorySupported      respjson.Field
		Name                  respjson.Field
		Parameters            respjson.Field
		RequiredRoles         respjson.Field
		RestSupported         respjson.Field
		SortSupported         respjson.Field
		TypeName              respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchDetectionQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchDetectionQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data to analyze launch detections.
type LaunchDetectionTupleResponse struct {
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
	DataMode LaunchDetectionTupleResponseDataMode `json:"dataMode,required"`
	// Type of message for the data.
	MessageType string `json:"messageType,required"`
	// Latitude of launch vehicle at observation time (in Degrees). -90 to 90 degrees
	// (negative values south of equator).
	ObservationLatitude float64 `json:"observationLatitude,required"`
	// Longitude of launch vehicle at observation time (in Degrees).
	ObservationLongitude float64 `json:"observationLongitude,required"`
	// Time of observation.
	ObservationTime time.Time `json:"observationTime,required" format:"date-time"`
	// Integer indicating how messages should be sequenced for a specific event.
	SequenceNumber int64 `json:"sequenceNumber,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Time the row was created in the database.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// Id to be able to correlate different messages to a specific event.
	EventID string `json:"eventId"`
	// Flag indicating that the Launch azimuth is uncertain due to near vertical flight
	// path.
	HighZenithAzimuth bool `json:"highZenithAzimuth"`
	// Orbit Inclination (in Degrees).
	Inclination float64 `json:"inclination"`
	// Angle measured clockwise from North for the launch heading (in Degrees).
	LaunchAzimuth float64 `json:"launchAzimuth"`
	// Geodetic Latitude of launch origin (in Degrees). -90 to 90 degrees (negative
	// values south of equator).
	LaunchLatitude float64 `json:"launchLatitude"`
	// Geodetic Longitude of launch origin (in Degrees). -180 to 180 degrees (negative
	// values west of Prime Meridian).
	LaunchLongitude float64 `json:"launchLongitude"`
	// Time of Launch.
	LaunchTime time.Time `json:"launchTime" format:"date-time"`
	// Altitude of launch vehicle at observation time (in KM).
	ObservationAltitude float64 `json:"observationAltitude"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Orbit Right Ascension of Ascending Node (in Degrees).
	Raan float64 `json:"raan"`
	// Flag indicating multiple observers were used.
	StereoFlag bool `json:"stereoFlag"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// Read-only time the row was updated in the database, set automatically by the
	// system on update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who last updated the row in the database, set by the system
	// automatically and ignored on create/edit operations.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		MessageType           respjson.Field
		ObservationLatitude   respjson.Field
		ObservationLongitude  respjson.Field
		ObservationTime       respjson.Field
		SequenceNumber        respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		EventID               respjson.Field
		HighZenithAzimuth     respjson.Field
		Inclination           respjson.Field
		LaunchAzimuth         respjson.Field
		LaunchLatitude        respjson.Field
		LaunchLongitude       respjson.Field
		LaunchTime            respjson.Field
		ObservationAltitude   respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Raan                  respjson.Field
		StereoFlag            respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchDetectionTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchDetectionTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaunchDetectionTupleResponseDataMode string

const (
	LaunchDetectionTupleResponseDataModeReal      LaunchDetectionTupleResponseDataMode = "REAL"
	LaunchDetectionTupleResponseDataModeTest      LaunchDetectionTupleResponseDataMode = "TEST"
	LaunchDetectionTupleResponseDataModeSimulated LaunchDetectionTupleResponseDataMode = "SIMULATED"
	LaunchDetectionTupleResponseDataModeExercise  LaunchDetectionTupleResponseDataMode = "EXERCISE"
)

type LaunchDetectionNewParams struct {
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
	DataMode LaunchDetectionNewParamsDataMode `json:"dataMode,omitzero,required"`
	// Type of message for the data.
	MessageType string `json:"messageType,required"`
	// Latitude of launch vehicle at observation time (in Degrees). -90 to 90 degrees
	// (negative values south of equator).
	ObservationLatitude float64 `json:"observationLatitude,required"`
	// Longitude of launch vehicle at observation time (in Degrees).
	ObservationLongitude float64 `json:"observationLongitude,required"`
	// Time of observation.
	ObservationTime time.Time `json:"observationTime,required" format:"date-time"`
	// Integer indicating how messages should be sequenced for a specific event.
	SequenceNumber int64 `json:"sequenceNumber,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Id to be able to correlate different messages to a specific event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Flag indicating that the Launch azimuth is uncertain due to near vertical flight
	// path.
	HighZenithAzimuth param.Opt[bool] `json:"highZenithAzimuth,omitzero"`
	// Orbit Inclination (in Degrees).
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Angle measured clockwise from North for the launch heading (in Degrees).
	LaunchAzimuth param.Opt[float64] `json:"launchAzimuth,omitzero"`
	// Geodetic Latitude of launch origin (in Degrees). -90 to 90 degrees (negative
	// values south of equator).
	LaunchLatitude param.Opt[float64] `json:"launchLatitude,omitzero"`
	// Geodetic Longitude of launch origin (in Degrees). -180 to 180 degrees (negative
	// values west of Prime Meridian).
	LaunchLongitude param.Opt[float64] `json:"launchLongitude,omitzero"`
	// Time of Launch.
	LaunchTime param.Opt[time.Time] `json:"launchTime,omitzero" format:"date-time"`
	// Altitude of launch vehicle at observation time (in KM).
	ObservationAltitude param.Opt[float64] `json:"observationAltitude,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Orbit Right Ascension of Ascending Node (in Degrees).
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// Flag indicating multiple observers were used.
	StereoFlag param.Opt[bool] `json:"stereoFlag,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LaunchDetectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchDetectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchDetectionNewParams) UnmarshalJSON(data []byte) error {
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
type LaunchDetectionNewParamsDataMode string

const (
	LaunchDetectionNewParamsDataModeReal      LaunchDetectionNewParamsDataMode = "REAL"
	LaunchDetectionNewParamsDataModeTest      LaunchDetectionNewParamsDataMode = "TEST"
	LaunchDetectionNewParamsDataModeSimulated LaunchDetectionNewParamsDataMode = "SIMULATED"
	LaunchDetectionNewParamsDataModeExercise  LaunchDetectionNewParamsDataMode = "EXERCISE"
)

type LaunchDetectionUpdateParams struct {
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
	DataMode LaunchDetectionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// Type of message for the data.
	MessageType string `json:"messageType,required"`
	// Latitude of launch vehicle at observation time (in Degrees). -90 to 90 degrees
	// (negative values south of equator).
	ObservationLatitude float64 `json:"observationLatitude,required"`
	// Longitude of launch vehicle at observation time (in Degrees).
	ObservationLongitude float64 `json:"observationLongitude,required"`
	// Time of observation.
	ObservationTime time.Time `json:"observationTime,required" format:"date-time"`
	// Integer indicating how messages should be sequenced for a specific event.
	SequenceNumber int64 `json:"sequenceNumber,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// Id to be able to correlate different messages to a specific event.
	EventID param.Opt[string] `json:"eventId,omitzero"`
	// Flag indicating that the Launch azimuth is uncertain due to near vertical flight
	// path.
	HighZenithAzimuth param.Opt[bool] `json:"highZenithAzimuth,omitzero"`
	// Orbit Inclination (in Degrees).
	Inclination param.Opt[float64] `json:"inclination,omitzero"`
	// Angle measured clockwise from North for the launch heading (in Degrees).
	LaunchAzimuth param.Opt[float64] `json:"launchAzimuth,omitzero"`
	// Geodetic Latitude of launch origin (in Degrees). -90 to 90 degrees (negative
	// values south of equator).
	LaunchLatitude param.Opt[float64] `json:"launchLatitude,omitzero"`
	// Geodetic Longitude of launch origin (in Degrees). -180 to 180 degrees (negative
	// values west of Prime Meridian).
	LaunchLongitude param.Opt[float64] `json:"launchLongitude,omitzero"`
	// Time of Launch.
	LaunchTime param.Opt[time.Time] `json:"launchTime,omitzero" format:"date-time"`
	// Altitude of launch vehicle at observation time (in KM).
	ObservationAltitude param.Opt[float64] `json:"observationAltitude,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Orbit Right Ascension of Ascending Node (in Degrees).
	Raan param.Opt[float64] `json:"raan,omitzero"`
	// Flag indicating multiple observers were used.
	StereoFlag param.Opt[bool] `json:"stereoFlag,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LaunchDetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchDetectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaunchDetectionUpdateParams) UnmarshalJSON(data []byte) error {
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
type LaunchDetectionUpdateParamsDataMode string

const (
	LaunchDetectionUpdateParamsDataModeReal      LaunchDetectionUpdateParamsDataMode = "REAL"
	LaunchDetectionUpdateParamsDataModeTest      LaunchDetectionUpdateParamsDataMode = "TEST"
	LaunchDetectionUpdateParamsDataModeSimulated LaunchDetectionUpdateParamsDataMode = "SIMULATED"
	LaunchDetectionUpdateParamsDataModeExercise  LaunchDetectionUpdateParamsDataMode = "EXERCISE"
)

type LaunchDetectionListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchDetectionListParams]'s query parameters as
// `url.Values`.
func (r LaunchDetectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchDetectionCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchDetectionCountParams]'s query parameters as
// `url.Values`.
func (r LaunchDetectionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchDetectionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchDetectionGetParams]'s query parameters as
// `url.Values`.
func (r LaunchDetectionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchDetectionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaunchDetectionTupleParams]'s query parameters as
// `url.Values`.
func (r LaunchDetectionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
