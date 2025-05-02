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

func TestGnssobservationsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Gnssobservationset.List(context.TODO(), unifieddatalibrary.GnssobservationsetListParams{
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

func TestGnssobservationsetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Gnssobservationset.Count(context.TODO(), unifieddatalibrary.GnssobservationsetCountParams{
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

func TestGnssobservationsetNewBulk(t *testing.T) {
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
	err := client.Gnssobservationset.NewBulk(context.TODO(), unifieddatalibrary.GnssobservationsetNewBulkParams{
		Body: []unifieddatalibrary.GnssobservationsetNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("GNSSObSet-ID"),
			AgcState:              unifieddatalibrary.Int(20),
			Alt:                   unifieddatalibrary.Float(200),
			Boresight:             []float64{1.23, 3.23},
			EsID:                  unifieddatalibrary.String("ES-ID"),
			EventID:               unifieddatalibrary.String("2f2205c9-7bc2-4e1a-8416-2f80cc71f64b"),
			GDop:                  unifieddatalibrary.Float(0.33),
			GnssObservationList: []unifieddatalibrary.GnssobservationsetNewBulkParamsBodyGnssObservationList{{
				AgcState:       unifieddatalibrary.Int(20),
				GnssSatID:      unifieddatalibrary.String("GEJ"),
				Ob:             []float64{42.1, 1000, 0.9},
				ObsCodeSet:     []string{"S1C", "C1C", "C1D"},
				TrackingStatus: unifieddatalibrary.Int(0),
			}},
			HDop:             unifieddatalibrary.Float(0.03),
			Lat:              unifieddatalibrary.Float(32.021),
			Lon:              unifieddatalibrary.Float(125.123),
			MarkerType:       unifieddatalibrary.String("SPACEBORNE"),
			NavigationStatus: unifieddatalibrary.String("degraded"),
			ObsCodes:         []string{"ACL"},
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:     unifieddatalibrary.String("ORIGOBJECT-ID"),
			Outage:           unifieddatalibrary.Int(200),
			PDop:             unifieddatalibrary.Float(0.002),
			Quat:             []float64{0.03, 0.02, 0.01, 0.012},
			Receiver:         unifieddatalibrary.String("RECEIVER-ID"),
			SatNo:            unifieddatalibrary.Int(2),
			SatPosition:      []float64{1625.71954, 6782.15396, -1721.34267},
			SatVelocity:      []float64{2.03, 0.003, 0.12},
			SrcIDs:           []string{"SV_ID", "SV_ID"},
			SrcTyps:          []string{"SV", "SV"},
			Tags:             []string{"TAG1", "TAG2"},
			TDop:             unifieddatalibrary.Float(0.05),
			TrackingStatus:   unifieddatalibrary.Int(0),
			TransactionID:    unifieddatalibrary.String("TRANSACTION-ID"),
			VDop:             unifieddatalibrary.Float(0.03),
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

func TestGnssobservationsetQueryhelp(t *testing.T) {
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
	err := client.Gnssobservationset.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGnssobservationsetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Gnssobservationset.Tuple(context.TODO(), unifieddatalibrary.GnssobservationsetTupleParams{
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

func TestGnssobservationsetUnvalidatedPublish(t *testing.T) {
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
	err := client.Gnssobservationset.UnvalidatedPublish(context.TODO(), unifieddatalibrary.GnssobservationsetUnvalidatedPublishParams{
		Body: []unifieddatalibrary.GnssobservationsetUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("GNSSObSet-ID"),
			AgcState:              unifieddatalibrary.Int(20),
			Alt:                   unifieddatalibrary.Float(200),
			Boresight:             []float64{1.23, 3.23},
			EsID:                  unifieddatalibrary.String("ES-ID"),
			EventID:               unifieddatalibrary.String("2f2205c9-7bc2-4e1a-8416-2f80cc71f64b"),
			GDop:                  unifieddatalibrary.Float(0.33),
			GnssObservationList: []unifieddatalibrary.GnssobservationsetUnvalidatedPublishParamsBodyGnssObservationList{{
				AgcState:       unifieddatalibrary.Int(20),
				GnssSatID:      unifieddatalibrary.String("GEJ"),
				Ob:             []float64{42.1, 1000, 0.9},
				ObsCodeSet:     []string{"S1C", "C1C", "C1D"},
				TrackingStatus: unifieddatalibrary.Int(0),
			}},
			HDop:             unifieddatalibrary.Float(0.03),
			Lat:              unifieddatalibrary.Float(32.021),
			Lon:              unifieddatalibrary.Float(125.123),
			MarkerType:       unifieddatalibrary.String("SPACEBORNE"),
			NavigationStatus: unifieddatalibrary.String("degraded"),
			ObsCodes:         []string{"ACL"},
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:     unifieddatalibrary.String("ORIGOBJECT-ID"),
			Outage:           unifieddatalibrary.Int(200),
			PDop:             unifieddatalibrary.Float(0.002),
			Quat:             []float64{0.03, 0.02, 0.01, 0.012},
			Receiver:         unifieddatalibrary.String("RECEIVER-ID"),
			SatNo:            unifieddatalibrary.Int(2),
			SatPosition:      []float64{1625.71954, 6782.15396, -1721.34267},
			SatVelocity:      []float64{2.03, 0.003, 0.12},
			SrcIDs:           []string{"SV_ID", "SV_ID"},
			SrcTyps:          []string{"SV", "SV"},
			Tags:             []string{"TAG1", "TAG2"},
			TDop:             unifieddatalibrary.Float(0.05),
			TrackingStatus:   unifieddatalibrary.Int(0),
			TransactionID:    unifieddatalibrary.String("TRANSACTION-ID"),
			VDop:             unifieddatalibrary.Float(0.03),
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
