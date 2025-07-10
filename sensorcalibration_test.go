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

func TestSensorCalibrationNewWithOptionalParams(t *testing.T) {
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
	err := client.Sensor.Calibration.New(context.TODO(), unifieddatalibrary.SensorCalibrationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SensorCalibrationNewParamsDataModeTest,
		IDSensor:              "09f2c68c-5e24-4b72-9cc8-ba9b1efa82f0",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AzRaAccelBias:         unifieddatalibrary.Float(0.0123),
		AzRaAccelSigma:        unifieddatalibrary.Float(0.0123),
		AzRaBias:              unifieddatalibrary.Float(0.327883),
		AzRaRateBias:          unifieddatalibrary.Float(0.123),
		AzRaRateSigma:         unifieddatalibrary.Float(0.123),
		AzRaRms:               unifieddatalibrary.Float(0.605333),
		AzRaSigma:             unifieddatalibrary.Float(0.000381),
		CalAngleRef:           unifieddatalibrary.String("AZEL"),
		CalTrackMode:          unifieddatalibrary.String("INTRA_TRACK"),
		CalType:               unifieddatalibrary.String("OPERATIONAL"),
		ConfidenceNoiseBias:   unifieddatalibrary.Float(0.001477),
		Duration:              unifieddatalibrary.Float(14.125),
		Ecr:                   []float64{352815.1, -5852915.3, 3255189},
		ElDecAccelBias:        unifieddatalibrary.Float(0.0123),
		ElDecAccelSigma:       unifieddatalibrary.Float(0.0123),
		ElDecBias:             unifieddatalibrary.Float(0.012555),
		ElDecRateBias:         unifieddatalibrary.Float(0.123),
		ElDecRateSigma:        unifieddatalibrary.Float(0.123),
		ElDecRms:              unifieddatalibrary.Float(0.080505),
		ElDecSigma:            unifieddatalibrary.Float(0.00265),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		NumAzRaObs:            unifieddatalibrary.Int(339),
		NumElDecObs:           unifieddatalibrary.Int(339),
		NumObs:                unifieddatalibrary.Int(341),
		NumPhotoObs:           unifieddatalibrary.Int(77),
		NumRangeObs:           unifieddatalibrary.Int(341),
		NumRangeRateObs:       unifieddatalibrary.Int(341),
		NumRcsObs:             unifieddatalibrary.Int(325),
		NumTimeObs:            unifieddatalibrary.Int(307),
		NumTracks:             unifieddatalibrary.Int(85),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PhotoBias:             unifieddatalibrary.Float(0.123),
		PhotoSigma:            unifieddatalibrary.Float(0.0123),
		RangeAccelBias:        unifieddatalibrary.Float(0.123),
		RangeAccelSigma:       unifieddatalibrary.Float(0.0123),
		RangeBias:             unifieddatalibrary.Float(0.024777),
		RangeRateBias:         unifieddatalibrary.Float(0.105333),
		RangeRateRms:          unifieddatalibrary.Float(0.000227),
		RangeRateSigma:        unifieddatalibrary.Float(0.000321),
		RangeRms:              unifieddatalibrary.Float(0.0123),
		RangeSigma:            unifieddatalibrary.Float(0.042644),
		RcsBias:               unifieddatalibrary.Float(0.123),
		RcsSigma:              unifieddatalibrary.Float(0.0123),
		RefTargets:            []string{"xx", "yy", "zz"},
		RefType:               unifieddatalibrary.String("SLR"),
		SenType:               unifieddatalibrary.String("PHASED ARRAY"),
		TimeBias:              unifieddatalibrary.Float(0.000372),
		TimeBiasSigma:         unifieddatalibrary.Float(15.333212),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorCalibrationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Calibration.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorCalibrationGetParams{
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

func TestSensorCalibrationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Calibration.Count(context.TODO(), unifieddatalibrary.SensorCalibrationCountParams{
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

func TestSensorCalibrationNewBulk(t *testing.T) {
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
	err := client.Sensor.Calibration.NewBulk(context.TODO(), unifieddatalibrary.SensorCalibrationNewBulkParams{
		Body: []unifieddatalibrary.SensorCalibrationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSensor:              "09f2c68c-5e24-4b72-9cc8-ba9b1efa82f0",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AzRaAccelBias:         unifieddatalibrary.Float(0.0123),
			AzRaAccelSigma:        unifieddatalibrary.Float(0.0123),
			AzRaBias:              unifieddatalibrary.Float(0.327883),
			AzRaRateBias:          unifieddatalibrary.Float(0.123),
			AzRaRateSigma:         unifieddatalibrary.Float(0.123),
			AzRaRms:               unifieddatalibrary.Float(0.605333),
			AzRaSigma:             unifieddatalibrary.Float(0.000381),
			CalAngleRef:           unifieddatalibrary.String("AZEL"),
			CalTrackMode:          unifieddatalibrary.String("INTRA_TRACK"),
			CalType:               unifieddatalibrary.String("OPERATIONAL"),
			ConfidenceNoiseBias:   unifieddatalibrary.Float(0.001477),
			Duration:              unifieddatalibrary.Float(14.125),
			Ecr:                   []float64{352815.1, -5852915.3, 3255189},
			ElDecAccelBias:        unifieddatalibrary.Float(0.0123),
			ElDecAccelSigma:       unifieddatalibrary.Float(0.0123),
			ElDecBias:             unifieddatalibrary.Float(0.012555),
			ElDecRateBias:         unifieddatalibrary.Float(0.123),
			ElDecRateSigma:        unifieddatalibrary.Float(0.123),
			ElDecRms:              unifieddatalibrary.Float(0.080505),
			ElDecSigma:            unifieddatalibrary.Float(0.00265),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			NumAzRaObs:            unifieddatalibrary.Int(339),
			NumElDecObs:           unifieddatalibrary.Int(339),
			NumObs:                unifieddatalibrary.Int(341),
			NumPhotoObs:           unifieddatalibrary.Int(77),
			NumRangeObs:           unifieddatalibrary.Int(341),
			NumRangeRateObs:       unifieddatalibrary.Int(341),
			NumRcsObs:             unifieddatalibrary.Int(325),
			NumTimeObs:            unifieddatalibrary.Int(307),
			NumTracks:             unifieddatalibrary.Int(85),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PhotoBias:             unifieddatalibrary.Float(0.123),
			PhotoSigma:            unifieddatalibrary.Float(0.0123),
			RangeAccelBias:        unifieddatalibrary.Float(0.123),
			RangeAccelSigma:       unifieddatalibrary.Float(0.0123),
			RangeBias:             unifieddatalibrary.Float(0.024777),
			RangeRateBias:         unifieddatalibrary.Float(0.105333),
			RangeRateRms:          unifieddatalibrary.Float(0.000227),
			RangeRateSigma:        unifieddatalibrary.Float(0.000321),
			RangeRms:              unifieddatalibrary.Float(0.0123),
			RangeSigma:            unifieddatalibrary.Float(0.042644),
			RcsBias:               unifieddatalibrary.Float(0.123),
			RcsSigma:              unifieddatalibrary.Float(0.0123),
			RefTargets:            []string{"xx", "yy", "zz"},
			RefType:               unifieddatalibrary.String("SLR"),
			SenType:               unifieddatalibrary.String("PHASED ARRAY"),
			TimeBias:              unifieddatalibrary.Float(0.000372),
			TimeBiasSigma:         unifieddatalibrary.Float(15.333212),
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

func TestSensorCalibrationQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Calibration.Query(context.TODO(), unifieddatalibrary.SensorCalibrationQueryParams{
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

func TestSensorCalibrationQueryHelp(t *testing.T) {
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
	_, err := client.Sensor.Calibration.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorCalibrationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Calibration.Tuple(context.TODO(), unifieddatalibrary.SensorCalibrationTupleParams{
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

func TestSensorCalibrationUnvalidatedPublish(t *testing.T) {
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
	err := client.Sensor.Calibration.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SensorCalibrationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SensorCalibrationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSensor:              "09f2c68c-5e24-4b72-9cc8-ba9b1efa82f0",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AzRaAccelBias:         unifieddatalibrary.Float(0.0123),
			AzRaAccelSigma:        unifieddatalibrary.Float(0.0123),
			AzRaBias:              unifieddatalibrary.Float(0.327883),
			AzRaRateBias:          unifieddatalibrary.Float(0.123),
			AzRaRateSigma:         unifieddatalibrary.Float(0.123),
			AzRaRms:               unifieddatalibrary.Float(0.605333),
			AzRaSigma:             unifieddatalibrary.Float(0.000381),
			CalAngleRef:           unifieddatalibrary.String("AZEL"),
			CalTrackMode:          unifieddatalibrary.String("INTRA_TRACK"),
			CalType:               unifieddatalibrary.String("OPERATIONAL"),
			ConfidenceNoiseBias:   unifieddatalibrary.Float(0.001477),
			Duration:              unifieddatalibrary.Float(14.125),
			Ecr:                   []float64{352815.1, -5852915.3, 3255189},
			ElDecAccelBias:        unifieddatalibrary.Float(0.0123),
			ElDecAccelSigma:       unifieddatalibrary.Float(0.0123),
			ElDecBias:             unifieddatalibrary.Float(0.012555),
			ElDecRateBias:         unifieddatalibrary.Float(0.123),
			ElDecRateSigma:        unifieddatalibrary.Float(0.123),
			ElDecRms:              unifieddatalibrary.Float(0.080505),
			ElDecSigma:            unifieddatalibrary.Float(0.00265),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			NumAzRaObs:            unifieddatalibrary.Int(339),
			NumElDecObs:           unifieddatalibrary.Int(339),
			NumObs:                unifieddatalibrary.Int(341),
			NumPhotoObs:           unifieddatalibrary.Int(77),
			NumRangeObs:           unifieddatalibrary.Int(341),
			NumRangeRateObs:       unifieddatalibrary.Int(341),
			NumRcsObs:             unifieddatalibrary.Int(325),
			NumTimeObs:            unifieddatalibrary.Int(307),
			NumTracks:             unifieddatalibrary.Int(85),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PhotoBias:             unifieddatalibrary.Float(0.123),
			PhotoSigma:            unifieddatalibrary.Float(0.0123),
			RangeAccelBias:        unifieddatalibrary.Float(0.123),
			RangeAccelSigma:       unifieddatalibrary.Float(0.0123),
			RangeBias:             unifieddatalibrary.Float(0.024777),
			RangeRateBias:         unifieddatalibrary.Float(0.105333),
			RangeRateRms:          unifieddatalibrary.Float(0.000227),
			RangeRateSigma:        unifieddatalibrary.Float(0.000321),
			RangeRms:              unifieddatalibrary.Float(0.0123),
			RangeSigma:            unifieddatalibrary.Float(0.042644),
			RcsBias:               unifieddatalibrary.Float(0.123),
			RcsSigma:              unifieddatalibrary.Float(0.0123),
			RefTargets:            []string{"xx", "yy", "zz"},
			RefType:               unifieddatalibrary.String("SLR"),
			SenType:               unifieddatalibrary.String("PHASED ARRAY"),
			TimeBias:              unifieddatalibrary.Float(0.000372),
			TimeBiasSigma:         unifieddatalibrary.Float(15.333212),
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
