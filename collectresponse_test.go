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

func TestCollectResponseNewWithOptionalParams(t *testing.T) {
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
	err := client.CollectResponses.New(context.TODO(), unifieddatalibrary.CollectResponseNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.CollectResponseNewParamsDataModeTest,
		IDRequest:             "REF-REQUEST-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("COLLECTRESPONSE-ID"),
		ActualEndTime:         unifieddatalibrary.Time(time.Now()),
		ActualStartTime:       unifieddatalibrary.Time(time.Now()),
		AltEndTime:            unifieddatalibrary.Time(time.Now()),
		AltStartTime:          unifieddatalibrary.Time(time.Now()),
		ErrCode:               unifieddatalibrary.String("ERROR CODE"),
		ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
		IDPlan:                unifieddatalibrary.String("REF-PLAN-ID"),
		IDSensor:              unifieddatalibrary.String("REF-SENSOR-ID"),
		Notes:                 unifieddatalibrary.String("Example notes"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		SatNo:                 unifieddatalibrary.Int(101),
		SrcIDs:                []string{"DOA_ID", "DWELL_ID"},
		SrcTyps:               []string{"DOA", "DWELL"},
		Status:                unifieddatalibrary.String("ACCEPTED"),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TaskID:                unifieddatalibrary.String("TASK-ID"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectResponseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.CollectResponses.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.CollectResponseGetParams{
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

func TestCollectResponseListWithOptionalParams(t *testing.T) {
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
	_, err := client.CollectResponses.List(context.TODO(), unifieddatalibrary.CollectResponseListParams{
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

func TestCollectResponseCountWithOptionalParams(t *testing.T) {
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
	_, err := client.CollectResponses.Count(context.TODO(), unifieddatalibrary.CollectResponseCountParams{
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

func TestCollectResponseNewBulk(t *testing.T) {
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
	err := client.CollectResponses.NewBulk(context.TODO(), unifieddatalibrary.CollectResponseNewBulkParams{
		Body: []unifieddatalibrary.CollectResponseNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDRequest:             "REF-REQUEST-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("COLLECTRESPONSE-ID"),
			ActualEndTime:         unifieddatalibrary.Time(time.Now()),
			ActualStartTime:       unifieddatalibrary.Time(time.Now()),
			AltEndTime:            unifieddatalibrary.Time(time.Now()),
			AltStartTime:          unifieddatalibrary.Time(time.Now()),
			ErrCode:               unifieddatalibrary.String("ERROR CODE"),
			ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
			IDPlan:                unifieddatalibrary.String("REF-PLAN-ID"),
			IDSensor:              unifieddatalibrary.String("REF-SENSOR-ID"),
			Notes:                 unifieddatalibrary.String("Example notes"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			SatNo:                 unifieddatalibrary.Int(101),
			SrcIDs:                []string{"DOA_ID", "DWELL_ID"},
			SrcTyps:               []string{"DOA", "DWELL"},
			Status:                unifieddatalibrary.String("ACCEPTED"),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
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

func TestCollectResponseQueryHelp(t *testing.T) {
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
	_, err := client.CollectResponses.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectResponseUnvalidatedPublish(t *testing.T) {
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
	err := client.CollectResponses.UnvalidatedPublish(context.TODO(), unifieddatalibrary.CollectResponseUnvalidatedPublishParams{
		Body: []unifieddatalibrary.CollectResponseUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDRequest:             "REF-REQUEST-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("COLLECTRESPONSE-ID"),
			ActualEndTime:         unifieddatalibrary.Time(time.Now()),
			ActualStartTime:       unifieddatalibrary.Time(time.Now()),
			AltEndTime:            unifieddatalibrary.Time(time.Now()),
			AltStartTime:          unifieddatalibrary.Time(time.Now()),
			ErrCode:               unifieddatalibrary.String("ERROR CODE"),
			ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
			IDPlan:                unifieddatalibrary.String("REF-PLAN-ID"),
			IDSensor:              unifieddatalibrary.String("REF-SENSOR-ID"),
			Notes:                 unifieddatalibrary.String("Example notes"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			SatNo:                 unifieddatalibrary.Int(101),
			SrcIDs:                []string{"DOA_ID", "DWELL_ID"},
			SrcTyps:               []string{"DOA", "DWELL"},
			Status:                unifieddatalibrary.String("ACCEPTED"),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
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
