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

func TestAirfieldSlotNewWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldSlots.New(context.TODO(), unifieddatalibrary.AirfieldSlotNewParams{
		AirfieldName:          "USAF Academy AFLD",
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirfieldSlotNewParamsDataModeTest,
		Name:                  "Apron 5",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		AcSlotCat:             unifieddatalibrary.AirfieldSlotNewParamsAcSlotCatWide,
		AltAirfieldID:         unifieddatalibrary.String("ALT-AIRFIELD-ID"),
		Capacity:              unifieddatalibrary.Int(5),
		EndTime:               unifieddatalibrary.String("2359Z"),
		Icao:                  unifieddatalibrary.String("KCOS"),
		IDAirfield:            unifieddatalibrary.String("3136498f-2969-3535-1432-e984b2e2e686"),
		MinSeparation:         unifieddatalibrary.Int(7),
		Notes:                 unifieddatalibrary.String("Notes for an airfield slot."),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		StartTime:             unifieddatalibrary.String("0000Z"),
		Type:                  unifieddatalibrary.AirfieldSlotNewParamsTypeWorking,
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldSlotGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlots.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldSlotGetParams{
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

func TestAirfieldSlotUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldSlots.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldSlotUpdateParams{
			AirfieldName:          "USAF Academy AFLD",
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirfieldSlotUpdateParamsDataModeTest,
			Name:                  "Apron 5",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AcSlotCat:             unifieddatalibrary.AirfieldSlotUpdateParamsAcSlotCatWide,
			AltAirfieldID:         unifieddatalibrary.String("ALT-AIRFIELD-ID"),
			Capacity:              unifieddatalibrary.Int(5),
			EndTime:               unifieddatalibrary.String("2359Z"),
			Icao:                  unifieddatalibrary.String("KCOS"),
			IDAirfield:            unifieddatalibrary.String("3136498f-2969-3535-1432-e984b2e2e686"),
			MinSeparation:         unifieddatalibrary.Int(7),
			Notes:                 unifieddatalibrary.String("Notes for an airfield slot."),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			StartTime:             unifieddatalibrary.String("0000Z"),
			Type:                  unifieddatalibrary.AirfieldSlotUpdateParamsTypeWorking,
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

func TestAirfieldSlotListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlots.List(context.TODO(), unifieddatalibrary.AirfieldSlotListParams{
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

func TestAirfieldSlotDelete(t *testing.T) {
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
	err := client.AirfieldSlots.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldSlotCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlots.Count(context.TODO(), unifieddatalibrary.AirfieldSlotCountParams{
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

func TestAirfieldSlotQueryhelp(t *testing.T) {
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
	_, err := client.AirfieldSlots.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldSlotTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlots.Tuple(context.TODO(), unifieddatalibrary.AirfieldSlotTupleParams{
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
