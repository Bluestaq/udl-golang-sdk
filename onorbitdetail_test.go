// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestOnorbitdetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitdetails.New(context.TODO(), unifieddatalibrary.OnorbitdetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitdetailNewParamsDataModeTest,
		IDOnOrbit:             "REF-ONORBIT-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ONORBITDETAILS-ID"),
		AdditionalMass:        unifieddatalibrary.Float(10.23),
		AdeptRadius:           unifieddatalibrary.Float(10.23),
		BolDeltaV:             unifieddatalibrary.Float(1000.1),
		BolFuelMass:           unifieddatalibrary.Float(10.23),
		BusCrossSection:       unifieddatalibrary.Float(10.23),
		BusType:               unifieddatalibrary.String("A2100"),
		ColaRadius:            unifieddatalibrary.Float(10.23),
		CrossSection:          unifieddatalibrary.Float(10.23),
		CurrentMass:           unifieddatalibrary.Float(500),
		DeltaVUnc:             unifieddatalibrary.Float(50.1),
		DepEstMasses:          []float64{20, 21},
		DepMassUncs:           []float64{10, 5},
		DepNames:              []string{"GOES-18A", "GOES-18B"},
		DriftRate:             unifieddatalibrary.Float(1.23),
		DryMass:               unifieddatalibrary.Float(10.23),
		EstDeltaVDuration:     unifieddatalibrary.Float(10.23),
		FuelRemaining:         unifieddatalibrary.Float(10.23),
		GeoSlot:               unifieddatalibrary.Float(90.23),
		LastObSource:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		LastObTime:            unifieddatalibrary.Time(time.Now()),
		LaunchMass:            unifieddatalibrary.Float(10.23),
		LaunchMassMax:         unifieddatalibrary.Float(15.23),
		LaunchMassMin:         unifieddatalibrary.Float(5.23),
		Maneuverable:          unifieddatalibrary.Bool(false),
		MaxDeltaV:             unifieddatalibrary.Float(10.23),
		MaxRadius:             unifieddatalibrary.Float(10.23),
		MissionTypes:          []string{"Weather", "Space Weather"},
		NumDeployable:         unifieddatalibrary.Int(2),
		NumMission:            unifieddatalibrary.Int(2),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Rcs:                   unifieddatalibrary.Float(10.23),
		RcsMax:                unifieddatalibrary.Float(15.23),
		RcsMean:               unifieddatalibrary.Float(10.23),
		RcsMin:                unifieddatalibrary.Float(5.23),
		RefSource:             unifieddatalibrary.String("Wikipedia"),
		SolarArrayArea:        unifieddatalibrary.Float(10.23),
		TotalMassUnc:          unifieddatalibrary.Float(50.1),
		Vismag:                unifieddatalibrary.Float(10.23),
		VismagMax:             unifieddatalibrary.Float(15.23),
		VismagMean:            unifieddatalibrary.Float(10.23),
		VismagMin:             unifieddatalibrary.Float(5.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitdetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitdetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitdetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OnorbitdetailUpdateParamsDataModeTest,
			IDOnOrbit:             "REF-ONORBIT-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ONORBITDETAILS-ID"),
			AdditionalMass:        unifieddatalibrary.Float(10.23),
			AdeptRadius:           unifieddatalibrary.Float(10.23),
			BolDeltaV:             unifieddatalibrary.Float(1000.1),
			BolFuelMass:           unifieddatalibrary.Float(10.23),
			BusCrossSection:       unifieddatalibrary.Float(10.23),
			BusType:               unifieddatalibrary.String("A2100"),
			ColaRadius:            unifieddatalibrary.Float(10.23),
			CrossSection:          unifieddatalibrary.Float(10.23),
			CurrentMass:           unifieddatalibrary.Float(500),
			DeltaVUnc:             unifieddatalibrary.Float(50.1),
			DepEstMasses:          []float64{20, 21},
			DepMassUncs:           []float64{10, 5},
			DepNames:              []string{"GOES-18A", "GOES-18B"},
			DriftRate:             unifieddatalibrary.Float(1.23),
			DryMass:               unifieddatalibrary.Float(10.23),
			EstDeltaVDuration:     unifieddatalibrary.Float(10.23),
			FuelRemaining:         unifieddatalibrary.Float(10.23),
			GeoSlot:               unifieddatalibrary.Float(90.23),
			LastObSource:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			LastObTime:            unifieddatalibrary.Time(time.Now()),
			LaunchMass:            unifieddatalibrary.Float(10.23),
			LaunchMassMax:         unifieddatalibrary.Float(15.23),
			LaunchMassMin:         unifieddatalibrary.Float(5.23),
			Maneuverable:          unifieddatalibrary.Bool(false),
			MaxDeltaV:             unifieddatalibrary.Float(10.23),
			MaxRadius:             unifieddatalibrary.Float(10.23),
			MissionTypes:          []string{"Weather", "Space Weather"},
			NumDeployable:         unifieddatalibrary.Int(2),
			NumMission:            unifieddatalibrary.Int(2),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Rcs:                   unifieddatalibrary.Float(10.23),
			RcsMax:                unifieddatalibrary.Float(15.23),
			RcsMean:               unifieddatalibrary.Float(10.23),
			RcsMin:                unifieddatalibrary.Float(5.23),
			RefSource:             unifieddatalibrary.String("Wikipedia"),
			SolarArrayArea:        unifieddatalibrary.Float(10.23),
			TotalMassUnc:          unifieddatalibrary.Float(50.1),
			Vismag:                unifieddatalibrary.Float(10.23),
			VismagMax:             unifieddatalibrary.Float(15.23),
			VismagMean:            unifieddatalibrary.Float(10.23),
			VismagMin:             unifieddatalibrary.Float(5.23),
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

func TestOnorbitdetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitdetails.List(context.TODO(), unifieddatalibrary.OnorbitdetailListParams{
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

func TestOnorbitdetailDelete(t *testing.T) {
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
	err := client.Onorbitdetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitdetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitdetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitdetailGetParams{
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
