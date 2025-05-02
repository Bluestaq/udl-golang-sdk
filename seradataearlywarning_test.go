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

func TestSeradataearlywarningNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradataearlywarning.New(context.TODO(), unifieddatalibrary.SeradataearlywarningNewParams{
		ClassificationMarking:              "U",
		DataMode:                           unifieddatalibrary.SeradataearlywarningNewParamsDataModeTest,
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

func TestSeradataearlywarningUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradataearlywarning.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataearlywarningUpdateParams{
			ClassificationMarking:              "U",
			DataMode:                           unifieddatalibrary.SeradataearlywarningUpdateParamsDataModeTest,
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

func TestSeradataearlywarningListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataearlywarning.List(context.TODO(), unifieddatalibrary.SeradataearlywarningListParams{
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

func TestSeradataearlywarningDelete(t *testing.T) {
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
	err := client.Seradataearlywarning.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataearlywarningCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataearlywarning.Count(context.TODO(), unifieddatalibrary.SeradataearlywarningCountParams{
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

func TestSeradataearlywarningGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataearlywarning.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataearlywarningGetParams{
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

func TestSeradataearlywarningQueryhelp(t *testing.T) {
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
	err := client.Seradataearlywarning.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataearlywarningTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataearlywarning.Tuple(context.TODO(), unifieddatalibrary.SeradataearlywarningTupleParams{
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
