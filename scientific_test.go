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

func TestScientificNewWithOptionalParams(t *testing.T) {
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
	err := client.Scientific.New(context.TODO(), unifieddatalibrary.ScientificNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ScientificNewParamsDataModeTest,
		Name:                  "SEM/MAG",
		Source:                "Bluestaq",
		SpacecraftID:          "REF-SPACECRAFT-ID",
		ID:                    unifieddatalibrary.String("SCIENTIFIC-ID"),
		Entity: unifieddatalibrary.EntityIngestParam{
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
		FrequencyBand:         unifieddatalibrary.String("Gamma"),
		HostedForCompanyOrgID: unifieddatalibrary.String("REF-HOSTEDFORCOMPANYORG-ID"),
		IDEntity:              unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		ManufacturerOrgID:     unifieddatalibrary.String("REF-MANUFACTURERORG-ID"),
		Notes:                 unifieddatalibrary.String("NOTES"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PayloadCategory:       unifieddatalibrary.String("Sensor"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScientificUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Scientific.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.ScientificUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ScientificUpdateParamsDataModeTest,
			Name:                  "SEM/MAG",
			Source:                "Bluestaq",
			SpacecraftID:          "REF-SPACECRAFT-ID",
			ID:                    unifieddatalibrary.String("SCIENTIFIC-ID"),
			Entity: unifieddatalibrary.EntityIngestParam{
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
			FrequencyBand:         unifieddatalibrary.String("Gamma"),
			HostedForCompanyOrgID: unifieddatalibrary.String("REF-HOSTEDFORCOMPANYORG-ID"),
			IDEntity:              unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			ManufacturerOrgID:     unifieddatalibrary.String("REF-MANUFACTURERORG-ID"),
			Notes:                 unifieddatalibrary.String("NOTES"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PayloadCategory:       unifieddatalibrary.String("Sensor"),
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

func TestScientificListWithOptionalParams(t *testing.T) {
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
	_, err := client.Scientific.List(context.TODO(), unifieddatalibrary.ScientificListParams{
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

func TestScientificDelete(t *testing.T) {
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
	err := client.Scientific.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScientificCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Scientific.Count(context.TODO(), unifieddatalibrary.ScientificCountParams{
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

func TestScientificGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Scientific.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ScientificGetParams{
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

func TestScientificQueryhelp(t *testing.T) {
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
	_, err := client.Scientific.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScientificTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Scientific.Tuple(context.TODO(), unifieddatalibrary.ScientificTupleParams{
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
