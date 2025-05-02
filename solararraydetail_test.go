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

func TestSolararraydetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Solararraydetails.New(context.TODO(), unifieddatalibrary.SolararraydetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SolararraydetailNewParamsDataModeTest,
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

func TestSolararraydetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Solararraydetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SolararraydetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SolararraydetailUpdateParamsDataModeTest,
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

func TestSolararraydetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Solararraydetails.List(context.TODO(), unifieddatalibrary.SolararraydetailListParams{
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

func TestSolararraydetailDelete(t *testing.T) {
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
	err := client.Solararraydetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSolararraydetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Solararraydetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SolararraydetailGetParams{
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
