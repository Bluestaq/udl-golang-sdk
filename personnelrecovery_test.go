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

func TestPersonnelrecoveryNewWithOptionalParams(t *testing.T) {
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
	err := client.Personnelrecovery.New(context.TODO(), unifieddatalibrary.PersonnelrecoveryNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.PersonnelrecoveryNewParamsDataModeTest,
		MsgTime:               time.Now(),
		PickupLat:             75.1234,
		PickupLon:             175.1234,
		Source:                "Bluestaq",
		Type:                  "MEDICAL",
		ID:                    unifieddatalibrary.String("PERSONNEL_RECOVERY-ID"),
		AuthMethod:            unifieddatalibrary.String("PASSPORT"),
		AuthStatus:            unifieddatalibrary.String("NO STATEMENT"),
		BeaconInd:             unifieddatalibrary.Bool(false),
		CallSign:              unifieddatalibrary.String("BADGER"),
		CommEq1:               unifieddatalibrary.String("LL PHONE"),
		CommEq2:               unifieddatalibrary.String("LL PHONE"),
		CommEq3:               unifieddatalibrary.String("LL PHONE"),
		ExecutionInfo: unifieddatalibrary.PersonnelrecoveryNewParamsExecutionInfo{
			Egress:      unifieddatalibrary.Float(66.53),
			EgressPoint: []float64{107.23, 30.455},
			EscortVehicle: unifieddatalibrary.PersonnelrecoveryNewParamsExecutionInfoEscortVehicle{
				CallSign:    unifieddatalibrary.String("FALCO"),
				PrimaryFreq: unifieddatalibrary.Float(34.55),
				Strength:    unifieddatalibrary.Int(5),
				Type:        unifieddatalibrary.String("C17"),
			},
			Ingress:      unifieddatalibrary.Float(35.66),
			InitialPoint: []float64{103.23, 30.445},
			ObjStrategy:  unifieddatalibrary.String("Description of strategy plan."),
			RecoveryVehicle: unifieddatalibrary.PersonnelrecoveryNewParamsExecutionInfoRecoveryVehicle{
				CallSign:    unifieddatalibrary.String("FALCO"),
				PrimaryFreq: unifieddatalibrary.Float(34.55),
				Strength:    unifieddatalibrary.Int(5),
				Type:        unifieddatalibrary.String("C17"),
			},
		},
		Identity:             unifieddatalibrary.String("NEUTRAL CIVILIAN"),
		IDWeatherReport:      unifieddatalibrary.String("WEATHER_REPORT-ID"),
		MilClass:             unifieddatalibrary.String("CIVILIAN"),
		NatAlliance:          unifieddatalibrary.Int(1),
		NatAlliance1:         unifieddatalibrary.Int(0),
		NumAmbulatory:        unifieddatalibrary.Int(1),
		NumAmbulatoryInjured: unifieddatalibrary.Int(2),
		NumNonAmbulatory:     unifieddatalibrary.Int(0),
		NumPersons:           unifieddatalibrary.Int(1),
		ObjectiveAreaInfo: unifieddatalibrary.PersonnelrecoveryNewParamsObjectiveAreaInfo{
			EnemyData: []unifieddatalibrary.PersonnelrecoveryNewParamsObjectiveAreaInfoEnemyData{{
				DirToEnemy:        unifieddatalibrary.String("NORTHWEST"),
				FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
				HlzRemarks:        unifieddatalibrary.String("Hot Landing Zone remarks."),
				HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
			}},
			OscCallSign: unifieddatalibrary.String("STARFOX"),
			OscFreq:     unifieddatalibrary.Float(12.55),
			PzDesc:      unifieddatalibrary.String("Near the lake."),
			PzLocation:  []float64{103.23, 30.445},
		},
		Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PickupAlt:        unifieddatalibrary.Float(30.1234),
		RecovID:          unifieddatalibrary.String("RECOV-ID"),
		RxFreq:           unifieddatalibrary.Float(5.5),
		SurvivorMessages: unifieddatalibrary.String("UNINJURED CANT MOVE HOSTILES NEARBY"),
		SurvivorRadio:    unifieddatalibrary.String("NO STATEMENT"),
		TermInd:          unifieddatalibrary.Bool(true),
		TextMsg:          unifieddatalibrary.String("Additional message from survivor."),
		TxFreq:           unifieddatalibrary.Float(5.5),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPersonnelrecoveryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Personnelrecovery.List(context.TODO(), unifieddatalibrary.PersonnelrecoveryListParams{
		MsgTime:     time.Now(),
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

func TestPersonnelrecoveryCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Personnelrecovery.Count(context.TODO(), unifieddatalibrary.PersonnelrecoveryCountParams{
		MsgTime:     time.Now(),
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

func TestPersonnelrecoveryNewBulk(t *testing.T) {
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
	err := client.Personnelrecovery.NewBulk(context.TODO(), unifieddatalibrary.PersonnelrecoveryNewBulkParams{
		Body: []unifieddatalibrary.PersonnelrecoveryNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			MsgTime:               time.Now(),
			PickupLat:             75.1234,
			PickupLon:             175.1234,
			Source:                "Bluestaq",
			Type:                  "MEDICAL",
			ID:                    unifieddatalibrary.String("PERSONNEL_RECOVERY-ID"),
			AuthMethod:            unifieddatalibrary.String("PASSPORT"),
			AuthStatus:            unifieddatalibrary.String("NO STATEMENT"),
			BeaconInd:             unifieddatalibrary.Bool(false),
			CallSign:              unifieddatalibrary.String("BADGER"),
			CommEq1:               unifieddatalibrary.String("LL PHONE"),
			CommEq2:               unifieddatalibrary.String("LL PHONE"),
			CommEq3:               unifieddatalibrary.String("LL PHONE"),
			ExecutionInfo: unifieddatalibrary.PersonnelrecoveryNewBulkParamsBodyExecutionInfo{
				Egress:      unifieddatalibrary.Float(66.53),
				EgressPoint: []float64{107.23, 30.455},
				EscortVehicle: unifieddatalibrary.PersonnelrecoveryNewBulkParamsBodyExecutionInfoEscortVehicle{
					CallSign:    unifieddatalibrary.String("FALCO"),
					PrimaryFreq: unifieddatalibrary.Float(34.55),
					Strength:    unifieddatalibrary.Int(5),
					Type:        unifieddatalibrary.String("C17"),
				},
				Ingress:      unifieddatalibrary.Float(35.66),
				InitialPoint: []float64{103.23, 30.445},
				ObjStrategy:  unifieddatalibrary.String("Description of strategy plan."),
				RecoveryVehicle: unifieddatalibrary.PersonnelrecoveryNewBulkParamsBodyExecutionInfoRecoveryVehicle{
					CallSign:    unifieddatalibrary.String("FALCO"),
					PrimaryFreq: unifieddatalibrary.Float(34.55),
					Strength:    unifieddatalibrary.Int(5),
					Type:        unifieddatalibrary.String("C17"),
				},
			},
			Identity:             unifieddatalibrary.String("NEUTRAL CIVILIAN"),
			IDWeatherReport:      unifieddatalibrary.String("WEATHER_REPORT-ID"),
			MilClass:             unifieddatalibrary.String("CIVILIAN"),
			NatAlliance:          unifieddatalibrary.Int(1),
			NatAlliance1:         unifieddatalibrary.Int(0),
			NumAmbulatory:        unifieddatalibrary.Int(1),
			NumAmbulatoryInjured: unifieddatalibrary.Int(2),
			NumNonAmbulatory:     unifieddatalibrary.Int(0),
			NumPersons:           unifieddatalibrary.Int(1),
			ObjectiveAreaInfo: unifieddatalibrary.PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfo{
				EnemyData: []unifieddatalibrary.PersonnelrecoveryNewBulkParamsBodyObjectiveAreaInfoEnemyData{{
					DirToEnemy:        unifieddatalibrary.String("NORTHWEST"),
					FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
					HlzRemarks:        unifieddatalibrary.String("Hot Landing Zone remarks."),
					HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
				}},
				OscCallSign: unifieddatalibrary.String("STARFOX"),
				OscFreq:     unifieddatalibrary.Float(12.55),
				PzDesc:      unifieddatalibrary.String("Near the lake."),
				PzLocation:  []float64{103.23, 30.445},
			},
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PickupAlt:        unifieddatalibrary.Float(30.1234),
			RecovID:          unifieddatalibrary.String("RECOV-ID"),
			RxFreq:           unifieddatalibrary.Float(5.5),
			SurvivorMessages: unifieddatalibrary.String("UNINJURED CANT MOVE HOSTILES NEARBY"),
			SurvivorRadio:    unifieddatalibrary.String("NO STATEMENT"),
			TermInd:          unifieddatalibrary.Bool(true),
			TextMsg:          unifieddatalibrary.String("Additional message from survivor."),
			TxFreq:           unifieddatalibrary.Float(5.5),
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

func TestPersonnelrecoveryFileNew(t *testing.T) {
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
	err := client.Personnelrecovery.FileNew(context.TODO(), unifieddatalibrary.PersonnelrecoveryFileNewParams{
		Body: []unifieddatalibrary.PersonnelrecoveryFileNewParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			MsgTime:               time.Now(),
			PickupLat:             75.1234,
			PickupLon:             175.1234,
			Source:                "Bluestaq",
			Type:                  "MEDICAL",
			ID:                    unifieddatalibrary.String("PERSONNEL_RECOVERY-ID"),
			AuthMethod:            unifieddatalibrary.String("PASSPORT"),
			AuthStatus:            unifieddatalibrary.String("NO STATEMENT"),
			BeaconInd:             unifieddatalibrary.Bool(false),
			CallSign:              unifieddatalibrary.String("BADGER"),
			CommEq1:               unifieddatalibrary.String("LL PHONE"),
			CommEq2:               unifieddatalibrary.String("LL PHONE"),
			CommEq3:               unifieddatalibrary.String("LL PHONE"),
			ExecutionInfo: unifieddatalibrary.PersonnelrecoveryFileNewParamsBodyExecutionInfo{
				Egress:      unifieddatalibrary.Float(66.53),
				EgressPoint: []float64{107.23, 30.455},
				EscortVehicle: unifieddatalibrary.PersonnelrecoveryFileNewParamsBodyExecutionInfoEscortVehicle{
					CallSign:    unifieddatalibrary.String("FALCO"),
					PrimaryFreq: unifieddatalibrary.Float(34.55),
					Strength:    unifieddatalibrary.Int(5),
					Type:        unifieddatalibrary.String("C17"),
				},
				Ingress:      unifieddatalibrary.Float(35.66),
				InitialPoint: []float64{103.23, 30.445},
				ObjStrategy:  unifieddatalibrary.String("Description of strategy plan."),
				RecoveryVehicle: unifieddatalibrary.PersonnelrecoveryFileNewParamsBodyExecutionInfoRecoveryVehicle{
					CallSign:    unifieddatalibrary.String("FALCO"),
					PrimaryFreq: unifieddatalibrary.Float(34.55),
					Strength:    unifieddatalibrary.Int(5),
					Type:        unifieddatalibrary.String("C17"),
				},
			},
			Identity:             unifieddatalibrary.String("NEUTRAL CIVILIAN"),
			IDWeatherReport:      unifieddatalibrary.String("WEATHER_REPORT-ID"),
			MilClass:             unifieddatalibrary.String("CIVILIAN"),
			NatAlliance:          unifieddatalibrary.Int(1),
			NatAlliance1:         unifieddatalibrary.Int(0),
			NumAmbulatory:        unifieddatalibrary.Int(1),
			NumAmbulatoryInjured: unifieddatalibrary.Int(2),
			NumNonAmbulatory:     unifieddatalibrary.Int(0),
			NumPersons:           unifieddatalibrary.Int(1),
			ObjectiveAreaInfo: unifieddatalibrary.PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfo{
				EnemyData: []unifieddatalibrary.PersonnelrecoveryFileNewParamsBodyObjectiveAreaInfoEnemyData{{
					DirToEnemy:        unifieddatalibrary.String("NORTHWEST"),
					FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
					HlzRemarks:        unifieddatalibrary.String("Hot Landing Zone remarks."),
					HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
				}},
				OscCallSign: unifieddatalibrary.String("STARFOX"),
				OscFreq:     unifieddatalibrary.Float(12.55),
				PzDesc:      unifieddatalibrary.String("Near the lake."),
				PzLocation:  []float64{103.23, 30.445},
			},
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PickupAlt:        unifieddatalibrary.Float(30.1234),
			RecovID:          unifieddatalibrary.String("RECOV-ID"),
			RxFreq:           unifieddatalibrary.Float(5.5),
			SurvivorMessages: unifieddatalibrary.String("UNINJURED CANT MOVE HOSTILES NEARBY"),
			SurvivorRadio:    unifieddatalibrary.String("NO STATEMENT"),
			TermInd:          unifieddatalibrary.Bool(true),
			TextMsg:          unifieddatalibrary.String("Additional message from survivor."),
			TxFreq:           unifieddatalibrary.Float(5.5),
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

func TestPersonnelrecoveryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Personnelrecovery.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.PersonnelrecoveryGetParams{
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

func TestPersonnelrecoveryQueryhelp(t *testing.T) {
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
	err := client.Personnelrecovery.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPersonnelrecoveryTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Personnelrecovery.Tuple(context.TODO(), unifieddatalibrary.PersonnelrecoveryTupleParams{
		Columns:     "columns",
		MsgTime:     time.Now(),
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
