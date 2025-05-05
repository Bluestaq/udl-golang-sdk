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

func TestSiteStatusNewWithOptionalParams(t *testing.T) {
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
	err := client.SiteStatus.New(context.TODO(), unifieddatalibrary.SiteStatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SiteStatusNewParamsDataModeTest,
		IDSite:                "41e3e554-9790-40b9-bd7b-f30d864dcad8",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SITESTATUS-ID"),
		Cat:                   unifieddatalibrary.SiteStatusNewParamsCatCold,
		ColdInventory:         unifieddatalibrary.Int(1),
		CommImpairment:        unifieddatalibrary.String("commImpairment"),
		Cpcon:                 unifieddatalibrary.SiteStatusNewParamsCpcon4,
		Eoc:                   unifieddatalibrary.SiteStatusNewParamsEocWarm,
		Fpcon:                 unifieddatalibrary.SiteStatusNewParamsFpconBravo,
		HotInventory:          unifieddatalibrary.Int(1),
		Hpcon:                 unifieddatalibrary.SiteStatusNewParamsHpconCharlie,
		InstStatus:            unifieddatalibrary.SiteStatusNewParamsInstStatusPmc,
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

func TestSiteStatusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SiteStatus.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteStatusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SiteStatusUpdateParamsDataModeTest,
			IDSite:                "41e3e554-9790-40b9-bd7b-f30d864dcad8",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SITESTATUS-ID"),
			Cat:                   unifieddatalibrary.SiteStatusUpdateParamsCatCold,
			ColdInventory:         unifieddatalibrary.Int(1),
			CommImpairment:        unifieddatalibrary.String("commImpairment"),
			Cpcon:                 unifieddatalibrary.SiteStatusUpdateParamsCpcon4,
			Eoc:                   unifieddatalibrary.SiteStatusUpdateParamsEocWarm,
			Fpcon:                 unifieddatalibrary.SiteStatusUpdateParamsFpconBravo,
			HotInventory:          unifieddatalibrary.Int(1),
			Hpcon:                 unifieddatalibrary.SiteStatusUpdateParamsHpconCharlie,
			InstStatus:            unifieddatalibrary.SiteStatusUpdateParamsInstStatusPmc,
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

func TestSiteStatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.SiteStatus.List(context.TODO(), unifieddatalibrary.SiteStatusListParams{
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

func TestSiteStatusDelete(t *testing.T) {
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
	err := client.SiteStatus.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteStatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SiteStatus.Count(context.TODO(), unifieddatalibrary.SiteStatusCountParams{
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

func TestSiteStatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SiteStatus.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteStatusGetParams{
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

func TestSiteStatusQueryhelp(t *testing.T) {
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
	err := client.SiteStatus.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteStatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SiteStatus.Tuple(context.TODO(), unifieddatalibrary.SiteStatusTupleParams{
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
