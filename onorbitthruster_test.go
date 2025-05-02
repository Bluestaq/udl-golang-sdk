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

func TestOnorbitthrusterNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitthruster.New(context.TODO(), unifieddatalibrary.OnorbitthrusterNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitthrusterNewParamsDataModeTest,
		IDEngine:              "ENGINE-ID",
		IDOnOrbit:             "ONORBIT-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ONORBITTHRUSTER-ID"),
		Engine: unifieddatalibrary.OnorbitthrusterNewParamsEngine{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Name:                  "ENGINE_VARIANT1",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ENGINE-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		},
		Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Quantity: unifieddatalibrary.Int(10),
		Type:     unifieddatalibrary.String("Hydrazine REA"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitthruster.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitthrusterUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OnorbitthrusterUpdateParamsDataModeTest,
			IDEngine:              "ENGINE-ID",
			IDOnOrbit:             "ONORBIT-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ONORBITTHRUSTER-ID"),
			Engine: unifieddatalibrary.OnorbitthrusterUpdateParamsEngine{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Name:                  "ENGINE_VARIANT1",
				Source:                "Bluestaq",
				ID:                    unifieddatalibrary.String("ENGINE-ID"),
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Quantity: unifieddatalibrary.Int(10),
			Type:     unifieddatalibrary.String("Hydrazine REA"),
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

func TestOnorbitthrusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthruster.List(context.TODO(), unifieddatalibrary.OnorbitthrusterListParams{
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

func TestOnorbitthrusterDelete(t *testing.T) {
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
	err := client.Onorbitthruster.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthruster.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitthrusterGetParams{
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
