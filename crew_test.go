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

func TestCrewNewWithOptionalParams(t *testing.T) {
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
	err := client.Crew.New(context.TODO(), unifieddatalibrary.CrewNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.CrewNewParamsDataModeTest,
		OrigCrewID:            "JHJDHjhuu929o92",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("bdad6945-c9e4-b829-f7be-1ad075541921"),
		AdjReturnTime:         unifieddatalibrary.Time(time.Now()),
		AdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
		AircraftMds:           unifieddatalibrary.String("C017A"),
		AlertedTime:           unifieddatalibrary.Time(time.Now()),
		AlertType:             unifieddatalibrary.String("ALPHA"),
		ArmsCrewUnit:          unifieddatalibrary.String("00016ALSQ"),
		AssignedQualCode:      []string{"AL", "CS"},
		CommanderID:           unifieddatalibrary.String("763a1c1e8d2f3c16af825a11e3f1f579"),
		CommanderLast4Ssn:     unifieddatalibrary.String("1234"),
		CommanderName:         unifieddatalibrary.String("John Doe"),
		CrewHome:              unifieddatalibrary.Bool(false),
		CrewMembers: []unifieddatalibrary.CrewNewParamsCrewMember{{
			Alerted:                     unifieddatalibrary.Bool(true),
			AllSortie:                   unifieddatalibrary.Bool(true),
			Approved:                    unifieddatalibrary.Bool(true),
			Attached:                    unifieddatalibrary.Bool(true),
			Branch:                      unifieddatalibrary.String("Air Force"),
			Civilian:                    unifieddatalibrary.Bool(false),
			Commander:                   unifieddatalibrary.Bool(false),
			CrewPosition:                unifieddatalibrary.String("EP A"),
			DodID:                       unifieddatalibrary.String("0123456789"),
			DutyPosition:                unifieddatalibrary.String("IP"),
			DutyStatus:                  unifieddatalibrary.String("AGR"),
			Emailed:                     unifieddatalibrary.Bool(true),
			ExtraTime:                   unifieddatalibrary.Bool(true),
			FirstName:                   unifieddatalibrary.String("Freddie"),
			FltCurrencyExp:              unifieddatalibrary.Time(time.Now()),
			FltCurrencyExpID:            unifieddatalibrary.String("SS05AM"),
			FltRecDate:                  unifieddatalibrary.Time(time.Now()),
			FltRecDue:                   unifieddatalibrary.Time(time.Now()),
			FlySquadron:                 unifieddatalibrary.String("141ARS"),
			Funded:                      unifieddatalibrary.Bool(true),
			Gender:                      unifieddatalibrary.String("F"),
			GndCurrencyExp:              unifieddatalibrary.Time(time.Now()),
			GndCurrencyExpID:            unifieddatalibrary.String("AH03YM"),
			Grounded:                    unifieddatalibrary.Bool(true),
			GuestStart:                  unifieddatalibrary.Time(time.Now()),
			GuestStop:                   unifieddatalibrary.Time(time.Now()),
			Last4Ssn:                    unifieddatalibrary.String("1234"),
			LastFltDate:                 unifieddatalibrary.Time(time.Now()),
			LastName:                    unifieddatalibrary.String("Smith"),
			LoanedTo:                    unifieddatalibrary.String("Thunderbirds"),
			Lodging:                     unifieddatalibrary.String("Peterson SFB"),
			MemberActualAlertTime:       unifieddatalibrary.Time(time.Now()),
			MemberAdjReturnTime:         unifieddatalibrary.Time(time.Now()),
			MemberAdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
			MemberID:                    unifieddatalibrary.String("12345678abc"),
			MemberInitStartTime:         unifieddatalibrary.Time(time.Now()),
			MemberLastAlertTime:         unifieddatalibrary.Time(time.Now()),
			MemberLegalAlertTime:        unifieddatalibrary.Time(time.Now()),
			MemberPickupTime:            unifieddatalibrary.Time(time.Now()),
			MemberPostRestOffset:        unifieddatalibrary.String("+05:00"),
			MemberPostRestTime:          unifieddatalibrary.Time(time.Now()),
			MemberPreRestTime:           unifieddatalibrary.Time(time.Now()),
			MemberRemarks:               unifieddatalibrary.String("Crew member remark"),
			MemberReturnTime:            unifieddatalibrary.Time(time.Now()),
			MemberSchedAlertTime:        unifieddatalibrary.Time(time.Now()),
			MemberSource:                unifieddatalibrary.String("ACTIVE"),
			MemberStageName:             unifieddatalibrary.String("Falcon Squadron"),
			MemberTransportReq:          unifieddatalibrary.Bool(true),
			MemberType:                  unifieddatalibrary.String("AIRCREW"),
			MiddleInitial:               unifieddatalibrary.String("G"),
			Notified:                    unifieddatalibrary.Bool(true),
			PhoneNumber:                 unifieddatalibrary.String("+14155552671"),
			PhysAvCode:                  unifieddatalibrary.String("D"),
			PhysAvStatus:                unifieddatalibrary.String("OVERDUE"),
			PhysDue:                     unifieddatalibrary.Time(time.Now()),
			Rank:                        unifieddatalibrary.String("Capt"),
			RemarkCode:                  unifieddatalibrary.String("ABE33"),
			RmsMds:                      unifieddatalibrary.String("C017A"),
			ShowTime:                    unifieddatalibrary.Time(time.Now()),
			Squadron:                    unifieddatalibrary.String("21AS"),
			TrainingDate:                unifieddatalibrary.Time(time.Now()),
			Username:                    unifieddatalibrary.String("fgsmith"),
			Wing:                        unifieddatalibrary.String("60AMW"),
		}},
		CrewName:          unifieddatalibrary.String("falcon"),
		CrewRms:           unifieddatalibrary.String("ARMS"),
		CrewRole:          unifieddatalibrary.String("DEADHEAD"),
		CrewSource:        unifieddatalibrary.String("ACTIVE"),
		CrewSquadron:      unifieddatalibrary.String("21AS"),
		CrewType:          unifieddatalibrary.String("AIRLAND"),
		CrewUnit:          unifieddatalibrary.String("00016ALSQ"),
		CrewWing:          unifieddatalibrary.String("60AMW"),
		CurrentIcao:       unifieddatalibrary.String("KCOS"),
		FdpEligType:       unifieddatalibrary.String("A"),
		FdpType:           unifieddatalibrary.String("A"),
		FemaleEnlistedQty: unifieddatalibrary.Int(2),
		FemaleOfficerQty:  unifieddatalibrary.Int(1),
		FltAuthNum:        unifieddatalibrary.String("KT001"),
		IDSiteCurrent:     unifieddatalibrary.String("b677cf3b-d44d-450e-8b8f-d23f997f8778"),
		IDSortie:          unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
		InitStartTime:     unifieddatalibrary.Time(time.Now()),
		LastAlertTime:     unifieddatalibrary.Time(time.Now()),
		LegalAlertTime:    unifieddatalibrary.Time(time.Now()),
		LegalBravoTime:    unifieddatalibrary.Time(time.Now()),
		LinkedTask:        unifieddatalibrary.Bool(false),
		MaleEnlistedQty:   unifieddatalibrary.Int(3),
		MaleOfficerQty:    unifieddatalibrary.Int(1),
		MissionAlias:      unifieddatalibrary.String("PACIFIC DEPLOY / CHAP 3 MOVEMENT"),
		MissionID:         unifieddatalibrary.String("AJM123456123"),
		Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PersonnelType:     unifieddatalibrary.String("AIRCREW"),
		PickupTime:        unifieddatalibrary.Time(time.Now()),
		PostRestApplied:   unifieddatalibrary.Bool(false),
		PostRestEnd:       unifieddatalibrary.Time(time.Now()),
		PostRestOffset:    unifieddatalibrary.String("+05:00"),
		PreRestApplied:    unifieddatalibrary.Bool(false),
		PreRestStart:      unifieddatalibrary.Time(time.Now()),
		ReqQualCode:       []string{"AL", "CS"},
		ReturnTime:        unifieddatalibrary.Time(time.Now()),
		Stage1Qual:        unifieddatalibrary.String("1AXXX"),
		Stage2Qual:        unifieddatalibrary.String("2AXXX"),
		Stage3Qual:        unifieddatalibrary.String("3AXXX"),
		StageName:         unifieddatalibrary.String("Falcon Squadron"),
		StageTime:         unifieddatalibrary.Time(time.Now()),
		Status:            unifieddatalibrary.String("APPROVED"),
		TransportReq:      unifieddatalibrary.Bool(true),
		TripKit:           unifieddatalibrary.String("TK-1234"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrewGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crew.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.CrewGetParams{
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

func TestCrewUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Crew.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.CrewUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.CrewUpdateParamsDataModeTest,
			OrigCrewID:            "JHJDHjhuu929o92",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("bdad6945-c9e4-b829-f7be-1ad075541921"),
			AdjReturnTime:         unifieddatalibrary.Time(time.Now()),
			AdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
			AircraftMds:           unifieddatalibrary.String("C017A"),
			AlertedTime:           unifieddatalibrary.Time(time.Now()),
			AlertType:             unifieddatalibrary.String("ALPHA"),
			ArmsCrewUnit:          unifieddatalibrary.String("00016ALSQ"),
			AssignedQualCode:      []string{"AL", "CS"},
			CommanderID:           unifieddatalibrary.String("763a1c1e8d2f3c16af825a11e3f1f579"),
			CommanderLast4Ssn:     unifieddatalibrary.String("1234"),
			CommanderName:         unifieddatalibrary.String("John Doe"),
			CrewHome:              unifieddatalibrary.Bool(false),
			CrewMembers: []unifieddatalibrary.CrewUpdateParamsCrewMember{{
				Alerted:                     unifieddatalibrary.Bool(true),
				AllSortie:                   unifieddatalibrary.Bool(true),
				Approved:                    unifieddatalibrary.Bool(true),
				Attached:                    unifieddatalibrary.Bool(true),
				Branch:                      unifieddatalibrary.String("Air Force"),
				Civilian:                    unifieddatalibrary.Bool(false),
				Commander:                   unifieddatalibrary.Bool(false),
				CrewPosition:                unifieddatalibrary.String("EP A"),
				DodID:                       unifieddatalibrary.String("0123456789"),
				DutyPosition:                unifieddatalibrary.String("IP"),
				DutyStatus:                  unifieddatalibrary.String("AGR"),
				Emailed:                     unifieddatalibrary.Bool(true),
				ExtraTime:                   unifieddatalibrary.Bool(true),
				FirstName:                   unifieddatalibrary.String("Freddie"),
				FltCurrencyExp:              unifieddatalibrary.Time(time.Now()),
				FltCurrencyExpID:            unifieddatalibrary.String("SS05AM"),
				FltRecDate:                  unifieddatalibrary.Time(time.Now()),
				FltRecDue:                   unifieddatalibrary.Time(time.Now()),
				FlySquadron:                 unifieddatalibrary.String("141ARS"),
				Funded:                      unifieddatalibrary.Bool(true),
				Gender:                      unifieddatalibrary.String("F"),
				GndCurrencyExp:              unifieddatalibrary.Time(time.Now()),
				GndCurrencyExpID:            unifieddatalibrary.String("AH03YM"),
				Grounded:                    unifieddatalibrary.Bool(true),
				GuestStart:                  unifieddatalibrary.Time(time.Now()),
				GuestStop:                   unifieddatalibrary.Time(time.Now()),
				Last4Ssn:                    unifieddatalibrary.String("1234"),
				LastFltDate:                 unifieddatalibrary.Time(time.Now()),
				LastName:                    unifieddatalibrary.String("Smith"),
				LoanedTo:                    unifieddatalibrary.String("Thunderbirds"),
				Lodging:                     unifieddatalibrary.String("Peterson SFB"),
				MemberActualAlertTime:       unifieddatalibrary.Time(time.Now()),
				MemberAdjReturnTime:         unifieddatalibrary.Time(time.Now()),
				MemberAdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
				MemberID:                    unifieddatalibrary.String("12345678abc"),
				MemberInitStartTime:         unifieddatalibrary.Time(time.Now()),
				MemberLastAlertTime:         unifieddatalibrary.Time(time.Now()),
				MemberLegalAlertTime:        unifieddatalibrary.Time(time.Now()),
				MemberPickupTime:            unifieddatalibrary.Time(time.Now()),
				MemberPostRestOffset:        unifieddatalibrary.String("+05:00"),
				MemberPostRestTime:          unifieddatalibrary.Time(time.Now()),
				MemberPreRestTime:           unifieddatalibrary.Time(time.Now()),
				MemberRemarks:               unifieddatalibrary.String("Crew member remark"),
				MemberReturnTime:            unifieddatalibrary.Time(time.Now()),
				MemberSchedAlertTime:        unifieddatalibrary.Time(time.Now()),
				MemberSource:                unifieddatalibrary.String("ACTIVE"),
				MemberStageName:             unifieddatalibrary.String("Falcon Squadron"),
				MemberTransportReq:          unifieddatalibrary.Bool(true),
				MemberType:                  unifieddatalibrary.String("AIRCREW"),
				MiddleInitial:               unifieddatalibrary.String("G"),
				Notified:                    unifieddatalibrary.Bool(true),
				PhoneNumber:                 unifieddatalibrary.String("+14155552671"),
				PhysAvCode:                  unifieddatalibrary.String("D"),
				PhysAvStatus:                unifieddatalibrary.String("OVERDUE"),
				PhysDue:                     unifieddatalibrary.Time(time.Now()),
				Rank:                        unifieddatalibrary.String("Capt"),
				RemarkCode:                  unifieddatalibrary.String("ABE33"),
				RmsMds:                      unifieddatalibrary.String("C017A"),
				ShowTime:                    unifieddatalibrary.Time(time.Now()),
				Squadron:                    unifieddatalibrary.String("21AS"),
				TrainingDate:                unifieddatalibrary.Time(time.Now()),
				Username:                    unifieddatalibrary.String("fgsmith"),
				Wing:                        unifieddatalibrary.String("60AMW"),
			}},
			CrewName:          unifieddatalibrary.String("falcon"),
			CrewRms:           unifieddatalibrary.String("ARMS"),
			CrewRole:          unifieddatalibrary.String("DEADHEAD"),
			CrewSource:        unifieddatalibrary.String("ACTIVE"),
			CrewSquadron:      unifieddatalibrary.String("21AS"),
			CrewType:          unifieddatalibrary.String("AIRLAND"),
			CrewUnit:          unifieddatalibrary.String("00016ALSQ"),
			CrewWing:          unifieddatalibrary.String("60AMW"),
			CurrentIcao:       unifieddatalibrary.String("KCOS"),
			FdpEligType:       unifieddatalibrary.String("A"),
			FdpType:           unifieddatalibrary.String("A"),
			FemaleEnlistedQty: unifieddatalibrary.Int(2),
			FemaleOfficerQty:  unifieddatalibrary.Int(1),
			FltAuthNum:        unifieddatalibrary.String("KT001"),
			IDSiteCurrent:     unifieddatalibrary.String("b677cf3b-d44d-450e-8b8f-d23f997f8778"),
			IDSortie:          unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
			InitStartTime:     unifieddatalibrary.Time(time.Now()),
			LastAlertTime:     unifieddatalibrary.Time(time.Now()),
			LegalAlertTime:    unifieddatalibrary.Time(time.Now()),
			LegalBravoTime:    unifieddatalibrary.Time(time.Now()),
			LinkedTask:        unifieddatalibrary.Bool(false),
			MaleEnlistedQty:   unifieddatalibrary.Int(3),
			MaleOfficerQty:    unifieddatalibrary.Int(1),
			MissionAlias:      unifieddatalibrary.String("PACIFIC DEPLOY / CHAP 3 MOVEMENT"),
			MissionID:         unifieddatalibrary.String("AJM123456123"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PersonnelType:     unifieddatalibrary.String("AIRCREW"),
			PickupTime:        unifieddatalibrary.Time(time.Now()),
			PostRestApplied:   unifieddatalibrary.Bool(false),
			PostRestEnd:       unifieddatalibrary.Time(time.Now()),
			PostRestOffset:    unifieddatalibrary.String("+05:00"),
			PreRestApplied:    unifieddatalibrary.Bool(false),
			PreRestStart:      unifieddatalibrary.Time(time.Now()),
			ReqQualCode:       []string{"AL", "CS"},
			ReturnTime:        unifieddatalibrary.Time(time.Now()),
			Stage1Qual:        unifieddatalibrary.String("1AXXX"),
			Stage2Qual:        unifieddatalibrary.String("2AXXX"),
			Stage3Qual:        unifieddatalibrary.String("3AXXX"),
			StageName:         unifieddatalibrary.String("Falcon Squadron"),
			StageTime:         unifieddatalibrary.Time(time.Now()),
			Status:            unifieddatalibrary.String("APPROVED"),
			TransportReq:      unifieddatalibrary.Bool(true),
			TripKit:           unifieddatalibrary.String("TK-1234"),
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

func TestCrewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Crew.List(context.TODO(), unifieddatalibrary.CrewListParams{
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

func TestCrewCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Crew.Count(context.TODO(), unifieddatalibrary.CrewCountParams{
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

func TestCrewQueryhelp(t *testing.T) {
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
	_, err := client.Crew.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrewTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Crew.Tuple(context.TODO(), unifieddatalibrary.CrewTupleParams{
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

func TestCrewUnvalidatedPublish(t *testing.T) {
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
	err := client.Crew.UnvalidatedPublish(context.TODO(), unifieddatalibrary.CrewUnvalidatedPublishParams{
		Body: []unifieddatalibrary.CrewUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			OrigCrewID:            "JHJDHjhuu929o92",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("bdad6945-c9e4-b829-f7be-1ad075541921"),
			AdjReturnTime:         unifieddatalibrary.Time(time.Now()),
			AdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
			AircraftMds:           unifieddatalibrary.String("C017A"),
			AlertedTime:           unifieddatalibrary.Time(time.Now()),
			AlertType:             unifieddatalibrary.String("ALPHA"),
			ArmsCrewUnit:          unifieddatalibrary.String("00016ALSQ"),
			AssignedQualCode:      []string{"AL", "CS"},
			CommanderID:           unifieddatalibrary.String("763a1c1e8d2f3c16af825a11e3f1f579"),
			CommanderLast4Ssn:     unifieddatalibrary.String("1234"),
			CommanderName:         unifieddatalibrary.String("John Doe"),
			CrewHome:              unifieddatalibrary.Bool(false),
			CrewMembers: []unifieddatalibrary.CrewUnvalidatedPublishParamsBodyCrewMember{{
				Alerted:                     unifieddatalibrary.Bool(true),
				AllSortie:                   unifieddatalibrary.Bool(true),
				Approved:                    unifieddatalibrary.Bool(true),
				Attached:                    unifieddatalibrary.Bool(true),
				Branch:                      unifieddatalibrary.String("Air Force"),
				Civilian:                    unifieddatalibrary.Bool(false),
				Commander:                   unifieddatalibrary.Bool(false),
				CrewPosition:                unifieddatalibrary.String("EP A"),
				DodID:                       unifieddatalibrary.String("0123456789"),
				DutyPosition:                unifieddatalibrary.String("IP"),
				DutyStatus:                  unifieddatalibrary.String("AGR"),
				Emailed:                     unifieddatalibrary.Bool(true),
				ExtraTime:                   unifieddatalibrary.Bool(true),
				FirstName:                   unifieddatalibrary.String("Freddie"),
				FltCurrencyExp:              unifieddatalibrary.Time(time.Now()),
				FltCurrencyExpID:            unifieddatalibrary.String("SS05AM"),
				FltRecDate:                  unifieddatalibrary.Time(time.Now()),
				FltRecDue:                   unifieddatalibrary.Time(time.Now()),
				FlySquadron:                 unifieddatalibrary.String("141ARS"),
				Funded:                      unifieddatalibrary.Bool(true),
				Gender:                      unifieddatalibrary.String("F"),
				GndCurrencyExp:              unifieddatalibrary.Time(time.Now()),
				GndCurrencyExpID:            unifieddatalibrary.String("AH03YM"),
				Grounded:                    unifieddatalibrary.Bool(true),
				GuestStart:                  unifieddatalibrary.Time(time.Now()),
				GuestStop:                   unifieddatalibrary.Time(time.Now()),
				Last4Ssn:                    unifieddatalibrary.String("1234"),
				LastFltDate:                 unifieddatalibrary.Time(time.Now()),
				LastName:                    unifieddatalibrary.String("Smith"),
				LoanedTo:                    unifieddatalibrary.String("Thunderbirds"),
				Lodging:                     unifieddatalibrary.String("Peterson SFB"),
				MemberActualAlertTime:       unifieddatalibrary.Time(time.Now()),
				MemberAdjReturnTime:         unifieddatalibrary.Time(time.Now()),
				MemberAdjReturnTimeApprover: unifieddatalibrary.String("Smith"),
				MemberID:                    unifieddatalibrary.String("12345678abc"),
				MemberInitStartTime:         unifieddatalibrary.Time(time.Now()),
				MemberLastAlertTime:         unifieddatalibrary.Time(time.Now()),
				MemberLegalAlertTime:        unifieddatalibrary.Time(time.Now()),
				MemberPickupTime:            unifieddatalibrary.Time(time.Now()),
				MemberPostRestOffset:        unifieddatalibrary.String("+05:00"),
				MemberPostRestTime:          unifieddatalibrary.Time(time.Now()),
				MemberPreRestTime:           unifieddatalibrary.Time(time.Now()),
				MemberRemarks:               unifieddatalibrary.String("Crew member remark"),
				MemberReturnTime:            unifieddatalibrary.Time(time.Now()),
				MemberSchedAlertTime:        unifieddatalibrary.Time(time.Now()),
				MemberSource:                unifieddatalibrary.String("ACTIVE"),
				MemberStageName:             unifieddatalibrary.String("Falcon Squadron"),
				MemberTransportReq:          unifieddatalibrary.Bool(true),
				MemberType:                  unifieddatalibrary.String("AIRCREW"),
				MiddleInitial:               unifieddatalibrary.String("G"),
				Notified:                    unifieddatalibrary.Bool(true),
				PhoneNumber:                 unifieddatalibrary.String("+14155552671"),
				PhysAvCode:                  unifieddatalibrary.String("D"),
				PhysAvStatus:                unifieddatalibrary.String("OVERDUE"),
				PhysDue:                     unifieddatalibrary.Time(time.Now()),
				Rank:                        unifieddatalibrary.String("Capt"),
				RemarkCode:                  unifieddatalibrary.String("ABE33"),
				RmsMds:                      unifieddatalibrary.String("C017A"),
				ShowTime:                    unifieddatalibrary.Time(time.Now()),
				Squadron:                    unifieddatalibrary.String("21AS"),
				TrainingDate:                unifieddatalibrary.Time(time.Now()),
				Username:                    unifieddatalibrary.String("fgsmith"),
				Wing:                        unifieddatalibrary.String("60AMW"),
			}},
			CrewName:          unifieddatalibrary.String("falcon"),
			CrewRms:           unifieddatalibrary.String("ARMS"),
			CrewRole:          unifieddatalibrary.String("DEADHEAD"),
			CrewSource:        unifieddatalibrary.String("ACTIVE"),
			CrewSquadron:      unifieddatalibrary.String("21AS"),
			CrewType:          unifieddatalibrary.String("AIRLAND"),
			CrewUnit:          unifieddatalibrary.String("00016ALSQ"),
			CrewWing:          unifieddatalibrary.String("60AMW"),
			CurrentIcao:       unifieddatalibrary.String("KCOS"),
			FdpEligType:       unifieddatalibrary.String("A"),
			FdpType:           unifieddatalibrary.String("A"),
			FemaleEnlistedQty: unifieddatalibrary.Int(2),
			FemaleOfficerQty:  unifieddatalibrary.Int(1),
			FltAuthNum:        unifieddatalibrary.String("KT001"),
			IDSiteCurrent:     unifieddatalibrary.String("b677cf3b-d44d-450e-8b8f-d23f997f8778"),
			IDSortie:          unifieddatalibrary.String("4ef3d1e8-ab08-ab70-498f-edc479734e5c"),
			InitStartTime:     unifieddatalibrary.Time(time.Now()),
			LastAlertTime:     unifieddatalibrary.Time(time.Now()),
			LegalAlertTime:    unifieddatalibrary.Time(time.Now()),
			LegalBravoTime:    unifieddatalibrary.Time(time.Now()),
			LinkedTask:        unifieddatalibrary.Bool(false),
			MaleEnlistedQty:   unifieddatalibrary.Int(3),
			MaleOfficerQty:    unifieddatalibrary.Int(1),
			MissionAlias:      unifieddatalibrary.String("PACIFIC DEPLOY / CHAP 3 MOVEMENT"),
			MissionID:         unifieddatalibrary.String("AJM123456123"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PersonnelType:     unifieddatalibrary.String("AIRCREW"),
			PickupTime:        unifieddatalibrary.Time(time.Now()),
			PostRestApplied:   unifieddatalibrary.Bool(false),
			PostRestEnd:       unifieddatalibrary.Time(time.Now()),
			PostRestOffset:    unifieddatalibrary.String("+05:00"),
			PreRestApplied:    unifieddatalibrary.Bool(false),
			PreRestStart:      unifieddatalibrary.Time(time.Now()),
			ReqQualCode:       []string{"AL", "CS"},
			ReturnTime:        unifieddatalibrary.Time(time.Now()),
			Stage1Qual:        unifieddatalibrary.String("1AXXX"),
			Stage2Qual:        unifieddatalibrary.String("2AXXX"),
			Stage3Qual:        unifieddatalibrary.String("3AXXX"),
			StageName:         unifieddatalibrary.String("Falcon Squadron"),
			StageTime:         unifieddatalibrary.Time(time.Now()),
			Status:            unifieddatalibrary.String("APPROVED"),
			TransportReq:      unifieddatalibrary.Bool(true),
			TripKit:           unifieddatalibrary.String("TK-1234"),
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
