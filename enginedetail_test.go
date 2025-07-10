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

func TestEngineDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.EngineDetails.New(context.TODO(), unifieddatalibrary.EngineDetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EngineDetailNewParamsDataModeTest,
		IDEngine:              "ENGINE-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ENGINEDETAILS-ID"),
		BurnTime:              unifieddatalibrary.Float(1.1),
		ChamberPressure:       unifieddatalibrary.Float(1.1),
		CharacteristicType:    unifieddatalibrary.String("Electric"),
		CycleType:             unifieddatalibrary.String("Pressure Fed"),
		Family:                unifieddatalibrary.String("ENGINE_TYPE1"),
		ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
		MaxFirings:            unifieddatalibrary.Int(5),
		Notes:                 unifieddatalibrary.String("Example notes"),
		NozzleExpansionRatio:  unifieddatalibrary.Float(1.1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Oxidizer:              unifieddatalibrary.String("Liquid Oxygen"),
		Propellant:            unifieddatalibrary.String("Liquid"),
		SeaLevelThrust:        unifieddatalibrary.Float(1.1),
		SpecificImpulse:       unifieddatalibrary.Float(1.1),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		VacuumThrust:          unifieddatalibrary.Float(1.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEngineDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EngineDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EngineDetailGetParams{
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

func TestEngineDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.EngineDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.EngineDetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.EngineDetailUpdateParamsDataModeTest,
			IDEngine:              "ENGINE-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ENGINEDETAILS-ID"),
			BurnTime:              unifieddatalibrary.Float(1.1),
			ChamberPressure:       unifieddatalibrary.Float(1.1),
			CharacteristicType:    unifieddatalibrary.String("Electric"),
			CycleType:             unifieddatalibrary.String("Pressure Fed"),
			Family:                unifieddatalibrary.String("ENGINE_TYPE1"),
			ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTURERORG-ID"),
			MaxFirings:            unifieddatalibrary.Int(5),
			Notes:                 unifieddatalibrary.String("Example notes"),
			NozzleExpansionRatio:  unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Oxidizer:              unifieddatalibrary.String("Liquid Oxygen"),
			Propellant:            unifieddatalibrary.String("Liquid"),
			SeaLevelThrust:        unifieddatalibrary.Float(1.1),
			SpecificImpulse:       unifieddatalibrary.Float(1.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			VacuumThrust:          unifieddatalibrary.Float(1.1),
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

func TestEngineDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.EngineDetails.List(context.TODO(), unifieddatalibrary.EngineDetailListParams{
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

func TestEngineDetailDelete(t *testing.T) {
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
	err := client.EngineDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
