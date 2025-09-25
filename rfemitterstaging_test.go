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

func TestRfEmitterStagingNewWithOptionalParams(t *testing.T) {
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
	err := client.RfEmitter.Staging.New(context.TODO(), unifieddatalibrary.RfEmitterStagingNewParams{
		ClassificationMarking: "U",
		Name:                  "RF_NAME",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
		Altitude:              unifieddatalibrary.Float(1.57543),
		ExtSysID:              unifieddatalibrary.String("EXTSYS-ID"),
		Lat:                   unifieddatalibrary.Float(48.6732),
		LocationCountry:       unifieddatalibrary.String("US"),
		Lon:                   unifieddatalibrary.Float(22.8455),
		OwnerCountry:          unifieddatalibrary.String("US"),
		Subtype:               unifieddatalibrary.String("BLOCK_1"),
		Type:                  unifieddatalibrary.String("TYPE_OF_EMITTER"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfEmitterStagingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitter.Staging.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RfEmitterStagingGetParams{
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

func TestRfEmitterStagingUpdateWithOptionalParams(t *testing.T) {
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
	err := client.RfEmitter.Staging.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RfEmitterStagingUpdateParams{
			ClassificationMarking: "U",
			Name:                  "RF_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(1.57543),
			ExtSysID:              unifieddatalibrary.String("EXTSYS-ID"),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("US"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("US"),
			Subtype:               unifieddatalibrary.String("BLOCK_1"),
			Type:                  unifieddatalibrary.String("TYPE_OF_EMITTER"),
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

func TestRfEmitterStagingListWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitter.Staging.List(context.TODO(), unifieddatalibrary.RfEmitterStagingListParams{
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

func TestRfEmitterStagingDelete(t *testing.T) {
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
	err := client.RfEmitter.Staging.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfEmitterStagingNewBulk(t *testing.T) {
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
	err := client.RfEmitter.Staging.NewBulk(context.TODO(), unifieddatalibrary.RfEmitterStagingNewBulkParams{
		Body: []unifieddatalibrary.RfEmitterStagingNewBulkParamsBody{{
			ClassificationMarking: "U",
			Name:                  "RF_NAME",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			Altitude:              unifieddatalibrary.Float(1.57543),
			ExtSysID:              unifieddatalibrary.String("EXTSYS-ID"),
			Lat:                   unifieddatalibrary.Float(48.6732),
			LocationCountry:       unifieddatalibrary.String("US"),
			Lon:                   unifieddatalibrary.Float(22.8455),
			OwnerCountry:          unifieddatalibrary.String("US"),
			Subtype:               unifieddatalibrary.String("BLOCK_1"),
			Type:                  unifieddatalibrary.String("TYPE_OF_EMITTER"),
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

func TestRfEmitterStagingQueryhelp(t *testing.T) {
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
	_, err := client.RfEmitter.Staging.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
