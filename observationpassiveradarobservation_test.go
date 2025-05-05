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

func TestObservationPassiveRadarObservationNewWithOptionalParams(t *testing.T) {
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
	err := client.Observations.PassiveRadarObservation.New(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ObservationPassiveRadarObservationNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("bdcacfb0-3c47-4bd0-9d6c-9fa7d2c4fbb0"),
		Accel:                 unifieddatalibrary.Float(1.23),
		AccelUnc:              unifieddatalibrary.Float(0.1),
		Alt:                   unifieddatalibrary.Float(478.056378),
		Azimuth:               unifieddatalibrary.Float(134.5),
		AzimuthBias:           unifieddatalibrary.Float(0.123),
		AzimuthRate:           unifieddatalibrary.Float(0.5),
		AzimuthUnc:            unifieddatalibrary.Float(0.5),
		BistaticRange:         unifieddatalibrary.Float(754.8212),
		BistaticRangeAccel:    unifieddatalibrary.Float(1.23),
		BistaticRangeAccelUnc: unifieddatalibrary.Float(0.1),
		BistaticRangeBias:     unifieddatalibrary.Float(2.34),
		BistaticRangeRate:     unifieddatalibrary.Float(-0.30222),
		BistaticRangeRateUnc:  unifieddatalibrary.Float(0.123),
		BistaticRangeUnc:      unifieddatalibrary.Float(5.1),
		Coning:                unifieddatalibrary.Float(60.1),
		ConingUnc:             unifieddatalibrary.Float(0.5),
		Declination:           unifieddatalibrary.Float(10.23),
		Delay:                 unifieddatalibrary.Float(0.00505820232809312),
		DelayBias:             unifieddatalibrary.Float(0.00000123),
		DelayUnc:              unifieddatalibrary.Float(0.0000031),
		Descriptor:            unifieddatalibrary.String("Descriptor"),
		Doppler:               unifieddatalibrary.Float(-101.781641000597),
		DopplerUnc:            unifieddatalibrary.Float(0.2),
		Elevation:             unifieddatalibrary.Float(76.1),
		ElevationBias:         unifieddatalibrary.Float(0.123),
		ElevationRate:         unifieddatalibrary.Float(0.5),
		ElevationUnc:          unifieddatalibrary.Float(0.5),
		ExtObservationID:      unifieddatalibrary.String("26892"),
		IDRfEmitter:           unifieddatalibrary.String("RED_CLIFFS_3ABCRN"),
		IDSensor:              unifieddatalibrary.String("OCULUSA"),
		IDSensorRefReceiver:   unifieddatalibrary.String("OculusRef1"),
		Lat:                   unifieddatalibrary.Float(-35.1181763996856),
		Lon:                   unifieddatalibrary.Float(139.613567052763),
		ObPosition:            unifieddatalibrary.String("FIRST"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		OrthogonalRcs:         unifieddatalibrary.Float(10.23),
		OrthogonalRcsUnc:      unifieddatalibrary.Float(1.23),
		Ra:                    unifieddatalibrary.Float(1.23),
		Rcs:                   unifieddatalibrary.Float(100.23),
		RcsUnc:                unifieddatalibrary.Float(1.23),
		SatNo:                 unifieddatalibrary.Int(40699),
		Snr:                   unifieddatalibrary.Float(17.292053),
		Tags:                  []string{"TAG1", "TAG2"},
		TaskID:                unifieddatalibrary.String("TASK-ID"),
		TimingBias:            unifieddatalibrary.Float(1.23),
		Tof:                   unifieddatalibrary.Float(0.00592856674135648),
		TofBias:               unifieddatalibrary.Float(0.00000123),
		TofUnc:                unifieddatalibrary.Float(0.0000031),
		TrackID:               unifieddatalibrary.String("12212"),
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		Uct:                   unifieddatalibrary.Bool(false),
		Xvel:                  unifieddatalibrary.Float(1.23),
		Yvel:                  unifieddatalibrary.Float(3.21),
		Zvel:                  unifieddatalibrary.Float(3.12),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationPassiveRadarObservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.PassiveRadarObservation.List(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationListParams{
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

func TestObservationPassiveRadarObservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.PassiveRadarObservation.Count(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationCountParams{
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

func TestObservationPassiveRadarObservationNewBulk(t *testing.T) {
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
	err := client.Observations.PassiveRadarObservation.NewBulk(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationNewBulkParams{
		Body: []unifieddatalibrary.ObservationPassiveRadarObservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("bdcacfb0-3c47-4bd0-9d6c-9fa7d2c4fbb0"),
			Accel:                 unifieddatalibrary.Float(1.23),
			AccelUnc:              unifieddatalibrary.Float(0.1),
			Alt:                   unifieddatalibrary.Float(478.056378),
			Azimuth:               unifieddatalibrary.Float(134.5),
			AzimuthBias:           unifieddatalibrary.Float(0.123),
			AzimuthRate:           unifieddatalibrary.Float(0.5),
			AzimuthUnc:            unifieddatalibrary.Float(0.5),
			BistaticRange:         unifieddatalibrary.Float(754.8212),
			BistaticRangeAccel:    unifieddatalibrary.Float(1.23),
			BistaticRangeAccelUnc: unifieddatalibrary.Float(0.1),
			BistaticRangeBias:     unifieddatalibrary.Float(2.34),
			BistaticRangeRate:     unifieddatalibrary.Float(-0.30222),
			BistaticRangeRateUnc:  unifieddatalibrary.Float(0.123),
			BistaticRangeUnc:      unifieddatalibrary.Float(5.1),
			Coning:                unifieddatalibrary.Float(60.1),
			ConingUnc:             unifieddatalibrary.Float(0.5),
			Declination:           unifieddatalibrary.Float(10.23),
			Delay:                 unifieddatalibrary.Float(0.00505820232809312),
			DelayBias:             unifieddatalibrary.Float(0.00000123),
			DelayUnc:              unifieddatalibrary.Float(0.0000031),
			Descriptor:            unifieddatalibrary.String("Descriptor"),
			Doppler:               unifieddatalibrary.Float(-101.781641000597),
			DopplerUnc:            unifieddatalibrary.Float(0.2),
			Elevation:             unifieddatalibrary.Float(76.1),
			ElevationBias:         unifieddatalibrary.Float(0.123),
			ElevationRate:         unifieddatalibrary.Float(0.5),
			ElevationUnc:          unifieddatalibrary.Float(0.5),
			ExtObservationID:      unifieddatalibrary.String("26892"),
			IDRfEmitter:           unifieddatalibrary.String("RED_CLIFFS_3ABCRN"),
			IDSensor:              unifieddatalibrary.String("OCULUSA"),
			IDSensorRefReceiver:   unifieddatalibrary.String("OculusRef1"),
			Lat:                   unifieddatalibrary.Float(-35.1181763996856),
			Lon:                   unifieddatalibrary.Float(139.613567052763),
			ObPosition:            unifieddatalibrary.String("FIRST"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			OrthogonalRcs:         unifieddatalibrary.Float(10.23),
			OrthogonalRcsUnc:      unifieddatalibrary.Float(1.23),
			Ra:                    unifieddatalibrary.Float(1.23),
			Rcs:                   unifieddatalibrary.Float(100.23),
			RcsUnc:                unifieddatalibrary.Float(1.23),
			SatNo:                 unifieddatalibrary.Int(40699),
			Snr:                   unifieddatalibrary.Float(17.292053),
			Tags:                  []string{"TAG1", "TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TimingBias:            unifieddatalibrary.Float(1.23),
			Tof:                   unifieddatalibrary.Float(0.00592856674135648),
			TofBias:               unifieddatalibrary.Float(0.00000123),
			TofUnc:                unifieddatalibrary.Float(0.0000031),
			TrackID:               unifieddatalibrary.String("12212"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
			Xvel:                  unifieddatalibrary.Float(1.23),
			Yvel:                  unifieddatalibrary.Float(3.21),
			Zvel:                  unifieddatalibrary.Float(3.12),
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

func TestObservationPassiveRadarObservationFileNew(t *testing.T) {
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
	err := client.Observations.PassiveRadarObservation.FileNew(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationFileNewParams{
		Body: []unifieddatalibrary.ObservationPassiveRadarObservationFileNewParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("bdcacfb0-3c47-4bd0-9d6c-9fa7d2c4fbb0"),
			Accel:                 unifieddatalibrary.Float(1.23),
			AccelUnc:              unifieddatalibrary.Float(0.1),
			Alt:                   unifieddatalibrary.Float(478.056378),
			Azimuth:               unifieddatalibrary.Float(134.5),
			AzimuthBias:           unifieddatalibrary.Float(0.123),
			AzimuthRate:           unifieddatalibrary.Float(0.5),
			AzimuthUnc:            unifieddatalibrary.Float(0.5),
			BistaticRange:         unifieddatalibrary.Float(754.8212),
			BistaticRangeAccel:    unifieddatalibrary.Float(1.23),
			BistaticRangeAccelUnc: unifieddatalibrary.Float(0.1),
			BistaticRangeBias:     unifieddatalibrary.Float(2.34),
			BistaticRangeRate:     unifieddatalibrary.Float(-0.30222),
			BistaticRangeRateUnc:  unifieddatalibrary.Float(0.123),
			BistaticRangeUnc:      unifieddatalibrary.Float(5.1),
			Coning:                unifieddatalibrary.Float(60.1),
			ConingUnc:             unifieddatalibrary.Float(0.5),
			Declination:           unifieddatalibrary.Float(10.23),
			Delay:                 unifieddatalibrary.Float(0.00505820232809312),
			DelayBias:             unifieddatalibrary.Float(0.00000123),
			DelayUnc:              unifieddatalibrary.Float(0.0000031),
			Descriptor:            unifieddatalibrary.String("Descriptor"),
			Doppler:               unifieddatalibrary.Float(-101.781641000597),
			DopplerUnc:            unifieddatalibrary.Float(0.2),
			Elevation:             unifieddatalibrary.Float(76.1),
			ElevationBias:         unifieddatalibrary.Float(0.123),
			ElevationRate:         unifieddatalibrary.Float(0.5),
			ElevationUnc:          unifieddatalibrary.Float(0.5),
			ExtObservationID:      unifieddatalibrary.String("26892"),
			IDRfEmitter:           unifieddatalibrary.String("RED_CLIFFS_3ABCRN"),
			IDSensor:              unifieddatalibrary.String("OCULUSA"),
			IDSensorRefReceiver:   unifieddatalibrary.String("OculusRef1"),
			Lat:                   unifieddatalibrary.Float(-35.1181763996856),
			Lon:                   unifieddatalibrary.Float(139.613567052763),
			ObPosition:            unifieddatalibrary.String("FIRST"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			OrthogonalRcs:         unifieddatalibrary.Float(10.23),
			OrthogonalRcsUnc:      unifieddatalibrary.Float(1.23),
			Ra:                    unifieddatalibrary.Float(1.23),
			Rcs:                   unifieddatalibrary.Float(100.23),
			RcsUnc:                unifieddatalibrary.Float(1.23),
			SatNo:                 unifieddatalibrary.Int(40699),
			Snr:                   unifieddatalibrary.Float(17.292053),
			Tags:                  []string{"TAG1", "TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TimingBias:            unifieddatalibrary.Float(1.23),
			Tof:                   unifieddatalibrary.Float(0.00592856674135648),
			TofBias:               unifieddatalibrary.Float(0.00000123),
			TofUnc:                unifieddatalibrary.Float(0.0000031),
			TrackID:               unifieddatalibrary.String("12212"),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
			Xvel:                  unifieddatalibrary.Float(1.23),
			Yvel:                  unifieddatalibrary.Float(3.21),
			Zvel:                  unifieddatalibrary.Float(3.12),
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

func TestObservationPassiveRadarObservationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.PassiveRadarObservation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObservationPassiveRadarObservationGetParams{
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

func TestObservationPassiveRadarObservationQueryhelp(t *testing.T) {
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
	err := client.Observations.PassiveRadarObservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationPassiveRadarObservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.PassiveRadarObservation.Tuple(context.TODO(), unifieddatalibrary.ObservationPassiveRadarObservationTupleParams{
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
