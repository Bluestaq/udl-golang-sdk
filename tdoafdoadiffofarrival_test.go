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

func TestTdoaFdoaDiffofarrivalNewWithOptionalParams(t *testing.T) {
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
	err := client.TdoaFdoa.Diffofarrival.New(context.TODO(), unifieddatalibrary.TdoaFdoaDiffofarrivalNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.TdoaFdoaDiffofarrivalNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("DIFFOFARRIVAL-ID"),
		Bandwidth:             unifieddatalibrary.Float(1.1),
		CollectionMode:        unifieddatalibrary.String("SURVEY"),
		DeltaRange:            unifieddatalibrary.Float(1.1),
		DeltaRangeRate:        unifieddatalibrary.Float(1.1),
		DeltaRangeRateUnc:     unifieddatalibrary.Float(1.1),
		DeltaRangeUnc:         unifieddatalibrary.Float(1.1),
		Descriptor:            unifieddatalibrary.String("Example descriptor"),
		Fdoa:                  unifieddatalibrary.Float(1.1),
		FdoaUnc:               unifieddatalibrary.Float(1.1),
		Frequency:             unifieddatalibrary.Float(1.1),
		IDSensor1:             unifieddatalibrary.String("SENSOR1-ID"),
		IDSensor2:             unifieddatalibrary.String("SENSOR2-ID"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorId1:         unifieddatalibrary.String("ORIGSENSOR1-ID"),
		OrigSensorId2:         unifieddatalibrary.String("ORIGSENSOR2-ID"),
		RawFileUri:            unifieddatalibrary.String("rawFileURI"),
		SatNo:                 unifieddatalibrary.Int(25544),
		Sen2alt:               unifieddatalibrary.Float(1.1),
		Sen2lat:               unifieddatalibrary.Float(1.1),
		Sen2lon:               unifieddatalibrary.Float(1.1),
		Senalt:                unifieddatalibrary.Float(1.1),
		Senlat:                unifieddatalibrary.Float(45.1),
		Senlon:                unifieddatalibrary.Float(120.1),
		Sensor1Delay:          unifieddatalibrary.Float(1.1),
		Sensor2Delay:          unifieddatalibrary.Float(1.1),
		Snr:                   unifieddatalibrary.Float(1.1),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TaskID:                unifieddatalibrary.String("TASK-ID"),
		Tdoa:                  unifieddatalibrary.Float(1.1),
		TdoaUnc:               unifieddatalibrary.Float(1.1),
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		Uct:                   unifieddatalibrary.Bool(false),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTdoaFdoaDiffofarrivalListWithOptionalParams(t *testing.T) {
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
	_, err := client.TdoaFdoa.Diffofarrival.List(context.TODO(), unifieddatalibrary.TdoaFdoaDiffofarrivalListParams{
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

func TestTdoaFdoaDiffofarrivalCountWithOptionalParams(t *testing.T) {
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
	_, err := client.TdoaFdoa.Diffofarrival.Count(context.TODO(), unifieddatalibrary.TdoaFdoaDiffofarrivalCountParams{
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

func TestTdoaFdoaDiffofarrivalNewBulk(t *testing.T) {
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
	err := client.TdoaFdoa.Diffofarrival.NewBulk(context.TODO(), unifieddatalibrary.TdoaFdoaDiffofarrivalNewBulkParams{
		Body: []unifieddatalibrary.TdoaFdoaDiffofarrivalNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("DIFFOFARRIVAL-ID"),
			Bandwidth:             unifieddatalibrary.Float(1.1),
			CollectionMode:        unifieddatalibrary.String("SURVEY"),
			DeltaRange:            unifieddatalibrary.Float(1.1),
			DeltaRangeRate:        unifieddatalibrary.Float(1.1),
			DeltaRangeRateUnc:     unifieddatalibrary.Float(1.1),
			DeltaRangeUnc:         unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("Example descriptor"),
			Fdoa:                  unifieddatalibrary.Float(1.1),
			FdoaUnc:               unifieddatalibrary.Float(1.1),
			Frequency:             unifieddatalibrary.Float(1.1),
			IDSensor1:             unifieddatalibrary.String("SENSOR1-ID"),
			IDSensor2:             unifieddatalibrary.String("SENSOR2-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorId1:         unifieddatalibrary.String("ORIGSENSOR1-ID"),
			OrigSensorId2:         unifieddatalibrary.String("ORIGSENSOR2-ID"),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			SatNo:                 unifieddatalibrary.Int(25544),
			Sen2alt:               unifieddatalibrary.Float(1.1),
			Sen2lat:               unifieddatalibrary.Float(1.1),
			Sen2lon:               unifieddatalibrary.Float(1.1),
			Senalt:                unifieddatalibrary.Float(1.1),
			Senlat:                unifieddatalibrary.Float(45.1),
			Senlon:                unifieddatalibrary.Float(120.1),
			Sensor1Delay:          unifieddatalibrary.Float(1.1),
			Sensor2Delay:          unifieddatalibrary.Float(1.1),
			Snr:                   unifieddatalibrary.Float(1.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			Tdoa:                  unifieddatalibrary.Float(1.1),
			TdoaUnc:               unifieddatalibrary.Float(1.1),
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
