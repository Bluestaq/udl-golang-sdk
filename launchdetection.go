// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
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
)

// LaunchdetectionService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaunchdetectionService] method instead.
type LaunchdetectionService struct {
	Options []option.RequestOption
}

// NewLaunchdetectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaunchdetectionService(opts ...option.RequestOption) (r LaunchdetectionService) {
	r = LaunchdetectionService{}
	r.Options = opts
	return
}

// Service operation to take a single launch detection as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *LaunchdetectionService) New(ctx context.Context, body LaunchdetectionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchdetection"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single launch detection. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *LaunchdetectionService) Update(ctx context.Context, id string, body LaunchdetectionUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *LaunchdetectionService) List(ctx context.Context, query LaunchdetectionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LaunchdetectionListResponse], err error) {
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
func (r *LaunchdetectionService) ListAutoPaging(ctx context.Context, query LaunchdetectionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LaunchdetectionListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a launch detection object specified by the passed ID
// path parameter. A specific role is required to perform this service operation.
// Please contact the UDL team for assistance.
func (r *LaunchdetectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
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
func (r *LaunchdetectionService) Count(ctx context.Context, query LaunchdetectionCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/launchdetection/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to get a single launch detection record by its unique ID
// passed as a path parameter.
func (r *LaunchdetectionService) Get(ctx context.Context, id string, query LaunchdetectionGetParams, opts ...option.RequestOption) (res *LaunchdetectionGetResponse, err error) {
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
func (r *LaunchdetectionService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/launchdetection/queryhelp"
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
func (r *LaunchdetectionService) Tuple(ctx context.Context, query LaunchdetectionTupleParams, opts ...option.RequestOption) (res *[]LaunchdetectionTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/launchdetection/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Data to analyze launch detections.
type LaunchdetectionListResponse struct {
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
	DataMode LaunchdetectionListResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		MessageType           resp.Field
		ObservationLatitude   resp.Field
		ObservationLongitude  resp.Field
		ObservationTime       resp.Field
		SequenceNumber        resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		EventID               resp.Field
		HighZenithAzimuth     resp.Field
		Inclination           resp.Field
		LaunchAzimuth         resp.Field
		LaunchLatitude        resp.Field
		LaunchLongitude       resp.Field
		LaunchTime            resp.Field
		ObservationAltitude   resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Raan                  resp.Field
		StereoFlag            resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchdetectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchdetectionListResponse) UnmarshalJSON(data []byte) error {
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
type LaunchdetectionListResponseDataMode string

const (
	LaunchdetectionListResponseDataModeReal      LaunchdetectionListResponseDataMode = "REAL"
	LaunchdetectionListResponseDataModeTest      LaunchdetectionListResponseDataMode = "TEST"
	LaunchdetectionListResponseDataModeSimulated LaunchdetectionListResponseDataMode = "SIMULATED"
	LaunchdetectionListResponseDataModeExercise  LaunchdetectionListResponseDataMode = "EXERCISE"
)

// Data to analyze launch detections.
type LaunchdetectionGetResponse struct {
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
	DataMode LaunchdetectionGetResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		MessageType           resp.Field
		ObservationLatitude   resp.Field
		ObservationLongitude  resp.Field
		ObservationTime       resp.Field
		SequenceNumber        resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		EventID               resp.Field
		HighZenithAzimuth     resp.Field
		Inclination           resp.Field
		LaunchAzimuth         resp.Field
		LaunchLatitude        resp.Field
		LaunchLongitude       resp.Field
		LaunchTime            resp.Field
		ObservationAltitude   resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Raan                  resp.Field
		StereoFlag            resp.Field
		Tags                  resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchdetectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchdetectionGetResponse) UnmarshalJSON(data []byte) error {
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
type LaunchdetectionGetResponseDataMode string

const (
	LaunchdetectionGetResponseDataModeReal      LaunchdetectionGetResponseDataMode = "REAL"
	LaunchdetectionGetResponseDataModeTest      LaunchdetectionGetResponseDataMode = "TEST"
	LaunchdetectionGetResponseDataModeSimulated LaunchdetectionGetResponseDataMode = "SIMULATED"
	LaunchdetectionGetResponseDataModeExercise  LaunchdetectionGetResponseDataMode = "EXERCISE"
)

// Data to analyze launch detections.
type LaunchdetectionTupleResponse struct {
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
	DataMode LaunchdetectionTupleResponseDataMode `json:"dataMode,required"`
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClassificationMarking resp.Field
		DataMode              resp.Field
		MessageType           resp.Field
		ObservationLatitude   resp.Field
		ObservationLongitude  resp.Field
		ObservationTime       resp.Field
		SequenceNumber        resp.Field
		Source                resp.Field
		ID                    resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		Descriptor            resp.Field
		EventID               resp.Field
		HighZenithAzimuth     resp.Field
		Inclination           resp.Field
		LaunchAzimuth         resp.Field
		LaunchLatitude        resp.Field
		LaunchLongitude       resp.Field
		LaunchTime            resp.Field
		ObservationAltitude   resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Raan                  resp.Field
		StereoFlag            resp.Field
		Tags                  resp.Field
		UpdatedAt             resp.Field
		UpdatedBy             resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaunchdetectionTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *LaunchdetectionTupleResponse) UnmarshalJSON(data []byte) error {
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
type LaunchdetectionTupleResponseDataMode string

const (
	LaunchdetectionTupleResponseDataModeReal      LaunchdetectionTupleResponseDataMode = "REAL"
	LaunchdetectionTupleResponseDataModeTest      LaunchdetectionTupleResponseDataMode = "TEST"
	LaunchdetectionTupleResponseDataModeSimulated LaunchdetectionTupleResponseDataMode = "SIMULATED"
	LaunchdetectionTupleResponseDataModeExercise  LaunchdetectionTupleResponseDataMode = "EXERCISE"
)

type LaunchdetectionNewParams struct {
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
	DataMode LaunchdetectionNewParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LaunchdetectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchdetectionNewParams
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
type LaunchdetectionNewParamsDataMode string

const (
	LaunchdetectionNewParamsDataModeReal      LaunchdetectionNewParamsDataMode = "REAL"
	LaunchdetectionNewParamsDataModeTest      LaunchdetectionNewParamsDataMode = "TEST"
	LaunchdetectionNewParamsDataModeSimulated LaunchdetectionNewParamsDataMode = "SIMULATED"
	LaunchdetectionNewParamsDataModeExercise  LaunchdetectionNewParamsDataMode = "EXERCISE"
)

type LaunchdetectionUpdateParams struct {
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
	DataMode LaunchdetectionUpdateParamsDataMode `json:"dataMode,omitzero,required"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LaunchdetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaunchdetectionUpdateParams
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
type LaunchdetectionUpdateParamsDataMode string

const (
	LaunchdetectionUpdateParamsDataModeReal      LaunchdetectionUpdateParamsDataMode = "REAL"
	LaunchdetectionUpdateParamsDataModeTest      LaunchdetectionUpdateParamsDataMode = "TEST"
	LaunchdetectionUpdateParamsDataModeSimulated LaunchdetectionUpdateParamsDataMode = "SIMULATED"
	LaunchdetectionUpdateParamsDataModeExercise  LaunchdetectionUpdateParamsDataMode = "EXERCISE"
)

type LaunchdetectionListParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LaunchdetectionListParams]'s query parameters as
// `url.Values`.
func (r LaunchdetectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchdetectionCountParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LaunchdetectionCountParams]'s query parameters as
// `url.Values`.
func (r LaunchdetectionCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchdetectionGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LaunchdetectionGetParams]'s query parameters as
// `url.Values`.
func (r LaunchdetectionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaunchdetectionTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns     string           `query:"columns,required" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LaunchdetectionTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LaunchdetectionTupleParams]'s query parameters as
// `url.Values`.
func (r LaunchdetectionTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
