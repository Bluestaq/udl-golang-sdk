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

func TestEcpedrNewWithOptionalParams(t *testing.T) {
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
	err := client.Ecpedr.New(context.TODO(), unifieddatalibrary.EcpedrNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EcpedrNewParamsDataModeTest,
		EcpedrMeasurements: []unifieddatalibrary.EcpedrNewParamsEcpedrMeasurement{{
			ObType:         "FLUX",
			ObUoM:          "#/MeV/cm^2/sr/s",
			ChanEnergyHigh: unifieddatalibrary.Float(0.003495),
			ChanEnergyLow:  unifieddatalibrary.Float(58.4),
			ChanID:         unifieddatalibrary.String("H05E"),
			ChanType:       unifieddatalibrary.String("INTEGRAL"),
			ChanUnit:       unifieddatalibrary.String("keV"),
			MsgNumber:      unifieddatalibrary.Int(21),
			ObValue:        unifieddatalibrary.Float(31473.9),
			Species:        unifieddatalibrary.String("ELECTRON"),
		}},
		ObTime:            time.Now(),
		Source:            "Bluestaq",
		ID:                unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		GenSystem:         unifieddatalibrary.String("cpuch2.aero.org"),
		GenTime:           unifieddatalibrary.Time(time.Now()),
		IDSensor:          unifieddatalibrary.String("REACH-101"),
		Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:      unifieddatalibrary.String("WSF-M SV1"),
		OrigSensorID:      unifieddatalibrary.String("CEASE-3"),
		SatNo:             unifieddatalibrary.Int(101),
		SenPos:            []float64{5893.74, 1408.8, 3899.38},
		SenReferenceFrame: unifieddatalibrary.EcpedrNewParamsSenReferenceFrameTeme,
		Tags:              []string{"TAG1", "TAG2"},
		TransactionID:     unifieddatalibrary.String("TRANSACTION-ID"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEcpedrListWithOptionalParams(t *testing.T) {
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
	_, err := client.Ecpedr.List(context.TODO(), unifieddatalibrary.EcpedrListParams{
		ObTime:      time.Now(),
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

func TestEcpedrCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Ecpedr.Count(context.TODO(), unifieddatalibrary.EcpedrCountParams{
		ObTime:      time.Now(),
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

func TestEcpedrNewBulk(t *testing.T) {
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
	err := client.Ecpedr.NewBulk(context.TODO(), unifieddatalibrary.EcpedrNewBulkParams{
		Body: []unifieddatalibrary.EcpedrNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EcpedrMeasurements: []unifieddatalibrary.EcpedrNewBulkParamsBodyEcpedrMeasurement{{
				ObType:         "FLUX",
				ObUoM:          "#/MeV/cm^2/sr/s",
				ChanEnergyHigh: unifieddatalibrary.Float(0.003495),
				ChanEnergyLow:  unifieddatalibrary.Float(58.4),
				ChanID:         unifieddatalibrary.String("H05E"),
				ChanType:       unifieddatalibrary.String("INTEGRAL"),
				ChanUnit:       unifieddatalibrary.String("keV"),
				MsgNumber:      unifieddatalibrary.Int(21),
				ObValue:        unifieddatalibrary.Float(31473.9),
				Species:        unifieddatalibrary.String("ELECTRON"),
			}},
			ObTime:            time.Now(),
			Source:            "Bluestaq",
			ID:                unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			GenSystem:         unifieddatalibrary.String("cpuch2.aero.org"),
			GenTime:           unifieddatalibrary.Time(time.Now()),
			IDSensor:          unifieddatalibrary.String("REACH-101"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:      unifieddatalibrary.String("WSF-M SV1"),
			OrigSensorID:      unifieddatalibrary.String("CEASE-3"),
			SatNo:             unifieddatalibrary.Int(101),
			SenPos:            []float64{5893.74, 1408.8, 3899.38},
			SenReferenceFrame: "TEME",
			Tags:              []string{"TAG1", "TAG2"},
			TransactionID:     unifieddatalibrary.String("TRANSACTION-ID"),
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

func TestEcpedrQueryhelp(t *testing.T) {
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
	_, err := client.Ecpedr.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEcpedrTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Ecpedr.Tuple(context.TODO(), unifieddatalibrary.EcpedrTupleParams{
		Columns:     "columns",
		ObTime:      time.Now(),
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

func TestEcpedrUnvalidatedPublish(t *testing.T) {
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
	err := client.Ecpedr.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EcpedrUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EcpedrUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EcpedrMeasurements: []unifieddatalibrary.EcpedrUnvalidatedPublishParamsBodyEcpedrMeasurement{{
				ObType:         "FLUX",
				ObUoM:          "#/MeV/cm^2/sr/s",
				ChanEnergyHigh: unifieddatalibrary.Float(0.003495),
				ChanEnergyLow:  unifieddatalibrary.Float(58.4),
				ChanID:         unifieddatalibrary.String("H05E"),
				ChanType:       unifieddatalibrary.String("INTEGRAL"),
				ChanUnit:       unifieddatalibrary.String("keV"),
				MsgNumber:      unifieddatalibrary.Int(21),
				ObValue:        unifieddatalibrary.Float(31473.9),
				Species:        unifieddatalibrary.String("ELECTRON"),
			}},
			ObTime:            time.Now(),
			Source:            "Bluestaq",
			ID:                unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			GenSystem:         unifieddatalibrary.String("cpuch2.aero.org"),
			GenTime:           unifieddatalibrary.Time(time.Now()),
			IDSensor:          unifieddatalibrary.String("REACH-101"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:      unifieddatalibrary.String("WSF-M SV1"),
			OrigSensorID:      unifieddatalibrary.String("CEASE-3"),
			SatNo:             unifieddatalibrary.Int(101),
			SenPos:            []float64{5893.74, 1408.8, 3899.38},
			SenReferenceFrame: "TEME",
			Tags:              []string{"TAG1", "TAG2"},
			TransactionID:     unifieddatalibrary.String("TRANSACTION-ID"),
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
