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

func TestSolarArrayDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.SolarArrayDetails.New(context.TODO(), unifieddatalibrary.SolarArrayDetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SolarArrayDetailNewParamsDataModeTest,
		IDSolarArray:          "SOLARARRAY-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SOLARARRAYDETAILS-ID"),
		Area:                  unifieddatalibrary.Float(123.4),
		Description:           unifieddatalibrary.String("Example notes"),
		JunctionTechnology:    unifieddatalibrary.String("Triple"),
		ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Span:                  unifieddatalibrary.Float(123.4),
		Tags:                  []string{"TAG1", "TAG2"},
		Technology:            unifieddatalibrary.String("Ga-As"),
		Type:                  unifieddatalibrary.String("U Shaped"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSolarArrayDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SolarArrayDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SolarArrayDetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SolarArrayDetailUpdateParamsDataModeTest,
			IDSolarArray:          "SOLARARRAY-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SOLARARRAYDETAILS-ID"),
			Area:                  unifieddatalibrary.Float(123.4),
			Description:           unifieddatalibrary.String("Example notes"),
			JunctionTechnology:    unifieddatalibrary.String("Triple"),
			ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Span:                  unifieddatalibrary.Float(123.4),
			Tags:                  []string{"TAG1", "TAG2"},
			Technology:            unifieddatalibrary.String("Ga-As"),
			Type:                  unifieddatalibrary.String("U Shaped"),
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

func TestSolarArrayDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.SolarArrayDetails.List(context.TODO(), unifieddatalibrary.SolarArrayDetailListParams{
		ClassificationMarking: unifieddatalibrary.String("classificationMarking"),
		DataMode:              unifieddatalibrary.String("dataMode"),
		FirstResult:           unifieddatalibrary.Int(0),
		MaxResults:            unifieddatalibrary.Int(0),
		Source:                unifieddatalibrary.String("source"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSolarArrayDetailDelete(t *testing.T) {
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
	err := client.SolarArrayDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSolarArrayDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SolarArrayDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SolarArrayDetailGetParams{
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
