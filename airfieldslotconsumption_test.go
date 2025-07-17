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

func TestAirfieldSlotConsumptionNewWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldSlotConsumptions.New(context.TODO(), unifieddatalibrary.AirfieldSlotConsumptionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirfieldSlotConsumptionNewParamsDataModeTest,
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
		Status:                unifieddatalibrary.AirfieldSlotConsumptionNewParamsStatusApproved,
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

func TestAirfieldSlotConsumptionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlotConsumptions.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldSlotConsumptionGetParams{
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

func TestAirfieldSlotConsumptionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirfieldSlotConsumptions.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirfieldSlotConsumptionUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirfieldSlotConsumptionUpdateParamsDataModeTest,
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
			Status:                unifieddatalibrary.AirfieldSlotConsumptionUpdateParamsStatusApproved,
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

func TestAirfieldSlotConsumptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlotConsumptions.List(context.TODO(), unifieddatalibrary.AirfieldSlotConsumptionListParams{
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

func TestAirfieldSlotConsumptionDelete(t *testing.T) {
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
	err := client.AirfieldSlotConsumptions.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldSlotConsumptionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlotConsumptions.Count(context.TODO(), unifieddatalibrary.AirfieldSlotConsumptionCountParams{
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

func TestAirfieldSlotConsumptionQueryhelp(t *testing.T) {
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
	_, err := client.AirfieldSlotConsumptions.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirfieldSlotConsumptionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirfieldSlotConsumptions.Tuple(context.TODO(), unifieddatalibrary.AirfieldSlotConsumptionTupleParams{
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
