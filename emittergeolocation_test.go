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

func TestEmitterGeolocationNewWithOptionalParams(t *testing.T) {
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
	err := client.EmitterGeolocation.New(context.TODO(), unifieddatalibrary.EmitterGeolocationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EmitterGeolocationNewParamsDataModeTest,
		SignalOfInterestType:  "RF",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
		AlgVersion:            unifieddatalibrary.String("v1.0-3-gps_nb_3ball"),
		Andims:                unifieddatalibrary.Int(3),
		Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Asrid:                 unifieddatalibrary.Int(3),
		Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Atype:                 unifieddatalibrary.String("MultiPolygon"),
		CenterFreq:            unifieddatalibrary.Float(1575.42),
		Cluster:               unifieddatalibrary.String("CONSTELLATION1-F"),
		ConfArea:              unifieddatalibrary.Float(81577480.056),
		Constellation:         unifieddatalibrary.String("HawkEye360"),
		CreatedTs:             unifieddatalibrary.Time(time.Now()),
		DetectAlt:             unifieddatalibrary.Float(123.456),
		DetectLat:             unifieddatalibrary.Float(41.172),
		DetectLon:             unifieddatalibrary.Float(37.019),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		ErrEllp:               []float64{1.23, 2.34, 3.45},
		ExternalID:            unifieddatalibrary.String("780180925"),
		IDRfEmitter:           unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		IDSensor:              unifieddatalibrary.String("OCULUSA"),
		MaxFreq:               unifieddatalibrary.Float(1575.42),
		MinFreq:               unifieddatalibrary.Float(1575.42),
		NumBursts:             unifieddatalibrary.Int(17),
		OrderID:               unifieddatalibrary.String("155240"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigRfEmitterID:       unifieddatalibrary.String("12345678"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		PassGroupID:           unifieddatalibrary.String("80fd25a8-8b41-448d-888a-91c9dfcd940b"),
		ReceivedTs:            unifieddatalibrary.Time(time.Now()),
		SatNo:                 unifieddatalibrary.Int(101),
		SignalOfInterest:      unifieddatalibrary.String("GPS"),
		Tags:                  []string{"TAG1", "TAG2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmitterGeolocationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EmitterGeolocation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EmitterGeolocationGetParams{
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

func TestEmitterGeolocationDelete(t *testing.T) {
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
	err := client.EmitterGeolocation.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmitterGeolocationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.EmitterGeolocation.Count(context.TODO(), unifieddatalibrary.EmitterGeolocationCountParams{
		StartTime:   time.Now(),
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

func TestEmitterGeolocationNewBulk(t *testing.T) {
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
	err := client.EmitterGeolocation.NewBulk(context.TODO(), unifieddatalibrary.EmitterGeolocationNewBulkParams{
		Body: []unifieddatalibrary.EmitterGeolocationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			SignalOfInterestType:  "RF",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			AlgVersion:            unifieddatalibrary.String("v1.0-3-gps_nb_3ball"),
			Andims:                unifieddatalibrary.Int(3),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(3),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("MultiPolygon"),
			CenterFreq:            unifieddatalibrary.Float(1575.42),
			Cluster:               unifieddatalibrary.String("CONSTELLATION1-F"),
			ConfArea:              unifieddatalibrary.Float(81577480.056),
			Constellation:         unifieddatalibrary.String("HawkEye360"),
			CreatedTs:             unifieddatalibrary.Time(time.Now()),
			DetectAlt:             unifieddatalibrary.Float(123.456),
			DetectLat:             unifieddatalibrary.Float(41.172),
			DetectLon:             unifieddatalibrary.Float(37.019),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ErrEllp:               []float64{1.23, 2.34, 3.45},
			ExternalID:            unifieddatalibrary.String("780180925"),
			IDRfEmitter:           unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			IDSensor:              unifieddatalibrary.String("OCULUSA"),
			MaxFreq:               unifieddatalibrary.Float(1575.42),
			MinFreq:               unifieddatalibrary.Float(1575.42),
			NumBursts:             unifieddatalibrary.Int(17),
			OrderID:               unifieddatalibrary.String("155240"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigRfEmitterID:       unifieddatalibrary.String("12345678"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			PassGroupID:           unifieddatalibrary.String("80fd25a8-8b41-448d-888a-91c9dfcd940b"),
			ReceivedTs:            unifieddatalibrary.Time(time.Now()),
			SatNo:                 unifieddatalibrary.Int(101),
			SignalOfInterest:      unifieddatalibrary.String("GPS"),
			Tags:                  []string{"TAG1", "TAG2"},
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

func TestEmitterGeolocationQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.EmitterGeolocation.Query(context.TODO(), unifieddatalibrary.EmitterGeolocationQueryParams{
		StartTime:   time.Now(),
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

func TestEmitterGeolocationQueryHelp(t *testing.T) {
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
	err := client.EmitterGeolocation.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmitterGeolocationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.EmitterGeolocation.Tuple(context.TODO(), unifieddatalibrary.EmitterGeolocationTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestEmitterGeolocationUnvalidatedPublish(t *testing.T) {
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
	err := client.EmitterGeolocation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EmitterGeolocationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EmitterGeolocationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			SignalOfInterestType:  "RF",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			AlgVersion:            unifieddatalibrary.String("v1.0-3-gps_nb_3ball"),
			Andims:                unifieddatalibrary.Int(3),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(3),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("MultiPolygon"),
			CenterFreq:            unifieddatalibrary.Float(1575.42),
			Cluster:               unifieddatalibrary.String("CONSTELLATION1-F"),
			ConfArea:              unifieddatalibrary.Float(81577480.056),
			Constellation:         unifieddatalibrary.String("HawkEye360"),
			CreatedTs:             unifieddatalibrary.Time(time.Now()),
			DetectAlt:             unifieddatalibrary.Float(123.456),
			DetectLat:             unifieddatalibrary.Float(41.172),
			DetectLon:             unifieddatalibrary.Float(37.019),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ErrEllp:               []float64{1.23, 2.34, 3.45},
			ExternalID:            unifieddatalibrary.String("780180925"),
			IDRfEmitter:           unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			IDSensor:              unifieddatalibrary.String("OCULUSA"),
			MaxFreq:               unifieddatalibrary.Float(1575.42),
			MinFreq:               unifieddatalibrary.Float(1575.42),
			NumBursts:             unifieddatalibrary.Int(17),
			OrderID:               unifieddatalibrary.String("155240"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigRfEmitterID:       unifieddatalibrary.String("12345678"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			PassGroupID:           unifieddatalibrary.String("80fd25a8-8b41-448d-888a-91c9dfcd940b"),
			ReceivedTs:            unifieddatalibrary.Time(time.Now()),
			SatNo:                 unifieddatalibrary.Int(101),
			SignalOfInterest:      unifieddatalibrary.String("GPS"),
			Tags:                  []string{"TAG1", "TAG2"},
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
