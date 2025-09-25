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

func TestObservationObscorrelationNewWithOptionalParams(t *testing.T) {
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
	err := client.Observations.Obscorrelation.New(context.TODO(), unifieddatalibrary.ObservationObscorrelationNewParams{
		ClassificationMarking: "U",
		CorrType:              unifieddatalibrary.ObservationObscorrelationNewParamsCorrTypeObservation,
		DataMode:              unifieddatalibrary.ObservationObscorrelationNewParamsDataModeTest,
		MsgTs:                 time.Now(),
		ObID:                  "e69c6734-30a1-4c4f-8fe2-7187e7012e8c",
		ObType:                unifieddatalibrary.ObservationObscorrelationNewParamsObTypeEo,
		ReferenceOrbitID:      "21826de2-5639-4c2a-b68f-30b67753b983",
		ReferenceOrbitType:    unifieddatalibrary.ObservationObscorrelationNewParamsReferenceOrbitTypeElset,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AlgorithmCorrType:     unifieddatalibrary.String("ROTAS"),
		AltCatalog:            unifieddatalibrary.String("CATALOG"),
		AltNamespace:          unifieddatalibrary.String("18SDS"),
		AltObjectID:           unifieddatalibrary.String("26900"),
		AltUct:                unifieddatalibrary.Bool(false),
		Astat:                 unifieddatalibrary.Int(2),
		CorrQuality:           unifieddatalibrary.Float(0.96),
		IDParentCorrelation:   unifieddatalibrary.String("ID-PARENT-CORRELATION"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		SatNo:                 unifieddatalibrary.Int(12100),
		Tags:                  []string{"TAG1", "TAG2"},
		TrackID:               unifieddatalibrary.String("TRACK-ID"),
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationObscorrelationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Obscorrelation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObservationObscorrelationGetParams{
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

func TestObservationObscorrelationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Obscorrelation.List(context.TODO(), unifieddatalibrary.ObservationObscorrelationListParams{
		MsgTs:       time.Now(),
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

func TestObservationObscorrelationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Obscorrelation.Count(context.TODO(), unifieddatalibrary.ObservationObscorrelationCountParams{
		MsgTs:       time.Now(),
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

func TestObservationObscorrelationNewBulk(t *testing.T) {
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
	err := client.Observations.Obscorrelation.NewBulk(context.TODO(), unifieddatalibrary.ObservationObscorrelationNewBulkParams{
		Body: []unifieddatalibrary.ObservationObscorrelationNewBulkParamsBody{{
			ClassificationMarking: "U",
			CorrType:              "OBSERVATION",
			DataMode:              "TEST",
			MsgTs:                 time.Now(),
			ObID:                  "e69c6734-30a1-4c4f-8fe2-7187e7012e8c",
			ObType:                "EO",
			ReferenceOrbitID:      "21826de2-5639-4c2a-b68f-30b67753b983",
			ReferenceOrbitType:    "ELSET",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AlgorithmCorrType:     unifieddatalibrary.String("ROTAS"),
			AltCatalog:            unifieddatalibrary.String("CATALOG"),
			AltNamespace:          unifieddatalibrary.String("18SDS"),
			AltObjectID:           unifieddatalibrary.String("26900"),
			AltUct:                unifieddatalibrary.Bool(false),
			Astat:                 unifieddatalibrary.Int(2),
			CorrQuality:           unifieddatalibrary.Float(0.96),
			IDParentCorrelation:   unifieddatalibrary.String("ID-PARENT-CORRELATION"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			SatNo:                 unifieddatalibrary.Int(12100),
			Tags:                  []string{"TAG1", "TAG2"},
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
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

func TestObservationObscorrelationQueryHelp(t *testing.T) {
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
	_, err := client.Observations.Obscorrelation.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationObscorrelationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Obscorrelation.Tuple(context.TODO(), unifieddatalibrary.ObservationObscorrelationTupleParams{
		Columns:     "columns",
		MsgTs:       time.Now(),
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

func TestObservationObscorrelationUnvalidatedPublish(t *testing.T) {
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
	err := client.Observations.Obscorrelation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ObservationObscorrelationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ObservationObscorrelationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			CorrType:              "OBSERVATION",
			DataMode:              "TEST",
			MsgTs:                 time.Now(),
			ObID:                  "e69c6734-30a1-4c4f-8fe2-7187e7012e8c",
			ObType:                "EO",
			ReferenceOrbitID:      "21826de2-5639-4c2a-b68f-30b67753b983",
			ReferenceOrbitType:    "ELSET",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AlgorithmCorrType:     unifieddatalibrary.String("ROTAS"),
			AltCatalog:            unifieddatalibrary.String("CATALOG"),
			AltNamespace:          unifieddatalibrary.String("18SDS"),
			AltObjectID:           unifieddatalibrary.String("26900"),
			AltUct:                unifieddatalibrary.Bool(false),
			Astat:                 unifieddatalibrary.Int(2),
			CorrQuality:           unifieddatalibrary.Float(0.96),
			IDParentCorrelation:   unifieddatalibrary.String("ID-PARENT-CORRELATION"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			SatNo:                 unifieddatalibrary.Int(12100),
			Tags:                  []string{"TAG1", "TAG2"},
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
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
