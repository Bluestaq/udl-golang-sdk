// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestStageNewWithOptionalParams(t *testing.T) {
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
	err := client.Stage.New(context.TODO(), unifieddatalibrary.StageNewParams{
		ClassificationMarking:    "U",
		DataMode:                 unifieddatalibrary.StageNewParamsDataModeTest,
		IDEngine:                 "ENGINE-ID",
		IDLaunchVehicle:          "LAUNCHVEHICLE-ID",
		Source:                   "Bluestaq",
		ID:                       unifieddatalibrary.String("STAGE-ID"),
		AvionicsNotes:            unifieddatalibrary.String("Sample Notes"),
		BurnTime:                 unifieddatalibrary.Float(256.3),
		ControlThruster1:         unifieddatalibrary.String("controlThruster1"),
		ControlThruster2:         unifieddatalibrary.String("controlThruster2"),
		Diameter:                 unifieddatalibrary.Float(3.95),
		Length:                   unifieddatalibrary.Float(25.13),
		MainEngineThrustSeaLevel: unifieddatalibrary.Float(733.4),
		MainEngineThrustVacuum:   unifieddatalibrary.Float(733.4),
		ManufacturerOrgID:        unifieddatalibrary.String("5feed5d7-d131-57e5-a3fd-acc173bca736"),
		Mass:                     unifieddatalibrary.Float(9956.1),
		Notes:                    unifieddatalibrary.String("Sample Notes"),
		NumBurns:                 unifieddatalibrary.Int(1),
		NumControlThruster1:      unifieddatalibrary.Int(1),
		NumControlThruster2:      unifieddatalibrary.Int(1),
		NumEngines:               unifieddatalibrary.Int(1),
		NumStageElements:         unifieddatalibrary.Int(2),
		NumVernier:               unifieddatalibrary.Int(3),
		Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PhotoURLs:                []string{"photoURL"},
		Restartable:              unifieddatalibrary.Bool(true),
		Reusable:                 unifieddatalibrary.Bool(true),
		StageNumber:              unifieddatalibrary.Int(2),
		Tags:                     []string{"TAG1", "TAG2"},
		ThrustSeaLevel:           unifieddatalibrary.Float(733.4),
		ThrustVacuum:             unifieddatalibrary.Float(733.4),
		Type:                     unifieddatalibrary.String("Electrostatic Ion"),
		Vernier:                  unifieddatalibrary.String("vernier"),
		VernierBurnTime:          unifieddatalibrary.Float(1.1),
		VernierNumBurns:          unifieddatalibrary.Int(4),
		VernierThrustSeaLevel:    unifieddatalibrary.Float(4.1),
		VernierThrustVacuum:      unifieddatalibrary.Float(3.2),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStageUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Stage.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.StageUpdateParams{
			ClassificationMarking:    "U",
			DataMode:                 unifieddatalibrary.StageUpdateParamsDataModeTest,
			IDEngine:                 "ENGINE-ID",
			IDLaunchVehicle:          "LAUNCHVEHICLE-ID",
			Source:                   "Bluestaq",
			ID:                       unifieddatalibrary.String("STAGE-ID"),
			AvionicsNotes:            unifieddatalibrary.String("Sample Notes"),
			BurnTime:                 unifieddatalibrary.Float(256.3),
			ControlThruster1:         unifieddatalibrary.String("controlThruster1"),
			ControlThruster2:         unifieddatalibrary.String("controlThruster2"),
			Diameter:                 unifieddatalibrary.Float(3.95),
			Length:                   unifieddatalibrary.Float(25.13),
			MainEngineThrustSeaLevel: unifieddatalibrary.Float(733.4),
			MainEngineThrustVacuum:   unifieddatalibrary.Float(733.4),
			ManufacturerOrgID:        unifieddatalibrary.String("5feed5d7-d131-57e5-a3fd-acc173bca736"),
			Mass:                     unifieddatalibrary.Float(9956.1),
			Notes:                    unifieddatalibrary.String("Sample Notes"),
			NumBurns:                 unifieddatalibrary.Int(1),
			NumControlThruster1:      unifieddatalibrary.Int(1),
			NumControlThruster2:      unifieddatalibrary.Int(1),
			NumEngines:               unifieddatalibrary.Int(1),
			NumStageElements:         unifieddatalibrary.Int(2),
			NumVernier:               unifieddatalibrary.Int(3),
			Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PhotoURLs:                []string{"photoURL"},
			Restartable:              unifieddatalibrary.Bool(true),
			Reusable:                 unifieddatalibrary.Bool(true),
			StageNumber:              unifieddatalibrary.Int(2),
			Tags:                     []string{"TAG1", "TAG2"},
			ThrustSeaLevel:           unifieddatalibrary.Float(733.4),
			ThrustVacuum:             unifieddatalibrary.Float(733.4),
			Type:                     unifieddatalibrary.String("Electrostatic Ion"),
			Vernier:                  unifieddatalibrary.String("vernier"),
			VernierBurnTime:          unifieddatalibrary.Float(1.1),
			VernierNumBurns:          unifieddatalibrary.Int(4),
			VernierThrustSeaLevel:    unifieddatalibrary.Float(4.1),
			VernierThrustVacuum:      unifieddatalibrary.Float(3.2),
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

func TestStageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Stage.List(context.TODO(), unifieddatalibrary.StageListParams{
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

func TestStageDelete(t *testing.T) {
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
	err := client.Stage.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStageCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Stage.Count(context.TODO(), unifieddatalibrary.StageCountParams{
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

func TestStageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Stage.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.StageGetParams{
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

func TestStageQueryhelp(t *testing.T) {
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
	_, err := client.Stage.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStageTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Stage.Tuple(context.TODO(), unifieddatalibrary.StageTupleParams{
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
