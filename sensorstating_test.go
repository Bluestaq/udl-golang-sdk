// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestSensorStatingNewWithOptionalParams(t *testing.T) {
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
	err := client.SensorStating.New(context.TODO(), unifieddatalibrary.SensorStatingNewParams{
		ClassificationMarking: "U",
		SensorName:            "SENSOR_NAME",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
		Altitude:              unifieddatalibrary.Float(157.543),
		Lat:                   unifieddatalibrary.Float(48.6732),
		LocationCountry:       unifieddatalibrary.String("UA"),
		Lon:                   unifieddatalibrary.Float(22.8455),
		OwnerCountry:          unifieddatalibrary.String("UA"),
		SensorNumber:          unifieddatalibrary.Int(1234),
		SensorObservationType: unifieddatalibrary.String("5"),
		SensorType:            unifieddatalibrary.String("Space Borne"),
		ShortName:             unifieddatalibrary.String("SNR-1"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorStatingUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SensorStating.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorStatingUpdateParams{
			ClassificationMarking: "U",
			SensorName:            "SENSOR_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(157.543),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("UA"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("UA"),
			SensorNumber:          unifieddatalibrary.Int(1234),
			SensorObservationType: unifieddatalibrary.String("5"),
			SensorType:            unifieddatalibrary.String("Space Borne"),
			ShortName:             unifieddatalibrary.String("SNR-1"),
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

func TestSensorStatingListWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorStating.List(context.TODO(), unifieddatalibrary.SensorStatingListParams{
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

func TestSensorStatingDelete(t *testing.T) {
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
	err := client.SensorStating.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorStatingNewBulk(t *testing.T) {
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
	err := client.SensorStating.NewBulk(context.TODO(), unifieddatalibrary.SensorStatingNewBulkParams{
		Body: []unifieddatalibrary.SensorStatingNewBulkParamsBody{{
			ClassificationMarking: "U",
			SensorName:            "SENSOR_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(157.543),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("UA"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("UA"),
			SensorNumber:          unifieddatalibrary.Int(1234),
			SensorObservationType: unifieddatalibrary.String("5"),
			SensorType:            unifieddatalibrary.String("Space Borne"),
			ShortName:             unifieddatalibrary.String("SNR-1"),
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

func TestSensorStatingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorStating.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorStatingGetParams{
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

func TestSensorStatingQueryhelp(t *testing.T) {
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
	_, err := client.SensorStating.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
