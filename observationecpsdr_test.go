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

func TestObservationEcpsdrNewWithOptionalParams(t *testing.T) {
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
	err := client.Observations.Ecpsdr.New(context.TODO(), unifieddatalibrary.ObservationEcpsdrNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ObservationEcpsdrNewParamsDataModeTest,
		MsgTime:               time.Now(),
		Source:                "Bluestaq",
		Type:                  "STANDARD",
		ID:                    unifieddatalibrary.String("ECPSDR-ID"),
		Asl5VCurrMon:          unifieddatalibrary.Int(12),
		CdsPlateVMon:          unifieddatalibrary.Int(12),
		CdsRefVMon:            unifieddatalibrary.Int(12),
		CdsThreshold:          unifieddatalibrary.Int(12),
		CdsThrottle:           unifieddatalibrary.Int(12),
		Checksum:              unifieddatalibrary.Int(12),
		DosBias:               unifieddatalibrary.Int(12),
		Dsl5VCurrMon:          unifieddatalibrary.Int(12),
		EsdTrigCountH:         unifieddatalibrary.Int(12),
		EsdTrigCountL:         unifieddatalibrary.Int(12),
		HiLetL:                unifieddatalibrary.Int(2),
		HiLetM:                unifieddatalibrary.Int(2),
		IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
		LowLetL:               unifieddatalibrary.Int(2),
		LowLetM:               unifieddatalibrary.Int(2),
		MedLet1L:              unifieddatalibrary.Int(2),
		MedLet1M:              unifieddatalibrary.Int(2),
		MedLet2L:              unifieddatalibrary.Int(2),
		MedLet2M:              unifieddatalibrary.Int(2),
		MedLet3L:              unifieddatalibrary.Int(2),
		MedLet3M:              unifieddatalibrary.Int(2),
		MedLet4L:              unifieddatalibrary.Int(2),
		MedLet4M:              unifieddatalibrary.Int(2),
		MpTemp:                unifieddatalibrary.Int(12),
		ObTime:                unifieddatalibrary.Time(time.Now()),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		Pd1SigLev:             unifieddatalibrary.Int(12),
		Pd2SigLev:             unifieddatalibrary.Int(12),
		PsTempMon:             unifieddatalibrary.Int(12),
		Retransmit:            unifieddatalibrary.Bool(true),
		SatNo:                 unifieddatalibrary.Int(101),
		SenMode:               unifieddatalibrary.String("TEST"),
		SurfDosChargeH:        unifieddatalibrary.Int(12),
		SurfDosChargeL:        unifieddatalibrary.Int(12),
		SurfDosH:              unifieddatalibrary.Int(12),
		SurfDosL:              unifieddatalibrary.Int(12),
		SurfDosM:              unifieddatalibrary.Int(12),
		SurfDosStat:           unifieddatalibrary.Int(2),
		TransientData:         []int64{1, 2, 3},
		VRef:                  unifieddatalibrary.Int(12),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationEcpsdrGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Ecpsdr.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObservationEcpsdrGetParams{
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

func TestObservationEcpsdrListWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Ecpsdr.List(context.TODO(), unifieddatalibrary.ObservationEcpsdrListParams{
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

func TestObservationEcpsdrCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Ecpsdr.Count(context.TODO(), unifieddatalibrary.ObservationEcpsdrCountParams{
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

func TestObservationEcpsdrNewBulk(t *testing.T) {
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
	err := client.Observations.Ecpsdr.NewBulk(context.TODO(), unifieddatalibrary.ObservationEcpsdrNewBulkParams{
		Body: []unifieddatalibrary.ObservationEcpsdrNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			MsgTime:               time.Now(),
			Source:                "Bluestaq",
			Type:                  "STANDARD",
			ID:                    unifieddatalibrary.String("ECPSDR-ID"),
			Asl5VCurrMon:          unifieddatalibrary.Int(12),
			CdsPlateVMon:          unifieddatalibrary.Int(12),
			CdsRefVMon:            unifieddatalibrary.Int(12),
			CdsThreshold:          unifieddatalibrary.Int(12),
			CdsThrottle:           unifieddatalibrary.Int(12),
			Checksum:              unifieddatalibrary.Int(12),
			DosBias:               unifieddatalibrary.Int(12),
			Dsl5VCurrMon:          unifieddatalibrary.Int(12),
			EsdTrigCountH:         unifieddatalibrary.Int(12),
			EsdTrigCountL:         unifieddatalibrary.Int(12),
			HiLetL:                unifieddatalibrary.Int(2),
			HiLetM:                unifieddatalibrary.Int(2),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			LowLetL:               unifieddatalibrary.Int(2),
			LowLetM:               unifieddatalibrary.Int(2),
			MedLet1L:              unifieddatalibrary.Int(2),
			MedLet1M:              unifieddatalibrary.Int(2),
			MedLet2L:              unifieddatalibrary.Int(2),
			MedLet2M:              unifieddatalibrary.Int(2),
			MedLet3L:              unifieddatalibrary.Int(2),
			MedLet3M:              unifieddatalibrary.Int(2),
			MedLet4L:              unifieddatalibrary.Int(2),
			MedLet4M:              unifieddatalibrary.Int(2),
			MpTemp:                unifieddatalibrary.Int(12),
			ObTime:                unifieddatalibrary.Time(time.Now()),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			Pd1SigLev:             unifieddatalibrary.Int(12),
			Pd2SigLev:             unifieddatalibrary.Int(12),
			PsTempMon:             unifieddatalibrary.Int(12),
			Retransmit:            unifieddatalibrary.Bool(true),
			SatNo:                 unifieddatalibrary.Int(101),
			SenMode:               unifieddatalibrary.String("TEST"),
			SurfDosChargeH:        unifieddatalibrary.Int(12),
			SurfDosChargeL:        unifieddatalibrary.Int(12),
			SurfDosH:              unifieddatalibrary.Int(12),
			SurfDosL:              unifieddatalibrary.Int(12),
			SurfDosM:              unifieddatalibrary.Int(12),
			SurfDosStat:           unifieddatalibrary.Int(2),
			TransientData:         []int64{1, 2, 3},
			VRef:                  unifieddatalibrary.Int(12),
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

func TestObservationEcpsdrQueryHelp(t *testing.T) {
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
	_, err := client.Observations.Ecpsdr.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationEcpsdrTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.Ecpsdr.Tuple(context.TODO(), unifieddatalibrary.ObservationEcpsdrTupleParams{
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
