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

func TestObjectOfInterestNewWithOptionalParams(t *testing.T) {
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
	err := client.ObjectOfInterest.New(context.TODO(), unifieddatalibrary.ObjectOfInterestNewParams{
		ClassificationMarking:  "U",
		DataMode:               unifieddatalibrary.ObjectOfInterestNewParamsDataModeTest,
		IDOnOrbit:              "REF-ONORBIT-ID",
		SensorTaskingStartTime: time.Now(),
		Source:                 "Bluestaq",
		StatusDate:             time.Now(),
		ID:                     unifieddatalibrary.String("OBJECTOFINTEREST-ID"),
		AffectedObjects:        []string{"AFFECTEDOBJECT1-ID", "AFFECTEDOBJECT2-ID"},
		Apogee:                 unifieddatalibrary.Float(123.4),
		ArgOfPerigee:           unifieddatalibrary.Float(123.4),
		BStar:                  unifieddatalibrary.Float(123.4),
		DeltaTs:                []float64{1.1, 2.2, 3.3},
		DeltaVs:                []float64{1.1, 2.2, 3.3},
		Description:            unifieddatalibrary.String("Example description"),
		Eccentricity:           unifieddatalibrary.Float(123.4),
		ElsetEpoch:             unifieddatalibrary.Time(time.Now()),
		Inclination:            unifieddatalibrary.Float(123.4),
		LastObTime:             unifieddatalibrary.Time(time.Now()),
		MeanAnomaly:            unifieddatalibrary.Float(123.4),
		MeanMotion:             unifieddatalibrary.Float(123.4),
		MeanMotionDDot:         unifieddatalibrary.Float(123.4),
		MeanMotionDot:          unifieddatalibrary.Float(123.4),
		MissedObTime:           unifieddatalibrary.Time(time.Now()),
		Name:                   unifieddatalibrary.String("Example_name"),
		Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Perigee:                unifieddatalibrary.Float(123.4),
		Period:                 unifieddatalibrary.Float(123.4),
		Priority:               unifieddatalibrary.Int(7),
		Raan:                   unifieddatalibrary.Float(123.4),
		RevNo:                  unifieddatalibrary.Int(123),
		SatNo:                  unifieddatalibrary.Int(12),
		SemiMajorAxis:          unifieddatalibrary.Float(123.4),
		SensorTaskingStopTime:  unifieddatalibrary.Time(time.Now()),
		Status:                 unifieddatalibrary.String("OPEN"),
		SvEpoch:                unifieddatalibrary.Time(time.Now()),
		X:                      unifieddatalibrary.Float(123.4),
		Xvel:                   unifieddatalibrary.Float(123.4),
		Y:                      unifieddatalibrary.Float(123.4),
		Yvel:                   unifieddatalibrary.Float(123.4),
		Z:                      unifieddatalibrary.Float(123.4),
		Zvel:                   unifieddatalibrary.Float(123.4),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectOfInterestUpdateWithOptionalParams(t *testing.T) {
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
	err := client.ObjectOfInterest.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.ObjectOfInterestUpdateParams{
			ClassificationMarking:  "U",
			DataMode:               unifieddatalibrary.ObjectOfInterestUpdateParamsDataModeTest,
			IDOnOrbit:              "REF-ONORBIT-ID",
			SensorTaskingStartTime: time.Now(),
			Source:                 "Bluestaq",
			StatusDate:             time.Now(),
			ID:                     unifieddatalibrary.String("OBJECTOFINTEREST-ID"),
			AffectedObjects:        []string{"AFFECTEDOBJECT1-ID", "AFFECTEDOBJECT2-ID"},
			Apogee:                 unifieddatalibrary.Float(123.4),
			ArgOfPerigee:           unifieddatalibrary.Float(123.4),
			BStar:                  unifieddatalibrary.Float(123.4),
			DeltaTs:                []float64{1.1, 2.2, 3.3},
			DeltaVs:                []float64{1.1, 2.2, 3.3},
			Description:            unifieddatalibrary.String("Example description"),
			Eccentricity:           unifieddatalibrary.Float(123.4),
			ElsetEpoch:             unifieddatalibrary.Time(time.Now()),
			Inclination:            unifieddatalibrary.Float(123.4),
			LastObTime:             unifieddatalibrary.Time(time.Now()),
			MeanAnomaly:            unifieddatalibrary.Float(123.4),
			MeanMotion:             unifieddatalibrary.Float(123.4),
			MeanMotionDDot:         unifieddatalibrary.Float(123.4),
			MeanMotionDot:          unifieddatalibrary.Float(123.4),
			MissedObTime:           unifieddatalibrary.Time(time.Now()),
			Name:                   unifieddatalibrary.String("Example_name"),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Perigee:                unifieddatalibrary.Float(123.4),
			Period:                 unifieddatalibrary.Float(123.4),
			Priority:               unifieddatalibrary.Int(7),
			Raan:                   unifieddatalibrary.Float(123.4),
			RevNo:                  unifieddatalibrary.Int(123),
			SatNo:                  unifieddatalibrary.Int(12),
			SemiMajorAxis:          unifieddatalibrary.Float(123.4),
			SensorTaskingStopTime:  unifieddatalibrary.Time(time.Now()),
			Status:                 unifieddatalibrary.String("OPEN"),
			SvEpoch:                unifieddatalibrary.Time(time.Now()),
			X:                      unifieddatalibrary.Float(123.4),
			Xvel:                   unifieddatalibrary.Float(123.4),
			Y:                      unifieddatalibrary.Float(123.4),
			Yvel:                   unifieddatalibrary.Float(123.4),
			Z:                      unifieddatalibrary.Float(123.4),
			Zvel:                   unifieddatalibrary.Float(123.4),
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

func TestObjectOfInterestListWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectOfInterest.List(context.TODO(), unifieddatalibrary.ObjectOfInterestListParams{
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

func TestObjectOfInterestDelete(t *testing.T) {
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
	err := client.ObjectOfInterest.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectOfInterestCountWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectOfInterest.Count(context.TODO(), unifieddatalibrary.ObjectOfInterestCountParams{
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

func TestObjectOfInterestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectOfInterest.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObjectOfInterestGetParams{
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

func TestObjectOfInterestQueryhelp(t *testing.T) {
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
	err := client.ObjectOfInterest.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectOfInterestTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectOfInterest.Tuple(context.TODO(), unifieddatalibrary.ObjectOfInterestTupleParams{
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
