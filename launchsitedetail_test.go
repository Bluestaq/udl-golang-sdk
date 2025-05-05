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

func TestLaunchSiteDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.LaunchSiteDetails.New(context.TODO(), unifieddatalibrary.LaunchSiteDetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LaunchSiteDetailNewParamsDataModeTest,
		IDLaunchSite:          "LAUNCHSITE-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("LAUNCHSITEDETAILS-ID"),
		AvailableInclinations: []float64{10.23, 10.23, 12.23, 14.23},
		Description:           unifieddatalibrary.String("Example notes"),
		IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
		LaunchGroup:           unifieddatalibrary.String("Example-group-name"),
		Location: unifieddatalibrary.LocationIngestParam{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LocationIngestDataModeTest,
			Name:                  "Example location",
			Source:                "Bluestaq",
			Altitude:              unifieddatalibrary.Float(10.23),
			CountryCode:           unifieddatalibrary.String("US"),
			IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
			Lat:                   unifieddatalibrary.Float(45.23),
			Lon:                   unifieddatalibrary.Float(179.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		},
		Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchSiteDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.LaunchSiteDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchSiteDetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LaunchSiteDetailUpdateParamsDataModeTest,
			IDLaunchSite:          "LAUNCHSITE-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LAUNCHSITEDETAILS-ID"),
			AvailableInclinations: []float64{10.23, 10.23, 12.23, 14.23},
			Description:           unifieddatalibrary.String("Example notes"),
			IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
			LaunchGroup:           unifieddatalibrary.String("Example-group-name"),
			Location: unifieddatalibrary.LocationIngestParam{
				ClassificationMarking: "U",
				DataMode:              unifieddatalibrary.LocationIngestDataModeTest,
				Name:                  "Example location",
				Source:                "Bluestaq",
				Altitude:              unifieddatalibrary.Float(10.23),
				CountryCode:           unifieddatalibrary.String("US"),
				IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
				Lat:                   unifieddatalibrary.Float(45.23),
				Lon:                   unifieddatalibrary.Float(179.1),
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestLaunchSiteDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchSiteDetails.List(context.TODO(), unifieddatalibrary.LaunchSiteDetailListParams{
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

func TestLaunchSiteDetailDelete(t *testing.T) {
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
	err := client.LaunchSiteDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchSiteDetailFindBySourceWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchSiteDetails.FindBySource(context.TODO(), unifieddatalibrary.LaunchSiteDetailFindBySourceParams{
		Source:      "source",
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

func TestLaunchSiteDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchSiteDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchSiteDetailGetParams{
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
