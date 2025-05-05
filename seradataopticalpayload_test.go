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

func TestSeradataOpticalPayloadNewWithOptionalParams(t *testing.T) {
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
	err := client.SeradataOpticalPayload.New(context.TODO(), unifieddatalibrary.SeradataOpticalPayloadNewParams{
		ClassificationMarking:       "U",
		DataMode:                    unifieddatalibrary.SeradataOpticalPayloadNewParamsDataModeTest,
		Source:                      "Bluestaq",
		SpacecraftID:                "spacecraftId",
		ID:                          unifieddatalibrary.String("SERADATAOPTICALPAYLOAD-ID"),
		BestResolution:              unifieddatalibrary.Float(1.23),
		FieldOfRegard:               unifieddatalibrary.Float(1.23),
		FieldOfView:                 unifieddatalibrary.Float(1.23),
		GroundStationLocations:      unifieddatalibrary.String("groundStationLocations"),
		GroundStations:              unifieddatalibrary.String("groundStations"),
		HostedForCompanyOrgID:       unifieddatalibrary.String("hostedForCompanyOrgId"),
		IDSensor:                    unifieddatalibrary.String("idSensor"),
		ImagingPayloadCategory:      unifieddatalibrary.String("Infrared"),
		ManufacturerOrgID:           unifieddatalibrary.String("manufacturerOrgId"),
		Name:                        unifieddatalibrary.String("TOURNESOL"),
		Notes:                       unifieddatalibrary.String("Sample Notes"),
		NumberOfFilmReturnCanisters: unifieddatalibrary.Int(1),
		Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PointingMethod:              unifieddatalibrary.String("Spacecraft"),
		RecorderSize:                unifieddatalibrary.String("1024"),
		SpectralBand:                unifieddatalibrary.String("Green"),
		SpectralFrequencyLimits:     unifieddatalibrary.String("0.51"),
		SwathWidth:                  unifieddatalibrary.Float(1.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataOpticalPayloadUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeradataOpticalPayload.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataOpticalPayloadUpdateParams{
			ClassificationMarking:       "U",
			DataMode:                    unifieddatalibrary.SeradataOpticalPayloadUpdateParamsDataModeTest,
			Source:                      "Bluestaq",
			SpacecraftID:                "spacecraftId",
			ID:                          unifieddatalibrary.String("SERADATAOPTICALPAYLOAD-ID"),
			BestResolution:              unifieddatalibrary.Float(1.23),
			FieldOfRegard:               unifieddatalibrary.Float(1.23),
			FieldOfView:                 unifieddatalibrary.Float(1.23),
			GroundStationLocations:      unifieddatalibrary.String("groundStationLocations"),
			GroundStations:              unifieddatalibrary.String("groundStations"),
			HostedForCompanyOrgID:       unifieddatalibrary.String("hostedForCompanyOrgId"),
			IDSensor:                    unifieddatalibrary.String("idSensor"),
			ImagingPayloadCategory:      unifieddatalibrary.String("Infrared"),
			ManufacturerOrgID:           unifieddatalibrary.String("manufacturerOrgId"),
			Name:                        unifieddatalibrary.String("TOURNESOL"),
			Notes:                       unifieddatalibrary.String("Sample Notes"),
			NumberOfFilmReturnCanisters: unifieddatalibrary.Int(1),
			Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PointingMethod:              unifieddatalibrary.String("Spacecraft"),
			RecorderSize:                unifieddatalibrary.String("1024"),
			SpectralBand:                unifieddatalibrary.String("Green"),
			SpectralFrequencyLimits:     unifieddatalibrary.String("0.51"),
			SwathWidth:                  unifieddatalibrary.Float(1.23),
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

func TestSeradataOpticalPayloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataOpticalPayload.List(context.TODO(), unifieddatalibrary.SeradataOpticalPayloadListParams{
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

func TestSeradataOpticalPayloadDelete(t *testing.T) {
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
	err := client.SeradataOpticalPayload.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataOpticalPayloadCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataOpticalPayload.Count(context.TODO(), unifieddatalibrary.SeradataOpticalPayloadCountParams{
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

func TestSeradataOpticalPayloadGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataOpticalPayload.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataOpticalPayloadGetParams{
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

func TestSeradataOpticalPayloadQueryhelp(t *testing.T) {
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
	err := client.SeradataOpticalPayload.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataOpticalPayloadTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataOpticalPayload.Tuple(context.TODO(), unifieddatalibrary.SeradataOpticalPayloadTupleParams{
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
