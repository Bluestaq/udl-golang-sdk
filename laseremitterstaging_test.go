// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestLaseremitterStagingNewWithOptionalParams(t *testing.T) {
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
	err := client.Laseremitter.Staging.New(context.TODO(), unifieddatalibrary.LaseremitterStagingNewParams{
		ClassificationMarking: "U",
		LaserName:             "LASER_NAME",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
		Altitude:              unifieddatalibrary.Float(1.57543),
		LaserType:             unifieddatalibrary.String("PULSED"),
		Lat:                   unifieddatalibrary.Float(48.6732),
		LocationCountry:       unifieddatalibrary.String("US"),
		Lon:                   unifieddatalibrary.Float(22.8455),
		OwnerCountry:          unifieddatalibrary.String("US"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaseremitterStagingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Laseremitter.Staging.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaseremitterStagingGetParams{
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

func TestLaseremitterStagingUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Laseremitter.Staging.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LaseremitterStagingUpdateParams{
			ClassificationMarking: "U",
			LaserName:             "LASER_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(1.57543),
			LaserType:             unifieddatalibrary.String("PULSED"),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("US"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("US"),
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

func TestLaseremitterStagingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Laseremitter.Staging.List(context.TODO(), unifieddatalibrary.LaseremitterStagingListParams{
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

func TestLaseremitterStagingDelete(t *testing.T) {
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
	err := client.Laseremitter.Staging.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaseremitterStagingNewBulk(t *testing.T) {
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
	err := client.Laseremitter.Staging.NewBulk(context.TODO(), unifieddatalibrary.LaseremitterStagingNewBulkParams{
		Body: []unifieddatalibrary.LaseremitterStagingNewBulkParamsBody{{
			ClassificationMarking: "U",
			LaserName:             "LASER_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(1.57543),
			LaserType:             unifieddatalibrary.String("PULSED"),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("US"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("US"),
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

func TestLaseremitterStagingQueryhelp(t *testing.T) {
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
	_, err := client.Laseremitter.Staging.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
