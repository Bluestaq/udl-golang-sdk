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

func TestOnboardnavigationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onboardnavigation.List(context.TODO(), unifieddatalibrary.OnboardnavigationListParams{
		StartTime:   time.Now(),
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

func TestOnboardnavigationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Onboardnavigation.Count(context.TODO(), unifieddatalibrary.OnboardnavigationCountParams{
		StartTime:   time.Now(),
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

func TestOnboardnavigationNewBulk(t *testing.T) {
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
	err := client.Onboardnavigation.NewBulk(context.TODO(), unifieddatalibrary.OnboardnavigationNewBulkParams{
		Body: []unifieddatalibrary.OnboardnavigationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("ONBOARD-NAVIGATION-ID"),
			DeltaPos:              [][]float64{{1.1, 2.2, 3.3}},
			EndTime:               unifieddatalibrary.Time(time.Now()),
			EsID:                  unifieddatalibrary.String("EPHEMERISSET-ID"),
			IDStateVector:         unifieddatalibrary.String("STATE-VECTOR-ID"),
			Mag:                   [][]float64{{1.1, 2.2, 3.3}},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			ReferenceFrame:        "J2000",
			SatNo:                 unifieddatalibrary.Int(101),
			StarCatLoadTime:       unifieddatalibrary.Time(time.Now()),
			StarCatName:           unifieddatalibrary.String("STAR-CAT-NAME"),
			StarTracker:           [][]float64{{1.1, 2.2, 3.3}},
			SunSensor:             [][]float64{{1.1, 2.2, 3.3}},
			Ts:                    []time.Time{time.Now()},
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

func TestOnboardnavigationQueryhelp(t *testing.T) {
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
	_, err := client.Onboardnavigation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnboardnavigationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Onboardnavigation.Tuple(context.TODO(), unifieddatalibrary.OnboardnavigationTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestOnboardnavigationUnvalidatedPublish(t *testing.T) {
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
	err := client.Onboardnavigation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.OnboardnavigationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.OnboardnavigationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("ONBOARD-NAVIGATION-ID"),
			DeltaPos:              [][]float64{{1.1, 2.2, 3.3}},
			EndTime:               unifieddatalibrary.Time(time.Now()),
			EsID:                  unifieddatalibrary.String("EPHEMERISSET-ID"),
			IDStateVector:         unifieddatalibrary.String("STATE-VECTOR-ID"),
			Mag:                   [][]float64{{1.1, 2.2, 3.3}},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			ReferenceFrame:        "J2000",
			SatNo:                 unifieddatalibrary.Int(101),
			StarCatLoadTime:       unifieddatalibrary.Time(time.Now()),
			StarCatName:           unifieddatalibrary.String("STAR-CAT-NAME"),
			StarTracker:           [][]float64{{1.1, 2.2, 3.3}},
			SunSensor:             [][]float64{{1.1, 2.2, 3.3}},
			Ts:                    []time.Time{time.Now()},
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
