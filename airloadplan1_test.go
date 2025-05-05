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

func TestAirloadPlanUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirloadPlans.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirloadPlanUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirloadPlanUpdateParamsDataModeTest,
			EstDepTime:            time.Now(),
			Source:                "source",
			ID:                    unifieddatalibrary.String("0457f578-e29c-312e-85aa-0a04a430bdd0"),
			ACLOnboard:            unifieddatalibrary.Float(500.1),
			ACLReleased:           unifieddatalibrary.Float(200.1),
			AircraftMds:           unifieddatalibrary.String("C17A"),
			AirLoadPlanHazmatActuals: []unifieddatalibrary.AirloadPlanUpdateParamsAirLoadPlanHazmatActual{{
				Ashc:           unifieddatalibrary.String("RFL"),
				Cgc:            unifieddatalibrary.String("A"),
				ClassDiv:       unifieddatalibrary.String("1.1"),
				HazDescription: unifieddatalibrary.String("CORROSIVE OXIDIZER"),
				HazmatRemarks:  unifieddatalibrary.String("Hazmat remarks"),
				HazNum:         unifieddatalibrary.String("2031"),
				HazNumType:     unifieddatalibrary.String("UN"),
				HazOffIcao:     unifieddatalibrary.String("MBPV"),
				HazOffItin:     unifieddatalibrary.Int(300),
				HazOnIcao:      unifieddatalibrary.String("LIRQ"),
				HazOnItin:      unifieddatalibrary.Int(50),
				HazPieces:      unifieddatalibrary.Int(29),
				HazTcn:         unifieddatalibrary.String("M1358232245912XXX"),
				HazWeight:      unifieddatalibrary.Float(22.1),
				ItemName:       unifieddatalibrary.String("NITRIC ACID"),
				LotNum:         unifieddatalibrary.String("1234A"),
				NetExpWt:       unifieddatalibrary.Float(12.1),
			}},
			AirLoadPlanHr: []unifieddatalibrary.AirloadPlanUpdateParamsAirLoadPlanHr{{
				Container:    unifieddatalibrary.String("Metal"),
				Escort:       unifieddatalibrary.String("Jane Doe"),
				HrEstArrTime: unifieddatalibrary.Time(time.Now()),
				HrOffIcao:    unifieddatalibrary.String("KDEN"),
				HrOffItin:    unifieddatalibrary.Int(200),
				HrOnIcao:     unifieddatalibrary.String("KCOS"),
				HrOnItin:     unifieddatalibrary.Int(100),
				HrRemarks:    unifieddatalibrary.String("HR remarks"),
				Name:         unifieddatalibrary.String("John Doe"),
				Rank:         unifieddatalibrary.String("Captain"),
				RecAgency:    unifieddatalibrary.String("Agency name"),
				Service:      unifieddatalibrary.String("Air Force"),
				Viewable:     unifieddatalibrary.Bool(true),
			}},
			AirLoadPlanPalletDetails: []unifieddatalibrary.AirloadPlanUpdateParamsAirLoadPlanPalletDetail{{
				Category:        unifieddatalibrary.String("AMCMICAP"),
				Pp:              unifieddatalibrary.String("2"),
				PpDescription:   unifieddatalibrary.String("Ammunition"),
				PpOffIcao:       unifieddatalibrary.String("MBPV"),
				PpPieces:        unifieddatalibrary.Int(3),
				PpRemarks:       unifieddatalibrary.String("Pallet remarks"),
				PpTcn:           unifieddatalibrary.String("M1358232245912XXX"),
				PpWeight:        unifieddatalibrary.Float(100.1),
				SpecialInterest: unifieddatalibrary.Bool(true),
			}},
			AirLoadPlanPaxCargo: []unifieddatalibrary.AirloadPlanUpdateParamsAirLoadPlanPaxCargo{{
				AmbPax:           unifieddatalibrary.Int(5),
				AttPax:           unifieddatalibrary.Int(6),
				AvailablePax:     unifieddatalibrary.Int(20),
				BagWeight:        unifieddatalibrary.Float(2000.1),
				CivPax:           unifieddatalibrary.Int(3),
				DvPax:            unifieddatalibrary.Int(2),
				FnPax:            unifieddatalibrary.Int(1),
				GroupCargoWeight: unifieddatalibrary.Float(5000.1),
				GroupType:        unifieddatalibrary.String("OFFTHIS"),
				LitPax:           unifieddatalibrary.Int(4),
				MailWeight:       unifieddatalibrary.Float(200.1),
				NumPallet:        unifieddatalibrary.Int(20),
				PalletWeight:     unifieddatalibrary.Float(400.1),
				PaxWeight:        unifieddatalibrary.Float(8000.1),
				RequiredPax:      unifieddatalibrary.Int(20),
			}},
			AirLoadPlanUlnActuals: []unifieddatalibrary.AirloadPlanUpdateParamsAirLoadPlanUlnActual{{
				NumAmbulatory:  unifieddatalibrary.Int(10),
				NumAttendant:   unifieddatalibrary.Int(10),
				NumLitter:      unifieddatalibrary.Int(10),
				NumPax:         unifieddatalibrary.Int(44),
				OffloadID:      unifieddatalibrary.Int(300),
				OffloadLoCode:  unifieddatalibrary.String("KHOP"),
				OnloadID:       unifieddatalibrary.Int(200),
				OnloadLoCode:   unifieddatalibrary.String("KCHS"),
				Oplan:          unifieddatalibrary.String("5027A"),
				ProjName:       unifieddatalibrary.String("CENTINTRA21"),
				Uln:            unifieddatalibrary.String("T01ME01"),
				UlnCargoWeight: unifieddatalibrary.Float(1000.1),
				UlnRemarks:     unifieddatalibrary.String("ULN actuals remark"),
			}},
			ArrAirfield:          unifieddatalibrary.String("W99"),
			ArrIcao:              unifieddatalibrary.String("ETAR"),
			AvailableTime:        unifieddatalibrary.Time(time.Now()),
			BasicMoment:          unifieddatalibrary.Float(2500.1),
			BasicWeight:          unifieddatalibrary.Float(100.1),
			BriefTime:            unifieddatalibrary.Time(time.Now()),
			CallSign:             unifieddatalibrary.String("RCH1234"),
			CargoBayFsMax:        unifieddatalibrary.Float(20.1),
			CargoBayFsMin:        unifieddatalibrary.Float(10.1),
			CargoBayWidth:        unifieddatalibrary.Float(3.1),
			CargoConfig:          unifieddatalibrary.String("C-1"),
			CargoMoment:          unifieddatalibrary.Float(2500.1),
			CargoVolume:          unifieddatalibrary.Float(50.1),
			CargoWeight:          unifieddatalibrary.Float(100.1),
			CrewSize:             unifieddatalibrary.Int(5),
			DepAirfield:          unifieddatalibrary.String("W99"),
			DepIcao:              unifieddatalibrary.String("KCHS"),
			EquipConfig:          unifieddatalibrary.String("Standard"),
			EstArrTime:           unifieddatalibrary.Time(time.Now()),
			EstLandingFuelMoment: unifieddatalibrary.Float(2500.1),
			EstLandingFuelWeight: unifieddatalibrary.Float(100.1),
			ExternalID:           unifieddatalibrary.String("dec7a61a-cd97-4af0-b7bc-f4c3bb33341b"),
			FuelMoment:           unifieddatalibrary.Float(2500.1),
			FuelWeight:           unifieddatalibrary.Float(100.1),
			GrossCg:              unifieddatalibrary.Float(38.8),
			GrossMoment:          unifieddatalibrary.Float(2500.1),
			GrossWeight:          unifieddatalibrary.Float(100.1),
			IDMission:            unifieddatalibrary.String("412bebb6-a45e-029c-ca51-e29f8a442b12"),
			IDSortie:             unifieddatalibrary.String("823acfbe6-f36a-157b-ef32-b47c9b589c4"),
			LandingCg:            unifieddatalibrary.Float(38.2),
			LandingMoment:        unifieddatalibrary.Float(2500.1),
			LandingWeight:        unifieddatalibrary.Float(100.1),
			LegNum:               unifieddatalibrary.Int(200),
			LoadmasterName:       unifieddatalibrary.String("John Smith"),
			LoadmasterRank:       unifieddatalibrary.String("Staff Sergeant"),
			LoadRemarks:          unifieddatalibrary.String("Load remarks"),
			MissionNumber:        unifieddatalibrary.String("AJM123456123"),
			OperatingMoment:      unifieddatalibrary.Float(2500.1),
			OperatingWeight:      unifieddatalibrary.Float(100.1),
			Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PpOnboard:            unifieddatalibrary.Int(18),
			PpReleased:           unifieddatalibrary.Int(5),
			SchedTime:            unifieddatalibrary.Time(time.Now()),
			SeatsOnboard:         unifieddatalibrary.Int(20),
			SeatsReleased:        unifieddatalibrary.Int(15),
			TailNumber:           unifieddatalibrary.String("77187"),
			TankConfig:           unifieddatalibrary.String("ER"),
			UtilCode:             unifieddatalibrary.String("AD"),
			ZeroFuelCg:           unifieddatalibrary.Float(39.5),
			ZeroFuelMoment:       unifieddatalibrary.Float(2500.1),
			ZeroFuelWeight:       unifieddatalibrary.Float(100.1),
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

func TestAirloadPlanDelete(t *testing.T) {
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
	err := client.AirloadPlans.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
