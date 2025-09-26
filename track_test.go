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

func TestTrackListWithOptionalParams(t *testing.T) {
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
	_, err := client.Track.List(context.TODO(), unifieddatalibrary.TrackListParams{
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

func TestTrackCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Track.Count(context.TODO(), unifieddatalibrary.TrackCountParams{
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

func TestTrackNewBulk(t *testing.T) {
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
	err := client.Track.NewBulk(context.TODO(), unifieddatalibrary.TrackNewBulkParams{
		Body: []unifieddatalibrary.TrackNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("TRACK-ID"),
			Alt:                   unifieddatalibrary.Float(2024.15),
			Asset:                 unifieddatalibrary.String("asset"),
			AssetNat:              unifieddatalibrary.String("US"),
			Attitude:              []float64{10.1, 0.1, 1.1},
			AttitudeRate:          []float64{0.0003, 1e-7, 0.00003},
			CallSign:              unifieddatalibrary.String("callSign"),
			Cntct:                 unifieddatalibrary.String("Contact Info"),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
			Course:                unifieddatalibrary.Float(215.2),
			Cov:                   []float64{1.1, 2.2, 3.3},
			EcefAcc:               []float64{-0.0265, -0.2764, -0.1773},
			EcefPos:               []float64{-1273011.47, -48108810.77, 3979360.49},
			EcefVel:               []float64{25.635, -1.097, 57.373},
			ENuAcc:                []float64{0.0003, 0.014, 0.0003},
			ENuGroundVel:          []float64{1.23, 2.34, 3.45},
			ENuPos:                []float64{1.23, 2.34, 3.45},
			ENuVel:                []float64{1.23, 2.34, 3.45},
			Env:                   unifieddatalibrary.String("LAND"),
			EnvConf:               unifieddatalibrary.Float(1.23),
			ErrEllp:               []float64{1.23, 2.34, 3.45},
			Hdng:                  unifieddatalibrary.Float(215.7),
			IdentAmp:              unifieddatalibrary.String("ZOMBIE"),
			IdentCred:             unifieddatalibrary.Int(0),
			IdentRel:              unifieddatalibrary.Int(0),
			JSeries:               unifieddatalibrary.String("J12.5"),
			Lat:                   unifieddatalibrary.Float(38.8353),
			LcAcc:                 []float64{1.23, 2.34, 3.45},
			Lco:                   []float64{1.23, 2.34, 3.45},
			LcPos:                 []float64{1.23, 2.34, 3.45},
			Lcs:                   []float64{1.23, 2.34, 3.45},
			LcVel:                 []float64{1.23, 2.34, 3.45},
			Lon:                   unifieddatalibrary.Float(-104.8216),
			M1:                    unifieddatalibrary.Int(11),
			M1v:                   unifieddatalibrary.Int(1),
			M2:                    unifieddatalibrary.Int(1234),
			M2v:                   unifieddatalibrary.Int(1),
			M3a:                   unifieddatalibrary.Int(2636),
			M3av:                  unifieddatalibrary.Int(1),
			ModType:               unifieddatalibrary.String("MASINT"),
			MsgTs:                 unifieddatalibrary.Time(time.Now()),
			MsnID:                 unifieddatalibrary.String("msnId"),
			MultiSource:           unifieddatalibrary.Bool(true),
			ObjAct:                unifieddatalibrary.String("HOLDING"),
			ObjDescription:        unifieddatalibrary.String("Object description text."),
			ObjID:                 unifieddatalibrary.String("objId"),
			ObjIdent:              unifieddatalibrary.String("FRIEND"),
			ObjNat:                unifieddatalibrary.String("NATO"),
			ObjPlat:               unifieddatalibrary.String("COMBAT_VEHICLE"),
			ObjSpec:               unifieddatalibrary.String("LIGHT_TANK"),
			ObjType:               unifieddatalibrary.String("WATERCRAFT"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Sen:                   unifieddatalibrary.String("sen"),
			SenQual:               unifieddatalibrary.String("senQual"),
			Spd:                   unifieddatalibrary.Float(62.9921),
			SrcIDs:                []string{"f7c70cc8-f9b7-4467-b4ad-3904e360e842", "1da3fab000014e3133709830937387405"},
			SrcTyps:               []string{"MTI", "POI"},
			Strength:              unifieddatalibrary.Int(14),
			Tags:                  []string{"TAG1", "TAG2"},
			TrkConf:               unifieddatalibrary.Float(0.67),
			TrkID:                 unifieddatalibrary.String("trkId"),
			TrkItmID:              unifieddatalibrary.String("38f3a71f-4bba-4d58-9765-c1b93212aff1"),
			TrkNum:                unifieddatalibrary.String("trkNum"),
			TrkPtType:             unifieddatalibrary.String("MEASURED"),
			TrkQual:               unifieddatalibrary.Int(0),
			TrkStat:               unifieddatalibrary.String("INITIATING"),
			VertUnc:               unifieddatalibrary.Float(4.56),
			WanderAng:             unifieddatalibrary.Float(1.23),
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

func TestTrackQueryhelp(t *testing.T) {
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
	_, err := client.Track.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Track.Tuple(context.TODO(), unifieddatalibrary.TrackTupleParams{
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

func TestTrackUnvalidatedPublish(t *testing.T) {
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
	err := client.Track.UnvalidatedPublish(context.TODO(), unifieddatalibrary.TrackUnvalidatedPublishParams{
		Body: []unifieddatalibrary.TrackUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("TRACK-ID"),
			Alt:                   unifieddatalibrary.Float(2024.15),
			Asset:                 unifieddatalibrary.String("asset"),
			AssetNat:              unifieddatalibrary.String("US"),
			Attitude:              []float64{10.1, 0.1, 1.1},
			AttitudeRate:          []float64{0.0003, 1e-7, 0.00003},
			CallSign:              unifieddatalibrary.String("callSign"),
			Cntct:                 unifieddatalibrary.String("Contact Info"),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer", "velTs", "accelTs"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b", "2024-06-07T14:17:39.234Z", "2024-06-07T14:17:39.123Z"},
			Course:                unifieddatalibrary.Float(215.2),
			Cov:                   []float64{1.1, 2.2, 3.3},
			EcefAcc:               []float64{-0.0265, -0.2764, -0.1773},
			EcefPos:               []float64{-1273011.47, -48108810.77, 3979360.49},
			EcefVel:               []float64{25.635, -1.097, 57.373},
			ENuAcc:                []float64{0.0003, 0.014, 0.0003},
			ENuGroundVel:          []float64{1.23, 2.34, 3.45},
			ENuPos:                []float64{1.23, 2.34, 3.45},
			ENuVel:                []float64{1.23, 2.34, 3.45},
			Env:                   unifieddatalibrary.String("LAND"),
			EnvConf:               unifieddatalibrary.Float(1.23),
			ErrEllp:               []float64{1.23, 2.34, 3.45},
			Hdng:                  unifieddatalibrary.Float(215.7),
			IdentAmp:              unifieddatalibrary.String("ZOMBIE"),
			IdentCred:             unifieddatalibrary.Int(0),
			IdentRel:              unifieddatalibrary.Int(0),
			JSeries:               unifieddatalibrary.String("J12.5"),
			Lat:                   unifieddatalibrary.Float(38.8353),
			LcAcc:                 []float64{1.23, 2.34, 3.45},
			Lco:                   []float64{1.23, 2.34, 3.45},
			LcPos:                 []float64{1.23, 2.34, 3.45},
			Lcs:                   []float64{1.23, 2.34, 3.45},
			LcVel:                 []float64{1.23, 2.34, 3.45},
			Lon:                   unifieddatalibrary.Float(-104.8216),
			M1:                    unifieddatalibrary.Int(11),
			M1v:                   unifieddatalibrary.Int(1),
			M2:                    unifieddatalibrary.Int(1234),
			M2v:                   unifieddatalibrary.Int(1),
			M3a:                   unifieddatalibrary.Int(2636),
			M3av:                  unifieddatalibrary.Int(1),
			ModType:               unifieddatalibrary.String("MASINT"),
			MsgTs:                 unifieddatalibrary.Time(time.Now()),
			MsnID:                 unifieddatalibrary.String("msnId"),
			MultiSource:           unifieddatalibrary.Bool(true),
			ObjAct:                unifieddatalibrary.String("HOLDING"),
			ObjDescription:        unifieddatalibrary.String("Object description text."),
			ObjID:                 unifieddatalibrary.String("objId"),
			ObjIdent:              unifieddatalibrary.String("FRIEND"),
			ObjNat:                unifieddatalibrary.String("NATO"),
			ObjPlat:               unifieddatalibrary.String("COMBAT_VEHICLE"),
			ObjSpec:               unifieddatalibrary.String("LIGHT_TANK"),
			ObjType:               unifieddatalibrary.String("WATERCRAFT"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Sen:                   unifieddatalibrary.String("sen"),
			SenQual:               unifieddatalibrary.String("senQual"),
			Spd:                   unifieddatalibrary.Float(62.9921),
			SrcIDs:                []string{"f7c70cc8-f9b7-4467-b4ad-3904e360e842", "1da3fab000014e3133709830937387405"},
			SrcTyps:               []string{"MTI", "POI"},
			Strength:              unifieddatalibrary.Int(14),
			Tags:                  []string{"TAG1", "TAG2"},
			TrkConf:               unifieddatalibrary.Float(0.67),
			TrkID:                 unifieddatalibrary.String("trkId"),
			TrkItmID:              unifieddatalibrary.String("38f3a71f-4bba-4d58-9765-c1b93212aff1"),
			TrkNum:                unifieddatalibrary.String("trkNum"),
			TrkPtType:             unifieddatalibrary.String("MEASURED"),
			TrkQual:               unifieddatalibrary.Int(0),
			TrkStat:               unifieddatalibrary.String("INITIATING"),
			VertUnc:               unifieddatalibrary.Float(4.56),
			WanderAng:             unifieddatalibrary.Float(1.23),
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
