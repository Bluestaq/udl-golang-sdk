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

func TestTrackDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackDetails.List(context.TODO(), unifieddatalibrary.TrackDetailListParams{
		Ts:          time.Now(),
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

func TestTrackDetailCountWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackDetails.Count(context.TODO(), unifieddatalibrary.TrackDetailCountParams{
		Ts:          time.Now(),
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

func TestTrackDetailNewBulk(t *testing.T) {
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
	err := client.TrackDetails.NewBulk(context.TODO(), unifieddatalibrary.TrackDetailNewBulkParams{
		Body: []unifieddatalibrary.TrackDetailNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Lat:                   19.88550102,
			Lon:                   46.74596844,
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("TRACK-DETAILS-ID"),
			AcftSubType:           unifieddatalibrary.String("SLBM"),
			AddInfo:               unifieddatalibrary.String("Additional information"),
			Alert:                 unifieddatalibrary.String("TGT"),
			Alt:                   unifieddatalibrary.Float(153.01),
			AngElev:               unifieddatalibrary.Float(15.2),
			AntennaRefDimensions:  []float64{50.1, 50.1, 20.1, 20.1},
			AouRptData:            []float64{34.3, 26.5, 1.2},
			AouRptType:            unifieddatalibrary.String("ELLIPSE"),
			AppGrp:                unifieddatalibrary.String("GP1"),
			ArrCargo:              unifieddatalibrary.String("Freight"),
			ArrFlag:               unifieddatalibrary.String("USA"),
			ArrPort:               unifieddatalibrary.String("Lanshan"),
			ArrTime:               unifieddatalibrary.Time(time.Now()),
			Aton:                  unifieddatalibrary.String("Cardinal Mark N"),
			AvgSpd:                unifieddatalibrary.Float(18.25),
			AzCorrArcWidth:        unifieddatalibrary.Float(71.76),
			AzCorrCenterLine:      unifieddatalibrary.Float(12.876),
			BeNumber:              unifieddatalibrary.String("ENC-123"),
			Boosting:              unifieddatalibrary.Bool(false),
			BurnoutAlt:            unifieddatalibrary.Float(30567.452),
			CallSign:              unifieddatalibrary.String("Charlie"),
			CargoType:             unifieddatalibrary.String("Freight"),
			CI:                    unifieddatalibrary.String("BB"),
			Containment:           unifieddatalibrary.Float(97),
			CoopLocInd:            unifieddatalibrary.String("COOPERATIVE"),
			Course:                unifieddatalibrary.Float(4.3580153),
			Cpa:                   unifieddatalibrary.Float(500),
			DepCargo:              unifieddatalibrary.String("Freight"),
			DepFlag:               unifieddatalibrary.String("USA"),
			DepPort:               unifieddatalibrary.String("Lanshan"),
			DesCargo:              unifieddatalibrary.String("Freight"),
			DesFlag:               unifieddatalibrary.String("USA"),
			Destination:           unifieddatalibrary.String("USCLE"),
			DisID:                 unifieddatalibrary.String("7670"),
			Draught:               unifieddatalibrary.Float(21.1),
			DropPtInd:             unifieddatalibrary.Bool(false),
			Dummy:                 unifieddatalibrary.Bool(false),
			EcefPos:               []float64{1.23, 2.35, 3.42},
			EcefVel:               []float64{1.23, 2.35, 3.42},
			Elnot1:                unifieddatalibrary.String("A123A"),
			Elnot2:                unifieddatalibrary.String("A123B"),
			EmgInd:                unifieddatalibrary.Bool(false),
			EmitterName:           unifieddatalibrary.String("RAY1500"),
			Env:                   "LAND",
			ErrAreaOrient:         unifieddatalibrary.Float(69.6),
			ErrGeoAreaSwitch:      unifieddatalibrary.String("CIRCLE_ELLIPSE"),
			ErrSemiIntAxis:        unifieddatalibrary.Float(7010.882),
			ErrSemiMajElev:        unifieddatalibrary.Float(168.8),
			Eta:                   unifieddatalibrary.Time(time.Now()),
			Etd:                   unifieddatalibrary.Time(time.Now()),
			EvalRating:            unifieddatalibrary.String("A1"),
			Feint:                 unifieddatalibrary.Bool(false),
			Freq:                  unifieddatalibrary.Float(63.65),
			Ftn:                   unifieddatalibrary.String("FTN"),
			FtnCmd:                unifieddatalibrary.String("TRUETT"),
			FtnMsgTs:              unifieddatalibrary.Time(time.Now()),
			Harmonics:             unifieddatalibrary.String("8,12,4"),
			Hdng:                  unifieddatalibrary.Float(19.7),
			Hq:                    unifieddatalibrary.Bool(false),
			HullNum:               unifieddatalibrary.String("A30081"),
			HullProf:              unifieddatalibrary.String("Raised 1-23"),
			IdentAmp:              unifieddatalibrary.String("JOKER"),
			Iff:                   unifieddatalibrary.String("ID Mode"),
			Imon:                  unifieddatalibrary.Int(9015462),
			ImpactAouData:         []float64{34.3, 26.5, 1.2},
			ImpactAouType:         unifieddatalibrary.String("ELLIPSE"),
			ImpactLat:             unifieddatalibrary.Float(19.88550102),
			ImpactLon:             unifieddatalibrary.Float(46.74550102),
			ImpactTime:            unifieddatalibrary.Time(time.Now()),
			InfoSource:            unifieddatalibrary.String("S1"),
			Installation:          unifieddatalibrary.Bool(false),
			LaunchAouData:         []float64{34.3, 26.5, 1.2},
			LaunchAouType:         unifieddatalibrary.String("ELLIPSE"),
			LaunchLat:             unifieddatalibrary.Float(19.88550102),
			LaunchLon:             unifieddatalibrary.Float(46.74550102),
			LaunchTime:            unifieddatalibrary.Time(time.Now()),
			Length:                unifieddatalibrary.Float(511.1),
			LostTrkInd:            unifieddatalibrary.Bool(false),
			ManeuverInd:           unifieddatalibrary.String("POST_BOOST_NONE"),
			MaxFreq:               unifieddatalibrary.Float(10324.53),
			MidbCat:               unifieddatalibrary.String("20345"),
			Mil2525Bstr:           unifieddatalibrary.String("SHP*S----------"),
			Mmsi:                  unifieddatalibrary.Int(304010417),
			MsgType:               unifieddatalibrary.String("PLATFORM"),
			MslStatus:             unifieddatalibrary.String("AT LAUNCH"),
			MuidSrc:               unifieddatalibrary.String("MUID-SRC"),
			MuidSrcTrk:            unifieddatalibrary.String("MUID-SRC-ID"),
			Name:                  unifieddatalibrary.String("TRACK-NAME"),
			NavStatus:             unifieddatalibrary.String("Underway Using Engine"),
			Ntds:                  unifieddatalibrary.String("ZZ777"),
			NumBlades:             unifieddatalibrary.Int(4),
			NumShafts:             unifieddatalibrary.Int(3),
			ObjAct:                unifieddatalibrary.String("HOLDING"),
			ObjIdent:              "FRIEND",
			ObjNat:                unifieddatalibrary.String("USA"),
			ObjPlat:               unifieddatalibrary.String("WEAPON"),
			ObjType:               unifieddatalibrary.String("TRACKED"),
			OffPosInd:             unifieddatalibrary.String("ON"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigXref:              unifieddatalibrary.String("INT"),
			OSuffix:               unifieddatalibrary.String("AA125"),
			Pif:                   unifieddatalibrary.String("0137"),
			Pin:                   unifieddatalibrary.String("E12345012"),
			PolarSingLocLat:       unifieddatalibrary.Float(19.88550102),
			PolarSingLocLon:       unifieddatalibrary.Float(46.74550102),
			PosDeviceType:         unifieddatalibrary.String("GPS"),
			Prf:                   unifieddatalibrary.Float(17.65),
			Pri:                   unifieddatalibrary.Float(56657.2238),
			PropRpm:               unifieddatalibrary.Float(8.2),
			PropType:              unifieddatalibrary.String("Diesel"),
			Pw:                    unifieddatalibrary.Float(1347.45),
			Reduced:               unifieddatalibrary.Bool(false),
			Reinforced:            unifieddatalibrary.Bool(false),
			RptArchived:           unifieddatalibrary.Bool(false),
			RptChxref:             unifieddatalibrary.String("INT"),
			Rtn:                   []string{"ex-a"},
			RtnCmd:                unifieddatalibrary.String("YORKTOWN"),
			RtnMsgTs:              []time.Time{time.Now()},
			RtnTrkState:           unifieddatalibrary.String("Local_RTN"),
			ScanRate:              unifieddatalibrary.Float(12.01),
			ScanType:              unifieddatalibrary.String("UNK"),
			Scn:                   unifieddatalibrary.Int(1474305),
			Sconum:                unifieddatalibrary.String("B45524"),
			SelfReport:            unifieddatalibrary.Bool(false),
			Sen:                   unifieddatalibrary.String("OTH"),
			ShipClass:             unifieddatalibrary.String("Nimitz"),
			ShortName:             unifieddatalibrary.String("COMMSCHECK"),
			SourceUid:             unifieddatalibrary.String("MCS"),
			SpaceAmp:              unifieddatalibrary.String("NUCLEAR WARHEAD"),
			SpaceAmpConf:          unifieddatalibrary.Int(6),
			SpaceSpecType:         unifieddatalibrary.String("SS-21_MOD_2_CRBM"),
			Spd:                   unifieddatalibrary.Float(15.03443),
			StaffCmts:             unifieddatalibrary.String("Staff Comments"),
			SternType:             unifieddatalibrary.String("Cruiser"),
			TaskForce:             unifieddatalibrary.Bool(false),
			Tcpa:                  unifieddatalibrary.Time(time.Now()),
			TesEventID:            unifieddatalibrary.String("6217"),
			Tol:                   unifieddatalibrary.Float(4.1),
			Tpk:                   unifieddatalibrary.Float(2.65),
			TrkConf:               unifieddatalibrary.Float(0.95),
			TrkID:                 unifieddatalibrary.String("trkId"),
			TrkNum:                unifieddatalibrary.String("trkNum"),
			TrkQual:               unifieddatalibrary.Int(1),
			TrkScope:              unifieddatalibrary.String("OTH"),
			TrnspdrID:             unifieddatalibrary.String("11"),
			TrnspdrType:           unifieddatalibrary.String("AFTS"),
			VslWt:                 unifieddatalibrary.Float(3423.76),
			Width:                 unifieddatalibrary.Float(24.1),
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

func TestTrackDetailQueryhelp(t *testing.T) {
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
	err := client.TrackDetails.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackDetailTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackDetails.Tuple(context.TODO(), unifieddatalibrary.TrackDetailTupleParams{
		Columns:     "columns",
		Ts:          time.Now(),
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
