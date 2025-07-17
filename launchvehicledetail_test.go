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

func TestLaunchVehicleDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.LaunchVehicleDetails.New(context.TODO(), unifieddatalibrary.LaunchVehicleDetailNewParams{
		ClassificationMarking:               "U",
		DataMode:                            unifieddatalibrary.LaunchVehicleDetailNewParamsDataModeTest,
		IDLaunchVehicle:                     "LAUNCHVEHICLE-ID",
		Source:                              "Bluestaq",
		ID:                                  unifieddatalibrary.String("LAUNCHVEHICLEDETAILS-ID"),
		AttitudeAccuracy:                    unifieddatalibrary.Float(10.23),
		Category:                            unifieddatalibrary.String("Example-category"),
		DeploymentRotationRate:              unifieddatalibrary.Float(10.23),
		Diameter:                            unifieddatalibrary.Float(10.23),
		EstLaunchPrice:                      unifieddatalibrary.Float(10.23),
		EstLaunchPriceTypical:               unifieddatalibrary.Float(10.23),
		FairingExternalDiameter:             unifieddatalibrary.Float(10.23),
		FairingInternalDiameter:             unifieddatalibrary.Float(10.23),
		FairingLength:                       unifieddatalibrary.Float(10.23),
		FairingMass:                         unifieddatalibrary.Float(10.23),
		FairingMaterial:                     unifieddatalibrary.String("Example-fairing-material"),
		FairingName:                         unifieddatalibrary.String("Example-fairing-name"),
		FairingNotes:                        unifieddatalibrary.String("Example notes"),
		Family:                              unifieddatalibrary.String("Example-family"),
		GeoPayloadMass:                      unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyApogeeMargin:      unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyApogeeTarget:      unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyInclinationMargin: unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyInclinationTarget: unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyPerigeeMargin:     unifieddatalibrary.Float(10.23),
		GtoInj3SigAccuracyPerigeeTarget:     unifieddatalibrary.Float(10.23),
		GtoPayloadMass:                      unifieddatalibrary.Float(10.23),
		LaunchMass:                          unifieddatalibrary.Float(10.23),
		LaunchPrefix:                        unifieddatalibrary.String("AX011"),
		Length:                              unifieddatalibrary.Float(10.23),
		LeoPayloadMass:                      unifieddatalibrary.Float(10.23),
		ManufacturerOrgID:                   unifieddatalibrary.String("MANUFACTURERORG-ID"),
		MaxAccelLoad:                        unifieddatalibrary.Float(10.23),
		MaxAcousticLevel:                    unifieddatalibrary.Float(10.23),
		MaxAcousticLevelRange:               unifieddatalibrary.Float(10.23),
		MaxFairingPressureChange:            unifieddatalibrary.Float(10.23),
		MaxFlightShockForce:                 unifieddatalibrary.Float(10.23),
		MaxFlightShockFreq:                  unifieddatalibrary.Float(10.23),
		MaxPayloadFreqLat:                   unifieddatalibrary.Float(10.23),
		MaxPayloadFreqLon:                   unifieddatalibrary.Float(10.23),
		MinorVariant:                        unifieddatalibrary.String("Example-minor-variant"),
		Notes:                               unifieddatalibrary.String("Example notes"),
		Origin:                              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Oxidizer:                            unifieddatalibrary.String("Bromine"),
		PayloadNotes:                        unifieddatalibrary.String("Example notes"),
		PayloadSeparationRate:               unifieddatalibrary.Float(10.23),
		Propellant:                          unifieddatalibrary.String("Nitrogen"),
		SoundPressureLevel:                  unifieddatalibrary.Float(10.23),
		SourceURL:                           unifieddatalibrary.String("Example URL"),
		SSOPayloadMass:                      unifieddatalibrary.Float(10.23),
		Tags:                                []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		Variant:                             unifieddatalibrary.String("Example-variant"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchVehicleDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.LaunchVehicleDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchVehicleDetailUpdateParams{
			ClassificationMarking:               "U",
			DataMode:                            unifieddatalibrary.LaunchVehicleDetailUpdateParamsDataModeTest,
			IDLaunchVehicle:                     "LAUNCHVEHICLE-ID",
			Source:                              "Bluestaq",
			ID:                                  unifieddatalibrary.String("LAUNCHVEHICLEDETAILS-ID"),
			AttitudeAccuracy:                    unifieddatalibrary.Float(10.23),
			Category:                            unifieddatalibrary.String("Example-category"),
			DeploymentRotationRate:              unifieddatalibrary.Float(10.23),
			Diameter:                            unifieddatalibrary.Float(10.23),
			EstLaunchPrice:                      unifieddatalibrary.Float(10.23),
			EstLaunchPriceTypical:               unifieddatalibrary.Float(10.23),
			FairingExternalDiameter:             unifieddatalibrary.Float(10.23),
			FairingInternalDiameter:             unifieddatalibrary.Float(10.23),
			FairingLength:                       unifieddatalibrary.Float(10.23),
			FairingMass:                         unifieddatalibrary.Float(10.23),
			FairingMaterial:                     unifieddatalibrary.String("Example-fairing-material"),
			FairingName:                         unifieddatalibrary.String("Example-fairing-name"),
			FairingNotes:                        unifieddatalibrary.String("Example notes"),
			Family:                              unifieddatalibrary.String("Example-family"),
			GeoPayloadMass:                      unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyApogeeMargin:      unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyApogeeTarget:      unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyInclinationMargin: unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyInclinationTarget: unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyPerigeeMargin:     unifieddatalibrary.Float(10.23),
			GtoInj3SigAccuracyPerigeeTarget:     unifieddatalibrary.Float(10.23),
			GtoPayloadMass:                      unifieddatalibrary.Float(10.23),
			LaunchMass:                          unifieddatalibrary.Float(10.23),
			LaunchPrefix:                        unifieddatalibrary.String("AX011"),
			Length:                              unifieddatalibrary.Float(10.23),
			LeoPayloadMass:                      unifieddatalibrary.Float(10.23),
			ManufacturerOrgID:                   unifieddatalibrary.String("MANUFACTURERORG-ID"),
			MaxAccelLoad:                        unifieddatalibrary.Float(10.23),
			MaxAcousticLevel:                    unifieddatalibrary.Float(10.23),
			MaxAcousticLevelRange:               unifieddatalibrary.Float(10.23),
			MaxFairingPressureChange:            unifieddatalibrary.Float(10.23),
			MaxFlightShockForce:                 unifieddatalibrary.Float(10.23),
			MaxFlightShockFreq:                  unifieddatalibrary.Float(10.23),
			MaxPayloadFreqLat:                   unifieddatalibrary.Float(10.23),
			MaxPayloadFreqLon:                   unifieddatalibrary.Float(10.23),
			MinorVariant:                        unifieddatalibrary.String("Example-minor-variant"),
			Notes:                               unifieddatalibrary.String("Example notes"),
			Origin:                              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Oxidizer:                            unifieddatalibrary.String("Bromine"),
			PayloadNotes:                        unifieddatalibrary.String("Example notes"),
			PayloadSeparationRate:               unifieddatalibrary.Float(10.23),
			Propellant:                          unifieddatalibrary.String("Nitrogen"),
			SoundPressureLevel:                  unifieddatalibrary.Float(10.23),
			SourceURL:                           unifieddatalibrary.String("Example URL"),
			SSOPayloadMass:                      unifieddatalibrary.Float(10.23),
			Tags:                                []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			Variant:                             unifieddatalibrary.String("Example-variant"),
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

func TestLaunchVehicleDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchVehicleDetails.List(context.TODO(), unifieddatalibrary.LaunchVehicleDetailListParams{
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

func TestLaunchVehicleDetailDelete(t *testing.T) {
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
	err := client.LaunchVehicleDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaunchVehicleDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LaunchVehicleDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaunchVehicleDetailGetParams{
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
