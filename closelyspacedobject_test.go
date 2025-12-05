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

func TestCloselyspacedobjectNewWithOptionalParams(t *testing.T) {
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
	err := client.Closelyspacedobjects.New(context.TODO(), unifieddatalibrary.CloselyspacedobjectNewParams{
		ClassificationMarking: "U",
		CsoState:              "POSSIBLE",
		DataMode:              unifieddatalibrary.CloselyspacedobjectNewParamsDataModeTest,
		EventStartTime:        time.Now(),
		EventType:             "RENDEZVOUS",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		ActorSvEpoch:          unifieddatalibrary.Time(time.Now()),
		AnalysisDuration:      unifieddatalibrary.Float(60.1),
		AnalysisEpoch:         unifieddatalibrary.Time(time.Now()),
		CompType:              unifieddatalibrary.String("LONGITUDE"),
		ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer"},
		ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b"},
		CsoDetails: []unifieddatalibrary.CloselyspacedobjectNewParamsCsoDetail{{
			ObjectEvent:           "MEAN",
			ObjectType:            "DELTA",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Apogee:                unifieddatalibrary.Float(1.1),
			ClassificationMarking: unifieddatalibrary.String("U"),
			DataMode:              "TEST",
			IDCso:                 unifieddatalibrary.String("CSO-ID"),
			Inclination:           unifieddatalibrary.Float(45.1),
			Longitude:             unifieddatalibrary.Float(45.1),
			Perigee:               unifieddatalibrary.Float(1.1),
		}},
		DeltaVTol:                unifieddatalibrary.Float(0.123),
		DurationThreshold:        unifieddatalibrary.Float(60.1),
		EventEndTime:             unifieddatalibrary.Time(time.Now()),
		EventIntervalCoverage:    unifieddatalibrary.Float(22.3),
		ExtID:                    unifieddatalibrary.String("EXTERNAL-ID"),
		HohmannDeltaV:            unifieddatalibrary.Float(0.012),
		IDActorSv:                unifieddatalibrary.String("ACTOR-SV-ID"),
		IDOnOrbit1:               unifieddatalibrary.String("ONORBIT1-ID"),
		IDOnOrbit2:               unifieddatalibrary.String("ONORBIT2-ID"),
		IDTargetSv:               unifieddatalibrary.String("TARGET-SV-ID"),
		InclinationDeltaV:        unifieddatalibrary.Float(0.012),
		IndicationSource:         unifieddatalibrary.String("Manually input"),
		LonTol:                   unifieddatalibrary.Float(30.1),
		MaxRange:                 unifieddatalibrary.Float(233.266),
		MinPlaneSepAngle:         unifieddatalibrary.Float(30.1),
		MinPlaneSepEpoch:         unifieddatalibrary.Time(time.Now()),
		MinRange:                 unifieddatalibrary.Float(0.5),
		MinRangeAnalysisDuration: unifieddatalibrary.Float(60.1),
		MinRangeEpoch:            unifieddatalibrary.Time(time.Now()),
		Notes:                    unifieddatalibrary.String("FREE-TEXT"),
		NumSubIntervals:          unifieddatalibrary.Int(0),
		OrbitAlignDel:            unifieddatalibrary.Float(12.3),
		OrbitPlaneTol:            unifieddatalibrary.Float(1.23),
		Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectId1:            unifieddatalibrary.String("ORIGONORBIT1-ID"),
		OrigObjectId2:            unifieddatalibrary.String("ORIGONORBIT2-ID"),
		RangeThreshold:           unifieddatalibrary.Float(0.1),
		RangeTol:                 unifieddatalibrary.Float(0.123),
		RelPos:                   []float64{0.12, 0.23, -0.12},
		RelPosMag:                unifieddatalibrary.Float(0.12),
		RelSpeedMag:              unifieddatalibrary.Float(1.23),
		RelVel:                   []float64{0.12, 0.23, -0.12},
		SatNo1:                   unifieddatalibrary.Int(1),
		SatNo2:                   unifieddatalibrary.Int(2),
		StationLimLonTol:         unifieddatalibrary.Float(12.5),
		TargetSvEpoch:            unifieddatalibrary.Time(time.Now()),
		TotalDeltaV:              unifieddatalibrary.Float(2.46),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloselyspacedobjectGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Closelyspacedobjects.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.CloselyspacedobjectGetParams{
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

func TestCloselyspacedobjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Closelyspacedobjects.List(context.TODO(), unifieddatalibrary.CloselyspacedobjectListParams{
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloselyspacedobjectCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Closelyspacedobjects.Count(context.TODO(), unifieddatalibrary.CloselyspacedobjectCountParams{
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloselyspacedobjectNewBulk(t *testing.T) {
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
	err := client.Closelyspacedobjects.NewBulk(context.TODO(), unifieddatalibrary.CloselyspacedobjectNewBulkParams{
		Body: []unifieddatalibrary.CloselyspacedobjectNewBulkParamsBody{{
			ClassificationMarking: "U",
			CsoState:              "POSSIBLE",
			DataMode:              "TEST",
			EventStartTime:        time.Now(),
			EventType:             "RENDEZVOUS",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActorSvEpoch:          unifieddatalibrary.Time(time.Now()),
			AnalysisDuration:      unifieddatalibrary.Float(60.1),
			AnalysisEpoch:         unifieddatalibrary.Time(time.Now()),
			CompType:              unifieddatalibrary.String("LONGITUDE"),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b"},
			CsoDetails: []unifieddatalibrary.CloselyspacedobjectNewBulkParamsBodyCsoDetail{{
				ObjectEvent:           "MEAN",
				ObjectType:            "DELTA",
				ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
				Apogee:                unifieddatalibrary.Float(1.1),
				ClassificationMarking: unifieddatalibrary.String("U"),
				DataMode:              "TEST",
				IDCso:                 unifieddatalibrary.String("CSO-ID"),
				Inclination:           unifieddatalibrary.Float(45.1),
				Longitude:             unifieddatalibrary.Float(45.1),
				Perigee:               unifieddatalibrary.Float(1.1),
			}},
			DeltaVTol:                unifieddatalibrary.Float(0.123),
			DurationThreshold:        unifieddatalibrary.Float(60.1),
			EventEndTime:             unifieddatalibrary.Time(time.Now()),
			EventIntervalCoverage:    unifieddatalibrary.Float(22.3),
			ExtID:                    unifieddatalibrary.String("EXTERNAL-ID"),
			HohmannDeltaV:            unifieddatalibrary.Float(0.012),
			IDActorSv:                unifieddatalibrary.String("ACTOR-SV-ID"),
			IDOnOrbit1:               unifieddatalibrary.String("ONORBIT1-ID"),
			IDOnOrbit2:               unifieddatalibrary.String("ONORBIT2-ID"),
			IDTargetSv:               unifieddatalibrary.String("TARGET-SV-ID"),
			InclinationDeltaV:        unifieddatalibrary.Float(0.012),
			IndicationSource:         unifieddatalibrary.String("Manually input"),
			LonTol:                   unifieddatalibrary.Float(30.1),
			MaxRange:                 unifieddatalibrary.Float(233.266),
			MinPlaneSepAngle:         unifieddatalibrary.Float(30.1),
			MinPlaneSepEpoch:         unifieddatalibrary.Time(time.Now()),
			MinRange:                 unifieddatalibrary.Float(0.5),
			MinRangeAnalysisDuration: unifieddatalibrary.Float(60.1),
			MinRangeEpoch:            unifieddatalibrary.Time(time.Now()),
			Notes:                    unifieddatalibrary.String("FREE-TEXT"),
			NumSubIntervals:          unifieddatalibrary.Int(0),
			OrbitAlignDel:            unifieddatalibrary.Float(12.3),
			OrbitPlaneTol:            unifieddatalibrary.Float(1.23),
			Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectId1:            unifieddatalibrary.String("ORIGONORBIT1-ID"),
			OrigObjectId2:            unifieddatalibrary.String("ORIGONORBIT2-ID"),
			RangeThreshold:           unifieddatalibrary.Float(0.1),
			RangeTol:                 unifieddatalibrary.Float(0.123),
			RelPos:                   []float64{0.12, 0.23, -0.12},
			RelPosMag:                unifieddatalibrary.Float(0.12),
			RelSpeedMag:              unifieddatalibrary.Float(1.23),
			RelVel:                   []float64{0.12, 0.23, -0.12},
			SatNo1:                   unifieddatalibrary.Int(1),
			SatNo2:                   unifieddatalibrary.Int(2),
			StationLimLonTol:         unifieddatalibrary.Float(12.5),
			TargetSvEpoch:            unifieddatalibrary.Time(time.Now()),
			TotalDeltaV:              unifieddatalibrary.Float(2.46),
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

func TestCloselyspacedobjectQueryHelp(t *testing.T) {
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
	_, err := client.Closelyspacedobjects.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloselyspacedobjectTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Closelyspacedobjects.Tuple(context.TODO(), unifieddatalibrary.CloselyspacedobjectTupleParams{
		Columns:        "columns",
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloselyspacedobjectUnvalidatedPublish(t *testing.T) {
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
	err := client.Closelyspacedobjects.UnvalidatedPublish(context.TODO(), unifieddatalibrary.CloselyspacedobjectUnvalidatedPublishParams{
		Body: []unifieddatalibrary.CloselyspacedobjectUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			CsoState:              "POSSIBLE",
			DataMode:              "TEST",
			EventStartTime:        time.Now(),
			EventType:             "RENDEZVOUS",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActorSvEpoch:          unifieddatalibrary.Time(time.Now()),
			AnalysisDuration:      unifieddatalibrary.Float(60.1),
			AnalysisEpoch:         unifieddatalibrary.Time(time.Now()),
			CompType:              unifieddatalibrary.String("LONGITUDE"),
			ContextKeys:           []string{"MsnID_DescLabel", "msnVer", "serVer"},
			ContextValues:         []string{"MissionID Descriptive Label text", "1", "001.9b"},
			CsoDetails: []unifieddatalibrary.CloselyspacedobjectUnvalidatedPublishParamsBodyCsoDetail{{
				ObjectEvent:           "MEAN",
				ObjectType:            "DELTA",
				ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
				Apogee:                unifieddatalibrary.Float(1.1),
				ClassificationMarking: unifieddatalibrary.String("U"),
				DataMode:              "TEST",
				IDCso:                 unifieddatalibrary.String("CSO-ID"),
				Inclination:           unifieddatalibrary.Float(45.1),
				Longitude:             unifieddatalibrary.Float(45.1),
				Perigee:               unifieddatalibrary.Float(1.1),
			}},
			DeltaVTol:                unifieddatalibrary.Float(0.123),
			DurationThreshold:        unifieddatalibrary.Float(60.1),
			EventEndTime:             unifieddatalibrary.Time(time.Now()),
			EventIntervalCoverage:    unifieddatalibrary.Float(22.3),
			ExtID:                    unifieddatalibrary.String("EXTERNAL-ID"),
			HohmannDeltaV:            unifieddatalibrary.Float(0.012),
			IDActorSv:                unifieddatalibrary.String("ACTOR-SV-ID"),
			IDOnOrbit1:               unifieddatalibrary.String("ONORBIT1-ID"),
			IDOnOrbit2:               unifieddatalibrary.String("ONORBIT2-ID"),
			IDTargetSv:               unifieddatalibrary.String("TARGET-SV-ID"),
			InclinationDeltaV:        unifieddatalibrary.Float(0.012),
			IndicationSource:         unifieddatalibrary.String("Manually input"),
			LonTol:                   unifieddatalibrary.Float(30.1),
			MaxRange:                 unifieddatalibrary.Float(233.266),
			MinPlaneSepAngle:         unifieddatalibrary.Float(30.1),
			MinPlaneSepEpoch:         unifieddatalibrary.Time(time.Now()),
			MinRange:                 unifieddatalibrary.Float(0.5),
			MinRangeAnalysisDuration: unifieddatalibrary.Float(60.1),
			MinRangeEpoch:            unifieddatalibrary.Time(time.Now()),
			Notes:                    unifieddatalibrary.String("FREE-TEXT"),
			NumSubIntervals:          unifieddatalibrary.Int(0),
			OrbitAlignDel:            unifieddatalibrary.Float(12.3),
			OrbitPlaneTol:            unifieddatalibrary.Float(1.23),
			Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectId1:            unifieddatalibrary.String("ORIGONORBIT1-ID"),
			OrigObjectId2:            unifieddatalibrary.String("ORIGONORBIT2-ID"),
			RangeThreshold:           unifieddatalibrary.Float(0.1),
			RangeTol:                 unifieddatalibrary.Float(0.123),
			RelPos:                   []float64{0.12, 0.23, -0.12},
			RelPosMag:                unifieddatalibrary.Float(0.12),
			RelSpeedMag:              unifieddatalibrary.Float(1.23),
			RelVel:                   []float64{0.12, 0.23, -0.12},
			SatNo1:                   unifieddatalibrary.Int(1),
			SatNo2:                   unifieddatalibrary.Int(2),
			StationLimLonTol:         unifieddatalibrary.Float(12.5),
			TargetSvEpoch:            unifieddatalibrary.Time(time.Now()),
			TotalDeltaV:              unifieddatalibrary.Float(2.46),
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
