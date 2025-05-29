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

func TestEntityNewWithOptionalParams(t *testing.T) {
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
	err := client.Entities.New(context.TODO(), unifieddatalibrary.EntityNewParams{
		EntityIngest: unifieddatalibrary.EntityIngestParam{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.EntityIngestDataModeTest,
			Name:                  "Example name",
			Source:                "Bluestaq",
			Type:                  unifieddatalibrary.EntityIngestTypeOnorbit,
			CountryCode:           unifieddatalibrary.String("US"),
			IDEntity:              unifieddatalibrary.String("ENTITY-ID"),
			IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
			IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
			IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
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
			OnOrbit: unifieddatalibrary.EntityIngestOnOrbitParam{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				SatNo:                 1,
				Source:                "Bluestaq",
				AltName:               unifieddatalibrary.String("Alternate Name"),
				Category:              "Lunar",
				CommonName:            unifieddatalibrary.String("Example common name"),
				Constellation:         unifieddatalibrary.String("Big Dipper"),
				CountryCode:           unifieddatalibrary.String("US"),
				DecayDate:             unifieddatalibrary.Time(time.Now()),
				IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
				IntlDes:               unifieddatalibrary.String("2021123ABC"),
				LaunchDate:            unifieddatalibrary.Time(time.Now()),
				LaunchSiteID:          unifieddatalibrary.String("LAUNCHSITE-ID"),
				LifetimeYears:         unifieddatalibrary.Int(10),
				MissionNumber:         unifieddatalibrary.String("Expedition 1"),
				ObjectType:            "PAYLOAD",
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			Origin:    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OwnerType: unifieddatalibrary.EntityIngestOwnerTypeCommercial,
			Taskable:  unifieddatalibrary.Bool(false),
			URLs:      []string{"URL1", "URL2"},
		},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EntityGetParams{
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

func TestEntityUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Entities.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.EntityUpdateParams{
			EntityIngest: unifieddatalibrary.EntityIngestParam{
				ClassificationMarking: "U",
				DataMode:              unifieddatalibrary.EntityIngestDataModeTest,
				Name:                  "Example name",
				Source:                "Bluestaq",
				Type:                  unifieddatalibrary.EntityIngestTypeOnorbit,
				CountryCode:           unifieddatalibrary.String("US"),
				IDEntity:              unifieddatalibrary.String("ENTITY-ID"),
				IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
				IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
				IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
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
				OnOrbit: unifieddatalibrary.EntityIngestOnOrbitParam{
					ClassificationMarking: "U",
					DataMode:              "TEST",
					SatNo:                 1,
					Source:                "Bluestaq",
					AltName:               unifieddatalibrary.String("Alternate Name"),
					Category:              "Lunar",
					CommonName:            unifieddatalibrary.String("Example common name"),
					Constellation:         unifieddatalibrary.String("Big Dipper"),
					CountryCode:           unifieddatalibrary.String("US"),
					DecayDate:             unifieddatalibrary.Time(time.Now()),
					IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
					IntlDes:               unifieddatalibrary.String("2021123ABC"),
					LaunchDate:            unifieddatalibrary.Time(time.Now()),
					LaunchSiteID:          unifieddatalibrary.String("LAUNCHSITE-ID"),
					LifetimeYears:         unifieddatalibrary.Int(10),
					MissionNumber:         unifieddatalibrary.String("Expedition 1"),
					ObjectType:            "PAYLOAD",
					Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				},
				Origin:    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				OwnerType: unifieddatalibrary.EntityIngestOwnerTypeCommercial,
				Taskable:  unifieddatalibrary.Bool(false),
				URLs:      []string{"URL1", "URL2"},
			},
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

func TestEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.List(context.TODO(), unifieddatalibrary.EntityListParams{
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

func TestEntityDelete(t *testing.T) {
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
	err := client.Entities.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Count(context.TODO(), unifieddatalibrary.EntityCountParams{
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

func TestEntityGetAllTypesWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.GetAllTypes(context.TODO(), unifieddatalibrary.EntityGetAllTypesParams{
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

func TestEntityQueryHelp(t *testing.T) {
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
	_, err := client.Entities.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Tuple(context.TODO(), unifieddatalibrary.EntityTupleParams{
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
