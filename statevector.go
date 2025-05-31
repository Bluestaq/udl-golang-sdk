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
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/respjson"
	"github.com/stainless-sdks/unifieddatalibrary-go/shared"
)

// StateVectorService contains methods and other services that help with
// interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStateVectorService] method instead.
type StateVectorService struct {
	Options []option.RequestOption
	History StateVectorHistoryService
	Current StateVectorCurrentService
}

// NewStateVectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStateVectorService(opts ...option.RequestOption) (r StateVectorService) {
	r = StateVectorService{}
	r.Options = opts
	r.History = NewStateVectorHistoryService(opts...)
	r.Current = NewStateVectorCurrentService(opts...)
	return
}

// Service operation to take a single state vector as a POST body and ingest into
// the database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *StateVectorService) New(ctx context.Context, body StateVectorNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/statevector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *StateVectorService) List(ctx context.Context, query StateVectorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[StateVectorAbridged], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/statevector"
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
func (r *StateVectorService) ListAutoPaging(ctx context.Context, query StateVectorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[StateVectorAbridged] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *StateVectorService) Count(ctx context.Context, query StateVectorCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/statevector/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of state
// vectors as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *StateVectorService) NewBulk(ctx context.Context, body StateVectorNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/statevector/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single state vector by its unique ID passed as a path
// parameter.
func (r *StateVectorService) Get(ctx context.Context, id string, query StateVectorGetParams, opts ...option.RequestOption) (res *StateVectorFull, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/statevector/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *StateVectorService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *StateVectorQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/statevector/queryhelp"
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
func (r *StateVectorService) Tuple(ctx context.Context, query StateVectorTupleParams, opts ...option.RequestOption) (res *[]StateVectorFull, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/statevector/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to take multiple state vectors as a POST body and ingest into
// the database. This operation is intended to be used for automated feeds into
// UDL. A specific role is required to perform this service operation. Please
// contact the UDL team for assistance.
func (r *StateVectorService) UnvalidatedPublish(ctx context.Context, body StateVectorUnvalidatedPublishParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sv"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type StateVectorAbridged struct {
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
	DataMode StateVectorAbridgedDataMode `json:"dataMode,required"`
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
	CovReferenceFrame StateVectorAbridgedCovReferenceFrame `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame StateVectorAbridgedReferenceFrame `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		Source                respjson.Field
		ActualOdSpan          respjson.Field
		Algorithm             respjson.Field
		Alt1ReferenceFrame    respjson.Field
		Alt2ReferenceFrame    respjson.Field
		Area                  respjson.Field
		BDot                  respjson.Field
		CmOffset              respjson.Field
		Cov                   respjson.Field
		CovMethod             respjson.Field
		CovReferenceFrame     respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DragArea              respjson.Field
		DragCoeff             respjson.Field
		DragModel             respjson.Field
		Edr                   respjson.Field
		EqCov                 respjson.Field
		ErrorControl          respjson.Field
		FixedStep             respjson.Field
		GeopotentialModel     respjson.Field
		Iau1980Terms          respjson.Field
		IDOnOrbit             respjson.Field
		IDOrbitDetermination  respjson.Field
		IDStateVector         respjson.Field
		IntegratorMode        respjson.Field
		InTrackThrust         respjson.Field
		LastObEnd             respjson.Field
		LastObStart           respjson.Field
		LeapSecondTime        respjson.Field
		LunarSolar            respjson.Field
		Mass                  respjson.Field
		ObsAvailable          respjson.Field
		ObsUsed               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		Partials              respjson.Field
		Pedigree              respjson.Field
		PolarMotionX          respjson.Field
		PolarMotionY          respjson.Field
		PosUnc                respjson.Field
		RecOdSpan             respjson.Field
		ReferenceFrame        respjson.Field
		ResidualsAcc          respjson.Field
		RevNo                 respjson.Field
		Rms                   respjson.Field
		SatNo                 respjson.Field
		SigmaPosUvw           respjson.Field
		SigmaVelUvw           respjson.Field
		SolarFluxApAvg        respjson.Field
		SolarFluxF10          respjson.Field
		SolarFluxF10Avg       respjson.Field
		SolarRadPress         respjson.Field
		SolarRadPressCoeff    respjson.Field
		SolidEarthTides       respjson.Field
		SourceDl              respjson.Field
		SrpArea               respjson.Field
		StepMode              respjson.Field
		StepSize              respjson.Field
		StepSizeSelection     respjson.Field
		TaiUtc                respjson.Field
		ThrustAccel           respjson.Field
		TracksAvail           respjson.Field
		TracksUsed            respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		Ut1Rate               respjson.Field
		Ut1Utc                respjson.Field
		VelUnc                respjson.Field
		Xaccel                respjson.Field
		Xpos                  respjson.Field
		XposAlt1              respjson.Field
		XposAlt2              respjson.Field
		Xvel                  respjson.Field
		XvelAlt1              respjson.Field
		XvelAlt2              respjson.Field
		Yaccel                respjson.Field
		Ypos                  respjson.Field
		YposAlt1              respjson.Field
		YposAlt2              respjson.Field
		Yvel                  respjson.Field
		YvelAlt1              respjson.Field
		YvelAlt2              respjson.Field
		Zaccel                respjson.Field
		Zpos                  respjson.Field
		ZposAlt1              respjson.Field
		ZposAlt2              respjson.Field
		Zvel                  respjson.Field
		ZvelAlt1              respjson.Field
		ZvelAlt2              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StateVectorAbridged) RawJSON() string { return r.JSON.raw }
func (r *StateVectorAbridged) UnmarshalJSON(data []byte) error {
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
type StateVectorAbridgedDataMode string

const (
	StateVectorAbridgedDataModeReal      StateVectorAbridgedDataMode = "REAL"
	StateVectorAbridgedDataModeTest      StateVectorAbridgedDataMode = "TEST"
	StateVectorAbridgedDataModeSimulated StateVectorAbridgedDataMode = "SIMULATED"
	StateVectorAbridgedDataModeExercise  StateVectorAbridgedDataMode = "EXERCISE"
)

// The reference frame of the covariance matrix elements. If the covReferenceFrame
// is null it is assumed to be J2000.
type StateVectorAbridgedCovReferenceFrame string

const (
	StateVectorAbridgedCovReferenceFrameJ2000  StateVectorAbridgedCovReferenceFrame = "J2000"
	StateVectorAbridgedCovReferenceFrameUvw    StateVectorAbridgedCovReferenceFrame = "UVW"
	StateVectorAbridgedCovReferenceFrameEfgTdr StateVectorAbridgedCovReferenceFrame = "EFG/TDR"
	StateVectorAbridgedCovReferenceFrameTeme   StateVectorAbridgedCovReferenceFrame = "TEME"
	StateVectorAbridgedCovReferenceFrameGcrf   StateVectorAbridgedCovReferenceFrame = "GCRF"
)

// The reference frame of the cartesian orbital states. If the referenceFrame is
// null it is assumed to be J2000.
type StateVectorAbridgedReferenceFrame string

const (
	StateVectorAbridgedReferenceFrameJ2000   StateVectorAbridgedReferenceFrame = "J2000"
	StateVectorAbridgedReferenceFrameEfgTdr  StateVectorAbridgedReferenceFrame = "EFG/TDR"
	StateVectorAbridgedReferenceFrameEcrEcef StateVectorAbridgedReferenceFrame = "ECR/ECEF"
	StateVectorAbridgedReferenceFrameTeme    StateVectorAbridgedReferenceFrame = "TEME"
	StateVectorAbridgedReferenceFrameItrf    StateVectorAbridgedReferenceFrame = "ITRF"
	StateVectorAbridgedReferenceFrameGcrf    StateVectorAbridgedReferenceFrame = "GCRF"
)

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
type StateVectorFull struct {
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
	DataMode StateVectorFullDataMode `json:"dataMode,required"`
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan float64 `json:"actualODSpan"`
	// Optional algorithm used to produce this record.
	Algorithm string `json:"algorithm"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame string `json:"alt1ReferenceFrame"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame string `json:"alt2ReferenceFrame"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area float64 `json:"area"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot float64 `json:"bDot"`
	// Model parameter value for center of mass offset (m).
	CmOffset float64 `json:"cmOffset"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod string `json:"covMethod"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
	CovReferenceFrame StateVectorFullCovReferenceFrame `json:"covReferenceFrame"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor string `json:"descriptor"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea float64 `json:"dragArea"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff float64 `json:"dragCoeff"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel string `json:"dragModel"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr float64 `json:"edr"`
	// Start time at which this state vector was the 'current' state vector for its
	// satellite.
	EffectiveFrom time.Time `json:"effectiveFrom" format:"date-time"`
	// End time at which this state vector was no longer the 'current' state vector for
	// its satellite.
	EffectiveUntil time.Time `json:"effectiveUntil" format:"date-time"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov"`
	// Integrator error control.
	ErrorControl float64 `json:"errorControl"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep bool `json:"fixedStep"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel string `json:"geopotentialModel"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms int64 `json:"iau1980Terms"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit string `json:"idOnOrbit"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination string `json:"idOrbitDetermination"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector string `json:"idStateVector"`
	// Integrator Mode.
	IntegratorMode string `json:"integratorMode"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust bool `json:"inTrackThrust"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd time.Time `json:"lastObEnd" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart time.Time `json:"lastObStart" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime time.Time `json:"leapSecondTime" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar bool `json:"lunarSolar"`
	// The mass of the object, in kilograms.
	Mass float64 `json:"mass"`
	// The number of observations available for the OD of the object.
	ObsAvailable int64 `json:"obsAvailable"`
	// The number of observations accepted for the OD of the object.
	ObsUsed int64 `json:"obsUsed"`
	// Model object representing on-orbit objects or satellites in the system.
	OnOrbit shared.OnorbitFull `json:"onOrbit"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID string `json:"origObjectId"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials string `json:"partials"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree string `json:"pedigree"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX float64 `json:"polarMotionX"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY float64 `json:"polarMotionY"`
	// One sigma position uncertainty, in kilometers.
	PosUnc float64 `json:"posUnc"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri string `json:"rawFileURI"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan float64 `json:"recODSpan"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame StateVectorFullReferenceFrame `json:"referenceFrame"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc float64 `json:"residualsAcc"`
	// Epoch revolution number.
	RevNo int64 `json:"revNo"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms float64 `json:"rms"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo int64 `json:"satNo"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg float64 `json:"solarFluxAPAvg"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 float64 `json:"solarFluxF10"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg float64 `json:"solarFluxF10Avg"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress bool `json:"solarRadPress"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff float64 `json:"solarRadPressCoeff"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides bool `json:"solidEarthTides"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData"`
	// Optional array of UDL observation data types used to build this state vector
	// (e.g. EO, RADAR, RF, DOA). See the associated sourcedData array for the specific
	// UUIDs of observations for the positionally corresponding data types in this
	// array (the two arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl string `json:"sourceDL"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea float64 `json:"srpArea"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode string `json:"stepMode"`
	// Initial integration step size (seconds).
	StepSize float64 `json:"stepSize"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection string `json:"stepSizeSelection"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc float64 `json:"taiUtc"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel float64 `json:"thrustAccel"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail int64 `json:"tracksAvail"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed int64 `json:"tracksUsed"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID string `json:"transactionId"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct bool `json:"uct"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate float64 `json:"ut1Rate"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc float64 `json:"ut1Utc"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc float64 `json:"velUnc"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel float64 `json:"xaccel"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos float64 `json:"xpos"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 float64 `json:"xposAlt1"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 float64 `json:"xposAlt2"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel float64 `json:"xvel"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 float64 `json:"xvelAlt1"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 float64 `json:"xvelAlt2"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel float64 `json:"yaccel"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos float64 `json:"ypos"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 float64 `json:"yposAlt1"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 float64 `json:"yposAlt2"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel float64 `json:"yvel"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 float64 `json:"yvelAlt1"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 float64 `json:"yvelAlt2"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel float64 `json:"zaccel"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos float64 `json:"zpos"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 float64 `json:"zposAlt1"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 float64 `json:"zposAlt2"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel float64 `json:"zvel"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 float64 `json:"zvelAlt1"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 float64 `json:"zvelAlt2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		Epoch                 respjson.Field
		Source                respjson.Field
		ActualOdSpan          respjson.Field
		Algorithm             respjson.Field
		Alt1ReferenceFrame    respjson.Field
		Alt2ReferenceFrame    respjson.Field
		Area                  respjson.Field
		BDot                  respjson.Field
		CmOffset              respjson.Field
		Cov                   respjson.Field
		CovMethod             respjson.Field
		CovReferenceFrame     respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Descriptor            respjson.Field
		DragArea              respjson.Field
		DragCoeff             respjson.Field
		DragModel             respjson.Field
		Edr                   respjson.Field
		EffectiveFrom         respjson.Field
		EffectiveUntil        respjson.Field
		EqCov                 respjson.Field
		ErrorControl          respjson.Field
		FixedStep             respjson.Field
		GeopotentialModel     respjson.Field
		Iau1980Terms          respjson.Field
		IDOnOrbit             respjson.Field
		IDOrbitDetermination  respjson.Field
		IDStateVector         respjson.Field
		IntegratorMode        respjson.Field
		InTrackThrust         respjson.Field
		LastObEnd             respjson.Field
		LastObStart           respjson.Field
		LeapSecondTime        respjson.Field
		LunarSolar            respjson.Field
		Mass                  respjson.Field
		ObsAvailable          respjson.Field
		ObsUsed               respjson.Field
		OnOrbit               respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		OrigObjectID          respjson.Field
		Partials              respjson.Field
		Pedigree              respjson.Field
		PolarMotionX          respjson.Field
		PolarMotionY          respjson.Field
		PosUnc                respjson.Field
		RawFileUri            respjson.Field
		RecOdSpan             respjson.Field
		ReferenceFrame        respjson.Field
		ResidualsAcc          respjson.Field
		RevNo                 respjson.Field
		Rms                   respjson.Field
		SatNo                 respjson.Field
		SigmaPosUvw           respjson.Field
		SigmaVelUvw           respjson.Field
		SolarFluxApAvg        respjson.Field
		SolarFluxF10          respjson.Field
		SolarFluxF10Avg       respjson.Field
		SolarRadPress         respjson.Field
		SolarRadPressCoeff    respjson.Field
		SolidEarthTides       respjson.Field
		SourcedData           respjson.Field
		SourcedDataTypes      respjson.Field
		SourceDl              respjson.Field
		SrpArea               respjson.Field
		StepMode              respjson.Field
		StepSize              respjson.Field
		StepSizeSelection     respjson.Field
		Tags                  respjson.Field
		TaiUtc                respjson.Field
		ThrustAccel           respjson.Field
		TracksAvail           respjson.Field
		TracksUsed            respjson.Field
		TransactionID         respjson.Field
		Uct                   respjson.Field
		Ut1Rate               respjson.Field
		Ut1Utc                respjson.Field
		VelUnc                respjson.Field
		Xaccel                respjson.Field
		Xpos                  respjson.Field
		XposAlt1              respjson.Field
		XposAlt2              respjson.Field
		Xvel                  respjson.Field
		XvelAlt1              respjson.Field
		XvelAlt2              respjson.Field
		Yaccel                respjson.Field
		Ypos                  respjson.Field
		YposAlt1              respjson.Field
		YposAlt2              respjson.Field
		Yvel                  respjson.Field
		YvelAlt1              respjson.Field
		YvelAlt2              respjson.Field
		Zaccel                respjson.Field
		Zpos                  respjson.Field
		ZposAlt1              respjson.Field
		ZposAlt2              respjson.Field
		Zvel                  respjson.Field
		ZvelAlt1              respjson.Field
		ZvelAlt2              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StateVectorFull) RawJSON() string { return r.JSON.raw }
func (r *StateVectorFull) UnmarshalJSON(data []byte) error {
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
type StateVectorFullDataMode string

const (
	StateVectorFullDataModeReal      StateVectorFullDataMode = "REAL"
	StateVectorFullDataModeTest      StateVectorFullDataMode = "TEST"
	StateVectorFullDataModeSimulated StateVectorFullDataMode = "SIMULATED"
	StateVectorFullDataModeExercise  StateVectorFullDataMode = "EXERCISE"
)

// The reference frame of the covariance matrix elements. If the covReferenceFrame
// is null it is assumed to be J2000.
type StateVectorFullCovReferenceFrame string

const (
	StateVectorFullCovReferenceFrameJ2000  StateVectorFullCovReferenceFrame = "J2000"
	StateVectorFullCovReferenceFrameUvw    StateVectorFullCovReferenceFrame = "UVW"
	StateVectorFullCovReferenceFrameEfgTdr StateVectorFullCovReferenceFrame = "EFG/TDR"
	StateVectorFullCovReferenceFrameTeme   StateVectorFullCovReferenceFrame = "TEME"
	StateVectorFullCovReferenceFrameGcrf   StateVectorFullCovReferenceFrame = "GCRF"
)

// The reference frame of the cartesian orbital states. If the referenceFrame is
// null it is assumed to be J2000.
type StateVectorFullReferenceFrame string

const (
	StateVectorFullReferenceFrameJ2000   StateVectorFullReferenceFrame = "J2000"
	StateVectorFullReferenceFrameEfgTdr  StateVectorFullReferenceFrame = "EFG/TDR"
	StateVectorFullReferenceFrameEcrEcef StateVectorFullReferenceFrame = "ECR/ECEF"
	StateVectorFullReferenceFrameTeme    StateVectorFullReferenceFrame = "TEME"
	StateVectorFullReferenceFrameItrf    StateVectorFullReferenceFrame = "ITRF"
	StateVectorFullReferenceFrameGcrf    StateVectorFullReferenceFrame = "GCRF"
)

// This service provides operations for querying and manipulation of state vectors
// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
// velocity (v) that, together with their time (epoch) (t), uniquely determine the
// trajectory of the orbiting body in space. J2000 is the preferred coordinate
// frame for all state vector positions/velocities in UDL, but in some cases data
// may be in another frame depending on the provider and/or datatype. Please see
// the 'Discover' tab in the storefront to confirm coordinate frames by data
// provider.
//
// The properties ClassificationMarking, DataMode, Epoch, Source are required.
type StateVectorIngestParam struct {
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
	DataMode StateVectorIngestDataMode `json:"dataMode,omitzero,required"`
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision.
	Epoch time.Time `json:"epoch,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// The actual time span used for the OD of the object, expressed in days.
	ActualOdSpan param.Opt[float64] `json:"actualODSpan,omitzero"`
	// Optional algorithm used to produce this record.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// The reference frame of the alternate1 (Alt1) cartesian orbital state.
	Alt1ReferenceFrame param.Opt[string] `json:"alt1ReferenceFrame,omitzero"`
	// The reference frame of the alternate2 (Alt2) cartesian orbital state.
	Alt2ReferenceFrame param.Opt[string] `json:"alt2ReferenceFrame,omitzero"`
	// The actual area of the object at it's largest cross-section, expressed in
	// meters^2.
	Area param.Opt[float64] `json:"area,omitzero"`
	// First derivative of drag/ballistic coefficient (m2/kg-s).
	BDot param.Opt[float64] `json:"bDot,omitzero"`
	// Model parameter value for center of mass offset (m).
	CmOffset param.Opt[float64] `json:"cmOffset,omitzero"`
	// The method used to generate the covariance during the orbit determination (OD)
	// that produced the state vector, or whether an arbitrary, non-calculated default
	// value was used (CALCULATED, DEFAULT).
	CovMethod param.Opt[string] `json:"covMethod,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Optional source-provided and searchable metadata or descriptor of the data.
	Descriptor param.Opt[string] `json:"descriptor,omitzero"`
	// The effective area of the object exposed to atmospheric drag, expressed in
	// meters^2.
	DragArea param.Opt[float64] `json:"dragArea,omitzero"`
	// Area-to-mass ratio coefficient for atmospheric ballistic drag (m2/kg).
	DragCoeff param.Opt[float64] `json:"dragCoeff,omitzero"`
	// The Drag Model used for this vector (e.g. HARRIS-PRIESTER, JAC70, JBH09, MSIS90,
	// NONE, etc.).
	DragModel param.Opt[string] `json:"dragModel,omitzero"`
	// Model parameter value for energy dissipation rate (EDR) (w/kg).
	Edr param.Opt[float64] `json:"edr,omitzero"`
	// Integrator error control.
	ErrorControl param.Opt[float64] `json:"errorControl,omitzero"`
	// Boolean indicating use of fixed step size for this vector.
	FixedStep param.Opt[bool] `json:"fixedStep,omitzero"`
	// Geopotential model used for this vector (e.g. EGM-96, WGS-84, WGS-72, JGM-2, or
	// GEM-T3), including mm degree zonals, nn degree/order tesserals. E.g. EGM-96
	// 24Z,24T.
	GeopotentialModel param.Opt[string] `json:"geopotentialModel,omitzero"`
	// Number of terms used in the IAU 1980 nutation model (4, 50, or 106).
	Iau1980Terms param.Opt[int64] `json:"iau1980Terms,omitzero"`
	// Unique identifier of the satellite on-orbit object, if correlated. For the
	// public catalog, the idOnOrbit is typically the satellite number as a string, but
	// may be a UUID for analyst or other unknown or untracked satellites.
	IDOnOrbit param.Opt[string] `json:"idOnOrbit,omitzero"`
	// Unique identifier of the OD solution record that produced this state vector.
	// This ID can be used to obtain additional information on an OrbitDetermination
	// object using the 'get by ID' operation (e.g. /udl/orbitdetermination/{id}). For
	// example, the OrbitDetermination with idOrbitDetermination = abc would be queries
	// as /udl/orbitdetermination/abc.
	IDOrbitDetermination param.Opt[string] `json:"idOrbitDetermination,omitzero"`
	// Unique identifier of the record, auto-generated by the system.
	IDStateVector param.Opt[string] `json:"idStateVector,omitzero"`
	// Integrator Mode.
	IntegratorMode param.Opt[string] `json:"integratorMode,omitzero"`
	// Boolean indicating use of in-track thrust perturbations for this vector.
	InTrackThrust param.Opt[bool] `json:"inTrackThrust,omitzero"`
	// The end of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObEnd param.Opt[time.Time] `json:"lastObEnd,omitzero" format:"date-time"`
	// The start of the time interval containing the time of the last accepted
	// observation, in ISO 8601 UTC format with microsecond precision. For an exact
	// observation time, the firstObTime and lastObTime are the same.
	LastObStart param.Opt[time.Time] `json:"lastObStart,omitzero" format:"date-time"`
	// Time of the next leap second after epoch in ISO 8601 UTC time. If the next leap
	// second is not known, the time of the previous leap second is used.
	LeapSecondTime param.Opt[time.Time] `json:"leapSecondTime,omitzero" format:"date-time"`
	// Boolean indicating use of lunar/solar perturbations for this vector.
	LunarSolar param.Opt[bool] `json:"lunarSolar,omitzero"`
	// The mass of the object, in kilograms.
	Mass param.Opt[float64] `json:"mass,omitzero"`
	// The number of observations available for the OD of the object.
	ObsAvailable param.Opt[int64] `json:"obsAvailable,omitzero"`
	// The number of observations accepted for the OD of the object.
	ObsUsed param.Opt[int64] `json:"obsUsed,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Optional identifier provided by state vector source to indicate the target
	// onorbit object of this state vector. This may be an internal identifier and not
	// necessarily map to a valid satellite number.
	OrigObjectID param.Opt[string] `json:"origObjectId,omitzero"`
	// Type of partial derivatives used (ANALYTIC, FULL NUM, or FAST NUM).
	Partials param.Opt[string] `json:"partials,omitzero"`
	// The pedigree of state vector, or methods used for its generation to include
	// state update/orbit determination, propagation from another state, or a state
	// from a calibration satellite (e.g. ORBIT_UPDATE, PROPAGATION, CALIBRATION,
	// CONJUNCTION, FLIGHT_PLAN).
	Pedigree param.Opt[string] `json:"pedigree,omitzero"`
	// Polar Wander Motion X (arc seconds).
	PolarMotionX param.Opt[float64] `json:"polarMotionX,omitzero"`
	// Polar Wander Motion Y (arc seconds).
	PolarMotionY param.Opt[float64] `json:"polarMotionY,omitzero"`
	// One sigma position uncertainty, in kilometers.
	PosUnc param.Opt[float64] `json:"posUnc,omitzero"`
	// Optional URI location in the document repository of the raw file parsed by the
	// system to produce this record. To download the raw file, prepend
	// https://udl-hostname/scs/download?id= to this value.
	RawFileUri param.Opt[string] `json:"rawFileURI,omitzero"`
	// The recommended OD time span calculated for the object, expressed in days.
	RecOdSpan param.Opt[float64] `json:"recODSpan,omitzero"`
	// The percentage of residuals accepted in the OD of the object.
	ResidualsAcc param.Opt[float64] `json:"residualsAcc,omitzero"`
	// Epoch revolution number.
	RevNo param.Opt[int64] `json:"revNo,omitzero"`
	// The Weighted Root Mean Squared (RMS) of the differential correction on the
	// target object that produced this vector. WRMS is a quality indicator of the
	// state vector update, with a value of 1.00 being optimal. WRMS applies to Batch
	// Least Squares (BLS) processes.
	Rms param.Opt[float64] `json:"rms,omitzero"`
	// Satellite/Catalog number of the target OnOrbit object.
	SatNo param.Opt[int64] `json:"satNo,omitzero"`
	// Average solar flux geomagnetic index.
	SolarFluxApAvg param.Opt[float64] `json:"solarFluxAPAvg,omitzero"`
	// F10 (10.7 cm) solar flux value.
	SolarFluxF10 param.Opt[float64] `json:"solarFluxF10,omitzero"`
	// F10 (10.7 cm) solar flux 81-day average value.
	SolarFluxF10Avg param.Opt[float64] `json:"solarFluxF10Avg,omitzero"`
	// Boolean indicating use of solar radiation pressure perturbations for this
	// vector.
	SolarRadPress param.Opt[bool] `json:"solarRadPress,omitzero"`
	// Area-to-mass ratio coefficient for solar radiation pressure.
	SolarRadPressCoeff param.Opt[float64] `json:"solarRadPressCoeff,omitzero"`
	// Boolean indicating use of solid earth tide perturbations for this vector.
	SolidEarthTides param.Opt[bool] `json:"solidEarthTides,omitzero"`
	// The source data library from which this record was received. This could be a
	// remote or tactical UDL or another data library. If null, the record should be
	// assumed to have originated from the primary Enterprise UDL.
	SourceDl param.Opt[string] `json:"sourceDL,omitzero"`
	// The effective area of the object exposed to solar radiation pressure, expressed
	// in meters^2.
	SrpArea param.Opt[float64] `json:"srpArea,omitzero"`
	// Integrator step mode (AUTO, TIME, or S).
	StepMode param.Opt[string] `json:"stepMode,omitzero"`
	// Initial integration step size (seconds).
	StepSize param.Opt[float64] `json:"stepSize,omitzero"`
	// Initial step size selection (AUTO or MANUAL).
	StepSizeSelection param.Opt[string] `json:"stepSizeSelection,omitzero"`
	// TAI (Temps Atomique International) minus UTC (Universal Time Coordinates) offset
	// in seconds.
	TaiUtc param.Opt[float64] `json:"taiUtc,omitzero"`
	// Model parameter value for thrust acceleration (m/s2).
	ThrustAccel param.Opt[float64] `json:"thrustAccel,omitzero"`
	// The number of sensor tracks available for the OD of the object.
	TracksAvail param.Opt[int64] `json:"tracksAvail,omitzero"`
	// The number of sensor tracks accepted for the OD of the object.
	TracksUsed param.Opt[int64] `json:"tracksUsed,omitzero"`
	// Optional identifier to track a commercial or marketplace transaction executed to
	// produce this data.
	TransactionID param.Opt[string] `json:"transactionId,omitzero"`
	// Boolean indicating this state vector was unable to be correlated to a known
	// object. This flag should only be set to true by data providers after an attempt
	// to correlate to an OnOrbit object was made and failed. If unable to correlate,
	// the 'origObjectId' field may be populated with an internal data provider
	// specific identifier.
	Uct param.Opt[bool] `json:"uct,omitzero"`
	// Rate of change of UT1 (milliseconds/day) - first derivative of ut1Utc.
	Ut1Rate param.Opt[float64] `json:"ut1Rate,omitzero"`
	// Universal Time-1 (UT1) minus UTC offset, in seconds.
	Ut1Utc param.Opt[float64] `json:"ut1Utc,omitzero"`
	// One sigma velocity uncertainty, in kilometers/second.
	VelUnc param.Opt[float64] `json:"velUnc,omitzero"`
	// Cartesian X acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xaccel param.Opt[float64] `json:"xaccel,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xpos param.Opt[float64] `json:"xpos,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt1 param.Opt[float64] `json:"xposAlt1,omitzero"`
	// Cartesian X position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XposAlt2 param.Opt[float64] `json:"xposAlt2,omitzero"`
	// Cartesian X velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Xvel param.Opt[float64] `json:"xvel,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt1 param.Opt[float64] `json:"xvelAlt1,omitzero"`
	// Cartesian X velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	XvelAlt2 param.Opt[float64] `json:"xvelAlt2,omitzero"`
	// Cartesian Y acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yaccel param.Opt[float64] `json:"yaccel,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Ypos param.Opt[float64] `json:"ypos,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt1 param.Opt[float64] `json:"yposAlt1,omitzero"`
	// Cartesian Y position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YposAlt2 param.Opt[float64] `json:"yposAlt2,omitzero"`
	// Cartesian Y velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Yvel param.Opt[float64] `json:"yvel,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt1 param.Opt[float64] `json:"yvelAlt1,omitzero"`
	// Cartesian Y velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	YvelAlt2 param.Opt[float64] `json:"yvelAlt2,omitzero"`
	// Cartesian Z acceleration of target, in kilometers/second^2, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zaccel param.Opt[float64] `json:"zaccel,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zpos param.Opt[float64] `json:"zpos,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt1 param.Opt[float64] `json:"zposAlt1,omitzero"`
	// Cartesian Z position of the target, in kilometers, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZposAlt2 param.Opt[float64] `json:"zposAlt2,omitzero"`
	// Cartesian Z velocity of target, in kilometers/second, in the specified
	// referenceFrame. If referenceFrame is null then J2K should be assumed.
	Zvel param.Opt[float64] `json:"zvel,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt1ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt1 param.Opt[float64] `json:"zvelAlt1,omitzero"`
	// Cartesian Z velocity of the target, in kilometers/second, in the specified
	// alt2ReferenceFrame. Alternate reference frames are optional and are intended to
	// allow a data source to provide an equivalent vector in a different cartesian
	// frame than the primary vector.
	ZvelAlt2 param.Opt[float64] `json:"zvelAlt2,omitzero"`
	// Covariance matrix, in kilometer and second based units, in the specified
	// covReferenceFrame. If the covReferenceFrame is null it is assumed to be J2000.
	// The array values (1-21) represent the lower triangular half of the
	// position-velocity covariance matrix. The size of the covariance matrix is
	// dynamic, depending on whether the covariance for position only or position &
	// velocity. The covariance elements are position dependent within the array with
	// values ordered as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR&nbsp;&nbsp;
	//
	// x&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1
	//
	// y&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2&nbsp;&nbsp;&nbsp;&nbsp;3
	//
	// z&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4&nbsp;&nbsp;&nbsp;&nbsp;5&nbsp;&nbsp;&nbsp;&nbsp;6
	//
	// x'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7&nbsp;&nbsp;&nbsp;&nbsp;8&nbsp;&nbsp;&nbsp;&nbsp;9&nbsp;&nbsp;&nbsp;10
	//
	// y'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;11&nbsp;&nbsp;12&nbsp;&nbsp;13&nbsp;&nbsp;14&nbsp;&nbsp;15
	//
	// z'&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;16&nbsp;&nbsp;17&nbsp;&nbsp;18&nbsp;&nbsp;19&nbsp;&nbsp;20&nbsp;&nbsp;&nbsp;21&nbsp;
	//
	// The cov array should contain only the lower left triangle values from top left
	// down to bottom right, in order.
	//
	// If additional covariance terms are included for DRAG, SRP, and/or THRUST, the
	// matrix can be extended with the following order of elements:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;&nbsp;&nbsp;&nbsp;z&nbsp;&nbsp;&nbsp;&nbsp;x'&nbsp;&nbsp;&nbsp;&nbsp;y'&nbsp;&nbsp;&nbsp;&nbsp;z'&nbsp;&nbsp;&nbsp;&nbsp;DRG&nbsp;&nbsp;&nbsp;&nbsp;SRP&nbsp;&nbsp;&nbsp;&nbsp;THR
	//
	// DRG&nbsp;&nbsp;&nbsp;22&nbsp;&nbsp;23&nbsp;&nbsp;24&nbsp;&nbsp;25&nbsp;&nbsp;26&nbsp;&nbsp;&nbsp;27&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;28&nbsp;&nbsp;
	//
	// SRP&nbsp;&nbsp;&nbsp;29&nbsp;&nbsp;30&nbsp;&nbsp;31&nbsp;&nbsp;32&nbsp;&nbsp;33&nbsp;&nbsp;&nbsp;34&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;35&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;36&nbsp;&nbsp;
	//
	// THR&nbsp;&nbsp;&nbsp;37&nbsp;&nbsp;38&nbsp;&nbsp;39&nbsp;&nbsp;40&nbsp;&nbsp;41&nbsp;&nbsp;&nbsp;42&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;43&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;44&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;45&nbsp;
	Cov []float64 `json:"cov,omitzero"`
	// The reference frame of the covariance matrix elements. If the covReferenceFrame
	// is null it is assumed to be J2000.
	//
	// Any of "J2000", "UVW", "EFG/TDR", "TEME", "GCRF".
	CovReferenceFrame StateVectorIngestCovReferenceFrame `json:"covReferenceFrame,omitzero"`
	// The covariance matrix values represent the lower triangular half of the
	// covariance matrix in terms of equinoctial elements.&nbsp; The size of the
	// covariance matrix is dynamic.&nbsp; The values are outputted in order across
	// each row, i.e.:
	//
	// 1&nbsp;&nbsp; 2&nbsp;&nbsp; 3&nbsp;&nbsp; 4&nbsp;&nbsp; 5
	//
	// 6&nbsp;&nbsp; 7&nbsp;&nbsp; 8&nbsp;&nbsp; 9&nbsp; 10
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// 51&nbsp; 52&nbsp; 53&nbsp; 54&nbsp; 55
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :&nbsp;&nbsp; :
	//
	// The ordering of values is as follows:
	//
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Af&nbsp;&nbsp;
	// Ag&nbsp;&nbsp; L&nbsp;&nbsp;&nbsp; N&nbsp;&nbsp; Chi&nbsp; Psi&nbsp;&nbsp;
	// B&nbsp;&nbsp; BDOT AGOM&nbsp; T&nbsp;&nbsp; C1&nbsp;&nbsp; C2&nbsp; ...
	//
	// Af&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 1
	//
	// Ag&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2&nbsp;&nbsp;&nbsp; 3
	//
	// L&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 4&nbsp;&nbsp;&nbsp; 5&nbsp;&nbsp;&nbsp; 6
	//
	// N&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// 7&nbsp;&nbsp;&nbsp; 8&nbsp;&nbsp;&nbsp; 9&nbsp;&nbsp; 10
	//
	// Chi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 11&nbsp;&nbsp; 12&nbsp;&nbsp;
	// 13&nbsp;&nbsp; 14&nbsp;&nbsp; 15
	//
	// Psi&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 16&nbsp;&nbsp; 17&nbsp;&nbsp;
	// 18&nbsp;&nbsp; 19&nbsp;&nbsp; 20&nbsp;&nbsp; 21
	//
	// B&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 22&nbsp;&nbsp;
	// 23&nbsp;&nbsp; 24 &nbsp;&nbsp;25&nbsp;&nbsp; 26&nbsp;&nbsp; 27&nbsp;&nbsp; 28
	//
	// BDOT&nbsp;&nbsp; 29&nbsp;&nbsp; 30&nbsp;&nbsp; 31&nbsp;&nbsp; 32&nbsp;&nbsp;
	// 33&nbsp;&nbsp; 34&nbsp;&nbsp; 35&nbsp;&nbsp; 36
	//
	// AGOM&nbsp; 37&nbsp;&nbsp; 38&nbsp;&nbsp; 39&nbsp;&nbsp; 40&nbsp;&nbsp;
	// 41&nbsp;&nbsp; 42&nbsp;&nbsp; 43&nbsp;&nbsp; 44&nbsp;&nbsp; 45
	//
	// T&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 46&nbsp;&nbsp;
	// 47&nbsp;&nbsp; 48&nbsp;&nbsp; 49&nbsp;&nbsp; 50&nbsp;&nbsp; 51&nbsp;&nbsp;
	// 52&nbsp;&nbsp; 53&nbsp;&nbsp; 54&nbsp;&nbsp; 55
	//
	// C1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 56&nbsp;&nbsp; 57&nbsp;&nbsp;
	// 58&nbsp;&nbsp; 59&nbsp;&nbsp; 60&nbsp;&nbsp; 61&nbsp;&nbsp; 62&nbsp;&nbsp;
	// 63&nbsp;&nbsp; 64&nbsp;&nbsp; 65&nbsp;&nbsp; 66
	//
	// C2&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 67&nbsp;&nbsp; 68&nbsp;&nbsp;
	// 69&nbsp;&nbsp; 70&nbsp;&nbsp; 71&nbsp; &nbsp;72&nbsp;&nbsp; 73&nbsp;&nbsp;
	// 74&nbsp;&nbsp; 75&nbsp;&nbsp; 76&nbsp;&nbsp; 77&nbsp;&nbsp; 78
	//
	// :
	//
	// :
	//
	// where C1, C2, etc, are the "consider parameters" that may be added to the
	// covariance matrix.&nbsp; The covariance matrix will be as large as the last
	// element/model parameter needed.&nbsp; In other words, if the DC solved for all 6
	// elements plus AGOM, the covariance matrix will be 9x9 (and the rows for B and
	// BDOT will be all zeros).&nbsp; If the covariance matrix is unavailable, the size
	// will be set to 0x0, and no data will follow.&nbsp; The cov field should contain
	// only the lower left triangle values from top left down to bottom right, in
	// order.
	EqCov []float64 `json:"eqCov,omitzero"`
	// The reference frame of the cartesian orbital states. If the referenceFrame is
	// null it is assumed to be J2000.
	//
	// Any of "J2000", "EFG/TDR", "ECR/ECEF", "TEME", "ITRF", "GCRF".
	ReferenceFrame StateVectorIngestReferenceFrame `json:"referenceFrame,omitzero"`
	// Array containing the standard deviation of error in target object position, U, V
	// and W direction respectively (km).
	SigmaPosUvw []float64 `json:"sigmaPosUVW,omitzero"`
	// Array containing the standard deviation of error in target object velocity, U, V
	// and W direction respectively (km/sec).
	SigmaVelUvw []float64 `json:"sigmaVelUVW,omitzero"`
	// Optional array of UDL data (observation) UUIDs used to build this state vector.
	// See the associated sourcedDataTypes array for the specific types of observations
	// for the positionally corresponding UUIDs in this array (the two arrays must
	// match in size).
	SourcedData []string `json:"sourcedData,omitzero"`
	// Optional array of UDL observation data types used to build this state vector
	// (e.g. EO, RADAR, RF, DOA). See the associated sourcedData array for the specific
	// UUIDs of observations for the positionally corresponding data types in this
	// array (the two arrays must match in size).
	//
	// Any of "EO", "RADAR", "RF", "DOA", "ELSET", "SV".
	SourcedDataTypes []string `json:"sourcedDataTypes,omitzero"`
	// Optional array of provider/source specific tags for this data, where each
	// element is no longer than 32 characters, used for implementing data owner
	// conditional access controls to restrict access to the data. Should be left null
	// by data providers unless conditional access controls are coordinated with the
	// UDL team.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r StateVectorIngestParam) MarshalJSON() (data []byte, err error) {
	type shadow StateVectorIngestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StateVectorIngestParam) UnmarshalJSON(data []byte) error {
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
type StateVectorIngestDataMode string

const (
	StateVectorIngestDataModeReal      StateVectorIngestDataMode = "REAL"
	StateVectorIngestDataModeTest      StateVectorIngestDataMode = "TEST"
	StateVectorIngestDataModeSimulated StateVectorIngestDataMode = "SIMULATED"
	StateVectorIngestDataModeExercise  StateVectorIngestDataMode = "EXERCISE"
)

// The reference frame of the covariance matrix elements. If the covReferenceFrame
// is null it is assumed to be J2000.
type StateVectorIngestCovReferenceFrame string

const (
	StateVectorIngestCovReferenceFrameJ2000  StateVectorIngestCovReferenceFrame = "J2000"
	StateVectorIngestCovReferenceFrameUvw    StateVectorIngestCovReferenceFrame = "UVW"
	StateVectorIngestCovReferenceFrameEfgTdr StateVectorIngestCovReferenceFrame = "EFG/TDR"
	StateVectorIngestCovReferenceFrameTeme   StateVectorIngestCovReferenceFrame = "TEME"
	StateVectorIngestCovReferenceFrameGcrf   StateVectorIngestCovReferenceFrame = "GCRF"
)

// The reference frame of the cartesian orbital states. If the referenceFrame is
// null it is assumed to be J2000.
type StateVectorIngestReferenceFrame string

const (
	StateVectorIngestReferenceFrameJ2000   StateVectorIngestReferenceFrame = "J2000"
	StateVectorIngestReferenceFrameEfgTdr  StateVectorIngestReferenceFrame = "EFG/TDR"
	StateVectorIngestReferenceFrameEcrEcef StateVectorIngestReferenceFrame = "ECR/ECEF"
	StateVectorIngestReferenceFrameTeme    StateVectorIngestReferenceFrame = "TEME"
	StateVectorIngestReferenceFrameItrf    StateVectorIngestReferenceFrame = "ITRF"
	StateVectorIngestReferenceFrameGcrf    StateVectorIngestReferenceFrame = "GCRF"
)

type StateVectorQueryhelpResponse struct {
	AodrSupported         bool                                    `json:"aodrSupported"`
	ClassificationMarking string                                  `json:"classificationMarking"`
	Description           string                                  `json:"description"`
	HistorySupported      bool                                    `json:"historySupported"`
	Name                  string                                  `json:"name"`
	Parameters            []StateVectorQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                                `json:"requiredRoles"`
	RestSupported         bool                                    `json:"restSupported"`
	SortSupported         bool                                    `json:"sortSupported"`
	TypeName              string                                  `json:"typeName"`
	Uri                   string                                  `json:"uri"`
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
func (r StateVectorQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *StateVectorQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StateVectorQueryhelpResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StateVectorQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *StateVectorQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StateVectorNewParams struct {
	// This service provides operations for querying and manipulation of state vectors
	// for OnOrbit objects. State vectors are cartesian vectors of position (r) and
	// velocity (v) that, together with their time (epoch) (t), uniquely determine the
	// trajectory of the orbiting body in space. J2000 is the preferred coordinate
	// frame for all state vector positions/velocities in UDL, but in some cases data
	// may be in another frame depending on the provider and/or datatype. Please see
	// the 'Discover' tab in the storefront to confirm coordinate frames by data
	// provider.
	StateVectorIngest StateVectorIngestParam
	paramObj
}

func (r StateVectorNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.StateVectorIngest)
}
func (r *StateVectorNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.StateVectorIngest)
}

type StateVectorListParams struct {
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorListParams]'s query parameters as `url.Values`.
func (r StateVectorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StateVectorCountParams struct {
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorCountParams]'s query parameters as `url.Values`.
func (r StateVectorCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StateVectorNewBulkParams struct {
	Body []StateVectorIngestParam
	paramObj
}

func (r StateVectorNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *StateVectorNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type StateVectorGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorGetParams]'s query parameters as `url.Values`.
func (r StateVectorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StateVectorTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Time of validity for state vector in ISO 8601 UTC datetime format, with
	// microsecond precision. (YYYY-MM-DDTHH:MM:SS.ssssssZ)
	Epoch       time.Time        `query:"epoch,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StateVectorTupleParams]'s query parameters as `url.Values`.
func (r StateVectorTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StateVectorUnvalidatedPublishParams struct {
	Body []StateVectorIngestParam
	paramObj
}

func (r StateVectorUnvalidatedPublishParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *StateVectorUnvalidatedPublishParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
