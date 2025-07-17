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

func TestAirEventNewWithOptionalParams(t *testing.T) {
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
	err := client.AirEvents.New(context.TODO(), unifieddatalibrary.AirEventNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirEventNewParamsDataModeTest,
		Source:                "Bluestaq",
		Type:                  "FUEL TRANSFER",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		ActualArrTime:         unifieddatalibrary.Time(time.Now()),
		ActualDepTime:         unifieddatalibrary.Time(time.Now()),
		Arct:                  unifieddatalibrary.Time(time.Now()),
		ArEventType:           unifieddatalibrary.String("V"),
		ArrPurpose:            unifieddatalibrary.String("A"),
		ArTrackID:             unifieddatalibrary.String("CH61"),
		ArTrackName:           unifieddatalibrary.String("CH61 POST"),
		BaseAlt:               unifieddatalibrary.Float(28000.1),
		Cancelled:             unifieddatalibrary.Bool(false),
		DepPurpose:            unifieddatalibrary.String("Q"),
		EstArrTime:            unifieddatalibrary.Time(time.Now()),
		EstDepTime:            unifieddatalibrary.Time(time.Now()),
		ExternalAirEventID:    unifieddatalibrary.String("MB014313032022407540"),
		ExternalArTrackID:     unifieddatalibrary.String("6418a4b68e5c3896bf024cc79aa4174c"),
		IDMission:             unifieddatalibrary.String("190dea6d-2a90-45a2-a276-be9047d9b96c"),
		IDSortie:              unifieddatalibrary.String("b9866c03-2397-4506-8153-852e72d9b54f"),
		LegNum:                unifieddatalibrary.Int(825),
		Location:              unifieddatalibrary.String("901EW"),
		NumTankers:            unifieddatalibrary.Int(1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
		PlannedDepTime:        unifieddatalibrary.Time(time.Now()),
		Priority:              unifieddatalibrary.String("1A2"),
		Receivers: []unifieddatalibrary.AirEventNewParamsReceiver{{
			AltReceiverMissionID:   unifieddatalibrary.String("1UN05201L121"),
			AmcReceiverMissionID:   unifieddatalibrary.String("8PH000B1S052"),
			ExternalReceiverID:     unifieddatalibrary.String("3fb8169f-adc1-4667-acab-8415a012d766"),
			FuelOn:                 unifieddatalibrary.Float(15000000.1),
			IDReceiverAirfield:     unifieddatalibrary.String("96c4c2ba-a031-4e58-9b8e-3c6fb90a7534"),
			IDReceiverMission:      unifieddatalibrary.String("ce99757d-f733-461f-8939-3939d4f05946"),
			IDReceiverSortie:       unifieddatalibrary.String("1d03e85a-1fb9-4f6e-86a0-593306b6e3f0"),
			NumRecAircraft:         unifieddatalibrary.Int(3),
			PackageID:              unifieddatalibrary.String("135"),
			ReceiverCallSign:       unifieddatalibrary.String("BAKER"),
			ReceiverCellPosition:   unifieddatalibrary.Int(2),
			ReceiverCoord:          unifieddatalibrary.String("TTC601"),
			ReceiverDeliveryMethod: unifieddatalibrary.String("DROGUE"),
			ReceiverDeployedIcao:   unifieddatalibrary.String("KOFF"),
			ReceiverExercise:       unifieddatalibrary.String("NATO19"),
			ReceiverFuelType:       unifieddatalibrary.String("JP8"),
			ReceiverLegNum:         unifieddatalibrary.Int(825),
			ReceiverMds:            unifieddatalibrary.String("KC135R"),
			ReceiverOwner:          unifieddatalibrary.String("117ARW"),
			ReceiverPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
			RecOrg:                 unifieddatalibrary.String("AMC"),
			SequenceNum:            unifieddatalibrary.String("1018"),
		}},
		Remarks: []unifieddatalibrary.AirEventNewParamsRemark{{
			Date:             unifieddatalibrary.Time(time.Now()),
			ExternalRemarkID: unifieddatalibrary.String("23ea2877a6f74d7d8f309567a5896441"),
			Text:             unifieddatalibrary.String("Example air event remarks."),
			User:             unifieddatalibrary.String("John Doe"),
		}},
		RevTrack:   unifieddatalibrary.Bool(true),
		Rzct:       unifieddatalibrary.Time(time.Now()),
		RzPoint:    unifieddatalibrary.String("AN"),
		RzType:     unifieddatalibrary.String("PP"),
		ShortTrack: unifieddatalibrary.Bool(true),
		StatusCode: unifieddatalibrary.String("R"),
		Tankers: []unifieddatalibrary.AirEventNewParamsTanker{{
			AltTankerMissionID:   unifieddatalibrary.String("1UN05201L121"),
			AmcTankerMissionID:   unifieddatalibrary.String("8PH000B1S052"),
			DualRole:             unifieddatalibrary.Bool(true),
			ExternalTankerID:     unifieddatalibrary.String("ca673c580fb949a5b733f0e0b67ffab2"),
			FuelOff:              unifieddatalibrary.Float(15000000.1),
			IDTankerAirfield:     unifieddatalibrary.String("b33955d2-67d3-42be-8316-263e284ce6cc"),
			IDTankerMission:      unifieddatalibrary.String("edef700c-9917-4dbf-a153-89ffd4446fe9"),
			IDTankerSortie:       unifieddatalibrary.String("d833a4bc-756b-41d5-8845-f146fe563387"),
			TankerCallSign:       unifieddatalibrary.String("BAKER"),
			TankerCellPosition:   unifieddatalibrary.Int(2),
			TankerCoord:          unifieddatalibrary.String("TTC601"),
			TankerDeliveryMethod: unifieddatalibrary.String("DROGUE"),
			TankerDeployedIcao:   unifieddatalibrary.String("KOFF"),
			TankerFuelType:       unifieddatalibrary.String("JP8"),
			TankerLegNum:         unifieddatalibrary.Int(825),
			TankerMds:            unifieddatalibrary.String("KC135R"),
			TankerOwner:          unifieddatalibrary.String("117ARW"),
			TankerPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
		}},
		TrackTime: unifieddatalibrary.Float(1.5),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirEventUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirEvents.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirEventUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirEventUpdateParamsDataModeTest,
			Source:                "Bluestaq",
			Type:                  "FUEL TRANSFER",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActualArrTime:         unifieddatalibrary.Time(time.Now()),
			ActualDepTime:         unifieddatalibrary.Time(time.Now()),
			Arct:                  unifieddatalibrary.Time(time.Now()),
			ArEventType:           unifieddatalibrary.String("V"),
			ArrPurpose:            unifieddatalibrary.String("A"),
			ArTrackID:             unifieddatalibrary.String("CH61"),
			ArTrackName:           unifieddatalibrary.String("CH61 POST"),
			BaseAlt:               unifieddatalibrary.Float(28000.1),
			Cancelled:             unifieddatalibrary.Bool(false),
			DepPurpose:            unifieddatalibrary.String("Q"),
			EstArrTime:            unifieddatalibrary.Time(time.Now()),
			EstDepTime:            unifieddatalibrary.Time(time.Now()),
			ExternalAirEventID:    unifieddatalibrary.String("MB014313032022407540"),
			ExternalArTrackID:     unifieddatalibrary.String("6418a4b68e5c3896bf024cc79aa4174c"),
			IDMission:             unifieddatalibrary.String("190dea6d-2a90-45a2-a276-be9047d9b96c"),
			IDSortie:              unifieddatalibrary.String("b9866c03-2397-4506-8153-852e72d9b54f"),
			LegNum:                unifieddatalibrary.Int(825),
			Location:              unifieddatalibrary.String("901EW"),
			NumTankers:            unifieddatalibrary.Int(1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
			PlannedDepTime:        unifieddatalibrary.Time(time.Now()),
			Priority:              unifieddatalibrary.String("1A2"),
			Receivers: []unifieddatalibrary.AirEventUpdateParamsReceiver{{
				AltReceiverMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcReceiverMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				ExternalReceiverID:     unifieddatalibrary.String("3fb8169f-adc1-4667-acab-8415a012d766"),
				FuelOn:                 unifieddatalibrary.Float(15000000.1),
				IDReceiverAirfield:     unifieddatalibrary.String("96c4c2ba-a031-4e58-9b8e-3c6fb90a7534"),
				IDReceiverMission:      unifieddatalibrary.String("ce99757d-f733-461f-8939-3939d4f05946"),
				IDReceiverSortie:       unifieddatalibrary.String("1d03e85a-1fb9-4f6e-86a0-593306b6e3f0"),
				NumRecAircraft:         unifieddatalibrary.Int(3),
				PackageID:              unifieddatalibrary.String("135"),
				ReceiverCallSign:       unifieddatalibrary.String("BAKER"),
				ReceiverCellPosition:   unifieddatalibrary.Int(2),
				ReceiverCoord:          unifieddatalibrary.String("TTC601"),
				ReceiverDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				ReceiverDeployedIcao:   unifieddatalibrary.String("KOFF"),
				ReceiverExercise:       unifieddatalibrary.String("NATO19"),
				ReceiverFuelType:       unifieddatalibrary.String("JP8"),
				ReceiverLegNum:         unifieddatalibrary.Int(825),
				ReceiverMds:            unifieddatalibrary.String("KC135R"),
				ReceiverOwner:          unifieddatalibrary.String("117ARW"),
				ReceiverPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
				RecOrg:                 unifieddatalibrary.String("AMC"),
				SequenceNum:            unifieddatalibrary.String("1018"),
			}},
			Remarks: []unifieddatalibrary.AirEventUpdateParamsRemark{{
				Date:             unifieddatalibrary.Time(time.Now()),
				ExternalRemarkID: unifieddatalibrary.String("23ea2877a6f74d7d8f309567a5896441"),
				Text:             unifieddatalibrary.String("Example air event remarks."),
				User:             unifieddatalibrary.String("John Doe"),
			}},
			RevTrack:   unifieddatalibrary.Bool(true),
			Rzct:       unifieddatalibrary.Time(time.Now()),
			RzPoint:    unifieddatalibrary.String("AN"),
			RzType:     unifieddatalibrary.String("PP"),
			ShortTrack: unifieddatalibrary.Bool(true),
			StatusCode: unifieddatalibrary.String("R"),
			Tankers: []unifieddatalibrary.AirEventUpdateParamsTanker{{
				AltTankerMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcTankerMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				DualRole:             unifieddatalibrary.Bool(true),
				ExternalTankerID:     unifieddatalibrary.String("ca673c580fb949a5b733f0e0b67ffab2"),
				FuelOff:              unifieddatalibrary.Float(15000000.1),
				IDTankerAirfield:     unifieddatalibrary.String("b33955d2-67d3-42be-8316-263e284ce6cc"),
				IDTankerMission:      unifieddatalibrary.String("edef700c-9917-4dbf-a153-89ffd4446fe9"),
				IDTankerSortie:       unifieddatalibrary.String("d833a4bc-756b-41d5-8845-f146fe563387"),
				TankerCallSign:       unifieddatalibrary.String("BAKER"),
				TankerCellPosition:   unifieddatalibrary.Int(2),
				TankerCoord:          unifieddatalibrary.String("TTC601"),
				TankerDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				TankerDeployedIcao:   unifieddatalibrary.String("KOFF"),
				TankerFuelType:       unifieddatalibrary.String("JP8"),
				TankerLegNum:         unifieddatalibrary.Int(825),
				TankerMds:            unifieddatalibrary.String("KC135R"),
				TankerOwner:          unifieddatalibrary.String("117ARW"),
				TankerPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
			}},
			TrackTime: unifieddatalibrary.Float(1.5),
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

func TestAirEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirEvents.List(context.TODO(), unifieddatalibrary.AirEventListParams{
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

func TestAirEventDelete(t *testing.T) {
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
	err := client.AirEvents.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirEventCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirEvents.Count(context.TODO(), unifieddatalibrary.AirEventCountParams{
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

func TestAirEventNewBulk(t *testing.T) {
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
	err := client.AirEvents.NewBulk(context.TODO(), unifieddatalibrary.AirEventNewBulkParams{
		Body: []unifieddatalibrary.AirEventNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Type:                  "FUEL TRANSFER",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActualArrTime:         unifieddatalibrary.Time(time.Now()),
			ActualDepTime:         unifieddatalibrary.Time(time.Now()),
			Arct:                  unifieddatalibrary.Time(time.Now()),
			ArEventType:           unifieddatalibrary.String("V"),
			ArrPurpose:            unifieddatalibrary.String("A"),
			ArTrackID:             unifieddatalibrary.String("CH61"),
			ArTrackName:           unifieddatalibrary.String("CH61 POST"),
			BaseAlt:               unifieddatalibrary.Float(28000.1),
			Cancelled:             unifieddatalibrary.Bool(false),
			DepPurpose:            unifieddatalibrary.String("Q"),
			EstArrTime:            unifieddatalibrary.Time(time.Now()),
			EstDepTime:            unifieddatalibrary.Time(time.Now()),
			ExternalAirEventID:    unifieddatalibrary.String("MB014313032022407540"),
			ExternalArTrackID:     unifieddatalibrary.String("6418a4b68e5c3896bf024cc79aa4174c"),
			IDMission:             unifieddatalibrary.String("190dea6d-2a90-45a2-a276-be9047d9b96c"),
			IDSortie:              unifieddatalibrary.String("b9866c03-2397-4506-8153-852e72d9b54f"),
			LegNum:                unifieddatalibrary.Int(825),
			Location:              unifieddatalibrary.String("901EW"),
			NumTankers:            unifieddatalibrary.Int(1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
			PlannedDepTime:        unifieddatalibrary.Time(time.Now()),
			Priority:              unifieddatalibrary.String("1A2"),
			Receivers: []unifieddatalibrary.AirEventNewBulkParamsBodyReceiver{{
				AltReceiverMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcReceiverMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				ExternalReceiverID:     unifieddatalibrary.String("3fb8169f-adc1-4667-acab-8415a012d766"),
				FuelOn:                 unifieddatalibrary.Float(15000000.1),
				IDReceiverAirfield:     unifieddatalibrary.String("96c4c2ba-a031-4e58-9b8e-3c6fb90a7534"),
				IDReceiverMission:      unifieddatalibrary.String("ce99757d-f733-461f-8939-3939d4f05946"),
				IDReceiverSortie:       unifieddatalibrary.String("1d03e85a-1fb9-4f6e-86a0-593306b6e3f0"),
				NumRecAircraft:         unifieddatalibrary.Int(3),
				PackageID:              unifieddatalibrary.String("135"),
				ReceiverCallSign:       unifieddatalibrary.String("BAKER"),
				ReceiverCellPosition:   unifieddatalibrary.Int(2),
				ReceiverCoord:          unifieddatalibrary.String("TTC601"),
				ReceiverDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				ReceiverDeployedIcao:   unifieddatalibrary.String("KOFF"),
				ReceiverExercise:       unifieddatalibrary.String("NATO19"),
				ReceiverFuelType:       unifieddatalibrary.String("JP8"),
				ReceiverLegNum:         unifieddatalibrary.Int(825),
				ReceiverMds:            unifieddatalibrary.String("KC135R"),
				ReceiverOwner:          unifieddatalibrary.String("117ARW"),
				ReceiverPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
				RecOrg:                 unifieddatalibrary.String("AMC"),
				SequenceNum:            unifieddatalibrary.String("1018"),
			}},
			Remarks: []unifieddatalibrary.AirEventNewBulkParamsBodyRemark{{
				Date:             unifieddatalibrary.Time(time.Now()),
				ExternalRemarkID: unifieddatalibrary.String("23ea2877a6f74d7d8f309567a5896441"),
				Text:             unifieddatalibrary.String("Example air event remarks."),
				User:             unifieddatalibrary.String("John Doe"),
			}},
			RevTrack:   unifieddatalibrary.Bool(true),
			Rzct:       unifieddatalibrary.Time(time.Now()),
			RzPoint:    unifieddatalibrary.String("AN"),
			RzType:     unifieddatalibrary.String("PP"),
			ShortTrack: unifieddatalibrary.Bool(true),
			StatusCode: unifieddatalibrary.String("R"),
			Tankers: []unifieddatalibrary.AirEventNewBulkParamsBodyTanker{{
				AltTankerMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcTankerMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				DualRole:             unifieddatalibrary.Bool(true),
				ExternalTankerID:     unifieddatalibrary.String("ca673c580fb949a5b733f0e0b67ffab2"),
				FuelOff:              unifieddatalibrary.Float(15000000.1),
				IDTankerAirfield:     unifieddatalibrary.String("b33955d2-67d3-42be-8316-263e284ce6cc"),
				IDTankerMission:      unifieddatalibrary.String("edef700c-9917-4dbf-a153-89ffd4446fe9"),
				IDTankerSortie:       unifieddatalibrary.String("d833a4bc-756b-41d5-8845-f146fe563387"),
				TankerCallSign:       unifieddatalibrary.String("BAKER"),
				TankerCellPosition:   unifieddatalibrary.Int(2),
				TankerCoord:          unifieddatalibrary.String("TTC601"),
				TankerDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				TankerDeployedIcao:   unifieddatalibrary.String("KOFF"),
				TankerFuelType:       unifieddatalibrary.String("JP8"),
				TankerLegNum:         unifieddatalibrary.Int(825),
				TankerMds:            unifieddatalibrary.String("KC135R"),
				TankerOwner:          unifieddatalibrary.String("117ARW"),
				TankerPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
			}},
			TrackTime: unifieddatalibrary.Float(1.5),
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

func TestAirEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirEvents.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirEventGetParams{
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

func TestAirEventQueryhelp(t *testing.T) {
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
	_, err := client.AirEvents.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirEventTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirEvents.Tuple(context.TODO(), unifieddatalibrary.AirEventTupleParams{
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

func TestAirEventUnvalidatedPublish(t *testing.T) {
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
	err := client.AirEvents.UnvalidatedPublish(context.TODO(), unifieddatalibrary.AirEventUnvalidatedPublishParams{
		Body: []unifieddatalibrary.AirEventUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Type:                  "FUEL TRANSFER",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActualArrTime:         unifieddatalibrary.Time(time.Now()),
			ActualDepTime:         unifieddatalibrary.Time(time.Now()),
			Arct:                  unifieddatalibrary.Time(time.Now()),
			ArEventType:           unifieddatalibrary.String("V"),
			ArrPurpose:            unifieddatalibrary.String("A"),
			ArTrackID:             unifieddatalibrary.String("CH61"),
			ArTrackName:           unifieddatalibrary.String("CH61 POST"),
			BaseAlt:               unifieddatalibrary.Float(28000.1),
			Cancelled:             unifieddatalibrary.Bool(false),
			DepPurpose:            unifieddatalibrary.String("Q"),
			EstArrTime:            unifieddatalibrary.Time(time.Now()),
			EstDepTime:            unifieddatalibrary.Time(time.Now()),
			ExternalAirEventID:    unifieddatalibrary.String("MB014313032022407540"),
			ExternalArTrackID:     unifieddatalibrary.String("6418a4b68e5c3896bf024cc79aa4174c"),
			IDMission:             unifieddatalibrary.String("190dea6d-2a90-45a2-a276-be9047d9b96c"),
			IDSortie:              unifieddatalibrary.String("b9866c03-2397-4506-8153-852e72d9b54f"),
			LegNum:                unifieddatalibrary.Int(825),
			Location:              unifieddatalibrary.String("901EW"),
			NumTankers:            unifieddatalibrary.Int(1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlannedArrTime:        unifieddatalibrary.Time(time.Now()),
			PlannedDepTime:        unifieddatalibrary.Time(time.Now()),
			Priority:              unifieddatalibrary.String("1A2"),
			Receivers: []unifieddatalibrary.AirEventUnvalidatedPublishParamsBodyReceiver{{
				AltReceiverMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcReceiverMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				ExternalReceiverID:     unifieddatalibrary.String("3fb8169f-adc1-4667-acab-8415a012d766"),
				FuelOn:                 unifieddatalibrary.Float(15000000.1),
				IDReceiverAirfield:     unifieddatalibrary.String("96c4c2ba-a031-4e58-9b8e-3c6fb90a7534"),
				IDReceiverMission:      unifieddatalibrary.String("ce99757d-f733-461f-8939-3939d4f05946"),
				IDReceiverSortie:       unifieddatalibrary.String("1d03e85a-1fb9-4f6e-86a0-593306b6e3f0"),
				NumRecAircraft:         unifieddatalibrary.Int(3),
				PackageID:              unifieddatalibrary.String("135"),
				ReceiverCallSign:       unifieddatalibrary.String("BAKER"),
				ReceiverCellPosition:   unifieddatalibrary.Int(2),
				ReceiverCoord:          unifieddatalibrary.String("TTC601"),
				ReceiverDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				ReceiverDeployedIcao:   unifieddatalibrary.String("KOFF"),
				ReceiverExercise:       unifieddatalibrary.String("NATO19"),
				ReceiverFuelType:       unifieddatalibrary.String("JP8"),
				ReceiverLegNum:         unifieddatalibrary.Int(825),
				ReceiverMds:            unifieddatalibrary.String("KC135R"),
				ReceiverOwner:          unifieddatalibrary.String("117ARW"),
				ReceiverPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
				RecOrg:                 unifieddatalibrary.String("AMC"),
				SequenceNum:            unifieddatalibrary.String("1018"),
			}},
			Remarks: []unifieddatalibrary.AirEventUnvalidatedPublishParamsBodyRemark{{
				Date:             unifieddatalibrary.Time(time.Now()),
				ExternalRemarkID: unifieddatalibrary.String("23ea2877a6f74d7d8f309567a5896441"),
				Text:             unifieddatalibrary.String("Example air event remarks."),
				User:             unifieddatalibrary.String("John Doe"),
			}},
			RevTrack:   unifieddatalibrary.Bool(true),
			Rzct:       unifieddatalibrary.Time(time.Now()),
			RzPoint:    unifieddatalibrary.String("AN"),
			RzType:     unifieddatalibrary.String("PP"),
			ShortTrack: unifieddatalibrary.Bool(true),
			StatusCode: unifieddatalibrary.String("R"),
			Tankers: []unifieddatalibrary.AirEventUnvalidatedPublishParamsBodyTanker{{
				AltTankerMissionID:   unifieddatalibrary.String("1UN05201L121"),
				AmcTankerMissionID:   unifieddatalibrary.String("8PH000B1S052"),
				DualRole:             unifieddatalibrary.Bool(true),
				ExternalTankerID:     unifieddatalibrary.String("ca673c580fb949a5b733f0e0b67ffab2"),
				FuelOff:              unifieddatalibrary.Float(15000000.1),
				IDTankerAirfield:     unifieddatalibrary.String("b33955d2-67d3-42be-8316-263e284ce6cc"),
				IDTankerMission:      unifieddatalibrary.String("edef700c-9917-4dbf-a153-89ffd4446fe9"),
				IDTankerSortie:       unifieddatalibrary.String("d833a4bc-756b-41d5-8845-f146fe563387"),
				TankerCallSign:       unifieddatalibrary.String("BAKER"),
				TankerCellPosition:   unifieddatalibrary.Int(2),
				TankerCoord:          unifieddatalibrary.String("TTC601"),
				TankerDeliveryMethod: unifieddatalibrary.String("DROGUE"),
				TankerDeployedIcao:   unifieddatalibrary.String("KOFF"),
				TankerFuelType:       unifieddatalibrary.String("JP8"),
				TankerLegNum:         unifieddatalibrary.Int(825),
				TankerMds:            unifieddatalibrary.String("KC135R"),
				TankerOwner:          unifieddatalibrary.String("117ARW"),
				TankerPoc:            unifieddatalibrary.String("JOHN SMITH (555)555-5555"),
			}},
			TrackTime: unifieddatalibrary.Float(1.5),
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
