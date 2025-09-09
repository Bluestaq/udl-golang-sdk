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

func TestMissileTrackListWithOptionalParams(t *testing.T) {
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
	_, err := client.MissileTracks.List(context.TODO(), unifieddatalibrary.MissileTrackListParams{
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

func TestMissileTrackCountWithOptionalParams(t *testing.T) {
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
	_, err := client.MissileTracks.Count(context.TODO(), unifieddatalibrary.MissileTrackCountParams{
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

func TestMissileTrackNewBulk(t *testing.T) {
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
	err := client.MissileTracks.NewBulk(context.TODO(), unifieddatalibrary.MissileTrackNewBulkParams{
		Body: []unifieddatalibrary.MissileTrackNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("MissileTrack_ID"),
			AcftSubType:           unifieddatalibrary.String("SLBM"),
			Alert:                 unifieddatalibrary.String("HIT"),
			AngElev:               unifieddatalibrary.Float(15.2),
			AouRptData:            []float64{34.3, 26.5, 1.2},
			AouRptType:            unifieddatalibrary.String("ELLIPSE"),
			AzCorr:                unifieddatalibrary.Float(12.876),
			Boosting:              unifieddatalibrary.Bool(true),
			BurnoutAlt:            unifieddatalibrary.Float(30567.452),
			CallSign:              unifieddatalibrary.String("Charlie"),
			Containment:           unifieddatalibrary.Float(90.64),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
			DropPtInd:             unifieddatalibrary.Bool(true),
			EmgInd:                unifieddatalibrary.Bool(true),
			Env:                   "AIR",
			ImpactAlt:             unifieddatalibrary.Float(0.02),
			ImpactAouData:         []float64{34.3, 26.5, 1.2},
			ImpactAouType:         unifieddatalibrary.String("ELLIPSE"),
			ImpactConf:            unifieddatalibrary.Float(99.9),
			ImpactLat:             unifieddatalibrary.Float(19.88550102),
			ImpactLon:             unifieddatalibrary.Float(46.74596844),
			ImpactTime:            unifieddatalibrary.Time(time.Now()),
			InfoSource:            unifieddatalibrary.String("S1"),
			LaunchAlt:             unifieddatalibrary.Float(0.01),
			LaunchAouData:         []float64{1.23, 2.34, 3.45},
			LaunchAouType:         unifieddatalibrary.String("ELLIPSE"),
			LaunchAz:              unifieddatalibrary.Float(99.1),
			LaunchAzUnc:           unifieddatalibrary.Float(2.4),
			LaunchConf:            unifieddatalibrary.Float(90.7),
			LaunchLat:             unifieddatalibrary.Float(19.88550102),
			LaunchLon:             unifieddatalibrary.Float(46.74596844),
			LaunchTime:            unifieddatalibrary.Time(time.Now()),
			LostTrkInd:            unifieddatalibrary.Bool(false),
			ManeuverEnd:           unifieddatalibrary.Time(time.Now()),
			ManeuverStart:         unifieddatalibrary.Time(time.Now()),
			MsgCreateDate:         unifieddatalibrary.Time(time.Now()),
			MsgSubType:            unifieddatalibrary.String("Update"),
			MsgType:               unifieddatalibrary.String("MSG-TYPE"),
			MslStatus:             unifieddatalibrary.String("AT LAUNCH"),
			MuidSrc:               unifieddatalibrary.String("MUID-SRC"),
			MuidSrcTrk:            unifieddatalibrary.String("MUID-SRC-ID"),
			Name:                  unifieddatalibrary.String("TRACK-NAME"),
			ObjAct:                unifieddatalibrary.String("HOLDING"),
			ObjIdent:              "FRIEND",
			ObjPlat:               unifieddatalibrary.String("WEAPON"),
			ObjType:               unifieddatalibrary.String("Ballistic"),
			ObjTypeConf:           unifieddatalibrary.Int(90),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ParentTrackID:         unifieddatalibrary.String("102288"),
			PolarSingLocLat:       unifieddatalibrary.Float(19.88550102),
			PolarSingLocLon:       unifieddatalibrary.Float(46.74596844),
			SenMode:               unifieddatalibrary.String("OBSBO"),
			SpaceAmp:              unifieddatalibrary.String("NUCLEAR WARHEAD"),
			SpaceAmpConf:          unifieddatalibrary.Int(6),
			SpaceSpecType:         unifieddatalibrary.String("SS-21_MOD_2_CRBM"),
			TrackID:               unifieddatalibrary.String("102288"),
			TrkConf:               unifieddatalibrary.Float(0.95),
			TrkQual:               unifieddatalibrary.Int(1),
			Vectors: []unifieddatalibrary.MissileTrackNewBulkParamsBodyVector{{
				Epoch:             time.Now(),
				Accel:             []float64{0.59236, -0.03537, 0.35675},
				Confidence:        unifieddatalibrary.Int(100),
				ContextKeys:       []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
				ContextValues:     []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
				Course:            unifieddatalibrary.Float(7.3580153),
				Cov:               []float64{1.1, 2.2, 3.3},
				CovReferenceFrame: "ECEF",
				FlightAz:          unifieddatalibrary.Float(45.23),
				IDSensor:          unifieddatalibrary.String("a7e99418-b6d6-29ab-e767-440a989cce26"),
				Object:            unifieddatalibrary.String("TARGET"),
				OrigSensorID:      unifieddatalibrary.String("ORIGSENSOR-ID"),
				Pos:               []float64{-1456.91592, -2883.54041, 6165.55186},
				Propagated:        unifieddatalibrary.Bool(false),
				Quat:              []float64{0.03, 0.02, 0.01, 0.012},
				Range:             unifieddatalibrary.Float(12.3),
				ReferenceFrame:    unifieddatalibrary.String("ECEF"),
				Spd:               unifieddatalibrary.Float(15.03443),
				Status:            unifieddatalibrary.String("INITIAL"),
				TimeSource:        unifieddatalibrary.String("Sensor 1"),
				Type:              unifieddatalibrary.String("STATE"),
				VectorAlt:         unifieddatalibrary.Float(25),
				VectorLat:         unifieddatalibrary.Float(45),
				VectorLon:         unifieddatalibrary.Float(150),
				VectorTrackID:     unifieddatalibrary.String("102288"),
				Vel:               []float64{-1.21981, -6.60208, -3.36515},
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

func TestMissileTrackQueryhelp(t *testing.T) {
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
	_, err := client.MissileTracks.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMissileTrackTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.MissileTracks.Tuple(context.TODO(), unifieddatalibrary.MissileTrackTupleParams{
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

func TestMissileTrackUnvalidatedPublish(t *testing.T) {
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
	err := client.MissileTracks.UnvalidatedPublish(context.TODO(), unifieddatalibrary.MissileTrackUnvalidatedPublishParams{
		Body: []unifieddatalibrary.MissileTrackUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("MissileTrack_ID"),
			AcftSubType:           unifieddatalibrary.String("SLBM"),
			Alert:                 unifieddatalibrary.String("HIT"),
			AngElev:               unifieddatalibrary.Float(15.2),
			AouRptData:            []float64{34.3, 26.5, 1.2},
			AouRptType:            unifieddatalibrary.String("ELLIPSE"),
			AzCorr:                unifieddatalibrary.Float(12.876),
			Boosting:              unifieddatalibrary.Bool(true),
			BurnoutAlt:            unifieddatalibrary.Float(30567.452),
			CallSign:              unifieddatalibrary.String("Charlie"),
			Containment:           unifieddatalibrary.Float(90.64),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
			DropPtInd:             unifieddatalibrary.Bool(true),
			EmgInd:                unifieddatalibrary.Bool(true),
			Env:                   "AIR",
			ImpactAlt:             unifieddatalibrary.Float(0.02),
			ImpactAouData:         []float64{34.3, 26.5, 1.2},
			ImpactAouType:         unifieddatalibrary.String("ELLIPSE"),
			ImpactConf:            unifieddatalibrary.Float(99.9),
			ImpactLat:             unifieddatalibrary.Float(19.88550102),
			ImpactLon:             unifieddatalibrary.Float(46.74596844),
			ImpactTime:            unifieddatalibrary.Time(time.Now()),
			InfoSource:            unifieddatalibrary.String("S1"),
			LaunchAlt:             unifieddatalibrary.Float(0.01),
			LaunchAouData:         []float64{1.23, 2.34, 3.45},
			LaunchAouType:         unifieddatalibrary.String("ELLIPSE"),
			LaunchAz:              unifieddatalibrary.Float(99.1),
			LaunchAzUnc:           unifieddatalibrary.Float(2.4),
			LaunchConf:            unifieddatalibrary.Float(90.7),
			LaunchLat:             unifieddatalibrary.Float(19.88550102),
			LaunchLon:             unifieddatalibrary.Float(46.74596844),
			LaunchTime:            unifieddatalibrary.Time(time.Now()),
			LostTrkInd:            unifieddatalibrary.Bool(false),
			ManeuverEnd:           unifieddatalibrary.Time(time.Now()),
			ManeuverStart:         unifieddatalibrary.Time(time.Now()),
			MsgCreateDate:         unifieddatalibrary.Time(time.Now()),
			MsgSubType:            unifieddatalibrary.String("Update"),
			MsgType:               unifieddatalibrary.String("MSG-TYPE"),
			MslStatus:             unifieddatalibrary.String("AT LAUNCH"),
			MuidSrc:               unifieddatalibrary.String("MUID-SRC"),
			MuidSrcTrk:            unifieddatalibrary.String("MUID-SRC-ID"),
			Name:                  unifieddatalibrary.String("TRACK-NAME"),
			ObjAct:                unifieddatalibrary.String("HOLDING"),
			ObjIdent:              "FRIEND",
			ObjPlat:               unifieddatalibrary.String("WEAPON"),
			ObjType:               unifieddatalibrary.String("Ballistic"),
			ObjTypeConf:           unifieddatalibrary.Int(90),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ParentTrackID:         unifieddatalibrary.String("102288"),
			PolarSingLocLat:       unifieddatalibrary.Float(19.88550102),
			PolarSingLocLon:       unifieddatalibrary.Float(46.74596844),
			SenMode:               unifieddatalibrary.String("OBSBO"),
			SpaceAmp:              unifieddatalibrary.String("NUCLEAR WARHEAD"),
			SpaceAmpConf:          unifieddatalibrary.Int(6),
			SpaceSpecType:         unifieddatalibrary.String("SS-21_MOD_2_CRBM"),
			TrackID:               unifieddatalibrary.String("102288"),
			TrkConf:               unifieddatalibrary.Float(0.95),
			TrkQual:               unifieddatalibrary.Int(1),
			Vectors: []unifieddatalibrary.MissileTrackUnvalidatedPublishParamsBodyVector{{
				Epoch:             time.Now(),
				Accel:             []float64{0.59236, -0.03537, 0.35675},
				Confidence:        unifieddatalibrary.Int(100),
				ContextKeys:       []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
				ContextValues:     []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
				Course:            unifieddatalibrary.Float(7.3580153),
				Cov:               []float64{1.1, 2.2, 3.3},
				CovReferenceFrame: "ECEF",
				FlightAz:          unifieddatalibrary.Float(45.23),
				IDSensor:          unifieddatalibrary.String("a7e99418-b6d6-29ab-e767-440a989cce26"),
				Object:            unifieddatalibrary.String("TARGET"),
				OrigSensorID:      unifieddatalibrary.String("ORIGSENSOR-ID"),
				Pos:               []float64{-1456.91592, -2883.54041, 6165.55186},
				Propagated:        unifieddatalibrary.Bool(false),
				Quat:              []float64{0.03, 0.02, 0.01, 0.012},
				Range:             unifieddatalibrary.Float(12.3),
				ReferenceFrame:    unifieddatalibrary.String("ECEF"),
				Spd:               unifieddatalibrary.Float(15.03443),
				Status:            unifieddatalibrary.String("INITIAL"),
				TimeSource:        unifieddatalibrary.String("Sensor 1"),
				Type:              unifieddatalibrary.String("STATE"),
				VectorAlt:         unifieddatalibrary.Float(25),
				VectorLat:         unifieddatalibrary.Float(45),
				VectorLon:         unifieddatalibrary.Float(150),
				VectorTrackID:     unifieddatalibrary.String("102288"),
				Vel:               []float64{-1.21981, -6.60208, -3.36515},
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
