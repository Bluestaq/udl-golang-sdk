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

func TestMonoradarListWithOptionalParams(t *testing.T) {
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
	_, err := client.Monoradar.List(context.TODO(), unifieddatalibrary.MonoradarListParams{
		Ts:          time.Now(),
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

func TestMonoradarCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Monoradar.Count(context.TODO(), unifieddatalibrary.MonoradarCountParams{
		Ts:          time.Now(),
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

func TestMonoradarNewBulk(t *testing.T) {
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
	err := client.Monoradar.NewBulk(context.TODO(), unifieddatalibrary.MonoradarNewBulkParams{
		Body: []unifieddatalibrary.MonoradarNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Msgfmt:                "CAT48",
			Msgts:                 time.Now(),
			Msgtyp:                "BCN",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("MONORADAR-ID"),
			Acp:                   unifieddatalibrary.Int(100),
			Addr:                  unifieddatalibrary.String("ADDR-ID"),
			Af:                    unifieddatalibrary.Bool(true),
			Aims:                  unifieddatalibrary.Bool(true),
			Alt3d:                 unifieddatalibrary.Float(100.23),
			Artsqual:              unifieddatalibrary.String("QUALITY"),
			Az:                    unifieddatalibrary.Float(100.23),
			Azdelt:                unifieddatalibrary.Float(44.23),
			Bcnhits:               unifieddatalibrary.Int(12),
			Cartpos:               []float64{1.2, 2.2},
			Cdm:                   unifieddatalibrary.String("CDM"),
			Code7500:              unifieddatalibrary.Bool(false),
			Code7600:              unifieddatalibrary.Bool(false),
			Code7700:              unifieddatalibrary.Bool(false),
			Faa:                   unifieddatalibrary.Bool(true),
			Grndspd:               unifieddatalibrary.Float(30.23),
			Hdng:                  unifieddatalibrary.Float(30.23),
			IDSensor:              unifieddatalibrary.String("REF-SENSOR-ID"),
			M1:                    unifieddatalibrary.String("MISSION_CODE"),
			M1g:                   unifieddatalibrary.Bool(true),
			M1v:                   unifieddatalibrary.String("M1V"),
			M2:                    unifieddatalibrary.String("MILITARY_ID_CODE"),
			M2g:                   unifieddatalibrary.Bool(true),
			M2v:                   unifieddatalibrary.String("M2V"),
			M2xv:                  unifieddatalibrary.String("M2XV"),
			M3a:                   unifieddatalibrary.String("AIRCRAFT_ID"),
			M3ag:                  unifieddatalibrary.Bool(true),
			M3av:                  unifieddatalibrary.String("M3AV"),
			M3axv:                 unifieddatalibrary.String("M3AXV"),
			M4:                    unifieddatalibrary.String("ID_FRIEND"),
			M4d1d2:                unifieddatalibrary.String("STATUS"),
			M4v:                   unifieddatalibrary.String("M4V"),
			Mah:                   unifieddatalibrary.String("MAH"),
			Mc:                    unifieddatalibrary.Float(100.23),
			Mcg:                   unifieddatalibrary.Bool(true),
			Mcv:                   unifieddatalibrary.String("MCV"),
			Milemrgcy:             unifieddatalibrary.Bool(false),
			Mrgrpt:                unifieddatalibrary.Bool(true),
			Mscommb:               unifieddatalibrary.String("MSCOMMB"),
			Mti:                   unifieddatalibrary.Bool(true),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			Psrrl:                 unifieddatalibrary.Float(44.23),
			Rad:                   unifieddatalibrary.String("RAD"),
			Rng:                   unifieddatalibrary.Float(100.23),
			Rngdelt:               unifieddatalibrary.Float(44.23),
			Sac:                   unifieddatalibrary.Int(10),
			Senalt:                unifieddatalibrary.Float(100.23),
			Senlat:                unifieddatalibrary.Float(45.23),
			Senlon:                unifieddatalibrary.Float(45.23),
			Sic:                   unifieddatalibrary.Int(40),
			Spi:                   unifieddatalibrary.Bool(true),
			Ssrl:                  unifieddatalibrary.Float(44.23),
			Tags:                  []string{"TAG1", "TAG2"},
			Tgtconf:               unifieddatalibrary.String("CONFIDENCE"),
			Tgtcorr:               unifieddatalibrary.String("CORRELATION"),
			Tgtid:                 unifieddatalibrary.String("TGT-ID"),
			Tis:                   unifieddatalibrary.Float(0.4),
			Trkelig:               unifieddatalibrary.String("ELIGIBILITY"),
			Trknum:                unifieddatalibrary.Int(30),
			Tti:                   unifieddatalibrary.String("TTI"),
			Wectc:                 []string{"WARNING", "WARNING"},
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

func TestMonoradarQueryhelp(t *testing.T) {
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
	err := client.Monoradar.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonoradarTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Monoradar.Tuple(context.TODO(), unifieddatalibrary.MonoradarTupleParams{
		Columns:     "columns",
		Ts:          time.Now(),
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
