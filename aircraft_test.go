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

func TestAircraftNewWithOptionalParams(t *testing.T) {
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
	err := client.Aircraft.New(context.TODO(), unifieddatalibrary.AircraftNewParams{
		AircraftMds:           "E-2C HAWKEYE",
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AircraftNewParamsDataModeTest,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		Category:              unifieddatalibrary.String("M"),
		Command:               unifieddatalibrary.String("HQACC"),
		CruiseSpeed:           unifieddatalibrary.Float(915),
		Dtd:                   unifieddatalibrary.String("005"),
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
		IDEntity:       unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		MaxSpeed:       unifieddatalibrary.Float(2655.1),
		MinReqRunwayFt: unifieddatalibrary.Int(3000),
		MinReqRunwayM:  unifieddatalibrary.Int(1000),
		NominalTaTime:  unifieddatalibrary.Int(500),
		Notes:          unifieddatalibrary.String("Notes for this aircraft"),
		Origin:         unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Owner:          unifieddatalibrary.String("437AW"),
		SerialNumber:   unifieddatalibrary.String("7007187"),
		TailNumber:     unifieddatalibrary.String("N702JG"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAircraftGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Aircraft.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AircraftGetParams{
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

func TestAircraftUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Aircraft.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AircraftUpdateParams{
			AircraftMds:           "E-2C HAWKEYE",
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AircraftUpdateParamsDataModeTest,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			Category:              unifieddatalibrary.String("M"),
			Command:               unifieddatalibrary.String("HQACC"),
			CruiseSpeed:           unifieddatalibrary.Float(915),
			Dtd:                   unifieddatalibrary.String("005"),
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
			IDEntity:       unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			MaxSpeed:       unifieddatalibrary.Float(2655.1),
			MinReqRunwayFt: unifieddatalibrary.Int(3000),
			MinReqRunwayM:  unifieddatalibrary.Int(1000),
			NominalTaTime:  unifieddatalibrary.Int(500),
			Notes:          unifieddatalibrary.String("Notes for this aircraft"),
			Origin:         unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Owner:          unifieddatalibrary.String("437AW"),
			SerialNumber:   unifieddatalibrary.String("7007187"),
			TailNumber:     unifieddatalibrary.String("N702JG"),
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

func TestAircraftListWithOptionalParams(t *testing.T) {
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
	_, err := client.Aircraft.List(context.TODO(), unifieddatalibrary.AircraftListParams{
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

func TestAircraftCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Aircraft.Count(context.TODO(), unifieddatalibrary.AircraftCountParams{
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

func TestAircraftQueryhelp(t *testing.T) {
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
	_, err := client.Aircraft.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAircraftTupleQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Aircraft.TupleQuery(context.TODO(), unifieddatalibrary.AircraftTupleQueryParams{
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
