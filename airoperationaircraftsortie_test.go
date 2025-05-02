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

func TestAirOperationAircraftSortieNewWithOptionalParams(t *testing.T) {
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
	err := client.AirOperations.AircraftSortie.New(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirOperationAircraftSortieNewParamsDataModeTest,
		PlannedDepTime:        time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("AIRCRAFTSORTIE-ID"),
		ActualArrTime:         unifieddatalibrary.Time(time.Now()),
		ActualBlockInTime:     unifieddatalibrary.Time(time.Now()),
		ActualBlockOutTime:    unifieddatalibrary.Time(time.Now()),
		ActualDepTime:         unifieddatalibrary.Time(time.Now()),
		AircraftAdsb:          unifieddatalibrary.String("AE123C"),
		AircraftAltID:         unifieddatalibrary.String("ALT-AIRCRAFT-ID"),
		AircraftEvent:         unifieddatalibrary.String("Example event"),
		AircraftMds:           unifieddatalibrary.String("C017A"),
		AircraftRemarks:       unifieddatalibrary.String("Some remark about aircraft A"),
		AlertStatus:           unifieddatalibrary.Int(22),
		AlertStatusCode:       unifieddatalibrary.String("C1"),
		AmcMsnNum:             unifieddatalibrary.String("AJM512571333"),
		AmcMsnType:            unifieddatalibrary.String("SAAM"),
		ArrFaa:                unifieddatalibrary.String("FAA1"),
		ArrIata:               unifieddatalibrary.String("AAA"),
		ArrIcao:               unifieddatalibrary.String("KCOS"),
		ArrItinerary:          unifieddatalibrary.Int(101),
		ArrPurposeCode:        unifieddatalibrary.String("O"),
		CallSign:              unifieddatalibrary.String("BAKER"),
		CargoConfig:           unifieddatalibrary.String("C-1"),
		CommanderName:         unifieddatalibrary.String("Smith"),
		CurrentState:          unifieddatalibrary.String("Park"),
		DelayCode:             unifieddatalibrary.String("500"),
		DepFaa:                unifieddatalibrary.String("FAA1"),
		DepIata:               unifieddatalibrary.String("AAA"),
		DepIcao:               unifieddatalibrary.String("KCOS"),
		DepItinerary:          unifieddatalibrary.Int(100),
		DepPurposeCode:        unifieddatalibrary.String("P"),
		Dhd:                   unifieddatalibrary.Time(time.Now()),
		DhdReason:             unifieddatalibrary.String("Due for maintenance"),
		EstArrTime:            unifieddatalibrary.Time(time.Now()),
		EstBlockInTime:        unifieddatalibrary.Time(time.Now()),
		EstBlockOutTime:       unifieddatalibrary.Time(time.Now()),
		EstDepTime:            unifieddatalibrary.Time(time.Now()),
		FlightTime:            unifieddatalibrary.Float(104.5),
		FmDeskNum:             unifieddatalibrary.String("7198675309"),
		FmName:                unifieddatalibrary.String("Smith"),
		FuelReq:               unifieddatalibrary.Float(20000.1),
		GndTime:               unifieddatalibrary.Float(387.8),
		IDAircraft:            unifieddatalibrary.String("REF-AIRCRAFT-ID"),
		IDMission:             unifieddatalibrary.String("fa18d96e-91ea-60da-a7a8-1af6500066c8"),
		JcsPriority:           unifieddatalibrary.String("1A3"),
		LegNum:                unifieddatalibrary.Int(14),
		LineNumber:            unifieddatalibrary.Int(99),
		MissionID:             unifieddatalibrary.String("ABLE"),
		MissionUpdate:         unifieddatalibrary.Time(time.Now()),
		ObjectiveRemarks:      unifieddatalibrary.String("Some objective remark about aircraft A"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigSortieID:          unifieddatalibrary.String("A0640"),
		OxyOnCrew:             unifieddatalibrary.Float(12.3),
		OxyOnPax:              unifieddatalibrary.Float(12.3),
		OxyReqCrew:            unifieddatalibrary.Float(12.3),
		OxyReqPax:             unifieddatalibrary.Float(12.3),
		ParkingLoc:            unifieddatalibrary.String("KCOS"),
		Passengers:            unifieddatalibrary.Int(17),
		PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
		PprStatus:             unifieddatalibrary.AirOperationAircraftSortieNewParamsPprStatusPending,
		PrimaryScl:            unifieddatalibrary.String("ABC"),
		ReqConfig:             unifieddatalibrary.String("C-1"),
		ResultRemarks:         unifieddatalibrary.String("Some remark about aircraft A"),
		RvnReq:                unifieddatalibrary.AirOperationAircraftSortieNewParamsRvnReqR,
		ScheduleRemarks:       unifieddatalibrary.String("Some schedule remark about aircraft A"),
		SecondaryScl:          unifieddatalibrary.String("ABC"),
		Soe:                   unifieddatalibrary.String("OPS"),
		SortieDate:            unifieddatalibrary.Time(time.Now()),
		TailNumber:            unifieddatalibrary.String("Tail_1"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAircraftSortieListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AircraftSortie.List(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieListParams{
		PlannedDepTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAircraftSortieCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AircraftSortie.Count(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieCountParams{
		PlannedDepTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAircraftSortieNewBulk(t *testing.T) {
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
	err := client.AirOperations.AircraftSortie.NewBulk(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieNewBulkParams{
		Body: []unifieddatalibrary.AirOperationAircraftSortieNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			PlannedDepTime:        time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("AIRCRAFTSORTIE-ID"),
			ActualArrTime:         unifieddatalibrary.Time(time.Now()),
			ActualBlockInTime:     unifieddatalibrary.Time(time.Now()),
			ActualBlockOutTime:    unifieddatalibrary.Time(time.Now()),
			ActualDepTime:         unifieddatalibrary.Time(time.Now()),
			AircraftAdsb:          unifieddatalibrary.String("AE123C"),
			AircraftAltID:         unifieddatalibrary.String("ALT-AIRCRAFT-ID"),
			AircraftEvent:         unifieddatalibrary.String("Example event"),
			AircraftMds:           unifieddatalibrary.String("C017A"),
			AircraftRemarks:       unifieddatalibrary.String("Some remark about aircraft A"),
			AlertStatus:           unifieddatalibrary.Int(22),
			AlertStatusCode:       unifieddatalibrary.String("C1"),
			AmcMsnNum:             unifieddatalibrary.String("AJM512571333"),
			AmcMsnType:            unifieddatalibrary.String("SAAM"),
			ArrFaa:                unifieddatalibrary.String("FAA1"),
			ArrIata:               unifieddatalibrary.String("AAA"),
			ArrIcao:               unifieddatalibrary.String("KCOS"),
			ArrItinerary:          unifieddatalibrary.Int(101),
			ArrPurposeCode:        unifieddatalibrary.String("O"),
			CallSign:              unifieddatalibrary.String("BAKER"),
			CargoConfig:           unifieddatalibrary.String("C-1"),
			CommanderName:         unifieddatalibrary.String("Smith"),
			CurrentState:          unifieddatalibrary.String("Park"),
			DelayCode:             unifieddatalibrary.String("500"),
			DepFaa:                unifieddatalibrary.String("FAA1"),
			DepIata:               unifieddatalibrary.String("AAA"),
			DepIcao:               unifieddatalibrary.String("KCOS"),
			DepItinerary:          unifieddatalibrary.Int(100),
			DepPurposeCode:        unifieddatalibrary.String("P"),
			Dhd:                   unifieddatalibrary.Time(time.Now()),
			DhdReason:             unifieddatalibrary.String("Due for maintenance"),
			EstArrTime:            unifieddatalibrary.Time(time.Now()),
			EstBlockInTime:        unifieddatalibrary.Time(time.Now()),
			EstBlockOutTime:       unifieddatalibrary.Time(time.Now()),
			EstDepTime:            unifieddatalibrary.Time(time.Now()),
			FlightTime:            unifieddatalibrary.Float(104.5),
			FmDeskNum:             unifieddatalibrary.String("7198675309"),
			FmName:                unifieddatalibrary.String("Smith"),
			FuelReq:               unifieddatalibrary.Float(20000.1),
			GndTime:               unifieddatalibrary.Float(387.8),
			IDAircraft:            unifieddatalibrary.String("REF-AIRCRAFT-ID"),
			IDMission:             unifieddatalibrary.String("fa18d96e-91ea-60da-a7a8-1af6500066c8"),
			JcsPriority:           unifieddatalibrary.String("1A3"),
			LegNum:                unifieddatalibrary.Int(14),
			LineNumber:            unifieddatalibrary.Int(99),
			MissionID:             unifieddatalibrary.String("ABLE"),
			MissionUpdate:         unifieddatalibrary.Time(time.Now()),
			ObjectiveRemarks:      unifieddatalibrary.String("Some objective remark about aircraft A"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSortieID:          unifieddatalibrary.String("A0640"),
			OxyOnCrew:             unifieddatalibrary.Float(12.3),
			OxyOnPax:              unifieddatalibrary.Float(12.3),
			OxyReqCrew:            unifieddatalibrary.Float(12.3),
			OxyReqPax:             unifieddatalibrary.Float(12.3),
			ParkingLoc:            unifieddatalibrary.String("KCOS"),
			Passengers:            unifieddatalibrary.Int(17),
			PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
			PprStatus:             "PENDING",
			PrimaryScl:            unifieddatalibrary.String("ABC"),
			ReqConfig:             unifieddatalibrary.String("C-1"),
			ResultRemarks:         unifieddatalibrary.String("Some remark about aircraft A"),
			RvnReq:                "R",
			ScheduleRemarks:       unifieddatalibrary.String("Some schedule remark about aircraft A"),
			SecondaryScl:          unifieddatalibrary.String("ABC"),
			Soe:                   unifieddatalibrary.String("OPS"),
			SortieDate:            unifieddatalibrary.Time(time.Now()),
			TailNumber:            unifieddatalibrary.String("Tail_1"),
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

func TestAirOperationAircraftSortieHistoryAodrWithOptionalParams(t *testing.T) {
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
	err := client.AirOperations.AircraftSortie.HistoryAodr(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieHistoryAodrParams{
		PlannedDepTime:  time.Now(),
		Columns:         unifieddatalibrary.String("columns"),
		FirstResult:     unifieddatalibrary.Int(0),
		MaxResults:      unifieddatalibrary.Int(0),
		Notification:    unifieddatalibrary.String("notification"),
		OutputDelimiter: unifieddatalibrary.String("outputDelimiter"),
		OutputFormat:    unifieddatalibrary.String("outputFormat"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAircraftSortieHistoryCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AircraftSortie.HistoryCount(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieHistoryCountParams{
		PlannedDepTime: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirOperationAircraftSortieHistoryQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.AirOperations.AircraftSortie.HistoryQuery(context.TODO(), unifieddatalibrary.AirOperationAircraftSortieHistoryQueryParams{
		PlannedDepTime: time.Now(),
		Columns:        unifieddatalibrary.String("columns"),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
