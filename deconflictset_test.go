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

func TestDeconflictsetNewWithOptionalParams(t *testing.T) {
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
	err := client.Deconflictset.New(context.TODO(), unifieddatalibrary.DeconflictsetNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.DeconflictsetNewParamsDataModeTest,
		EventStartTime:        time.Now(),
		NumWindows:            250001,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("123dd511-8ba5-47d3-9909-836149f87434"),
		CalculationEndTime:    unifieddatalibrary.Time(time.Now()),
		CalculationID:         unifieddatalibrary.String("3856c0a0-585f-4232-af5d-93bad320fac6"),
		CalculationStartTime:  unifieddatalibrary.Time(time.Now()),
		DeconflictWindows: []unifieddatalibrary.DeconflictsetNewParamsDeconflictWindow{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EventStartTime:        time.Now(),
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			StopTime:              time.Now(),
			ID:                    unifieddatalibrary.String("123dd511-8ba5-47d3-9909-836149f87434"),
			AngleOfEntry:          unifieddatalibrary.Float(0.65),
			AngleOfExit:           unifieddatalibrary.Float(0.65),
			EntryCoords:           []float64{-191500.74728263554, -987729.0529358581, 6735105.853234725},
			EventType:             unifieddatalibrary.String("LASER"),
			ExitCoords:            []float64{-361767.9896431379, -854021.6371921108, 6746208.020741149},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Target:                unifieddatalibrary.String("41715"),
			TargetType:            unifieddatalibrary.String("VICTIM"),
			Victim:                unifieddatalibrary.String("55914"),
			WindowType:            unifieddatalibrary.String("CLOSED"),
		}},
		Errors:                   []string{"ERROR1", "ERROR2"},
		EventEndTime:             unifieddatalibrary.Time(time.Now()),
		EventType:                unifieddatalibrary.String("LASER"),
		IDLaserDeconflictRequest: unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ReferenceFrame:           unifieddatalibrary.String("J2000"),
		Tags:                     []string{"TAG1", "TAG2"},
		TransactionID:            unifieddatalibrary.String("TRANSACTION-ID"),
		Warnings:                 []string{"WARNING1", "WARNING2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeconflictsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Deconflictset.List(context.TODO(), unifieddatalibrary.DeconflictsetListParams{
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeconflictsetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Deconflictset.Count(context.TODO(), unifieddatalibrary.DeconflictsetCountParams{
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeconflictsetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Deconflictset.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.DeconflictsetGetParams{
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

func TestDeconflictsetQueryhelp(t *testing.T) {
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
	_, err := client.Deconflictset.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeconflictsetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Deconflictset.Tuple(context.TODO(), unifieddatalibrary.DeconflictsetTupleParams{
		Columns:        "columns",
		EventStartTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeconflictsetUnvalidatedPublishWithOptionalParams(t *testing.T) {
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
	err := client.Deconflictset.UnvalidatedPublish(context.TODO(), unifieddatalibrary.DeconflictsetUnvalidatedPublishParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.DeconflictsetUnvalidatedPublishParamsDataModeTest,
		EventStartTime:        time.Now(),
		NumWindows:            250001,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("123dd511-8ba5-47d3-9909-836149f87434"),
		CalculationEndTime:    unifieddatalibrary.Time(time.Now()),
		CalculationID:         unifieddatalibrary.String("3856c0a0-585f-4232-af5d-93bad320fac6"),
		CalculationStartTime:  unifieddatalibrary.Time(time.Now()),
		DeconflictWindows: []unifieddatalibrary.DeconflictsetUnvalidatedPublishParamsDeconflictWindow{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EventStartTime:        time.Now(),
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			StopTime:              time.Now(),
			ID:                    unifieddatalibrary.String("123dd511-8ba5-47d3-9909-836149f87434"),
			AngleOfEntry:          unifieddatalibrary.Float(0.65),
			AngleOfExit:           unifieddatalibrary.Float(0.65),
			EntryCoords:           []float64{-191500.74728263554, -987729.0529358581, 6735105.853234725},
			EventType:             unifieddatalibrary.String("LASER"),
			ExitCoords:            []float64{-361767.9896431379, -854021.6371921108, 6746208.020741149},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Target:                unifieddatalibrary.String("41715"),
			TargetType:            unifieddatalibrary.String("VICTIM"),
			Victim:                unifieddatalibrary.String("55914"),
			WindowType:            unifieddatalibrary.String("CLOSED"),
		}},
		Errors:                   []string{"ERROR1", "ERROR2"},
		EventEndTime:             unifieddatalibrary.Time(time.Now()),
		EventType:                unifieddatalibrary.String("LASER"),
		IDLaserDeconflictRequest: unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ReferenceFrame:           unifieddatalibrary.String("J2000"),
		Tags:                     []string{"TAG1", "TAG2"},
		TransactionID:            unifieddatalibrary.String("TRANSACTION-ID"),
		Warnings:                 []string{"WARNING1", "WARNING2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
