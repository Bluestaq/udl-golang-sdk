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

func TestAircraftStatusNewWithOptionalParams(t *testing.T) {
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
	err := client.AircraftStatuses.New(context.TODO(), unifieddatalibrary.AircraftStatusNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AircraftStatusNewParamsDataModeTest,
		IDAircraft:            "29232269-e4c2-45c9-aa21-039a33209340",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		AdditionalSys:         []string{"ATOMS", "TUDL", "BLOS1"},
		AirToAirStatus:        unifieddatalibrary.AircraftStatusNewParamsAirToAirStatusOperational,
		AirToGroundStatus:     unifieddatalibrary.AircraftStatusNewParamsAirToGroundStatusOperational,
		AlphaStatusCode:       unifieddatalibrary.String("A2"),
		AltAircraftID:         unifieddatalibrary.String("ORIG-AIRCRAFT-ID"),
		ContaminationStatus:   unifieddatalibrary.String("CLEAR"),
		CurrentIcao:           unifieddatalibrary.String("KCHS"),
		CurrentState:          unifieddatalibrary.String("AVAILABLE"),
		EarliestTaEndTime:     unifieddatalibrary.Time(time.Now()),
		Etic:                  unifieddatalibrary.Time(time.Now()),
		FlightPhase:           unifieddatalibrary.String("Landing"),
		Fuel:                  unifieddatalibrary.Int(10),
		FuelFunction:          unifieddatalibrary.String("Burn"),
		FuelStatus:            unifieddatalibrary.String("DELIVERED"),
		GeoLoc:                unifieddatalibrary.String("AJJY"),
		GroundStatus:          unifieddatalibrary.String("ALERT"),
		GunCapable:            unifieddatalibrary.Bool(true),
		GunRdsMax:             unifieddatalibrary.Int(550),
		GunRdsMin:             unifieddatalibrary.Int(150),
		GunRdsType:            unifieddatalibrary.String("7.62 MM"),
		IDAirfield:            unifieddatalibrary.String("b89430e3-97d9-408c-9c89-fd3840c4b84d"),
		IDPoi:                 unifieddatalibrary.String("0e52f081-a2e3-4b73-b822-88b882232691"),
		Inventory:             []string{"AIM-9 SIDEWINDER", "AIM-120 AMRAAM"},
		InventoryMax:          []int64{2, 2},
		InventoryMin:          []int64{1, 2},
		LastInspectionDate:    unifieddatalibrary.Time(time.Now()),
		LastUpdatedBy:         unifieddatalibrary.String("some.user"),
		MaintPoc:              unifieddatalibrary.String("PSUP NIGHT SHIFT 800-555-4412"),
		MaintPriority:         unifieddatalibrary.String("1"),
		MaintStatus:           unifieddatalibrary.String("maintenance status"),
		MaintStatusDriver:     unifieddatalibrary.String("SCREW STUCK IN LEFT NLG TIRE"),
		MaintStatusUpdate:     unifieddatalibrary.Time(time.Now()),
		MissionReadiness:      unifieddatalibrary.String("ABLE"),
		MxRemark:              unifieddatalibrary.String("COM2 INOP"),
		NextIcao:              unifieddatalibrary.String("PHNL"),
		Notes:                 unifieddatalibrary.String("Some notes for aircraft A"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ParkLocation:          unifieddatalibrary.String("B1"),
		ParkLocationSystem:    unifieddatalibrary.String("GDSS"),
		PreviousIcao:          unifieddatalibrary.String("EGLL"),
		TaStartTime:           unifieddatalibrary.Time(time.Now()),
		TroubleshootEtic:      unifieddatalibrary.Time(time.Now()),
		UnavailableSys:        []string{"CMDS", "AOC"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAircraftStatusGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AircraftStatuses.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AircraftStatusGetParams{
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

func TestAircraftStatusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AircraftStatuses.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AircraftStatusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AircraftStatusUpdateParamsDataModeTest,
			IDAircraft:            "29232269-e4c2-45c9-aa21-039a33209340",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AdditionalSys:         []string{"ATOMS", "TUDL", "BLOS1"},
			AirToAirStatus:        unifieddatalibrary.AircraftStatusUpdateParamsAirToAirStatusOperational,
			AirToGroundStatus:     unifieddatalibrary.AircraftStatusUpdateParamsAirToGroundStatusOperational,
			AlphaStatusCode:       unifieddatalibrary.String("A2"),
			AltAircraftID:         unifieddatalibrary.String("ORIG-AIRCRAFT-ID"),
			ContaminationStatus:   unifieddatalibrary.String("CLEAR"),
			CurrentIcao:           unifieddatalibrary.String("KCHS"),
			CurrentState:          unifieddatalibrary.String("AVAILABLE"),
			EarliestTaEndTime:     unifieddatalibrary.Time(time.Now()),
			Etic:                  unifieddatalibrary.Time(time.Now()),
			FlightPhase:           unifieddatalibrary.String("Landing"),
			Fuel:                  unifieddatalibrary.Int(10),
			FuelFunction:          unifieddatalibrary.String("Burn"),
			FuelStatus:            unifieddatalibrary.String("DELIVERED"),
			GeoLoc:                unifieddatalibrary.String("AJJY"),
			GroundStatus:          unifieddatalibrary.String("ALERT"),
			GunCapable:            unifieddatalibrary.Bool(true),
			GunRdsMax:             unifieddatalibrary.Int(550),
			GunRdsMin:             unifieddatalibrary.Int(150),
			GunRdsType:            unifieddatalibrary.String("7.62 MM"),
			IDAirfield:            unifieddatalibrary.String("b89430e3-97d9-408c-9c89-fd3840c4b84d"),
			IDPoi:                 unifieddatalibrary.String("0e52f081-a2e3-4b73-b822-88b882232691"),
			Inventory:             []string{"AIM-9 SIDEWINDER", "AIM-120 AMRAAM"},
			InventoryMax:          []int64{2, 2},
			InventoryMin:          []int64{1, 2},
			LastInspectionDate:    unifieddatalibrary.Time(time.Now()),
			LastUpdatedBy:         unifieddatalibrary.String("some.user"),
			MaintPoc:              unifieddatalibrary.String("PSUP NIGHT SHIFT 800-555-4412"),
			MaintPriority:         unifieddatalibrary.String("1"),
			MaintStatus:           unifieddatalibrary.String("maintenance status"),
			MaintStatusDriver:     unifieddatalibrary.String("SCREW STUCK IN LEFT NLG TIRE"),
			MaintStatusUpdate:     unifieddatalibrary.Time(time.Now()),
			MissionReadiness:      unifieddatalibrary.String("ABLE"),
			MxRemark:              unifieddatalibrary.String("COM2 INOP"),
			NextIcao:              unifieddatalibrary.String("PHNL"),
			Notes:                 unifieddatalibrary.String("Some notes for aircraft A"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ParkLocation:          unifieddatalibrary.String("B1"),
			ParkLocationSystem:    unifieddatalibrary.String("GDSS"),
			PreviousIcao:          unifieddatalibrary.String("EGLL"),
			TaStartTime:           unifieddatalibrary.Time(time.Now()),
			TroubleshootEtic:      unifieddatalibrary.Time(time.Now()),
			UnavailableSys:        []string{"CMDS", "AOC"},
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

func TestAircraftStatusListWithOptionalParams(t *testing.T) {
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
	_, err := client.AircraftStatuses.List(context.TODO(), unifieddatalibrary.AircraftStatusListParams{
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

func TestAircraftStatusDelete(t *testing.T) {
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
	err := client.AircraftStatuses.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAircraftStatusCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AircraftStatuses.Count(context.TODO(), unifieddatalibrary.AircraftStatusCountParams{
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

func TestAircraftStatusQueryhelp(t *testing.T) {
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
	_, err := client.AircraftStatuses.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAircraftStatusTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AircraftStatuses.Tuple(context.TODO(), unifieddatalibrary.AircraftStatusTupleParams{
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
