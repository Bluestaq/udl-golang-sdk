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

func TestLaunchdetectionNewWithOptionalParams(t *testing.T) {
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
	err := client.Launchdetection.New(context.TODO(), unifieddatalibrary.LaunchdetectionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LaunchdetectionNewParamsDataModeTest,
		MessageType:           "Example-Msg-Type",
		ObservationLatitude:   45.23,
		ObservationLongitude:  1.23,
		ObservationTime:       time.Now(),
		SequenceNumber:        5,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("LAUNCHDETECTION-ID"),
		Descriptor:            unifieddatalibrary.String("Example descriptor"),
		EventID:               unifieddatalibrary.String("EVENT-ID"),
		HighZenithAzimuth:     unifieddatalibrary.Bool(false),
		Inclination:           unifieddatalibrary.Float(1.23),
		LaunchAzimuth:         unifieddatalibrary.Float(1.23),
		LaunchLatitude:        unifieddatalibrary.Float(1.23),
		LaunchLongitude:       unifieddatalibrary.Float(1.23),
		LaunchTime:            unifieddatalibrary.Time(time.Now()),
		ObservationAltitude:   unifieddatalibrary.Float(1.23),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Raan:                  unifieddatalibrary.Float(1.23),
		StereoFlag:            unifieddatalibrary.Bool(false),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchdetectionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Launchdetection.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchdetectionUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LaunchdetectionUpdateParamsDataModeTest,
			MessageType:           "Example-Msg-Type",
			ObservationLatitude:   45.23,
			ObservationLongitude:  1.23,
			ObservationTime:       time.Now(),
			SequenceNumber:        5,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LAUNCHDETECTION-ID"),
			Descriptor:            unifieddatalibrary.String("Example descriptor"),
			EventID:               unifieddatalibrary.String("EVENT-ID"),
			HighZenithAzimuth:     unifieddatalibrary.Bool(false),
			Inclination:           unifieddatalibrary.Float(1.23),
			LaunchAzimuth:         unifieddatalibrary.Float(1.23),
			LaunchLatitude:        unifieddatalibrary.Float(1.23),
			LaunchLongitude:       unifieddatalibrary.Float(1.23),
			LaunchTime:            unifieddatalibrary.Time(time.Now()),
			ObservationAltitude:   unifieddatalibrary.Float(1.23),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Raan:                  unifieddatalibrary.Float(1.23),
			StereoFlag:            unifieddatalibrary.Bool(false),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
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

func TestLaunchdetectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchdetection.List(context.TODO(), unifieddatalibrary.LaunchdetectionListParams{
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

func TestLaunchdetectionDelete(t *testing.T) {
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
	err := client.Launchdetection.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchdetectionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchdetection.Count(context.TODO(), unifieddatalibrary.LaunchdetectionCountParams{
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

func TestLaunchdetectionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchdetection.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchdetectionGetParams{
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

func TestLaunchdetectionQueryhelp(t *testing.T) {
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
	err := client.Launchdetection.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchdetectionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Launchdetection.Tuple(context.TODO(), unifieddatalibrary.LaunchdetectionTupleParams{
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
