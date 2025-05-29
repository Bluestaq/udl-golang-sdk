// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestSeradataSigintPayloadNewWithOptionalParams(t *testing.T) {
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
	err := client.SeradataSigintPayload.New(context.TODO(), unifieddatalibrary.SeradataSigintPayloadNewParams{
		ClassificationMarking:  "U",
		DataMode:               unifieddatalibrary.SeradataSigintPayloadNewParamsDataModeTest,
		Source:                 "Bluestaq",
		SpacecraftID:           "spacecraftId",
		ID:                     unifieddatalibrary.String("SERADATASIGINTPAYLOAD-ID"),
		FrequencyCoverage:      unifieddatalibrary.String("1.1 to 3.3"),
		GroundStationLocations: unifieddatalibrary.String("groundStationLocations"),
		GroundStations:         unifieddatalibrary.String("groundStations"),
		HostedForCompanyOrgID:  unifieddatalibrary.String("hostedForCompanyOrgId"),
		IDSensor:               unifieddatalibrary.String("0c5ec9c0-10cd-1d35-c46b-3764c4d76e13"),
		InterceptParameters:    unifieddatalibrary.String("interceptParameters"),
		ManufacturerOrgID:      unifieddatalibrary.String("manufacturerOrgId"),
		Name:                   unifieddatalibrary.String("Sensor Name"),
		Notes:                  unifieddatalibrary.String("Sample Notes"),
		Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PositionalAccuracy:     unifieddatalibrary.String("positionalAccuracy"),
		SwathWidth:             unifieddatalibrary.Float(1.23),
		Type:                   unifieddatalibrary.String("Comint"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataSigintPayloadUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeradataSigintPayload.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataSigintPayloadUpdateParams{
			ClassificationMarking:  "U",
			DataMode:               unifieddatalibrary.SeradataSigintPayloadUpdateParamsDataModeTest,
			Source:                 "Bluestaq",
			SpacecraftID:           "spacecraftId",
			ID:                     unifieddatalibrary.String("SERADATASIGINTPAYLOAD-ID"),
			FrequencyCoverage:      unifieddatalibrary.String("1.1 to 3.3"),
			GroundStationLocations: unifieddatalibrary.String("groundStationLocations"),
			GroundStations:         unifieddatalibrary.String("groundStations"),
			HostedForCompanyOrgID:  unifieddatalibrary.String("hostedForCompanyOrgId"),
			IDSensor:               unifieddatalibrary.String("0c5ec9c0-10cd-1d35-c46b-3764c4d76e13"),
			InterceptParameters:    unifieddatalibrary.String("interceptParameters"),
			ManufacturerOrgID:      unifieddatalibrary.String("manufacturerOrgId"),
			Name:                   unifieddatalibrary.String("Sensor Name"),
			Notes:                  unifieddatalibrary.String("Sample Notes"),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PositionalAccuracy:     unifieddatalibrary.String("positionalAccuracy"),
			SwathWidth:             unifieddatalibrary.Float(1.23),
			Type:                   unifieddatalibrary.String("Comint"),
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

func TestSeradataSigintPayloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataSigintPayload.List(context.TODO(), unifieddatalibrary.SeradataSigintPayloadListParams{
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

func TestSeradataSigintPayloadDelete(t *testing.T) {
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
	err := client.SeradataSigintPayload.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataSigintPayloadCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataSigintPayload.Count(context.TODO(), unifieddatalibrary.SeradataSigintPayloadCountParams{
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

func TestSeradataSigintPayloadGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataSigintPayload.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataSigintPayloadGetParams{
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

func TestSeradataSigintPayloadQueryhelp(t *testing.T) {
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
	_, err := client.SeradataSigintPayload.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataSigintPayloadTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeradataSigintPayload.Tuple(context.TODO(), unifieddatalibrary.SeradataSigintPayloadTupleParams{
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
