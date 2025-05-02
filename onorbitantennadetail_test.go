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

func TestOnorbitAntennaDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbit.AntennaDetails.New(context.TODO(), unifieddatalibrary.OnorbitAntennaDetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitAntennaDetailNewParamsDataModeTest,
		IDAntenna:             "ANTENNA-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ANTENNADETAILS-ID"),
		BeamForming:           unifieddatalibrary.Bool(false),
		Beamwidth:             unifieddatalibrary.Float(14.1),
		Description:           unifieddatalibrary.String("Description of antenna A"),
		Diameter:              unifieddatalibrary.Float(0.01),
		EndFrequency:          unifieddatalibrary.Float(3.3),
		Gain:                  unifieddatalibrary.Float(20.1),
		GainTolerance:         unifieddatalibrary.Float(5.1),
		ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTUREORG-ID"),
		Mode:                  unifieddatalibrary.OnorbitAntennaDetailNewParamsModeTx,
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Polarization:          unifieddatalibrary.Float(45.1),
		Position:              unifieddatalibrary.String("Top"),
		Size:                  []float64{0.03, 0.05},
		StartFrequency:        unifieddatalibrary.Float(2.1),
		Steerable:             unifieddatalibrary.Bool(false),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		Type:                  unifieddatalibrary.String("Reflector"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitAntennaDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbit.AntennaDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitAntennaDetailGetParams{
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

func TestOnorbitAntennaDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbit.AntennaDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitAntennaDetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OnorbitAntennaDetailUpdateParamsDataModeTest,
			IDAntenna:             "ANTENNA-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ANTENNADETAILS-ID"),
			BeamForming:           unifieddatalibrary.Bool(false),
			Beamwidth:             unifieddatalibrary.Float(14.1),
			Description:           unifieddatalibrary.String("Description of antenna A"),
			Diameter:              unifieddatalibrary.Float(0.01),
			EndFrequency:          unifieddatalibrary.Float(3.3),
			Gain:                  unifieddatalibrary.Float(20.1),
			GainTolerance:         unifieddatalibrary.Float(5.1),
			ManufacturerOrgID:     unifieddatalibrary.String("MANUFACTUREORG-ID"),
			Mode:                  unifieddatalibrary.OnorbitAntennaDetailUpdateParamsModeTx,
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Polarization:          unifieddatalibrary.Float(45.1),
			Position:              unifieddatalibrary.String("Top"),
			Size:                  []float64{0.03, 0.05},
			StartFrequency:        unifieddatalibrary.Float(2.1),
			Steerable:             unifieddatalibrary.Bool(false),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			Type:                  unifieddatalibrary.String("Reflector"),
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

func TestOnorbitAntennaDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbit.AntennaDetails.List(context.TODO(), unifieddatalibrary.OnorbitAntennaDetailListParams{
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

func TestOnorbitAntennaDetailDelete(t *testing.T) {
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
	err := client.Onorbit.AntennaDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
