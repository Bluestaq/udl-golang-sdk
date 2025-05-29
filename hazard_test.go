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

func TestHazardNewWithOptionalParams(t *testing.T) {
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
	err := client.Hazard.New(context.TODO(), unifieddatalibrary.HazardNewParams{
		Alarms:                []string{"Alarm1", "Alarm2"},
		AlarmValues:           []float64{2.7, 2.9},
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.HazardNewParamsDataModeTest,
		DetectTime:            time.Now(),
		DetectType:            "Chemical",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("HAZARD-ID"),
		A:                     unifieddatalibrary.Int(238),
		Activity:              unifieddatalibrary.Float(120.1),
		BottleID:              unifieddatalibrary.String("6264"),
		CasRn:                 unifieddatalibrary.String("64-17-5"),
		Channel:               unifieddatalibrary.String("Skin"),
		CtrnTime:              unifieddatalibrary.Float(1.077),
		Density:               unifieddatalibrary.Float(18900.2),
		Dep:                   unifieddatalibrary.Float(1.084),
		DepCtrn:               unifieddatalibrary.Float(86.1),
		Dose:                  unifieddatalibrary.Float(1.12),
		DoseRate:              unifieddatalibrary.Float(1.0000001865),
		Duration:              unifieddatalibrary.Int(14400),
		GBar:                  unifieddatalibrary.Float(2.5),
		Harmful:               unifieddatalibrary.Bool(false),
		HBar:                  unifieddatalibrary.Float(3.1),
		IDPoi:                 unifieddatalibrary.String("POI-ID"),
		IDTrack:               unifieddatalibrary.String("TRACK-ID"),
		MassFrac:              unifieddatalibrary.Float(0.029),
		MatCat:                unifieddatalibrary.Int(3),
		MatClass:              unifieddatalibrary.String("Nerve Agent"),
		MatName:               unifieddatalibrary.String("VX"),
		MatType:               unifieddatalibrary.String("21"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Ppm:                   unifieddatalibrary.Int(27129),
		RadCtrn:               unifieddatalibrary.Float(1.31),
		Readings:              []string{"Rad1", "Rad2"},
		ReadingUnits:          []string{"Gray", "Gray"},
		ReadingValues:         []float64{107.2, 124.1},
		Z:                     unifieddatalibrary.Int(92),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHazardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Hazard.List(context.TODO(), unifieddatalibrary.HazardListParams{
		DetectTime:  time.Now(),
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

func TestHazardCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Hazard.Count(context.TODO(), unifieddatalibrary.HazardCountParams{
		DetectTime:  time.Now(),
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

func TestHazardNewBulk(t *testing.T) {
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
	err := client.Hazard.NewBulk(context.TODO(), unifieddatalibrary.HazardNewBulkParams{
		Body: []unifieddatalibrary.HazardNewBulkParamsBody{{
			Alarms:                []string{"Alarm1", "Alarm2"},
			AlarmValues:           []float64{2.7, 2.9},
			ClassificationMarking: "U",
			DataMode:              "TEST",
			DetectTime:            time.Now(),
			DetectType:            "Chemical",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("HAZARD-ID"),
			A:                     unifieddatalibrary.Int(238),
			Activity:              unifieddatalibrary.Float(120.1),
			BottleID:              unifieddatalibrary.String("6264"),
			CasRn:                 unifieddatalibrary.String("64-17-5"),
			Channel:               unifieddatalibrary.String("Skin"),
			CtrnTime:              unifieddatalibrary.Float(1.077),
			Density:               unifieddatalibrary.Float(18900.2),
			Dep:                   unifieddatalibrary.Float(1.084),
			DepCtrn:               unifieddatalibrary.Float(86.1),
			Dose:                  unifieddatalibrary.Float(1.12),
			DoseRate:              unifieddatalibrary.Float(1.0000001865),
			Duration:              unifieddatalibrary.Int(14400),
			GBar:                  unifieddatalibrary.Float(2.5),
			Harmful:               unifieddatalibrary.Bool(false),
			HBar:                  unifieddatalibrary.Float(3.1),
			IDPoi:                 unifieddatalibrary.String("POI-ID"),
			IDTrack:               unifieddatalibrary.String("TRACK-ID"),
			MassFrac:              unifieddatalibrary.Float(0.029),
			MatCat:                unifieddatalibrary.Int(3),
			MatClass:              unifieddatalibrary.String("Nerve Agent"),
			MatName:               unifieddatalibrary.String("VX"),
			MatType:               unifieddatalibrary.String("21"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Ppm:                   unifieddatalibrary.Int(27129),
			RadCtrn:               unifieddatalibrary.Float(1.31),
			Readings:              []string{"Rad1", "Rad2"},
			ReadingUnits:          []string{"Gray", "Gray"},
			ReadingValues:         []float64{107.2, 124.1},
			Z:                     unifieddatalibrary.Int(92),
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

func TestHazardGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Hazard.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.HazardGetParams{
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

func TestHazardQueryhelp(t *testing.T) {
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
	_, err := client.Hazard.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHazardTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Hazard.Tuple(context.TODO(), unifieddatalibrary.HazardTupleParams{
		Columns:     "columns",
		DetectTime:  time.Now(),
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
