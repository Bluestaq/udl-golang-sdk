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

func TestOnorbiteventNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitevent.New(context.TODO(), unifieddatalibrary.OnorbiteventNewParams{
		ClassificationMarking:         "U",
		DataMode:                      unifieddatalibrary.OnorbiteventNewParamsDataModeTest,
		EventTime:                     time.Now(),
		Source:                        "Bluestaq",
		ID:                            unifieddatalibrary.String("ONORBITEVENT-ID"),
		AchievedFlightPhase:           unifieddatalibrary.String("Phase 2"),
		AgeAtEvent:                    unifieddatalibrary.Float(5.23),
		CapabilityLoss:                unifieddatalibrary.Float(0.5),
		CapabilityLossNotes:           unifieddatalibrary.String("Example notes"),
		CapacityLoss:                  unifieddatalibrary.Float(0.5),
		ConsequentialEquipmentFailure: unifieddatalibrary.String("Example Equipment"),
		DeclassificationDate:          unifieddatalibrary.Time(time.Now()),
		DeclassificationString:        unifieddatalibrary.String("DECLASS_STRING"),
		DerivedFrom:                   unifieddatalibrary.String("DERIVED_SOURCE"),
		Description:                   unifieddatalibrary.String("Example notes"),
		EquipmentAtFault:              unifieddatalibrary.String("Example Equipment"),
		EquipmentCausingLossNotes:     unifieddatalibrary.String("Example notes"),
		EquipmentPartAtFault:          unifieddatalibrary.String("Example Equipment"),
		EquipmentTypeAtFault:          unifieddatalibrary.String("Example Equipment"),
		EventResult:                   unifieddatalibrary.String("Example results"),
		EventTimeNotes:                unifieddatalibrary.String("Notes on validity"),
		EventType:                     unifieddatalibrary.String("Type1"),
		GeoPosition:                   unifieddatalibrary.Float(45.23),
		IDOnOrbit:                     unifieddatalibrary.String("ONORBIT-ID"),
		Inclined:                      unifieddatalibrary.Bool(false),
		Injured:                       unifieddatalibrary.Int(1),
		InsuranceCarriedNotes:         unifieddatalibrary.String("Insurance notes"),
		InsuranceLoss:                 unifieddatalibrary.Float(0.5),
		InsuranceLossNotes:            unifieddatalibrary.String("Insurance notes"),
		Killed:                        unifieddatalibrary.Int(23),
		LesseeOrgID:                   unifieddatalibrary.String("LESSEEORG-ID"),
		LifeLost:                      unifieddatalibrary.Float(0.5),
		NetAmount:                     unifieddatalibrary.Float(10000.23),
		ObjectStatus:                  unifieddatalibrary.String("Status1"),
		OccurrenceFlightPhase:         unifieddatalibrary.String("Phase 2"),
		OfficialLossDate:              unifieddatalibrary.Time(time.Now()),
		OperatedOnBehalfOfOrgID:       unifieddatalibrary.String("OPERATEDONBEHALFOFORG-ID"),
		OperatorOrgID:                 unifieddatalibrary.String("OPERATORORG-ID"),
		Origin:                        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:                  unifieddatalibrary.String("ORIGOBJECT-ID"),
		OwnerOrgID:                    unifieddatalibrary.String("OWNERORG-ID"),
		PlaneNumber:                   unifieddatalibrary.String("PL_1"),
		PlaneSlot:                     unifieddatalibrary.String("example_slot"),
		PositionStatus:                unifieddatalibrary.String("Stable"),
		Remarks:                       unifieddatalibrary.String("Example remarks"),
		SatellitePosition:             unifieddatalibrary.String("Example description"),
		SatNo:                         unifieddatalibrary.Int(1),
		StageAtFault:                  unifieddatalibrary.String("Phase 2"),
		ThirdPartyInsuranceLoss:       unifieddatalibrary.Float(10000.23),
		UnderlyingCause:               unifieddatalibrary.String("CAUSE_EXAMPLE"),
		UntilTime:                     unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbiteventUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitevent.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbiteventUpdateParams{
			ClassificationMarking:         "U",
			DataMode:                      unifieddatalibrary.OnorbiteventUpdateParamsDataModeTest,
			EventTime:                     time.Now(),
			Source:                        "Bluestaq",
			ID:                            unifieddatalibrary.String("ONORBITEVENT-ID"),
			AchievedFlightPhase:           unifieddatalibrary.String("Phase 2"),
			AgeAtEvent:                    unifieddatalibrary.Float(5.23),
			CapabilityLoss:                unifieddatalibrary.Float(0.5),
			CapabilityLossNotes:           unifieddatalibrary.String("Example notes"),
			CapacityLoss:                  unifieddatalibrary.Float(0.5),
			ConsequentialEquipmentFailure: unifieddatalibrary.String("Example Equipment"),
			DeclassificationDate:          unifieddatalibrary.Time(time.Now()),
			DeclassificationString:        unifieddatalibrary.String("DECLASS_STRING"),
			DerivedFrom:                   unifieddatalibrary.String("DERIVED_SOURCE"),
			Description:                   unifieddatalibrary.String("Example notes"),
			EquipmentAtFault:              unifieddatalibrary.String("Example Equipment"),
			EquipmentCausingLossNotes:     unifieddatalibrary.String("Example notes"),
			EquipmentPartAtFault:          unifieddatalibrary.String("Example Equipment"),
			EquipmentTypeAtFault:          unifieddatalibrary.String("Example Equipment"),
			EventResult:                   unifieddatalibrary.String("Example results"),
			EventTimeNotes:                unifieddatalibrary.String("Notes on validity"),
			EventType:                     unifieddatalibrary.String("Type1"),
			GeoPosition:                   unifieddatalibrary.Float(45.23),
			IDOnOrbit:                     unifieddatalibrary.String("ONORBIT-ID"),
			Inclined:                      unifieddatalibrary.Bool(false),
			Injured:                       unifieddatalibrary.Int(1),
			InsuranceCarriedNotes:         unifieddatalibrary.String("Insurance notes"),
			InsuranceLoss:                 unifieddatalibrary.Float(0.5),
			InsuranceLossNotes:            unifieddatalibrary.String("Insurance notes"),
			Killed:                        unifieddatalibrary.Int(23),
			LesseeOrgID:                   unifieddatalibrary.String("LESSEEORG-ID"),
			LifeLost:                      unifieddatalibrary.Float(0.5),
			NetAmount:                     unifieddatalibrary.Float(10000.23),
			ObjectStatus:                  unifieddatalibrary.String("Status1"),
			OccurrenceFlightPhase:         unifieddatalibrary.String("Phase 2"),
			OfficialLossDate:              unifieddatalibrary.Time(time.Now()),
			OperatedOnBehalfOfOrgID:       unifieddatalibrary.String("OPERATEDONBEHALFOFORG-ID"),
			OperatorOrgID:                 unifieddatalibrary.String("OPERATORORG-ID"),
			Origin:                        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:                  unifieddatalibrary.String("ORIGOBJECT-ID"),
			OwnerOrgID:                    unifieddatalibrary.String("OWNERORG-ID"),
			PlaneNumber:                   unifieddatalibrary.String("PL_1"),
			PlaneSlot:                     unifieddatalibrary.String("example_slot"),
			PositionStatus:                unifieddatalibrary.String("Stable"),
			Remarks:                       unifieddatalibrary.String("Example remarks"),
			SatellitePosition:             unifieddatalibrary.String("Example description"),
			SatNo:                         unifieddatalibrary.Int(1),
			StageAtFault:                  unifieddatalibrary.String("Phase 2"),
			ThirdPartyInsuranceLoss:       unifieddatalibrary.Float(10000.23),
			UnderlyingCause:               unifieddatalibrary.String("CAUSE_EXAMPLE"),
			UntilTime:                     unifieddatalibrary.Time(time.Now()),
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

func TestOnorbiteventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitevent.List(context.TODO(), unifieddatalibrary.OnorbiteventListParams{
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

func TestOnorbiteventDelete(t *testing.T) {
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
	err := client.Onorbitevent.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbiteventCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitevent.Count(context.TODO(), unifieddatalibrary.OnorbiteventCountParams{
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

func TestOnorbiteventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitevent.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbiteventGetParams{
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

func TestOnorbiteventQueryhelp(t *testing.T) {
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
	err := client.Onorbitevent.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbiteventTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitevent.Tuple(context.TODO(), unifieddatalibrary.OnorbiteventTupleParams{
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
