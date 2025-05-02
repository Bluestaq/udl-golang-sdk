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

func TestSeradatasigintpayloadNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradatasigintpayload.New(context.TODO(), unifieddatalibrary.SeradatasigintpayloadNewParams{
		ClassificationMarking:  "U",
		DataMode:               unifieddatalibrary.SeradatasigintpayloadNewParamsDataModeTest,
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

func TestSeradatasigintpayloadUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradatasigintpayload.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatasigintpayloadUpdateParams{
			ClassificationMarking:  "U",
			DataMode:               unifieddatalibrary.SeradatasigintpayloadUpdateParamsDataModeTest,
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

func TestSeradatasigintpayloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatasigintpayload.List(context.TODO(), unifieddatalibrary.SeradatasigintpayloadListParams{
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

func TestSeradatasigintpayloadDelete(t *testing.T) {
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
	err := client.Seradatasigintpayload.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatasigintpayloadCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatasigintpayload.Count(context.TODO(), unifieddatalibrary.SeradatasigintpayloadCountParams{
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

func TestSeradatasigintpayloadGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatasigintpayload.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatasigintpayloadGetParams{
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

func TestSeradatasigintpayloadQueryhelp(t *testing.T) {
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
	err := client.Seradatasigintpayload.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatasigintpayloadTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatasigintpayload.Tuple(context.TODO(), unifieddatalibrary.SeradatasigintpayloadTupleParams{
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
