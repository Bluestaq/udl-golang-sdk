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

func TestRfbandNewWithOptionalParams(t *testing.T) {
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
	err := client.Rfband.New(context.TODO(), unifieddatalibrary.RfbandNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RfbandNewParamsDataModeTest,
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
		Mode:                  unifieddatalibrary.RfbandNewParamsModeTx,
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PeakGain:              unifieddatalibrary.Float(120.23),
		Polarization:          unifieddatalibrary.RfbandNewParamsPolarizationH,
		Purpose:               unifieddatalibrary.RfbandNewParamsPurposeTtc,
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfbandUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Rfband.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RfbandUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.RfbandUpdateParamsDataModeTest,
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
			Mode:                  unifieddatalibrary.RfbandUpdateParamsModeTx,
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PeakGain:              unifieddatalibrary.Float(120.23),
			Polarization:          unifieddatalibrary.RfbandUpdateParamsPolarizationH,
			Purpose:               unifieddatalibrary.RfbandUpdateParamsPurposeTtc,
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

func TestRfbandListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rfband.List(context.TODO(), unifieddatalibrary.RfbandListParams{
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

func TestRfbandDelete(t *testing.T) {
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
	err := client.Rfband.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfbandCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Rfband.Count(context.TODO(), unifieddatalibrary.RfbandCountParams{
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

func TestRfbandGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rfband.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RfbandGetParams{
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

func TestRfbandQueryhelp(t *testing.T) {
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
	err := client.Rfband.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfbandTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Rfband.Tuple(context.TODO(), unifieddatalibrary.RfbandTupleParams{
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
