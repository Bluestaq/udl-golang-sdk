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

func TestFeatureAssessmentNewWithOptionalParams(t *testing.T) {
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
	err := client.FeatureAssessment.New(context.TODO(), unifieddatalibrary.FeatureAssessmentNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.FeatureAssessmentNewParamsDataModeTest,
		FeatureTs:             time.Now(),
		FeatureUoM:            "MHz",
		IDAnalyticImagery:     "fa1509ae-c19d-432e-9542-e5d1e0f47bc3",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Agjson:                unifieddatalibrary.String(`{"type":"Point","coordinates":[52.23486096929749 16.191937138595005]}`),
		Andims:                unifieddatalibrary.Int(1),
		AnnLims:               [][]int64{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
		AnnText:               []string{"rec1", "rec2"},
		Area:                  unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
		Asrid:                 unifieddatalibrary.Int(4326),
		Assessment:            unifieddatalibrary.String("Vessel bigger than other small fishing boats commonly found along the coastline"),
		Atext:                 unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
		Atype:                 unifieddatalibrary.String("POINT"),
		Confidence:            unifieddatalibrary.Float(0.85),
		ExternalID:            unifieddatalibrary.String("2024-06-22-17-53-05_UMBRA-05_GEC"),
		FeatureArray:          []float64{1227.6, 1575.42},
		FeatureBool:           unifieddatalibrary.Bool(true),
		FeatureString:         unifieddatalibrary.String("TRANSMITTING FREQUENCIES"),
		FeatureStringArray:    []string{"String1", "String2"},
		FeatureValue:          unifieddatalibrary.Float(1227.6),
		Heading:               unifieddatalibrary.Float(97.1),
		Height:                unifieddatalibrary.Float(7.25),
		Length:                unifieddatalibrary.Float(10.54),
		Name:                  unifieddatalibrary.String("HEADING"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Speed:                 unifieddatalibrary.Float(0.1),
		SrcIDs:                []string{"b008c63b-ad89-4493-80e0-77bc982bef77", "3565a6dd-654e-4969-89e0-ee7c51ab1e1b"},
		SrcTs:                 []time.Time{time.Now(), time.Now()},
		SrcTyps:               []string{"SAR", "AIS"},
		Tags:                  []string{"TAG1", "TAG2"},
		TransactionID:         unifieddatalibrary.String("c3bdef1f-5a4f-4716-bee4-7a1e0ec7d37d"),
		Type:                  unifieddatalibrary.String("VESSEL"),
		Width:                 unifieddatalibrary.Float(3.74),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFeatureAssessmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.FeatureAssessment.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.FeatureAssessmentGetParams{
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

func TestFeatureAssessmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.FeatureAssessment.List(context.TODO(), unifieddatalibrary.FeatureAssessmentListParams{
		IDAnalyticImagery: "idAnalyticImagery",
		FirstResult:       unifieddatalibrary.Int(0),
		MaxResults:        unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFeatureAssessmentCountWithOptionalParams(t *testing.T) {
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
	_, err := client.FeatureAssessment.Count(context.TODO(), unifieddatalibrary.FeatureAssessmentCountParams{
		IDAnalyticImagery: "idAnalyticImagery",
		FirstResult:       unifieddatalibrary.Int(0),
		MaxResults:        unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFeatureAssessmentNewBulk(t *testing.T) {
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
	err := client.FeatureAssessment.NewBulk(context.TODO(), unifieddatalibrary.FeatureAssessmentNewBulkParams{
		Body: []unifieddatalibrary.FeatureAssessmentNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			FeatureTs:             time.Now(),
			FeatureUoM:            "MHz",
			IDAnalyticImagery:     "fa1509ae-c19d-432e-9542-e5d1e0f47bc3",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Agjson:                unifieddatalibrary.String(`{"type":"Point","coordinates":[52.23486096929749 16.191937138595005]}`),
			Andims:                unifieddatalibrary.Int(1),
			AnnLims:               [][]int64{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
			AnnText:               []string{"rec1", "rec2"},
			Area:                  unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
			Asrid:                 unifieddatalibrary.Int(4326),
			Assessment:            unifieddatalibrary.String("Vessel bigger than other small fishing boats commonly found along the coastline"),
			Atext:                 unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
			Atype:                 unifieddatalibrary.String("POINT"),
			Confidence:            unifieddatalibrary.Float(0.85),
			ExternalID:            unifieddatalibrary.String("2024-06-22-17-53-05_UMBRA-05_GEC"),
			FeatureArray:          []float64{1227.6, 1575.42},
			FeatureBool:           unifieddatalibrary.Bool(true),
			FeatureString:         unifieddatalibrary.String("TRANSMITTING FREQUENCIES"),
			FeatureStringArray:    []string{"String1", "String2"},
			FeatureValue:          unifieddatalibrary.Float(1227.6),
			Heading:               unifieddatalibrary.Float(97.1),
			Height:                unifieddatalibrary.Float(7.25),
			Length:                unifieddatalibrary.Float(10.54),
			Name:                  unifieddatalibrary.String("HEADING"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Speed:                 unifieddatalibrary.Float(0.1),
			SrcIDs:                []string{"b008c63b-ad89-4493-80e0-77bc982bef77", "3565a6dd-654e-4969-89e0-ee7c51ab1e1b"},
			SrcTs:                 []time.Time{time.Now(), time.Now()},
			SrcTyps:               []string{"SAR", "AIS"},
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("c3bdef1f-5a4f-4716-bee4-7a1e0ec7d37d"),
			Type:                  unifieddatalibrary.String("VESSEL"),
			Width:                 unifieddatalibrary.Float(3.74),
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

func TestFeatureAssessmentQueryHelp(t *testing.T) {
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
	_, err := client.FeatureAssessment.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFeatureAssessmentTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.FeatureAssessment.Tuple(context.TODO(), unifieddatalibrary.FeatureAssessmentTupleParams{
		Columns:           "columns",
		IDAnalyticImagery: "idAnalyticImagery",
		FirstResult:       unifieddatalibrary.Int(0),
		MaxResults:        unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFeatureAssessmentUnvalidatedPublish(t *testing.T) {
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
	err := client.FeatureAssessment.UnvalidatedPublish(context.TODO(), unifieddatalibrary.FeatureAssessmentUnvalidatedPublishParams{
		Body: []unifieddatalibrary.FeatureAssessmentUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			FeatureTs:             time.Now(),
			FeatureUoM:            "MHz",
			IDAnalyticImagery:     "fa1509ae-c19d-432e-9542-e5d1e0f47bc3",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Agjson:                unifieddatalibrary.String(`{"type":"Point","coordinates":[52.23486096929749 16.191937138595005]}`),
			Andims:                unifieddatalibrary.Int(1),
			AnnLims:               [][]int64{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
			AnnText:               []string{"rec1", "rec2"},
			Area:                  unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
			Asrid:                 unifieddatalibrary.Int(4326),
			Assessment:            unifieddatalibrary.String("Vessel bigger than other small fishing boats commonly found along the coastline"),
			Atext:                 unifieddatalibrary.String("POINT(52.23486096929749 16.191937138595005)"),
			Atype:                 unifieddatalibrary.String("POINT"),
			Confidence:            unifieddatalibrary.Float(0.85),
			ExternalID:            unifieddatalibrary.String("2024-06-22-17-53-05_UMBRA-05_GEC"),
			FeatureArray:          []float64{1227.6, 1575.42},
			FeatureBool:           unifieddatalibrary.Bool(true),
			FeatureString:         unifieddatalibrary.String("TRANSMITTING FREQUENCIES"),
			FeatureStringArray:    []string{"String1", "String2"},
			FeatureValue:          unifieddatalibrary.Float(1227.6),
			Heading:               unifieddatalibrary.Float(97.1),
			Height:                unifieddatalibrary.Float(7.25),
			Length:                unifieddatalibrary.Float(10.54),
			Name:                  unifieddatalibrary.String("HEADING"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Speed:                 unifieddatalibrary.Float(0.1),
			SrcIDs:                []string{"b008c63b-ad89-4493-80e0-77bc982bef77", "3565a6dd-654e-4969-89e0-ee7c51ab1e1b"},
			SrcTs:                 []time.Time{time.Now(), time.Now()},
			SrcTyps:               []string{"SAR", "AIS"},
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("c3bdef1f-5a4f-4716-bee4-7a1e0ec7d37d"),
			Type:                  unifieddatalibrary.String("VESSEL"),
			Width:                 unifieddatalibrary.Float(3.74),
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
