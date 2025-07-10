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

func TestLaunchDetectionNewWithOptionalParams(t *testing.T) {
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
	err := client.LaunchDetection.New(context.TODO(), unifieddatalibrary.LaunchDetectionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LaunchDetectionNewParamsDataModeTest,
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

func TestLaunchDetectionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.LaunchDetection.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchDetectionUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LaunchDetectionUpdateParamsDataModeTest,
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

func TestLaunchDetectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchDetection.List(context.TODO(), unifieddatalibrary.LaunchDetectionListParams{
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

func TestLaunchDetectionDelete(t *testing.T) {
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
	err := client.LaunchDetection.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchDetectionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchDetection.Count(context.TODO(), unifieddatalibrary.LaunchDetectionCountParams{
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

func TestLaunchDetectionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchDetection.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchDetectionGetParams{
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

func TestLaunchDetectionQueryhelp(t *testing.T) {
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
	_, err := client.LaunchDetection.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchDetectionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchDetection.Tuple(context.TODO(), unifieddatalibrary.LaunchDetectionTupleParams{
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
