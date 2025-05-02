// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestBatterydetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Batterydetails.New(context.TODO(), unifieddatalibrary.BatterydetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.BatterydetailNewParamsDataModeTest,
		IDBattery:             "BATTERY-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("BATTERYDETAILS-ID"),
		Capacity:              unifieddatalibrary.Float(10.1),
		Description:           unifieddatalibrary.String("example notes"),
		DischargeDepth:        unifieddatalibrary.Float(0.2),
		ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
		Model:                 unifieddatalibrary.String("11212"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		Technology:            unifieddatalibrary.String("Ni-Cd"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBatterydetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Batterydetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.BatterydetailGetParams{
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

func TestBatterydetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Batterydetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.BatterydetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.BatterydetailUpdateParamsDataModeTest,
			IDBattery:             "BATTERY-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("BATTERYDETAILS-ID"),
			Capacity:              unifieddatalibrary.Float(10.1),
			Description:           unifieddatalibrary.String("example notes"),
			DischargeDepth:        unifieddatalibrary.Float(0.2),
			ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
			Model:                 unifieddatalibrary.String("11212"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			Technology:            unifieddatalibrary.String("Ni-Cd"),
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

func TestBatterydetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Batterydetails.List(context.TODO(), unifieddatalibrary.BatterydetailListParams{
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

func TestBatterydetailDelete(t *testing.T) {
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
	err := client.Batterydetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
