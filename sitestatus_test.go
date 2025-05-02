// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestSitestatusNewWithOptionalParams(t *testing.T) {
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
	err := client.Sitestatus.New(context.TODO(), unifieddatalibrary.SitestatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SitestatusNewParamsDataModeTest,
		IDSite:                "41e3e554-9790-40b9-bd7b-f30d864dcad8",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SITESTATUS-ID"),
		Cat:                   unifieddatalibrary.SitestatusNewParamsCatCold,
		ColdInventory:         unifieddatalibrary.Int(1),
		CommImpairment:        unifieddatalibrary.String("commImpairment"),
		Cpcon:                 unifieddatalibrary.SitestatusNewParamsCpcon4,
		Eoc:                   unifieddatalibrary.SitestatusNewParamsEocWarm,
		Fpcon:                 unifieddatalibrary.SitestatusNewParamsFpconBravo,
		HotInventory:          unifieddatalibrary.Int(1),
		Hpcon:                 unifieddatalibrary.SitestatusNewParamsHpconCharlie,
		InstStatus:            unifieddatalibrary.SitestatusNewParamsInstStatusPmc,
		Link:                  []string{"ATDL", "IJMS", "LINK-1"},
		LinkStatus:            []string{"AVAILABLE", "DEGRADED", "NOT AVAILABLE"},
		Missile:               []string{"GMD", "HARPOON", "JAVELIN"},
		MissileInventory:      []int64{5, 10, 100},
		MobileAltID:           unifieddatalibrary.String("MOBILEALT-ID"),
		OpsCapability:         unifieddatalibrary.String("Fully Operational"),
		OpsImpairment:         unifieddatalibrary.String("opsImpairment"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Pes:                   unifieddatalibrary.Bool(true),
		Poiid:                 unifieddatalibrary.String("d4a91864-6140-4b8d-67cd-45421c75f696"),
		RadarStatus:           []string{"OPERATIONAL", "OFF", "NON-OPERATIONAL"},
		RadarSystem:           []string{"ILLUMINATING", "MODE-4", "MODE-3"},
		RadiateMode:           unifieddatalibrary.String("Active"),
		ReportTime:            unifieddatalibrary.Time(time.Now()),
		SamMode:               unifieddatalibrary.String("Initialization"),
		SiteType:              unifieddatalibrary.String("ADOC"),
		TimeFunction:          unifieddatalibrary.String("Activation"),
		TrackID:               unifieddatalibrary.String("PCM4923-1656174732-4-1-257"),
		TrackRefL16:           unifieddatalibrary.String("TrkNm"),
		WeatherMessage:        unifieddatalibrary.String("Heavy rain"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSitestatusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sitestatus.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SitestatusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SitestatusUpdateParamsDataModeTest,
			IDSite:                "41e3e554-9790-40b9-bd7b-f30d864dcad8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SITESTATUS-ID"),
			Cat:                   unifieddatalibrary.SitestatusUpdateParamsCatCold,
			ColdInventory:         unifieddatalibrary.Int(1),
			CommImpairment:        unifieddatalibrary.String("commImpairment"),
			Cpcon:                 unifieddatalibrary.SitestatusUpdateParamsCpcon4,
			Eoc:                   unifieddatalibrary.SitestatusUpdateParamsEocWarm,
			Fpcon:                 unifieddatalibrary.SitestatusUpdateParamsFpconBravo,
			HotInventory:          unifieddatalibrary.Int(1),
			Hpcon:                 unifieddatalibrary.SitestatusUpdateParamsHpconCharlie,
			InstStatus:            unifieddatalibrary.SitestatusUpdateParamsInstStatusPmc,
			Link:                  []string{"ATDL", "IJMS", "LINK-1"},
			LinkStatus:            []string{"AVAILABLE", "DEGRADED", "NOT AVAILABLE"},
			Missile:               []string{"GMD", "HARPOON", "JAVELIN"},
			MissileInventory:      []int64{5, 10, 100},
			MobileAltID:           unifieddatalibrary.String("MOBILEALT-ID"),
			OpsCapability:         unifieddatalibrary.String("Fully Operational"),
			OpsImpairment:         unifieddatalibrary.String("opsImpairment"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Pes:                   unifieddatalibrary.Bool(true),
			Poiid:                 unifieddatalibrary.String("d4a91864-6140-4b8d-67cd-45421c75f696"),
			RadarStatus:           []string{"OPERATIONAL", "OFF", "NON-OPERATIONAL"},
			RadarSystem:           []string{"ILLUMINATING", "MODE-4", "MODE-3"},
			RadiateMode:           unifieddatalibrary.String("Active"),
			ReportTime:            unifieddatalibrary.Time(time.Now()),
			SamMode:               unifieddatalibrary.String("Initialization"),
			SiteType:              unifieddatalibrary.String("ADOC"),
			TimeFunction:          unifieddatalibrary.String("Activation"),
			TrackID:               unifieddatalibrary.String("PCM4923-1656174732-4-1-257"),
			TrackRefL16:           unifieddatalibrary.String("TrkNm"),
			WeatherMessage:        unifieddatalibrary.String("Heavy rain"),
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

func TestSitestatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sitestatus.List(context.TODO(), unifieddatalibrary.SitestatusListParams{
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

func TestSitestatusDelete(t *testing.T) {
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
	err := client.Sitestatus.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSitestatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sitestatus.Count(context.TODO(), unifieddatalibrary.SitestatusCountParams{
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

func TestSitestatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sitestatus.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SitestatusGetParams{
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

func TestSitestatusQueryhelp(t *testing.T) {
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
	err := client.Sitestatus.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSitestatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sitestatus.Tuple(context.TODO(), unifieddatalibrary.SitestatusTupleParams{
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
