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

func TestSgiNewWithOptionalParams(t *testing.T) {
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
	err := client.Sgi.New(context.TODO(), unifieddatalibrary.SgiNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SgiNewParamsDataModeTest,
		EffectiveDate:         time.Now(),
		SgiDate:               time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SGI-ID"),
		AnalyzerAttenuation:   unifieddatalibrary.Float(5.1),
		Ap:                    unifieddatalibrary.Float(1.23),
		ApDuration:            unifieddatalibrary.Int(3),
		CoeffDegree:           []int64{1, 2, 3},
		CoeffOrder:            []int64{1, 2, 3},
		Ctce:                  []float64{1.23, 342.3, 1.32},
		Ctci:                  []float64{1.23, 342.3, 1.32},
		Dst:                   unifieddatalibrary.Float(1.23),
		Dtc:                   unifieddatalibrary.Float(1.23),
		E10:                   unifieddatalibrary.Float(1.23),
		E54:                   unifieddatalibrary.Float(1.23),
		F10:                   unifieddatalibrary.Float(1.23),
		F10High:               unifieddatalibrary.Float(187.5),
		F10Low:                unifieddatalibrary.Float(185.5),
		F54:                   unifieddatalibrary.Float(1.23),
		F81:                   unifieddatalibrary.Float(1.23),
		Frequencies:           []float64{25, 25.125, 25.25, 25.375, 25.5, 25.625, 25.75, 25.875},
		Gamma:                 unifieddatalibrary.Int(25),
		IDSensor:              unifieddatalibrary.String("57c96c97-e076-48af-a068-73ee2cb37e65"),
		KIndex:                unifieddatalibrary.Int(1),
		Kp:                    unifieddatalibrary.Float(4.66),
		KpDuration:            unifieddatalibrary.Int(3),
		M10:                   unifieddatalibrary.Float(1.23),
		M54:                   unifieddatalibrary.Float(1.23),
		Mode:                  unifieddatalibrary.Int(1),
		NormFactor:            unifieddatalibrary.Float(2.12679e-7),
		ObservedBaseline:      []int64{15, 32, 25, 134, 0, 6, 19, 8},
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		Powers:                []float64{67.1, 65.2, 68.1, 74.3, 68.1, 96.4, 97.3, 68.1},
		Precedence:            unifieddatalibrary.SgiNewParamsPrecedenceR,
		RawFileUri:            unifieddatalibrary.String("rawFileURI"),
		RbDuration:            unifieddatalibrary.Int(24),
		RbIndex:               unifieddatalibrary.Float(1.02947164506),
		RbRegionCode:          unifieddatalibrary.Int(2),
		S10:                   unifieddatalibrary.Float(1.23),
		S54:                   unifieddatalibrary.Float(1.23),
		State:                 unifieddatalibrary.SgiNewParamsStateI,
		StationName:           unifieddatalibrary.String("Boulder"),
		Stce:                  []float64{1.23, 342.3, 1.32},
		Stci:                  []float64{1.23, 342.3, 1.32},
		SunspotNum:            unifieddatalibrary.Float(151.1),
		SunspotNumHigh:        unifieddatalibrary.Float(152.1),
		SunspotNumLow:         unifieddatalibrary.Float(150.1),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		Type:                  unifieddatalibrary.String("JBH09"),
		Y10:                   unifieddatalibrary.Float(1.23),
		Y54:                   unifieddatalibrary.Float(1.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sgi.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SgiUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SgiUpdateParamsDataModeTest,
			EffectiveDate:         time.Now(),
			SgiDate:               time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SGI-ID"),
			AnalyzerAttenuation:   unifieddatalibrary.Float(5.1),
			Ap:                    unifieddatalibrary.Float(1.23),
			ApDuration:            unifieddatalibrary.Int(3),
			CoeffDegree:           []int64{1, 2, 3},
			CoeffOrder:            []int64{1, 2, 3},
			Ctce:                  []float64{1.23, 342.3, 1.32},
			Ctci:                  []float64{1.23, 342.3, 1.32},
			Dst:                   unifieddatalibrary.Float(1.23),
			Dtc:                   unifieddatalibrary.Float(1.23),
			E10:                   unifieddatalibrary.Float(1.23),
			E54:                   unifieddatalibrary.Float(1.23),
			F10:                   unifieddatalibrary.Float(1.23),
			F10High:               unifieddatalibrary.Float(187.5),
			F10Low:                unifieddatalibrary.Float(185.5),
			F54:                   unifieddatalibrary.Float(1.23),
			F81:                   unifieddatalibrary.Float(1.23),
			Frequencies:           []float64{25, 25.125, 25.25, 25.375, 25.5, 25.625, 25.75, 25.875},
			Gamma:                 unifieddatalibrary.Int(25),
			IDSensor:              unifieddatalibrary.String("57c96c97-e076-48af-a068-73ee2cb37e65"),
			KIndex:                unifieddatalibrary.Int(1),
			Kp:                    unifieddatalibrary.Float(4.66),
			KpDuration:            unifieddatalibrary.Int(3),
			M10:                   unifieddatalibrary.Float(1.23),
			M54:                   unifieddatalibrary.Float(1.23),
			Mode:                  unifieddatalibrary.Int(1),
			NormFactor:            unifieddatalibrary.Float(2.12679e-7),
			ObservedBaseline:      []int64{15, 32, 25, 134, 0, 6, 19, 8},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			Powers:                []float64{67.1, 65.2, 68.1, 74.3, 68.1, 96.4, 97.3, 68.1},
			Precedence:            unifieddatalibrary.SgiUpdateParamsPrecedenceR,
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RbDuration:            unifieddatalibrary.Int(24),
			RbIndex:               unifieddatalibrary.Float(1.02947164506),
			RbRegionCode:          unifieddatalibrary.Int(2),
			S10:                   unifieddatalibrary.Float(1.23),
			S54:                   unifieddatalibrary.Float(1.23),
			State:                 unifieddatalibrary.SgiUpdateParamsStateI,
			StationName:           unifieddatalibrary.String("Boulder"),
			Stce:                  []float64{1.23, 342.3, 1.32},
			Stci:                  []float64{1.23, 342.3, 1.32},
			SunspotNum:            unifieddatalibrary.Float(151.1),
			SunspotNumHigh:        unifieddatalibrary.Float(152.1),
			SunspotNumLow:         unifieddatalibrary.Float(150.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Type:                  unifieddatalibrary.String("JBH09"),
			Y10:                   unifieddatalibrary.Float(1.23),
			Y54:                   unifieddatalibrary.Float(1.23),
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

func TestSgiListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sgi.List(context.TODO(), unifieddatalibrary.SgiListParams{
		EffectiveDate: unifieddatalibrary.Time(time.Now()),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
		SgiDate:       unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiDelete(t *testing.T) {
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
	err := client.Sgi.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sgi.Count(context.TODO(), unifieddatalibrary.SgiCountParams{
		EffectiveDate: unifieddatalibrary.Time(time.Now()),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
		SgiDate:       unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiNewBulk(t *testing.T) {
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
	err := client.Sgi.NewBulk(context.TODO(), unifieddatalibrary.SgiNewBulkParams{
		Body: []unifieddatalibrary.SgiNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EffectiveDate:         time.Now(),
			SgiDate:               time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SGI-ID"),
			AnalyzerAttenuation:   unifieddatalibrary.Float(5.1),
			Ap:                    unifieddatalibrary.Float(1.23),
			ApDuration:            unifieddatalibrary.Int(3),
			CoeffDegree:           []int64{1, 2, 3},
			CoeffOrder:            []int64{1, 2, 3},
			Ctce:                  []float64{1.23, 342.3, 1.32},
			Ctci:                  []float64{1.23, 342.3, 1.32},
			Dst:                   unifieddatalibrary.Float(1.23),
			Dtc:                   unifieddatalibrary.Float(1.23),
			E10:                   unifieddatalibrary.Float(1.23),
			E54:                   unifieddatalibrary.Float(1.23),
			F10:                   unifieddatalibrary.Float(1.23),
			F10High:               unifieddatalibrary.Float(187.5),
			F10Low:                unifieddatalibrary.Float(185.5),
			F54:                   unifieddatalibrary.Float(1.23),
			F81:                   unifieddatalibrary.Float(1.23),
			Frequencies:           []float64{25, 25.125, 25.25, 25.375, 25.5, 25.625, 25.75, 25.875},
			Gamma:                 unifieddatalibrary.Int(25),
			IDSensor:              unifieddatalibrary.String("57c96c97-e076-48af-a068-73ee2cb37e65"),
			KIndex:                unifieddatalibrary.Int(1),
			Kp:                    unifieddatalibrary.Float(4.66),
			KpDuration:            unifieddatalibrary.Int(3),
			M10:                   unifieddatalibrary.Float(1.23),
			M54:                   unifieddatalibrary.Float(1.23),
			Mode:                  unifieddatalibrary.Int(1),
			NormFactor:            unifieddatalibrary.Float(2.12679e-7),
			ObservedBaseline:      []int64{15, 32, 25, 134, 0, 6, 19, 8},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			Powers:                []float64{67.1, 65.2, 68.1, 74.3, 68.1, 96.4, 97.3, 68.1},
			Precedence:            "R",
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RbDuration:            unifieddatalibrary.Int(24),
			RbIndex:               unifieddatalibrary.Float(1.02947164506),
			RbRegionCode:          unifieddatalibrary.Int(2),
			S10:                   unifieddatalibrary.Float(1.23),
			S54:                   unifieddatalibrary.Float(1.23),
			State:                 "I",
			StationName:           unifieddatalibrary.String("Boulder"),
			Stce:                  []float64{1.23, 342.3, 1.32},
			Stci:                  []float64{1.23, 342.3, 1.32},
			SunspotNum:            unifieddatalibrary.Float(151.1),
			SunspotNumHigh:        unifieddatalibrary.Float(152.1),
			SunspotNumLow:         unifieddatalibrary.Float(150.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Type:                  unifieddatalibrary.String("JBH09"),
			Y10:                   unifieddatalibrary.Float(1.23),
			Y54:                   unifieddatalibrary.Float(1.23),
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

func TestSgiGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sgi.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SgiGetParams{
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

func TestSgiGetDataByEffectiveAsOfDateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sgi.GetDataByEffectiveAsOfDate(context.TODO(), unifieddatalibrary.SgiGetDataByEffectiveAsOfDateParams{
		EffectiveDate: unifieddatalibrary.Time(time.Now()),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
		SgiDate:       unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiQueryhelp(t *testing.T) {
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
	_, err := client.Sgi.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sgi.Tuple(context.TODO(), unifieddatalibrary.SgiTupleParams{
		Columns:       "columns",
		EffectiveDate: unifieddatalibrary.Time(time.Now()),
		FirstResult:   unifieddatalibrary.Int(0),
		MaxResults:    unifieddatalibrary.Int(0),
		SgiDate:       unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSgiUnvalidatedPublish(t *testing.T) {
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
	err := client.Sgi.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SgiUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SgiUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EffectiveDate:         time.Now(),
			SgiDate:               time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SGI-ID"),
			AnalyzerAttenuation:   unifieddatalibrary.Float(5.1),
			Ap:                    unifieddatalibrary.Float(1.23),
			ApDuration:            unifieddatalibrary.Int(3),
			CoeffDegree:           []int64{1, 2, 3},
			CoeffOrder:            []int64{1, 2, 3},
			Ctce:                  []float64{1.23, 342.3, 1.32},
			Ctci:                  []float64{1.23, 342.3, 1.32},
			Dst:                   unifieddatalibrary.Float(1.23),
			Dtc:                   unifieddatalibrary.Float(1.23),
			E10:                   unifieddatalibrary.Float(1.23),
			E54:                   unifieddatalibrary.Float(1.23),
			F10:                   unifieddatalibrary.Float(1.23),
			F10High:               unifieddatalibrary.Float(187.5),
			F10Low:                unifieddatalibrary.Float(185.5),
			F54:                   unifieddatalibrary.Float(1.23),
			F81:                   unifieddatalibrary.Float(1.23),
			Frequencies:           []float64{25, 25.125, 25.25, 25.375, 25.5, 25.625, 25.75, 25.875},
			Gamma:                 unifieddatalibrary.Int(25),
			IDSensor:              unifieddatalibrary.String("57c96c97-e076-48af-a068-73ee2cb37e65"),
			KIndex:                unifieddatalibrary.Int(1),
			Kp:                    unifieddatalibrary.Float(4.66),
			KpDuration:            unifieddatalibrary.Int(3),
			M10:                   unifieddatalibrary.Float(1.23),
			M54:                   unifieddatalibrary.Float(1.23),
			Mode:                  unifieddatalibrary.Int(1),
			NormFactor:            unifieddatalibrary.Float(2.12679e-7),
			ObservedBaseline:      []int64{15, 32, 25, 134, 0, 6, 19, 8},
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			Powers:                []float64{67.1, 65.2, 68.1, 74.3, 68.1, 96.4, 97.3, 68.1},
			Precedence:            "R",
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RbDuration:            unifieddatalibrary.Int(24),
			RbIndex:               unifieddatalibrary.Float(1.02947164506),
			RbRegionCode:          unifieddatalibrary.Int(2),
			S10:                   unifieddatalibrary.Float(1.23),
			S54:                   unifieddatalibrary.Float(1.23),
			State:                 "I",
			StationName:           unifieddatalibrary.String("Boulder"),
			Stce:                  []float64{1.23, 342.3, 1.32},
			Stci:                  []float64{1.23, 342.3, 1.32},
			SunspotNum:            unifieddatalibrary.Float(151.1),
			SunspotNumHigh:        unifieddatalibrary.Float(152.1),
			SunspotNumLow:         unifieddatalibrary.Float(150.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Type:                  unifieddatalibrary.String("JBH09"),
			Y10:                   unifieddatalibrary.Float(1.23),
			Y54:                   unifieddatalibrary.Float(1.23),
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
