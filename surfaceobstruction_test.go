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

func TestSurfaceobstructionNewWithOptionalParams(t *testing.T) {
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
	err := client.Surfaceobstruction.New(context.TODO(), unifieddatalibrary.SurfaceobstructionNewParams{
		ClassificationMarking:     "U",
		DataMode:                  unifieddatalibrary.SurfaceobstructionNewParamsDataModeTest,
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

func TestSurfaceobstructionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Surfaceobstruction.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceobstructionUpdateParams{
			ClassificationMarking:     "U",
			DataMode:                  unifieddatalibrary.SurfaceobstructionUpdateParamsDataModeTest,
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

func TestSurfaceobstructionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Surfaceobstruction.List(context.TODO(), unifieddatalibrary.SurfaceobstructionListParams{
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

func TestSurfaceobstructionDelete(t *testing.T) {
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
	err := client.Surfaceobstruction.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceobstructionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Surfaceobstruction.Count(context.TODO(), unifieddatalibrary.SurfaceobstructionCountParams{
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

func TestSurfaceobstructionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Surfaceobstruction.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceobstructionGetParams{
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

func TestSurfaceobstructionQueryhelp(t *testing.T) {
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
	err := client.Surfaceobstruction.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceobstructionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Surfaceobstruction.Tuple(context.TODO(), unifieddatalibrary.SurfaceobstructionTupleParams{
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

func TestSurfaceobstructionUnvalidatedPublish(t *testing.T) {
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
	err := client.Surfaceobstruction.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SurfaceobstructionUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SurfaceobstructionUnvalidatedPublishParamsBody{{
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
