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

func TestOnorbitthrusterstatusNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitthrusterstatus.New(context.TODO(), unifieddatalibrary.OnorbitthrusterstatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitthrusterstatusNewParamsDataModeTest,
		IDOnorbitThruster:     "ff7dc909-e8b4-4a54-8529-1963d4e9b353",
		Source:                "Bluestaq",
		StatusTime:            time.Now(),
		ID:                    unifieddatalibrary.String("af103c-1f917dc-002c1bd"),
		EstDeltaV:             unifieddatalibrary.Float(10.1),
		FuelMass:              unifieddatalibrary.Float(100.1),
		FuelMassUnc:           unifieddatalibrary.Float(10.1),
		Isp:                   unifieddatalibrary.Float(300.1),
		MaxDeltaV:             unifieddatalibrary.Float(100.1),
		MinDeltaV:             unifieddatalibrary.Float(0.1),
		Name:                  unifieddatalibrary.String("REA1"),
		Operational:           unifieddatalibrary.Bool(true),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PropMassAvg:           unifieddatalibrary.Float(907.6),
		PropMassMax:           unifieddatalibrary.Float(2333.3),
		PropMassMedian:        unifieddatalibrary.Float(200.1),
		PropMassMin:           unifieddatalibrary.Float(0.1),
		ThrustMax:             unifieddatalibrary.Float(22.1),
		TotalDeltaV:           unifieddatalibrary.Float(100.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterstatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthrusterstatus.List(context.TODO(), unifieddatalibrary.OnorbitthrusterstatusListParams{
		FirstResult:       unifieddatalibrary.Int(0),
		IDOnorbitThruster: unifieddatalibrary.String("idOnorbitThruster"),
		MaxResults:        unifieddatalibrary.Int(0),
		StatusTime:        unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterstatusDelete(t *testing.T) {
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
	err := client.Onorbitthrusterstatus.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterstatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthrusterstatus.Count(context.TODO(), unifieddatalibrary.OnorbitthrusterstatusCountParams{
		FirstResult:       unifieddatalibrary.Int(0),
		IDOnorbitThruster: unifieddatalibrary.String("idOnorbitThruster"),
		MaxResults:        unifieddatalibrary.Int(0),
		StatusTime:        unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterstatusNewBulk(t *testing.T) {
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
	err := client.Onorbitthrusterstatus.NewBulk(context.TODO(), unifieddatalibrary.OnorbitthrusterstatusNewBulkParams{
		Body: []unifieddatalibrary.OnorbitthrusterstatusNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDOnorbitThruster:     "ff7dc909-e8b4-4a54-8529-1963d4e9b353",
			Source:                "Bluestaq",
			StatusTime:            time.Now(),
			ID:                    unifieddatalibrary.String("af103c-1f917dc-002c1bd"),
			EstDeltaV:             unifieddatalibrary.Float(10.1),
			FuelMass:              unifieddatalibrary.Float(100.1),
			FuelMassUnc:           unifieddatalibrary.Float(10.1),
			Isp:                   unifieddatalibrary.Float(300.1),
			MaxDeltaV:             unifieddatalibrary.Float(100.1),
			MinDeltaV:             unifieddatalibrary.Float(0.1),
			Name:                  unifieddatalibrary.String("REA1"),
			Operational:           unifieddatalibrary.Bool(true),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PropMassAvg:           unifieddatalibrary.Float(907.6),
			PropMassMax:           unifieddatalibrary.Float(2333.3),
			PropMassMedian:        unifieddatalibrary.Float(200.1),
			PropMassMin:           unifieddatalibrary.Float(0.1),
			ThrustMax:             unifieddatalibrary.Float(22.1),
			TotalDeltaV:           unifieddatalibrary.Float(100.1),
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

func TestOnorbitthrusterstatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthrusterstatus.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitthrusterstatusGetParams{
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

func TestOnorbitthrusterstatusQueryhelp(t *testing.T) {
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
	_, err := client.Onorbitthrusterstatus.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitthrusterstatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitthrusterstatus.Tuple(context.TODO(), unifieddatalibrary.OnorbitthrusterstatusTupleParams{
		Columns:           "columns",
		FirstResult:       unifieddatalibrary.Int(0),
		IDOnorbitThruster: unifieddatalibrary.String("idOnorbitThruster"),
		MaxResults:        unifieddatalibrary.Int(0),
		StatusTime:        unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
