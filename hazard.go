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

// HazardService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHazardService] method instead.
type HazardService struct {
	Options []option.RequestOption
	History HazardHistoryService
}

// NewHazardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHazardService(opts ...option.RequestOption) (r HazardService) {
	r = HazardService{}
	r.Options = opts
	r.History = NewHazardHistoryService(opts...)
	return
}

// Service operation to take a single hazard as a POST body and ingest into the
// database. This operation is not intended to be used for automated feeds into
// UDL. Data providers should contact the UDL team for specific role assignments
// and for instructions on setting up a permanent feed through an alternate
// mechanism.
func (r *HazardService) New(ctx context.Context, body HazardNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/hazard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *HazardService) List(ctx context.Context, query HazardListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[HazardListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/hazard"
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
func (r *HazardService) ListAutoPaging(ctx context.Context, query HazardListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[HazardListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *HazardService) Count(ctx context.Context, query HazardCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/hazard/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// hazards as a POST body and ingest into the database. This operation is not
// intended to be used for automated feeds into UDL. Data providers should contact
// the UDL team for specific role assignments and for instructions on setting up a
// permanent feed through an alternate mechanism.
func (r *HazardService) NewBulk(ctx context.Context, body HazardNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/hazard/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to get a single Hazard by its unique ID passed as a path
// parameter.
func (r *HazardService) Get(ctx context.Context, id string, query HazardGetParams, opts ...option.RequestOption) (res *HazardGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("udl/hazard/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *HazardService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/hazard/queryhelp"
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
func (r *HazardService) Tuple(ctx context.Context, query HazardTupleParams, opts ...option.RequestOption) (res *[]HazardTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/hazard/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Hazard contains information concerning the exposure of a geographic area to a
// Chemical, Biological, Radiological, or Nuclear (CBRN) contaminant. The Hazard
// schema includes the detection time and type of contamination as well as optional
// information regarding specific material properties, the extent of contamination,
// and identification of affected regions.
type HazardListResponse struct {
	// Array of the specific alarms associated with this detection. The alarms and
	// alarmValues arrays must contain the same number of elements.
	Alarms []string `json:"alarms,required"`
	// Array of the values that correspond to each of the alarms contained in alarms.
	// The alarms and alarmValues arrays must contain the same number of elements.
	AlarmValues []float64 `json:"alarmValues,required"`
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
	DataMode HazardListResponseDataMode `json:"dataMode,required"`
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	DetectTime time.Time `json:"detectTime,required" format:"date-time"`
	// The type of hazard (Chemical, Biological, Radiological, Nuclear) detect
	// associated with this record.
	DetectType string `json:"detectType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The (rounded) Mass Number of the material associated with this detection. The
	// rounded value is the mass number of the most abundant isotope of the element.
	A int64 `json:"a"`
	// The radioactivity measurement of the material associated with this detection, in
	// becquerels (Bq). One becquerel (Bq) is equal to one nuclear decay per second.
	Activity float64 `json:"activity"`
	// The specific bottle ID associated with this detection.
	BottleID string `json:"bottleId"`
	// The CAS Registry Number, also referred to as CAS Number or CAS RN, associated
	// with the this detection. The CAS Number is a unique numerical identifier
	// assigned by the Chemical Abstracts Service (CAS), to every chemical substance
	// described in the open scientific literature. It includes organic and inorganic
	// compounds, minerals, isotopes, alloys, mixtures, and nonstructurable materials
	// (UVCBs, substances of unknown or variable composition, complex reaction
	// products, or biological origin). For further information, reference
	// https://www.cas.org/cas-data/cas-registry.
	CasRn string `json:"casRN"`
	// The applicable channel involved in this biological material detection (e.g.
	// Digestive, Eyes, Respiratory, Skin, etc.) .
	Channel string `json:"channel"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The concentration time, in (kg/sec)/m^3, associated with this material
	// detection.
	CtrnTime float64 `json:"ctrnTime"`
	// Measure of density of the material associated with this detection, in kg/m^3.
	Density float64 `json:"density"`
	// The deposition measurement of the contaminant to surface area, in kg/m^2.
	Dep float64 `json:"dep"`
	// The deposition concentration of the contaminant to surface area, in
	// becquerels/m^2.
	DepCtrn float64 `json:"depCtrn"`
	// The dose associated with this detection, in gray. Dose is the measure of the
	// energy deposited in matter by ionizing radiation per unit mass. One gray is
	// defined as one Joule of energy absorbed per kilogram of matter.
	Dose float64 `json:"dose"`
	// The dose rate associated with this detection, in gray/sec. One gray is defined
	// as one Joule of energy absorbed per kilogram of matter.
	DoseRate float64 `json:"doseRate"`
	// The known or projected hazard duration, in seconds, associated with this
	// material detection.
	Duration int64 `json:"duration"`
	// Chemical Agent Monitor (CAM) G-type agent measurement, in number of display
	// bars. In G-mode, CAMs monitor for G-series nerve agents.
	GBar float64 `json:"gBar"`
	// Flag indicating whether this detection is harmful to humans.
	Harmful bool `json:"harmful"`
	// Chemical Agent Monitor (CAM) H-type agent measurement, in number of display
	// bars. In H-mode, CAMs monitor for blister agents.
	HBar float64 `json:"hBar"`
	// ID of the Point of Interest (POI) record related to this hazard record.
	IDPoi string `json:"idPOI"`
	// ID of the Track record related to this hazard record.
	IDTrack string `json:"idTrack"`
	// Ratio of the chemical substance mass to the total mass of the mixture.
	MassFrac float64 `json:"massFrac"`
	// The Radiological Category (1 - 5) which applies to the material associated with
	// this detection, according to the following definitions:
	//
	// Category 1: If not safely or securely managed, would be likely to cause
	// permanent injury to a person who handled them or was otherwise in contact with
	// them for more than a few minutes. It would probably be fatal to be close to this
	// amount of unshielded material for a period of a few minutes to an hour.
	//
	// Category 2: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for a short time
	// (minutes to hours). It could possibly be fatal to be close to this amount of
	// unshielded radioactive material for a period of hours to days.
	//
	// Category 3: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for hours. It
	// could possibly - although it is unlikely to be - fatal to be close to this
	// amount of unshielded radioactive material for a period of days to weeks.
	//
	// Category 4: If not safely managed or securely protected, could possibly cause
	// temporary injury to someone who handled them or was otherwise in contact with or
	// close to them for a period of many weeks, though this is unlikely. It is very
	// unlikely anyone would be permanently injured by this amount of radioactive
	// material.
	//
	// Category 5: Cannot cause permanent injury. This category applies to x-ray
	// fluorescence devices and electron capture devices.
	MatCat int64 `json:"matCat"`
	// The specific Material Class for the material associated with this detect. The
	// material class is generally associated with chemical and biological detections.
	MatClass string `json:"matClass"`
	// The material common name associated with this detection.
	MatName string `json:"matName"`
	// The specific material type (MT) or MT Code involved in this detection, when
	// applicable. The material type is generally associated with radiological and/or
	// nuclear detections. For further information, reference Nuclear Materials
	// Management and Safeguards System (NMMSS) Users Guide Rev. 2.1.
	MatType string `json:"matType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Measure of the concentration of the material associated with this detection, in
	// parts per million (units of contaminant mass per million parts of total mass).
	Ppm int64 `json:"ppm"`
	// Measure of radioactive concentration of the material associated with this
	// detection, in becquerels/m^3. One becquerel (Bq) is equal to one nuclear decay
	// per second.
	RadCtrn float64 `json:"radCtrn"`
	// Array of the specific readings associated with this detection. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	Readings []string `json:"readings"`
	// Array of the units that correspond to each of the readingValues. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	ReadingUnits []string `json:"readingUnits"`
	// Array of the values that correspond to each of the readings contained in
	// readings. The readings, readingUnits, and readingValues arrays must contain the
	// same number of elements.
	ReadingValues []float64 `json:"readingValues"`
	// The Atomic Number of the material associated with this detection.
	Z int64 `json:"z"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Alarms                resp.Field
		AlarmValues           resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		DetectTime            resp.Field
		DetectType            resp.Field
		Source                resp.Field
		ID                    resp.Field
		A                     resp.Field
		Activity              resp.Field
		BottleID              resp.Field
		CasRn                 resp.Field
		Channel               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CtrnTime              resp.Field
		Density               resp.Field
		Dep                   resp.Field
		DepCtrn               resp.Field
		Dose                  resp.Field
		DoseRate              resp.Field
		Duration              resp.Field
		GBar                  resp.Field
		Harmful               resp.Field
		HBar                  resp.Field
		IDPoi                 resp.Field
		IDTrack               resp.Field
		MassFrac              resp.Field
		MatCat                resp.Field
		MatClass              resp.Field
		MatName               resp.Field
		MatType               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Ppm                   resp.Field
		RadCtrn               resp.Field
		Readings              resp.Field
		ReadingUnits          resp.Field
		ReadingValues         resp.Field
		Z                     resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HazardListResponse) RawJSON() string { return r.JSON.raw }
func (r *HazardListResponse) UnmarshalJSON(data []byte) error {
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
type HazardListResponseDataMode string

const (
	HazardListResponseDataModeReal      HazardListResponseDataMode = "REAL"
	HazardListResponseDataModeTest      HazardListResponseDataMode = "TEST"
	HazardListResponseDataModeSimulated HazardListResponseDataMode = "SIMULATED"
	HazardListResponseDataModeExercise  HazardListResponseDataMode = "EXERCISE"
)

// Hazard contains information concerning the exposure of a geographic area to a
// Chemical, Biological, Radiological, or Nuclear (CBRN) contaminant. The Hazard
// schema includes the detection time and type of contamination as well as optional
// information regarding specific material properties, the extent of contamination,
// and identification of affected regions.
type HazardGetResponse struct {
	// Array of the specific alarms associated with this detection. The alarms and
	// alarmValues arrays must contain the same number of elements.
	Alarms []string `json:"alarms,required"`
	// Array of the values that correspond to each of the alarms contained in alarms.
	// The alarms and alarmValues arrays must contain the same number of elements.
	AlarmValues []float64 `json:"alarmValues,required"`
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
	DataMode HazardGetResponseDataMode `json:"dataMode,required"`
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	DetectTime time.Time `json:"detectTime,required" format:"date-time"`
	// The type of hazard (Chemical, Biological, Radiological, Nuclear) detect
	// associated with this record.
	DetectType string `json:"detectType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The (rounded) Mass Number of the material associated with this detection. The
	// rounded value is the mass number of the most abundant isotope of the element.
	A int64 `json:"a"`
	// The radioactivity measurement of the material associated with this detection, in
	// becquerels (Bq). One becquerel (Bq) is equal to one nuclear decay per second.
	Activity float64 `json:"activity"`
	// The specific bottle ID associated with this detection.
	BottleID string `json:"bottleId"`
	// The CAS Registry Number, also referred to as CAS Number or CAS RN, associated
	// with the this detection. The CAS Number is a unique numerical identifier
	// assigned by the Chemical Abstracts Service (CAS), to every chemical substance
	// described in the open scientific literature. It includes organic and inorganic
	// compounds, minerals, isotopes, alloys, mixtures, and nonstructurable materials
	// (UVCBs, substances of unknown or variable composition, complex reaction
	// products, or biological origin). For further information, reference
	// https://www.cas.org/cas-data/cas-registry.
	CasRn string `json:"casRN"`
	// The applicable channel involved in this biological material detection (e.g.
	// Digestive, Eyes, Respiratory, Skin, etc.) .
	Channel string `json:"channel"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The concentration time, in (kg/sec)/m^3, associated with this material
	// detection.
	CtrnTime float64 `json:"ctrnTime"`
	// Measure of density of the material associated with this detection, in kg/m^3.
	Density float64 `json:"density"`
	// The deposition measurement of the contaminant to surface area, in kg/m^2.
	Dep float64 `json:"dep"`
	// The deposition concentration of the contaminant to surface area, in
	// becquerels/m^2.
	DepCtrn float64 `json:"depCtrn"`
	// The dose associated with this detection, in gray. Dose is the measure of the
	// energy deposited in matter by ionizing radiation per unit mass. One gray is
	// defined as one Joule of energy absorbed per kilogram of matter.
	Dose float64 `json:"dose"`
	// The dose rate associated with this detection, in gray/sec. One gray is defined
	// as one Joule of energy absorbed per kilogram of matter.
	DoseRate float64 `json:"doseRate"`
	// The known or projected hazard duration, in seconds, associated with this
	// material detection.
	Duration int64 `json:"duration"`
	// Chemical Agent Monitor (CAM) G-type agent measurement, in number of display
	// bars. In G-mode, CAMs monitor for G-series nerve agents.
	GBar float64 `json:"gBar"`
	// Flag indicating whether this detection is harmful to humans.
	Harmful bool `json:"harmful"`
	// Chemical Agent Monitor (CAM) H-type agent measurement, in number of display
	// bars. In H-mode, CAMs monitor for blister agents.
	HBar float64 `json:"hBar"`
	// ID of the Point of Interest (POI) record related to this hazard record.
	IDPoi string `json:"idPOI"`
	// ID of the Track record related to this hazard record.
	IDTrack string `json:"idTrack"`
	// Ratio of the chemical substance mass to the total mass of the mixture.
	MassFrac float64 `json:"massFrac"`
	// The Radiological Category (1 - 5) which applies to the material associated with
	// this detection, according to the following definitions:
	//
	// Category 1: If not safely or securely managed, would be likely to cause
	// permanent injury to a person who handled them or was otherwise in contact with
	// them for more than a few minutes. It would probably be fatal to be close to this
	// amount of unshielded material for a period of a few minutes to an hour.
	//
	// Category 2: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for a short time
	// (minutes to hours). It could possibly be fatal to be close to this amount of
	// unshielded radioactive material for a period of hours to days.
	//
	// Category 3: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for hours. It
	// could possibly - although it is unlikely to be - fatal to be close to this
	// amount of unshielded radioactive material for a period of days to weeks.
	//
	// Category 4: If not safely managed or securely protected, could possibly cause
	// temporary injury to someone who handled them or was otherwise in contact with or
	// close to them for a period of many weeks, though this is unlikely. It is very
	// unlikely anyone would be permanently injured by this amount of radioactive
	// material.
	//
	// Category 5: Cannot cause permanent injury. This category applies to x-ray
	// fluorescence devices and electron capture devices.
	MatCat int64 `json:"matCat"`
	// The specific Material Class for the material associated with this detect. The
	// material class is generally associated with chemical and biological detections.
	MatClass string `json:"matClass"`
	// The material common name associated with this detection.
	MatName string `json:"matName"`
	// The specific material type (MT) or MT Code involved in this detection, when
	// applicable. The material type is generally associated with radiological and/or
	// nuclear detections. For further information, reference Nuclear Materials
	// Management and Safeguards System (NMMSS) Users Guide Rev. 2.1.
	MatType string `json:"matType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Measure of the concentration of the material associated with this detection, in
	// parts per million (units of contaminant mass per million parts of total mass).
	Ppm int64 `json:"ppm"`
	// Measure of radioactive concentration of the material associated with this
	// detection, in becquerels/m^3. One becquerel (Bq) is equal to one nuclear decay
	// per second.
	RadCtrn float64 `json:"radCtrn"`
	// Array of the specific readings associated with this detection. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	Readings []string `json:"readings"`
	// Array of the units that correspond to each of the readingValues. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	ReadingUnits []string `json:"readingUnits"`
	// Array of the values that correspond to each of the readings contained in
	// readings. The readings, readingUnits, and readingValues arrays must contain the
	// same number of elements.
	ReadingValues []float64 `json:"readingValues"`
	// The Atomic Number of the material associated with this detection.
	Z int64 `json:"z"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Alarms                resp.Field
		AlarmValues           resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		DetectTime            resp.Field
		DetectType            resp.Field
		Source                resp.Field
		ID                    resp.Field
		A                     resp.Field
		Activity              resp.Field
		BottleID              resp.Field
		CasRn                 resp.Field
		Channel               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CtrnTime              resp.Field
		Density               resp.Field
		Dep                   resp.Field
		DepCtrn               resp.Field
		Dose                  resp.Field
		DoseRate              resp.Field
		Duration              resp.Field
		GBar                  resp.Field
		Harmful               resp.Field
		HBar                  resp.Field
		IDPoi                 resp.Field
		IDTrack               resp.Field
		MassFrac              resp.Field
		MatCat                resp.Field
		MatClass              resp.Field
		MatName               resp.Field
		MatType               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Ppm                   resp.Field
		RadCtrn               resp.Field
		Readings              resp.Field
		ReadingUnits          resp.Field
		ReadingValues         resp.Field
		Z                     resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HazardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *HazardGetResponse) UnmarshalJSON(data []byte) error {
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
type HazardGetResponseDataMode string

const (
	HazardGetResponseDataModeReal      HazardGetResponseDataMode = "REAL"
	HazardGetResponseDataModeTest      HazardGetResponseDataMode = "TEST"
	HazardGetResponseDataModeSimulated HazardGetResponseDataMode = "SIMULATED"
	HazardGetResponseDataModeExercise  HazardGetResponseDataMode = "EXERCISE"
)

// Hazard contains information concerning the exposure of a geographic area to a
// Chemical, Biological, Radiological, or Nuclear (CBRN) contaminant. The Hazard
// schema includes the detection time and type of contamination as well as optional
// information regarding specific material properties, the extent of contamination,
// and identification of affected regions.
type HazardTupleResponse struct {
	// Array of the specific alarms associated with this detection. The alarms and
	// alarmValues arrays must contain the same number of elements.
	Alarms []string `json:"alarms,required"`
	// Array of the values that correspond to each of the alarms contained in alarms.
	// The alarms and alarmValues arrays must contain the same number of elements.
	AlarmValues []float64 `json:"alarmValues,required"`
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
	DataMode HazardTupleResponseDataMode `json:"dataMode,required"`
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	DetectTime time.Time `json:"detectTime,required" format:"date-time"`
	// The type of hazard (Chemical, Biological, Radiological, Nuclear) detect
	// associated with this record.
	DetectType string `json:"detectType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// The (rounded) Mass Number of the material associated with this detection. The
	// rounded value is the mass number of the most abundant isotope of the element.
	A int64 `json:"a"`
	// The radioactivity measurement of the material associated with this detection, in
	// becquerels (Bq). One becquerel (Bq) is equal to one nuclear decay per second.
	Activity float64 `json:"activity"`
	// The specific bottle ID associated with this detection.
	BottleID string `json:"bottleId"`
	// The CAS Registry Number, also referred to as CAS Number or CAS RN, associated
	// with the this detection. The CAS Number is a unique numerical identifier
	// assigned by the Chemical Abstracts Service (CAS), to every chemical substance
	// described in the open scientific literature. It includes organic and inorganic
	// compounds, minerals, isotopes, alloys, mixtures, and nonstructurable materials
	// (UVCBs, substances of unknown or variable composition, complex reaction
	// products, or biological origin). For further information, reference
	// https://www.cas.org/cas-data/cas-registry.
	CasRn string `json:"casRN"`
	// The applicable channel involved in this biological material detection (e.g.
	// Digestive, Eyes, Respiratory, Skin, etc.) .
	Channel string `json:"channel"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The concentration time, in (kg/sec)/m^3, associated with this material
	// detection.
	CtrnTime float64 `json:"ctrnTime"`
	// Measure of density of the material associated with this detection, in kg/m^3.
	Density float64 `json:"density"`
	// The deposition measurement of the contaminant to surface area, in kg/m^2.
	Dep float64 `json:"dep"`
	// The deposition concentration of the contaminant to surface area, in
	// becquerels/m^2.
	DepCtrn float64 `json:"depCtrn"`
	// The dose associated with this detection, in gray. Dose is the measure of the
	// energy deposited in matter by ionizing radiation per unit mass. One gray is
	// defined as one Joule of energy absorbed per kilogram of matter.
	Dose float64 `json:"dose"`
	// The dose rate associated with this detection, in gray/sec. One gray is defined
	// as one Joule of energy absorbed per kilogram of matter.
	DoseRate float64 `json:"doseRate"`
	// The known or projected hazard duration, in seconds, associated with this
	// material detection.
	Duration int64 `json:"duration"`
	// Chemical Agent Monitor (CAM) G-type agent measurement, in number of display
	// bars. In G-mode, CAMs monitor for G-series nerve agents.
	GBar float64 `json:"gBar"`
	// Flag indicating whether this detection is harmful to humans.
	Harmful bool `json:"harmful"`
	// Chemical Agent Monitor (CAM) H-type agent measurement, in number of display
	// bars. In H-mode, CAMs monitor for blister agents.
	HBar float64 `json:"hBar"`
	// ID of the Point of Interest (POI) record related to this hazard record.
	IDPoi string `json:"idPOI"`
	// ID of the Track record related to this hazard record.
	IDTrack string `json:"idTrack"`
	// Ratio of the chemical substance mass to the total mass of the mixture.
	MassFrac float64 `json:"massFrac"`
	// The Radiological Category (1 - 5) which applies to the material associated with
	// this detection, according to the following definitions:
	//
	// Category 1: If not safely or securely managed, would be likely to cause
	// permanent injury to a person who handled them or was otherwise in contact with
	// them for more than a few minutes. It would probably be fatal to be close to this
	// amount of unshielded material for a period of a few minutes to an hour.
	//
	// Category 2: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for a short time
	// (minutes to hours). It could possibly be fatal to be close to this amount of
	// unshielded radioactive material for a period of hours to days.
	//
	// Category 3: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for hours. It
	// could possibly - although it is unlikely to be - fatal to be close to this
	// amount of unshielded radioactive material for a period of days to weeks.
	//
	// Category 4: If not safely managed or securely protected, could possibly cause
	// temporary injury to someone who handled them or was otherwise in contact with or
	// close to them for a period of many weeks, though this is unlikely. It is very
	// unlikely anyone would be permanently injured by this amount of radioactive
	// material.
	//
	// Category 5: Cannot cause permanent injury. This category applies to x-ray
	// fluorescence devices and electron capture devices.
	MatCat int64 `json:"matCat"`
	// The specific Material Class for the material associated with this detect. The
	// material class is generally associated with chemical and biological detections.
	MatClass string `json:"matClass"`
	// The material common name associated with this detection.
	MatName string `json:"matName"`
	// The specific material type (MT) or MT Code involved in this detection, when
	// applicable. The material type is generally associated with radiological and/or
	// nuclear detections. For further information, reference Nuclear Materials
	// Management and Safeguards System (NMMSS) Users Guide Rev. 2.1.
	MatType string `json:"matType"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// Measure of the concentration of the material associated with this detection, in
	// parts per million (units of contaminant mass per million parts of total mass).
	Ppm int64 `json:"ppm"`
	// Measure of radioactive concentration of the material associated with this
	// detection, in becquerels/m^3. One becquerel (Bq) is equal to one nuclear decay
	// per second.
	RadCtrn float64 `json:"radCtrn"`
	// Array of the specific readings associated with this detection. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	Readings []string `json:"readings"`
	// Array of the units that correspond to each of the readingValues. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	ReadingUnits []string `json:"readingUnits"`
	// Array of the values that correspond to each of the readings contained in
	// readings. The readings, readingUnits, and readingValues arrays must contain the
	// same number of elements.
	ReadingValues []float64 `json:"readingValues"`
	// The Atomic Number of the material associated with this detection.
	Z int64 `json:"z"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Alarms                resp.Field
		AlarmValues           resp.Field
		ClassificationMarking resp.Field
		DataMode              resp.Field
		DetectTime            resp.Field
		DetectType            resp.Field
		Source                resp.Field
		ID                    resp.Field
		A                     resp.Field
		Activity              resp.Field
		BottleID              resp.Field
		CasRn                 resp.Field
		Channel               resp.Field
		CreatedAt             resp.Field
		CreatedBy             resp.Field
		CtrnTime              resp.Field
		Density               resp.Field
		Dep                   resp.Field
		DepCtrn               resp.Field
		Dose                  resp.Field
		DoseRate              resp.Field
		Duration              resp.Field
		GBar                  resp.Field
		Harmful               resp.Field
		HBar                  resp.Field
		IDPoi                 resp.Field
		IDTrack               resp.Field
		MassFrac              resp.Field
		MatCat                resp.Field
		MatClass              resp.Field
		MatName               resp.Field
		MatType               resp.Field
		Origin                resp.Field
		OrigNetwork           resp.Field
		Ppm                   resp.Field
		RadCtrn               resp.Field
		Readings              resp.Field
		ReadingUnits          resp.Field
		ReadingValues         resp.Field
		Z                     resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HazardTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *HazardTupleResponse) UnmarshalJSON(data []byte) error {
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
type HazardTupleResponseDataMode string

const (
	HazardTupleResponseDataModeReal      HazardTupleResponseDataMode = "REAL"
	HazardTupleResponseDataModeTest      HazardTupleResponseDataMode = "TEST"
	HazardTupleResponseDataModeSimulated HazardTupleResponseDataMode = "SIMULATED"
	HazardTupleResponseDataModeExercise  HazardTupleResponseDataMode = "EXERCISE"
)

type HazardNewParams struct {
	// Array of the specific alarms associated with this detection. The alarms and
	// alarmValues arrays must contain the same number of elements.
	Alarms []string `json:"alarms,omitzero,required"`
	// Array of the values that correspond to each of the alarms contained in alarms.
	// The alarms and alarmValues arrays must contain the same number of elements.
	AlarmValues []float64 `json:"alarmValues,omitzero,required"`
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
	DataMode HazardNewParamsDataMode `json:"dataMode,omitzero,required"`
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	DetectTime time.Time `json:"detectTime,required" format:"date-time"`
	// The type of hazard (Chemical, Biological, Radiological, Nuclear) detect
	// associated with this record.
	DetectType string `json:"detectType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The (rounded) Mass Number of the material associated with this detection. The
	// rounded value is the mass number of the most abundant isotope of the element.
	A param.Opt[int64] `json:"a,omitzero"`
	// The radioactivity measurement of the material associated with this detection, in
	// becquerels (Bq). One becquerel (Bq) is equal to one nuclear decay per second.
	Activity param.Opt[float64] `json:"activity,omitzero"`
	// The specific bottle ID associated with this detection.
	BottleID param.Opt[string] `json:"bottleId,omitzero"`
	// The CAS Registry Number, also referred to as CAS Number or CAS RN, associated
	// with the this detection. The CAS Number is a unique numerical identifier
	// assigned by the Chemical Abstracts Service (CAS), to every chemical substance
	// described in the open scientific literature. It includes organic and inorganic
	// compounds, minerals, isotopes, alloys, mixtures, and nonstructurable materials
	// (UVCBs, substances of unknown or variable composition, complex reaction
	// products, or biological origin). For further information, reference
	// https://www.cas.org/cas-data/cas-registry.
	CasRn param.Opt[string] `json:"casRN,omitzero"`
	// The applicable channel involved in this biological material detection (e.g.
	// Digestive, Eyes, Respiratory, Skin, etc.) .
	Channel param.Opt[string] `json:"channel,omitzero"`
	// The concentration time, in (kg/sec)/m^3, associated with this material
	// detection.
	CtrnTime param.Opt[float64] `json:"ctrnTime,omitzero"`
	// Measure of density of the material associated with this detection, in kg/m^3.
	Density param.Opt[float64] `json:"density,omitzero"`
	// The deposition measurement of the contaminant to surface area, in kg/m^2.
	Dep param.Opt[float64] `json:"dep,omitzero"`
	// The deposition concentration of the contaminant to surface area, in
	// becquerels/m^2.
	DepCtrn param.Opt[float64] `json:"depCtrn,omitzero"`
	// The dose associated with this detection, in gray. Dose is the measure of the
	// energy deposited in matter by ionizing radiation per unit mass. One gray is
	// defined as one Joule of energy absorbed per kilogram of matter.
	Dose param.Opt[float64] `json:"dose,omitzero"`
	// The dose rate associated with this detection, in gray/sec. One gray is defined
	// as one Joule of energy absorbed per kilogram of matter.
	DoseRate param.Opt[float64] `json:"doseRate,omitzero"`
	// The known or projected hazard duration, in seconds, associated with this
	// material detection.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Chemical Agent Monitor (CAM) G-type agent measurement, in number of display
	// bars. In G-mode, CAMs monitor for G-series nerve agents.
	GBar param.Opt[float64] `json:"gBar,omitzero"`
	// Flag indicating whether this detection is harmful to humans.
	Harmful param.Opt[bool] `json:"harmful,omitzero"`
	// Chemical Agent Monitor (CAM) H-type agent measurement, in number of display
	// bars. In H-mode, CAMs monitor for blister agents.
	HBar param.Opt[float64] `json:"hBar,omitzero"`
	// ID of the Point of Interest (POI) record related to this hazard record.
	IDPoi param.Opt[string] `json:"idPOI,omitzero"`
	// ID of the Track record related to this hazard record.
	IDTrack param.Opt[string] `json:"idTrack,omitzero"`
	// Ratio of the chemical substance mass to the total mass of the mixture.
	MassFrac param.Opt[float64] `json:"massFrac,omitzero"`
	// The Radiological Category (1 - 5) which applies to the material associated with
	// this detection, according to the following definitions:
	//
	// Category 1: If not safely or securely managed, would be likely to cause
	// permanent injury to a person who handled them or was otherwise in contact with
	// them for more than a few minutes. It would probably be fatal to be close to this
	// amount of unshielded material for a period of a few minutes to an hour.
	//
	// Category 2: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for a short time
	// (minutes to hours). It could possibly be fatal to be close to this amount of
	// unshielded radioactive material for a period of hours to days.
	//
	// Category 3: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for hours. It
	// could possibly - although it is unlikely to be - fatal to be close to this
	// amount of unshielded radioactive material for a period of days to weeks.
	//
	// Category 4: If not safely managed or securely protected, could possibly cause
	// temporary injury to someone who handled them or was otherwise in contact with or
	// close to them for a period of many weeks, though this is unlikely. It is very
	// unlikely anyone would be permanently injured by this amount of radioactive
	// material.
	//
	// Category 5: Cannot cause permanent injury. This category applies to x-ray
	// fluorescence devices and electron capture devices.
	MatCat param.Opt[int64] `json:"matCat,omitzero"`
	// The specific Material Class for the material associated with this detect. The
	// material class is generally associated with chemical and biological detections.
	MatClass param.Opt[string] `json:"matClass,omitzero"`
	// The material common name associated with this detection.
	MatName param.Opt[string] `json:"matName,omitzero"`
	// The specific material type (MT) or MT Code involved in this detection, when
	// applicable. The material type is generally associated with radiological and/or
	// nuclear detections. For further information, reference Nuclear Materials
	// Management and Safeguards System (NMMSS) Users Guide Rev. 2.1.
	MatType param.Opt[string] `json:"matType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Measure of the concentration of the material associated with this detection, in
	// parts per million (units of contaminant mass per million parts of total mass).
	Ppm param.Opt[int64] `json:"ppm,omitzero"`
	// Measure of radioactive concentration of the material associated with this
	// detection, in becquerels/m^3. One becquerel (Bq) is equal to one nuclear decay
	// per second.
	RadCtrn param.Opt[float64] `json:"radCtrn,omitzero"`
	// The Atomic Number of the material associated with this detection.
	Z param.Opt[int64] `json:"z,omitzero"`
	// Array of the specific readings associated with this detection. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	Readings []string `json:"readings,omitzero"`
	// Array of the units that correspond to each of the readingValues. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	ReadingUnits []string `json:"readingUnits,omitzero"`
	// Array of the values that correspond to each of the readings contained in
	// readings. The readings, readingUnits, and readingValues arrays must contain the
	// same number of elements.
	ReadingValues []float64 `json:"readingValues,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r HazardNewParams) MarshalJSON() (data []byte, err error) {
	type shadow HazardNewParams
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
type HazardNewParamsDataMode string

const (
	HazardNewParamsDataModeReal      HazardNewParamsDataMode = "REAL"
	HazardNewParamsDataModeTest      HazardNewParamsDataMode = "TEST"
	HazardNewParamsDataModeSimulated HazardNewParamsDataMode = "SIMULATED"
	HazardNewParamsDataModeExercise  HazardNewParamsDataMode = "EXERCISE"
)

type HazardListParams struct {
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	DetectTime  time.Time        `query:"detectTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [HazardListParams]'s query parameters as `url.Values`.
func (r HazardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HazardCountParams struct {
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	DetectTime  time.Time        `query:"detectTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardCountParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [HazardCountParams]'s query parameters as `url.Values`.
func (r HazardCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HazardNewBulkParams struct {
	Body []HazardNewBulkParamsBody
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardNewBulkParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r HazardNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// Hazard contains information concerning the exposure of a geographic area to a
// Chemical, Biological, Radiological, or Nuclear (CBRN) contaminant. The Hazard
// schema includes the detection time and type of contamination as well as optional
// information regarding specific material properties, the extent of contamination,
// and identification of affected regions.
//
// The properties Alarms, AlarmValues, ClassificationMarking, DataMode, DetectTime,
// DetectType, Source are required.
type HazardNewBulkParamsBody struct {
	// Array of the specific alarms associated with this detection. The alarms and
	// alarmValues arrays must contain the same number of elements.
	Alarms []string `json:"alarms,omitzero,required"`
	// Array of the values that correspond to each of the alarms contained in alarms.
	// The alarms and alarmValues arrays must contain the same number of elements.
	AlarmValues []float64 `json:"alarmValues,omitzero,required"`
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
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	DetectTime time.Time `json:"detectTime,required" format:"date-time"`
	// The type of hazard (Chemical, Biological, Radiological, Nuclear) detect
	// associated with this record.
	DetectType string `json:"detectType,required"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// The (rounded) Mass Number of the material associated with this detection. The
	// rounded value is the mass number of the most abundant isotope of the element.
	A param.Opt[int64] `json:"a,omitzero"`
	// The radioactivity measurement of the material associated with this detection, in
	// becquerels (Bq). One becquerel (Bq) is equal to one nuclear decay per second.
	Activity param.Opt[float64] `json:"activity,omitzero"`
	// The specific bottle ID associated with this detection.
	BottleID param.Opt[string] `json:"bottleId,omitzero"`
	// The CAS Registry Number, also referred to as CAS Number or CAS RN, associated
	// with the this detection. The CAS Number is a unique numerical identifier
	// assigned by the Chemical Abstracts Service (CAS), to every chemical substance
	// described in the open scientific literature. It includes organic and inorganic
	// compounds, minerals, isotopes, alloys, mixtures, and nonstructurable materials
	// (UVCBs, substances of unknown or variable composition, complex reaction
	// products, or biological origin). For further information, reference
	// https://www.cas.org/cas-data/cas-registry.
	CasRn param.Opt[string] `json:"casRN,omitzero"`
	// The applicable channel involved in this biological material detection (e.g.
	// Digestive, Eyes, Respiratory, Skin, etc.) .
	Channel param.Opt[string] `json:"channel,omitzero"`
	// Time the row was created in the database, auto-populated by the system.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The concentration time, in (kg/sec)/m^3, associated with this material
	// detection.
	CtrnTime param.Opt[float64] `json:"ctrnTime,omitzero"`
	// Measure of density of the material associated with this detection, in kg/m^3.
	Density param.Opt[float64] `json:"density,omitzero"`
	// The deposition measurement of the contaminant to surface area, in kg/m^2.
	Dep param.Opt[float64] `json:"dep,omitzero"`
	// The deposition concentration of the contaminant to surface area, in
	// becquerels/m^2.
	DepCtrn param.Opt[float64] `json:"depCtrn,omitzero"`
	// The dose associated with this detection, in gray. Dose is the measure of the
	// energy deposited in matter by ionizing radiation per unit mass. One gray is
	// defined as one Joule of energy absorbed per kilogram of matter.
	Dose param.Opt[float64] `json:"dose,omitzero"`
	// The dose rate associated with this detection, in gray/sec. One gray is defined
	// as one Joule of energy absorbed per kilogram of matter.
	DoseRate param.Opt[float64] `json:"doseRate,omitzero"`
	// The known or projected hazard duration, in seconds, associated with this
	// material detection.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Chemical Agent Monitor (CAM) G-type agent measurement, in number of display
	// bars. In G-mode, CAMs monitor for G-series nerve agents.
	GBar param.Opt[float64] `json:"gBar,omitzero"`
	// Flag indicating whether this detection is harmful to humans.
	Harmful param.Opt[bool] `json:"harmful,omitzero"`
	// Chemical Agent Monitor (CAM) H-type agent measurement, in number of display
	// bars. In H-mode, CAMs monitor for blister agents.
	HBar param.Opt[float64] `json:"hBar,omitzero"`
	// ID of the Point of Interest (POI) record related to this hazard record.
	IDPoi param.Opt[string] `json:"idPOI,omitzero"`
	// ID of the Track record related to this hazard record.
	IDTrack param.Opt[string] `json:"idTrack,omitzero"`
	// Ratio of the chemical substance mass to the total mass of the mixture.
	MassFrac param.Opt[float64] `json:"massFrac,omitzero"`
	// The Radiological Category (1 - 5) which applies to the material associated with
	// this detection, according to the following definitions:
	//
	// Category 1: If not safely or securely managed, would be likely to cause
	// permanent injury to a person who handled them or was otherwise in contact with
	// them for more than a few minutes. It would probably be fatal to be close to this
	// amount of unshielded material for a period of a few minutes to an hour.
	//
	// Category 2: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for a short time
	// (minutes to hours). It could possibly be fatal to be close to this amount of
	// unshielded radioactive material for a period of hours to days.
	//
	// Category 3: If not safely or securely managed, could cause permanent injury to a
	// person who handled them or was otherwise in contact with them for hours. It
	// could possibly - although it is unlikely to be - fatal to be close to this
	// amount of unshielded radioactive material for a period of days to weeks.
	//
	// Category 4: If not safely managed or securely protected, could possibly cause
	// temporary injury to someone who handled them or was otherwise in contact with or
	// close to them for a period of many weeks, though this is unlikely. It is very
	// unlikely anyone would be permanently injured by this amount of radioactive
	// material.
	//
	// Category 5: Cannot cause permanent injury. This category applies to x-ray
	// fluorescence devices and electron capture devices.
	MatCat param.Opt[int64] `json:"matCat,omitzero"`
	// The specific Material Class for the material associated with this detect. The
	// material class is generally associated with chemical and biological detections.
	MatClass param.Opt[string] `json:"matClass,omitzero"`
	// The material common name associated with this detection.
	MatName param.Opt[string] `json:"matName,omitzero"`
	// The specific material type (MT) or MT Code involved in this detection, when
	// applicable. The material type is generally associated with radiological and/or
	// nuclear detections. For further information, reference Nuclear Materials
	// Management and Safeguards System (NMMSS) Users Guide Rev. 2.1.
	MatType param.Opt[string] `json:"matType,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// Measure of the concentration of the material associated with this detection, in
	// parts per million (units of contaminant mass per million parts of total mass).
	Ppm param.Opt[int64] `json:"ppm,omitzero"`
	// Measure of radioactive concentration of the material associated with this
	// detection, in becquerels/m^3. One becquerel (Bq) is equal to one nuclear decay
	// per second.
	RadCtrn param.Opt[float64] `json:"radCtrn,omitzero"`
	// The Atomic Number of the material associated with this detection.
	Z param.Opt[int64] `json:"z,omitzero"`
	// Array of the specific readings associated with this detection. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	Readings []string `json:"readings,omitzero"`
	// Array of the units that correspond to each of the readingValues. The readings,
	// readingUnits, and readingValues arrays must contain the same number of elements.
	ReadingUnits []string `json:"readingUnits,omitzero"`
	// Array of the values that correspond to each of the readings contained in
	// readings. The readings, readingUnits, and readingValues arrays must contain the
	// same number of elements.
	ReadingValues []float64 `json:"readingValues,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardNewBulkParamsBody) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r HazardNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow HazardNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[HazardNewBulkParamsBody](
		"DataMode", false, "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type HazardGetParams struct {
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [HazardGetParams]'s query parameters as `url.Values`.
func (r HazardGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HazardTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// The detect time, in ISO 8601 UTC format, with millisecond precision.
	// (YYYY-MM-DDTHH:MM:SS.sssZ)
	DetectTime  time.Time        `query:"detectTime,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f HazardTupleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [HazardTupleParams]'s query parameters as `url.Values`.
func (r HazardTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
