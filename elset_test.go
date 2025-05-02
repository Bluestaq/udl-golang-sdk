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

func TestElsetNewWithOptionalParams(t *testing.T) {
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
	err := client.Elsets.New(context.TODO(), unifieddatalibrary.ElsetNewParams{
		ElsetIngest: unifieddatalibrary.ElsetIngestParam{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ElsetIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			Agom:                  unifieddatalibrary.Float(0.0126),
			Algorithm:             unifieddatalibrary.String("Example algorithm"),
			Apogee:                unifieddatalibrary.Float(1.1),
			ArgOfPerigee:          unifieddatalibrary.Float(1.1),
			BallisticCoeff:        unifieddatalibrary.Float(0.00815),
			BStar:                 unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("Example description"),
			Eccentricity:          unifieddatalibrary.Float(0.333),
			EphemType:             unifieddatalibrary.Int(1),
			IDElset:               unifieddatalibrary.String("ELSET-ID"),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Inclination:           unifieddatalibrary.Float(45.1),
			MeanAnomaly:           unifieddatalibrary.Float(179.1),
			MeanMotion:            unifieddatalibrary.Float(1.1),
			MeanMotionDDot:        unifieddatalibrary.Float(1.1),
			MeanMotionDot:         unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Perigee:               unifieddatalibrary.Float(1.1),
			Period:                unifieddatalibrary.Float(1.1),
			Raan:                  unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			RevNo:                 unifieddatalibrary.Int(111),
			SatNo:                 unifieddatalibrary.Int(12),
			SemiMajorAxis:         unifieddatalibrary.Float(1.1),
			SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
			SourcedDataTypes:      []string{"RADAR", "RF"},
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
		},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElsetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Elsets.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ElsetGetParams{
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

func TestElsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Elsets.List(context.TODO(), unifieddatalibrary.ElsetListParams{
		Epoch:       time.Now(),
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

func TestElsetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Elsets.Count(context.TODO(), unifieddatalibrary.ElsetCountParams{
		Epoch:       time.Now(),
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

func TestElsetNewBulkWithOptionalParams(t *testing.T) {
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
	err := client.Elsets.NewBulk(context.TODO(), unifieddatalibrary.ElsetNewBulkParams{
		Body: []unifieddatalibrary.ElsetIngestParam{{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ElsetIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			Agom:                  unifieddatalibrary.Float(0.0126),
			Algorithm:             unifieddatalibrary.String("Example algorithm"),
			Apogee:                unifieddatalibrary.Float(1.1),
			ArgOfPerigee:          unifieddatalibrary.Float(1.1),
			BallisticCoeff:        unifieddatalibrary.Float(0.00815),
			BStar:                 unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("Example description"),
			Eccentricity:          unifieddatalibrary.Float(0.333),
			EphemType:             unifieddatalibrary.Int(1),
			IDElset:               unifieddatalibrary.String("ELSET-ID"),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Inclination:           unifieddatalibrary.Float(45.1),
			MeanAnomaly:           unifieddatalibrary.Float(179.1),
			MeanMotion:            unifieddatalibrary.Float(1.1),
			MeanMotionDDot:        unifieddatalibrary.Float(1.1),
			MeanMotionDot:         unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Perigee:               unifieddatalibrary.Float(1.1),
			Period:                unifieddatalibrary.Float(1.1),
			Raan:                  unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			RevNo:                 unifieddatalibrary.Int(111),
			SatNo:                 unifieddatalibrary.Int(12),
			SemiMajorAxis:         unifieddatalibrary.Float(1.1),
			SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
			SourcedDataTypes:      []string{"RADAR", "RF"},
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
		}},
		DupeCheck: unifieddatalibrary.Bool(true),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElsetNewBulkFromTleWithOptionalParams(t *testing.T) {
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
	err := client.Elsets.NewBulkFromTle(context.TODO(), unifieddatalibrary.ElsetNewBulkFromTleParams{
		DataMode:       "dataMode",
		MakeCurrent:    true,
		Source:         "source",
		Body:           "body",
		AutoCreateSats: unifieddatalibrary.Bool(true),
		Control:        unifieddatalibrary.String("control"),
		Origin:         unifieddatalibrary.String("origin"),
		Tags:           unifieddatalibrary.String("tags"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElsetQueryCurrentElsetHelp(t *testing.T) {
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
	err := client.Elsets.QueryCurrentElsetHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElsetQueryhelp(t *testing.T) {
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
	err := client.Elsets.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElsetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Elsets.Tuple(context.TODO(), unifieddatalibrary.ElsetTupleParams{
		Columns:     "columns",
		Epoch:       time.Now(),
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

func TestElsetUnvalidatedPublish(t *testing.T) {
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
	err := client.Elsets.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ElsetUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ElsetIngestParam{{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ElsetIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			Agom:                  unifieddatalibrary.Float(0.0126),
			Algorithm:             unifieddatalibrary.String("Example algorithm"),
			Apogee:                unifieddatalibrary.Float(1.1),
			ArgOfPerigee:          unifieddatalibrary.Float(1.1),
			BallisticCoeff:        unifieddatalibrary.Float(0.00815),
			BStar:                 unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("Example description"),
			Eccentricity:          unifieddatalibrary.Float(0.333),
			EphemType:             unifieddatalibrary.Int(1),
			IDElset:               unifieddatalibrary.String("ELSET-ID"),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Inclination:           unifieddatalibrary.Float(45.1),
			MeanAnomaly:           unifieddatalibrary.Float(179.1),
			MeanMotion:            unifieddatalibrary.Float(1.1),
			MeanMotionDDot:        unifieddatalibrary.Float(1.1),
			MeanMotionDot:         unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Perigee:               unifieddatalibrary.Float(1.1),
			Period:                unifieddatalibrary.Float(1.1),
			Raan:                  unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			RevNo:                 unifieddatalibrary.Int(111),
			SatNo:                 unifieddatalibrary.Int(12),
			SemiMajorAxis:         unifieddatalibrary.Float(1.1),
			SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
			SourcedDataTypes:      []string{"RADAR", "RF"},
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
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
