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

func TestEopNewWithOptionalParams(t *testing.T) {
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
	err := client.Eop.New(context.TODO(), unifieddatalibrary.EopNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EopNewParamsDataModeTest,
		EopDate:               time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("EOP-ID"),
		DEpsilon:              unifieddatalibrary.Float(-0.917),
		DEpsilonB:             unifieddatalibrary.Float(-1.7),
		DEpsilonUnc:           unifieddatalibrary.Float(0.165),
		DPsi:                  unifieddatalibrary.Float(-10.437),
		DPsib:                 unifieddatalibrary.Float(-9.9),
		DPsiUnc:               unifieddatalibrary.Float(0.507),
		DX:                    unifieddatalibrary.Float(-0.086),
		DXb:                   unifieddatalibrary.Float(0.129),
		DXUnc:                 unifieddatalibrary.Float(0.202),
		DY:                    unifieddatalibrary.Float(0.13),
		DYb:                   unifieddatalibrary.Float(-0.653),
		DYUnc:                 unifieddatalibrary.Float(0.165),
		Lod:                   unifieddatalibrary.Float(1.8335),
		LodUnc:                unifieddatalibrary.Float(0.0201),
		NutationState:         unifieddatalibrary.EopNewParamsNutationStateI,
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PolarMotionState:      unifieddatalibrary.EopNewParamsPolarMotionStateI,
		PolarMotionX:          unifieddatalibrary.Float(0.182987),
		PolarMotionXb:         unifieddatalibrary.Float(0.1824),
		PolarMotionXUnc:       unifieddatalibrary.Float(0.000672),
		PolarMotionY:          unifieddatalibrary.Float(0.168775),
		PolarMotionYb:         unifieddatalibrary.Float(0.1679),
		PolarMotionYUnc:       unifieddatalibrary.Float(0.000345),
		PrecessionNutationStd: unifieddatalibrary.String("IAU1980"),
		RawFileUri:            unifieddatalibrary.String("Example URI"),
		Ut1Utc:                unifieddatalibrary.Float(-0.1251659),
		Ut1Utcb:               unifieddatalibrary.Float(-0.1253),
		Ut1UtcState:           unifieddatalibrary.EopNewParamsUt1UtcStateI,
		Ut1UtcUnc:             unifieddatalibrary.Float(0.0000207),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEopGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Eop.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EopGetParams{
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

func TestEopUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Eop.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.EopUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.EopUpdateParamsDataModeTest,
			EopDate:               time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("EOP-ID"),
			DEpsilon:              unifieddatalibrary.Float(-0.917),
			DEpsilonB:             unifieddatalibrary.Float(-1.7),
			DEpsilonUnc:           unifieddatalibrary.Float(0.165),
			DPsi:                  unifieddatalibrary.Float(-10.437),
			DPsib:                 unifieddatalibrary.Float(-9.9),
			DPsiUnc:               unifieddatalibrary.Float(0.507),
			DX:                    unifieddatalibrary.Float(-0.086),
			DXb:                   unifieddatalibrary.Float(0.129),
			DXUnc:                 unifieddatalibrary.Float(0.202),
			DY:                    unifieddatalibrary.Float(0.13),
			DYb:                   unifieddatalibrary.Float(-0.653),
			DYUnc:                 unifieddatalibrary.Float(0.165),
			Lod:                   unifieddatalibrary.Float(1.8335),
			LodUnc:                unifieddatalibrary.Float(0.0201),
			NutationState:         unifieddatalibrary.EopUpdateParamsNutationStateI,
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PolarMotionState:      unifieddatalibrary.EopUpdateParamsPolarMotionStateI,
			PolarMotionX:          unifieddatalibrary.Float(0.182987),
			PolarMotionXb:         unifieddatalibrary.Float(0.1824),
			PolarMotionXUnc:       unifieddatalibrary.Float(0.000672),
			PolarMotionY:          unifieddatalibrary.Float(0.168775),
			PolarMotionYb:         unifieddatalibrary.Float(0.1679),
			PolarMotionYUnc:       unifieddatalibrary.Float(0.000345),
			PrecessionNutationStd: unifieddatalibrary.String("IAU1980"),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			Ut1Utc:                unifieddatalibrary.Float(-0.1251659),
			Ut1Utcb:               unifieddatalibrary.Float(-0.1253),
			Ut1UtcState:           unifieddatalibrary.EopUpdateParamsUt1UtcStateI,
			Ut1UtcUnc:             unifieddatalibrary.Float(0.0000207),
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

func TestEopListWithOptionalParams(t *testing.T) {
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
	_, err := client.Eop.List(context.TODO(), unifieddatalibrary.EopListParams{
		EopDate:     time.Now(),
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

func TestEopDelete(t *testing.T) {
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
	err := client.Eop.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEopCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Eop.Count(context.TODO(), unifieddatalibrary.EopCountParams{
		EopDate:     time.Now(),
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

func TestEopListTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Eop.ListTuple(context.TODO(), unifieddatalibrary.EopListTupleParams{
		Columns:     "columns",
		EopDate:     time.Now(),
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

func TestEopQueryhelp(t *testing.T) {
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
	_, err := client.Eop.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
