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

func TestEffectRequestNewWithOptionalParams(t *testing.T) {
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
	err := client.EffectRequests.New(context.TODO(), unifieddatalibrary.EffectRequestNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EffectRequestNewParamsDataModeTest,
		EffectList:            []string{"COVER", "DECEIVE"},
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("EFFECTREQUEST-ID"),
		Context:               unifieddatalibrary.String("Example Notes"),
		DeadlineType:          unifieddatalibrary.String("NoLaterThan"),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		ExternalRequestID:     unifieddatalibrary.String("EXTERNALREQUEST-ID"),
		MetricTypes:           []string{"COST", "RISK"},
		MetricWeights:         []float64{0.5, 0.6},
		ModelClass:            unifieddatalibrary.String("Preference model"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Priority:              unifieddatalibrary.String("LOW"),
		StartTime:             unifieddatalibrary.Time(time.Now()),
		State:                 unifieddatalibrary.String("CREATED"),
		TargetSrcID:           unifieddatalibrary.String("TARGETSRC-ID"),
		TargetSrcType:         unifieddatalibrary.String("POI"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEffectRequestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectRequests.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EffectRequestGetParams{
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

func TestEffectRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectRequests.List(context.TODO(), unifieddatalibrary.EffectRequestListParams{
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

func TestEffectRequestCountWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectRequests.Count(context.TODO(), unifieddatalibrary.EffectRequestCountParams{
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

func TestEffectRequestNewBulk(t *testing.T) {
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
	err := client.EffectRequests.NewBulk(context.TODO(), unifieddatalibrary.EffectRequestNewBulkParams{
		Body: []unifieddatalibrary.EffectRequestNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EffectList:            []string{"COVER", "DECEIVE"},
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("EFFECTREQUEST-ID"),
			Context:               unifieddatalibrary.String("Example Notes"),
			DeadlineType:          unifieddatalibrary.String("NoLaterThan"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ExternalRequestID:     unifieddatalibrary.String("EXTERNALREQUEST-ID"),
			MetricTypes:           []string{"COST", "RISK"},
			MetricWeights:         []float64{0.5, 0.6},
			ModelClass:            unifieddatalibrary.String("Preference model"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Priority:              unifieddatalibrary.String("LOW"),
			StartTime:             unifieddatalibrary.Time(time.Now()),
			State:                 unifieddatalibrary.String("CREATED"),
			TargetSrcID:           unifieddatalibrary.String("TARGETSRC-ID"),
			TargetSrcType:         unifieddatalibrary.String("POI"),
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

func TestEffectRequestQueryHelp(t *testing.T) {
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
	_, err := client.EffectRequests.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEffectRequestTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectRequests.Tuple(context.TODO(), unifieddatalibrary.EffectRequestTupleParams{
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

func TestEffectRequestUnvalidatedPublish(t *testing.T) {
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
	err := client.EffectRequests.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EffectRequestUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EffectRequestUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EffectList:            []string{"COVER", "DECEIVE"},
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("EFFECTREQUEST-ID"),
			Context:               unifieddatalibrary.String("Example Notes"),
			DeadlineType:          unifieddatalibrary.String("NoLaterThan"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ExternalRequestID:     unifieddatalibrary.String("EXTERNALREQUEST-ID"),
			MetricTypes:           []string{"COST", "RISK"},
			MetricWeights:         []float64{0.5, 0.6},
			ModelClass:            unifieddatalibrary.String("Preference model"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Priority:              unifieddatalibrary.String("LOW"),
			StartTime:             unifieddatalibrary.Time(time.Now()),
			State:                 unifieddatalibrary.String("CREATED"),
			TargetSrcID:           unifieddatalibrary.String("TARGETSRC-ID"),
			TargetSrcType:         unifieddatalibrary.String("POI"),
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
