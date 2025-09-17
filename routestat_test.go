// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestRouteStatNewWithOptionalParams(t *testing.T) {
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
	err := client.RouteStats.New(context.TODO(), unifieddatalibrary.RouteStatNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RouteStatNewParamsDataModeTest,
		LocationEnd:           "KCOS",
		LocationStart:         "KDEN",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		AvgDuration:           unifieddatalibrary.Float(47.1),
		AvgSpeed:              unifieddatalibrary.Float(450.1),
		DataPtsUsed:           unifieddatalibrary.Int(6),
		Distance:              unifieddatalibrary.Float(63.1),
		DistUnit:              unifieddatalibrary.String("Nautical miles"),
		FirstPt:               unifieddatalibrary.Time(time.Now()),
		IdealDesc:             unifieddatalibrary.String("Block speed using great circle path"),
		IdealDuration:         unifieddatalibrary.Float(45.1),
		IDSiteEnd:             unifieddatalibrary.String("77b5550c-c0f4-47bd-94ce-d71cdaa52f62"),
		IDSiteStart:           unifieddatalibrary.String("23370387-5e8e-4a74-89db-ad81145aa4df"),
		LastPt:                unifieddatalibrary.Time(time.Now()),
		LocationType:          unifieddatalibrary.String("ICAO"),
		MaxDuration:           unifieddatalibrary.Float(52.1),
		MaxSpeed:              unifieddatalibrary.Float(470.1),
		MinDuration:           unifieddatalibrary.Float(42.1),
		MinSpeed:              unifieddatalibrary.Float(420.1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PartialDesc:           unifieddatalibrary.String("Performance speed using great circle path"),
		PartialDuration:       unifieddatalibrary.Float(38.1),
		SpeedUnit:             unifieddatalibrary.String("knots"),
		TimePeriod:            unifieddatalibrary.String("Q1"),
		VehicleCategory:       unifieddatalibrary.String("AIRCRAFT"),
		VehicleType:           unifieddatalibrary.String("C-17"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRouteStatGetWithOptionalParams(t *testing.T) {
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
	_, err := client.RouteStats.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RouteStatGetParams{
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

func TestRouteStatUpdateWithOptionalParams(t *testing.T) {
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
	err := client.RouteStats.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RouteStatUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.RouteStatUpdateParamsDataModeTest,
			LocationEnd:           "KCOS",
			LocationStart:         "KDEN",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AvgDuration:           unifieddatalibrary.Float(47.1),
			AvgSpeed:              unifieddatalibrary.Float(450.1),
			DataPtsUsed:           unifieddatalibrary.Int(6),
			Distance:              unifieddatalibrary.Float(63.1),
			DistUnit:              unifieddatalibrary.String("Nautical miles"),
			FirstPt:               unifieddatalibrary.Time(time.Now()),
			IdealDesc:             unifieddatalibrary.String("Block speed using great circle path"),
			IdealDuration:         unifieddatalibrary.Float(45.1),
			IDSiteEnd:             unifieddatalibrary.String("77b5550c-c0f4-47bd-94ce-d71cdaa52f62"),
			IDSiteStart:           unifieddatalibrary.String("23370387-5e8e-4a74-89db-ad81145aa4df"),
			LastPt:                unifieddatalibrary.Time(time.Now()),
			LocationType:          unifieddatalibrary.String("ICAO"),
			MaxDuration:           unifieddatalibrary.Float(52.1),
			MaxSpeed:              unifieddatalibrary.Float(470.1),
			MinDuration:           unifieddatalibrary.Float(42.1),
			MinSpeed:              unifieddatalibrary.Float(420.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartialDesc:           unifieddatalibrary.String("Performance speed using great circle path"),
			PartialDuration:       unifieddatalibrary.Float(38.1),
			SpeedUnit:             unifieddatalibrary.String("knots"),
			TimePeriod:            unifieddatalibrary.String("Q1"),
			VehicleCategory:       unifieddatalibrary.String("AIRCRAFT"),
			VehicleType:           unifieddatalibrary.String("C-17"),
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

func TestRouteStatListWithOptionalParams(t *testing.T) {
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
	_, err := client.RouteStats.List(context.TODO(), unifieddatalibrary.RouteStatListParams{
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

func TestRouteStatDelete(t *testing.T) {
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
	err := client.RouteStats.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRouteStatCountWithOptionalParams(t *testing.T) {
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
	_, err := client.RouteStats.Count(context.TODO(), unifieddatalibrary.RouteStatCountParams{
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

func TestRouteStatNewBulk(t *testing.T) {
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
	err := client.RouteStats.NewBulk(context.TODO(), unifieddatalibrary.RouteStatNewBulkParams{
		Body: []unifieddatalibrary.RouteStatNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			LocationEnd:           "KCOS",
			LocationStart:         "KDEN",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AvgDuration:           unifieddatalibrary.Float(47.1),
			AvgSpeed:              unifieddatalibrary.Float(450.1),
			DataPtsUsed:           unifieddatalibrary.Int(6),
			Distance:              unifieddatalibrary.Float(63.1),
			DistUnit:              unifieddatalibrary.String("Nautical miles"),
			FirstPt:               unifieddatalibrary.Time(time.Now()),
			IdealDesc:             unifieddatalibrary.String("Block speed using great circle path"),
			IdealDuration:         unifieddatalibrary.Float(45.1),
			IDSiteEnd:             unifieddatalibrary.String("77b5550c-c0f4-47bd-94ce-d71cdaa52f62"),
			IDSiteStart:           unifieddatalibrary.String("23370387-5e8e-4a74-89db-ad81145aa4df"),
			LastPt:                unifieddatalibrary.Time(time.Now()),
			LocationType:          unifieddatalibrary.String("ICAO"),
			MaxDuration:           unifieddatalibrary.Float(52.1),
			MaxSpeed:              unifieddatalibrary.Float(470.1),
			MinDuration:           unifieddatalibrary.Float(42.1),
			MinSpeed:              unifieddatalibrary.Float(420.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartialDesc:           unifieddatalibrary.String("Performance speed using great circle path"),
			PartialDuration:       unifieddatalibrary.Float(38.1),
			SpeedUnit:             unifieddatalibrary.String("knots"),
			TimePeriod:            unifieddatalibrary.String("Q1"),
			VehicleCategory:       unifieddatalibrary.String("AIRCRAFT"),
			VehicleType:           unifieddatalibrary.String("C-17"),
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRouteStatQueryHelp(t *testing.T) {
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
	_, err := client.RouteStats.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRouteStatTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.RouteStats.Tuple(context.TODO(), unifieddatalibrary.RouteStatTupleParams{
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

func TestRouteStatUnvalidatedPublish(t *testing.T) {
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
	err := client.RouteStats.UnvalidatedPublish(context.TODO(), unifieddatalibrary.RouteStatUnvalidatedPublishParams{
		Body: []unifieddatalibrary.RouteStatUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			LocationEnd:           "KCOS",
			LocationStart:         "KDEN",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AvgDuration:           unifieddatalibrary.Float(47.1),
			AvgSpeed:              unifieddatalibrary.Float(450.1),
			DataPtsUsed:           unifieddatalibrary.Int(6),
			Distance:              unifieddatalibrary.Float(63.1),
			DistUnit:              unifieddatalibrary.String("Nautical miles"),
			FirstPt:               unifieddatalibrary.Time(time.Now()),
			IdealDesc:             unifieddatalibrary.String("Block speed using great circle path"),
			IdealDuration:         unifieddatalibrary.Float(45.1),
			IDSiteEnd:             unifieddatalibrary.String("77b5550c-c0f4-47bd-94ce-d71cdaa52f62"),
			IDSiteStart:           unifieddatalibrary.String("23370387-5e8e-4a74-89db-ad81145aa4df"),
			LastPt:                unifieddatalibrary.Time(time.Now()),
			LocationType:          unifieddatalibrary.String("ICAO"),
			MaxDuration:           unifieddatalibrary.Float(52.1),
			MaxSpeed:              unifieddatalibrary.Float(470.1),
			MinDuration:           unifieddatalibrary.Float(42.1),
			MinSpeed:              unifieddatalibrary.Float(420.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PartialDesc:           unifieddatalibrary.String("Performance speed using great circle path"),
			PartialDuration:       unifieddatalibrary.Float(38.1),
			SpeedUnit:             unifieddatalibrary.String("knots"),
			TimePeriod:            unifieddatalibrary.String("Q1"),
			VehicleCategory:       unifieddatalibrary.String("AIRCRAFT"),
			VehicleType:           unifieddatalibrary.String("C-17"),
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
