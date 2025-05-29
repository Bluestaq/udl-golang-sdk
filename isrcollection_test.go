// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestIsrCollectionListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.IsrCollections.List(context.TODO(), unifieddatalibrary.IsrCollectionListParams{
		CreatedAt:   time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIsrCollectionCountWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.IsrCollections.Count(context.TODO(), unifieddatalibrary.IsrCollectionCountParams{
		CreatedAt:   time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIsrCollectionNewBulk(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	err := client.IsrCollections.NewBulk(context.TODO(), unifieddatalibrary.IsrCollectionNewBulkParams{
		Body: []unifieddatalibrary.IsrCollectionNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ISRCOLLECTION-ID"),
			CollectionRequirements: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyCollectionRequirement{{
				ID:          unifieddatalibrary.String("ISCRCOLLECTIONREQUIREMENTS"),
				Country:     unifieddatalibrary.String("VE"),
				CridNumbers: unifieddatalibrary.String("CRID"),
				CriticalTimes: unifieddatalibrary.IsrCollectionNewBulkParamsBodyCollectionRequirementCriticalTimes{
					EarliestImagingTime: time.Now(),
					LatestImagingTime:   time.Now(),
				},
				Emphasized: unifieddatalibrary.Bool(false),
				ExploitationRequirement: unifieddatalibrary.IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirement{
					ID:            unifieddatalibrary.String("ISRCOLLECTIONEXPLOITATIONREQUIREMENT"),
					Amplification: unifieddatalibrary.String("AMPLIFICATION"),
					Dissemination: unifieddatalibrary.String("EMAILS"),
					Eei:           unifieddatalibrary.String("ESSENTIAL_ELEMENTS"),
					Poc: unifieddatalibrary.IsrCollectionNewBulkParamsBodyCollectionRequirementExploitationRequirementPoc{
						ID:             unifieddatalibrary.String("ISRCOLLECTIONPOC-ID"),
						Callsign:       unifieddatalibrary.String("CALLSIGN"),
						ChatName:       unifieddatalibrary.String("CHAT_NAME"),
						ChatSystem:     unifieddatalibrary.String("CHAT"),
						Email:          unifieddatalibrary.String("EMAIL"),
						Name:           unifieddatalibrary.String("NAME"),
						Notes:          unifieddatalibrary.String("NOTES"),
						Phone:          unifieddatalibrary.String("PHONE"),
						RadioFrequency: unifieddatalibrary.Float(123.23),
						Unit:           unifieddatalibrary.String("UNIT"),
					},
					ReportingCriteria: unifieddatalibrary.String("CRITERIA"),
				},
				Hash:               unifieddatalibrary.String("HASH"),
				IntelDiscipline:    unifieddatalibrary.String("Sig"),
				IsPrismCr:          unifieddatalibrary.Bool(true),
				Operation:          unifieddatalibrary.String("NAME"),
				Priority:           unifieddatalibrary.Float(20.23),
				ReconSurvey:        unifieddatalibrary.String("SURVEY_INFO"),
				RecordID:           unifieddatalibrary.String("RECORD-ID"),
				Region:             unifieddatalibrary.String("REGION"),
				Secondary:          unifieddatalibrary.Bool(false),
				SpecialComGuidance: unifieddatalibrary.String("TEXT"),
				Start:              unifieddatalibrary.Time(time.Now()),
				Stop:               unifieddatalibrary.Time(time.Now()),
				Subregion:          unifieddatalibrary.String("SUBREGION"),
				SupportedUnit:      unifieddatalibrary.String("UNIT"),
				TargetList:         []string{"string"},
				Type:               unifieddatalibrary.String("COLLECTION_TYPE"),
			}},
			IdexVersion:                     unifieddatalibrary.Int(2),
			MissionAor:                      unifieddatalibrary.String("Kandahar"),
			MissionCollectionArea:           unifieddatalibrary.String("Example collection area"),
			MissionCountry:                  unifieddatalibrary.String("US"),
			MissionEmphasis:                 unifieddatalibrary.String("Mission emphasis"),
			MissionID:                       unifieddatalibrary.String("myTask-2020-04-23T00:00:00.000Z"),
			MissionJoa:                      unifieddatalibrary.String("Operation area"),
			MissionOperation:                unifieddatalibrary.String("OP-HONEY-BADGER"),
			MissionPrimaryIntelDiscipline:   unifieddatalibrary.String("Sig"),
			MissionPrimarySubCategory:       unifieddatalibrary.String("FMV"),
			MissionPriority:                 unifieddatalibrary.Int(1),
			MissionRegion:                   unifieddatalibrary.String("Example Region"),
			MissionRole:                     unifieddatalibrary.String("Targeting of Lead Vehicle"),
			MissionSecondaryIntelDiscipline: unifieddatalibrary.String("Intelligence_2"),
			MissionSecondarySubCategory:     unifieddatalibrary.String("Convoy"),
			MissionStartPointLat:            unifieddatalibrary.Float(45.23),
			MissionStartPointLong:           unifieddatalibrary.Float(80.23),
			MissionSubRegion:                unifieddatalibrary.String("Example Subregion"),
			MissionSupportedUnit:            unifieddatalibrary.String("ENVOYS"),
			MissionSyncMatrixBin:            unifieddatalibrary.String("MATRIX"),
			Name:                            unifieddatalibrary.String("Example mission name"),
			Origin:                          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Taskings: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyTasking{{
				ID: unifieddatalibrary.String("ISRCOLLECTIONTASKINGS-ID"),
				CollectionPeriods: unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingCollectionPeriods{
					Actual: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsActual{{
						ID:    unifieddatalibrary.String("ISRCOLLECTIONACTUAL-ID"),
						Start: unifieddatalibrary.Time(time.Now()),
						Stop:  unifieddatalibrary.Time(time.Now()),
					}},
					Planned: unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlanned{
						Additional: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingCollectionPeriodsPlannedAdditional{{
							ID:    unifieddatalibrary.String("ISRCOLLECTIONADDITIONAL"),
							Start: unifieddatalibrary.Time(time.Now()),
							Stop:  unifieddatalibrary.Time(time.Now()),
						}},
						Start: unifieddatalibrary.Time(time.Now()),
						Stop:  unifieddatalibrary.Time(time.Now()),
					},
				},
				CollectionType:        "Simultaneous",
				EightLine:             unifieddatalibrary.String("eightLine"),
				SpecialComGuidance:    unifieddatalibrary.String("TEXT"),
				SroTrack:              unifieddatalibrary.String("SRO"),
				TaskingAor:            unifieddatalibrary.String("Kandahar"),
				TaskingCollectionArea: unifieddatalibrary.String("AREA"),
				TaskingCollectionRequirements: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirement{{
					ID:          unifieddatalibrary.String("ISCRCOLLECTIONREQUIREMENTS"),
					Country:     unifieddatalibrary.String("VE"),
					CridNumbers: unifieddatalibrary.String("CRID"),
					CriticalTimes: unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementCriticalTimes{
						EarliestImagingTime: time.Now(),
						LatestImagingTime:   time.Now(),
					},
					Emphasized: unifieddatalibrary.Bool(false),
					ExploitationRequirement: unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement{
						ID:            unifieddatalibrary.String("ISRCOLLECTIONEXPLOITATIONREQUIREMENT"),
						Amplification: unifieddatalibrary.String("AMPLIFICATION"),
						Dissemination: unifieddatalibrary.String("EMAILS"),
						Eei:           unifieddatalibrary.String("ESSENTIAL_ELEMENTS"),
						Poc: unifieddatalibrary.IsrCollectionNewBulkParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc{
							ID:             unifieddatalibrary.String("ISRCOLLECTIONPOC-ID"),
							Callsign:       unifieddatalibrary.String("CALLSIGN"),
							ChatName:       unifieddatalibrary.String("CHAT_NAME"),
							ChatSystem:     unifieddatalibrary.String("CHAT"),
							Email:          unifieddatalibrary.String("EMAIL"),
							Name:           unifieddatalibrary.String("NAME"),
							Notes:          unifieddatalibrary.String("NOTES"),
							Phone:          unifieddatalibrary.String("PHONE"),
							RadioFrequency: unifieddatalibrary.Float(123.23),
							Unit:           unifieddatalibrary.String("UNIT"),
						},
						ReportingCriteria: unifieddatalibrary.String("CRITERIA"),
					},
					Hash:               unifieddatalibrary.String("HASH"),
					IntelDiscipline:    unifieddatalibrary.String("Sig"),
					IsPrismCr:          unifieddatalibrary.Bool(true),
					Operation:          unifieddatalibrary.String("NAME"),
					Priority:           unifieddatalibrary.Float(20.23),
					ReconSurvey:        unifieddatalibrary.String("SURVEY_INFO"),
					RecordID:           unifieddatalibrary.String("RECORD-ID"),
					Region:             unifieddatalibrary.String("REGION"),
					Secondary:          unifieddatalibrary.Bool(false),
					SpecialComGuidance: unifieddatalibrary.String("TEXT"),
					Start:              unifieddatalibrary.Time(time.Now()),
					Stop:               unifieddatalibrary.Time(time.Now()),
					Subregion:          unifieddatalibrary.String("SUBREGION"),
					SupportedUnit:      unifieddatalibrary.String("UNIT"),
					TargetList:         []string{"string"},
					Type:               unifieddatalibrary.String("COLLECTION_TYPE"),
				}},
				TaskingCountry:                  unifieddatalibrary.String("CODE"),
				TaskingEmphasis:                 unifieddatalibrary.String("EMPHASIS"),
				TaskingJoa:                      unifieddatalibrary.String("AREA"),
				TaskingOperation:                unifieddatalibrary.String("OP-HONEY-BADGER"),
				TaskingPrimaryIntelDiscipline:   unifieddatalibrary.String("Sig"),
				TaskingPrimarySubCategory:       unifieddatalibrary.String("FMV"),
				TaskingPriority:                 unifieddatalibrary.Float(10.23),
				TaskingRegion:                   unifieddatalibrary.String("REGION"),
				TaskingRetaskTime:               unifieddatalibrary.Time(time.Now()),
				TaskingRole:                     unifieddatalibrary.String("Track Lead Vehicle"),
				TaskingSecondaryIntelDiscipline: unifieddatalibrary.String("Intelligence_2"),
				TaskingSecondarySubCategory:     unifieddatalibrary.String("Convoy"),
				TaskingStartPointLat:            unifieddatalibrary.Float(45.23),
				TaskingStartPointLong:           unifieddatalibrary.Float(45.23),
				TaskingSubRegion:                unifieddatalibrary.String("SUBREGION"),
				TaskingSupportedUnit:            unifieddatalibrary.String("ENVOYS"),
				TaskingSyncMatrixBin:            unifieddatalibrary.String("MATRIX"),
				Type:                            "Deliberate",
			}},
			Transit: []unifieddatalibrary.IsrCollectionNewBulkParamsBodyTransit{{
				ID:       unifieddatalibrary.String("ISRCOLLECTIONTRANSIT-ID"),
				Base:     unifieddatalibrary.String("ENVOYS"),
				Duration: unifieddatalibrary.Float(200.23),
			}},
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIsrCollectionQueryhelp(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.IsrCollections.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIsrCollectionTupleWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.IsrCollections.Tuple(context.TODO(), unifieddatalibrary.IsrCollectionTupleParams{
		Columns:     "columns",
		CreatedAt:   time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIsrCollectionUnvalidatedPublish(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	err := client.IsrCollections.UnvalidatedPublish(context.TODO(), unifieddatalibrary.IsrCollectionUnvalidatedPublishParams{
		Body: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ISRCOLLECTION-ID"),
			CollectionRequirements: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirement{{
				ID:          unifieddatalibrary.String("ISCRCOLLECTIONREQUIREMENTS"),
				Country:     unifieddatalibrary.String("VE"),
				CridNumbers: unifieddatalibrary.String("CRID"),
				CriticalTimes: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementCriticalTimes{
					EarliestImagingTime: time.Now(),
					LatestImagingTime:   time.Now(),
				},
				Emphasized: unifieddatalibrary.Bool(false),
				ExploitationRequirement: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirement{
					ID:            unifieddatalibrary.String("ISRCOLLECTIONEXPLOITATIONREQUIREMENT"),
					Amplification: unifieddatalibrary.String("AMPLIFICATION"),
					Dissemination: unifieddatalibrary.String("EMAILS"),
					Eei:           unifieddatalibrary.String("ESSENTIAL_ELEMENTS"),
					Poc: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyCollectionRequirementExploitationRequirementPoc{
						ID:             unifieddatalibrary.String("ISRCOLLECTIONPOC-ID"),
						Callsign:       unifieddatalibrary.String("CALLSIGN"),
						ChatName:       unifieddatalibrary.String("CHAT_NAME"),
						ChatSystem:     unifieddatalibrary.String("CHAT"),
						Email:          unifieddatalibrary.String("EMAIL"),
						Name:           unifieddatalibrary.String("NAME"),
						Notes:          unifieddatalibrary.String("NOTES"),
						Phone:          unifieddatalibrary.String("PHONE"),
						RadioFrequency: unifieddatalibrary.Float(123.23),
						Unit:           unifieddatalibrary.String("UNIT"),
					},
					ReportingCriteria: unifieddatalibrary.String("CRITERIA"),
				},
				Hash:               unifieddatalibrary.String("HASH"),
				IntelDiscipline:    unifieddatalibrary.String("Sig"),
				IsPrismCr:          unifieddatalibrary.Bool(true),
				Operation:          unifieddatalibrary.String("NAME"),
				Priority:           unifieddatalibrary.Float(20.23),
				ReconSurvey:        unifieddatalibrary.String("SURVEY_INFO"),
				RecordID:           unifieddatalibrary.String("RECORD-ID"),
				Region:             unifieddatalibrary.String("REGION"),
				Secondary:          unifieddatalibrary.Bool(false),
				SpecialComGuidance: unifieddatalibrary.String("TEXT"),
				Start:              unifieddatalibrary.Time(time.Now()),
				Stop:               unifieddatalibrary.Time(time.Now()),
				Subregion:          unifieddatalibrary.String("SUBREGION"),
				SupportedUnit:      unifieddatalibrary.String("UNIT"),
				TargetList:         []string{"string"},
				Type:               unifieddatalibrary.String("COLLECTION_TYPE"),
			}},
			IdexVersion:                     unifieddatalibrary.Int(2),
			MissionAor:                      unifieddatalibrary.String("Kandahar"),
			MissionCollectionArea:           unifieddatalibrary.String("Example collection area"),
			MissionCountry:                  unifieddatalibrary.String("US"),
			MissionEmphasis:                 unifieddatalibrary.String("Mission emphasis"),
			MissionID:                       unifieddatalibrary.String("myTask-2020-04-23T00:00:00.000Z"),
			MissionJoa:                      unifieddatalibrary.String("Operation area"),
			MissionOperation:                unifieddatalibrary.String("OP-HONEY-BADGER"),
			MissionPrimaryIntelDiscipline:   unifieddatalibrary.String("Sig"),
			MissionPrimarySubCategory:       unifieddatalibrary.String("FMV"),
			MissionPriority:                 unifieddatalibrary.Int(1),
			MissionRegion:                   unifieddatalibrary.String("Example Region"),
			MissionRole:                     unifieddatalibrary.String("Targeting of Lead Vehicle"),
			MissionSecondaryIntelDiscipline: unifieddatalibrary.String("Intelligence_2"),
			MissionSecondarySubCategory:     unifieddatalibrary.String("Convoy"),
			MissionStartPointLat:            unifieddatalibrary.Float(45.23),
			MissionStartPointLong:           unifieddatalibrary.Float(80.23),
			MissionSubRegion:                unifieddatalibrary.String("Example Subregion"),
			MissionSupportedUnit:            unifieddatalibrary.String("ENVOYS"),
			MissionSyncMatrixBin:            unifieddatalibrary.String("MATRIX"),
			Name:                            unifieddatalibrary.String("Example mission name"),
			Origin:                          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Taskings: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTasking{{
				ID: unifieddatalibrary.String("ISRCOLLECTIONTASKINGS-ID"),
				CollectionPeriods: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriods{
					Actual: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsActual{{
						ID:    unifieddatalibrary.String("ISRCOLLECTIONACTUAL-ID"),
						Start: unifieddatalibrary.Time(time.Now()),
						Stop:  unifieddatalibrary.Time(time.Now()),
					}},
					Planned: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlanned{
						Additional: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingCollectionPeriodsPlannedAdditional{{
							ID:    unifieddatalibrary.String("ISRCOLLECTIONADDITIONAL"),
							Start: unifieddatalibrary.Time(time.Now()),
							Stop:  unifieddatalibrary.Time(time.Now()),
						}},
						Start: unifieddatalibrary.Time(time.Now()),
						Stop:  unifieddatalibrary.Time(time.Now()),
					},
				},
				CollectionType:        "Simultaneous",
				EightLine:             unifieddatalibrary.String("eightLine"),
				SpecialComGuidance:    unifieddatalibrary.String("TEXT"),
				SroTrack:              unifieddatalibrary.String("SRO"),
				TaskingAor:            unifieddatalibrary.String("Kandahar"),
				TaskingCollectionArea: unifieddatalibrary.String("AREA"),
				TaskingCollectionRequirements: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirement{{
					ID:          unifieddatalibrary.String("ISCRCOLLECTIONREQUIREMENTS"),
					Country:     unifieddatalibrary.String("VE"),
					CridNumbers: unifieddatalibrary.String("CRID"),
					CriticalTimes: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementCriticalTimes{
						EarliestImagingTime: time.Now(),
						LatestImagingTime:   time.Now(),
					},
					Emphasized: unifieddatalibrary.Bool(false),
					ExploitationRequirement: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirement{
						ID:            unifieddatalibrary.String("ISRCOLLECTIONEXPLOITATIONREQUIREMENT"),
						Amplification: unifieddatalibrary.String("AMPLIFICATION"),
						Dissemination: unifieddatalibrary.String("EMAILS"),
						Eei:           unifieddatalibrary.String("ESSENTIAL_ELEMENTS"),
						Poc: unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTaskingTaskingCollectionRequirementExploitationRequirementPoc{
							ID:             unifieddatalibrary.String("ISRCOLLECTIONPOC-ID"),
							Callsign:       unifieddatalibrary.String("CALLSIGN"),
							ChatName:       unifieddatalibrary.String("CHAT_NAME"),
							ChatSystem:     unifieddatalibrary.String("CHAT"),
							Email:          unifieddatalibrary.String("EMAIL"),
							Name:           unifieddatalibrary.String("NAME"),
							Notes:          unifieddatalibrary.String("NOTES"),
							Phone:          unifieddatalibrary.String("PHONE"),
							RadioFrequency: unifieddatalibrary.Float(123.23),
							Unit:           unifieddatalibrary.String("UNIT"),
						},
						ReportingCriteria: unifieddatalibrary.String("CRITERIA"),
					},
					Hash:               unifieddatalibrary.String("HASH"),
					IntelDiscipline:    unifieddatalibrary.String("Sig"),
					IsPrismCr:          unifieddatalibrary.Bool(true),
					Operation:          unifieddatalibrary.String("NAME"),
					Priority:           unifieddatalibrary.Float(20.23),
					ReconSurvey:        unifieddatalibrary.String("SURVEY_INFO"),
					RecordID:           unifieddatalibrary.String("RECORD-ID"),
					Region:             unifieddatalibrary.String("REGION"),
					Secondary:          unifieddatalibrary.Bool(false),
					SpecialComGuidance: unifieddatalibrary.String("TEXT"),
					Start:              unifieddatalibrary.Time(time.Now()),
					Stop:               unifieddatalibrary.Time(time.Now()),
					Subregion:          unifieddatalibrary.String("SUBREGION"),
					SupportedUnit:      unifieddatalibrary.String("UNIT"),
					TargetList:         []string{"string"},
					Type:               unifieddatalibrary.String("COLLECTION_TYPE"),
				}},
				TaskingCountry:                  unifieddatalibrary.String("CODE"),
				TaskingEmphasis:                 unifieddatalibrary.String("EMPHASIS"),
				TaskingJoa:                      unifieddatalibrary.String("AREA"),
				TaskingOperation:                unifieddatalibrary.String("OP-HONEY-BADGER"),
				TaskingPrimaryIntelDiscipline:   unifieddatalibrary.String("Sig"),
				TaskingPrimarySubCategory:       unifieddatalibrary.String("FMV"),
				TaskingPriority:                 unifieddatalibrary.Float(10.23),
				TaskingRegion:                   unifieddatalibrary.String("REGION"),
				TaskingRetaskTime:               unifieddatalibrary.Time(time.Now()),
				TaskingRole:                     unifieddatalibrary.String("Track Lead Vehicle"),
				TaskingSecondaryIntelDiscipline: unifieddatalibrary.String("Intelligence_2"),
				TaskingSecondarySubCategory:     unifieddatalibrary.String("Convoy"),
				TaskingStartPointLat:            unifieddatalibrary.Float(45.23),
				TaskingStartPointLong:           unifieddatalibrary.Float(45.23),
				TaskingSubRegion:                unifieddatalibrary.String("SUBREGION"),
				TaskingSupportedUnit:            unifieddatalibrary.String("ENVOYS"),
				TaskingSyncMatrixBin:            unifieddatalibrary.String("MATRIX"),
				Type:                            "Deliberate",
			}},
			Transit: []unifieddatalibrary.IsrCollectionUnvalidatedPublishParamsBodyTransit{{
				ID:       unifieddatalibrary.String("ISRCOLLECTIONTRANSIT-ID"),
				Base:     unifieddatalibrary.String("ENVOYS"),
				Duration: unifieddatalibrary.Float(200.23),
			}},
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
