// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestAviationRiskManagementNewWithOptionalParams(t *testing.T) {
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
	err := client.AviationRiskManagement.New(context.TODO(), unifieddatalibrary.AviationRiskManagementNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AviationRiskManagementNewParamsDataModeTest,
		IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecord{{
			MissionDate:     time.Now(),
			AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
			ApprovalPending: unifieddatalibrary.Bool(true),
			Approved:        unifieddatalibrary.Bool(false),
			AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
				ApprovalDate: unifieddatalibrary.Time(time.Now()),
				ApprovedBy:   unifieddatalibrary.String("John Smith"),
				ApprovedCode: unifieddatalibrary.Int(0),
				AviationRiskManagementSortie: []unifieddatalibrary.AviationRiskManagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
					ExtSortieID: unifieddatalibrary.String("MB014313032022407540"),
					IDSortie:    unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
					LegNum:      unifieddatalibrary.Int(100),
					SortieScore: unifieddatalibrary.Int(3),
				}},
				ExtScoreID:      unifieddatalibrary.String("BM022301191649232740"),
				RiskCategory:    unifieddatalibrary.String("Crew/Aircraft"),
				RiskDescription: unifieddatalibrary.String("Upgrade training"),
				RiskKey:         unifieddatalibrary.String("26"),
				RiskName:        unifieddatalibrary.String("Crew Qualification"),
				Score:           unifieddatalibrary.Int(1),
				ScoreRemark:     unifieddatalibrary.String("Worksheet score remark."),
			}},
			DispositionComments: unifieddatalibrary.String("Disposition comment."),
			ExtRecordID:         unifieddatalibrary.String("B022401191649232716"),
			Itinerary:           unifieddatalibrary.String("RJTY-PGUA-RJTY"),
			LastUpdatedAt:       unifieddatalibrary.Time(time.Now()),
			Remarks:             unifieddatalibrary.String("Worksheet record remark."),
			SeverityLevel:       unifieddatalibrary.Int(0),
			SubmissionDate:      unifieddatalibrary.Time(time.Now()),
			TierNumber:          unifieddatalibrary.Int(2),
			TotalScore:          unifieddatalibrary.Int(11),
			UserID:              unifieddatalibrary.String("TIER0SCORING"),
		}},
		ExtMissionID:  unifieddatalibrary.String("MCD04250106123509230"),
		MissionNumber: unifieddatalibrary.String("LVM134412001"),
		OrgID:         unifieddatalibrary.String("50000002"),
		Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		UnitID:        unifieddatalibrary.String("63"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationRiskManagementGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AviationRiskManagement.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AviationRiskManagementGetParams{
			FirstResult: unifieddatalibrary.Int(0),
			MaxResults:  unifieddatalibrary.Int(0),
		},
	)
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationRiskManagementUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AviationRiskManagement.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AviationRiskManagementUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AviationRiskManagementUpdateParamsDataModeTest,
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationRiskManagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
						ExtSortieID: unifieddatalibrary.String("MB014313032022407540"),
						IDSortie:    unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
						LegNum:      unifieddatalibrary.Int(100),
						SortieScore: unifieddatalibrary.Int(3),
					}},
					ExtScoreID:      unifieddatalibrary.String("BM022301191649232740"),
					RiskCategory:    unifieddatalibrary.String("Crew/Aircraft"),
					RiskDescription: unifieddatalibrary.String("Upgrade training"),
					RiskKey:         unifieddatalibrary.String("26"),
					RiskName:        unifieddatalibrary.String("Crew Qualification"),
					Score:           unifieddatalibrary.Int(1),
					ScoreRemark:     unifieddatalibrary.String("Worksheet score remark."),
				}},
				DispositionComments: unifieddatalibrary.String("Disposition comment."),
				ExtRecordID:         unifieddatalibrary.String("B022401191649232716"),
				Itinerary:           unifieddatalibrary.String("RJTY-PGUA-RJTY"),
				LastUpdatedAt:       unifieddatalibrary.Time(time.Now()),
				Remarks:             unifieddatalibrary.String("Worksheet record remark."),
				SeverityLevel:       unifieddatalibrary.Int(0),
				SubmissionDate:      unifieddatalibrary.Time(time.Now()),
				TierNumber:          unifieddatalibrary.Int(2),
				TotalScore:          unifieddatalibrary.Int(11),
				UserID:              unifieddatalibrary.String("TIER0SCORING"),
			}},
			ExtMissionID:  unifieddatalibrary.String("MCD04250106123509230"),
			MissionNumber: unifieddatalibrary.String("LVM134412001"),
			OrgID:         unifieddatalibrary.String("50000002"),
			Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			UnitID:        unifieddatalibrary.String("63"),
		},
	)
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationRiskManagementListWithOptionalParams(t *testing.T) {
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
	_, err := client.AviationRiskManagement.List(context.TODO(), unifieddatalibrary.AviationRiskManagementListParams{
		IDMission:   "idMission",
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

func TestAviationRiskManagementDelete(t *testing.T) {
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
	err := client.AviationRiskManagement.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationRiskManagementCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AviationRiskManagement.Count(context.TODO(), unifieddatalibrary.AviationRiskManagementCountParams{
		IDMission:   "idMission",
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

func TestAviationRiskManagementNewBulk(t *testing.T) {
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
	err := client.AviationRiskManagement.NewBulk(context.TODO(), unifieddatalibrary.AviationRiskManagementNewBulkParams{
		Body: []unifieddatalibrary.AviationRiskManagementNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationRiskManagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
						ExtSortieID: unifieddatalibrary.String("MB014313032022407540"),
						IDSortie:    unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
						LegNum:      unifieddatalibrary.Int(100),
						SortieScore: unifieddatalibrary.Int(3),
					}},
					ExtScoreID:      unifieddatalibrary.String("BM022301191649232740"),
					RiskCategory:    unifieddatalibrary.String("Crew/Aircraft"),
					RiskDescription: unifieddatalibrary.String("Upgrade training"),
					RiskKey:         unifieddatalibrary.String("26"),
					RiskName:        unifieddatalibrary.String("Crew Qualification"),
					Score:           unifieddatalibrary.Int(1),
					ScoreRemark:     unifieddatalibrary.String("Worksheet score remark."),
				}},
				DispositionComments: unifieddatalibrary.String("Disposition comment."),
				ExtRecordID:         unifieddatalibrary.String("B022401191649232716"),
				Itinerary:           unifieddatalibrary.String("RJTY-PGUA-RJTY"),
				LastUpdatedAt:       unifieddatalibrary.Time(time.Now()),
				Remarks:             unifieddatalibrary.String("Worksheet record remark."),
				SeverityLevel:       unifieddatalibrary.Int(0),
				SubmissionDate:      unifieddatalibrary.Time(time.Now()),
				TierNumber:          unifieddatalibrary.Int(2),
				TotalScore:          unifieddatalibrary.Int(11),
				UserID:              unifieddatalibrary.String("TIER0SCORING"),
			}},
			ExtMissionID:  unifieddatalibrary.String("MCD04250106123509230"),
			MissionNumber: unifieddatalibrary.String("LVM134412001"),
			OrgID:         unifieddatalibrary.String("50000002"),
			Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			UnitID:        unifieddatalibrary.String("63"),
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

func TestAviationRiskManagementQueryHelp(t *testing.T) {
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
	_, err := client.AviationRiskManagement.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationRiskManagementTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AviationRiskManagement.Tuple(context.TODO(), unifieddatalibrary.AviationRiskManagementTupleParams{
		Columns:     "columns",
		IDMission:   "idMission",
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

func TestAviationRiskManagementUnvalidatedPublish(t *testing.T) {
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
	err := client.AviationRiskManagement.UnvalidatedPublish(context.TODO(), unifieddatalibrary.AviationRiskManagementUnvalidatedPublishParams{
		Body: []unifieddatalibrary.AviationRiskManagementUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationRiskManagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
						ExtSortieID: unifieddatalibrary.String("MB014313032022407540"),
						IDSortie:    unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
						LegNum:      unifieddatalibrary.Int(100),
						SortieScore: unifieddatalibrary.Int(3),
					}},
					ExtScoreID:      unifieddatalibrary.String("BM022301191649232740"),
					RiskCategory:    unifieddatalibrary.String("Crew/Aircraft"),
					RiskDescription: unifieddatalibrary.String("Upgrade training"),
					RiskKey:         unifieddatalibrary.String("26"),
					RiskName:        unifieddatalibrary.String("Crew Qualification"),
					Score:           unifieddatalibrary.Int(1),
					ScoreRemark:     unifieddatalibrary.String("Worksheet score remark."),
				}},
				DispositionComments: unifieddatalibrary.String("Disposition comment."),
				ExtRecordID:         unifieddatalibrary.String("B022401191649232716"),
				Itinerary:           unifieddatalibrary.String("RJTY-PGUA-RJTY"),
				LastUpdatedAt:       unifieddatalibrary.Time(time.Now()),
				Remarks:             unifieddatalibrary.String("Worksheet record remark."),
				SeverityLevel:       unifieddatalibrary.Int(0),
				SubmissionDate:      unifieddatalibrary.Time(time.Now()),
				TierNumber:          unifieddatalibrary.Int(2),
				TotalScore:          unifieddatalibrary.Int(11),
				UserID:              unifieddatalibrary.String("TIER0SCORING"),
			}},
			ExtMissionID:  unifieddatalibrary.String("MCD04250106123509230"),
			MissionNumber: unifieddatalibrary.String("LVM134412001"),
			OrgID:         unifieddatalibrary.String("50000002"),
			Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			UnitID:        unifieddatalibrary.String("63"),
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
