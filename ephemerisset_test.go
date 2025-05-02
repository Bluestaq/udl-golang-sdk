// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestEphemerisSetNewWithOptionalParams(t *testing.T) {
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
	err := client.EphemerisSets.New(context.TODO(), unifieddatalibrary.EphemerisSetNewParams{
		Category:              "ANALYST",
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EphemerisSetNewParamsDataModeTest,
		NumPoints:             1,
		PointEndTime:          time.Now(),
		PointStartTime:        time.Now(),
		Source:                "Bluestaq",
		Type:                  "LAUNCH",
		ID:                    unifieddatalibrary.String("EPHEMERISSET-ID"),
		BDot:                  unifieddatalibrary.Float(1.1),
		CentBody:              unifieddatalibrary.String("Earth"),
		Comments:              unifieddatalibrary.String("Example notes"),
		CovReferenceFrame:     unifieddatalibrary.EphemerisSetNewParamsCovReferenceFrameJ2000,
		Description:           unifieddatalibrary.String("Example notes"),
		Descriptor:            unifieddatalibrary.String("Example descriptor"),
		DragModel:             unifieddatalibrary.String("JAC70"),
		Edr:                   unifieddatalibrary.Float(1.1),
		EphemerisList: []unifieddatalibrary.EphemerisSetNewParamsEphemerisList{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			Xpos:                  1.1,
			Xvel:                  1.1,
			Ypos:                  1.1,
			Yvel:                  1.1,
			Zpos:                  1.1,
			Zvel:                  1.1,
			ID:                    unifieddatalibrary.String("EPHEMERIS-ID"),
			Cov:                   []float64{1.1, 2.4, 3.8, 4.2, 5.5, 6},
			EsID:                  unifieddatalibrary.String("ES-ID"),
			IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Xaccel:                unifieddatalibrary.Float(1.1),
			Yaccel:                unifieddatalibrary.Float(1.1),
			Zaccel:                unifieddatalibrary.Float(1.1),
		}},
		Filename:            unifieddatalibrary.String("Example file name"),
		GeopotentialModel:   unifieddatalibrary.String("GEM-T3"),
		HasAccel:            unifieddatalibrary.Bool(false),
		HasCov:              unifieddatalibrary.Bool(false),
		HasMnvr:             unifieddatalibrary.Bool(false),
		IDManeuvers:         []string{"EXAMPLE_ID1", "EXAMPLE_ID2"},
		IDOnOrbit:           unifieddatalibrary.String("ONORBIT-ID"),
		IDStateVector:       unifieddatalibrary.String("STATEVECTOR-ID"),
		Integrator:          unifieddatalibrary.String("COWELL"),
		Interpolation:       unifieddatalibrary.String("LINEAR"),
		InterpolationDegree: unifieddatalibrary.Int(5),
		LunarSolar:          unifieddatalibrary.Bool(false),
		Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:        unifieddatalibrary.String("ORIGOBJECT-ID"),
		Pedigree:            unifieddatalibrary.String("PROPAGATED"),
		ReferenceFrame:      unifieddatalibrary.EphemerisSetNewParamsReferenceFrameJ2000,
		SatNo:               unifieddatalibrary.Int(2),
		SolidEarthTides:     unifieddatalibrary.Bool(false),
		StepSize:            unifieddatalibrary.Int(1),
		Tags:                []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TransactionID:       unifieddatalibrary.String("TRANSACTION-ID"),
		UsableEndTime:       unifieddatalibrary.Time(time.Now()),
		UsableStartTime:     unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEphemerisSetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EphemerisSets.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EphemerisSetGetParams{
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

func TestEphemerisSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.EphemerisSets.List(context.TODO(), unifieddatalibrary.EphemerisSetListParams{
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
		PointEndTime:   unifieddatalibrary.Time(time.Now()),
		PointStartTime: unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEphemerisSetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.EphemerisSets.Count(context.TODO(), unifieddatalibrary.EphemerisSetCountParams{
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
		PointEndTime:   unifieddatalibrary.Time(time.Now()),
		PointStartTime: unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEphemerisSetFileGetWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	resp, err := client.EphemerisSets.FileGet(
		context.TODO(),
		"id",
		unifieddatalibrary.EphemerisSetFileGetParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestEphemerisSetQueryhelp(t *testing.T) {
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
	err := client.EphemerisSets.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEphemerisSetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.EphemerisSets.Tuple(context.TODO(), unifieddatalibrary.EphemerisSetTupleParams{
		Columns:        "columns",
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
		PointEndTime:   unifieddatalibrary.Time(time.Now()),
		PointStartTime: unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
