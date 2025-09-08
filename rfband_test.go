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
		ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
		Band:                  unifieddatalibrary.String("Ku"),
		Bandwidth:             unifieddatalibrary.Float(100.23),
		BandwidthSettings:     []float64{250.1, 500.1},
		Beamwidth:             unifieddatalibrary.Float(45.23),
		BeamwidthSettings:     []float64{5.1, 10.1},
		CenterFreq:            unifieddatalibrary.Float(1000.23),
		DelaySettings:         []float64{2.77, 5.64},
		EdgeGain:              unifieddatalibrary.Float(100.23),
		Eirp:                  unifieddatalibrary.Float(2.23),
		Erp:                   unifieddatalibrary.Float(2.23),
		FreqMax:               unifieddatalibrary.Float(2000.23),
		FreqMin:               unifieddatalibrary.Float(50.23),
		FrequencySettings:     []float64{12250.1, 15000.1},
		GainSettings:          []float64{2.77, 5.64},
		Mode:                  unifieddatalibrary.RfBandNewParamsModeTx,
		NoiseSettings:         []float64{0.00033, 0.0033},
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
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Band:                  unifieddatalibrary.String("Ku"),
			Bandwidth:             unifieddatalibrary.Float(100.23),
			BandwidthSettings:     []float64{250.1, 500.1},
			Beamwidth:             unifieddatalibrary.Float(45.23),
			BeamwidthSettings:     []float64{5.1, 10.1},
			CenterFreq:            unifieddatalibrary.Float(1000.23),
			DelaySettings:         []float64{2.77, 5.64},
			EdgeGain:              unifieddatalibrary.Float(100.23),
			Eirp:                  unifieddatalibrary.Float(2.23),
			Erp:                   unifieddatalibrary.Float(2.23),
			FreqMax:               unifieddatalibrary.Float(2000.23),
			FreqMin:               unifieddatalibrary.Float(50.23),
			FrequencySettings:     []float64{12250.1, 15000.1},
			GainSettings:          []float64{2.77, 5.64},
			Mode:                  unifieddatalibrary.RfBandUpdateParamsModeTx,
			NoiseSettings:         []float64{0.00033, 0.0033},
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
