// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/internal/apiform"
	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/apiquery"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/pagination"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// SigactService contains methods and other services that help with interacting
// with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSigactService] method instead.
type SigactService struct {
	Options []option.RequestOption
	History SigactHistoryService
}

// NewSigactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSigactService(opts ...option.RequestOption) (r SigactService) {
	r = SigactService{}
	r.Options = opts
	r.History = NewSigactHistoryService(opts...)
	return
}

// Service operation to dynamically query data by a variety of query parameters not
// specified in this API documentation. See the queryhelp operation
// (/udl/&lt;datatype&gt;/queryhelp) for more details on valid/required query
// parameter information.
func (r *SigactService) List(ctx context.Context, query SigactListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SigactListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "udl/sigact"
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
func (r *SigactService) ListAutoPaging(ctx context.Context, query SigactListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SigactListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Service operation to return the count of records satisfying the specified query
// parameters. This operation is useful to determine how many records pass a
// particular query criteria without retrieving large amounts of data. See the
// queryhelp operation (/udl/&lt;datatype&gt;/queryhelp) for more details on
// valid/required query parameter information.
func (r *SigactService) Count(ctx context.Context, query SigactCountParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "udl/sigact/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Service operation intended for initial integration only, to take a list of
// SigAct records as a POST body and ingest into the database. Requires specific
// roles, please contact the UDL team to gain access. This operation is not
// intended to be used for automated feeds into UDL...data providers should contact
// the UDL team for instructions on setting up a permanent feed through an
// alternate mechanism.
func (r *SigactService) NewBulk(ctx context.Context, body SigactNewBulkParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/sigact/createBulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SigactService) Queryhelp(ctx context.Context, opts ...option.RequestOption) (res *SigactQueryhelpResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sigact/queryhelp"
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
func (r *SigactService) Tuple(ctx context.Context, query SigactTupleParams, opts ...option.RequestOption) (res *[]SigactTupleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/sigact/tuple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a text file with its metadata. This operation bypasses the length
// constraints of the `eventDescription` field.
//
// The request body requires a zip file containing exactly two files:\
// 1\) A file with the `.json` file extension whose content conforms to the `SigAct_Ingest`
// schema.\
// 2\) A UTF-8 encoded file with the `.txt` file extension.
//
// The JSON and text files will be associated with each other via the `id` field.
// Query the metadata via `GET /udl/sigact` and use `GET /udl/sigact/getFile/{id}`
// to retrieve the text file.
//
// This operation only accepts application/zip media. The application/json request
// body is documented to provide a convenient reference to the ingest schema.
//
// This operation is intended to be used for automated feeds into UDL. A specific
// role is required to perform this service operation. Please contact the UDL team
// for assistance.
func (r *SigactService) UploadZip(ctx context.Context, body SigactUploadZipParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "filedrop/udl-sigact-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Provides information on the dates, actors, locations, fatalities, and types of
// all reported political violence and protest events across the world.
type SigactListResponse struct {
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
	DataMode SigactListResponseDataMode `json:"dataMode,required"`
	// Date of the report or filing.
	ReportDate time.Time `json:"reportDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Estimate of the accuracy that this event occurred as described/reported.
	Accuracy int64 `json:"accuracy"`
	// A list of one or more actors involved in the event.
	Actors []string `json:"actors"`
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
	// Type of region as projected on the ground.
	Atype string `json:"atype"`
	// This is the average tone of all documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. The score
	// ranges from -100 (extremely negative) to +100 (extremely positive). Common
	// values range between -10 and +10, with 0 indicating neutral.
	AvgTone float64 `json:"avgTone"`
	// CAMEO event codes are defined in a three-level taxonomy. For events at level
	// three in the taxonomy, this yields its level two leaf root node. For example,
	// code 0251 (Appeal for easing of administrative sanctions) would yield an
	// EventBaseCode of 025 (Appeal to yield). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoBaseCode string `json:"cameoBaseCode"`
	// This is the raw CAMEO action code describing the action that Actor1 performed
	// upon Actor2. Additional information about Cameo Codes can be obtained from the
	// GDELT project documentation here:
	// https://www.gdeltproject.org/data.html#documentation.
	CameoCode string `json:"cameoCode"`
	// Similar to EventBaseCode, this defines the root-level category the event code
	// falls under. For example, code 0251 (Appeal for easing of administrative
	// sanctions) has a root code of 02 (Appeal). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoRootCode string `json:"cameoRootCode"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// The city in or near which this event occurred.
	City string `json:"city"`
	// Number of civilians abducted in the activity.
	CivAbd int64 `json:"civAbd"`
	// Number of civilians detained in the activity.
	CivDet int64 `json:"civDet"`
	// Number of civilians killed in the activity.
	CivKia int64 `json:"civKIA"`
	// Number of civilians wounded in the activity.
	CivWound int64 `json:"civWound"`
	// 1 (high) for events where the reporting allows the coder to identify the event
	// in full. That is, events where the individual happening is described by the
	// original source in a sufficiently detailed way as to identify individual
	// incidents, i.e. separate activities of fighting in a single location:
	//
	// 2 (lower) for events where an aggregation of information was already made by the
	// source material that is impossible to undo in the coding process. Such events
	// are described by the original source only as aggregates (totals) of multiple
	// separate activities of fighting spanning over a longer period than a single,
	// clearly defined day.
	Clarity int64 `json:"clarity"`
	// Number of coalition members abducted in the activity.
	CoalAbd int64 `json:"coalAbd"`
	// Number of coalition members detained in the activity.
	CoalDet int64 `json:"coalDet"`
	// Number of coalition members killed in the activity.
	CoalKia int64 `json:"coalKIA"`
	// Number of coalition members wounded in the activity.
	CoalWound int64 `json:"coalWound"`
	// Flag indicating that this attack was of a complex or coordinated nature.
	ComplexAttack bool `json:"complexAttack"`
	// Estimate of the confidence that this event occurred.
	Confidence int64 `json:"confidence"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The district in which this event occurred.
	District string `json:"district"`
	// The filename of the document or report.
	DocumentFilename string `json:"documentFilename"`
	// The source of the document or report.
	DocumentSource string `json:"documentSource"`
	// Number of enemy combatants abducted in the activity.
	EnemyAbd int64 `json:"enemyAbd"`
	// Number of enemy combatants detained in the activity.
	EnemyDet int64 `json:"enemyDet"`
	// Number of enemy combatants killed in the activity.
	EnemyKia int64 `json:"enemyKIA"`
	// A description of the event.
	EventDescription string `json:"eventDescription"`
	// The approximate end time of the event, in ISO 8601 UTC format.
	EventEnd time.Time `json:"eventEnd" format:"date-time"`
	// The approximate start time of the event, in ISO 8601 UTC format.
	EventStart time.Time `json:"eventStart" format:"date-time"`
	// The type of event (e.g. Military, Natural, Political, Social, etc.).
	EventType string `json:"eventType"`
	// Size of the associated text file. Units in bytes. If filesize is provided
	// without an associated file, it defaults to 0.
	Filesize int64 `json:"filesize"`
	// Number of friendlies abducted in the activity.
	FriendlyAbd int64 `json:"friendlyAbd"`
	// Number of friendlies in the activity.
	FriendlyDet int64 `json:"friendlyDet"`
	// Number of friendlies killed in the activity.
	FriendlyKia int64 `json:"friendlyKIA"`
	// Number of friendlies wounded in the activity.
	FriendlyWound int64 `json:"friendlyWound"`
	// Each CAMEO event code is assigned a numeric score from -10 to +10, capturing the
	// theoretical potential impact that type of event will have on the stability of a
	// country. This is known as the Goldstein Scale. NOTE: this score is based on the
	// type of event, not the specifics of the actual event record being recorded thus
	// two riots, one with 10 people and one with 10,000, will both receive the same
	// Goldstein score. This can be aggregated to various levels of time resolution to
	// yield an approximation of the stability of a location over time.
	Goldstein float64 `json:"goldstein"`
	// Flag indicating this SigAct record has an associated txt file stored in the UDL.
	// Retrieve the txt file by using the GET/udl/sigact/getFile/{id} where id is the
	// same as the SigAct record id. The maximum file size for this service is
	// 10,000,000 bytes (10MB). Files exceeding the maximum size will be rejected.
	HasAttachment bool `json:"hasAttachment"`
	// Number of Host Nation members abducted in the activity.
	HostNatAbd int64 `json:"hostNatAbd"`
	// Number of Host Nation members detained in the activity.
	HostNatDet int64 `json:"hostNatDet"`
	// Number of Host Nation members killed in the activity.
	HostNatKia int64 `json:"hostNatKIA"`
	// Number of Host Nation members wounded in the activity.
	HostNatWound int64 `json:"hostNatWound"`
	// Unique identifier assigned to each event record that uniquely identifies it in
	// the master dataset. This ID is provided for convenience of mapping to external
	// systems.
	IDNumber string `json:"idNumber"`
	// WGS-84 centroid latitude of the event location, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat"`
	// WGS-84 centroid longitude of the event location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// &nbsp;&nbsp;4Q (grid zone designator, GZD)
	//
	// &nbsp;&nbsp;FJ (the 100,000-meter square identifier)
	//
	// &nbsp;&nbsp;12345678 (numerical location; easting is 1234 and northing is 5678,
	// in this case specifying a location with 10 m resolution).
	Milgrid string `json:"milgrid"`
	// Notes related to the documents or event.
	Notes string `json:"notes"`
	// This is the total number of source documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumArticles int64 `json:"numArticles"`
	// This is the total number of mentions of this event across all source documents
	// during the 15 minute update in which it was first seen. Multiple references to
	// an event within a single document also contribute to this count. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumMentions int64 `json:"numMentions"`
	// This is the total number of information sources containing one or more mentions
	// of this event during the 15 minute update in which it was first seen. This can
	// be used as a method of assessing the importance of an event: the more discussion
	// of that event, the more likely it is to be significant.
	NumSources int64 `json:"numSources"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The province in which this event occurred.
	Province string `json:"province"`
	// Related document ids.
	RelatedDocs []SigactListResponseRelatedDoc `json:"relatedDocs"`
	// The reporting unit.
	RepUnit string `json:"repUnit"`
	// The activity the unit was engaged in.
	RepUnitActivity string `json:"repUnitActivity"`
	// The reporting unit type.
	RepUnitType string `json:"repUnitType"`
	// Number of side A members abducted in the activity.
	SideAAbd int64 `json:"sideAAbd"`
	// Number of side A members detained in the activity.
	SideADet int64 `json:"sideADet"`
	// Number of side A members killed in the activity.
	SideAkia int64 `json:"sideAKIA"`
	// Number of side A members wounded in the activity.
	SideAWound int64 `json:"sideAWound"`
	// Number of side B members abducted in the activity.
	SideBAbd int64 `json:"sideBAbd"`
	// Number of side B members detained in the activity.
	SideBDet int64 `json:"sideBDet"`
	// Number of side B members killed in the activity.
	SideBkia int64 `json:"sideBKIA"`
	// Number of side B members wounded in the activity.
	SideBWound int64 `json:"sideBWound"`
	// The source language of the significant event using the ISO 639-3, 3 character
	// code definition.
	SourceLanguage string `json:"sourceLanguage"`
	// This field records the URL or citation of the first news report it found this
	// event in. In most cases this is the first report it saw the article in, but due
	// to the timing and flow of news reports through the processing pipeline, this may
	// not always be the very first report, but is at least in the first few reports.
	SourceURL string `json:"sourceUrl"`
	// A summary of the event.
	Summary string `json:"summary"`
	// The name of the target. The target may be an individual, an entity, or a
	// country/region.
	Target string `json:"target"`
	// Area in which important military events occur or are progressing. A theater can
	// include the entirety of the airspace, land and sea area that is or that may
	// potentially become involved in war operations.
	Theater string `json:"theater"`
	// The mode of this attack or event (e.g. Direct Fire, IED Explosion, etc.).
	TypeOfAttack string `json:"typeOfAttack"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ReportDate            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Accuracy              respjson.Field
		Actors                respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		AvgTone               respjson.Field
		CameoBaseCode         respjson.Field
		CameoCode             respjson.Field
		CameoRootCode         respjson.Field
		ChecksumValue         respjson.Field
		City                  respjson.Field
		CivAbd                respjson.Field
		CivDet                respjson.Field
		CivKia                respjson.Field
		CivWound              respjson.Field
		Clarity               respjson.Field
		CoalAbd               respjson.Field
		CoalDet               respjson.Field
		CoalKia               respjson.Field
		CoalWound             respjson.Field
		ComplexAttack         respjson.Field
		Confidence            respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		District              respjson.Field
		DocumentFilename      respjson.Field
		DocumentSource        respjson.Field
		EnemyAbd              respjson.Field
		EnemyDet              respjson.Field
		EnemyKia              respjson.Field
		EventDescription      respjson.Field
		EventEnd              respjson.Field
		EventStart            respjson.Field
		EventType             respjson.Field
		Filesize              respjson.Field
		FriendlyAbd           respjson.Field
		FriendlyDet           respjson.Field
		FriendlyKia           respjson.Field
		FriendlyWound         respjson.Field
		Goldstein             respjson.Field
		HasAttachment         respjson.Field
		HostNatAbd            respjson.Field
		HostNatDet            respjson.Field
		HostNatKia            respjson.Field
		HostNatWound          respjson.Field
		IDNumber              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Milgrid               respjson.Field
		Notes                 respjson.Field
		NumArticles           respjson.Field
		NumMentions           respjson.Field
		NumSources            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Province              respjson.Field
		RelatedDocs           respjson.Field
		RepUnit               respjson.Field
		RepUnitActivity       respjson.Field
		RepUnitType           respjson.Field
		SideAAbd              respjson.Field
		SideADet              respjson.Field
		SideAkia              respjson.Field
		SideAWound            respjson.Field
		SideBAbd              respjson.Field
		SideBDet              respjson.Field
		SideBkia              respjson.Field
		SideBWound            respjson.Field
		SourceLanguage        respjson.Field
		SourceURL             respjson.Field
		Summary               respjson.Field
		Target                respjson.Field
		Theater               respjson.Field
		TypeOfAttack          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactListResponse) RawJSON() string { return r.JSON.raw }
func (r *SigactListResponse) UnmarshalJSON(data []byte) error {
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
type SigactListResponseDataMode string

const (
	SigactListResponseDataModeReal      SigactListResponseDataMode = "REAL"
	SigactListResponseDataModeTest      SigactListResponseDataMode = "TEST"
	SigactListResponseDataModeSimulated SigactListResponseDataMode = "SIMULATED"
	SigactListResponseDataModeExercise  SigactListResponseDataMode = "EXERCISE"
)

type SigactListResponseRelatedDoc struct {
	// List of data sources related to this document.
	DataSourceRefs []SigactListResponseRelatedDocDataSourceRef `json:"dataSourceRefs"`
	// The document id of the related document.
	DocumentID string `json:"documentId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSourceRefs respjson.Field
		DocumentID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactListResponseRelatedDoc) RawJSON() string { return r.JSON.raw }
func (r *SigactListResponseRelatedDoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactListResponseRelatedDocDataSourceRef struct {
	// Data source id.
	DataSourceID string `json:"dataSourceId"`
	// end position.
	EndPosition string `json:"endPosition"`
	// paragraph number.
	ParagraphNumber string `json:"paragraphNumber"`
	// sentence number.
	SentenceNumber string `json:"sentenceNumber"`
	// start position.
	StartPosition string `json:"startPosition"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSourceID    respjson.Field
		EndPosition     respjson.Field
		ParagraphNumber respjson.Field
		SentenceNumber  respjson.Field
		StartPosition   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactListResponseRelatedDocDataSourceRef) RawJSON() string { return r.JSON.raw }
func (r *SigactListResponseRelatedDocDataSourceRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactQueryhelpResponse struct {
	AodrSupported         bool                               `json:"aodrSupported"`
	ClassificationMarking string                             `json:"classificationMarking"`
	Description           string                             `json:"description"`
	HistorySupported      bool                               `json:"historySupported"`
	Name                  string                             `json:"name"`
	Parameters            []SigactQueryhelpResponseParameter `json:"parameters"`
	RequiredRoles         []string                           `json:"requiredRoles"`
	RestSupported         bool                               `json:"restSupported"`
	SortSupported         bool                               `json:"sortSupported"`
	TypeName              string                             `json:"typeName"`
	Uri                   string                             `json:"uri"`
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
func (r SigactQueryhelpResponse) RawJSON() string { return r.JSON.raw }
func (r *SigactQueryhelpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactQueryhelpResponseParameter struct {
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
func (r SigactQueryhelpResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *SigactQueryhelpResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides information on the dates, actors, locations, fatalities, and types of
// all reported political violence and protest events across the world.
type SigactTupleResponse struct {
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
	DataMode SigactTupleResponseDataMode `json:"dataMode,required"`
	// Date of the report or filing.
	ReportDate time.Time `json:"reportDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID string `json:"id"`
	// Estimate of the accuracy that this event occurred as described/reported.
	Accuracy int64 `json:"accuracy"`
	// A list of one or more actors involved in the event.
	Actors []string `json:"actors"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson string `json:"agjson"`
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
	// This is the average tone of all documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. The score
	// ranges from -100 (extremely negative) to +100 (extremely positive). Common
	// values range between -10 and +10, with 0 indicating neutral.
	AvgTone float64 `json:"avgTone"`
	// CAMEO event codes are defined in a three-level taxonomy. For events at level
	// three in the taxonomy, this yields its level two leaf root node. For example,
	// code 0251 (Appeal for easing of administrative sanctions) would yield an
	// EventBaseCode of 025 (Appeal to yield). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoBaseCode string `json:"cameoBaseCode"`
	// This is the raw CAMEO action code describing the action that Actor1 performed
	// upon Actor2. Additional information about Cameo Codes can be obtained from the
	// GDELT project documentation here:
	// https://www.gdeltproject.org/data.html#documentation.
	CameoCode string `json:"cameoCode"`
	// Similar to EventBaseCode, this defines the root-level category the event code
	// falls under. For example, code 0251 (Appeal for easing of administrative
	// sanctions) has a root code of 02 (Appeal). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoRootCode string `json:"cameoRootCode"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue string `json:"checksumValue"`
	// The city in or near which this event occurred.
	City string `json:"city"`
	// Number of civilians abducted in the activity.
	CivAbd int64 `json:"civAbd"`
	// Number of civilians detained in the activity.
	CivDet int64 `json:"civDet"`
	// Number of civilians killed in the activity.
	CivKia int64 `json:"civKIA"`
	// Number of civilians wounded in the activity.
	CivWound int64 `json:"civWound"`
	// 1 (high) for events where the reporting allows the coder to identify the event
	// in full. That is, events where the individual happening is described by the
	// original source in a sufficiently detailed way as to identify individual
	// incidents, i.e. separate activities of fighting in a single location:
	//
	// 2 (lower) for events where an aggregation of information was already made by the
	// source material that is impossible to undo in the coding process. Such events
	// are described by the original source only as aggregates (totals) of multiple
	// separate activities of fighting spanning over a longer period than a single,
	// clearly defined day.
	Clarity int64 `json:"clarity"`
	// Number of coalition members abducted in the activity.
	CoalAbd int64 `json:"coalAbd"`
	// Number of coalition members detained in the activity.
	CoalDet int64 `json:"coalDet"`
	// Number of coalition members killed in the activity.
	CoalKia int64 `json:"coalKIA"`
	// Number of coalition members wounded in the activity.
	CoalWound int64 `json:"coalWound"`
	// Flag indicating that this attack was of a complex or coordinated nature.
	ComplexAttack bool `json:"complexAttack"`
	// Estimate of the confidence that this event occurred.
	Confidence int64 `json:"confidence"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode string `json:"countryCode"`
	// Time the row was created in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy string `json:"createdBy"`
	// The district in which this event occurred.
	District string `json:"district"`
	// The filename of the document or report.
	DocumentFilename string `json:"documentFilename"`
	// The source of the document or report.
	DocumentSource string `json:"documentSource"`
	// Number of enemy combatants abducted in the activity.
	EnemyAbd int64 `json:"enemyAbd"`
	// Number of enemy combatants detained in the activity.
	EnemyDet int64 `json:"enemyDet"`
	// Number of enemy combatants killed in the activity.
	EnemyKia int64 `json:"enemyKIA"`
	// A description of the event.
	EventDescription string `json:"eventDescription"`
	// The approximate end time of the event, in ISO 8601 UTC format.
	EventEnd time.Time `json:"eventEnd" format:"date-time"`
	// The approximate start time of the event, in ISO 8601 UTC format.
	EventStart time.Time `json:"eventStart" format:"date-time"`
	// The type of event (e.g. Military, Natural, Political, Social, etc.).
	EventType string `json:"eventType"`
	// Size of the associated text file. Units in bytes. If filesize is provided
	// without an associated file, it defaults to 0.
	Filesize int64 `json:"filesize"`
	// Number of friendlies abducted in the activity.
	FriendlyAbd int64 `json:"friendlyAbd"`
	// Number of friendlies in the activity.
	FriendlyDet int64 `json:"friendlyDet"`
	// Number of friendlies killed in the activity.
	FriendlyKia int64 `json:"friendlyKIA"`
	// Number of friendlies wounded in the activity.
	FriendlyWound int64 `json:"friendlyWound"`
	// Each CAMEO event code is assigned a numeric score from -10 to +10, capturing the
	// theoretical potential impact that type of event will have on the stability of a
	// country. This is known as the Goldstein Scale. NOTE: this score is based on the
	// type of event, not the specifics of the actual event record being recorded thus
	// two riots, one with 10 people and one with 10,000, will both receive the same
	// Goldstein score. This can be aggregated to various levels of time resolution to
	// yield an approximation of the stability of a location over time.
	Goldstein float64 `json:"goldstein"`
	// Flag indicating this SigAct record has an associated txt file stored in the UDL.
	// Retrieve the txt file by using the GET/udl/sigact/getFile/{id} where id is the
	// same as the SigAct record id. The maximum file size for this service is
	// 10,000,000 bytes (10MB). Files exceeding the maximum size will be rejected.
	HasAttachment bool `json:"hasAttachment"`
	// Number of Host Nation members abducted in the activity.
	HostNatAbd int64 `json:"hostNatAbd"`
	// Number of Host Nation members detained in the activity.
	HostNatDet int64 `json:"hostNatDet"`
	// Number of Host Nation members killed in the activity.
	HostNatKia int64 `json:"hostNatKIA"`
	// Number of Host Nation members wounded in the activity.
	HostNatWound int64 `json:"hostNatWound"`
	// Unique identifier assigned to each event record that uniquely identifies it in
	// the master dataset. This ID is provided for convenience of mapping to external
	// systems.
	IDNumber string `json:"idNumber"`
	// WGS-84 centroid latitude of the event location, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat float64 `json:"lat"`
	// WGS-84 centroid longitude of the event location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon float64 `json:"lon"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// &nbsp;&nbsp;4Q (grid zone designator, GZD)
	//
	// &nbsp;&nbsp;FJ (the 100,000-meter square identifier)
	//
	// &nbsp;&nbsp;12345678 (numerical location; easting is 1234 and northing is 5678,
	// in this case specifying a location with 10 m resolution).
	Milgrid string `json:"milgrid"`
	// Notes related to the documents or event.
	Notes string `json:"notes"`
	// This is the total number of source documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumArticles int64 `json:"numArticles"`
	// This is the total number of mentions of this event across all source documents
	// during the 15 minute update in which it was first seen. Multiple references to
	// an event within a single document also contribute to this count. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumMentions int64 `json:"numMentions"`
	// This is the total number of information sources containing one or more mentions
	// of this event during the 15 minute update in which it was first seen. This can
	// be used as a method of assessing the importance of an event: the more discussion
	// of that event, the more likely it is to be significant.
	NumSources int64 `json:"numSources"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin string `json:"origin"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork string `json:"origNetwork"`
	// The province in which this event occurred.
	Province string `json:"province"`
	// Related document ids.
	RelatedDocs []SigactTupleResponseRelatedDoc `json:"relatedDocs"`
	// The reporting unit.
	RepUnit string `json:"repUnit"`
	// The activity the unit was engaged in.
	RepUnitActivity string `json:"repUnitActivity"`
	// The reporting unit type.
	RepUnitType string `json:"repUnitType"`
	// Number of side A members abducted in the activity.
	SideAAbd int64 `json:"sideAAbd"`
	// Number of side A members detained in the activity.
	SideADet int64 `json:"sideADet"`
	// Number of side A members killed in the activity.
	SideAkia int64 `json:"sideAKIA"`
	// Number of side A members wounded in the activity.
	SideAWound int64 `json:"sideAWound"`
	// Number of side B members abducted in the activity.
	SideBAbd int64 `json:"sideBAbd"`
	// Number of side B members detained in the activity.
	SideBDet int64 `json:"sideBDet"`
	// Number of side B members killed in the activity.
	SideBkia int64 `json:"sideBKIA"`
	// Number of side B members wounded in the activity.
	SideBWound int64 `json:"sideBWound"`
	// The source language of the significant event using the ISO 639-3, 3 character
	// code definition.
	SourceLanguage string `json:"sourceLanguage"`
	// This field records the URL or citation of the first news report it found this
	// event in. In most cases this is the first report it saw the article in, but due
	// to the timing and flow of news reports through the processing pipeline, this may
	// not always be the very first report, but is at least in the first few reports.
	SourceURL string `json:"sourceUrl"`
	// A summary of the event.
	Summary string `json:"summary"`
	// The name of the target. The target may be an individual, an entity, or a
	// country/region.
	Target string `json:"target"`
	// Area in which important military events occur or are progressing. A theater can
	// include the entirety of the airspace, land and sea area that is or that may
	// potentially become involved in war operations.
	Theater string `json:"theater"`
	// The mode of this attack or event (e.g. Direct Fire, IED Explosion, etc.).
	TypeOfAttack string `json:"typeOfAttack"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		DataMode              respjson.Field
		ReportDate            respjson.Field
		Source                respjson.Field
		ID                    respjson.Field
		Accuracy              respjson.Field
		Actors                respjson.Field
		Agjson                respjson.Field
		Andims                respjson.Field
		Area                  respjson.Field
		Asrid                 respjson.Field
		Atext                 respjson.Field
		Atype                 respjson.Field
		AvgTone               respjson.Field
		CameoBaseCode         respjson.Field
		CameoCode             respjson.Field
		CameoRootCode         respjson.Field
		ChecksumValue         respjson.Field
		City                  respjson.Field
		CivAbd                respjson.Field
		CivDet                respjson.Field
		CivKia                respjson.Field
		CivWound              respjson.Field
		Clarity               respjson.Field
		CoalAbd               respjson.Field
		CoalDet               respjson.Field
		CoalKia               respjson.Field
		CoalWound             respjson.Field
		ComplexAttack         respjson.Field
		Confidence            respjson.Field
		CountryCode           respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		District              respjson.Field
		DocumentFilename      respjson.Field
		DocumentSource        respjson.Field
		EnemyAbd              respjson.Field
		EnemyDet              respjson.Field
		EnemyKia              respjson.Field
		EventDescription      respjson.Field
		EventEnd              respjson.Field
		EventStart            respjson.Field
		EventType             respjson.Field
		Filesize              respjson.Field
		FriendlyAbd           respjson.Field
		FriendlyDet           respjson.Field
		FriendlyKia           respjson.Field
		FriendlyWound         respjson.Field
		Goldstein             respjson.Field
		HasAttachment         respjson.Field
		HostNatAbd            respjson.Field
		HostNatDet            respjson.Field
		HostNatKia            respjson.Field
		HostNatWound          respjson.Field
		IDNumber              respjson.Field
		Lat                   respjson.Field
		Lon                   respjson.Field
		Milgrid               respjson.Field
		Notes                 respjson.Field
		NumArticles           respjson.Field
		NumMentions           respjson.Field
		NumSources            respjson.Field
		Origin                respjson.Field
		OrigNetwork           respjson.Field
		Province              respjson.Field
		RelatedDocs           respjson.Field
		RepUnit               respjson.Field
		RepUnitActivity       respjson.Field
		RepUnitType           respjson.Field
		SideAAbd              respjson.Field
		SideADet              respjson.Field
		SideAkia              respjson.Field
		SideAWound            respjson.Field
		SideBAbd              respjson.Field
		SideBDet              respjson.Field
		SideBkia              respjson.Field
		SideBWound            respjson.Field
		SourceLanguage        respjson.Field
		SourceURL             respjson.Field
		Summary               respjson.Field
		Target                respjson.Field
		Theater               respjson.Field
		TypeOfAttack          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactTupleResponse) RawJSON() string { return r.JSON.raw }
func (r *SigactTupleResponse) UnmarshalJSON(data []byte) error {
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
type SigactTupleResponseDataMode string

const (
	SigactTupleResponseDataModeReal      SigactTupleResponseDataMode = "REAL"
	SigactTupleResponseDataModeTest      SigactTupleResponseDataMode = "TEST"
	SigactTupleResponseDataModeSimulated SigactTupleResponseDataMode = "SIMULATED"
	SigactTupleResponseDataModeExercise  SigactTupleResponseDataMode = "EXERCISE"
)

type SigactTupleResponseRelatedDoc struct {
	// List of data sources related to this document.
	DataSourceRefs []SigactTupleResponseRelatedDocDataSourceRef `json:"dataSourceRefs"`
	// The document id of the related document.
	DocumentID string `json:"documentId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSourceRefs respjson.Field
		DocumentID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactTupleResponseRelatedDoc) RawJSON() string { return r.JSON.raw }
func (r *SigactTupleResponseRelatedDoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactTupleResponseRelatedDocDataSourceRef struct {
	// Data source id.
	DataSourceID string `json:"dataSourceId"`
	// end position.
	EndPosition string `json:"endPosition"`
	// paragraph number.
	ParagraphNumber string `json:"paragraphNumber"`
	// sentence number.
	SentenceNumber string `json:"sentenceNumber"`
	// start position.
	StartPosition string `json:"startPosition"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSourceID    respjson.Field
		EndPosition     respjson.Field
		ParagraphNumber respjson.Field
		SentenceNumber  respjson.Field
		StartPosition   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SigactTupleResponseRelatedDocDataSourceRef) RawJSON() string { return r.JSON.raw }
func (r *SigactTupleResponseRelatedDocDataSourceRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactListParams struct {
	// Date of the report or filing. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportDate  time.Time        `query:"reportDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SigactListParams]'s query parameters as `url.Values`.
func (r SigactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SigactCountParams struct {
	// Date of the report or filing. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportDate  time.Time        `query:"reportDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SigactCountParams]'s query parameters as `url.Values`.
func (r SigactCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SigactNewBulkParams struct {
	Body []SigactNewBulkParamsBody
	paramObj
}

func (r SigactNewBulkParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SigactNewBulkParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Provides information on the dates, actors, locations, fatalities, and types of
// all reported political violence and protest events across the world.
//
// The properties ClassificationMarking, DataMode, ReportDate, Source are required.
type SigactNewBulkParamsBody struct {
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
	// Date of the report or filing.
	ReportDate time.Time `json:"reportDate,required" format:"date-time"`
	// Source of the data.
	Source string `json:"source,required"`
	// Unique identifier of the record, auto-generated by the system.
	ID param.Opt[string] `json:"id,omitzero"`
	// Estimate of the accuracy that this event occurred as described/reported.
	Accuracy param.Opt[int64] `json:"accuracy,omitzero"`
	// Geographical region or polygon (lat/lon pairs), as depicted by the GeoJSON
	// representation of the geometry/geography, of the image as projected on the
	// ground. GeoJSON Reference: https://geojson.org/. Ignored if included with a POST
	// or PUT request that also specifies a valid 'area' or 'atext' field.
	Agjson param.Opt[string] `json:"agjson,omitzero"`
	// Number of dimensions of the geometry depicted by region.
	Andims param.Opt[int64] `json:"andims,omitzero"`
	// Optional geographical region or polygon (lat/lon pairs) of the area surrounding
	// the point of interest as projected on the ground.
	Area param.Opt[string] `json:"area,omitzero"`
	// Geographical spatial_ref_sys for region.
	Asrid param.Opt[int64] `json:"asrid,omitzero"`
	// Geographical region or polygon (lon/lat pairs), as depicted by the Well-Known
	// Text representation of the geometry/geography, of the image as projected on the
	// ground. WKT reference: https://www.opengeospatial.org/standards/wkt-crs. Ignored
	// if included with a POST or PUT request that also specifies a valid 'area' field.
	Atext param.Opt[string] `json:"atext,omitzero"`
	// Type of region as projected on the ground.
	Atype param.Opt[string] `json:"atype,omitzero"`
	// This is the average tone of all documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. The score
	// ranges from -100 (extremely negative) to +100 (extremely positive). Common
	// values range between -10 and +10, with 0 indicating neutral.
	AvgTone param.Opt[float64] `json:"avgTone,omitzero"`
	// CAMEO event codes are defined in a three-level taxonomy. For events at level
	// three in the taxonomy, this yields its level two leaf root node. For example,
	// code 0251 (Appeal for easing of administrative sanctions) would yield an
	// EventBaseCode of 025 (Appeal to yield). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoBaseCode param.Opt[string] `json:"cameoBaseCode,omitzero"`
	// This is the raw CAMEO action code describing the action that Actor1 performed
	// upon Actor2. Additional information about Cameo Codes can be obtained from the
	// GDELT project documentation here:
	// https://www.gdeltproject.org/data.html#documentation.
	CameoCode param.Opt[string] `json:"cameoCode,omitzero"`
	// Similar to EventBaseCode, this defines the root-level category the event code
	// falls under. For example, code 0251 (Appeal for easing of administrative
	// sanctions) has a root code of 02 (Appeal). This makes it possible to aggregate
	// events at various resolutions of specificity. For events at levels two or one,
	// this field will be set to EventCode.
	CameoRootCode param.Opt[string] `json:"cameoRootCode,omitzero"`
	// MD5 value of the file. The ingest/create operation will automatically generate
	// the value.
	ChecksumValue param.Opt[string] `json:"checksumValue,omitzero"`
	// The city in or near which this event occurred.
	City param.Opt[string] `json:"city,omitzero"`
	// Number of civilians abducted in the activity.
	CivAbd param.Opt[int64] `json:"civAbd,omitzero"`
	// Number of civilians detained in the activity.
	CivDet param.Opt[int64] `json:"civDet,omitzero"`
	// Number of civilians killed in the activity.
	CivKia param.Opt[int64] `json:"civKIA,omitzero"`
	// Number of civilians wounded in the activity.
	CivWound param.Opt[int64] `json:"civWound,omitzero"`
	// 1 (high) for events where the reporting allows the coder to identify the event
	// in full. That is, events where the individual happening is described by the
	// original source in a sufficiently detailed way as to identify individual
	// incidents, i.e. separate activities of fighting in a single location:
	//
	// 2 (lower) for events where an aggregation of information was already made by the
	// source material that is impossible to undo in the coding process. Such events
	// are described by the original source only as aggregates (totals) of multiple
	// separate activities of fighting spanning over a longer period than a single,
	// clearly defined day.
	Clarity param.Opt[int64] `json:"clarity,omitzero"`
	// Number of coalition members abducted in the activity.
	CoalAbd param.Opt[int64] `json:"coalAbd,omitzero"`
	// Number of coalition members detained in the activity.
	CoalDet param.Opt[int64] `json:"coalDet,omitzero"`
	// Number of coalition members killed in the activity.
	CoalKia param.Opt[int64] `json:"coalKIA,omitzero"`
	// Number of coalition members wounded in the activity.
	CoalWound param.Opt[int64] `json:"coalWound,omitzero"`
	// Flag indicating that this attack was of a complex or coordinated nature.
	ComplexAttack param.Opt[bool] `json:"complexAttack,omitzero"`
	// Estimate of the confidence that this event occurred.
	Confidence param.Opt[int64] `json:"confidence,omitzero"`
	// The country code. This value is typically the ISO 3166 Alpha-2 two-character
	// country code, however it can also represent various consortiums that do not
	// appear in the ISO document. The code must correspond to an existing country in
	// the UDL’s country API. Call udl/country/{code} to get any associated FIPS code,
	// ISO Alpha-3 code, or alternate code values that exist for the specified country
	// code.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Time the row was created in the database, auto-populated by the system, example
	// = 2018-01-01T16:00:00.123Z.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Application user who created the row in the database, auto-populated by the
	// system.
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// The district in which this event occurred.
	District param.Opt[string] `json:"district,omitzero"`
	// The filename of the document or report.
	DocumentFilename param.Opt[string] `json:"documentFilename,omitzero"`
	// The source of the document or report.
	DocumentSource param.Opt[string] `json:"documentSource,omitzero"`
	// Number of enemy combatants abducted in the activity.
	EnemyAbd param.Opt[int64] `json:"enemyAbd,omitzero"`
	// Number of enemy combatants detained in the activity.
	EnemyDet param.Opt[int64] `json:"enemyDet,omitzero"`
	// Number of enemy combatants killed in the activity.
	EnemyKia param.Opt[int64] `json:"enemyKIA,omitzero"`
	// A description of the event.
	EventDescription param.Opt[string] `json:"eventDescription,omitzero"`
	// The approximate end time of the event, in ISO 8601 UTC format.
	EventEnd param.Opt[time.Time] `json:"eventEnd,omitzero" format:"date-time"`
	// The approximate start time of the event, in ISO 8601 UTC format.
	EventStart param.Opt[time.Time] `json:"eventStart,omitzero" format:"date-time"`
	// The type of event (e.g. Military, Natural, Political, Social, etc.).
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// Size of the associated text file. Units in bytes. If filesize is provided
	// without an associated file, it defaults to 0.
	Filesize param.Opt[int64] `json:"filesize,omitzero"`
	// Number of friendlies abducted in the activity.
	FriendlyAbd param.Opt[int64] `json:"friendlyAbd,omitzero"`
	// Number of friendlies in the activity.
	FriendlyDet param.Opt[int64] `json:"friendlyDet,omitzero"`
	// Number of friendlies killed in the activity.
	FriendlyKia param.Opt[int64] `json:"friendlyKIA,omitzero"`
	// Number of friendlies wounded in the activity.
	FriendlyWound param.Opt[int64] `json:"friendlyWound,omitzero"`
	// Each CAMEO event code is assigned a numeric score from -10 to +10, capturing the
	// theoretical potential impact that type of event will have on the stability of a
	// country. This is known as the Goldstein Scale. NOTE: this score is based on the
	// type of event, not the specifics of the actual event record being recorded thus
	// two riots, one with 10 people and one with 10,000, will both receive the same
	// Goldstein score. This can be aggregated to various levels of time resolution to
	// yield an approximation of the stability of a location over time.
	Goldstein param.Opt[float64] `json:"goldstein,omitzero"`
	// Flag indicating this SigAct record has an associated txt file stored in the UDL.
	// Retrieve the txt file by using the GET/udl/sigact/getFile/{id} where id is the
	// same as the SigAct record id. The maximum file size for this service is
	// 10,000,000 bytes (10MB). Files exceeding the maximum size will be rejected.
	HasAttachment param.Opt[bool] `json:"hasAttachment,omitzero"`
	// Number of Host Nation members abducted in the activity.
	HostNatAbd param.Opt[int64] `json:"hostNatAbd,omitzero"`
	// Number of Host Nation members detained in the activity.
	HostNatDet param.Opt[int64] `json:"hostNatDet,omitzero"`
	// Number of Host Nation members killed in the activity.
	HostNatKia param.Opt[int64] `json:"hostNatKIA,omitzero"`
	// Number of Host Nation members wounded in the activity.
	HostNatWound param.Opt[int64] `json:"hostNatWound,omitzero"`
	// Unique identifier assigned to each event record that uniquely identifies it in
	// the master dataset. This ID is provided for convenience of mapping to external
	// systems.
	IDNumber param.Opt[string] `json:"idNumber,omitzero"`
	// WGS-84 centroid latitude of the event location, in degrees. -90 to 90 degrees
	// (negative values south of equator).
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// WGS-84 centroid longitude of the event location, in degrees. -180 to 180 degrees
	// (negative values west of Prime Meridian).
	Lon param.Opt[float64] `json:"lon,omitzero"`
	// The Military Grid Reference System is the geocoordinate standard used by NATO
	// militaries for locating points on Earth. The MGRS is derived from the Universal
	// Transverse Mercator (UTM) grid system and the Universal Polar Stereographic
	// (UPS) grid system, but uses a different labeling convention. The MGRS is used as
	// geocode for the entire Earth. Example of an milgrid coordinate, or grid
	// reference, would be 4QFJ12345678, which consists of three parts:
	//
	// &nbsp;&nbsp;4Q (grid zone designator, GZD)
	//
	// &nbsp;&nbsp;FJ (the 100,000-meter square identifier)
	//
	// &nbsp;&nbsp;12345678 (numerical location; easting is 1234 and northing is 5678,
	// in this case specifying a location with 10 m resolution).
	Milgrid param.Opt[string] `json:"milgrid,omitzero"`
	// Notes related to the documents or event.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// This is the total number of source documents containing one or more mentions of
	// this event during the 15 minute update in which it was first seen. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumArticles param.Opt[int64] `json:"numArticles,omitzero"`
	// This is the total number of mentions of this event across all source documents
	// during the 15 minute update in which it was first seen. Multiple references to
	// an event within a single document also contribute to this count. This can be
	// used as a method of assessing the importance of an event: the more discussion of
	// that event, the more likely it is to be significant.
	NumMentions param.Opt[int64] `json:"numMentions,omitzero"`
	// This is the total number of information sources containing one or more mentions
	// of this event during the 15 minute update in which it was first seen. This can
	// be used as a method of assessing the importance of an event: the more discussion
	// of that event, the more likely it is to be significant.
	NumSources param.Opt[int64] `json:"numSources,omitzero"`
	// Originating system or organization which produced the data, if different from
	// the source. The origin may be different than the source if the source was a
	// mediating system which forwarded the data on behalf of the origin system. If
	// null, the source may be assumed to be the origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// The originating source network on which this record was created, auto-populated
	// by the system.
	OrigNetwork param.Opt[string] `json:"origNetwork,omitzero"`
	// The province in which this event occurred.
	Province param.Opt[string] `json:"province,omitzero"`
	// The reporting unit.
	RepUnit param.Opt[string] `json:"repUnit,omitzero"`
	// The activity the unit was engaged in.
	RepUnitActivity param.Opt[string] `json:"repUnitActivity,omitzero"`
	// The reporting unit type.
	RepUnitType param.Opt[string] `json:"repUnitType,omitzero"`
	// Number of side A members abducted in the activity.
	SideAAbd param.Opt[int64] `json:"sideAAbd,omitzero"`
	// Number of side A members detained in the activity.
	SideADet param.Opt[int64] `json:"sideADet,omitzero"`
	// Number of side A members killed in the activity.
	SideAkia param.Opt[int64] `json:"sideAKIA,omitzero"`
	// Number of side A members wounded in the activity.
	SideAWound param.Opt[int64] `json:"sideAWound,omitzero"`
	// Number of side B members abducted in the activity.
	SideBAbd param.Opt[int64] `json:"sideBAbd,omitzero"`
	// Number of side B members detained in the activity.
	SideBDet param.Opt[int64] `json:"sideBDet,omitzero"`
	// Number of side B members killed in the activity.
	SideBkia param.Opt[int64] `json:"sideBKIA,omitzero"`
	// Number of side B members wounded in the activity.
	SideBWound param.Opt[int64] `json:"sideBWound,omitzero"`
	// The source language of the significant event using the ISO 639-3, 3 character
	// code definition.
	SourceLanguage param.Opt[string] `json:"sourceLanguage,omitzero"`
	// This field records the URL or citation of the first news report it found this
	// event in. In most cases this is the first report it saw the article in, but due
	// to the timing and flow of news reports through the processing pipeline, this may
	// not always be the very first report, but is at least in the first few reports.
	SourceURL param.Opt[string] `json:"sourceUrl,omitzero"`
	// A summary of the event.
	Summary param.Opt[string] `json:"summary,omitzero"`
	// The name of the target. The target may be an individual, an entity, or a
	// country/region.
	Target param.Opt[string] `json:"target,omitzero"`
	// Area in which important military events occur or are progressing. A theater can
	// include the entirety of the airspace, land and sea area that is or that may
	// potentially become involved in war operations.
	Theater param.Opt[string] `json:"theater,omitzero"`
	// The mode of this attack or event (e.g. Direct Fire, IED Explosion, etc.).
	TypeOfAttack param.Opt[string] `json:"typeOfAttack,omitzero"`
	// A list of one or more actors involved in the event.
	Actors []string `json:"actors,omitzero"`
	// Related document ids.
	RelatedDocs []SigactNewBulkParamsBodyRelatedDoc `json:"relatedDocs,omitzero"`
	paramObj
}

func (r SigactNewBulkParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SigactNewBulkParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SigactNewBulkParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SigactNewBulkParamsBody](
		"dataMode", "REAL", "TEST", "SIMULATED", "EXERCISE",
	)
}

type SigactNewBulkParamsBodyRelatedDoc struct {
	// The document id of the related document.
	DocumentID param.Opt[string] `json:"documentId,omitzero"`
	// List of data sources related to this document.
	DataSourceRefs []SigactNewBulkParamsBodyRelatedDocDataSourceRef `json:"dataSourceRefs,omitzero"`
	paramObj
}

func (r SigactNewBulkParamsBodyRelatedDoc) MarshalJSON() (data []byte, err error) {
	type shadow SigactNewBulkParamsBodyRelatedDoc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SigactNewBulkParamsBodyRelatedDoc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactNewBulkParamsBodyRelatedDocDataSourceRef struct {
	// Data source id.
	DataSourceID param.Opt[string] `json:"dataSourceId,omitzero"`
	// end position.
	EndPosition param.Opt[string] `json:"endPosition,omitzero"`
	// paragraph number.
	ParagraphNumber param.Opt[string] `json:"paragraphNumber,omitzero"`
	// sentence number.
	SentenceNumber param.Opt[string] `json:"sentenceNumber,omitzero"`
	// start position.
	StartPosition param.Opt[string] `json:"startPosition,omitzero"`
	paramObj
}

func (r SigactNewBulkParamsBodyRelatedDocDataSourceRef) MarshalJSON() (data []byte, err error) {
	type shadow SigactNewBulkParamsBodyRelatedDocDataSourceRef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SigactNewBulkParamsBodyRelatedDocDataSourceRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SigactTupleParams struct {
	// Comma-separated list of valid field names for this data type to be returned in
	// the response. Only the fields specified will be returned as well as the
	// classification marking of the data, if applicable. See the ‘queryhelp’ operation
	// for a complete list of possible fields.
	Columns string `query:"columns,required" json:"-"`
	// Date of the report or filing. (YYYY-MM-DDTHH:MM:SS.sssZ)
	ReportDate  time.Time        `query:"reportDate,required" format:"date-time" json:"-"`
	FirstResult param.Opt[int64] `query:"firstResult,omitzero" json:"-"`
	MaxResults  param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SigactTupleParams]'s query parameters as `url.Values`.
func (r SigactTupleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SigactUploadZipParams struct {
	// Zip file containing files described in the specification
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r SigactUploadZipParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
