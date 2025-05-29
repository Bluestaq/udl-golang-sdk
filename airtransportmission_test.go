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

func TestAirTransportMissionNewWithOptionalParams(t *testing.T) {
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
	err := client.AirTransportMissions.New(context.TODO(), unifieddatalibrary.AirTransportMissionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.AirTransportMissionNewParamsDataModeTest,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("bdad6945-c9e4-b829-f7be-1ad075541921"),
		Abp:                   unifieddatalibrary.String("ZZ12"),
		Alias:                 unifieddatalibrary.String("PACIFIC DEPLOY / CHAP 3 MOVEMENT"),
		AllocatedUnit:         unifieddatalibrary.String("437 AEW"),
		AmcMissionID:          unifieddatalibrary.String("AJM7939B1123"),
		ApacsID:               unifieddatalibrary.String("1083034"),
		AtoCallSign:           unifieddatalibrary.String("CHARLIE"),
		AtoMissionID:          unifieddatalibrary.String("8900"),
		CallSign:              unifieddatalibrary.String("RCH123"),
		Cw:                    unifieddatalibrary.Bool(true),
		DipWorksheetName:      unifieddatalibrary.String("G2-182402-AB"),
		FirstPickUp:           unifieddatalibrary.String("KFAY"),
		GdssMissionID:         unifieddatalibrary.String("1e6edeec-72e9-aaec-d33c-51147cb5ffdd"),
		HazMat: []unifieddatalibrary.AirTransportMissionNewParamsHazMat{{
			ApplicableNotes: unifieddatalibrary.String("11,12"),
			Cgc:             unifieddatalibrary.String("A"),
			Cgn:             unifieddatalibrary.String("4,5,7,8"),
			ClassDiv:        unifieddatalibrary.Float(1.1),
			ExtHazMatID:     unifieddatalibrary.String("cb6289e0f38534e01291ab6421d42724"),
			ItemName:        unifieddatalibrary.String("LITHIUM METAL BATTERIES"),
			NetExpWt:        unifieddatalibrary.Float(12.1),
			OffIcao:         unifieddatalibrary.String("MBPV"),
			OffItin:         unifieddatalibrary.Int(300),
			OnIcao:          unifieddatalibrary.String("LIRQ"),
			OnItin:          unifieddatalibrary.Int(50),
			Pieces:          unifieddatalibrary.Int(29),
			Planned:         unifieddatalibrary.String("P"),
			UnNum:           unifieddatalibrary.String("0181"),
			Weight:          unifieddatalibrary.Float(22.1),
		}},
		JcsPriority:      unifieddatalibrary.String("1A3"),
		LastDropOff:      unifieddatalibrary.String("PGUA"),
		LoadCategoryType: unifieddatalibrary.String("MIXED"),
		M1:               unifieddatalibrary.String("11"),
		M2:               unifieddatalibrary.String("3214"),
		M3a:              unifieddatalibrary.String("6655"),
		Naf:              unifieddatalibrary.String("18AF"),
		NextAmcMissionID: unifieddatalibrary.String("AJM7939B1124"),
		NextMissionID:    unifieddatalibrary.String("186e5658-1079-45c0-bccc-02d2fa31b663"),
		Node:             unifieddatalibrary.String("45TEST"),
		Objective:        unifieddatalibrary.String("Deliver water to island X."),
		Operation:        unifieddatalibrary.String("Golden Eye"),
		Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigMissionID:    unifieddatalibrary.String("614bebb6-a62e-053c-ca51-e79f8a402b28"),
		PrevAmcMissionID: unifieddatalibrary.String("AJM7939B1122"),
		PrevMissionID:    unifieddatalibrary.String("a77055df-edc3-4047-a5fa-604f80b9fe3c"),
		Purpose:          unifieddatalibrary.String("People at island X need water ASAP. Two previous attempts failed due to weather."),
		Remarks: []unifieddatalibrary.AirTransportMissionNewParamsRemark{{
			Date:         unifieddatalibrary.Time(time.Now()),
			GdssRemarkID: unifieddatalibrary.String("GDSSREMARK-ID"),
			ItineraryNum: unifieddatalibrary.Int(825),
			Text:         unifieddatalibrary.String("Example mission remarks."),
			Type:         unifieddatalibrary.String("MP"),
			User:         unifieddatalibrary.String("John Doe"),
		}},
		Requirements: []unifieddatalibrary.AirTransportMissionNewParamsRequirement{{
			BulkWeight:     unifieddatalibrary.Float(1.3),
			Ead:            unifieddatalibrary.Time(time.Now()),
			GdssReqID:      unifieddatalibrary.String("23a1fb67-cc2d-5ebe-6b99-68130cb1aa6c"),
			Lad:            unifieddatalibrary.Time(time.Now()),
			NumAmbulatory:  unifieddatalibrary.Int(10),
			NumAttendant:   unifieddatalibrary.Int(10),
			NumLitter:      unifieddatalibrary.Int(10),
			NumPax:         unifieddatalibrary.Int(44),
			OffloadID:      unifieddatalibrary.Int(300),
			OffloadLoCode:  unifieddatalibrary.String("KHOP"),
			OnloadID:       unifieddatalibrary.Int(200),
			OnloadLoCode:   unifieddatalibrary.String("KCHS"),
			Oplan:          unifieddatalibrary.String("5027"),
			OutsizeWeight:  unifieddatalibrary.Float(1.3),
			OversizeWeight: unifieddatalibrary.Float(1.3),
			ProjName:       unifieddatalibrary.String("CENTINTRA21"),
			TransReqNum:    unifieddatalibrary.String("T01ME01"),
			Uln:            unifieddatalibrary.String("T01ME01"),
		}},
		SourceSysDeviation: unifieddatalibrary.Float(-90.12),
		State:              unifieddatalibrary.String("EXECUTION"),
		Type:               unifieddatalibrary.String("SAAM"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirTransportMissionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AirTransportMissions.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.AirTransportMissionGetParams{
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

func TestAirTransportMissionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.AirTransportMissions.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AirTransportMissionUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AirTransportMissionUpdateParamsDataModeTest,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("bdad6945-c9e4-b829-f7be-1ad075541921"),
			Abp:                   unifieddatalibrary.String("ZZ12"),
			Alias:                 unifieddatalibrary.String("PACIFIC DEPLOY / CHAP 3 MOVEMENT"),
			AllocatedUnit:         unifieddatalibrary.String("437 AEW"),
			AmcMissionID:          unifieddatalibrary.String("AJM7939B1123"),
			ApacsID:               unifieddatalibrary.String("1083034"),
			AtoCallSign:           unifieddatalibrary.String("CHARLIE"),
			AtoMissionID:          unifieddatalibrary.String("8900"),
			CallSign:              unifieddatalibrary.String("RCH123"),
			Cw:                    unifieddatalibrary.Bool(true),
			DipWorksheetName:      unifieddatalibrary.String("G2-182402-AB"),
			FirstPickUp:           unifieddatalibrary.String("KFAY"),
			GdssMissionID:         unifieddatalibrary.String("1e6edeec-72e9-aaec-d33c-51147cb5ffdd"),
			HazMat: []unifieddatalibrary.AirTransportMissionUpdateParamsHazMat{{
				ApplicableNotes: unifieddatalibrary.String("11,12"),
				Cgc:             unifieddatalibrary.String("A"),
				Cgn:             unifieddatalibrary.String("4,5,7,8"),
				ClassDiv:        unifieddatalibrary.Float(1.1),
				ExtHazMatID:     unifieddatalibrary.String("cb6289e0f38534e01291ab6421d42724"),
				ItemName:        unifieddatalibrary.String("LITHIUM METAL BATTERIES"),
				NetExpWt:        unifieddatalibrary.Float(12.1),
				OffIcao:         unifieddatalibrary.String("MBPV"),
				OffItin:         unifieddatalibrary.Int(300),
				OnIcao:          unifieddatalibrary.String("LIRQ"),
				OnItin:          unifieddatalibrary.Int(50),
				Pieces:          unifieddatalibrary.Int(29),
				Planned:         unifieddatalibrary.String("P"),
				UnNum:           unifieddatalibrary.String("0181"),
				Weight:          unifieddatalibrary.Float(22.1),
			}},
			JcsPriority:      unifieddatalibrary.String("1A3"),
			LastDropOff:      unifieddatalibrary.String("PGUA"),
			LoadCategoryType: unifieddatalibrary.String("MIXED"),
			M1:               unifieddatalibrary.String("11"),
			M2:               unifieddatalibrary.String("3214"),
			M3a:              unifieddatalibrary.String("6655"),
			Naf:              unifieddatalibrary.String("18AF"),
			NextAmcMissionID: unifieddatalibrary.String("AJM7939B1124"),
			NextMissionID:    unifieddatalibrary.String("186e5658-1079-45c0-bccc-02d2fa31b663"),
			Node:             unifieddatalibrary.String("45TEST"),
			Objective:        unifieddatalibrary.String("Deliver water to island X."),
			Operation:        unifieddatalibrary.String("Golden Eye"),
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigMissionID:    unifieddatalibrary.String("614bebb6-a62e-053c-ca51-e79f8a402b28"),
			PrevAmcMissionID: unifieddatalibrary.String("AJM7939B1122"),
			PrevMissionID:    unifieddatalibrary.String("a77055df-edc3-4047-a5fa-604f80b9fe3c"),
			Purpose:          unifieddatalibrary.String("People at island X need water ASAP. Two previous attempts failed due to weather."),
			Remarks: []unifieddatalibrary.AirTransportMissionUpdateParamsRemark{{
				Date:         unifieddatalibrary.Time(time.Now()),
				GdssRemarkID: unifieddatalibrary.String("GDSSREMARK-ID"),
				ItineraryNum: unifieddatalibrary.Int(825),
				Text:         unifieddatalibrary.String("Example mission remarks."),
				Type:         unifieddatalibrary.String("MP"),
				User:         unifieddatalibrary.String("John Doe"),
			}},
			Requirements: []unifieddatalibrary.AirTransportMissionUpdateParamsRequirement{{
				BulkWeight:     unifieddatalibrary.Float(1.3),
				Ead:            unifieddatalibrary.Time(time.Now()),
				GdssReqID:      unifieddatalibrary.String("23a1fb67-cc2d-5ebe-6b99-68130cb1aa6c"),
				Lad:            unifieddatalibrary.Time(time.Now()),
				NumAmbulatory:  unifieddatalibrary.Int(10),
				NumAttendant:   unifieddatalibrary.Int(10),
				NumLitter:      unifieddatalibrary.Int(10),
				NumPax:         unifieddatalibrary.Int(44),
				OffloadID:      unifieddatalibrary.Int(300),
				OffloadLoCode:  unifieddatalibrary.String("KHOP"),
				OnloadID:       unifieddatalibrary.Int(200),
				OnloadLoCode:   unifieddatalibrary.String("KCHS"),
				Oplan:          unifieddatalibrary.String("5027"),
				OutsizeWeight:  unifieddatalibrary.Float(1.3),
				OversizeWeight: unifieddatalibrary.Float(1.3),
				ProjName:       unifieddatalibrary.String("CENTINTRA21"),
				TransReqNum:    unifieddatalibrary.String("T01ME01"),
				Uln:            unifieddatalibrary.String("T01ME01"),
			}},
			SourceSysDeviation: unifieddatalibrary.Float(-90.12),
			State:              unifieddatalibrary.String("EXECUTION"),
			Type:               unifieddatalibrary.String("SAAM"),
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

func TestAirTransportMissionListWithOptionalParams(t *testing.T) {
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
	_, err := client.AirTransportMissions.List(context.TODO(), unifieddatalibrary.AirTransportMissionListParams{
		CreatedAt:   time.Now(),
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

func TestAirTransportMissionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.AirTransportMissions.Count(context.TODO(), unifieddatalibrary.AirTransportMissionCountParams{
		CreatedAt:   time.Now(),
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

func TestAirTransportMissionQueryhelp(t *testing.T) {
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
	_, err := client.AirTransportMissions.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAirTransportMissionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.AirTransportMissions.Tuple(context.TODO(), unifieddatalibrary.AirTransportMissionTupleParams{
		Columns:     "columns",
		CreatedAt:   time.Now(),
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
