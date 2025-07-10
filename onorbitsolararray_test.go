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

func TestOnorbitsolararrayNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitsolararray.New(context.TODO(), unifieddatalibrary.OnorbitsolararrayNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitsolararrayNewParamsDataModeTest,
		IDOnOrbit:             "ONORBIT-ID",
		IDSolarArray:          "SOLARARRAY-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ONORBITSOLARARRAY-ID"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Quantity:              unifieddatalibrary.Int(10),
		SolarArray: unifieddatalibrary.OnorbitsolararrayNewParamsSolarArray{
			DataMode: "TEST",
			Name:     "Solar1",
			Source:   "Bluestaq",
			ID:       unifieddatalibrary.String("SOLARARRAY-ID"),
			Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitsolararrayUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitsolararray.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitsolararrayUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OnorbitsolararrayUpdateParamsDataModeTest,
			IDOnOrbit:             "ONORBIT-ID",
			IDSolarArray:          "SOLARARRAY-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ONORBITSOLARARRAY-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Quantity:              unifieddatalibrary.Int(10),
			SolarArray: unifieddatalibrary.OnorbitsolararrayUpdateParamsSolarArray{
				DataMode: "TEST",
				Name:     "Solar1",
				Source:   "Bluestaq",
				ID:       unifieddatalibrary.String("SOLARARRAY-ID"),
				Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
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

func TestOnorbitsolararrayListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitsolararray.List(context.TODO(), unifieddatalibrary.OnorbitsolararrayListParams{
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

func TestOnorbitsolararrayDelete(t *testing.T) {
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
	err := client.Onorbitsolararray.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitsolararrayGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitsolararray.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitsolararrayGetParams{
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
