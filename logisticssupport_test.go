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

func TestLogisticsSupportNewWithOptionalParams(t *testing.T) {
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
	err := client.LogisticsSupport.New(context.TODO(), unifieddatalibrary.LogisticsSupportNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LogisticsSupportNewParamsDataModeTest,
		RptCreatedTime:        time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
		AircraftMds:           unifieddatalibrary.String("CO17A"),
		CurrIcao:              unifieddatalibrary.String("KCOS"),
		Etic:                  unifieddatalibrary.Time(time.Now()),
		Etmc:                  unifieddatalibrary.Time(time.Now()),
		ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
		LogisticAction:        unifieddatalibrary.String("WA"),
		LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsDiscrepancyInfo{{
			ClosureTime:     unifieddatalibrary.Time(time.Now()),
			DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
			Jcn:             unifieddatalibrary.String("231942400"),
			JobStTime:       unifieddatalibrary.Time(time.Now()),
		}},
		LogisticsRecordID: unifieddatalibrary.String("L62017"),
		LogisticsRemarks: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsRemark{{
			LastChanged: unifieddatalibrary.Time(time.Now()),
			Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
			Username:    unifieddatalibrary.String("JSMITH"),
		}},
		LogisticsSupportItems: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsSupportItem{{
			Cannibalized:        unifieddatalibrary.Bool(true),
			DeployPlanNumber:    unifieddatalibrary.String("T89003"),
			Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
			ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
			JobControlNumber:    unifieddatalibrary.String("231942400"),
			LogisticsParts: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsSupportItemLogisticsPart{{
				FigureNumber:     unifieddatalibrary.String("3"),
				IndexNumber:      unifieddatalibrary.String("4"),
				LocationVerifier: unifieddatalibrary.String("JANE DOE"),
				LogisticsStocks: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsSupportItemLogisticsPartLogisticsStock{{
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
			LogisticsRemarks: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsSupportItemLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSpecialties: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsSupportItemLogisticsSpecialty{{
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
		LogisticsTransportationPlans: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsTransportationPlan{{
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
			LogisticsSegments: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsSegment{{
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
			LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticsSupportNewParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
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

func TestLogisticsSupportUpdateWithOptionalParams(t *testing.T) {
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
	err := client.LogisticsSupport.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LogisticsSupportUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LogisticsSupportUpdateParamsDataModeTest,
			RptCreatedTime:        time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LOGISTICS-SUPPORT-DETAILS UUID"),
			AircraftMds:           unifieddatalibrary.String("CO17A"),
			CurrIcao:              unifieddatalibrary.String("KCOS"),
			Etic:                  unifieddatalibrary.Time(time.Now()),
			Etmc:                  unifieddatalibrary.Time(time.Now()),
			ExtSystemID:           unifieddatalibrary.String("GDSSBL012307131347070165"),
			LogisticAction:        unifieddatalibrary.String("WA"),
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsPartLogisticsStock{{
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
				LogisticsRemarks: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsSupportItemLogisticsSpecialty{{
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
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsTransportationPlan{{
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
				LogisticsSegments: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsSegment{{
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
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticsSupportUpdateParamsLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
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

func TestLogisticsSupportListWithOptionalParams(t *testing.T) {
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
	_, err := client.LogisticsSupport.List(context.TODO(), unifieddatalibrary.LogisticsSupportListParams{
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

func TestLogisticsSupportCountWithOptionalParams(t *testing.T) {
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
	_, err := client.LogisticsSupport.Count(context.TODO(), unifieddatalibrary.LogisticsSupportCountParams{
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

func TestLogisticsSupportNewBulk(t *testing.T) {
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
	err := client.LogisticsSupport.NewBulk(context.TODO(), unifieddatalibrary.LogisticsSupportNewBulkParams{
		Body: []unifieddatalibrary.LogisticsSupportNewBulkParamsBody{{
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
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock{{
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
				LogisticsRemarks: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsSupportItemLogisticsSpecialty{{
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
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlan{{
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
				LogisticsSegments: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsSegment{{
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
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticsSupportNewBulkParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
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

func TestLogisticsSupportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LogisticsSupport.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LogisticsSupportGetParams{
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

func TestLogisticsSupportQueryhelp(t *testing.T) {
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
	_, err := client.LogisticsSupport.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogisticsSupportTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.LogisticsSupport.Tuple(context.TODO(), unifieddatalibrary.LogisticsSupportTupleParams{
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

func TestLogisticsSupportUnvalidatedPublish(t *testing.T) {
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
	err := client.LogisticsSupport.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LogisticsSupportUnvalidatedPublishParams{
		Body: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBody{{
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
			LogisticsDiscrepancyInfos: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsDiscrepancyInfo{{
				ClosureTime:     unifieddatalibrary.Time(time.Now()),
				DiscrepancyInfo: unifieddatalibrary.String("PILOT WINDSHIELD PANEL ASSY CRACKED, AND ARCING REQ R2 IAW 56.11.10"),
				Jcn:             unifieddatalibrary.String("231942400"),
				JobStTime:       unifieddatalibrary.Time(time.Now()),
			}},
			LogisticsRecordID: unifieddatalibrary.String("L62017"),
			LogisticsRemarks: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsRemark{{
				LastChanged: unifieddatalibrary.Time(time.Now()),
				Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
				Username:    unifieddatalibrary.String("JSMITH"),
			}},
			LogisticsSupportItems: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItem{{
				Cannibalized:        unifieddatalibrary.Bool(true),
				DeployPlanNumber:    unifieddatalibrary.String("T89003"),
				Description:         unifieddatalibrary.String("HOIST ADAPTER KIT"),
				ItemLastChangedDate: unifieddatalibrary.Time(time.Now()),
				JobControlNumber:    unifieddatalibrary.String("231942400"),
				LogisticsParts: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPart{{
					FigureNumber:     unifieddatalibrary.String("3"),
					IndexNumber:      unifieddatalibrary.String("4"),
					LocationVerifier: unifieddatalibrary.String("JANE DOE"),
					LogisticsStocks: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsPartLogisticsStock{{
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
				LogisticsRemarks: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsRemark{{
					LastChanged: unifieddatalibrary.Time(time.Now()),
					Remark:      unifieddatalibrary.String("EXAMPLE REMARK"),
					Username:    unifieddatalibrary.String("JSMITH"),
				}},
				LogisticsSpecialties: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsSupportItemLogisticsSpecialty{{
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
			LogisticsTransportationPlans: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlan{{
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
				LogisticsSegments: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsSegment{{
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
				LogisticsTransportationPlansRemarks: []unifieddatalibrary.LogisticsSupportUnvalidatedPublishParamsBodyLogisticsTransportationPlanLogisticsTransportationPlansRemark{{
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
