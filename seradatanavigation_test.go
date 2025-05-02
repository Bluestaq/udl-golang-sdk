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

func TestSeradatanavigationNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradatanavigation.New(context.TODO(), unifieddatalibrary.SeradatanavigationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SeradatanavigationNewParamsDataModeTest,
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

func TestSeradatanavigationUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradatanavigation.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatanavigationUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SeradatanavigationUpdateParamsDataModeTest,
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

func TestSeradatanavigationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatanavigation.List(context.TODO(), unifieddatalibrary.SeradatanavigationListParams{
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

func TestSeradatanavigationDelete(t *testing.T) {
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
	err := client.Seradatanavigation.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatanavigationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatanavigation.Count(context.TODO(), unifieddatalibrary.SeradatanavigationCountParams{
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

func TestSeradatanavigationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatanavigation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatanavigationGetParams{
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

func TestSeradatanavigationQueryhelp(t *testing.T) {
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
	err := client.Seradatanavigation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatanavigationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatanavigation.Tuple(context.TODO(), unifieddatalibrary.SeradatanavigationTupleParams{
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
