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

func TestBusNewWithOptionalParams(t *testing.T) {
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
	err := client.Buses.New(context.TODO(), unifieddatalibrary.BusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.BusNewParamsDataModeTest,
		Name:                  "Example name",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("BUS-ID"),
		AocsNotes:             unifieddatalibrary.String("Example notes"),
		AvgDryMass:            unifieddatalibrary.Float(2879.1),
		AvgPayloadMass:        unifieddatalibrary.Float(10.1),
		AvgPayloadPower:       unifieddatalibrary.Float(10.1),
		AvgSpacecraftPower:    unifieddatalibrary.Float(10.1),
		AvgWetMass:            unifieddatalibrary.Float(5246.1),
		BodyDimensionX:        unifieddatalibrary.Float(10.1),
		BodyDimensionY:        unifieddatalibrary.Float(10.1),
		BodyDimensionZ:        unifieddatalibrary.Float(10.1),
		BusKitDesignerOrgID:   unifieddatalibrary.String("BUSKITDESIGNERORG-ID"),
		CountryCode:           unifieddatalibrary.String("US"),
		Description:           unifieddatalibrary.String("Dedicated small spacecraft bus."),
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
		Generic:                            unifieddatalibrary.Bool(false),
		IDEntity:                           unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		LaunchEnvelopeDimensionX:           unifieddatalibrary.Float(10.1),
		LaunchEnvelopeDimensionY:           unifieddatalibrary.Float(10.1),
		LaunchEnvelopeDimensionZ:           unifieddatalibrary.Float(10.1),
		MainComputerManufacturerOrgID:      unifieddatalibrary.String("MAINCOMPUTERMANUFACTURERORG-ID"),
		ManufacturerOrgID:                  unifieddatalibrary.String("MANUFACTURERORG-ID"),
		MassCategory:                       unifieddatalibrary.String("Nanosatellite"),
		MaxBolPowerLower:                   unifieddatalibrary.Float(10.1),
		MaxBolPowerUpper:                   unifieddatalibrary.Float(10.1),
		MaxBolStationMass:                  unifieddatalibrary.Float(10.1),
		MaxDryMass:                         unifieddatalibrary.Float(2900.1),
		MaxEolPowerLower:                   unifieddatalibrary.Float(10.1),
		MaxEolPowerUpper:                   unifieddatalibrary.Float(10.1),
		MaxLaunchMassLower:                 unifieddatalibrary.Float(10.1),
		MaxLaunchMassUpper:                 unifieddatalibrary.Float(10.1),
		MaxPayloadMass:                     unifieddatalibrary.Float(10.1),
		MaxPayloadPower:                    unifieddatalibrary.Float(10.1),
		MaxSpacecraftPower:                 unifieddatalibrary.Float(10.1),
		MaxWetMass:                         unifieddatalibrary.Float(5300),
		MedianDryMass:                      unifieddatalibrary.Float(2950.1),
		MedianWetMass:                      unifieddatalibrary.Float(5260.1),
		MinDryMass:                         unifieddatalibrary.Float(2858.1),
		MinWetMass:                         unifieddatalibrary.Float(5192.1),
		NumOrbitType:                       unifieddatalibrary.Int(3),
		OapPayloadPower:                    unifieddatalibrary.Float(10.1),
		OapSpacecraftPower:                 unifieddatalibrary.Float(10.1),
		OrbitTypes:                         []string{"LEO", "HEO", "GEO"},
		Origin:                             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PayloadDimensionX:                  unifieddatalibrary.Float(1.1),
		PayloadDimensionY:                  unifieddatalibrary.Float(1.1),
		PayloadDimensionZ:                  unifieddatalibrary.Float(1.1),
		PayloadVolume:                      unifieddatalibrary.Float(1.1),
		PowerCategory:                      unifieddatalibrary.String("low power"),
		TelemetryTrackingManufacturerOrgID: unifieddatalibrary.String("TELEMETRYTRACKINGMANUFACTURERORG-ID"),
		Type:                               unifieddatalibrary.String("Example type"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Buses.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.BusGetParams{
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

func TestBusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Buses.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.BusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.BusUpdateParamsDataModeTest,
			Name:                  "Example name",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("BUS-ID"),
			AocsNotes:             unifieddatalibrary.String("Example notes"),
			AvgDryMass:            unifieddatalibrary.Float(2879.1),
			AvgPayloadMass:        unifieddatalibrary.Float(10.1),
			AvgPayloadPower:       unifieddatalibrary.Float(10.1),
			AvgSpacecraftPower:    unifieddatalibrary.Float(10.1),
			AvgWetMass:            unifieddatalibrary.Float(5246.1),
			BodyDimensionX:        unifieddatalibrary.Float(10.1),
			BodyDimensionY:        unifieddatalibrary.Float(10.1),
			BodyDimensionZ:        unifieddatalibrary.Float(10.1),
			BusKitDesignerOrgID:   unifieddatalibrary.String("BUSKITDESIGNERORG-ID"),
			CountryCode:           unifieddatalibrary.String("US"),
			Description:           unifieddatalibrary.String("Dedicated small spacecraft bus."),
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
			Generic:                            unifieddatalibrary.Bool(false),
			IDEntity:                           unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			LaunchEnvelopeDimensionX:           unifieddatalibrary.Float(10.1),
			LaunchEnvelopeDimensionY:           unifieddatalibrary.Float(10.1),
			LaunchEnvelopeDimensionZ:           unifieddatalibrary.Float(10.1),
			MainComputerManufacturerOrgID:      unifieddatalibrary.String("MAINCOMPUTERMANUFACTURERORG-ID"),
			ManufacturerOrgID:                  unifieddatalibrary.String("MANUFACTURERORG-ID"),
			MassCategory:                       unifieddatalibrary.String("Nanosatellite"),
			MaxBolPowerLower:                   unifieddatalibrary.Float(10.1),
			MaxBolPowerUpper:                   unifieddatalibrary.Float(10.1),
			MaxBolStationMass:                  unifieddatalibrary.Float(10.1),
			MaxDryMass:                         unifieddatalibrary.Float(2900.1),
			MaxEolPowerLower:                   unifieddatalibrary.Float(10.1),
			MaxEolPowerUpper:                   unifieddatalibrary.Float(10.1),
			MaxLaunchMassLower:                 unifieddatalibrary.Float(10.1),
			MaxLaunchMassUpper:                 unifieddatalibrary.Float(10.1),
			MaxPayloadMass:                     unifieddatalibrary.Float(10.1),
			MaxPayloadPower:                    unifieddatalibrary.Float(10.1),
			MaxSpacecraftPower:                 unifieddatalibrary.Float(10.1),
			MaxWetMass:                         unifieddatalibrary.Float(5300),
			MedianDryMass:                      unifieddatalibrary.Float(2950.1),
			MedianWetMass:                      unifieddatalibrary.Float(5260.1),
			MinDryMass:                         unifieddatalibrary.Float(2858.1),
			MinWetMass:                         unifieddatalibrary.Float(5192.1),
			NumOrbitType:                       unifieddatalibrary.Int(3),
			OapPayloadPower:                    unifieddatalibrary.Float(10.1),
			OapSpacecraftPower:                 unifieddatalibrary.Float(10.1),
			OrbitTypes:                         []string{"LEO", "HEO", "GEO"},
			Origin:                             unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PayloadDimensionX:                  unifieddatalibrary.Float(1.1),
			PayloadDimensionY:                  unifieddatalibrary.Float(1.1),
			PayloadDimensionZ:                  unifieddatalibrary.Float(1.1),
			PayloadVolume:                      unifieddatalibrary.Float(1.1),
			PowerCategory:                      unifieddatalibrary.String("low power"),
			TelemetryTrackingManufacturerOrgID: unifieddatalibrary.String("TELEMETRYTRACKINGMANUFACTURERORG-ID"),
			Type:                               unifieddatalibrary.String("Example type"),
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

func TestBusListWithOptionalParams(t *testing.T) {
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
	_, err := client.Buses.List(context.TODO(), unifieddatalibrary.BusListParams{
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

func TestBusDelete(t *testing.T) {
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
	err := client.Buses.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Buses.Count(context.TODO(), unifieddatalibrary.BusCountParams{
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

func TestBusQueryHelp(t *testing.T) {
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
	err := client.Buses.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Buses.Tuple(context.TODO(), unifieddatalibrary.BusTupleParams{
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
