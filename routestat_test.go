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

func TestRoutestatNewWithOptionalParams(t *testing.T) {
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
	err := client.Routestats.New(context.TODO(), unifieddatalibrary.RoutestatNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RoutestatNewParamsDataModeTest,
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

func TestRoutestatGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Routestats.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RoutestatGetParams{
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

func TestRoutestatUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Routestats.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RoutestatUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.RoutestatUpdateParamsDataModeTest,
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

func TestRoutestatDelete(t *testing.T) {
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
	err := client.Routestats.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutestatCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Routestats.Count(context.TODO(), unifieddatalibrary.RoutestatCountParams{
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

func TestRoutestatNewBulk(t *testing.T) {
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
	err := client.Routestats.NewBulk(context.TODO(), unifieddatalibrary.RoutestatNewBulkParams{
		Body: []unifieddatalibrary.RoutestatNewBulkParamsBody{{
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

func TestRoutestatQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Routestats.Query(context.TODO(), unifieddatalibrary.RoutestatQueryParams{
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

func TestRoutestatQueryHelp(t *testing.T) {
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
	err := client.Routestats.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutestatTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Routestats.Tuple(context.TODO(), unifieddatalibrary.RoutestatTupleParams{
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

func TestRoutestatUnvalidatedPublish(t *testing.T) {
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
	err := client.Routestats.UnvalidatedPublish(context.TODO(), unifieddatalibrary.RoutestatUnvalidatedPublishParams{
		Body: []unifieddatalibrary.RoutestatUnvalidatedPublishParamsBody{{
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
