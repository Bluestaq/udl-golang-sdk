// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestAirfieldStatusNewWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldStatus.New(context.TODO(), unifieddatalibrary.AirfieldStatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirfieldStatusNewParamsDataModeTest,
		IDAirfield:            "3136498f-2969-3535-1432-e984b2e2e686",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		AltAirfieldID:         unifieddatalibrary.String("AIRFIELD-ID"),
		ApprovedBy:            unifieddatalibrary.String("John Smith"),
		ApprovedDate:          unifieddatalibrary.Time(time.Now()),
		ArffCat:               unifieddatalibrary.String("FAA-A"),
		CargoMog:              unifieddatalibrary.Int(8),
		FleetServiceMog:       unifieddatalibrary.Int(4),
		FuelMog:               unifieddatalibrary.Int(9),
		FuelQtys:              []float64{263083.6, 286674.9, 18143.69},
		FuelTypes:             []string{"JP-8", "Jet A", "AVGAS"},
		GseTime:               unifieddatalibrary.Int(10),
		MedCap:                unifieddatalibrary.String("Large Field Hospital"),
		Message:               unifieddatalibrary.String("Status message about the airfield."),
		MheQtys:               []int64{1, 3, 1},
		MheTypes:              []string{"30k", "AT", "60k"},
		MxMog:                 unifieddatalibrary.Int(3),
		NarrowParkingMog:      unifieddatalibrary.Int(5),
		NarrowWorkingMog:      unifieddatalibrary.Int(4),
		NumCog:                unifieddatalibrary.Int(2),
		OperatingMog:          unifieddatalibrary.Int(4),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PassengerServiceMog:   unifieddatalibrary.Int(5),
		PriFreq:               unifieddatalibrary.Float(123.45),
		PriRwyNum:             unifieddatalibrary.String("35R"),
		ReviewedBy:            unifieddatalibrary.String("Jane Doe"),
		ReviewedDate:          unifieddatalibrary.Time(time.Now()),
		RwyCondReading:        unifieddatalibrary.Int(23),
		RwyFrictionFactor:     unifieddatalibrary.Int(10),
		RwyMarkings:           []string{"Aiming Point", "Threshold"},
		SlotTypesReq:          []string{"PARKING", "WORKING", "LANDING"},
		SurveyDate:            unifieddatalibrary.Time(time.Now()),
		WideParkingMog:        unifieddatalibrary.Int(7),
		WideWorkingMog:        unifieddatalibrary.Int(3),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldStatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldStatus.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldStatusGetParams{
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

func TestAirfieldStatusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldStatus.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldStatusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirfieldStatusUpdateParamsDataModeTest,
			IDAirfield:            "3136498f-2969-3535-1432-e984b2e2e686",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AltAirfieldID:         unifieddatalibrary.String("AIRFIELD-ID"),
			ApprovedBy:            unifieddatalibrary.String("John Smith"),
			ApprovedDate:          unifieddatalibrary.Time(time.Now()),
			ArffCat:               unifieddatalibrary.String("FAA-A"),
			CargoMog:              unifieddatalibrary.Int(8),
			FleetServiceMog:       unifieddatalibrary.Int(4),
			FuelMog:               unifieddatalibrary.Int(9),
			FuelQtys:              []float64{263083.6, 286674.9, 18143.69},
			FuelTypes:             []string{"JP-8", "Jet A", "AVGAS"},
			GseTime:               unifieddatalibrary.Int(10),
			MedCap:                unifieddatalibrary.String("Large Field Hospital"),
			Message:               unifieddatalibrary.String("Status message about the airfield."),
			MheQtys:               []int64{1, 3, 1},
			MheTypes:              []string{"30k", "AT", "60k"},
			MxMog:                 unifieddatalibrary.Int(3),
			NarrowParkingMog:      unifieddatalibrary.Int(5),
			NarrowWorkingMog:      unifieddatalibrary.Int(4),
			NumCog:                unifieddatalibrary.Int(2),
			OperatingMog:          unifieddatalibrary.Int(4),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PassengerServiceMog:   unifieddatalibrary.Int(5),
			PriFreq:               unifieddatalibrary.Float(123.45),
			PriRwyNum:             unifieddatalibrary.String("35R"),
			ReviewedBy:            unifieddatalibrary.String("Jane Doe"),
			ReviewedDate:          unifieddatalibrary.Time(time.Now()),
			RwyCondReading:        unifieddatalibrary.Int(23),
			RwyFrictionFactor:     unifieddatalibrary.Int(10),
			RwyMarkings:           []string{"Aiming Point", "Threshold"},
			SlotTypesReq:          []string{"PARKING", "WORKING", "LANDING"},
			SurveyDate:            unifieddatalibrary.Time(time.Now()),
			WideParkingMog:        unifieddatalibrary.Int(7),
			WideWorkingMog:        unifieddatalibrary.Int(3),
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

func TestAirfieldStatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldStatus.List(context.TODO(), unifieddatalibrary.AirfieldStatusListParams{
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

func TestAirfieldStatusDelete(t *testing.T) {
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
	err := client.AirfieldStatus.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldStatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldStatus.Count(context.TODO(), unifieddatalibrary.AirfieldStatusCountParams{
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

func TestAirfieldStatusQueryhelp(t *testing.T) {
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
	_, err := client.AirfieldStatus.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldStatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldStatus.Tuple(context.TODO(), unifieddatalibrary.AirfieldStatusTupleParams{
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
