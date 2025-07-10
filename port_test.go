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

func TestPortNewWithOptionalParams(t *testing.T) {
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
	err := client.Port.New(context.TODO(), unifieddatalibrary.PortNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.PortNewParamsDataModeTest,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AvgDuration:           unifieddatalibrary.Float(41.1),
		CountryCode:           unifieddatalibrary.String("US"),
		ExternalID:            unifieddatalibrary.String("fe4ad5dc-0128-4ce8-b09c-0b404322025e"),
		HarborSize:            unifieddatalibrary.Float(160.1),
		HarborType:            unifieddatalibrary.String("COASTAL NATURAL"),
		IDSite:                unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
		Lat:                   unifieddatalibrary.Float(45.23),
		Locode:                unifieddatalibrary.String("CAVAN"),
		Lon:                   unifieddatalibrary.Float(179.1),
		MaxDraught:            unifieddatalibrary.Float(18.1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PilotReqd:             unifieddatalibrary.Bool(true),
		PortName:              unifieddatalibrary.String("Vancouver"),
		Shelter:               unifieddatalibrary.String("EXCELLENT"),
		TideRange:             unifieddatalibrary.Float(4.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Port.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.PortUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.PortUpdateParamsDataModeTest,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AvgDuration:           unifieddatalibrary.Float(41.1),
			CountryCode:           unifieddatalibrary.String("US"),
			ExternalID:            unifieddatalibrary.String("fe4ad5dc-0128-4ce8-b09c-0b404322025e"),
			HarborSize:            unifieddatalibrary.Float(160.1),
			HarborType:            unifieddatalibrary.String("COASTAL NATURAL"),
			IDSite:                unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			Lat:                   unifieddatalibrary.Float(45.23),
			Locode:                unifieddatalibrary.String("CAVAN"),
			Lon:                   unifieddatalibrary.Float(179.1),
			MaxDraught:            unifieddatalibrary.Float(18.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PilotReqd:             unifieddatalibrary.Bool(true),
			PortName:              unifieddatalibrary.String("Vancouver"),
			Shelter:               unifieddatalibrary.String("EXCELLENT"),
			TideRange:             unifieddatalibrary.Float(4.1),
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

func TestPortListWithOptionalParams(t *testing.T) {
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
	_, err := client.Port.List(context.TODO(), unifieddatalibrary.PortListParams{
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

func TestPortCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Port.Count(context.TODO(), unifieddatalibrary.PortCountParams{
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

func TestPortNewBulk(t *testing.T) {
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
	err := client.Port.NewBulk(context.TODO(), unifieddatalibrary.PortNewBulkParams{
		Body: []unifieddatalibrary.PortNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AvgDuration:           unifieddatalibrary.Float(41.1),
			CountryCode:           unifieddatalibrary.String("US"),
			ExternalID:            unifieddatalibrary.String("fe4ad5dc-0128-4ce8-b09c-0b404322025e"),
			HarborSize:            unifieddatalibrary.Float(160.1),
			HarborType:            unifieddatalibrary.String("COASTAL NATURAL"),
			IDSite:                unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			Lat:                   unifieddatalibrary.Float(45.23),
			Locode:                unifieddatalibrary.String("CAVAN"),
			Lon:                   unifieddatalibrary.Float(179.1),
			MaxDraught:            unifieddatalibrary.Float(18.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PilotReqd:             unifieddatalibrary.Bool(true),
			PortName:              unifieddatalibrary.String("Vancouver"),
			Shelter:               unifieddatalibrary.String("EXCELLENT"),
			TideRange:             unifieddatalibrary.Float(4.1),
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

func TestPortGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Port.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.PortGetParams{
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

func TestPortQueryhelp(t *testing.T) {
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
	_, err := client.Port.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Port.Tuple(context.TODO(), unifieddatalibrary.PortTupleParams{
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
