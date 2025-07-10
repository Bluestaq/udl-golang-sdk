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

func TestGeoStatusNewWithOptionalParams(t *testing.T) {
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
	err := client.GeoStatus.New(context.TODO(), unifieddatalibrary.GeoStatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.GeoStatusNewParamsDataModeTest,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("GEOSTATUS-ID"),
		ConfidenceLevel:       unifieddatalibrary.String("Low"),
		LongitudeMax:          unifieddatalibrary.Float(1.01),
		LongitudeMin:          unifieddatalibrary.Float(180.1),
		LongitudeRate:         unifieddatalibrary.Float(1.1),
		LostFlag:              unifieddatalibrary.Bool(false),
		ObjectStatus:          unifieddatalibrary.String("Active"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		PlaneChangeStatus:     unifieddatalibrary.String("Current"),
		RawFileUri:            unifieddatalibrary.String("Example URI"),
		RelativeEnergy:        unifieddatalibrary.Float(1.1),
		SatNo:                 unifieddatalibrary.Int(21),
		Sc:                    unifieddatalibrary.Float(1.1),
		SemiAnnualCorrFlag:    unifieddatalibrary.Bool(false),
		SS:                    unifieddatalibrary.Float(1.1),
		TroughType:            unifieddatalibrary.String("West"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeoStatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.GeoStatus.List(context.TODO(), unifieddatalibrary.GeoStatusListParams{
		CreatedAt:   time.Now(),
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

func TestGeoStatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.GeoStatus.Count(context.TODO(), unifieddatalibrary.GeoStatusCountParams{
		CreatedAt:   time.Now(),
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

func TestGeoStatusNewBulk(t *testing.T) {
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
	err := client.GeoStatus.NewBulk(context.TODO(), unifieddatalibrary.GeoStatusNewBulkParams{
		Body: []unifieddatalibrary.GeoStatusNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("GEOSTATUS-ID"),
			ConfidenceLevel:       unifieddatalibrary.String("Low"),
			LongitudeMax:          unifieddatalibrary.Float(1.01),
			LongitudeMin:          unifieddatalibrary.Float(180.1),
			LongitudeRate:         unifieddatalibrary.Float(1.1),
			LostFlag:              unifieddatalibrary.Bool(false),
			ObjectStatus:          unifieddatalibrary.String("Active"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			PlaneChangeStatus:     unifieddatalibrary.String("Current"),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			RelativeEnergy:        unifieddatalibrary.Float(1.1),
			SatNo:                 unifieddatalibrary.Int(21),
			Sc:                    unifieddatalibrary.Float(1.1),
			SemiAnnualCorrFlag:    unifieddatalibrary.Bool(false),
			SS:                    unifieddatalibrary.Float(1.1),
			TroughType:            unifieddatalibrary.String("West"),
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

func TestGeoStatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.GeoStatus.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.GeoStatusGetParams{
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

func TestGeoStatusQueryhelp(t *testing.T) {
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
	_, err := client.GeoStatus.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeoStatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.GeoStatus.Tuple(context.TODO(), unifieddatalibrary.GeoStatusTupleParams{
		Columns:     "columns",
		CreatedAt:   time.Now(),
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
