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

func TestLogisticssupportNewWithOptionalParams(t *testing.T) {
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
	err := client.Logisticssupport.New(context.TODO(), unifieddatalibrary.LogisticssupportNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LogisticssupportNewParamsDataModeTest,
		RptCreatedTime:        time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
		AircraftMds:           unifieddatalibrary.String("CO17A"),
		CurrIcao:              unifieddatalibrary.String("KCOS"),
		Etic:                  unifieddatalibrary.Time(time.Now()),
		Etmc:                  unifieddatalibrary.Time(time.Now()),
		ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
		LogisticAction:        unifieddatalibrary.String("WA"),
		LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticssupportNewParamsLogisticsDiscrepancyInfo{{
			ClosureTime:     unifieddatalibrary.Time(time.Now()),
			DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
			Jcn:             unifieddatalibrary.String("231942400"),
			JobStTime:       unifieddatalibrary.Time(time.Now()),
		}},
		LogisticsRecordID: unifieddatalibrary.String("L62017"),
		LogisticsRemarks: []unifieddatalibrary.LogisticssupportNewParamsLogisticsRemark{{
			LastChanged: unifieddatalibrary.Time(time.Now()),
			Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
			Username:    unifieddatalibrary.String("JSMITH"),
		}},
		LogisticsSupportItems: []unifieddatalibrary.LogisticssupportNewParamsLogisticsSupportItem{{
			Cannibalized:        unifieddatalibrary.Bool(true),
			DeployPlanNumber:    unifieddatalibrary.String("T89003"),
			Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
			ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
			JobControlNumber:    unifieddatalibrary.String("231942400"),
			LogisticsParts: []unifieddatalibrary.LogisticssupportNewParamsLogisticsSupportItemLogisticsPart{{
				FigureNumber:     unifieddatalibrary.String("3"),
				IndexNumber:      unifieddatalibrary.String("4"),
				LocationVerifier: unifieddatalibrary.String("JANE DOE"),
				LogisticsStocks: []unifieddatalibrary.LogisticssupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock{{
					Quantity:       unifieddatalibrary.Int(4),
					SourceIcao:     unifieddatalibrary.String("PHIK"),
					StockCheckTime: unifieddatalibrary.Time(time.Now()),
					StockPoc:       unifieddatalibrary.String("SMITH, JOHN J"),
				}},
				MeasurementUnitCode:  unifieddatalibrary.String("EA"),
				NationalStockNumber:  unifieddatalibrary.String("5310-00-045-3299"),
				PartNumber:           unifieddatalibrary.String("MS35338-42"),
				RequestVerifier:      unifieddatalibrary.String("JOHN SMITH"),
				SupplyDocumentNumber: unifieddatalibrary.String("J223FU31908300"),
				TechnicalOrderText:   unifieddatalibrary.String("1C-17A-4"),
				WorkUnitCode:         unifieddatalibrary.String("5611UU001"),
			}},
			LogisticsRemarks: []unifieddatalibrary.LogisticssupportNewParamsLogisticsSupportItemLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSpecialties: []unifieddatalibrary.LogisticssupportNewParamsLogisticsSupportItemLogisticsSpecialty{{
				FirstName:    unifieddatalibrary.String("JOHN"),
				Last4Ssn:     unifieddatalibrary.String("9999"),
				LastName:     unifieddatalibrary.String("SMITH"),
				RankCode:     unifieddatalibrary.String("MAJ"),
				RoleTypeCode: unifieddatalibrary.String("TC"),
				SkillLevel:   unifieddatalibrary.Int(3),
				Specialty:    unifieddatalibrary.String("ELEN"),
			}},
			Quantity:                unifieddatalibrary.Int(1),
			ReadyTime:               unifieddatalibrary.Time(time.Now()),
			ReceivedTime:            unifieddatalibrary.Time(time.Now()),
			RecoveryRequestTypeCode: unifieddatalibrary.String("E"),
			RedeployPlanNumber:      unifieddatalibrary.String("T89003"),
			RedeployShipmentUnitID:  unifieddatalibrary.String("X400LA31949108"),
			RequestNumber:           unifieddatalibrary.String("89208"),
			ResupportFlag:           unifieddatalibrary.Bool(true),
			ShipmentUnitID:          unifieddatalibrary.String("FB44273196X501XXX"),
			SiPoc:                   unifieddatalibrary.String("SMITH, JOHN J"),
			SourceIcao:              unifieddatalibrary.String("PHIK"),
		}},
		LogisticsTransportationPlans: []unifieddatalibrary.LogisticssupportNewParamsLogisticsTransportationPlan{{
			ActDepTime:             unifieddatalibrary.Time(time.Now()),
			AircraftStatus:         unifieddatalibrary.String("NMCMU"),
			ApproxArrTime:          unifieddatalibrary.Time(time.Now()),
			CancelledDate:          unifieddatalibrary.Time(time.Now()),
			ClosedDate:             unifieddatalibrary.Time(time.Now()),
			Coordinator:            unifieddatalibrary.String("SMITH, JOHN"),
			CoordinatorUnit:        unifieddatalibrary.String("TACC"),
			DestinationIcao:        unifieddatalibrary.String("YBCS"),
			Duration:               unifieddatalibrary.String("086:20"),
			EstArrTime:             unifieddatalibrary.Time(time.Now()),
			EstDepTime:             unifieddatalibrary.Time(time.Now()),
			LastChangedDate:        unifieddatalibrary.Time(time.Now()),
			LogisticMasterRecordID: unifieddatalibrary.String("L62126"),
			LogisticsSegments: []unifieddatalibrary.LogisticssupportNewParamsLogisticsTransportationPlanLogisticsSegment{{
				ArrivalIcao:    unifieddatalibrary.String("YBCS"),
				DepartureIcao:  unifieddatalibrary.String("PHIK"),
				ExtMissionID:   unifieddatalibrary.String("2001101RF01202307062205"),
				IDMission:      unifieddatalibrary.String("EXAMPLE-UUID"),
				Itin:           unifieddatalibrary.Int(200),
				MissionNumber:  unifieddatalibrary.String("TAM308901196"),
				MissionType:    unifieddatalibrary.String("SAAM"),
				ModeCode:       unifieddatalibrary.String("A"),
				SegActArrTime:  unifieddatalibrary.Time(time.Now()),
				SegActDepTime:  unifieddatalibrary.Time(time.Now()),
				SegAircraftMds: unifieddatalibrary.String("B7772E"),
				SegEstArrTime:  unifieddatalibrary.Time(time.Now()),
				SegEstDepTime:  unifieddatalibrary.Time(time.Now()),
				SegmentNumber:  unifieddatalibrary.Int(3),
				SegTailNumber:  unifieddatalibrary.String("N819AX"),
			}},
			LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticssupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			Majcom:          unifieddatalibrary.String("HQAMC"),
			MissionChange:   unifieddatalibrary.Bool(false),
			NumEnrouteStops: unifieddatalibrary.Int(4),
			NumTransLoads:   unifieddatalibrary.Int(3),
			OriginIcao:      unifieddatalibrary.String("KATL"),
			PlanDefinition:  unifieddatalibrary.String("DEPLOY"),
			PlansNumber:     unifieddatalibrary.String("T89002"),
			SerialNumber:    unifieddatalibrary.String("9009209"),
			StatusCode:      unifieddatalibrary.String("N"),
			TpAircraftMds:   unifieddatalibrary.String("C17A"),
			TpTailNumber:    unifieddatalibrary.String("99209"),
		}},
		MaintStatusCode: unifieddatalibrary.String("NMCMU"),
		McTime:          unifieddatalibrary.Time(time.Now()),
		MeTime:          unifieddatalibrary.Time(time.Now()),
		Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Owner:           unifieddatalibrary.String("EXAMPLE_OWNER"),
		ReopenFlag:      unifieddatalibrary.Bool(true),
		RptClosedTime:   unifieddatalibrary.Time(time.Now()),
		SuppIcao:        unifieddatalibrary.String("KCOS"),
		TailNumber:      unifieddatalibrary.String("99290"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogisticssupportUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Logisticssupport.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LogisticssupportUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LogisticssupportUpdateParamsDataModeTest,
			RptCreatedTime:        time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
			AircraftMds:           unifieddatalibrary.String("CO17A"),
			CurrIcao:              unifieddatalibrary.String("KCOS"),
			Etic:                  unifieddatalibrary.Time(time.Now()),
			Etmc:                  unifieddatalibrary.Time(time.Now()),
			ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
			LogisticAction:        unifieddatalibrary.String("WA"),
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock{{
						Quantity:       unifieddatalibrary.Int(4),
						SourceIcao:     unifieddatalibrary.String("PHIK"),
						StockCheckTime: unifieddatalibrary.Time(time.Now()),
						StockPoc:       unifieddatalibrary.String("SMITH, JOHN J"),
					}},
					MeasurementUnitCode:  unifieddatalibrary.String("EA"),
					NationalStockNumber:  unifieddatalibrary.String("5310-00-045-3299"),
					PartNumber:           unifieddatalibrary.String("MS35338-42"),
					RequestVerifier:      unifieddatalibrary.String("JOHN SMITH"),
					SupplyDocumentNumber: unifieddatalibrary.String("J223FU31908300"),
					TechnicalOrderText:   unifieddatalibrary.String("1C-17A-4"),
					WorkUnitCode:         unifieddatalibrary.String("5611UU001"),
				}},
				LogisticsRemarks: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsSupportItemLogisticsSpecialty{{
					FirstName:    unifieddatalibrary.String("JOHN"),
					Last4Ssn:     unifieddatalibrary.String("9999"),
					LastName:     unifieddatalibrary.String("SMITH"),
					RankCode:     unifieddatalibrary.String("MAJ"),
					RoleTypeCode: unifieddatalibrary.String("TC"),
					SkillLevel:   unifieddatalibrary.Int(3),
					Specialty:    unifieddatalibrary.String("ELEN"),
				}},
				Quantity:                unifieddatalibrary.Int(1),
				ReadyTime:               unifieddatalibrary.Time(time.Now()),
				ReceivedTime:            unifieddatalibrary.Time(time.Now()),
				RecoveryRequestTypeCode: unifieddatalibrary.String("E"),
				RedeployPlanNumber:      unifieddatalibrary.String("T89003"),
				RedeployShipmentUnitID:  unifieddatalibrary.String("X400LA31949108"),
				RequestNumber:           unifieddatalibrary.String("89208"),
				ResupportFlag:           unifieddatalibrary.Bool(true),
				ShipmentUnitID:          unifieddatalibrary.String("FB44273196X501XXX"),
				SiPoc:                   unifieddatalibrary.String("SMITH, JOHN J"),
				SourceIcao:              unifieddatalibrary.String("PHIK"),
			}},
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsTransportationPlan{{
				ActDepTime:             unifieddatalibrary.Time(time.Now()),
				AircraftStatus:         unifieddatalibrary.String("NMCMU"),
				ApproxArrTime:          unifieddatalibrary.Time(time.Now()),
				CancelledDate:          unifieddatalibrary.Time(time.Now()),
				ClosedDate:             unifieddatalibrary.Time(time.Now()),
				Coordinator:            unifieddatalibrary.String("SMITH, JOHN"),
				CoordinatorUnit:        unifieddatalibrary.String("TACC"),
				DestinationIcao:        unifieddatalibrary.String("YBCS"),
				Duration:               unifieddatalibrary.String("086:20"),
				EstArrTime:             unifieddatalibrary.Time(time.Now()),
				EstDepTime:             unifieddatalibrary.Time(time.Now()),
				LastChangedDate:        unifieddatalibrary.Time(time.Now()),
				LogisticMasterRecordID: unifieddatalibrary.String("L62126"),
				LogisticsSegments: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsSegment{{
					ArrivalIcao:    unifieddatalibrary.String("YBCS"),
					DepartureIcao:  unifieddatalibrary.String("PHIK"),
					ExtMissionID:   unifieddatalibrary.String("2001101RF01202307062205"),
					IDMission:      unifieddatalibrary.String("EXAMPLE-UUID"),
					Itin:           unifieddatalibrary.Int(200),
					MissionNumber:  unifieddatalibrary.String("TAM308901196"),
					MissionType:    unifieddatalibrary.String("SAAM"),
					ModeCode:       unifieddatalibrary.String("A"),
					SegActArrTime:  unifieddatalibrary.Time(time.Now()),
					SegActDepTime:  unifieddatalibrary.Time(time.Now()),
					SegAircraftMds: unifieddatalibrary.String("B7772E"),
					SegEstArrTime:  unifieddatalibrary.Time(time.Now()),
					SegEstDepTime:  unifieddatalibrary.Time(time.Now()),
					SegmentNumber:  unifieddatalibrary.Int(3),
					SegTailNumber:  unifieddatalibrary.String("N819AX"),
				}},
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticssupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				Majcom:          unifieddatalibrary.String("HQAMC"),
				MissionChange:   unifieddatalibrary.Bool(false),
				NumEnrouteStops: unifieddatalibrary.Int(4),
				NumTransLoads:   unifieddatalibrary.Int(3),
				OriginIcao:      unifieddatalibrary.String("KATL"),
				PlanDefinition:  unifieddatalibrary.String("DEPLOY"),
				PlansNumber:     unifieddatalibrary.String("T89002"),
				SerialNumber:    unifieddatalibrary.String("9009209"),
				StatusCode:      unifieddatalibrary.String("N"),
				TpAircraftMds:   unifieddatalibrary.String("C17A"),
				TpTailNumber:    unifieddatalibrary.String("99209"),
			}},
			MaintStatusCode: unifieddatalibrary.String("NMCMU"),
			McTime:          unifieddatalibrary.Time(time.Now()),
			MeTime:          unifieddatalibrary.Time(time.Now()),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Owner:           unifieddatalibrary.String("EXAMPLE_OWNER"),
			ReopenFlag:      unifieddatalibrary.Bool(true),
			RptClosedTime:   unifieddatalibrary.Time(time.Now()),
			SuppIcao:        unifieddatalibrary.String("KCOS"),
			TailNumber:      unifieddatalibrary.String("99290"),
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

func TestLogisticssupportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Logisticssupport.List(context.TODO(), unifieddatalibrary.LogisticssupportListParams{
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

func TestLogisticssupportCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Logisticssupport.Count(context.TODO(), unifieddatalibrary.LogisticssupportCountParams{
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

func TestLogisticssupportNewBulk(t *testing.T) {
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
	err := client.Logisticssupport.NewBulk(context.TODO(), unifieddatalibrary.LogisticssupportNewBulkParams{
		Body: []unifieddatalibrary.LogisticssupportNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			RptCreatedTime:        time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
			AircraftMds:           unifieddatalibrary.String("CO17A"),
			CurrIcao:              unifieddatalibrary.String("KCOS"),
			Etic:                  unifieddatalibrary.Time(time.Now()),
			Etmc:                  unifieddatalibrary.Time(time.Now()),
			ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
			LogisticAction:        unifieddatalibrary.String("WA"),
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock{{
						Quantity:       unifieddatalibrary.Int(4),
						SourceIcao:     unifieddatalibrary.String("PHIK"),
						StockCheckTime: unifieddatalibrary.Time(time.Now()),
						StockPoc:       unifieddatalibrary.String("SMITH, JOHN J"),
					}},
					MeasurementUnitCode:  unifieddatalibrary.String("EA"),
					NationalStockNumber:  unifieddatalibrary.String("5310-00-045-3299"),
					PartNumber:           unifieddatalibrary.String("MS35338-42"),
					RequestVerifier:      unifieddatalibrary.String("JOHN SMITH"),
					SupplyDocumentNumber: unifieddatalibrary.String("J223FU31908300"),
					TechnicalOrderText:   unifieddatalibrary.String("1C-17A-4"),
					WorkUnitCode:         unifieddatalibrary.String("5611UU001"),
				}},
				LogisticsRemarks: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty{{
					FirstName:    unifieddatalibrary.String("JOHN"),
					Last4Ssn:     unifieddatalibrary.String("9999"),
					LastName:     unifieddatalibrary.String("SMITH"),
					RankCode:     unifieddatalibrary.String("MAJ"),
					RoleTypeCode: unifieddatalibrary.String("TC"),
					SkillLevel:   unifieddatalibrary.Int(3),
					Specialty:    unifieddatalibrary.String("ELEN"),
				}},
				Quantity:                unifieddatalibrary.Int(1),
				ReadyTime:               unifieddatalibrary.Time(time.Now()),
				ReceivedTime:            unifieddatalibrary.Time(time.Now()),
				RecoveryRequestTypeCode: unifieddatalibrary.String("E"),
				RedeployPlanNumber:      unifieddatalibrary.String("T89003"),
				RedeployShipmentUnitID:  unifieddatalibrary.String("X400LA31949108"),
				RequestNumber:           unifieddatalibrary.String("89208"),
				ResupportFlag:           unifieddatalibrary.Bool(true),
				ShipmentUnitID:          unifieddatalibrary.String("FB44273196X501XXX"),
				SiPoc:                   unifieddatalibrary.String("SMITH, JOHN J"),
				SourceIcao:              unifieddatalibrary.String("PHIK"),
			}},
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsTransportationPlan{{
				ActDepTime:             unifieddatalibrary.Time(time.Now()),
				AircraftStatus:         unifieddatalibrary.String("NMCMU"),
				ApproxArrTime:          unifieddatalibrary.Time(time.Now()),
				CancelledDate:          unifieddatalibrary.Time(time.Now()),
				ClosedDate:             unifieddatalibrary.Time(time.Now()),
				Coordinator:            unifieddatalibrary.String("SMITH, JOHN"),
				CoordinatorUnit:        unifieddatalibrary.String("TACC"),
				DestinationIcao:        unifieddatalibrary.String("YBCS"),
				Duration:               unifieddatalibrary.String("086:20"),
				EstArrTime:             unifieddatalibrary.Time(time.Now()),
				EstDepTime:             unifieddatalibrary.Time(time.Now()),
				LastChangedDate:        unifieddatalibrary.Time(time.Now()),
				LogisticMasterRecordID: unifieddatalibrary.String("L62126"),
				LogisticsSegments: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment{{
					ArrivalIcao:    unifieddatalibrary.String("YBCS"),
					DepartureIcao:  unifieddatalibrary.String("PHIK"),
					ExtMissionID:   unifieddatalibrary.String("2001101RF01202307062205"),
					IDMission:      unifieddatalibrary.String("EXAMPLE-UUID"),
					Itin:           unifieddatalibrary.Int(200),
					MissionNumber:  unifieddatalibrary.String("TAM308901196"),
					MissionType:    unifieddatalibrary.String("SAAM"),
					ModeCode:       unifieddatalibrary.String("A"),
					SegActArrTime:  unifieddatalibrary.Time(time.Now()),
					SegActDepTime:  unifieddatalibrary.Time(time.Now()),
					SegAircraftMds: unifieddatalibrary.String("B7772E"),
					SegEstArrTime:  unifieddatalibrary.Time(time.Now()),
					SegEstDepTime:  unifieddatalibrary.Time(time.Now()),
					SegmentNumber:  unifieddatalibrary.Int(3),
					SegTailNumber:  unifieddatalibrary.String("N819AX"),
				}},
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticssupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				Majcom:          unifieddatalibrary.String("HQAMC"),
				MissionChange:   unifieddatalibrary.Bool(false),
				NumEnrouteStops: unifieddatalibrary.Int(4),
				NumTransLoads:   unifieddatalibrary.Int(3),
				OriginIcao:      unifieddatalibrary.String("KATL"),
				PlanDefinition:  unifieddatalibrary.String("DEPLOY"),
				PlansNumber:     unifieddatalibrary.String("T89002"),
				SerialNumber:    unifieddatalibrary.String("9009209"),
				StatusCode:      unifieddatalibrary.String("N"),
				TpAircraftMds:   unifieddatalibrary.String("C17A"),
				TpTailNumber:    unifieddatalibrary.String("99209"),
			}},
			MaintStatusCode: unifieddatalibrary.String("NMCMU"),
			McTime:          unifieddatalibrary.Time(time.Now()),
			MeTime:          unifieddatalibrary.Time(time.Now()),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Owner:           unifieddatalibrary.String("EXAMPLE_OWNER"),
			ReopenFlag:      unifieddatalibrary.Bool(true),
			RptClosedTime:   unifieddatalibrary.Time(time.Now()),
			SuppIcao:        unifieddatalibrary.String("KCOS"),
			TailNumber:      unifieddatalibrary.String("99290"),
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

func TestLogisticssupportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logisticssupport.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LogisticssupportGetParams{
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

func TestLogisticssupportQueryhelp(t *testing.T) {
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
	err := client.Logisticssupport.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogisticssupportTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Logisticssupport.Tuple(context.TODO(), unifieddatalibrary.LogisticssupportTupleParams{
		Columns:     "columns",
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

func TestLogisticssupportUnvalidatedPublish(t *testing.T) {
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
	err := client.Logisticssupport.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LogisticssupportUnvalidatedPublishParams{
		Body: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			RptCreatedTime:        time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
			AircraftMds:           unifieddatalibrary.String("CO17A"),
			CurrIcao:              unifieddatalibrary.String("KCOS"),
			Etic:                  unifieddatalibrary.Time(time.Now()),
			Etmc:                  unifieddatalibrary.Time(time.Now()),
			ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
			LogisticAction:        unifieddatalibrary.String("WA"),
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock{{
						Quantity:       unifieddatalibrary.Int(4),
						SourceIcao:     unifieddatalibrary.String("PHIK"),
						StockCheckTime: unifieddatalibrary.Time(time.Now()),
						StockPoc:       unifieddatalibrary.String("SMITH, JOHN J"),
					}},
					MeasurementUnitCode:  unifieddatalibrary.String("EA"),
					NationalStockNumber:  unifieddatalibrary.String("5310-00-045-3299"),
					PartNumber:           unifieddatalibrary.String("MS35338-42"),
					RequestVerifier:      unifieddatalibrary.String("JOHN SMITH"),
					SupplyDocumentNumber: unifieddatalibrary.String("J223FU31908300"),
					TechnicalOrderText:   unifieddatalibrary.String("1C-17A-4"),
					WorkUnitCode:         unifieddatalibrary.String("5611UU001"),
				}},
				LogisticsRemarks: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty{{
					FirstName:    unifieddatalibrary.String("JOHN"),
					Last4Ssn:     unifieddatalibrary.String("9999"),
					LastName:     unifieddatalibrary.String("SMITH"),
					RankCode:     unifieddatalibrary.String("MAJ"),
					RoleTypeCode: unifieddatalibrary.String("TC"),
					SkillLevel:   unifieddatalibrary.Int(3),
					Specialty:    unifieddatalibrary.String("ELEN"),
				}},
				Quantity:                unifieddatalibrary.Int(1),
				ReadyTime:               unifieddatalibrary.Time(time.Now()),
				ReceivedTime:            unifieddatalibrary.Time(time.Now()),
				RecoveryRequestTypeCode: unifieddatalibrary.String("E"),
				RedeployPlanNumber:      unifieddatalibrary.String("T89003"),
				RedeployShipmentUnitID:  unifieddatalibrary.String("X400LA31949108"),
				RequestNumber:           unifieddatalibrary.String("89208"),
				ResupportFlag:           unifieddatalibrary.Bool(true),
				ShipmentUnitID:          unifieddatalibrary.String("FB44273196X501XXX"),
				SiPoc:                   unifieddatalibrary.String("SMITH, JOHN J"),
				SourceIcao:              unifieddatalibrary.String("PHIK"),
			}},
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan{{
				ActDepTime:             unifieddatalibrary.Time(time.Now()),
				AircraftStatus:         unifieddatalibrary.String("NMCMU"),
				ApproxArrTime:          unifieddatalibrary.Time(time.Now()),
				CancelledDate:          unifieddatalibrary.Time(time.Now()),
				ClosedDate:             unifieddatalibrary.Time(time.Now()),
				Coordinator:            unifieddatalibrary.String("SMITH, JOHN"),
				CoordinatorUnit:        unifieddatalibrary.String("TACC"),
				DestinationIcao:        unifieddatalibrary.String("YBCS"),
				Duration:               unifieddatalibrary.String("086:20"),
				EstArrTime:             unifieddatalibrary.Time(time.Now()),
				EstDepTime:             unifieddatalibrary.Time(time.Now()),
				LastChangedDate:        unifieddatalibrary.Time(time.Now()),
				LogisticMasterRecordID: unifieddatalibrary.String("L62126"),
				LogisticsSegments: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment{{
					ArrivalIcao:    unifieddatalibrary.String("YBCS"),
					DepartureIcao:  unifieddatalibrary.String("PHIK"),
					ExtMissionID:   unifieddatalibrary.String("2001101RF01202307062205"),
					IDMission:      unifieddatalibrary.String("EXAMPLE-UUID"),
					Itin:           unifieddatalibrary.Int(200),
					MissionNumber:  unifieddatalibrary.String("TAM308901196"),
					MissionType:    unifieddatalibrary.String("SAAM"),
					ModeCode:       unifieddatalibrary.String("A"),
					SegActArrTime:  unifieddatalibrary.Time(time.Now()),
					SegActDepTime:  unifieddatalibrary.Time(time.Now()),
					SegAircraftMds: unifieddatalibrary.String("B7772E"),
					SegEstArrTime:  unifieddatalibrary.Time(time.Now()),
					SegEstDepTime:  unifieddatalibrary.Time(time.Now()),
					SegmentNumber:  unifieddatalibrary.Int(3),
					SegTailNumber:  unifieddatalibrary.String("N819AX"),
				}},
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticssupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				Majcom:          unifieddatalibrary.String("HQAMC"),
				MissionChange:   unifieddatalibrary.Bool(false),
				NumEnrouteStops: unifieddatalibrary.Int(4),
				NumTransLoads:   unifieddatalibrary.Int(3),
				OriginIcao:      unifieddatalibrary.String("KATL"),
				PlanDefinition:  unifieddatalibrary.String("DEPLOY"),
				PlansNumber:     unifieddatalibrary.String("T89002"),
				SerialNumber:    unifieddatalibrary.String("9009209"),
				StatusCode:      unifieddatalibrary.String("N"),
				TpAircraftMds:   unifieddatalibrary.String("C17A"),
				TpTailNumber:    unifieddatalibrary.String("99209"),
			}},
			MaintStatusCode: unifieddatalibrary.String("NMCMU"),
			McTime:          unifieddatalibrary.Time(time.Now()),
			MeTime:          unifieddatalibrary.Time(time.Now()),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Owner:           unifieddatalibrary.String("EXAMPLE_OWNER"),
			ReopenFlag:      unifieddatalibrary.Bool(true),
			RptClosedTime:   unifieddatalibrary.Time(time.Now()),
			SuppIcao:        unifieddatalibrary.String("KCOS"),
			TailNumber:      unifieddatalibrary.String("99290"),
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
