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

func TestAIListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIs.List(context.TODO(), unifieddatalibrary.AIListParams{
		Ts:          time.Now(),
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

func TestAICountWithOptionalParams(t *testing.T) {
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
	_, err := client.AIs.Count(context.TODO(), unifieddatalibrary.AICountParams{
		Ts:          time.Now(),
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

func TestAINewBulk(t *testing.T) {
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
	err := client.AIs.NewBulk(context.TODO(), unifieddatalibrary.AINewBulkParams{
		Body: []unifieddatalibrary.AINewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("AIS-ID"),
			AntennaRefDimensions:  []float64{50.1, 50.1, 20.1, 20.1},
			AvgSpeed:              unifieddatalibrary.Float(12.1),
			CallSign:              unifieddatalibrary.String("V2OZ"),
			CargoType:             unifieddatalibrary.String("Freight"),
			Course:                unifieddatalibrary.Float(157.1),
			CurrentPortGuid:       unifieddatalibrary.String("0ABC"),
			CurrentPortLocode:     unifieddatalibrary.String("XF013"),
			Destination:           unifieddatalibrary.String("USCLE"),
			DestinationEta:        unifieddatalibrary.Time(time.Now()),
			DistanceToGo:          unifieddatalibrary.Float(150.5),
			DistanceTravelled:     unifieddatalibrary.Float(200.3),
			Draught:               unifieddatalibrary.Float(21.1),
			EngagedIn:             unifieddatalibrary.String("Cargo"),
			EtaCalculated:         unifieddatalibrary.Time(time.Now()),
			EtaUpdated:            unifieddatalibrary.Time(time.Now()),
			IDTrack:               unifieddatalibrary.String("TRACK-ID"),
			IDVessel:              unifieddatalibrary.String("VESSEL-ID"),
			Imon:                  unifieddatalibrary.Int(9015462),
			LastPortGuid:          unifieddatalibrary.String("0VAX"),
			LastPortLocode:        unifieddatalibrary.String("USSKY"),
			Lat:                   unifieddatalibrary.Float(47.758499),
			Length:                unifieddatalibrary.Float(511.1),
			Lon:                   unifieddatalibrary.Float(-5.154223),
			MaxSpeed:              unifieddatalibrary.Float(13.3),
			Mmsi:                  unifieddatalibrary.Int(304010417),
			NavStatus:             unifieddatalibrary.String("Underway Using Engine"),
			NextPortGuid:          unifieddatalibrary.String("0Z8Q"),
			NextPortLocode:        unifieddatalibrary.String("USCLE"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PosDeviceType:         unifieddatalibrary.String("GPS"),
			PosHiAccuracy:         unifieddatalibrary.Bool(true),
			PosHiLatency:          unifieddatalibrary.Bool(true),
			RateOfTurn:            unifieddatalibrary.Float(22.1),
			ShipDescription:       unifieddatalibrary.String("Search and rescue vessels"),
			ShipName:              unifieddatalibrary.String("DORNUM"),
			ShipType:              unifieddatalibrary.String("Passenger"),
			SpecialCraft:          unifieddatalibrary.String("Tug"),
			SpecialManeuver:       unifieddatalibrary.Bool(false),
			Speed:                 unifieddatalibrary.Float(10.5),
			TrueHeading:           unifieddatalibrary.Float(329.1),
			VesselFlag:            unifieddatalibrary.String("United States"),
			Width:                 unifieddatalibrary.Float(24.1),
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

func TestAIQueryhelp(t *testing.T) {
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
	_, err := client.AIs.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AIs.Tuple(context.TODO(), unifieddatalibrary.AITupleParams{
		Columns:     "columns",
		Ts:          time.Now(),
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
