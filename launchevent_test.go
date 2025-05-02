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

func TestLauncheventNewWithOptionalParams(t *testing.T) {
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
	err := client.Launchevent.New(context.TODO(), unifieddatalibrary.LauncheventNewParams{
		ClassificationMarking:  "U",
		DataMode:               unifieddatalibrary.LauncheventNewParamsDataModeTest,
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

func TestLauncheventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchevent.List(context.TODO(), unifieddatalibrary.LauncheventListParams{
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

func TestLauncheventCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchevent.Count(context.TODO(), unifieddatalibrary.LauncheventCountParams{
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

func TestLauncheventNewBulk(t *testing.T) {
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
	err := client.Launchevent.NewBulk(context.TODO(), unifieddatalibrary.LauncheventNewBulkParams{
		Body: []unifieddatalibrary.LauncheventNewBulkParamsBody{{
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

func TestLauncheventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchevent.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LauncheventGetParams{
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

func TestLauncheventQueryhelp(t *testing.T) {
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
	err := client.Launchevent.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLauncheventTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchevent.Tuple(context.TODO(), unifieddatalibrary.LauncheventTupleParams{
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

func TestLauncheventUnvalidatedPublish(t *testing.T) {
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
	err := client.Launchevent.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LauncheventUnvalidatedPublishParams{
		Body: []unifieddatalibrary.LauncheventUnvalidatedPublishParamsBody{{
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
