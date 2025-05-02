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

func TestSiteNewWithOptionalParams(t *testing.T) {
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
	err := client.Site.New(context.TODO(), unifieddatalibrary.SiteNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SiteNewParamsDataModeTest,
		Name:                  "Site Name",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SITE-ID"),
		Activity:              unifieddatalibrary.String("OCC"),
		AirDefArea:            unifieddatalibrary.String("AL006"),
		Allegiance:            unifieddatalibrary.String("OTHR"),
		AltAllegiance:         unifieddatalibrary.String("HL"),
		BeNumber:              unifieddatalibrary.String("0427RT1030"),
		CatCode:               unifieddatalibrary.String("20345"),
		CatText:               unifieddatalibrary.String("Radar Facility, General"),
		ClassRating:           unifieddatalibrary.String("1"),
		Condition:             unifieddatalibrary.String("RDY"),
		ConditionAvail:        unifieddatalibrary.String("A"),
		Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
		CoordDatum:            unifieddatalibrary.String("WGS"),
		CoordDerivAcc:         unifieddatalibrary.Float(12.345),
		ElevMsl:               unifieddatalibrary.Float(123.45),
		ElevMslConfLvl:        unifieddatalibrary.Int(50),
		ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
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
		Eval:            unifieddatalibrary.Int(7),
		Faa:             unifieddatalibrary.String("FAA1"),
		Fpa:             unifieddatalibrary.String("EOB"),
		FunctPrimary:    unifieddatalibrary.String("JG"),
		GeoArea:         unifieddatalibrary.String("E2"),
		GeoidalMslSep:   unifieddatalibrary.Float(12.34),
		Grade:           unifieddatalibrary.Int(5),
		Iata:            unifieddatalibrary.String("AAA"),
		Icao:            unifieddatalibrary.String("ICA1"),
		Ident:           unifieddatalibrary.String("FRIEND"),
		IDEntity:        unifieddatalibrary.String("ENTITY-ID"),
		IDParentSite:    unifieddatalibrary.String("ID-Parent-Site"),
		LzUsage:         unifieddatalibrary.String("AF"),
		MaxRunwayLength: unifieddatalibrary.Int(1000),
		MilGrid:         unifieddatalibrary.String("4QFJ12345678"),
		MilGridSys:      unifieddatalibrary.String("UTM"),
		MsnPrimary:      unifieddatalibrary.String("AA"),
		MsnPrimarySpec:  unifieddatalibrary.String("AB"),
		Notes:           unifieddatalibrary.String("Example Notes"),
		NucCap:          unifieddatalibrary.String("A"),
		OperStatus:      unifieddatalibrary.String("OPR"),
		Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigLzID:        unifieddatalibrary.String("ORIG-LZ-ID"),
		OrigSiteID:      unifieddatalibrary.String("ORIG-SITE-ID"),
		Osuffix:         unifieddatalibrary.String("BB002"),
		Pin:             unifieddatalibrary.String("25200"),
		PolSubdiv:       unifieddatalibrary.String("IZO7"),
		PopArea:         unifieddatalibrary.Bool(true),
		PopAreaProx:     unifieddatalibrary.Float(12.345),
		RecStatus:       unifieddatalibrary.String("A"),
		ReferenceDoc:    unifieddatalibrary.String("Provider Reference Documentation"),
		ResProd:         unifieddatalibrary.String("RT"),
		ReviewDate:      unifieddatalibrary.Time(time.Now()),
		Runways:         unifieddatalibrary.Int(5),
		SymCode:         unifieddatalibrary.String("SOGPU----------"),
		Type:            unifieddatalibrary.String("AIRBASE"),
		Usage:           unifieddatalibrary.String("MILITARY"),
		Utm:             unifieddatalibrary.String("19P4390691376966"),
		VegHt:           unifieddatalibrary.Float(3),
		VegType:         unifieddatalibrary.String("FR"),
		Wac:             unifieddatalibrary.String("0427"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Site.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SiteUpdateParamsDataModeTest,
			Name:                  "Site Name",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SITE-ID"),
			Activity:              unifieddatalibrary.String("OCC"),
			AirDefArea:            unifieddatalibrary.String("AL006"),
			Allegiance:            unifieddatalibrary.String("OTHR"),
			AltAllegiance:         unifieddatalibrary.String("HL"),
			BeNumber:              unifieddatalibrary.String("0427RT1030"),
			CatCode:               unifieddatalibrary.String("20345"),
			CatText:               unifieddatalibrary.String("Radar Facility, General"),
			ClassRating:           unifieddatalibrary.String("1"),
			Condition:             unifieddatalibrary.String("RDY"),
			ConditionAvail:        unifieddatalibrary.String("A"),
			Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
			CoordDatum:            unifieddatalibrary.String("WGS"),
			CoordDerivAcc:         unifieddatalibrary.Float(12.345),
			ElevMsl:               unifieddatalibrary.Float(123.45),
			ElevMslConfLvl:        unifieddatalibrary.Int(50),
			ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
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
			Eval:            unifieddatalibrary.Int(7),
			Faa:             unifieddatalibrary.String("FAA1"),
			Fpa:             unifieddatalibrary.String("EOB"),
			FunctPrimary:    unifieddatalibrary.String("JG"),
			GeoArea:         unifieddatalibrary.String("E2"),
			GeoidalMslSep:   unifieddatalibrary.Float(12.34),
			Grade:           unifieddatalibrary.Int(5),
			Iata:            unifieddatalibrary.String("AAA"),
			Icao:            unifieddatalibrary.String("ICA1"),
			Ident:           unifieddatalibrary.String("FRIEND"),
			IDEntity:        unifieddatalibrary.String("ENTITY-ID"),
			IDParentSite:    unifieddatalibrary.String("ID-Parent-Site"),
			LzUsage:         unifieddatalibrary.String("AF"),
			MaxRunwayLength: unifieddatalibrary.Int(1000),
			MilGrid:         unifieddatalibrary.String("4QFJ12345678"),
			MilGridSys:      unifieddatalibrary.String("UTM"),
			MsnPrimary:      unifieddatalibrary.String("AA"),
			MsnPrimarySpec:  unifieddatalibrary.String("AB"),
			Notes:           unifieddatalibrary.String("Example Notes"),
			NucCap:          unifieddatalibrary.String("A"),
			OperStatus:      unifieddatalibrary.String("OPR"),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigLzID:        unifieddatalibrary.String("ORIG-LZ-ID"),
			OrigSiteID:      unifieddatalibrary.String("ORIG-SITE-ID"),
			Osuffix:         unifieddatalibrary.String("BB002"),
			Pin:             unifieddatalibrary.String("25200"),
			PolSubdiv:       unifieddatalibrary.String("IZO7"),
			PopArea:         unifieddatalibrary.Bool(true),
			PopAreaProx:     unifieddatalibrary.Float(12.345),
			RecStatus:       unifieddatalibrary.String("A"),
			ReferenceDoc:    unifieddatalibrary.String("Provider Reference Documentation"),
			ResProd:         unifieddatalibrary.String("RT"),
			ReviewDate:      unifieddatalibrary.Time(time.Now()),
			Runways:         unifieddatalibrary.Int(5),
			SymCode:         unifieddatalibrary.String("SOGPU----------"),
			Type:            unifieddatalibrary.String("AIRBASE"),
			Usage:           unifieddatalibrary.String("MILITARY"),
			Utm:             unifieddatalibrary.String("19P4390691376966"),
			VegHt:           unifieddatalibrary.Float(3),
			VegType:         unifieddatalibrary.String("FR"),
			Wac:             unifieddatalibrary.String("0427"),
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

func TestSiteListWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.List(context.TODO(), unifieddatalibrary.SiteListParams{
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

func TestSiteCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Count(context.TODO(), unifieddatalibrary.SiteCountParams{
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

func TestSiteGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteGetParams{
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

func TestSiteQueryhelp(t *testing.T) {
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
	err := client.Site.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Tuple(context.TODO(), unifieddatalibrary.SiteTupleParams{
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
