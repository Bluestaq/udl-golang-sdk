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

func TestSeraDataEarlyWarningNewWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataEarlyWarning.New(context.TODO(), unifieddatalibrary.SeraDataEarlyWarningNewParams{
		ClassificationMarking:              "U",
		DataMode:                           unifieddatalibrary.SeraDataEarlyWarningNewParamsDataModeTest,
		Source:                             "Bluestaq",
		SpacecraftID:                       "spacecraftId",
		ID:                                 unifieddatalibrary.String("SERADATAEARLYWARNING-ID"),
		BestResolution:                     unifieddatalibrary.Float(1.23),
		EarthPointing:                      unifieddatalibrary.Bool(true),
		FrequencyLimits:                    unifieddatalibrary.String("frequencyLimits"),
		GroundStationLocations:             unifieddatalibrary.String("groundStationLocations"),
		GroundStations:                     unifieddatalibrary.String("groundStations"),
		HostedForCompanyOrgID:              unifieddatalibrary.String("hostedForCompanyOrgId"),
		IDIr:                               unifieddatalibrary.String("idIR"),
		ManufacturerOrgID:                  unifieddatalibrary.String("manufacturerOrgId"),
		MissileLaunchPhaseDetectionAbility: unifieddatalibrary.String("missileLaunchPhaseDetectionAbility"),
		Name:                               unifieddatalibrary.String("Infra red telescope"),
		Origin:                             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PartnerSpacecraftID:                unifieddatalibrary.String("partnerSpacecraftId"),
		PayloadNotes:                       unifieddatalibrary.String("Sample Notes"),
		SpectralBands:                      unifieddatalibrary.String("Infra-Red"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataEarlyWarningUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataEarlyWarning.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataEarlyWarningUpdateParams{
			ClassificationMarking:              "U",
			DataMode:                           unifieddatalibrary.SeraDataEarlyWarningUpdateParamsDataModeTest,
			Source:                             "Bluestaq",
			SpacecraftID:                       "spacecraftId",
			ID:                                 unifieddatalibrary.String("SERADATAEARLYWARNING-ID"),
			BestResolution:                     unifieddatalibrary.Float(1.23),
			EarthPointing:                      unifieddatalibrary.Bool(true),
			FrequencyLimits:                    unifieddatalibrary.String("frequencyLimits"),
			GroundStationLocations:             unifieddatalibrary.String("groundStationLocations"),
			GroundStations:                     unifieddatalibrary.String("groundStations"),
			HostedForCompanyOrgID:              unifieddatalibrary.String("hostedForCompanyOrgId"),
			IDIr:                               unifieddatalibrary.String("idIR"),
			ManufacturerOrgID:                  unifieddatalibrary.String("manufacturerOrgId"),
			MissileLaunchPhaseDetectionAbility: unifieddatalibrary.String("missileLaunchPhaseDetectionAbility"),
			Name:                               unifieddatalibrary.String("Infra red telescope"),
			Origin:                             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartnerSpacecraftID:                unifieddatalibrary.String("partnerSpacecraftId"),
			PayloadNotes:                       unifieddatalibrary.String("Sample Notes"),
			SpectralBands:                      unifieddatalibrary.String("Infra-Red"),
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

func TestSeraDataEarlyWarningListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataEarlyWarning.List(context.TODO(), unifieddatalibrary.SeraDataEarlyWarningListParams{
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

func TestSeraDataEarlyWarningDelete(t *testing.T) {
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
	err := client.SeraDataEarlyWarning.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataEarlyWarningCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataEarlyWarning.Count(context.TODO(), unifieddatalibrary.SeraDataEarlyWarningCountParams{
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

func TestSeraDataEarlyWarningGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataEarlyWarning.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataEarlyWarningGetParams{
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

func TestSeraDataEarlyWarningQueryhelp(t *testing.T) {
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
	_, err := client.SeraDataEarlyWarning.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataEarlyWarningTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataEarlyWarning.Tuple(context.TODO(), unifieddatalibrary.SeraDataEarlyWarningTupleParams{
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
