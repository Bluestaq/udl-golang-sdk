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

func TestSeradataradarpayloadNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradataradarpayload.New(context.TODO(), unifieddatalibrary.SeradataradarpayloadNewParams{
		ClassificationMarking:                  "U",
		DataMode:                               unifieddatalibrary.SeradataradarpayloadNewParamsDataModeTest,
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

func TestSeradataradarpayloadUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradataradarpayload.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataradarpayloadUpdateParams{
			ClassificationMarking:                  "U",
			DataMode:                               unifieddatalibrary.SeradataradarpayloadUpdateParamsDataModeTest,
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

func TestSeradataradarpayloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataradarpayload.List(context.TODO(), unifieddatalibrary.SeradataradarpayloadListParams{
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

func TestSeradataradarpayloadDelete(t *testing.T) {
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
	err := client.Seradataradarpayload.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataradarpayloadCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataradarpayload.Count(context.TODO(), unifieddatalibrary.SeradataradarpayloadCountParams{
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

func TestSeradataradarpayloadGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataradarpayload.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataradarpayloadGetParams{
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

func TestSeradataradarpayloadQueryhelp(t *testing.T) {
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
	err := client.Seradataradarpayload.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataradarpayloadTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataradarpayload.Tuple(context.TODO(), unifieddatalibrary.SeradataradarpayloadTupleParams{
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
