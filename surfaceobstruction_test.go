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

func TestSurfaceObstructionNewWithOptionalParams(t *testing.T) {
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
	err := client.SurfaceObstruction.New(context.TODO(), unifieddatalibrary.SurfaceObstructionNewParams{
		ClassificationMarking:     "U",
		DataMode:                  unifieddatalibrary.SurfaceObstructionNewParamsDataModeTest,
		IDSurface:                 "be831d39-1822-da9f-7ace-6cc5643397dc",
		Source:                    "Bluestaq",
		ID:                        unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		AdvisoryRequired:          []string{"C20", "C17"},
		ApprovalRequired:          []string{"C20", "C17"},
		DistanceFromCenterLine:    unifieddatalibrary.Float(17.8),
		DistanceFromEdge:          unifieddatalibrary.Float(15.8),
		DistanceFromThreshold:     unifieddatalibrary.Float(19.5),
		IDNavigationalObstruction: unifieddatalibrary.String("a2831d39-1822-da9f-7ace-6cc5643397da"),
		ObstructionDesc:           unifieddatalibrary.String("PYLON"),
		ObstructionHeight:         unifieddatalibrary.Float(35.25),
		ObstructionSideCode:       unifieddatalibrary.String("F"),
		Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceObstructionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SurfaceObstruction.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceObstructionUpdateParams{
			ClassificationMarking:     "U",
			DataMode:                  unifieddatalibrary.SurfaceObstructionUpdateParamsDataModeTest,
			IDSurface:                 "be831d39-1822-da9f-7ace-6cc5643397dc",
			Source:                    "Bluestaq",
			ID:                        unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AdvisoryRequired:          []string{"C20", "C17"},
			ApprovalRequired:          []string{"C20", "C17"},
			DistanceFromCenterLine:    unifieddatalibrary.Float(17.8),
			DistanceFromEdge:          unifieddatalibrary.Float(15.8),
			DistanceFromThreshold:     unifieddatalibrary.Float(19.5),
			IDNavigationalObstruction: unifieddatalibrary.String("a2831d39-1822-da9f-7ace-6cc5643397da"),
			ObstructionDesc:           unifieddatalibrary.String("PYLON"),
			ObstructionHeight:         unifieddatalibrary.Float(35.25),
			ObstructionSideCode:       unifieddatalibrary.String("F"),
			Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestSurfaceObstructionListWithOptionalParams(t *testing.T) {
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
	_, err := client.SurfaceObstruction.List(context.TODO(), unifieddatalibrary.SurfaceObstructionListParams{
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

func TestSurfaceObstructionDelete(t *testing.T) {
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
	err := client.SurfaceObstruction.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceObstructionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SurfaceObstruction.Count(context.TODO(), unifieddatalibrary.SurfaceObstructionCountParams{
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

func TestSurfaceObstructionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SurfaceObstruction.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceObstructionGetParams{
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

func TestSurfaceObstructionQueryhelp(t *testing.T) {
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
	_, err := client.SurfaceObstruction.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceObstructionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SurfaceObstruction.Tuple(context.TODO(), unifieddatalibrary.SurfaceObstructionTupleParams{
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

func TestSurfaceObstructionUnvalidatedPublish(t *testing.T) {
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
	err := client.SurfaceObstruction.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SurfaceObstructionUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SurfaceObstructionUnvalidatedPublishParamsBody{{
			ClassificationMarking:     "U",
			DataMode:                  "TEST",
			IDSurface:                 "be831d39-1822-da9f-7ace-6cc5643397dc",
			Source:                    "Bluestaq",
			ID:                        unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AdvisoryRequired:          []string{"C20", "C17"},
			ApprovalRequired:          []string{"C20", "C17"},
			DistanceFromCenterLine:    unifieddatalibrary.Float(17.8),
			DistanceFromEdge:          unifieddatalibrary.Float(15.8),
			DistanceFromThreshold:     unifieddatalibrary.Float(19.5),
			IDNavigationalObstruction: unifieddatalibrary.String("a2831d39-1822-da9f-7ace-6cc5643397da"),
			ObstructionDesc:           unifieddatalibrary.String("PYLON"),
			ObstructionHeight:         unifieddatalibrary.Float(35.25),
			ObstructionSideCode:       unifieddatalibrary.String("F"),
			Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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
