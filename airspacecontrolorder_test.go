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

func TestAirspaceControlOrderNewWithOptionalParams(t *testing.T) {
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
	err := client.AirspaceControlOrders.New(context.TODO(), unifieddatalibrary.AirspaceControlOrderNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirspaceControlOrderNewParamsDataModeTest,
		OpExName:              "DESERT WIND",
		Originator:            "USCENTCOM",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("c44b0a80-9fef-63d9-6267-79037fb93e4c"),
		AcoComments:           unifieddatalibrary.String("CHOKE POINTS"),
		AcoSerialNum:          unifieddatalibrary.String("27B"),
		AirspaceControlMeansStatus: []unifieddatalibrary.AirspaceControlOrderNewParamsAirspaceControlMeansStatus{{
			AirspaceControlMeans: []unifieddatalibrary.AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMean{{
				AirspaceControlPoint: []unifieddatalibrary.AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint{{
					CtrlPtAltitude: unifieddatalibrary.String("BRFL:MSL-FL230"),
					CtrlPtLocation: unifieddatalibrary.String("203632N0594256E"),
					CtrlPtName:     unifieddatalibrary.String("APPLE"),
					CtrlPtType:     unifieddatalibrary.String("CP"),
				}},
				AirspaceTimePeriod: []unifieddatalibrary.AirspaceControlOrderNewParamsAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod{{
					IntDur:    []string{"65WK"},
					IntFreq:   []string{"WEEKLY"},
					TimeEnd:   unifieddatalibrary.String("141325ZFEB2002"),
					TimeMode:  unifieddatalibrary.String("DISCRETE"),
					TimeStart: unifieddatalibrary.String("141325ZFEB2002"),
				}},
				Bearing0:       unifieddatalibrary.Float(330),
				Bearing1:       unifieddatalibrary.Float(160),
				CmID:           unifieddatalibrary.String("DESIG:C34"),
				CmShape:        "POLYARC",
				CmType:         unifieddatalibrary.String("cmType"),
				CntrlAuth:      unifieddatalibrary.String("RHEIN MAIN CP"),
				CntrlAuthFreqs: []string{"125.25MHZ"},
				Coord0:         unifieddatalibrary.String("152345N0505657E"),
				Coord1:         unifieddatalibrary.String("1523N05057E"),
				CorrWayPoints:  []string{"POB", "RDU", "IAD"},
				EffVDim:        unifieddatalibrary.String("BRRA:GL-100AGL"),
				FreeText:       unifieddatalibrary.String("1. CAPACITY: MDM TK, 50 VEHICLE CONVOY. 2. CHOKE POINTS: EXIT 5"),
				GenTextInd:     unifieddatalibrary.String("SITUATION"),
				GeoDatumAlt:    unifieddatalibrary.String("NAR"),
				Link16ID:       unifieddatalibrary.String("F3356"),
				OrbitAlignment: unifieddatalibrary.String("C"),
				PolyCoord:      []string{"203632N0594256E", "155000N0594815E", "155000N0591343E"},
				RadMag0:        unifieddatalibrary.Float(30.04),
				RadMag1:        unifieddatalibrary.Float(50.12),
				RadMagUnit:     unifieddatalibrary.String("NM"),
				TrackLeg:       unifieddatalibrary.Int(99),
				TransAltitude:  unifieddatalibrary.String("18000FT"),
				Usage:          unifieddatalibrary.String("usage"),
				Width:          unifieddatalibrary.Float(15.6),
				WidthLeft:      unifieddatalibrary.Float(5.2),
				WidthRight:     unifieddatalibrary.Float(10.4),
				WidthUnit:      unifieddatalibrary.String("KM"),
			}},
			CmStat:   unifieddatalibrary.String("ADD"),
			CmStatID: []string{"DESIGN:B35", "NAME:ERMA", "RANG:C21-C25"},
		}},
		AirspaceControlOrderReferences: []unifieddatalibrary.AirspaceControlOrderNewParamsAirspaceControlOrderReference{{
			RefOriginator:      unifieddatalibrary.String("SHAPE"),
			RefSerialNum:       unifieddatalibrary.String("100"),
			RefSiCs:            []string{"RCA", "FN:4503B"},
			RefSID:             unifieddatalibrary.String("A"),
			RefSpecialNotation: unifieddatalibrary.String("NOTAL"),
			RefTs:              unifieddatalibrary.Time(time.Now()),
			RefType:            unifieddatalibrary.String("NBC1"),
		}},
		AreaOfValidity:        unifieddatalibrary.String("FORT BRAGG"),
		ClassReasons:          []string{"15C", "10C"},
		ClassSource:           unifieddatalibrary.String("ORIG:USJFCOM"),
		DeclassExemptionCodes: []string{"X1", "X2"},
		DowngradeInsDates:     []string{"NST:AT EXERCISE ENDEX", "DATE:25NOV1997"},
		GeoDatum:              unifieddatalibrary.String("EUR-T"),
		Month:                 unifieddatalibrary.String("OCT"),
		OpExInfo:              unifieddatalibrary.String("CONTROL"),
		OpExInfoAlt:           unifieddatalibrary.String("ORANGE"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PlanOrigNum:           unifieddatalibrary.String("SACEUR 106"),
		Qualifier:             unifieddatalibrary.String("CHG"),
		QualSn:                unifieddatalibrary.Int(1),
		SerialNum:             unifieddatalibrary.String("1201003"),
		StopQualifier:         unifieddatalibrary.String("AFTER"),
		StopTime:              unifieddatalibrary.Time(time.Now()),
		UndLnkTrks:            []string{"A2467", "A3466", "AA232"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirspaceControlOrderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirspaceControlOrders.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirspaceControlOrderGetParams{
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

func TestAirspaceControlOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirspaceControlOrders.List(context.TODO(), unifieddatalibrary.AirspaceControlOrderListParams{
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

func TestAirspaceControlOrderCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirspaceControlOrders.Count(context.TODO(), unifieddatalibrary.AirspaceControlOrderCountParams{
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

func TestAirspaceControlOrderNewBulk(t *testing.T) {
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
	err := client.AirspaceControlOrders.NewBulk(context.TODO(), unifieddatalibrary.AirspaceControlOrderNewBulkParams{
		Body: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			OpExName:              "DESERT WIND",
			Originator:            "USCENTCOM",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("c44b0a80-9fef-63d9-6267-79037fb93e4c"),
			AcoComments:           unifieddatalibrary.String("CHOKE POINTS"),
			AcoSerialNum:          unifieddatalibrary.String("27B"),
			AirspaceControlMeansStatus: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatus{{
				AirspaceControlMeans: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMean{{
					AirspaceControlPoint: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceControlPoint{{
						CtrlPtAltitude: unifieddatalibrary.String("BRFL:MSL-FL230"),
						CtrlPtLocation: unifieddatalibrary.String("203632N0594256E"),
						CtrlPtName:     unifieddatalibrary.String("APPLE"),
						CtrlPtType:     unifieddatalibrary.String("CP"),
					}},
					AirspaceTimePeriod: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBodyAirspaceControlMeansStatusAirspaceControlMeanAirspaceTimePeriod{{
						IntDur:    []string{"65WK"},
						IntFreq:   []string{"WEEKLY"},
						TimeEnd:   unifieddatalibrary.String("141325ZFEB2002"),
						TimeMode:  unifieddatalibrary.String("DISCRETE"),
						TimeStart: unifieddatalibrary.String("141325ZFEB2002"),
					}},
					Bearing0:       unifieddatalibrary.Float(330),
					Bearing1:       unifieddatalibrary.Float(160),
					CmID:           unifieddatalibrary.String("DESIG:C34"),
					CmShape:        "POLYARC",
					CmType:         unifieddatalibrary.String("cmType"),
					CntrlAuth:      unifieddatalibrary.String("RHEIN MAIN CP"),
					CntrlAuthFreqs: []string{"125.25MHZ"},
					Coord0:         unifieddatalibrary.String("152345N0505657E"),
					Coord1:         unifieddatalibrary.String("1523N05057E"),
					CorrWayPoints:  []string{"POB", "RDU", "IAD"},
					EffVDim:        unifieddatalibrary.String("BRRA:GL-100AGL"),
					FreeText:       unifieddatalibrary.String("1. CAPACITY: MDM TK, 50 VEHICLE CONVOY. 2. CHOKE POINTS: EXIT 5"),
					GenTextInd:     unifieddatalibrary.String("SITUATION"),
					GeoDatumAlt:    unifieddatalibrary.String("NAR"),
					Link16ID:       unifieddatalibrary.String("F3356"),
					OrbitAlignment: unifieddatalibrary.String("C"),
					PolyCoord:      []string{"203632N0594256E", "155000N0594815E", "155000N0591343E"},
					RadMag0:        unifieddatalibrary.Float(30.04),
					RadMag1:        unifieddatalibrary.Float(50.12),
					RadMagUnit:     unifieddatalibrary.String("NM"),
					TrackLeg:       unifieddatalibrary.Int(99),
					TransAltitude:  unifieddatalibrary.String("18000FT"),
					Usage:          unifieddatalibrary.String("usage"),
					Width:          unifieddatalibrary.Float(15.6),
					WidthLeft:      unifieddatalibrary.Float(5.2),
					WidthRight:     unifieddatalibrary.Float(10.4),
					WidthUnit:      unifieddatalibrary.String("KM"),
				}},
				CmStat:   unifieddatalibrary.String("ADD"),
				CmStatID: []string{"DESIGN:B35", "NAME:ERMA", "RANG:C21-C25"},
			}},
			AirspaceControlOrderReferences: []unifieddatalibrary.AirspaceControlOrderNewBulkParamsBodyAirspaceControlOrderReference{{
				RefOriginator:      unifieddatalibrary.String("SHAPE"),
				RefSerialNum:       unifieddatalibrary.String("100"),
				RefSiCs:            []string{"RCA", "FN:4503B"},
				RefSID:             unifieddatalibrary.String("A"),
				RefSpecialNotation: unifieddatalibrary.String("NOTAL"),
				RefTs:              unifieddatalibrary.Time(time.Now()),
				RefType:            unifieddatalibrary.String("NBC1"),
			}},
			AreaOfValidity:        unifieddatalibrary.String("FORT BRAGG"),
			ClassReasons:          []string{"15C", "10C"},
			ClassSource:           unifieddatalibrary.String("ORIG:USJFCOM"),
			DeclassExemptionCodes: []string{"X1", "X2"},
			DowngradeInsDates:     []string{"NST:AT EXERCISE ENDEX", "DATE:25NOV1997"},
			GeoDatum:              unifieddatalibrary.String("EUR-T"),
			Month:                 unifieddatalibrary.String("OCT"),
			OpExInfo:              unifieddatalibrary.String("CONTROL"),
			OpExInfoAlt:           unifieddatalibrary.String("ORANGE"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlanOrigNum:           unifieddatalibrary.String("SACEUR 106"),
			Qualifier:             unifieddatalibrary.String("CHG"),
			QualSn:                unifieddatalibrary.Int(1),
			SerialNum:             unifieddatalibrary.String("1201003"),
			StopQualifier:         unifieddatalibrary.String("AFTER"),
			StopTime:              unifieddatalibrary.Time(time.Now()),
			UndLnkTrks:            []string{"A2467", "A3466", "AA232"},
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

func TestAirspaceControlOrderQueryHelp(t *testing.T) {
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
	_, err := client.AirspaceControlOrders.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirspaceControlOrderTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirspaceControlOrders.Tuple(context.TODO(), unifieddatalibrary.AirspaceControlOrderTupleParams{
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
