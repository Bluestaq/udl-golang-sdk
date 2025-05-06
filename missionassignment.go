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
)

// MissionAssignmentService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMissionAssignmentService] method instead.
type MissionAssignmentService struct {
	Options []option.RequestOption
	History MissionAssignmentHistoryService
}

// NewMissionAssignmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMissionAssignmentService(opts ...option.RequestOption) (r MissionAssignmentService) {
	r = MissionAssignmentService{}
	r.Options = opts
	r.History = NewMissionAssignmentHistoryService(opts...)
	return
}

// Service operation to take a single MissionAssignment as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *MissionAssignmentService) New(ctx context.Context, body MissionAssignmentNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/missionassignment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to update a single MissionAssignment. A specific role is
// required to perform this service operation. Please contact the UDL team for
// assistance.
func (r *MissionAssignmentService) Update(ctx context.Context, id string, body MissionAssignmentUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/missionassignment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *MissionAssignmentService) List(ctx context.Context, query MissionAssignmentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[MissionAssignmentListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/missionassignment"
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
func (r *MissionAssignmentService) ListAutoPaging(ctx context.Context, query MissionAssignmentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[MissionAssignmentListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to delete a MissionAssignment object specified by the passed
// ID path parameter. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *MissionAssignmentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/missionassignment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *MissionAssignmentService) Count(ctx context.Context, query MissionAssignmentCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/missionassignment/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple MissionAssignments as a POST body and ingest
// into the database. A specific role is required to perform this service
// operation. Please contact the UDL team for assistance.
func (r *MissionAssignmentService) NewBulk(ctx context.Context, body MissionAssignmentNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/missionassignment/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single MissionAssignment record by its unique ID
// passed as a path parameter.
func (r *MissionAssignmentService) Get(ctx context.Context, id string, query MissionAssignmentGetParams, opts ...option.RequestOption) (res *MissionAssignmentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/missionassignment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *MissionAssignmentService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/missionassignment/queryhelp"
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
func (r *MissionAssignmentService) Tuple(ctx context.Context, query MissionAssignmentTupleParams, opts ...option.RequestOption) (res *[]MissionAssignmentTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/missionassignment/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Platform mission assignment data.
type MissionAssignmentListResponse struct {
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
	DataMode MissionAssignmentListResponseDataMode `json:"dataMode,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis int64 `json:"c1associateddmpis"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air string `json:"c2air"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt int64 `json:"c2alt"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs int64 `json:"c2crs"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator string `json:"c2exerciseindicator"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof string `json:"c2exercisemof"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id string `json:"c2id"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor string `json:"c2idamplifyingdescriptor"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd string `json:"c2lnd"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc string `json:"c2spc"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd int64 `json:"c2spd"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator string `json:"c2specialinterestindicator"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur string `json:"c2sur"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv float64 `json:"c3elv"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat float64 `json:"c3lat"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon float64 `json:"c3lon"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl string `json:"c3ptl"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum string `json:"c3ptnum"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon int64 `json:"c4colon"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def string `json:"c4def"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress int64 `json:"c4egress"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod int64 `json:"c4mod"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores int64 `json:"c4numberofstores"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin int64 `json:"c4runin"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt string `json:"c4tgt"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete string `json:"c4timediscrete"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm int64 `json:"c4tm"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores int64 `json:"c4typeofstores"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon int64 `json:"c5colon"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs int64 `json:"c5elevationlsbs"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj int64 `json:"c5haeadj"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb int64 `json:"c5latlsb"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb int64 `json:"c5lonlsb"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng int64 `json:"c5tgtbrng"`
	// CONTINUATION WORD - time window.
	C5tw int64 `json:"c5tw"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc string `json:"c6dspc"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct string `json:"c6dspct"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm string `json:"c6fplpm"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel int64 `json:"c6intel"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser int64 `json:"c6laser"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm string `json:"c6longpm"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 int64 `json:"c6tnr3"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 float64 `json:"c7elang2"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p int64 `json:"c7in3p"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor string `json:"c7tnor"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Environment.
	Env string `json:"env"`
	// Index number.
	Index int64 `json:"index"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon float64 `json:"lon"`
	// Origin of index number.
	Orginx string `json:"orginx"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc string `json:"rc"`
	// Recurrence rate, receipt/compliance.
	Rr int64 `json:"rr"`
	// Strength.
	Sz string `json:"sz"`
	// Track number objective.
	Tno string `json:"tno"`
	// The track ID that the status is referencing, addressee.
	TrkID string `json:"trkId"`
	// Threat warning environment.
	Twenv string `json:"twenv"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		DataMode                   resp.Field
		Mad                        resp.Field
		Source                     resp.Field
		Ts                         resp.Field
		ID                         resp.Field
		C1associateddmpis          resp.Field
		C2air                      resp.Field
		C2alt                      resp.Field
		C2crs                      resp.Field
		C2exerciseindicator        resp.Field
		C2exercisemof              resp.Field
		C2id                       resp.Field
		C2idamplifyingdescriptor   resp.Field
		C2lnd                      resp.Field
		C2spc                      resp.Field
		C2spd                      resp.Field
		C2specialinterestindicator resp.Field
		C2sur                      resp.Field
		C3elv                      resp.Field
		C3lat                      resp.Field
		C3lon                      resp.Field
		C3ptl                      resp.Field
		C3ptnum                    resp.Field
		C4colon                    resp.Field
		C4def                      resp.Field
		C4egress                   resp.Field
		C4mod                      resp.Field
		C4numberofstores           resp.Field
		C4runin                    resp.Field
		C4tgt                      resp.Field
		C4timediscrete             resp.Field
		C4tm                       resp.Field
		C4typeofstores             resp.Field
		C5colon                    resp.Field
		C5elevationlsbs            resp.Field
		C5haeadj                   resp.Field
		C5latlsb                   resp.Field
		C5lonlsb                   resp.Field
		C5tgtbrng                  resp.Field
		C5tw                       resp.Field
		C6dspc                     resp.Field
		C6dspct                    resp.Field
		C6fplpm                    resp.Field
		C6intel                    resp.Field
		C6laser                    resp.Field
		C6longpm                   resp.Field
		C6tnr3                     resp.Field
		C7elang2                   resp.Field
		C7in3p                     resp.Field
		C7tnor                     resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		Env                        resp.Field
		Index                      resp.Field
		Lat                        resp.Field
		Lon                        resp.Field
		Orginx                     resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		Rc                         resp.Field
		Rr                         resp.Field
		Sz                         resp.Field
		Tno                        resp.Field
		TrkID                      resp.Field
		Twenv                      resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissionAssignmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *MissionAssignmentListResponse) UnmarshalJSON(data []byte) error {
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
type MissionAssignmentListResponseDataMode string

const (
	MissionAssignmentListResponseDataModeReal      MissionAssignmentListResponseDataMode = "REAL"
	MissionAssignmentListResponseDataModeTest      MissionAssignmentListResponseDataMode = "TEST"
	MissionAssignmentListResponseDataModeSimulated MissionAssignmentListResponseDataMode = "SIMULATED"
	MissionAssignmentListResponseDataModeExercise  MissionAssignmentListResponseDataMode = "EXERCISE"
)

// Platform mission assignment data.
type MissionAssignmentGetResponse struct {
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
	DataMode MissionAssignmentGetResponseDataMode `json:"dataMode,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis int64 `json:"c1associateddmpis"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air string `json:"c2air"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt int64 `json:"c2alt"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs int64 `json:"c2crs"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator string `json:"c2exerciseindicator"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof string `json:"c2exercisemof"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id string `json:"c2id"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor string `json:"c2idamplifyingdescriptor"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd string `json:"c2lnd"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc string `json:"c2spc"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd int64 `json:"c2spd"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator string `json:"c2specialinterestindicator"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur string `json:"c2sur"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv float64 `json:"c3elv"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat float64 `json:"c3lat"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon float64 `json:"c3lon"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl string `json:"c3ptl"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum string `json:"c3ptnum"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon int64 `json:"c4colon"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def string `json:"c4def"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress int64 `json:"c4egress"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod int64 `json:"c4mod"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores int64 `json:"c4numberofstores"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin int64 `json:"c4runin"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt string `json:"c4tgt"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete string `json:"c4timediscrete"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm int64 `json:"c4tm"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores int64 `json:"c4typeofstores"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon int64 `json:"c5colon"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs int64 `json:"c5elevationlsbs"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj int64 `json:"c5haeadj"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb int64 `json:"c5latlsb"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb int64 `json:"c5lonlsb"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng int64 `json:"c5tgtbrng"`
	// CONTINUATION WORD - time window.
	C5tw int64 `json:"c5tw"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc string `json:"c6dspc"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct string `json:"c6dspct"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm string `json:"c6fplpm"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel int64 `json:"c6intel"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser int64 `json:"c6laser"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm string `json:"c6longpm"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 int64 `json:"c6tnr3"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 float64 `json:"c7elang2"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p int64 `json:"c7in3p"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor string `json:"c7tnor"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Environment.
	Env string `json:"env"`
	// Index number.
	Index int64 `json:"index"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon float64 `json:"lon"`
	// Origin of index number.
	Orginx string `json:"orginx"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc string `json:"rc"`
	// Recurrence rate, receipt/compliance.
	Rr int64 `json:"rr"`
	// Strength.
	Sz string `json:"sz"`
	// Track number objective.
	Tno string `json:"tno"`
	// The track ID that the status is referencing, addressee.
	TrkID string `json:"trkId"`
	// Threat warning environment.
	Twenv string `json:"twenv"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		DataMode                   resp.Field
		Mad                        resp.Field
		Source                     resp.Field
		Ts                         resp.Field
		ID                         resp.Field
		C1associateddmpis          resp.Field
		C2air                      resp.Field
		C2alt                      resp.Field
		C2crs                      resp.Field
		C2exerciseindicator        resp.Field
		C2exercisemof              resp.Field
		C2id                       resp.Field
		C2idamplifyingdescriptor   resp.Field
		C2lnd                      resp.Field
		C2spc                      resp.Field
		C2spd                      resp.Field
		C2specialinterestindicator resp.Field
		C2sur                      resp.Field
		C3elv                      resp.Field
		C3lat                      resp.Field
		C3lon                      resp.Field
		C3ptl                      resp.Field
		C3ptnum                    resp.Field
		C4colon                    resp.Field
		C4def                      resp.Field
		C4egress                   resp.Field
		C4mod                      resp.Field
		C4numberofstores           resp.Field
		C4runin                    resp.Field
		C4tgt                      resp.Field
		C4timediscrete             resp.Field
		C4tm                       resp.Field
		C4typeofstores             resp.Field
		C5colon                    resp.Field
		C5elevationlsbs            resp.Field
		C5haeadj                   resp.Field
		C5latlsb                   resp.Field
		C5lonlsb                   resp.Field
		C5tgtbrng                  resp.Field
		C5tw                       resp.Field
		C6dspc                     resp.Field
		C6dspct                    resp.Field
		C6fplpm                    resp.Field
		C6intel                    resp.Field
		C6laser                    resp.Field
		C6longpm                   resp.Field
		C6tnr3                     resp.Field
		C7elang2                   resp.Field
		C7in3p                     resp.Field
		C7tnor                     resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		Env                        resp.Field
		Index                      resp.Field
		Lat                        resp.Field
		Lon                        resp.Field
		Orginx                     resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		Rc                         resp.Field
		Rr                         resp.Field
		Sz                         resp.Field
		Tno                        resp.Field
		TrkID                      resp.Field
		Twenv                      resp.Field
		UpdatedAt                  resp.Field
		UpdatedBy                  resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissionAssignmentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MissionAssignmentGetResponse) UnmarshalJSON(data []byte) error {
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
type MissionAssignmentGetResponseDataMode string

const (
	MissionAssignmentGetResponseDataModeReal      MissionAssignmentGetResponseDataMode = "REAL"
	MissionAssignmentGetResponseDataModeTest      MissionAssignmentGetResponseDataMode = "TEST"
	MissionAssignmentGetResponseDataModeSimulated MissionAssignmentGetResponseDataMode = "SIMULATED"
	MissionAssignmentGetResponseDataModeExercise  MissionAssignmentGetResponseDataMode = "EXERCISE"
)

// Platform mission assignment data.
type MissionAssignmentTupleResponse struct {
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
	DataMode MissionAssignmentTupleResponseDataMode `json:"dataMode,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis int64 `json:"c1associateddmpis"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air string `json:"c2air"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt int64 `json:"c2alt"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs int64 `json:"c2crs"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator string `json:"c2exerciseindicator"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof string `json:"c2exercisemof"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id string `json:"c2id"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor string `json:"c2idamplifyingdescriptor"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd string `json:"c2lnd"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc string `json:"c2spc"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd int64 `json:"c2spd"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator string `json:"c2specialinterestindicator"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur string `json:"c2sur"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv float64 `json:"c3elv"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat float64 `json:"c3lat"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon float64 `json:"c3lon"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl string `json:"c3ptl"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum string `json:"c3ptnum"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon int64 `json:"c4colon"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def string `json:"c4def"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress int64 `json:"c4egress"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod int64 `json:"c4mod"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores int64 `json:"c4numberofstores"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin int64 `json:"c4runin"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt string `json:"c4tgt"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete string `json:"c4timediscrete"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm int64 `json:"c4tm"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores int64 `json:"c4typeofstores"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon int64 `json:"c5colon"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs int64 `json:"c5elevationlsbs"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj int64 `json:"c5haeadj"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb int64 `json:"c5latlsb"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb int64 `json:"c5lonlsb"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng int64 `json:"c5tgtbrng"`
	// CONTINUATION WORD - time window.
	C5tw int64 `json:"c5tw"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc string `json:"c6dspc"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct string `json:"c6dspct"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm string `json:"c6fplpm"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel int64 `json:"c6intel"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser int64 `json:"c6laser"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm string `json:"c6longpm"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 int64 `json:"c6tnr3"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 float64 `json:"c7elang2"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p int64 `json:"c7in3p"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor string `json:"c7tnor"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Environment.
	Env string `json:"env"`
	// Index number.
	Index int64 `json:"index"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat float64 `json:"lat"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon float64 `json:"lon"`
	// Origin of index number.
	Orginx string `json:"orginx"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc string `json:"rc"`
	// Recurrence rate, receipt/compliance.
	Rr int64 `json:"rr"`
	// Strength.
	Sz string `json:"sz"`
	// Track number objective.
	Tno string `json:"tno"`
	// The track ID that the status is referencing, addressee.
	TrkID string `json:"trkId"`
	// Threat warning environment.
	Twenv string `json:"twenv"`
	// Time the row was updated in the database, auto-populated by the system.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Application user who updated the row in the database, auto-populated by the
	// system.
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		ClassificationMarking      resp.Field
		DataMode                   resp.Field
		Mad                        resp.Field
		Source                     resp.Field
		Ts                         resp.Field
		ID                         resp.Field
		C1associateddmpis          resp.Field
		C2air                      resp.Field
		C2alt                      resp.Field
		C2crs                      resp.Field
		C2exerciseindicator        resp.Field
		C2exercisemof              resp.Field
		C2id                       resp.Field
		C2idamplifyingdescriptor   resp.Field
		C2lnd                      resp.Field
		C2spc                      resp.Field
		C2spd                      resp.Field
		C2specialinterestindicator resp.Field
		C2sur                      resp.Field
		C3elv                      resp.Field
		C3lat                      resp.Field
		C3lon                      resp.Field
		C3ptl                      resp.Field
		C3ptnum                    resp.Field
		C4colon                    resp.Field
		C4def                      resp.Field
		C4egress                   resp.Field
		C4mod                      resp.Field
		C4numberofstores           resp.Field
		C4runin                    resp.Field
		C4tgt                      resp.Field
		C4timediscrete             resp.Field
		C4tm                       resp.Field
		C4typeofstores             resp.Field
		C5colon                    resp.Field
		C5elevationlsbs            resp.Field
		C5haeadj                   resp.Field
		C5latlsb                   resp.Field
		C5lonlsb                   resp.Field
		C5tgtbrng                  resp.Field
		C5tw                       resp.Field
		C6dspc                     resp.Field
		C6dspct                    resp.Field
		C6fplpm                    resp.Field
		C6intel                    resp.Field
		C6laser                    resp.Field
		C6longpm                   resp.Field
		C6tnr3                     resp.Field
		C7elang2                   resp.Field
		C7in3p                     resp.Field
		C7tnor                     resp.Field
		CreatedAt                  resp.Field
		CreatedBy                  resp.Field
		Env                        resp.Field
		Index                      resp.Field
		Lat                        resp.Field
		Lon                        resp.Field
		Orginx                     resp.Field
		Origin                     resp.Field
		OrigNetwork                resp.Field
		Rc                         resp.Field
		Rr                         resp.Field
		Sz                         resp.Field
		Tno                        resp.Field
		TrkID                      resp.Field
		Twenv                      resp.Field
		UpdatedAt                  resp.Field
		UpdatedBy                  resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissionAssignmentTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *MissionAssignmentTupleResponse) UnmarshalJSON(data []byte) error {
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
type MissionAssignmentTupleResponseDataMode string

const (
	MissionAssignmentTupleResponseDataModeReal      MissionAssignmentTupleResponseDataMode = "REAL"
	MissionAssignmentTupleResponseDataModeTest      MissionAssignmentTupleResponseDataMode = "TEST"
	MissionAssignmentTupleResponseDataModeSimulated MissionAssignmentTupleResponseDataMode = "SIMULATED"
	MissionAssignmentTupleResponseDataModeExercise  MissionAssignmentTupleResponseDataMode = "EXERCISE"
)

type MissionAssignmentNewParams struct {
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
	DataMode MissionAssignmentNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis param.Opt[int64] `json:"c1associateddmpis,omitzero"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air param.Opt[string] `json:"c2air,omitzero"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt param.Opt[int64] `json:"c2alt,omitzero"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs param.Opt[int64] `json:"c2crs,omitzero"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator param.Opt[string] `json:"c2exerciseindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof param.Opt[string] `json:"c2exercisemof,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id param.Opt[string] `json:"c2id,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor param.Opt[string] `json:"c2idamplifyingdescriptor,omitzero"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd param.Opt[string] `json:"c2lnd,omitzero"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc param.Opt[string] `json:"c2spc,omitzero"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd param.Opt[int64] `json:"c2spd,omitzero"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator param.Opt[string] `json:"c2specialinterestindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur param.Opt[string] `json:"c2sur,omitzero"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv param.Opt[float64] `json:"c3elv,omitzero"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat param.Opt[float64] `json:"c3lat,omitzero"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon param.Opt[float64] `json:"c3lon,omitzero"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl param.Opt[string] `json:"c3ptl,omitzero"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum param.Opt[string] `json:"c3ptnum,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon param.Opt[int64] `json:"c4colon,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def param.Opt[string] `json:"c4def,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress param.Opt[int64] `json:"c4egress,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod param.Opt[int64] `json:"c4mod,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores param.Opt[int64] `json:"c4numberofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin param.Opt[int64] `json:"c4runin,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt param.Opt[string] `json:"c4tgt,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete param.Opt[string] `json:"c4timediscrete,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm param.Opt[int64] `json:"c4tm,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores param.Opt[int64] `json:"c4typeofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon param.Opt[int64] `json:"c5colon,omitzero"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs param.Opt[int64] `json:"c5elevationlsbs,omitzero"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj param.Opt[int64] `json:"c5haeadj,omitzero"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb param.Opt[int64] `json:"c5latlsb,omitzero"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb param.Opt[int64] `json:"c5lonlsb,omitzero"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng param.Opt[int64] `json:"c5tgtbrng,omitzero"`
	// CONTINUATION WORD - time window.
	C5tw param.Opt[int64] `json:"c5tw,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc param.Opt[string] `json:"c6dspc,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct param.Opt[string] `json:"c6dspct,omitzero"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm param.Opt[string] `json:"c6fplpm,omitzero"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel param.Opt[int64] `json:"c6intel,omitzero"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser param.Opt[int64] `json:"c6laser,omitzero"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm param.Opt[string] `json:"c6longpm,omitzero"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 param.Opt[int64] `json:"c6tnr3,omitzero"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 param.Opt[float64] `json:"c7elang2,omitzero"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p param.Opt[int64] `json:"c7in3p,omitzero"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor param.Opt[string] `json:"c7tnor,omitzero"`
	// Environment.
	Env param.Opt[string] `json:"env,omitzero"`
	// Index number.
	Index param.Opt[int64] `json:"index,omitzero"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Origin of index number.
	Orginx param.Opt[string] `json:"orginx,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc param.Opt[string] `json:"rc,omitzero"`
	// Recurrence rate, receipt/compliance.
	Rr param.Opt[int64] `json:"rr,omitzero"`
	// Strength.
	Sz param.Opt[string] `json:"sz,omitzero"`
	// Track number objective.
	Tno param.Opt[string] `json:"tno,omitzero"`
	// The track ID that the status is referencing, addressee.
	TrkID param.Opt[string] `json:"trkId,omitzero"`
	// Threat warning environment.
	Twenv param.Opt[string] `json:"twenv,omitzero"`
	paramObj
}

func (r MissionAssignmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MissionAssignmentNewParams
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
type MissionAssignmentNewParamsDataMode string

const (
	MissionAssignmentNewParamsDataModeReal      MissionAssignmentNewParamsDataMode = "REAL"
	MissionAssignmentNewParamsDataModeTest      MissionAssignmentNewParamsDataMode = "TEST"
	MissionAssignmentNewParamsDataModeSimulated MissionAssignmentNewParamsDataMode = "SIMULATED"
	MissionAssignmentNewParamsDataModeExercise  MissionAssignmentNewParamsDataMode = "EXERCISE"
)

type MissionAssignmentUpdateParams struct {
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
	DataMode MissionAssignmentUpdateParamsDataMode `json:"dataMode,omitzero,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis param.Opt[int64] `json:"c1associateddmpis,omitzero"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air param.Opt[string] `json:"c2air,omitzero"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt param.Opt[int64] `json:"c2alt,omitzero"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs param.Opt[int64] `json:"c2crs,omitzero"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator param.Opt[string] `json:"c2exerciseindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof param.Opt[string] `json:"c2exercisemof,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id param.Opt[string] `json:"c2id,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor param.Opt[string] `json:"c2idamplifyingdescriptor,omitzero"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd param.Opt[string] `json:"c2lnd,omitzero"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc param.Opt[string] `json:"c2spc,omitzero"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd param.Opt[int64] `json:"c2spd,omitzero"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator param.Opt[string] `json:"c2specialinterestindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur param.Opt[string] `json:"c2sur,omitzero"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv param.Opt[float64] `json:"c3elv,omitzero"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat param.Opt[float64] `json:"c3lat,omitzero"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon param.Opt[float64] `json:"c3lon,omitzero"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl param.Opt[string] `json:"c3ptl,omitzero"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum param.Opt[string] `json:"c3ptnum,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon param.Opt[int64] `json:"c4colon,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def param.Opt[string] `json:"c4def,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress param.Opt[int64] `json:"c4egress,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod param.Opt[int64] `json:"c4mod,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores param.Opt[int64] `json:"c4numberofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin param.Opt[int64] `json:"c4runin,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt param.Opt[string] `json:"c4tgt,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete param.Opt[string] `json:"c4timediscrete,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm param.Opt[int64] `json:"c4tm,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores param.Opt[int64] `json:"c4typeofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon param.Opt[int64] `json:"c5colon,omitzero"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs param.Opt[int64] `json:"c5elevationlsbs,omitzero"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj param.Opt[int64] `json:"c5haeadj,omitzero"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb param.Opt[int64] `json:"c5latlsb,omitzero"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb param.Opt[int64] `json:"c5lonlsb,omitzero"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng param.Opt[int64] `json:"c5tgtbrng,omitzero"`
	// CONTINUATION WORD - time window.
	C5tw param.Opt[int64] `json:"c5tw,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc param.Opt[string] `json:"c6dspc,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct param.Opt[string] `json:"c6dspct,omitzero"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm param.Opt[string] `json:"c6fplpm,omitzero"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel param.Opt[int64] `json:"c6intel,omitzero"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser param.Opt[int64] `json:"c6laser,omitzero"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm param.Opt[string] `json:"c6longpm,omitzero"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 param.Opt[int64] `json:"c6tnr3,omitzero"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 param.Opt[float64] `json:"c7elang2,omitzero"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p param.Opt[int64] `json:"c7in3p,omitzero"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor param.Opt[string] `json:"c7tnor,omitzero"`
	// Environment.
	Env param.Opt[string] `json:"env,omitzero"`
	// Index number.
	Index param.Opt[int64] `json:"index,omitzero"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Origin of index number.
	Orginx param.Opt[string] `json:"orginx,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc param.Opt[string] `json:"rc,omitzero"`
	// Recurrence rate, receipt/compliance.
	Rr param.Opt[int64] `json:"rr,omitzero"`
	// Strength.
	Sz param.Opt[string] `json:"sz,omitzero"`
	// Track number objective.
	Tno param.Opt[string] `json:"tno,omitzero"`
	// The track ID that the status is referencing, addressee.
	TrkID param.Opt[string] `json:"trkId,omitzero"`
	// Threat warning environment.
	Twenv param.Opt[string] `json:"twenv,omitzero"`
	paramObj
}

func (r MissionAssignmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MissionAssignmentUpdateParams
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
type MissionAssignmentUpdateParamsDataMode string

const (
	MissionAssignmentUpdateParamsDataModeReal      MissionAssignmentUpdateParamsDataMode = "REAL"
	MissionAssignmentUpdateParamsDataModeTest      MissionAssignmentUpdateParamsDataMode = "TEST"
	MissionAssignmentUpdateParamsDataModeSimulated MissionAssignmentUpdateParamsDataMode = "SIMULATED"
	MissionAssignmentUpdateParamsDataModeExercise  MissionAssignmentUpdateParamsDataMode = "EXERCISE"
)

type MissionAssignmentListParams struct {
	// the timestamp of the mission data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissionAssignmentListParams]'s query parameters as
// `url.Values`.
func (r MissionAssignmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissionAssignmentCountParams struct {
	// the timestamp of the mission data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissionAssignmentCountParams]'s query parameters as
// `url.Values`.
func (r MissionAssignmentCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissionAssignmentNewBulkParams struct {
	Body []MissionAssignmentNewBulkParamsBody
	paramObj
}

func (r MissionAssignmentNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Platform mission assignment data.
//
// The properties ClassificationMarking, DataMode, Mad, Source, Ts are required.
type MissionAssignmentNewBulkParamsBody struct {
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
	DataMode string `json:"dataMode,omitzero,required"`
	// The mission assignment discrete value.
	Mad string `json:"mad,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// The timestamp of the mission data, in ISO 8601 UTC format.
	Ts time.Time `json:"ts,required" format:"date-time"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// TARGET POSITION CONTINUATION WORD - number of associated dmpis.
	C1associateddmpis param.Opt[int64] `json:"c1associateddmpis,omitzero"`
	// TARGET DATA CONTINUATION WORD - air specific type, see TABLE B-21.
	C2air param.Opt[string] `json:"c2air,omitzero"`
	// TARGET DATA CONTINUATION WORD - altitude, 100 FT, 2047=NS.
	C2alt param.Opt[int64] `json:"c2alt,omitzero"`
	// TARGET DATA CONTINUATION WORD - course in increments of 1 degree.
	C2crs param.Opt[int64] `json:"c2crs,omitzero"`
	// TARGET DATA CONTINUATION WORD - exercise indicator.
	C2exerciseindicator param.Opt[string] `json:"c2exerciseindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - method of fire.
	C2exercisemof param.Opt[string] `json:"c2exercisemof,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity.
	C2id param.Opt[string] `json:"c2id,omitzero"`
	// TARGET DATA CONTINUATION WORD - identity amplifying descriptor.
	C2idamplifyingdescriptor param.Opt[string] `json:"c2idamplifyingdescriptor,omitzero"`
	// TARGET DATA CONTINUATION WORD - land specific type, see TABLE B-21.
	C2lnd param.Opt[string] `json:"c2lnd,omitzero"`
	// TARGET DATA CONTINUATION WORD - space specific type, see TABLE B-39.
	C2spc param.Opt[string] `json:"c2spc,omitzero"`
	// TARGET DATA CONTINUATION WORD - speed in 2 DM/HR, 2047=NS.
	C2spd param.Opt[int64] `json:"c2spd,omitzero"`
	// TARGET DATA CONTINUATION WORD - special interest indicator.
	C2specialinterestindicator param.Opt[string] `json:"c2specialinterestindicator,omitzero"`
	// TARGET DATA CONTINUATION WORD - surface specific type, see TABLE B-21.
	C2sur param.Opt[string] `json:"c2sur,omitzero"`
	// POINT LOCATION CONTINUATION WORD - elevation, 25 FT, 1023=NS.
	C3elv param.Opt[float64] `json:"c3elv,omitzero"`
	// POINT LOCATION CONTINUATION WORD - latitude, 0.0013 MINUTE.
	C3lat param.Opt[float64] `json:"c3lat,omitzero"`
	// POINT LOCATION CONTINUATION WORD - longitude, 0.0013 MINUTE.
	C3lon param.Opt[float64] `json:"c3lon,omitzero"`
	// TARGET DATA CONTINUATION WORD - point type 1.
	C3ptl param.Opt[string] `json:"c3ptl,omitzero"`
	// TARGET DATA CONTINUATION WORD - point number.
	C3ptnum param.Opt[string] `json:"c3ptnum,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - minute.
	C4colon param.Opt[int64] `json:"c4colon,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target defenses.
	C4def param.Opt[string] `json:"c4def,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4egress param.Opt[int64] `json:"c4egress,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - mode of delivery.
	C4mod param.Opt[int64] `json:"c4mod,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - number of stores, NS=63.
	C4numberofstores param.Opt[int64] `json:"c4numberofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - run in heading, NS=511.
	C4runin param.Opt[int64] `json:"c4runin,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - target type - see TABLE B-32.
	C4tgt param.Opt[string] `json:"c4tgt,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - time discrete.
	C4timediscrete param.Opt[string] `json:"c4timediscrete,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - hour.
	C4tm param.Opt[int64] `json:"c4tm,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - type of stores.
	C4typeofstores param.Opt[int64] `json:"c4typeofstores,omitzero"`
	// SURFACE ATTACK CONTINUATION WORD - seconds in increments of 1 sec.
	C5colon param.Opt[int64] `json:"c5colon,omitzero"`
	// CONTINUATION WORD - used with c3_elv to double precision to approx 3 ft.
	C5elevationlsbs param.Opt[int64] `json:"c5elevationlsbs,omitzero"`
	// CONTINUATION WORD - hae adjustment, measured in 3.125 FT.
	C5haeadj param.Opt[int64] `json:"c5haeadj,omitzero"`
	// CONTINUATION WORD - used with c3_lat to double precision to approx 4 ft.
	C5latlsb param.Opt[int64] `json:"c5latlsb,omitzero"`
	// CONTINUATION WORD - used with c3_lon to double precision to approx 4 ft.
	C5lonlsb param.Opt[int64] `json:"c5lonlsb,omitzero"`
	// CONTINUATION WORD - target bearing.
	C5tgtbrng param.Opt[int64] `json:"c5tgtbrng,omitzero"`
	// CONTINUATION WORD - time window.
	C5tw param.Opt[int64] `json:"c5tw,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code.
	C6dspc param.Opt[string] `json:"c6dspc,omitzero"`
	// TARGETING CONTINUATION WORD - designator/seeker pulse code type.
	C6dspct param.Opt[string] `json:"c6dspct,omitzero"`
	// TARGETING CONTINUATION WORD - first pulse/last pulse mode.
	C6fplpm param.Opt[string] `json:"c6fplpm,omitzero"`
	// TARGETING CONTINUATION WORD - index number, related, 0=NS.
	C6intel param.Opt[int64] `json:"c6intel,omitzero"`
	// TARGETING CONTINUATION WORD - laser illuminator code.
	C6laser param.Opt[int64] `json:"c6laser,omitzero"`
	// TARGETING CONTINUATION WORD - long pulse mode.
	C6longpm param.Opt[string] `json:"c6longpm,omitzero"`
	// TARGETING CONTINUATION WORD - track number, related to 3.
	C6tnr3 param.Opt[int64] `json:"c6tnr3,omitzero"`
	// THIRD PARTY CONTINUATION WORD - elevation angle, 2.
	C7elang2 param.Opt[float64] `json:"c7elang2,omitzero"`
	// THIRD PARTY CONTINUATION WORD - index number, third party.
	C7in3p param.Opt[int64] `json:"c7in3p,omitzero"`
	// THIRD PARTY CONTINUATION WORD - track number, index originator.
	C7tnor param.Opt[string] `json:"c7tnor,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Environment.
	Env param.Opt[string] `json:"env,omitzero"`
	// Index number.
	Index param.Opt[int64] `json:"index,omitzero"`
	// WGS84 latitude, in degrees. -90 to 90 degrees (negative values south of
	// equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS84 longitude, in degrees. -180 to 180 degrees (negative values west of Prime
	// Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// Origin of index number.
	Orginx param.Opt[string] `json:"orginx,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Receipt/Compliance, values from TABLE B-9.
	Rc param.Opt[string] `json:"rc,omitzero"`
	// Recurrence rate, receipt/compliance.
	Rr param.Opt[int64] `json:"rr,omitzero"`
	// Strength.
	Sz param.Opt[string] `json:"sz,omitzero"`
	// Track number objective.
	Tno param.Opt[string] `json:"tno,omitzero"`
	// The track ID that the status is referencing, addressee.
	TrkID param.Opt[string] `json:"trkId,omitzero"`
	// Threat warning environment.
	Twenv param.Opt[string] `json:"twenv,omitzero"`
	paramObj
}

func (r MissionAssignmentNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow MissionAssignmentNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[MissionAssignmentNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type MissionAssignmentGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissionAssignmentGetParams]'s query parameters as
// `url.Values`.
func (r MissionAssignmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MissionAssignmentTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// the timestamp of the mission data, in ISO 8601 UTC format.
	// (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Ts          time.Time        `query:"ts,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissionAssignmentTupleParams]'s query parameters as
// `url.Values`.
func (r MissionAssignmentTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
