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

func TestSeraDataNavigationNewWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataNavigation.New(context.TODO(), unifieddatalibrary.SeraDataNavigationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SeraDataNavigationNewParamsDataModeTest,
		Source:                "Bluestaq",
		SpacecraftID:          "spacecraftId",
		ID:                    unifieddatalibrary.String("SERADATANAVIGATION-ID"),
		AreaCoverage:          unifieddatalibrary.String("Worldwide"),
		ClockType:             unifieddatalibrary.String("Rubidium"),
		HostedForCompanyOrgID: unifieddatalibrary.String("hostedForCompanyOrgId"),
		IDNavigation:          unifieddatalibrary.String("idNavigation"),
		LocationAccuracy:      unifieddatalibrary.Float(1.23),
		ManufacturerOrgID:     unifieddatalibrary.String("manufacturerOrgId"),
		ModeFrequency:         unifieddatalibrary.String("1234"),
		Modes:                 unifieddatalibrary.String("Military"),
		Name:                  unifieddatalibrary.String("WAAS GEO-5"),
		Notes:                 unifieddatalibrary.String("Sample Notes"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PartnerSpacecraftID:   unifieddatalibrary.String("partnerSpacecraftId"),
		PayloadType:           unifieddatalibrary.String("WAAS"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataNavigationUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataNavigation.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataNavigationUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SeraDataNavigationUpdateParamsDataModeTest,
			Source:                "Bluestaq",
			SpacecraftID:          "spacecraftId",
			ID:                    unifieddatalibrary.String("SERADATANAVIGATION-ID"),
			AreaCoverage:          unifieddatalibrary.String("Worldwide"),
			ClockType:             unifieddatalibrary.String("Rubidium"),
			HostedForCompanyOrgID: unifieddatalibrary.String("hostedForCompanyOrgId"),
			IDNavigation:          unifieddatalibrary.String("idNavigation"),
			LocationAccuracy:      unifieddatalibrary.Float(1.23),
			ManufacturerOrgID:     unifieddatalibrary.String("manufacturerOrgId"),
			ModeFrequency:         unifieddatalibrary.String("1234"),
			Modes:                 unifieddatalibrary.String("Military"),
			Name:                  unifieddatalibrary.String("WAAS GEO-5"),
			Notes:                 unifieddatalibrary.String("Sample Notes"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartnerSpacecraftID:   unifieddatalibrary.String("partnerSpacecraftId"),
			PayloadType:           unifieddatalibrary.String("WAAS"),
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

func TestSeraDataNavigationListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataNavigation.List(context.TODO(), unifieddatalibrary.SeraDataNavigationListParams{
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

func TestSeraDataNavigationDelete(t *testing.T) {
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
	err := client.SeraDataNavigation.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataNavigationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataNavigation.Count(context.TODO(), unifieddatalibrary.SeraDataNavigationCountParams{
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

func TestSeraDataNavigationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataNavigation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataNavigationGetParams{
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

func TestSeraDataNavigationQueryhelp(t *testing.T) {
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
	err := client.SeraDataNavigation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataNavigationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataNavigation.Tuple(context.TODO(), unifieddatalibrary.SeraDataNavigationTupleParams{
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
