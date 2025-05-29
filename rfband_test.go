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

func TestRfBandNewWithOptionalParams(t *testing.T) {
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
	err := client.RfBand.New(context.TODO(), unifieddatalibrary.RfBandNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RfBandNewParamsDataModeTest,
		IDEntity:              "ENTITY-ID",
		Name:                  "BAND_NAME",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("RFBAND-ID"),
		Band:                  unifieddatalibrary.String("Ku"),
		Bandwidth:             unifieddatalibrary.Float(100.23),
		Beamwidth:             unifieddatalibrary.Float(45.23),
		CenterFreq:            unifieddatalibrary.Float(1000.23),
		EdgeGain:              unifieddatalibrary.Float(100.23),
		Eirp:                  unifieddatalibrary.Float(2.23),
		Erp:                   unifieddatalibrary.Float(2.23),
		FreqMax:               unifieddatalibrary.Float(2000.23),
		FreqMin:               unifieddatalibrary.Float(50.23),
		Mode:                  unifieddatalibrary.RfBandNewParamsModeTx,
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PeakGain:              unifieddatalibrary.Float(120.23),
		Polarization:          unifieddatalibrary.RfBandNewParamsPolarizationH,
		Purpose:               unifieddatalibrary.RfBandNewParamsPurposeTtc,
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfBandUpdateWithOptionalParams(t *testing.T) {
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
	err := client.RfBand.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RfBandUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.RfBandUpdateParamsDataModeTest,
			IDEntity:              "ENTITY-ID",
			Name:                  "BAND_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("RFBAND-ID"),
			Band:                  unifieddatalibrary.String("Ku"),
			Bandwidth:             unifieddatalibrary.Float(100.23),
			Beamwidth:             unifieddatalibrary.Float(45.23),
			CenterFreq:            unifieddatalibrary.Float(1000.23),
			EdgeGain:              unifieddatalibrary.Float(100.23),
			Eirp:                  unifieddatalibrary.Float(2.23),
			Erp:                   unifieddatalibrary.Float(2.23),
			FreqMax:               unifieddatalibrary.Float(2000.23),
			FreqMin:               unifieddatalibrary.Float(50.23),
			Mode:                  unifieddatalibrary.RfBandUpdateParamsModeTx,
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PeakGain:              unifieddatalibrary.Float(120.23),
			Polarization:          unifieddatalibrary.RfBandUpdateParamsPolarizationH,
			Purpose:               unifieddatalibrary.RfBandUpdateParamsPurposeTtc,
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

func TestRfBandListWithOptionalParams(t *testing.T) {
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
	_, err := client.RfBand.List(context.TODO(), unifieddatalibrary.RfBandListParams{
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

func TestRfBandDelete(t *testing.T) {
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
	err := client.RfBand.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfBandCountWithOptionalParams(t *testing.T) {
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
	_, err := client.RfBand.Count(context.TODO(), unifieddatalibrary.RfBandCountParams{
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

func TestRfBandGetWithOptionalParams(t *testing.T) {
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
	_, err := client.RfBand.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RfBandGetParams{
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

func TestRfBandQueryhelp(t *testing.T) {
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
	_, err := client.RfBand.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfBandTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.RfBand.Tuple(context.TODO(), unifieddatalibrary.RfBandTupleParams{
		Columns:     "columns",
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
