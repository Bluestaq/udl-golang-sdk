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

func TestAirfieldNewWithOptionalParams(t *testing.T) {
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
	err := client.Airfields.New(context.TODO(), unifieddatalibrary.AirfieldNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirfieldNewParamsDataModeTest,
		Name:                  "Hickam Air Force Base",
		Source:                "Bluestaq",
		Type:                  "Commercial",
		ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
		AltAirfieldID:         unifieddatalibrary.String("45301"),
		AlternativeNames:      []string{"BELLEVILLE", "JONESTOWN"},
		City:                  unifieddatalibrary.String("Honolulu"),
		CountryCode:           unifieddatalibrary.String("US"),
		CountryName:           unifieddatalibrary.String("United States"),
		DstInfo:               unifieddatalibrary.String("SEE THE ENROUTE SUPP FOR INFORMATION"),
		ElevFt:                unifieddatalibrary.Float(33.562),
		ElevM:                 unifieddatalibrary.Float(10.29),
		Faa:                   unifieddatalibrary.String("FAA1"),
		Geoloc:                unifieddatalibrary.String("XLSX"),
		GmtOffset:             unifieddatalibrary.String("-4:30"),
		HostNatCode:           unifieddatalibrary.String("ZPU"),
		Iata:                  unifieddatalibrary.String("AAA"),
		Icao:                  unifieddatalibrary.String("KCOS"),
		IDSite:                unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
		InfoURL:               unifieddatalibrary.String("URL Link to the Airfield"),
		Lat:                   unifieddatalibrary.Float(45.23),
		Lon:                   unifieddatalibrary.Float(179.1),
		MagDec:                unifieddatalibrary.Float(7.35),
		MaxRunwayLength:       unifieddatalibrary.Int(1000),
		MiscCodes:             unifieddatalibrary.String("AMZ"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		RegionalAuthority:     unifieddatalibrary.String("18TH AF"),
		RegionName:            unifieddatalibrary.String("Hawaii"),
		Runways:               unifieddatalibrary.Int(5),
		SecondaryIcao:         unifieddatalibrary.String("PHNL"),
		State:                 unifieddatalibrary.String("Hawaii"),
		StateProvinceCode:     unifieddatalibrary.String("US15"),
		SuitabilityCodeDescs:  []string{"Suitable C-32", "Suitable C-5", "Suitable C-130"},
		SuitabilityCodes:      unifieddatalibrary.String("ABC"),
		WacInnr:               unifieddatalibrary.String("0409-00039"),
		ZarID:                 unifieddatalibrary.String("231"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfields.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldGetParams{
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

func TestAirfieldUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Airfields.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirfieldUpdateParamsDataModeTest,
			Name:                  "Hickam Air Force Base",
			Source:                "Bluestaq",
			Type:                  "Commercial",
			ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			AltAirfieldID:         unifieddatalibrary.String("45301"),
			AlternativeNames:      []string{"BELLEVILLE", "JONESTOWN"},
			City:                  unifieddatalibrary.String("Honolulu"),
			CountryCode:           unifieddatalibrary.String("US"),
			CountryName:           unifieddatalibrary.String("United States"),
			DstInfo:               unifieddatalibrary.String("SEE THE ENROUTE SUPP FOR INFORMATION"),
			ElevFt:                unifieddatalibrary.Float(33.562),
			ElevM:                 unifieddatalibrary.Float(10.29),
			Faa:                   unifieddatalibrary.String("FAA1"),
			Geoloc:                unifieddatalibrary.String("XLSX"),
			GmtOffset:             unifieddatalibrary.String("-4:30"),
			HostNatCode:           unifieddatalibrary.String("ZPU"),
			Iata:                  unifieddatalibrary.String("AAA"),
			Icao:                  unifieddatalibrary.String("KCOS"),
			IDSite:                unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			InfoURL:               unifieddatalibrary.String("URL Link to the Airfield"),
			Lat:                   unifieddatalibrary.Float(45.23),
			Lon:                   unifieddatalibrary.Float(179.1),
			MagDec:                unifieddatalibrary.Float(7.35),
			MaxRunwayLength:       unifieddatalibrary.Int(1000),
			MiscCodes:             unifieddatalibrary.String("AMZ"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			RegionalAuthority:     unifieddatalibrary.String("18TH AF"),
			RegionName:            unifieddatalibrary.String("Hawaii"),
			Runways:               unifieddatalibrary.Int(5),
			SecondaryIcao:         unifieddatalibrary.String("PHNL"),
			State:                 unifieddatalibrary.String("Hawaii"),
			StateProvinceCode:     unifieddatalibrary.String("US15"),
			SuitabilityCodeDescs:  []string{"Suitable C-32", "Suitable C-5", "Suitable C-130"},
			SuitabilityCodes:      unifieddatalibrary.String("ABC"),
			WacInnr:               unifieddatalibrary.String("0409-00039"),
			ZarID:                 unifieddatalibrary.String("231"),
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

func TestAirfieldListWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfields.List(context.TODO(), unifieddatalibrary.AirfieldListParams{
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

func TestAirfieldCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfields.Count(context.TODO(), unifieddatalibrary.AirfieldCountParams{
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

func TestAirfieldQueryhelp(t *testing.T) {
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
	err := client.Airfields.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfields.Tuple(context.TODO(), unifieddatalibrary.AirfieldTupleParams{
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
