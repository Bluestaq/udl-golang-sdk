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

func TestSeradataRadarPayloadNewWithOptionalParams(t *testing.T) {
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
	err := client.SeradataRadarPayload.New(context.TODO(), unifieddatalibrary.SeradataRadarPayloadNewParams{
		ClassificationMarking:                  "U",
		DataMode:                               unifieddatalibrary.SeradataRadarPayloadNewParamsDataModeTest,
		Source:                                 "Bluestaq",
		SpacecraftID:                           "12345",
		ID:                                     unifieddatalibrary.String("SERADATARADARPAYLOAD-ID"),
		Bandwidth:                              unifieddatalibrary.Float(1.23),
		BestResolution:                         unifieddatalibrary.Float(1.23),
		Category:                               unifieddatalibrary.String("SAR"),
		ConstellationInterferometricCapability: unifieddatalibrary.String("constellationInterferometricCapability"),
		DutyCycle:                              unifieddatalibrary.String("dutyCycle"),
		FieldOfRegard:                          unifieddatalibrary.Float(1.23),
		FieldOfView:                            unifieddatalibrary.Float(1.23),
		Frequency:                              unifieddatalibrary.Float(1.23),
		FrequencyBand:                          unifieddatalibrary.String("X"),
		GroundStationLocations:                 unifieddatalibrary.String("51,42N-44,35E"),
		GroundStations:                         unifieddatalibrary.String("groundStations"),
		HostedForCompanyOrgID:                  unifieddatalibrary.String("hostedForCompanyOrgId"),
		IDSensor:                               unifieddatalibrary.String("3c1ee9a0-90ad-1d75-c47b-2414e0a77e53"),
		ManufacturerOrgID:                      unifieddatalibrary.String("manufacturerOrgId"),
		Name:                                   unifieddatalibrary.String("ALT"),
		Notes:                                  unifieddatalibrary.String("Sample Notes"),
		Origin:                                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PartnerSpacecraft:                      unifieddatalibrary.String("partnerSpacecraft"),
		PointingMethod:                         unifieddatalibrary.String("Spacecraft"),
		ReceivePolarization:                    unifieddatalibrary.String("Lin Dual"),
		RecorderSize:                           unifieddatalibrary.String("256"),
		SwathWidth:                             unifieddatalibrary.Float(1.23),
		TransmitPolarization:                   unifieddatalibrary.String("Lin Dual"),
		WaveLength:                             unifieddatalibrary.Float(1.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataRadarPayloadUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeradataRadarPayload.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataRadarPayloadUpdateParams{
			ClassificationMarking:                  "U",
			DataMode:                               unifieddatalibrary.SeradataRadarPayloadUpdateParamsDataModeTest,
			Source:                                 "Bluestaq",
			SpacecraftID:                           "12345",
			ID:                                     unifieddatalibrary.String("SERADATARADARPAYLOAD-ID"),
			Bandwidth:                              unifieddatalibrary.Float(1.23),
			BestResolution:                         unifieddatalibrary.Float(1.23),
			Category:                               unifieddatalibrary.String("SAR"),
			ConstellationInterferometricCapability: unifieddatalibrary.String("constellationInterferometricCapability"),
			DutyCycle:                              unifieddatalibrary.String("dutyCycle"),
			FieldOfRegard:                          unifieddatalibrary.Float(1.23),
			FieldOfView:                            unifieddatalibrary.Float(1.23),
			Frequency:                              unifieddatalibrary.Float(1.23),
			FrequencyBand:                          unifieddatalibrary.String("X"),
			GroundStationLocations:                 unifieddatalibrary.String("51,42N-44,35E"),
			GroundStations:                         unifieddatalibrary.String("groundStations"),
			HostedForCompanyOrgID:                  unifieddatalibrary.String("hostedForCompanyOrgId"),
			IDSensor:                               unifieddatalibrary.String("3c1ee9a0-90ad-1d75-c47b-2414e0a77e53"),
			ManufacturerOrgID:                      unifieddatalibrary.String("manufacturerOrgId"),
			Name:                                   unifieddatalibrary.String("ALT"),
			Notes:                                  unifieddatalibrary.String("Sample Notes"),
			Origin:                                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartnerSpacecraft:                      unifieddatalibrary.String("partnerSpacecraft"),
			PointingMethod:                         unifieddatalibrary.String("Spacecraft"),
			ReceivePolarization:                    unifieddatalibrary.String("Lin Dual"),
			RecorderSize:                           unifieddatalibrary.String("256"),
			SwathWidth:                             unifieddatalibrary.Float(1.23),
			TransmitPolarization:                   unifieddatalibrary.String("Lin Dual"),
			WaveLength:                             unifieddatalibrary.Float(1.23),
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

func TestSeradataRadarPayloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataRadarPayload.List(context.TODO(), unifieddatalibrary.SeradataRadarPayloadListParams{
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

func TestSeradataRadarPayloadDelete(t *testing.T) {
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
	err := client.SeradataRadarPayload.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataRadarPayloadCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataRadarPayload.Count(context.TODO(), unifieddatalibrary.SeradataRadarPayloadCountParams{
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

func TestSeradataRadarPayloadGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataRadarPayload.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataRadarPayloadGetParams{
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

func TestSeradataRadarPayloadQueryhelp(t *testing.T) {
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
	err := client.SeradataRadarPayload.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataRadarPayloadTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataRadarPayload.Tuple(context.TODO(), unifieddatalibrary.SeradataRadarPayloadTupleParams{
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
