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

func TestOnorbitbatteryNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitbattery.New(context.TODO(), unifieddatalibrary.OnorbitbatteryNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitbatteryNewParamsDataModeTest,
		IDBattery:             "BATTERY-ID",
		IDOnOrbit:             "ONORBIT-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ONORBITBATTERY-ID"),
		Battery: unifieddatalibrary.OnorbitbatteryNewParamsBattery{
			DataMode: "TEST",
			Name:     "JAK-BATTERY-1479",
			Source:   "Bluestaq",
			ID:       unifieddatalibrary.String("BATTERY-ID"),
			Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		},
		Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Quantity: unifieddatalibrary.Int(5),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitbatteryUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitbattery.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitbatteryUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OnorbitbatteryUpdateParamsDataModeTest,
			IDBattery:             "BATTERY-ID",
			IDOnOrbit:             "ONORBIT-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ONORBITBATTERY-ID"),
			Battery: unifieddatalibrary.OnorbitbatteryUpdateParamsBattery{
				DataMode: "TEST",
				Name:     "JAK-BATTERY-1479",
				Source:   "Bluestaq",
				ID:       unifieddatalibrary.String("BATTERY-ID"),
				Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Quantity: unifieddatalibrary.Int(5),
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

func TestOnorbitbatteryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitbattery.List(context.TODO(), unifieddatalibrary.OnorbitbatteryListParams{
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

func TestOnorbitbatteryDelete(t *testing.T) {
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
	err := client.Onorbitbattery.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitbatteryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitbattery.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitbatteryGetParams{
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
