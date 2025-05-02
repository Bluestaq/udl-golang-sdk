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

func TestAirfieldslotconsumptionNewWithOptionalParams(t *testing.T) {
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
	err := client.Airfieldslotconsumptions.New(context.TODO(), unifieddatalibrary.AirfieldslotconsumptionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirfieldslotconsumptionNewParamsDataModeTest,
		IDAirfieldSlot:        "3136498f-2969-3535-1432-e984b2e2e686",
		NumAircraft:           1,
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		AltArrSortieID:        unifieddatalibrary.String("ALT-SORTIE-ID"),
		AltDepSortieID:        unifieddatalibrary.String("ALT-SORTIE-ID"),
		AppComment:            unifieddatalibrary.String("The request was denied due to inoperable fuel pumps."),
		AppInitials:           unifieddatalibrary.String("CB"),
		AppOrg:                unifieddatalibrary.String("KCHS/BOPS"),
		CallSigns:             []string{"RCH123", "ABC123", "LLS442"},
		Consumer:              unifieddatalibrary.String("APRON1-230401001"),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		IDArrSortie:           unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		IDDepSortie:           unifieddatalibrary.String("1e6edeec-72e9-aaec-d33c-51147cb5ffdd"),
		MissionID:             unifieddatalibrary.String("AJM123456123"),
		OccAircraftMds:        unifieddatalibrary.String("C017A"),
		OccStartTime:          unifieddatalibrary.Time(time.Now()),
		OccTailNumber:         unifieddatalibrary.String("N702JG"),
		Occupied:              unifieddatalibrary.Bool(true),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ReqComment:            unifieddatalibrary.String("Sorry for the late notice."),
		ReqInitials:           unifieddatalibrary.String("CB"),
		ReqOrg:                unifieddatalibrary.String("TACC"),
		ResAircraftMds:        unifieddatalibrary.String("C017A"),
		ResMissionID:          unifieddatalibrary.String("AJM123456123"),
		ResReason:             unifieddatalibrary.String("Maintenance needed"),
		ResTailNumber:         unifieddatalibrary.String("N702JG"),
		ResType:               unifieddatalibrary.String("M"),
		Status:                unifieddatalibrary.AirfieldslotconsumptionNewParamsStatusApproved,
		TargetTime:            unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldslotconsumptionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfieldslotconsumptions.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldslotconsumptionGetParams{
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

func TestAirfieldslotconsumptionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Airfieldslotconsumptions.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldslotconsumptionUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirfieldslotconsumptionUpdateParamsDataModeTest,
			IDAirfieldSlot:        "3136498f-2969-3535-1432-e984b2e2e686",
			NumAircraft:           1,
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AltArrSortieID:        unifieddatalibrary.String("ALT-SORTIE-ID"),
			AltDepSortieID:        unifieddatalibrary.String("ALT-SORTIE-ID"),
			AppComment:            unifieddatalibrary.String("The request was denied due to inoperable fuel pumps."),
			AppInitials:           unifieddatalibrary.String("CB"),
			AppOrg:                unifieddatalibrary.String("KCHS/BOPS"),
			CallSigns:             []string{"RCH123", "ABC123", "LLS442"},
			Consumer:              unifieddatalibrary.String("APRON1-230401001"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			IDArrSortie:           unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			IDDepSortie:           unifieddatalibrary.String("1e6edeec-72e9-aaec-d33c-51147cb5ffdd"),
			MissionID:             unifieddatalibrary.String("AJM123456123"),
			OccAircraftMds:        unifieddatalibrary.String("C017A"),
			OccStartTime:          unifieddatalibrary.Time(time.Now()),
			OccTailNumber:         unifieddatalibrary.String("N702JG"),
			Occupied:              unifieddatalibrary.Bool(true),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ReqComment:            unifieddatalibrary.String("Sorry for the late notice."),
			ReqInitials:           unifieddatalibrary.String("CB"),
			ReqOrg:                unifieddatalibrary.String("TACC"),
			ResAircraftMds:        unifieddatalibrary.String("C017A"),
			ResMissionID:          unifieddatalibrary.String("AJM123456123"),
			ResReason:             unifieddatalibrary.String("Maintenance needed"),
			ResTailNumber:         unifieddatalibrary.String("N702JG"),
			ResType:               unifieddatalibrary.String("M"),
			Status:                unifieddatalibrary.AirfieldslotconsumptionUpdateParamsStatusApproved,
			TargetTime:            unifieddatalibrary.Time(time.Now()),
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

func TestAirfieldslotconsumptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfieldslotconsumptions.List(context.TODO(), unifieddatalibrary.AirfieldslotconsumptionListParams{
		StartTime:   time.Now(),
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

func TestAirfieldslotconsumptionDelete(t *testing.T) {
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
	err := client.Airfieldslotconsumptions.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldslotconsumptionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfieldslotconsumptions.Count(context.TODO(), unifieddatalibrary.AirfieldslotconsumptionCountParams{
		StartTime:   time.Now(),
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

func TestAirfieldslotconsumptionQueryhelp(t *testing.T) {
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
	err := client.Airfieldslotconsumptions.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldslotconsumptionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Airfieldslotconsumptions.Tuple(context.TODO(), unifieddatalibrary.AirfieldslotconsumptionTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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
