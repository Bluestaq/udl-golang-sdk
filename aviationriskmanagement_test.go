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

func TestAviationriskmanagementNewWithOptionalParams(t *testing.T) {
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
	err := client.Aviationriskmanagement.New(context.TODO(), unifieddatalibrary.AviationriskmanagementNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AviationriskmanagementNewParamsDataModeTest,
		IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecord{{
			MissionDate:     time.Now(),
			AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
			ApprovalPending: unifieddatalibrary.Bool(true),
			Approved:        unifieddatalibrary.Bool(false),
			AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
				ApprovalDate: unifieddatalibrary.Time(time.Now()),
				ApprovedBy:   unifieddatalibrary.String("John Smith"),
				ApprovedCode: unifieddatalibrary.Int(0),
				AviationRiskManagementSortie: []unifieddatalibrary.AviationriskmanagementNewParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
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

func TestAviationriskmanagementGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Aviationriskmanagement.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AviationriskmanagementGetParams{
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

func TestAviationriskmanagementUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Aviationriskmanagement.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AviationriskmanagementUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AviationriskmanagementUpdateParamsDataModeTest,
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationriskmanagementUpdateParamsAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
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

func TestAviationriskmanagementDelete(t *testing.T) {
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
	err := client.Aviationriskmanagement.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationriskmanagementCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Aviationriskmanagement.Count(context.TODO(), unifieddatalibrary.AviationriskmanagementCountParams{
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

func TestAviationriskmanagementNewBulk(t *testing.T) {
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
	err := client.Aviationriskmanagement.NewBulk(context.TODO(), unifieddatalibrary.AviationriskmanagementNewBulkParams{
		Body: []unifieddatalibrary.AviationriskmanagementNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationriskmanagementNewBulkParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
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

func TestAviationriskmanagementQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Aviationriskmanagement.Query(context.TODO(), unifieddatalibrary.AviationriskmanagementQueryParams{
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

func TestAviationriskmanagementQueryHelp(t *testing.T) {
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
	err := client.Aviationriskmanagement.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAviationriskmanagementTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Aviationriskmanagement.Tuple(context.TODO(), unifieddatalibrary.AviationriskmanagementTupleParams{
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

func TestAviationriskmanagementUnvalidatedPublish(t *testing.T) {
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
	err := client.Aviationriskmanagement.UnvalidatedPublish(context.TODO(), unifieddatalibrary.AviationriskmanagementUnvalidatedPublishParams{
		Body: []unifieddatalibrary.AviationriskmanagementUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDMission:             "fa18d96e-91ea-60da-a7a8-1af6500066c8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AviationRiskManagementWorksheetRecord: []unifieddatalibrary.AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecord{{
				MissionDate:     time.Now(),
				AircraftMds:     unifieddatalibrary.String("E-2C HAWKEYE"),
				ApprovalPending: unifieddatalibrary.Bool(true),
				Approved:        unifieddatalibrary.Bool(false),
				AviationRiskManagementWorksheetScore: []unifieddatalibrary.AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScore{{
					ApprovalDate: unifieddatalibrary.Time(time.Now()),
					ApprovedBy:   unifieddatalibrary.String("John Smith"),
					ApprovedCode: unifieddatalibrary.Int(0),
					AviationRiskManagementSortie: []unifieddatalibrary.AviationriskmanagementUnvalidatedPublishParamsBodyAviationRiskManagementWorksheetRecordAviationRiskManagementWorksheetScoreAviationRiskManagementSortie{{
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
