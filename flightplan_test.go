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

func TestFlightplanNewWithOptionalParams(t *testing.T) {
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
	err := client.Flightplan.New(context.TODO(), unifieddatalibrary.FlightplanNewParams{
		ArrAirfield:           "KCHS",
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.FlightplanNewParamsDataModeTest,
		DepAirfield:           "KSLV",
		GenTs:                 time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("c44b0a80-9fef-63d9-6267-79037fb93e4c"),
		AircraftMds:           unifieddatalibrary.String("KC-130 HERCULES"),
		AirRefuelEvents: []unifieddatalibrary.FlightplanNewParamsAirRefuelEvent{{
			ArDegrade:       unifieddatalibrary.Float(3.1),
			ArExchangedFuel: unifieddatalibrary.Float(1500.1),
			ArNum:           unifieddatalibrary.Int(2),
			DivertFuel:      unifieddatalibrary.Float(143000.1),
			ExitFuel:        unifieddatalibrary.Float(160000.1),
		}},
		AmcMissionID:      unifieddatalibrary.String("AJM7939B1123"),
		AppLandingFuel:    unifieddatalibrary.Float(3000.1),
		ArrAlternate1:     unifieddatalibrary.String("EDDS"),
		ArrAlternate1Fuel: unifieddatalibrary.Float(6000.1),
		ArrAlternate2:     unifieddatalibrary.String("EDDM"),
		ArrAlternate2Fuel: unifieddatalibrary.Float(6000.1),
		ArrIceFuel:        unifieddatalibrary.Float(1000.1),
		ArrRunway:         unifieddatalibrary.String("05L"),
		AtcAddresses:      []string{"EYCBZMFO", "EUCHZMFP", "ETARYXYX", "EDUUZVZI"},
		AvgTempDev:        unifieddatalibrary.Float(16.1),
		BurnedFuel:        unifieddatalibrary.Float(145000.1),
		CallSign:          unifieddatalibrary.String("HKY629"),
		CargoRemark:       unifieddatalibrary.String("Expecting 55,000 lbs. If different, call us."),
		ClimbFuel:         unifieddatalibrary.Float(7000.1),
		ClimbTime:         unifieddatalibrary.String("00:13"),
		ContingencyFuel:   unifieddatalibrary.Float(3000.1),
		CountryCodes:      []string{"US", "CA", "UK"},
		DepAlternate:      unifieddatalibrary.String("LFPO"),
		DepressFuel:       unifieddatalibrary.Float(20000.1),
		DepRunway:         unifieddatalibrary.String("05L"),
		DragIndex:         unifieddatalibrary.Float(16.9),
		EarlyDescentFuel:  unifieddatalibrary.Float(500.1),
		EnduranceTime:     unifieddatalibrary.String("08:45"),
		EnrouteFuel:       unifieddatalibrary.Float(155000.1),
		EnrouteTime:       unifieddatalibrary.String("06:30"),
		Equipment:         unifieddatalibrary.String("SDFGHIRTUWXYZ/H"),
		EstDepTime:        unifieddatalibrary.Time(time.Now()),
		EtopsAirfields:    []string{"KHSV", "KISP", "KBG", "LTBS"},
		EtopsAltAirfields: []string{"KHSV", "KISP", "KBG", "LTBS"},
		EtopsRating:       unifieddatalibrary.String("85 MINUTES"),
		EtopsValWindow:    unifieddatalibrary.String("LPLA: 0317Z-0722Z"),
		ExternalID:        unifieddatalibrary.String("AFMAPP20322347140001"),
		FlightPlanMessages: []unifieddatalibrary.FlightplanNewParamsFlightPlanMessage{{
			MsgText:   unifieddatalibrary.String("Message text"),
			RoutePath: unifieddatalibrary.String("PRIMARY"),
			Severity:  unifieddatalibrary.String("SEVERE"),
			WpNum:     unifieddatalibrary.String("20"),
		}},
		FlightPlanPointGroups: []unifieddatalibrary.FlightplanNewParamsFlightPlanPointGroup{{
			AvgFuelFlow:        unifieddatalibrary.Float(19693.1),
			EtopsAvgWindFactor: unifieddatalibrary.Float(10.1),
			EtopsDistance:      unifieddatalibrary.Float(684.1),
			EtopsReqFuel:       unifieddatalibrary.Float(4412.1),
			EtopsTempDev:       unifieddatalibrary.Float(9.1),
			EtopsTime:          unifieddatalibrary.String("01:23"),
			FlightPlanPoints: []unifieddatalibrary.FlightplanNewParamsFlightPlanPointGroupFlightPlanPoint{{
				FppEta:     unifieddatalibrary.Time(time.Now()),
				FppLat:     unifieddatalibrary.Float(45.23),
				FppLon:     unifieddatalibrary.Float(179.1),
				FppReqFuel: unifieddatalibrary.Float(4250.1),
				PointName:  unifieddatalibrary.String("CRUISE ALTITUDE ETP"),
			}},
			FromTakeoffTime:   unifieddatalibrary.String("07:29"),
			FsafAvgWindFactor: unifieddatalibrary.Float(10.1),
			FsafDistance:      unifieddatalibrary.Float(684.1),
			FsafReqFuel:       unifieddatalibrary.Float(50380.1),
			FsafTempDev:       unifieddatalibrary.Float(9.1),
			FsafTime:          unifieddatalibrary.String("01:23"),
			FuelCalcAlt:       unifieddatalibrary.Float(100.1),
			FuelCalcSpd:       unifieddatalibrary.Float(365.1),
			LsafAvgWindFactor: unifieddatalibrary.Float(13.1),
			LsafDistance:      unifieddatalibrary.Float(684.1),
			LsafName:          unifieddatalibrary.String("LPPD"),
			LsafReqFuel:       unifieddatalibrary.Float(50787.1),
			LsafTempDev:       unifieddatalibrary.Float(9.1),
			LsafTime:          unifieddatalibrary.String("01:23"),
			PlannedFuel:       unifieddatalibrary.Float(190319.1),
			PointGroupName:    unifieddatalibrary.String("ETOPS_CF_POINT_1"),
			WorstFuelCase:     unifieddatalibrary.String("DEPRESSURIZED ENGINE OUT ETP"),
		}},
		FlightPlanWaypoints: []unifieddatalibrary.FlightplanNewParamsFlightPlanWaypoint{{
			Type:               "COMMENT",
			WaypointName:       "KCHS",
			AaTacanChannel:     unifieddatalibrary.String("31/94"),
			AirDistance:        unifieddatalibrary.Float(321.1),
			Airway:             unifieddatalibrary.String("W15"),
			Alt:                unifieddatalibrary.Float(27000.1),
			ArID:               unifieddatalibrary.String("AR202"),
			Arpt:               unifieddatalibrary.String("ARIP"),
			Ata:                unifieddatalibrary.Time(time.Now()),
			AvgCalAirspeed:     unifieddatalibrary.Float(200.1),
			AvgDriftAng:        unifieddatalibrary.Float(-3.2),
			AvgGroundSpeed:     unifieddatalibrary.Float(300.1),
			AvgTrueAirspeed:    unifieddatalibrary.Float(210.1),
			AvgWindDir:         unifieddatalibrary.Float(165.5),
			AvgWindSpeed:       unifieddatalibrary.Float(14.4),
			DayLowAlt:          unifieddatalibrary.Float(1500.1),
			Eta:                unifieddatalibrary.Time(time.Now()),
			ExchangedFuel:      unifieddatalibrary.Float(-30400.1),
			FuelFlow:           unifieddatalibrary.Float(17654.1),
			IceCat:             unifieddatalibrary.String("MODERATE"),
			Lat:                unifieddatalibrary.Float(45.23),
			LegAlternate:       unifieddatalibrary.String("KCHS"),
			LegDragIndex:       unifieddatalibrary.Float(1.2),
			LegFuelDegrade:     unifieddatalibrary.Float(10.1),
			LegMach:            unifieddatalibrary.Float(0.74),
			LegMsnIndex:        unifieddatalibrary.Float(65),
			LegWindFac:         unifieddatalibrary.Float(-32.1),
			Lon:                unifieddatalibrary.Float(179.1),
			MagCourse:          unifieddatalibrary.Float(338.1),
			MagHeading:         unifieddatalibrary.Float(212.1),
			MagVar:             unifieddatalibrary.Float(-13.2),
			Navaid:             unifieddatalibrary.String("HTO"),
			NightLowAlt:        unifieddatalibrary.Float(2300.1),
			NvgLowAlt:          unifieddatalibrary.Float(2450.1),
			PointWindDir:       unifieddatalibrary.Float(165.5),
			PointWindSpeed:     unifieddatalibrary.Float(14.4),
			PriFreq:            unifieddatalibrary.Float(357.5),
			SecFreq:            unifieddatalibrary.Float(357.5),
			TacanChannel:       unifieddatalibrary.String("83X"),
			TempDev:            unifieddatalibrary.Float(12.1),
			ThunderCat:         unifieddatalibrary.String("MODERATE"),
			TotalAirDistance:   unifieddatalibrary.Float(3251.1),
			TotalFlownDistance: unifieddatalibrary.Float(688.1),
			TotalRemDistance:   unifieddatalibrary.Float(1288.1),
			TotalRemFuel:       unifieddatalibrary.Float(30453.1),
			TotalTime:          unifieddatalibrary.String("08:45"),
			TotalTimeRem:       unifieddatalibrary.String("01:43"),
			TotalUsedFuel:      unifieddatalibrary.Float(70431.1),
			TotalWeight:        unifieddatalibrary.Float(207123.1),
			TrueCourse:         unifieddatalibrary.Float(328.1),
			TurbCat:            unifieddatalibrary.String("EXTREME"),
			VorFreq:            unifieddatalibrary.Float(113.6),
			WaypointNum:        unifieddatalibrary.Int(20),
			ZoneDistance:       unifieddatalibrary.Float(212.1),
			ZoneFuel:           unifieddatalibrary.Float(1120.1),
			ZoneTime:           unifieddatalibrary.Float(36.1),
		}},
		FlightRules:        unifieddatalibrary.String("l"),
		FlightType:         unifieddatalibrary.String("MILITARY"),
		FuelDegrade:        unifieddatalibrary.Float(10.3),
		GpsRaim:            unifieddatalibrary.String("Failed by FAA SAPT 184022AUG2022"),
		HoldDownFuel:       unifieddatalibrary.Float(500.1),
		HoldFuel:           unifieddatalibrary.Float(6000.1),
		HoldTime:           unifieddatalibrary.String("01:00"),
		IDAircraft:         unifieddatalibrary.String("4f4a67c6-40fd-11ee-be56-0242ac120002"),
		IDArrAirfield:      unifieddatalibrary.String("363080c2-40fd-11ee-be56-0242ac120002"),
		IDDepAirfield:      unifieddatalibrary.String("2a9020f6-40fd-11ee-be56-0242ac120002"),
		IdentExtraFuel:     unifieddatalibrary.Float(5000.1),
		IDSortie:           unifieddatalibrary.String("9d60c1b1-10b1-b2a7-e403-84c5d7eeb170"),
		InitialCruiseSpeed: unifieddatalibrary.String("N0305"),
		InitialFlightLevel: unifieddatalibrary.String("F270"),
		LandingFuel:        unifieddatalibrary.Float(19000.1),
		LegNum:             unifieddatalibrary.Int(100),
		MinDivertFuel:      unifieddatalibrary.Float(25000.1),
		MsnIndex:           unifieddatalibrary.Float(44.1),
		Notes:              unifieddatalibrary.String("STS/STATE PBN/A1B2B5C2C4D2D4 EUR/PROTECTED"),
		NumAircraft:        unifieddatalibrary.Int(1),
		OpConditionFuel:    unifieddatalibrary.Float(5000.1),
		OpWeight:           unifieddatalibrary.Float(251830.5),
		Origin:             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Originator:         unifieddatalibrary.String("ETARYXYX"),
		PlannerRemark:      unifieddatalibrary.String("Flight plan is good for 2 days before airspace closes over the UK."),
		RampFuel:           unifieddatalibrary.Float(180000.1),
		RemAlternate1Fuel:  unifieddatalibrary.Float(18000.1),
		RemAlternate2Fuel:  unifieddatalibrary.Float(18000.1),
		ReserveFuel:        unifieddatalibrary.Float(10000.1),
		RouteString:        unifieddatalibrary.String("RENV3B RENVI Y86 GOSVA/N0317F260 DCT EVLIT DCT UMUGI DCT NISIX DCT GIGOD DCT DIPEB DCT\nGORPI Z80 TILAV L87 RAKIT Z717 PODUS Z130 MAG/N0298F220 Z20 KENIG/N0319F220 Z20 ORTAG T177\nESEGU Z20 BEBLA DCT MASEK/N0300F200 DCT GISEM/N0319F200 DCT BOMBI/N0276F060 DCT RIDSU DCT"),
		Sid:                unifieddatalibrary.String("RENV3B"),
		Star:               unifieddatalibrary.String("ADANA"),
		Status:             unifieddatalibrary.String("APPROVED"),
		TailNumber:         unifieddatalibrary.String("77187"),
		TakeoffFuel:        unifieddatalibrary.Float(178500.1),
		TaxiFuel:           unifieddatalibrary.Float(1500.1),
		ThunderAvoidFuel:   unifieddatalibrary.Float(1000.1),
		TocFuel:            unifieddatalibrary.Float(160000.1),
		TocIceFuel:         unifieddatalibrary.Float(1000.1),
		TodFuel:            unifieddatalibrary.Float(32000.1),
		TodIceFuel:         unifieddatalibrary.Float(2000.1),
		UnidentExtraFuel:   unifieddatalibrary.Float(5000.1),
		UnusableFuel:       unifieddatalibrary.Float(2300.1),
		WakeTurbCat:        unifieddatalibrary.String("MEDIUM"),
		WindFac1:           unifieddatalibrary.Float(-1.1),
		WindFac2:           unifieddatalibrary.Float(10.1),
		WindFacAvg:         unifieddatalibrary.Float(5.1),
		WxValidEnd:         unifieddatalibrary.Time(time.Now()),
		WxValidStart:       unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFlightplanGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Flightplan.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.FlightplanGetParams{
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

func TestFlightplanUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Flightplan.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.FlightplanUpdateParams{
			ArrAirfield:           "KCHS",
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.FlightplanUpdateParamsDataModeTest,
			DepAirfield:           "KSLV",
			GenTs:                 time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("c44b0a80-9fef-63d9-6267-79037fb93e4c"),
			AircraftMds:           unifieddatalibrary.String("KC-130 HERCULES"),
			AirRefuelEvents: []unifieddatalibrary.FlightplanUpdateParamsAirRefuelEvent{{
				ArDegrade:       unifieddatalibrary.Float(3.1),
				ArExchangedFuel: unifieddatalibrary.Float(1500.1),
				ArNum:           unifieddatalibrary.Int(2),
				DivertFuel:      unifieddatalibrary.Float(143000.1),
				ExitFuel:        unifieddatalibrary.Float(160000.1),
			}},
			AmcMissionID:      unifieddatalibrary.String("AJM7939B1123"),
			AppLandingFuel:    unifieddatalibrary.Float(3000.1),
			ArrAlternate1:     unifieddatalibrary.String("EDDS"),
			ArrAlternate1Fuel: unifieddatalibrary.Float(6000.1),
			ArrAlternate2:     unifieddatalibrary.String("EDDM"),
			ArrAlternate2Fuel: unifieddatalibrary.Float(6000.1),
			ArrIceFuel:        unifieddatalibrary.Float(1000.1),
			ArrRunway:         unifieddatalibrary.String("05L"),
			AtcAddresses:      []string{"EYCBZMFO", "EUCHZMFP", "ETARYXYX", "EDUUZVZI"},
			AvgTempDev:        unifieddatalibrary.Float(16.1),
			BurnedFuel:        unifieddatalibrary.Float(145000.1),
			CallSign:          unifieddatalibrary.String("HKY629"),
			CargoRemark:       unifieddatalibrary.String("Expecting 55,000 lbs. If different, call us."),
			ClimbFuel:         unifieddatalibrary.Float(7000.1),
			ClimbTime:         unifieddatalibrary.String("00:13"),
			ContingencyFuel:   unifieddatalibrary.Float(3000.1),
			CountryCodes:      []string{"US", "CA", "UK"},
			DepAlternate:      unifieddatalibrary.String("LFPO"),
			DepressFuel:       unifieddatalibrary.Float(20000.1),
			DepRunway:         unifieddatalibrary.String("05L"),
			DragIndex:         unifieddatalibrary.Float(16.9),
			EarlyDescentFuel:  unifieddatalibrary.Float(500.1),
			EnduranceTime:     unifieddatalibrary.String("08:45"),
			EnrouteFuel:       unifieddatalibrary.Float(155000.1),
			EnrouteTime:       unifieddatalibrary.String("06:30"),
			Equipment:         unifieddatalibrary.String("SDFGHIRTUWXYZ/H"),
			EstDepTime:        unifieddatalibrary.Time(time.Now()),
			EtopsAirfields:    []string{"KHSV", "KISP", "KBG", "LTBS"},
			EtopsAltAirfields: []string{"KHSV", "KISP", "KBG", "LTBS"},
			EtopsRating:       unifieddatalibrary.String("85 MINUTES"),
			EtopsValWindow:    unifieddatalibrary.String("LPLA: 0317Z-0722Z"),
			ExternalID:        unifieddatalibrary.String("AFMAPP20322347140001"),
			FlightPlanMessages: []unifieddatalibrary.FlightplanUpdateParamsFlightPlanMessage{{
				MsgText:   unifieddatalibrary.String("Message text"),
				RoutePath: unifieddatalibrary.String("PRIMARY"),
				Severity:  unifieddatalibrary.String("SEVERE"),
				WpNum:     unifieddatalibrary.String("20"),
			}},
			FlightPlanPointGroups: []unifieddatalibrary.FlightplanUpdateParamsFlightPlanPointGroup{{
				AvgFuelFlow:        unifieddatalibrary.Float(19693.1),
				EtopsAvgWindFactor: unifieddatalibrary.Float(10.1),
				EtopsDistance:      unifieddatalibrary.Float(684.1),
				EtopsReqFuel:       unifieddatalibrary.Float(4412.1),
				EtopsTempDev:       unifieddatalibrary.Float(9.1),
				EtopsTime:          unifieddatalibrary.String("01:23"),
				FlightPlanPoints: []unifieddatalibrary.FlightplanUpdateParamsFlightPlanPointGroupFlightPlanPoint{{
					FppEta:     unifieddatalibrary.Time(time.Now()),
					FppLat:     unifieddatalibrary.Float(45.23),
					FppLon:     unifieddatalibrary.Float(179.1),
					FppReqFuel: unifieddatalibrary.Float(4250.1),
					PointName:  unifieddatalibrary.String("CRUISE ALTITUDE ETP"),
				}},
				FromTakeoffTime:   unifieddatalibrary.String("07:29"),
				FsafAvgWindFactor: unifieddatalibrary.Float(10.1),
				FsafDistance:      unifieddatalibrary.Float(684.1),
				FsafReqFuel:       unifieddatalibrary.Float(50380.1),
				FsafTempDev:       unifieddatalibrary.Float(9.1),
				FsafTime:          unifieddatalibrary.String("01:23"),
				FuelCalcAlt:       unifieddatalibrary.Float(100.1),
				FuelCalcSpd:       unifieddatalibrary.Float(365.1),
				LsafAvgWindFactor: unifieddatalibrary.Float(13.1),
				LsafDistance:      unifieddatalibrary.Float(684.1),
				LsafName:          unifieddatalibrary.String("LPPD"),
				LsafReqFuel:       unifieddatalibrary.Float(50787.1),
				LsafTempDev:       unifieddatalibrary.Float(9.1),
				LsafTime:          unifieddatalibrary.String("01:23"),
				PlannedFuel:       unifieddatalibrary.Float(190319.1),
				PointGroupName:    unifieddatalibrary.String("ETOPS_CF_POINT_1"),
				WorstFuelCase:     unifieddatalibrary.String("DEPRESSURIZED ENGINE OUT ETP"),
			}},
			FlightPlanWaypoints: []unifieddatalibrary.FlightplanUpdateParamsFlightPlanWaypoint{{
				Type:               "COMMENT",
				WaypointName:       "KCHS",
				AaTacanChannel:     unifieddatalibrary.String("31/94"),
				AirDistance:        unifieddatalibrary.Float(321.1),
				Airway:             unifieddatalibrary.String("W15"),
				Alt:                unifieddatalibrary.Float(27000.1),
				ArID:               unifieddatalibrary.String("AR202"),
				Arpt:               unifieddatalibrary.String("ARIP"),
				Ata:                unifieddatalibrary.Time(time.Now()),
				AvgCalAirspeed:     unifieddatalibrary.Float(200.1),
				AvgDriftAng:        unifieddatalibrary.Float(-3.2),
				AvgGroundSpeed:     unifieddatalibrary.Float(300.1),
				AvgTrueAirspeed:    unifieddatalibrary.Float(210.1),
				AvgWindDir:         unifieddatalibrary.Float(165.5),
				AvgWindSpeed:       unifieddatalibrary.Float(14.4),
				DayLowAlt:          unifieddatalibrary.Float(1500.1),
				Eta:                unifieddatalibrary.Time(time.Now()),
				ExchangedFuel:      unifieddatalibrary.Float(-30400.1),
				FuelFlow:           unifieddatalibrary.Float(17654.1),
				IceCat:             unifieddatalibrary.String("MODERATE"),
				Lat:                unifieddatalibrary.Float(45.23),
				LegAlternate:       unifieddatalibrary.String("KCHS"),
				LegDragIndex:       unifieddatalibrary.Float(1.2),
				LegFuelDegrade:     unifieddatalibrary.Float(10.1),
				LegMach:            unifieddatalibrary.Float(0.74),
				LegMsnIndex:        unifieddatalibrary.Float(65),
				LegWindFac:         unifieddatalibrary.Float(-32.1),
				Lon:                unifieddatalibrary.Float(179.1),
				MagCourse:          unifieddatalibrary.Float(338.1),
				MagHeading:         unifieddatalibrary.Float(212.1),
				MagVar:             unifieddatalibrary.Float(-13.2),
				Navaid:             unifieddatalibrary.String("HTO"),
				NightLowAlt:        unifieddatalibrary.Float(2300.1),
				NvgLowAlt:          unifieddatalibrary.Float(2450.1),
				PointWindDir:       unifieddatalibrary.Float(165.5),
				PointWindSpeed:     unifieddatalibrary.Float(14.4),
				PriFreq:            unifieddatalibrary.Float(357.5),
				SecFreq:            unifieddatalibrary.Float(357.5),
				TacanChannel:       unifieddatalibrary.String("83X"),
				TempDev:            unifieddatalibrary.Float(12.1),
				ThunderCat:         unifieddatalibrary.String("MODERATE"),
				TotalAirDistance:   unifieddatalibrary.Float(3251.1),
				TotalFlownDistance: unifieddatalibrary.Float(688.1),
				TotalRemDistance:   unifieddatalibrary.Float(1288.1),
				TotalRemFuel:       unifieddatalibrary.Float(30453.1),
				TotalTime:          unifieddatalibrary.String("08:45"),
				TotalTimeRem:       unifieddatalibrary.String("01:43"),
				TotalUsedFuel:      unifieddatalibrary.Float(70431.1),
				TotalWeight:        unifieddatalibrary.Float(207123.1),
				TrueCourse:         unifieddatalibrary.Float(328.1),
				TurbCat:            unifieddatalibrary.String("EXTREME"),
				VorFreq:            unifieddatalibrary.Float(113.6),
				WaypointNum:        unifieddatalibrary.Int(20),
				ZoneDistance:       unifieddatalibrary.Float(212.1),
				ZoneFuel:           unifieddatalibrary.Float(1120.1),
				ZoneTime:           unifieddatalibrary.Float(36.1),
			}},
			FlightRules:        unifieddatalibrary.String("l"),
			FlightType:         unifieddatalibrary.String("MILITARY"),
			FuelDegrade:        unifieddatalibrary.Float(10.3),
			GpsRaim:            unifieddatalibrary.String("Failed by FAA SAPT 184022AUG2022"),
			HoldDownFuel:       unifieddatalibrary.Float(500.1),
			HoldFuel:           unifieddatalibrary.Float(6000.1),
			HoldTime:           unifieddatalibrary.String("01:00"),
			IDAircraft:         unifieddatalibrary.String("4f4a67c6-40fd-11ee-be56-0242ac120002"),
			IDArrAirfield:      unifieddatalibrary.String("363080c2-40fd-11ee-be56-0242ac120002"),
			IDDepAirfield:      unifieddatalibrary.String("2a9020f6-40fd-11ee-be56-0242ac120002"),
			IdentExtraFuel:     unifieddatalibrary.Float(5000.1),
			IDSortie:           unifieddatalibrary.String("9d60c1b1-10b1-b2a7-e403-84c5d7eeb170"),
			InitialCruiseSpeed: unifieddatalibrary.String("N0305"),
			InitialFlightLevel: unifieddatalibrary.String("F270"),
			LandingFuel:        unifieddatalibrary.Float(19000.1),
			LegNum:             unifieddatalibrary.Int(100),
			MinDivertFuel:      unifieddatalibrary.Float(25000.1),
			MsnIndex:           unifieddatalibrary.Float(44.1),
			Notes:              unifieddatalibrary.String("STS/STATE PBN/A1B2B5C2C4D2D4 EUR/PROTECTED"),
			NumAircraft:        unifieddatalibrary.Int(1),
			OpConditionFuel:    unifieddatalibrary.Float(5000.1),
			OpWeight:           unifieddatalibrary.Float(251830.5),
			Origin:             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Originator:         unifieddatalibrary.String("ETARYXYX"),
			PlannerRemark:      unifieddatalibrary.String("Flight plan is good for 2 days before airspace closes over the UK."),
			RampFuel:           unifieddatalibrary.Float(180000.1),
			RemAlternate1Fuel:  unifieddatalibrary.Float(18000.1),
			RemAlternate2Fuel:  unifieddatalibrary.Float(18000.1),
			ReserveFuel:        unifieddatalibrary.Float(10000.1),
			RouteString:        unifieddatalibrary.String("RENV3B RENVI Y86 GOSVA/N0317F260 DCT EVLIT DCT UMUGI DCT NISIX DCT GIGOD DCT DIPEB DCT\nGORPI Z80 TILAV L87 RAKIT Z717 PODUS Z130 MAG/N0298F220 Z20 KENIG/N0319F220 Z20 ORTAG T177\nESEGU Z20 BEBLA DCT MASEK/N0300F200 DCT GISEM/N0319F200 DCT BOMBI/N0276F060 DCT RIDSU DCT"),
			Sid:                unifieddatalibrary.String("RENV3B"),
			Star:               unifieddatalibrary.String("ADANA"),
			Status:             unifieddatalibrary.String("APPROVED"),
			TailNumber:         unifieddatalibrary.String("77187"),
			TakeoffFuel:        unifieddatalibrary.Float(178500.1),
			TaxiFuel:           unifieddatalibrary.Float(1500.1),
			ThunderAvoidFuel:   unifieddatalibrary.Float(1000.1),
			TocFuel:            unifieddatalibrary.Float(160000.1),
			TocIceFuel:         unifieddatalibrary.Float(1000.1),
			TodFuel:            unifieddatalibrary.Float(32000.1),
			TodIceFuel:         unifieddatalibrary.Float(2000.1),
			UnidentExtraFuel:   unifieddatalibrary.Float(5000.1),
			UnusableFuel:       unifieddatalibrary.Float(2300.1),
			WakeTurbCat:        unifieddatalibrary.String("MEDIUM"),
			WindFac1:           unifieddatalibrary.Float(-1.1),
			WindFac2:           unifieddatalibrary.Float(10.1),
			WindFacAvg:         unifieddatalibrary.Float(5.1),
			WxValidEnd:         unifieddatalibrary.Time(time.Now()),
			WxValidStart:       unifieddatalibrary.Time(time.Now()),
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

func TestFlightplanListWithOptionalParams(t *testing.T) {
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
	_, err := client.Flightplan.List(context.TODO(), unifieddatalibrary.FlightplanListParams{
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

func TestFlightplanDelete(t *testing.T) {
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
	err := client.Flightplan.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFlightplanCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Flightplan.Count(context.TODO(), unifieddatalibrary.FlightplanCountParams{
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

func TestFlightplanQueryhelp(t *testing.T) {
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
	_, err := client.Flightplan.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFlightplanTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Flightplan.Tuple(context.TODO(), unifieddatalibrary.FlightplanTupleParams{
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

func TestFlightplanUnvalidatedPublish(t *testing.T) {
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
	err := client.Flightplan.UnvalidatedPublish(context.TODO(), unifieddatalibrary.FlightplanUnvalidatedPublishParams{
		Body: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBody{{
			ArrAirfield:           "KCHS",
			ClassificationMarking: "U",
			DataMode:              "TEST",
			DepAirfield:           "KSLV",
			GenTs:                 time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("c44b0a80-9fef-63d9-6267-79037fb93e4c"),
			AircraftMds:           unifieddatalibrary.String("KC-130 HERCULES"),
			AirRefuelEvents: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBodyAirRefuelEvent{{
				ArDegrade:       unifieddatalibrary.Float(3.1),
				ArExchangedFuel: unifieddatalibrary.Float(1500.1),
				ArNum:           unifieddatalibrary.Int(2),
				DivertFuel:      unifieddatalibrary.Float(143000.1),
				ExitFuel:        unifieddatalibrary.Float(160000.1),
			}},
			AmcMissionID:      unifieddatalibrary.String("AJM7939B1123"),
			AppLandingFuel:    unifieddatalibrary.Float(3000.1),
			ArrAlternate1:     unifieddatalibrary.String("EDDS"),
			ArrAlternate1Fuel: unifieddatalibrary.Float(6000.1),
			ArrAlternate2:     unifieddatalibrary.String("EDDM"),
			ArrAlternate2Fuel: unifieddatalibrary.Float(6000.1),
			ArrIceFuel:        unifieddatalibrary.Float(1000.1),
			ArrRunway:         unifieddatalibrary.String("05L"),
			AtcAddresses:      []string{"EYCBZMFO", "EUCHZMFP", "ETARYXYX", "EDUUZVZI"},
			AvgTempDev:        unifieddatalibrary.Float(16.1),
			BurnedFuel:        unifieddatalibrary.Float(145000.1),
			CallSign:          unifieddatalibrary.String("HKY629"),
			CargoRemark:       unifieddatalibrary.String("Expecting 55,000 lbs. If different, call us."),
			ClimbFuel:         unifieddatalibrary.Float(7000.1),
			ClimbTime:         unifieddatalibrary.String("00:13"),
			ContingencyFuel:   unifieddatalibrary.Float(3000.1),
			CountryCodes:      []string{"US", "CA", "UK"},
			DepAlternate:      unifieddatalibrary.String("LFPO"),
			DepressFuel:       unifieddatalibrary.Float(20000.1),
			DepRunway:         unifieddatalibrary.String("05L"),
			DragIndex:         unifieddatalibrary.Float(16.9),
			EarlyDescentFuel:  unifieddatalibrary.Float(500.1),
			EnduranceTime:     unifieddatalibrary.String("08:45"),
			EnrouteFuel:       unifieddatalibrary.Float(155000.1),
			EnrouteTime:       unifieddatalibrary.String("06:30"),
			Equipment:         unifieddatalibrary.String("SDFGHIRTUWXYZ/H"),
			EstDepTime:        unifieddatalibrary.Time(time.Now()),
			EtopsAirfields:    []string{"KHSV", "KISP", "KBG", "LTBS"},
			EtopsAltAirfields: []string{"KHSV", "KISP", "KBG", "LTBS"},
			EtopsRating:       unifieddatalibrary.String("85 MINUTES"),
			EtopsValWindow:    unifieddatalibrary.String("LPLA: 0317Z-0722Z"),
			ExternalID:        unifieddatalibrary.String("AFMAPP20322347140001"),
			FlightPlanMessages: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBodyFlightPlanMessage{{
				MsgText:   unifieddatalibrary.String("Message text"),
				RoutePath: unifieddatalibrary.String("PRIMARY"),
				Severity:  unifieddatalibrary.String("SEVERE"),
				WpNum:     unifieddatalibrary.String("20"),
			}},
			FlightPlanPointGroups: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroup{{
				AvgFuelFlow:        unifieddatalibrary.Float(19693.1),
				EtopsAvgWindFactor: unifieddatalibrary.Float(10.1),
				EtopsDistance:      unifieddatalibrary.Float(684.1),
				EtopsReqFuel:       unifieddatalibrary.Float(4412.1),
				EtopsTempDev:       unifieddatalibrary.Float(9.1),
				EtopsTime:          unifieddatalibrary.String("01:23"),
				FlightPlanPoints: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBodyFlightPlanPointGroupFlightPlanPoint{{
					FppEta:     unifieddatalibrary.Time(time.Now()),
					FppLat:     unifieddatalibrary.Float(45.23),
					FppLon:     unifieddatalibrary.Float(179.1),
					FppReqFuel: unifieddatalibrary.Float(4250.1),
					PointName:  unifieddatalibrary.String("CRUISE ALTITUDE ETP"),
				}},
				FromTakeoffTime:   unifieddatalibrary.String("07:29"),
				FsafAvgWindFactor: unifieddatalibrary.Float(10.1),
				FsafDistance:      unifieddatalibrary.Float(684.1),
				FsafReqFuel:       unifieddatalibrary.Float(50380.1),
				FsafTempDev:       unifieddatalibrary.Float(9.1),
				FsafTime:          unifieddatalibrary.String("01:23"),
				FuelCalcAlt:       unifieddatalibrary.Float(100.1),
				FuelCalcSpd:       unifieddatalibrary.Float(365.1),
				LsafAvgWindFactor: unifieddatalibrary.Float(13.1),
				LsafDistance:      unifieddatalibrary.Float(684.1),
				LsafName:          unifieddatalibrary.String("LPPD"),
				LsafReqFuel:       unifieddatalibrary.Float(50787.1),
				LsafTempDev:       unifieddatalibrary.Float(9.1),
				LsafTime:          unifieddatalibrary.String("01:23"),
				PlannedFuel:       unifieddatalibrary.Float(190319.1),
				PointGroupName:    unifieddatalibrary.String("ETOPS_CF_POINT_1"),
				WorstFuelCase:     unifieddatalibrary.String("DEPRESSURIZED ENGINE OUT ETP"),
			}},
			FlightPlanWaypoints: []unifieddatalibrary.FlightplanUnvalidatedPublishParamsBodyFlightPlanWaypoint{{
				Type:               "COMMENT",
				WaypointName:       "KCHS",
				AaTacanChannel:     unifieddatalibrary.String("31/94"),
				AirDistance:        unifieddatalibrary.Float(321.1),
				Airway:             unifieddatalibrary.String("W15"),
				Alt:                unifieddatalibrary.Float(27000.1),
				ArID:               unifieddatalibrary.String("AR202"),
				Arpt:               unifieddatalibrary.String("ARIP"),
				Ata:                unifieddatalibrary.Time(time.Now()),
				AvgCalAirspeed:     unifieddatalibrary.Float(200.1),
				AvgDriftAng:        unifieddatalibrary.Float(-3.2),
				AvgGroundSpeed:     unifieddatalibrary.Float(300.1),
				AvgTrueAirspeed:    unifieddatalibrary.Float(210.1),
				AvgWindDir:         unifieddatalibrary.Float(165.5),
				AvgWindSpeed:       unifieddatalibrary.Float(14.4),
				DayLowAlt:          unifieddatalibrary.Float(1500.1),
				Eta:                unifieddatalibrary.Time(time.Now()),
				ExchangedFuel:      unifieddatalibrary.Float(-30400.1),
				FuelFlow:           unifieddatalibrary.Float(17654.1),
				IceCat:             unifieddatalibrary.String("MODERATE"),
				Lat:                unifieddatalibrary.Float(45.23),
				LegAlternate:       unifieddatalibrary.String("KCHS"),
				LegDragIndex:       unifieddatalibrary.Float(1.2),
				LegFuelDegrade:     unifieddatalibrary.Float(10.1),
				LegMach:            unifieddatalibrary.Float(0.74),
				LegMsnIndex:        unifieddatalibrary.Float(65),
				LegWindFac:         unifieddatalibrary.Float(-32.1),
				Lon:                unifieddatalibrary.Float(179.1),
				MagCourse:          unifieddatalibrary.Float(338.1),
				MagHeading:         unifieddatalibrary.Float(212.1),
				MagVar:             unifieddatalibrary.Float(-13.2),
				Navaid:             unifieddatalibrary.String("HTO"),
				NightLowAlt:        unifieddatalibrary.Float(2300.1),
				NvgLowAlt:          unifieddatalibrary.Float(2450.1),
				PointWindDir:       unifieddatalibrary.Float(165.5),
				PointWindSpeed:     unifieddatalibrary.Float(14.4),
				PriFreq:            unifieddatalibrary.Float(357.5),
				SecFreq:            unifieddatalibrary.Float(357.5),
				TacanChannel:       unifieddatalibrary.String("83X"),
				TempDev:            unifieddatalibrary.Float(12.1),
				ThunderCat:         unifieddatalibrary.String("MODERATE"),
				TotalAirDistance:   unifieddatalibrary.Float(3251.1),
				TotalFlownDistance: unifieddatalibrary.Float(688.1),
				TotalRemDistance:   unifieddatalibrary.Float(1288.1),
				TotalRemFuel:       unifieddatalibrary.Float(30453.1),
				TotalTime:          unifieddatalibrary.String("08:45"),
				TotalTimeRem:       unifieddatalibrary.String("01:43"),
				TotalUsedFuel:      unifieddatalibrary.Float(70431.1),
				TotalWeight:        unifieddatalibrary.Float(207123.1),
				TrueCourse:         unifieddatalibrary.Float(328.1),
				TurbCat:            unifieddatalibrary.String("EXTREME"),
				VorFreq:            unifieddatalibrary.Float(113.6),
				WaypointNum:        unifieddatalibrary.Int(20),
				ZoneDistance:       unifieddatalibrary.Float(212.1),
				ZoneFuel:           unifieddatalibrary.Float(1120.1),
				ZoneTime:           unifieddatalibrary.Float(36.1),
			}},
			FlightRules:        unifieddatalibrary.String("l"),
			FlightType:         unifieddatalibrary.String("MILITARY"),
			FuelDegrade:        unifieddatalibrary.Float(10.3),
			GpsRaim:            unifieddatalibrary.String("Failed by FAA SAPT 184022AUG2022"),
			HoldDownFuel:       unifieddatalibrary.Float(500.1),
			HoldFuel:           unifieddatalibrary.Float(6000.1),
			HoldTime:           unifieddatalibrary.String("01:00"),
			IDAircraft:         unifieddatalibrary.String("4f4a67c6-40fd-11ee-be56-0242ac120002"),
			IDArrAirfield:      unifieddatalibrary.String("363080c2-40fd-11ee-be56-0242ac120002"),
			IDDepAirfield:      unifieddatalibrary.String("2a9020f6-40fd-11ee-be56-0242ac120002"),
			IdentExtraFuel:     unifieddatalibrary.Float(5000.1),
			IDSortie:           unifieddatalibrary.String("9d60c1b1-10b1-b2a7-e403-84c5d7eeb170"),
			InitialCruiseSpeed: unifieddatalibrary.String("N0305"),
			InitialFlightLevel: unifieddatalibrary.String("F270"),
			LandingFuel:        unifieddatalibrary.Float(19000.1),
			LegNum:             unifieddatalibrary.Int(100),
			MinDivertFuel:      unifieddatalibrary.Float(25000.1),
			MsnIndex:           unifieddatalibrary.Float(44.1),
			Notes:              unifieddatalibrary.String("STS/STATE PBN/A1B2B5C2C4D2D4 EUR/PROTECTED"),
			NumAircraft:        unifieddatalibrary.Int(1),
			OpConditionFuel:    unifieddatalibrary.Float(5000.1),
			OpWeight:           unifieddatalibrary.Float(251830.5),
			Origin:             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Originator:         unifieddatalibrary.String("ETARYXYX"),
			PlannerRemark:      unifieddatalibrary.String("Flight plan is good for 2 days before airspace closes over the UK."),
			RampFuel:           unifieddatalibrary.Float(180000.1),
			RemAlternate1Fuel:  unifieddatalibrary.Float(18000.1),
			RemAlternate2Fuel:  unifieddatalibrary.Float(18000.1),
			ReserveFuel:        unifieddatalibrary.Float(10000.1),
			RouteString:        unifieddatalibrary.String("RENV3B RENVI Y86 GOSVA/N0317F260 DCT EVLIT DCT UMUGI DCT NISIX DCT GIGOD DCT DIPEB DCT\nGORPI Z80 TILAV L87 RAKIT Z717 PODUS Z130 MAG/N0298F220 Z20 KENIG/N0319F220 Z20 ORTAG T177\nESEGU Z20 BEBLA DCT MASEK/N0300F200 DCT GISEM/N0319F200 DCT BOMBI/N0276F060 DCT RIDSU DCT"),
			Sid:                unifieddatalibrary.String("RENV3B"),
			Star:               unifieddatalibrary.String("ADANA"),
			Status:             unifieddatalibrary.String("APPROVED"),
			TailNumber:         unifieddatalibrary.String("77187"),
			TakeoffFuel:        unifieddatalibrary.Float(178500.1),
			TaxiFuel:           unifieddatalibrary.Float(1500.1),
			ThunderAvoidFuel:   unifieddatalibrary.Float(1000.1),
			TocFuel:            unifieddatalibrary.Float(160000.1),
			TocIceFuel:         unifieddatalibrary.Float(1000.1),
			TodFuel:            unifieddatalibrary.Float(32000.1),
			TodIceFuel:         unifieddatalibrary.Float(2000.1),
			UnidentExtraFuel:   unifieddatalibrary.Float(5000.1),
			UnusableFuel:       unifieddatalibrary.Float(2300.1),
			WakeTurbCat:        unifieddatalibrary.String("MEDIUM"),
			WindFac1:           unifieddatalibrary.Float(-1.1),
			WindFac2:           unifieddatalibrary.Float(10.1),
			WindFacAvg:         unifieddatalibrary.Float(5.1),
			WxValidEnd:         unifieddatalibrary.Time(time.Now()),
			WxValidStart:       unifieddatalibrary.Time(time.Now()),
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
