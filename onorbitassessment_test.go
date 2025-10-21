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

func TestOnorbitassessmentNewWithOptionalParams(t *testing.T) {
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
	err := client.Onorbitassessment.New(context.TODO(), unifieddatalibrary.OnorbitassessmentNewParams{
		AssmtTime:             time.Now(),
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OnorbitassessmentNewParamsDataModeTest,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Assessment:            unifieddatalibrary.String("NOMINAL"),
		AssmtDetails:          unifieddatalibrary.String("This spacecraft appears to be in a stable, 3-axis controlled state"),
		AssmtLevel:            unifieddatalibrary.String("SIGNATURE"),
		AssmtRotPeriod:        unifieddatalibrary.Float(72.5),
		AssmtSubType:          unifieddatalibrary.String("STABLE"),
		AssmtURL:              unifieddatalibrary.String("https://unifieddatalibrary.com"),
		AutoAssmt:             unifieddatalibrary.Bool(false),
		CollectionURL:         unifieddatalibrary.String("https://unifieddatalibrary.com"),
		Components:            []string{"THRUSTER", "RWA-2"},
		IDOnOrbit:             unifieddatalibrary.String("25544"),
		IDSensor:              unifieddatalibrary.String("211"),
		ObDuration:            unifieddatalibrary.Float(1.75),
		ObTime:                unifieddatalibrary.Time(time.Now()),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ISS"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		SatNo:                 unifieddatalibrary.Int(25544),
		SigDataType:           unifieddatalibrary.String("PHOTOMETRY"),
		SrcIDs:                []string{"49cf9dcf-e97e-43ed-8e21-22e2bb0e3da6", "da779fc4-3a37-4caa-a629-289671bc96e8"},
		SrcTyps:               []string{"EO", "SKYIMAGE"},
		Tags:                  []string{"TAG1", "TAG2"},
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitassessmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitassessment.List(context.TODO(), unifieddatalibrary.OnorbitassessmentListParams{
		AssmtTime:   time.Now(),
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

func TestOnorbitassessmentCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitassessment.Count(context.TODO(), unifieddatalibrary.OnorbitassessmentCountParams{
		AssmtTime:   time.Now(),
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

func TestOnorbitassessmentNewBulk(t *testing.T) {
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
	err := client.Onorbitassessment.NewBulk(context.TODO(), unifieddatalibrary.OnorbitassessmentNewBulkParams{
		Body: []unifieddatalibrary.OnorbitassessmentNewBulkParamsBody{{
			AssmtTime:             time.Now(),
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Assessment:            unifieddatalibrary.String("NOMINAL"),
			AssmtDetails:          unifieddatalibrary.String("This spacecraft appears to be in a stable, 3-axis controlled state"),
			AssmtLevel:            unifieddatalibrary.String("SIGNATURE"),
			AssmtRotPeriod:        unifieddatalibrary.Float(72.5),
			AssmtSubType:          unifieddatalibrary.String("STABLE"),
			AssmtURL:              unifieddatalibrary.String("https://unifieddatalibrary.com"),
			AutoAssmt:             unifieddatalibrary.Bool(false),
			CollectionURL:         unifieddatalibrary.String("https://unifieddatalibrary.com"),
			Components:            []string{"THRUSTER", "RWA-2"},
			IDOnOrbit:             unifieddatalibrary.String("25544"),
			IDSensor:              unifieddatalibrary.String("211"),
			ObDuration:            unifieddatalibrary.Float(1.75),
			ObTime:                unifieddatalibrary.Time(time.Now()),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ISS"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			SatNo:                 unifieddatalibrary.Int(25544),
			SigDataType:           unifieddatalibrary.String("PHOTOMETRY"),
			SrcIDs:                []string{"49cf9dcf-e97e-43ed-8e21-22e2bb0e3da6", "da779fc4-3a37-4caa-a629-289671bc96e8"},
			SrcTyps:               []string{"EO", "SKYIMAGE"},
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
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

func TestOnorbitassessmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitassessment.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OnorbitassessmentGetParams{
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

func TestOnorbitassessmentQueryhelp(t *testing.T) {
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
	_, err := client.Onorbitassessment.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnorbitassessmentTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Onorbitassessment.Tuple(context.TODO(), unifieddatalibrary.OnorbitassessmentTupleParams{
		AssmtTime:   time.Now(),
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

func TestOnorbitassessmentUnvalidatedPublish(t *testing.T) {
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
	err := client.Onorbitassessment.UnvalidatedPublish(context.TODO(), unifieddatalibrary.OnorbitassessmentUnvalidatedPublishParams{
		Body: []unifieddatalibrary.OnorbitassessmentUnvalidatedPublishParamsBody{{
			AssmtTime:             time.Now(),
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Assessment:            unifieddatalibrary.String("NOMINAL"),
			AssmtDetails:          unifieddatalibrary.String("This spacecraft appears to be in a stable, 3-axis controlled state"),
			AssmtLevel:            unifieddatalibrary.String("SIGNATURE"),
			AssmtRotPeriod:        unifieddatalibrary.Float(72.5),
			AssmtSubType:          unifieddatalibrary.String("STABLE"),
			AssmtURL:              unifieddatalibrary.String("https://unifieddatalibrary.com"),
			AutoAssmt:             unifieddatalibrary.Bool(false),
			CollectionURL:         unifieddatalibrary.String("https://unifieddatalibrary.com"),
			Components:            []string{"THRUSTER", "RWA-2"},
			IDOnOrbit:             unifieddatalibrary.String("25544"),
			IDSensor:              unifieddatalibrary.String("211"),
			ObDuration:            unifieddatalibrary.Float(1.75),
			ObTime:                unifieddatalibrary.Time(time.Now()),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ISS"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			SatNo:                 unifieddatalibrary.Int(25544),
			SigDataType:           unifieddatalibrary.String("PHOTOMETRY"),
			SrcIDs:                []string{"49cf9dcf-e97e-43ed-8e21-22e2bb0e3da6", "da779fc4-3a37-4caa-a629-289671bc96e8"},
			SrcTyps:               []string{"EO", "SKYIMAGE"},
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
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
