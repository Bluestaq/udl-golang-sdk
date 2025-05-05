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

func TestLaunchEventNewWithOptionalParams(t *testing.T) {
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
	err := client.LaunchEvent.New(context.TODO(), unifieddatalibrary.LaunchEventNewParams{
		ClassificationMarking:  "U",
		DataMode:               unifieddatalibrary.LaunchEventNewParamsDataModeTest,
		MsgCreateDate:          time.Now(),
		Source:                 "Bluestaq",
		ID:                     unifieddatalibrary.String("LAUNCHEVENT-ID"),
		BeNumber:               unifieddatalibrary.String("ENC-123"),
		DeclassificationDate:   unifieddatalibrary.Time(time.Now()),
		DeclassificationString: unifieddatalibrary.String("Example Declassification"),
		DerivedFrom:            unifieddatalibrary.String("Example source"),
		LaunchDate:             unifieddatalibrary.Time(time.Now()),
		LaunchFacilityName:     unifieddatalibrary.String("Example launch facility name"),
		LaunchFailureCode:      unifieddatalibrary.String("Example failure code"),
		Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
		OSuffix:                unifieddatalibrary.String("oSuffix"),
		SatNo:                  unifieddatalibrary.Int(12),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchEvent.List(context.TODO(), unifieddatalibrary.LaunchEventListParams{
		MsgCreateDate: time.Now(),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchEventCountWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchEvent.Count(context.TODO(), unifieddatalibrary.LaunchEventCountParams{
		MsgCreateDate: time.Now(),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchEventNewBulk(t *testing.T) {
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
	err := client.LaunchEvent.NewBulk(context.TODO(), unifieddatalibrary.LaunchEventNewBulkParams{
		Body: []unifieddatalibrary.LaunchEventNewBulkParamsBody{{
			ClassificationMarking:  "U",
			DataMode:               "TEST",
			MsgCreateDate:          time.Now(),
			Source:                 "Bluestaq",
			ID:                     unifieddatalibrary.String("LAUNCHEVENT-ID"),
			BeNumber:               unifieddatalibrary.String("ENC-123"),
			DeclassificationDate:   unifieddatalibrary.Time(time.Now()),
			DeclassificationString: unifieddatalibrary.String("Example Declassification"),
			DerivedFrom:            unifieddatalibrary.String("Example source"),
			LaunchDate:             unifieddatalibrary.Time(time.Now()),
			LaunchFacilityName:     unifieddatalibrary.String("Example launch facility name"),
			LaunchFailureCode:      unifieddatalibrary.String("Example failure code"),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
			OSuffix:                unifieddatalibrary.String("oSuffix"),
			SatNo:                  unifieddatalibrary.Int(12),
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

func TestLaunchEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchEvent.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchEventGetParams{
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

func TestLaunchEventQueryhelp(t *testing.T) {
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
	err := client.LaunchEvent.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchEventTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchEvent.Tuple(context.TODO(), unifieddatalibrary.LaunchEventTupleParams{
		Columns:       "columns",
		MsgCreateDate: time.Now(),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchEventUnvalidatedPublish(t *testing.T) {
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
	err := client.LaunchEvent.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LaunchEventUnvalidatedPublishParams{
		Body: []unifieddatalibrary.LaunchEventUnvalidatedPublishParamsBody{{
			ClassificationMarking:  "U",
			DataMode:               "TEST",
			MsgCreateDate:          time.Now(),
			Source:                 "Bluestaq",
			ID:                     unifieddatalibrary.String("LAUNCHEVENT-ID"),
			BeNumber:               unifieddatalibrary.String("ENC-123"),
			DeclassificationDate:   unifieddatalibrary.Time(time.Now()),
			DeclassificationString: unifieddatalibrary.String("Example Declassification"),
			DerivedFrom:            unifieddatalibrary.String("Example source"),
			LaunchDate:             unifieddatalibrary.Time(time.Now()),
			LaunchFacilityName:     unifieddatalibrary.String("Example launch facility name"),
			LaunchFailureCode:      unifieddatalibrary.String("Example failure code"),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
			OSuffix:                unifieddatalibrary.String("oSuffix"),
			SatNo:                  unifieddatalibrary.Int(12),
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
