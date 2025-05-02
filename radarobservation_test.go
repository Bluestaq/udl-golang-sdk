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

func TestRadarobservationNewWithOptionalParams(t *testing.T) {
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
	err := client.Radarobservation.New(context.TODO(), unifieddatalibrary.RadarobservationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RadarobservationNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("RADAROBSERVATION-ID"),
		Azimuth:               unifieddatalibrary.Float(45.23),
		AzimuthBias:           unifieddatalibrary.Float(45.23),
		AzimuthMeasured:       unifieddatalibrary.Bool(true),
		AzimuthRate:           unifieddatalibrary.Float(1.23),
		AzimuthUnc:            unifieddatalibrary.Float(45.23),
		Beam:                  unifieddatalibrary.Float(1.23),
		Declination:           unifieddatalibrary.Float(10.23),
		DeclinationMeasured:   unifieddatalibrary.Bool(true),
		Descriptor:            unifieddatalibrary.String("descriptor"),
		Doppler:               unifieddatalibrary.Float(10.23),
		DopplerUnc:            unifieddatalibrary.Float(1.23),
		Elevation:             unifieddatalibrary.Float(45.23),
		ElevationBias:         unifieddatalibrary.Float(1.23),
		ElevationMeasured:     unifieddatalibrary.Bool(true),
		ElevationRate:         unifieddatalibrary.Float(1.23),
		ElevationUnc:          unifieddatalibrary.Float(1.23),
		IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
		ObPosition:            unifieddatalibrary.String("FIRST"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		OrthogonalRcs:         unifieddatalibrary.Float(1.23),
		OrthogonalRcsUnc:      unifieddatalibrary.Float(10.23),
		Ra:                    unifieddatalibrary.Float(1.23),
		RaMeasured:            unifieddatalibrary.Bool(true),
		Range:                 unifieddatalibrary.Float(100.23),
		RangeAccel:            unifieddatalibrary.Float(10.23),
		RangeAccelUnc:         unifieddatalibrary.Float(1.23),
		RangeBias:             unifieddatalibrary.Float(1.23),
		RangeMeasured:         unifieddatalibrary.Bool(true),
		RangeRate:             unifieddatalibrary.Float(1.23),
		RangeRateMeasured:     unifieddatalibrary.Bool(true),
		RangeRateUnc:          unifieddatalibrary.Float(0.5),
		RangeUnc:              unifieddatalibrary.Float(1.23),
		RawFileUri:            unifieddatalibrary.String("rawFileURI"),
		Rcs:                   unifieddatalibrary.Float(100.23),
		RcsUnc:                unifieddatalibrary.Float(1.23),
		SatNo:                 unifieddatalibrary.Int(1),
		SenReferenceFrame:     unifieddatalibrary.RadarobservationNewParamsSenReferenceFrameJ2000,
		Senx:                  unifieddatalibrary.Float(45.23),
		Seny:                  unifieddatalibrary.Float(40.23),
		Senz:                  unifieddatalibrary.Float(35.23),
		Snr:                   unifieddatalibrary.Float(0.5),
		Tags:                  []string{"TAG1", "TAG2"},
		TaskID:                unifieddatalibrary.String("TASK-ID"),
		TimingBias:            unifieddatalibrary.Float(1.23),
		TrackID:               unifieddatalibrary.String("TRACK-ID"),
		TrackingState:         unifieddatalibrary.String("INIT ACQ"),
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		Uct:                   unifieddatalibrary.Bool(true),
		X:                     unifieddatalibrary.Float(50.23),
		Xvel:                  unifieddatalibrary.Float(1.23),
		Y:                     unifieddatalibrary.Float(50.23),
		Yvel:                  unifieddatalibrary.Float(5.23),
		Z:                     unifieddatalibrary.Float(50.23),
		Zvel:                  unifieddatalibrary.Float(5.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarobservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radarobservation.List(context.TODO(), unifieddatalibrary.RadarobservationListParams{
		ObTime:      time.Now(),
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

func TestRadarobservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Radarobservation.Count(context.TODO(), unifieddatalibrary.RadarobservationCountParams{
		ObTime:      time.Now(),
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

func TestRadarobservationNewBulk(t *testing.T) {
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
	err := client.Radarobservation.NewBulk(context.TODO(), unifieddatalibrary.RadarobservationNewBulkParams{
		Body: []unifieddatalibrary.RadarobservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("RADAROBSERVATION-ID"),
			Azimuth:               unifieddatalibrary.Float(45.23),
			AzimuthBias:           unifieddatalibrary.Float(45.23),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.23),
			AzimuthUnc:            unifieddatalibrary.Float(45.23),
			Beam:                  unifieddatalibrary.Float(1.23),
			Declination:           unifieddatalibrary.Float(10.23),
			DeclinationMeasured:   unifieddatalibrary.Bool(true),
			Descriptor:            unifieddatalibrary.String("descriptor"),
			Doppler:               unifieddatalibrary.Float(10.23),
			DopplerUnc:            unifieddatalibrary.Float(1.23),
			Elevation:             unifieddatalibrary.Float(45.23),
			ElevationBias:         unifieddatalibrary.Float(1.23),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(1.23),
			ElevationUnc:          unifieddatalibrary.Float(1.23),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			ObPosition:            unifieddatalibrary.String("FIRST"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			OrthogonalRcs:         unifieddatalibrary.Float(1.23),
			OrthogonalRcsUnc:      unifieddatalibrary.Float(10.23),
			Ra:                    unifieddatalibrary.Float(1.23),
			RaMeasured:            unifieddatalibrary.Bool(true),
			Range:                 unifieddatalibrary.Float(100.23),
			RangeAccel:            unifieddatalibrary.Float(10.23),
			RangeAccelUnc:         unifieddatalibrary.Float(1.23),
			RangeBias:             unifieddatalibrary.Float(1.23),
			RangeMeasured:         unifieddatalibrary.Bool(true),
			RangeRate:             unifieddatalibrary.Float(1.23),
			RangeRateMeasured:     unifieddatalibrary.Bool(true),
			RangeRateUnc:          unifieddatalibrary.Float(0.5),
			RangeUnc:              unifieddatalibrary.Float(1.23),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			Rcs:                   unifieddatalibrary.Float(100.23),
			RcsUnc:                unifieddatalibrary.Float(1.23),
			SatNo:                 unifieddatalibrary.Int(1),
			SenReferenceFrame:     "J2000",
			Senx:                  unifieddatalibrary.Float(45.23),
			Seny:                  unifieddatalibrary.Float(40.23),
			Senz:                  unifieddatalibrary.Float(35.23),
			Snr:                   unifieddatalibrary.Float(0.5),
			Tags:                  []string{"TAG1", "TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TimingBias:            unifieddatalibrary.Float(1.23),
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TrackingState:         unifieddatalibrary.String("INIT ACQ"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(true),
			X:                     unifieddatalibrary.Float(50.23),
			Xvel:                  unifieddatalibrary.Float(1.23),
			Y:                     unifieddatalibrary.Float(50.23),
			Yvel:                  unifieddatalibrary.Float(5.23),
			Z:                     unifieddatalibrary.Float(50.23),
			Zvel:                  unifieddatalibrary.Float(5.23),
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

func TestRadarobservationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radarobservation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RadarobservationGetParams{
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

func TestRadarobservationQueryhelp(t *testing.T) {
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
	err := client.Radarobservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarobservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Radarobservation.Tuple(context.TODO(), unifieddatalibrary.RadarobservationTupleParams{
		Columns:     "columns",
		ObTime:      time.Now(),
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

func TestRadarobservationUnvalidatedPublish(t *testing.T) {
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
	err := client.Radarobservation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.RadarobservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.RadarobservationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("RADAROBSERVATION-ID"),
			Azimuth:               unifieddatalibrary.Float(45.23),
			AzimuthBias:           unifieddatalibrary.Float(45.23),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.23),
			AzimuthUnc:            unifieddatalibrary.Float(45.23),
			Beam:                  unifieddatalibrary.Float(1.23),
			Declination:           unifieddatalibrary.Float(10.23),
			DeclinationMeasured:   unifieddatalibrary.Bool(true),
			Descriptor:            unifieddatalibrary.String("descriptor"),
			Doppler:               unifieddatalibrary.Float(10.23),
			DopplerUnc:            unifieddatalibrary.Float(1.23),
			Elevation:             unifieddatalibrary.Float(45.23),
			ElevationBias:         unifieddatalibrary.Float(1.23),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(1.23),
			ElevationUnc:          unifieddatalibrary.Float(1.23),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			ObPosition:            unifieddatalibrary.String("FIRST"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			OrthogonalRcs:         unifieddatalibrary.Float(1.23),
			OrthogonalRcsUnc:      unifieddatalibrary.Float(10.23),
			Ra:                    unifieddatalibrary.Float(1.23),
			RaMeasured:            unifieddatalibrary.Bool(true),
			Range:                 unifieddatalibrary.Float(100.23),
			RangeAccel:            unifieddatalibrary.Float(10.23),
			RangeAccelUnc:         unifieddatalibrary.Float(1.23),
			RangeBias:             unifieddatalibrary.Float(1.23),
			RangeMeasured:         unifieddatalibrary.Bool(true),
			RangeRate:             unifieddatalibrary.Float(1.23),
			RangeRateMeasured:     unifieddatalibrary.Bool(true),
			RangeRateUnc:          unifieddatalibrary.Float(0.5),
			RangeUnc:              unifieddatalibrary.Float(1.23),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			Rcs:                   unifieddatalibrary.Float(100.23),
			RcsUnc:                unifieddatalibrary.Float(1.23),
			SatNo:                 unifieddatalibrary.Int(1),
			SenReferenceFrame:     "J2000",
			Senx:                  unifieddatalibrary.Float(45.23),
			Seny:                  unifieddatalibrary.Float(40.23),
			Senz:                  unifieddatalibrary.Float(35.23),
			Snr:                   unifieddatalibrary.Float(0.5),
			Tags:                  []string{"TAG1", "TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TimingBias:            unifieddatalibrary.Float(1.23),
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TrackingState:         unifieddatalibrary.String("INIT ACQ"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(true),
			X:                     unifieddatalibrary.Float(50.23),
			Xvel:                  unifieddatalibrary.Float(1.23),
			Y:                     unifieddatalibrary.Float(50.23),
			Yvel:                  unifieddatalibrary.Float(5.23),
			Z:                     unifieddatalibrary.Float(50.23),
			Zvel:                  unifieddatalibrary.Float(5.23),
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
