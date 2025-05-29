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

func TestAirOperationAirTaskingOrderNewWithOptionalParams(t *testing.T) {
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
	err := client.AirOperations.AirTaskingOrders.New(context.TODO(), unifieddatalibrary.AirOperationAirTaskingOrderNewParams{
		BeginTs:               time.Now(),
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirOperationAirTaskingOrderNewParamsDataModeTest,
		OpExerName:            "DESERT WIND",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("POI-ID"),
		AckReqInd:             unifieddatalibrary.String("YES"),
		AckUnitInstructions:   unifieddatalibrary.String("INST:45TS"),
		AcMsnTasking: []unifieddatalibrary.AirOperationAirTaskingOrderNewParamsAcMsnTasking{{
			CountryCode:    "US",
			TaskedService:  "A",
			UnitDesignator: "AMPHIB5DIV",
			AcMsnLocSeg: []unifieddatalibrary.AirOperationAirTaskingOrderNewParamsAcMsnTaskingAcMsnLocSeg{{
				StartTime:    time.Now(),
				AirMsnPri:    unifieddatalibrary.String("1A"),
				Alt:          unifieddatalibrary.Int(210),
				AreaGeoRad:   unifieddatalibrary.Int(1000),
				EndTime:      unifieddatalibrary.Time(time.Now()),
				MsnLocName:   unifieddatalibrary.String("KLSV"),
				MsnLocPtBarT: unifieddatalibrary.String("330T-PT ALFA-50NM"),
				MsnLocPtLat:  unifieddatalibrary.Float(35.123),
				MsnLocPtLon:  unifieddatalibrary.Float(79.01),
				MsnLocPtName: unifieddatalibrary.String("PT ALFA"),
			}},
			AlertStatus: unifieddatalibrary.Int(30),
			AmcMsnNum:   unifieddatalibrary.String("AMC:JJXD123HA045"),
			DepLocLat:   unifieddatalibrary.Float(35.123),
			DepLocLon:   unifieddatalibrary.Float(79.2354),
			DepLocName:  unifieddatalibrary.String("ICAO:KBIF"),
			DepLocUtm:   unifieddatalibrary.String("32WDL123123"),
			DepTime:     unifieddatalibrary.Time(time.Now()),
			IndAcTasking: []unifieddatalibrary.AirOperationAirTaskingOrderNewParamsAcMsnTaskingIndAcTasking{{
				AcftType:        "F35A",
				CallSign:        unifieddatalibrary.String("EAGLE47"),
				IffSifMode1Code: unifieddatalibrary.String("111"),
				IffSifMode2Code: unifieddatalibrary.String("20147"),
				IffSifMode3Code: unifieddatalibrary.String("30147"),
				JuAddress:       []int64{12345, 65432},
				Link16CallSign:  unifieddatalibrary.String("EE47"),
				NumAcft:         unifieddatalibrary.Int(2),
				PriConfigCode:   unifieddatalibrary.String("6A2W3"),
				SecConfigCode:   unifieddatalibrary.String("2S2WG"),
				TacanChan:       unifieddatalibrary.Int(123),
			}},
			MsnCommander: unifieddatalibrary.String("MC"),
			MsnNum:       unifieddatalibrary.String("D123HA"),
			PkgID:        unifieddatalibrary.String("ZZ"),
			PriMsnType:   unifieddatalibrary.String("CAS"),
			RcvyLocLat:   []float64{48.8584, 40.7554},
			RcvyLocLon:   []float64{2.2945, -73.9866},
			RcvyLocName:  []string{"ARRLOC:KBIF", "ARRLOC:KDZ7"},
			RcvyLocUtm:   []string{"ARRUTMO:32WDL123123", "ARRUTMO:32WDL321321"},
			RcvyTime:     []time.Time{time.Now(), time.Now()},
			ResMsnInd:    unifieddatalibrary.String("N"),
			SecMsnType:   unifieddatalibrary.String("SEAD"),
			UnitLocName:  unifieddatalibrary.String("ICAO:KXXQ"),
		}},
		EndTs: unifieddatalibrary.Time(time.Now()),
		GenText: []unifieddatalibrary.AirOperationAirTaskingOrderNewParamsGenText{{
			Text:    unifieddatalibrary.String("FREE-TEXT"),
			TextInd: unifieddatalibrary.String("OPENING REMARKS"),
		}},
		MsgMonth:      unifieddatalibrary.String("OCT"),
		MsgOriginator: unifieddatalibrary.String("USCENTCOM"),
		MsgQualifier:  unifieddatalibrary.String("CHG"),
		MsgSn:         unifieddatalibrary.String("ATO A"),
		NavalFltOps: []unifieddatalibrary.AirOperationAirTaskingOrderNewParamsNavalFltOp{{
			ShipName:           "USS WASP",
			FltOpStart:         unifieddatalibrary.Time(time.Now()),
			FltOpStop:          unifieddatalibrary.Time(time.Now()),
			SchdLaunchRcvyTime: []time.Time{time.Now()},
		}},
		Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAirTaskingOrderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AirTaskingOrders.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirOperationAirTaskingOrderGetParams{
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

func TestAirOperationAirTaskingOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AirTaskingOrders.List(context.TODO(), unifieddatalibrary.AirOperationAirTaskingOrderListParams{
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

func TestAirOperationAirTaskingOrderCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AirTaskingOrders.Count(context.TODO(), unifieddatalibrary.AirOperationAirTaskingOrderCountParams{
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

func TestAirOperationAirTaskingOrderQueryHelp(t *testing.T) {
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
	_, err := client.AirOperations.AirTaskingOrders.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAirTaskingOrderTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AirTaskingOrders.Tuple(context.TODO(), unifieddatalibrary.AirOperationAirTaskingOrderTupleParams{
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

func TestAirOperationAirTaskingOrderUnvalidatedPublish(t *testing.T) {
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
	err := client.AirOperations.AirTaskingOrders.UnvalidatedPublish(context.TODO(), unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParams{
		Body: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBody{{
			BeginTs:               time.Now(),
			ClassificationMarking: "U",
			DataMode:              "TEST",
			OpExerName:            "DESERT WIND",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("POI-ID"),
			AckReqInd:             unifieddatalibrary.String("YES"),
			AckUnitInstructions:   unifieddatalibrary.String("INST:45TS"),
			AcMsnTasking: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTasking{{
				CountryCode:    "US",
				TaskedService:  "A",
				UnitDesignator: "AMPHIB5DIV",
				AcMsnLocSeg: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingAcMsnLocSeg{{
					StartTime:    time.Now(),
					AirMsnPri:    unifieddatalibrary.String("1A"),
					Alt:          unifieddatalibrary.Int(210),
					AreaGeoRad:   unifieddatalibrary.Int(1000),
					EndTime:      unifieddatalibrary.Time(time.Now()),
					MsnLocName:   unifieddatalibrary.String("KLSV"),
					MsnLocPtBarT: unifieddatalibrary.String("330T-PT ALFA-50NM"),
					MsnLocPtLat:  unifieddatalibrary.Float(35.123),
					MsnLocPtLon:  unifieddatalibrary.Float(79.01),
					MsnLocPtName: unifieddatalibrary.String("PT ALFA"),
				}},
				AlertStatus: unifieddatalibrary.Int(30),
				AmcMsnNum:   unifieddatalibrary.String("AMC:JJXD123HA045"),
				DepLocLat:   unifieddatalibrary.Float(35.123),
				DepLocLon:   unifieddatalibrary.Float(79.2354),
				DepLocName:  unifieddatalibrary.String("ICAO:KBIF"),
				DepLocUtm:   unifieddatalibrary.String("32WDL123123"),
				DepTime:     unifieddatalibrary.Time(time.Now()),
				IndAcTasking: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyAcMsnTaskingIndAcTasking{{
					AcftType:        "F35A",
					CallSign:        unifieddatalibrary.String("EAGLE47"),
					IffSifMode1Code: unifieddatalibrary.String("111"),
					IffSifMode2Code: unifieddatalibrary.String("20147"),
					IffSifMode3Code: unifieddatalibrary.String("30147"),
					JuAddress:       []int64{12345, 65432},
					Link16CallSign:  unifieddatalibrary.String("EE47"),
					NumAcft:         unifieddatalibrary.Int(2),
					PriConfigCode:   unifieddatalibrary.String("6A2W3"),
					SecConfigCode:   unifieddatalibrary.String("2S2WG"),
					TacanChan:       unifieddatalibrary.Int(123),
				}},
				MsnCommander: unifieddatalibrary.String("MC"),
				MsnNum:       unifieddatalibrary.String("D123HA"),
				PkgID:        unifieddatalibrary.String("ZZ"),
				PriMsnType:   unifieddatalibrary.String("CAS"),
				RcvyLocLat:   []float64{48.8584, 40.7554},
				RcvyLocLon:   []float64{2.2945, -73.9866},
				RcvyLocName:  []string{"ARRLOC:KBIF", "ARRLOC:KDZ7"},
				RcvyLocUtm:   []string{"ARRUTMO:32WDL123123", "ARRUTMO:32WDL321321"},
				RcvyTime:     []time.Time{time.Now(), time.Now()},
				ResMsnInd:    unifieddatalibrary.String("N"),
				SecMsnType:   unifieddatalibrary.String("SEAD"),
				UnitLocName:  unifieddatalibrary.String("ICAO:KXXQ"),
			}},
			EndTs: unifieddatalibrary.Time(time.Now()),
			GenText: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyGenText{{
				Text:    unifieddatalibrary.String("FREE-TEXT"),
				TextInd: unifieddatalibrary.String("OPENING REMARKS"),
			}},
			MsgMonth:      unifieddatalibrary.String("OCT"),
			MsgOriginator: unifieddatalibrary.String("USCENTCOM"),
			MsgQualifier:  unifieddatalibrary.String("CHG"),
			MsgSn:         unifieddatalibrary.String("ATO A"),
			NavalFltOps: []unifieddatalibrary.AirOperationAirTaskingOrderUnvalidatedPublishParamsBodyNavalFltOp{{
				ShipName:           "USS WASP",
				FltOpStart:         unifieddatalibrary.Time(time.Now()),
				FltOpStop:          unifieddatalibrary.Time(time.Now()),
				SchdLaunchRcvyTime: []time.Time{time.Now()},
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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
