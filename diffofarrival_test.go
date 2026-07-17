// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk/v2"
	"github.com/Bluestaq/udl-golang-sdk/v2/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
)

func TestDiffOfArrivalGetWithOptionalParams(t *testing.T) {
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
	_, err := client.DiffOfArrival.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.DiffOfArrivalGetParams{
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

func TestDiffOfArrivalQueryhelp(t *testing.T) {
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
	_, err := client.DiffOfArrival.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiffOfArrivalTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.DiffOfArrival.Tuple(context.TODO(), unifieddatalibrary.DiffOfArrivalTupleParams{
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

func TestDiffOfArrivalUnvalidatedPublish(t *testing.T) {
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
	err := client.DiffOfArrival.UnvalidatedPublish(context.TODO(), unifieddatalibrary.DiffOfArrivalUnvalidatedPublishParams{
		Body: []unifieddatalibrary.DiffOfArrivalUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("DIFFOFARRIVAL-ID"),
			Azimuth1:              unifieddatalibrary.Float(40.9),
			Azimuth2:              unifieddatalibrary.Float(43.4),
			Bandwidth:             unifieddatalibrary.Float(500000.1),
			CollectionMode:        unifieddatalibrary.String("SURVEY"),
			DeltaRange:            unifieddatalibrary.Float(1.7),
			DeltaRangeRate:        unifieddatalibrary.Float(0.018),
			DeltaRangeRateUnc:     unifieddatalibrary.Float(0.001),
			DeltaRangeUnc:         unifieddatalibrary.Float(0.2),
			Descriptor:            unifieddatalibrary.String("Example descriptor"),
			Elevation1:            unifieddatalibrary.Float(65.4),
			Elevation2:            unifieddatalibrary.Float(65.9),
			Fdoa:                  unifieddatalibrary.Float(970.8),
			FdoaUnc:               unifieddatalibrary.Float(20.1),
			Frequency:             unifieddatalibrary.Float(2240000000),
			IDSensor1:             unifieddatalibrary.String("SENSOR1-ID"),
			IDSensor2:             unifieddatalibrary.String("SENSOR2-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorId1:         unifieddatalibrary.String("ORIGSENSOR1-ID"),
			OrigSensorId2:         unifieddatalibrary.String("ORIGSENSOR2-ID"),
			PolarityType:          unifieddatalibrary.String("V"),
			RawFileUri:            unifieddatalibrary.String("/DOA/2026/01/01/filename.txt"),
			SatNo:                 unifieddatalibrary.Int(25544),
			Sen2alt:               unifieddatalibrary.Float(770.1),
			Sen2lat:               unifieddatalibrary.Float(32.2),
			Sen2lon:               unifieddatalibrary.Float(-110.9),
			Senalt:                unifieddatalibrary.Float(783.1),
			Senlat:                unifieddatalibrary.Float(45.1),
			Senlon:                unifieddatalibrary.Float(120.1),
			Sensor1Delay:          unifieddatalibrary.Float(0.000004),
			Sensor2Delay:          unifieddatalibrary.Float(0.000004),
			Snr:                   unifieddatalibrary.Float(18.1),
			Tags:                  []string{"TAG1", "TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			Tdoa:                  unifieddatalibrary.Float(0.0000136),
			TdoaAmb:               unifieddatalibrary.Int(1),
			TdoaUnc:               unifieddatalibrary.Float(2e-7),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
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
