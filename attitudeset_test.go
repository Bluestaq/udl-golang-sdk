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

func TestAttitudeSetNewWithOptionalParams(t *testing.T) {
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
	err := client.AttitudeSets.New(context.TODO(), unifieddatalibrary.AttitudeSetNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AttitudeSetNewParamsDataModeTest,
		EndTime:               time.Now(),
		Frame1:                "SCBODY",
		Frame2:                "J2000",
		NumPoints:             120,
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		Type:                  "AEM",
		ID:                    unifieddatalibrary.String("ATTITUDESET-ID"),
		AsRef:                 []string{"2ea97de6-4680-4767-a07e-35d16398ef60"},
		AttitudeList: []unifieddatalibrary.AttitudeSetNewParamsAttitudeList{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("ATTITUDEDATA-ID"),
			AsID:                  unifieddatalibrary.String("773c9887-e931-42eb-8155-f0fbd227b235"),
			ConingAngle:           unifieddatalibrary.Float(0.1),
			Declination:           unifieddatalibrary.Float(0.799),
			MotionType:            unifieddatalibrary.String("PROSOL_MOTION"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("D6593"),
			PrecPeriod:            unifieddatalibrary.Float(36.1),
			Q1:                    unifieddatalibrary.Float(0.0312),
			Q1Dot:                 unifieddatalibrary.Float(0.0043),
			Q2:                    unifieddatalibrary.Float(0.7854),
			Q2Dot:                 unifieddatalibrary.Float(0.06),
			Q3:                    unifieddatalibrary.Float(0.3916),
			Q3Dot:                 unifieddatalibrary.Float(0.499),
			Qc:                    unifieddatalibrary.Float(0.4783),
			QcDot:                 unifieddatalibrary.Float(0.011),
			Ra:                    unifieddatalibrary.Float(-173.75),
			SatNo:                 unifieddatalibrary.Int(41947),
			SpinPeriod:            unifieddatalibrary.Float(0.1),
			XAngle:                []float64{139.753},
			XRate:                 []float64{0.105},
			YAngle:                []float64{25.066},
			YRate:                 []float64{0.032},
			ZAngle:                []float64{-53.368},
			ZRate:                 []float64{0.022},
		}},
		EsID:               unifieddatalibrary.String("60f7a241-b7be-48d8-acf3-786670af53f9"),
		EulerRotSeq:        unifieddatalibrary.String("123"),
		IDSensor:           unifieddatalibrary.String("a7e99418-b6d6-29ab-e767-440a989cce26"),
		Interpolator:       unifieddatalibrary.String("LINEAR"),
		InterpolatorDegree: unifieddatalibrary.Int(2),
		Notes:              unifieddatalibrary.String("Notes for this attitude set"),
		Origin:             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:       unifieddatalibrary.String("D6593"),
		OrigSensorID:       unifieddatalibrary.String("ORIGSENSOR-ID"),
		PrecAngleInit:      unifieddatalibrary.Float(30.5),
		SatNo:              unifieddatalibrary.Int(41947),
		SpinAngleInit:      unifieddatalibrary.Float(25.5),
		StepSize:           unifieddatalibrary.Int(60),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttitudeSetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AttitudeSets.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AttitudeSetGetParams{
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

func TestAttitudeSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.AttitudeSets.List(context.TODO(), unifieddatalibrary.AttitudeSetListParams{
		StartTime:   time.Now(),
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

func TestAttitudeSetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AttitudeSets.Count(context.TODO(), unifieddatalibrary.AttitudeSetCountParams{
		StartTime:   time.Now(),
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

func TestAttitudeSetQueryHelp(t *testing.T) {
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
	_, err := client.AttitudeSets.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttitudeSetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AttitudeSets.Tuple(context.TODO(), unifieddatalibrary.AttitudeSetTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestAttitudeSetUnvalidatedPublishWithOptionalParams(t *testing.T) {
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
	err := client.AttitudeSets.UnvalidatedPublish(context.TODO(), unifieddatalibrary.AttitudeSetUnvalidatedPublishParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AttitudeSetUnvalidatedPublishParamsDataModeTest,
		EndTime:               time.Now(),
		Frame1:                "SCBODY",
		Frame2:                "J2000",
		NumPoints:             120,
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		Type:                  "AEM",
		ID:                    unifieddatalibrary.String("ATTITUDESET-ID"),
		AsRef:                 []string{"2ea97de6-4680-4767-a07e-35d16398ef60"},
		AttitudeList: []unifieddatalibrary.AttitudeSetUnvalidatedPublishParamsAttitudeList{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("ATTITUDEDATA-ID"),
			AsID:                  unifieddatalibrary.String("773c9887-e931-42eb-8155-f0fbd227b235"),
			ConingAngle:           unifieddatalibrary.Float(0.1),
			Declination:           unifieddatalibrary.Float(0.799),
			MotionType:            unifieddatalibrary.String("PROSOL_MOTION"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("D6593"),
			PrecPeriod:            unifieddatalibrary.Float(36.1),
			Q1:                    unifieddatalibrary.Float(0.0312),
			Q1Dot:                 unifieddatalibrary.Float(0.0043),
			Q2:                    unifieddatalibrary.Float(0.7854),
			Q2Dot:                 unifieddatalibrary.Float(0.06),
			Q3:                    unifieddatalibrary.Float(0.3916),
			Q3Dot:                 unifieddatalibrary.Float(0.499),
			Qc:                    unifieddatalibrary.Float(0.4783),
			QcDot:                 unifieddatalibrary.Float(0.011),
			Ra:                    unifieddatalibrary.Float(-173.75),
			SatNo:                 unifieddatalibrary.Int(41947),
			SpinPeriod:            unifieddatalibrary.Float(0.1),
			XAngle:                []float64{139.753},
			XRate:                 []float64{0.105},
			YAngle:                []float64{25.066},
			YRate:                 []float64{0.032},
			ZAngle:                []float64{-53.368},
			ZRate:                 []float64{0.022},
		}},
		EsID:               unifieddatalibrary.String("60f7a241-b7be-48d8-acf3-786670af53f9"),
		EulerRotSeq:        unifieddatalibrary.String("123"),
		IDSensor:           unifieddatalibrary.String("a7e99418-b6d6-29ab-e767-440a989cce26"),
		Interpolator:       unifieddatalibrary.String("LINEAR"),
		InterpolatorDegree: unifieddatalibrary.Int(2),
		Notes:              unifieddatalibrary.String("Notes for this attitude set"),
		Origin:             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:       unifieddatalibrary.String("D6593"),
		OrigSensorID:       unifieddatalibrary.String("ORIGSENSOR-ID"),
		PrecAngleInit:      unifieddatalibrary.Float(30.5),
		SatNo:              unifieddatalibrary.Int(41947),
		SpinAngleInit:      unifieddatalibrary.Float(25.5),
		StepSize:           unifieddatalibrary.Int(60),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
