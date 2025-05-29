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

func TestManifoldelsetNewWithOptionalParams(t *testing.T) {
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
	err := client.Manifoldelset.New(context.TODO(), unifieddatalibrary.ManifoldelsetNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ManifoldelsetNewParamsDataModeTest,
		Epoch:                 time.Now(),
		IDManifold:            "REF-MANIFOLD-ID",
		Source:                "Bluestaq",
		TmpSatNo:              10,
		ID:                    unifieddatalibrary.String("MANIFOLDELSET-ID"),
		Apogee:                unifieddatalibrary.Float(10.23),
		ArgOfPerigee:          unifieddatalibrary.Float(10.23),
		BStar:                 unifieddatalibrary.Float(10.23),
		Eccentricity:          unifieddatalibrary.Float(0.5),
		Inclination:           unifieddatalibrary.Float(90.23),
		MeanAnomaly:           unifieddatalibrary.Float(10.23),
		MeanMotion:            unifieddatalibrary.Float(10.23),
		MeanMotionDDot:        unifieddatalibrary.Float(10.23),
		MeanMotionDot:         unifieddatalibrary.Float(10.23),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Perigee:               unifieddatalibrary.Float(10.23),
		Period:                unifieddatalibrary.Float(10.23),
		Raan:                  unifieddatalibrary.Float(10.23),
		RevNo:                 unifieddatalibrary.Int(123),
		SemiMajorAxis:         unifieddatalibrary.Float(10.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManifoldelsetUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Manifoldelset.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.ManifoldelsetUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ManifoldelsetUpdateParamsDataModeTest,
			Epoch:                 time.Now(),
			IDManifold:            "REF-MANIFOLD-ID",
			Source:                "Bluestaq",
			TmpSatNo:              10,
			ID:                    unifieddatalibrary.String("MANIFOLDELSET-ID"),
			Apogee:                unifieddatalibrary.Float(10.23),
			ArgOfPerigee:          unifieddatalibrary.Float(10.23),
			BStar:                 unifieddatalibrary.Float(10.23),
			Eccentricity:          unifieddatalibrary.Float(0.5),
			Inclination:           unifieddatalibrary.Float(90.23),
			MeanAnomaly:           unifieddatalibrary.Float(10.23),
			MeanMotion:            unifieddatalibrary.Float(10.23),
			MeanMotionDDot:        unifieddatalibrary.Float(10.23),
			MeanMotionDot:         unifieddatalibrary.Float(10.23),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Perigee:               unifieddatalibrary.Float(10.23),
			Period:                unifieddatalibrary.Float(10.23),
			Raan:                  unifieddatalibrary.Float(10.23),
			RevNo:                 unifieddatalibrary.Int(123),
			SemiMajorAxis:         unifieddatalibrary.Float(10.23),
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

func TestManifoldelsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Manifoldelset.List(context.TODO(), unifieddatalibrary.ManifoldelsetListParams{
		Epoch:       time.Now(),
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

func TestManifoldelsetDelete(t *testing.T) {
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
	err := client.Manifoldelset.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManifoldelsetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Manifoldelset.Count(context.TODO(), unifieddatalibrary.ManifoldelsetCountParams{
		Epoch:       time.Now(),
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

func TestManifoldelsetNewBulk(t *testing.T) {
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
	err := client.Manifoldelset.NewBulk(context.TODO(), unifieddatalibrary.ManifoldelsetNewBulkParams{
		Body: []unifieddatalibrary.ManifoldelsetNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Epoch:                 time.Now(),
			IDManifold:            "REF-MANIFOLD-ID",
			Source:                "Bluestaq",
			TmpSatNo:              10,
			ID:                    unifieddatalibrary.String("MANIFOLDELSET-ID"),
			Apogee:                unifieddatalibrary.Float(10.23),
			ArgOfPerigee:          unifieddatalibrary.Float(10.23),
			BStar:                 unifieddatalibrary.Float(10.23),
			Eccentricity:          unifieddatalibrary.Float(0.5),
			Inclination:           unifieddatalibrary.Float(90.23),
			MeanAnomaly:           unifieddatalibrary.Float(10.23),
			MeanMotion:            unifieddatalibrary.Float(10.23),
			MeanMotionDDot:        unifieddatalibrary.Float(10.23),
			MeanMotionDot:         unifieddatalibrary.Float(10.23),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Perigee:               unifieddatalibrary.Float(10.23),
			Period:                unifieddatalibrary.Float(10.23),
			Raan:                  unifieddatalibrary.Float(10.23),
			RevNo:                 unifieddatalibrary.Int(123),
			SemiMajorAxis:         unifieddatalibrary.Float(10.23),
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

func TestManifoldelsetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Manifoldelset.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ManifoldelsetGetParams{
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

func TestManifoldelsetQueryhelp(t *testing.T) {
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
	_, err := client.Manifoldelset.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManifoldelsetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Manifoldelset.Tuple(context.TODO(), unifieddatalibrary.ManifoldelsetTupleParams{
		Columns:     "columns",
		Epoch:       time.Now(),
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
