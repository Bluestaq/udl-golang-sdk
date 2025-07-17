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

func TestSpaceEnvObservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.SpaceEnvObservation.List(context.TODO(), unifieddatalibrary.SpaceEnvObservationListParams{
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

func TestSpaceEnvObservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SpaceEnvObservation.Count(context.TODO(), unifieddatalibrary.SpaceEnvObservationCountParams{
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

func TestSpaceEnvObservationNewBulk(t *testing.T) {
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
	err := client.SpaceEnvObservation.NewBulk(context.TODO(), unifieddatalibrary.SpaceEnvObservationNewBulkParams{
		Body: []unifieddatalibrary.SpaceEnvObservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("f13f82b8-5d2a-407a-b3f3-1fe30ca79eae"),
			Alt:                   unifieddatalibrary.Float(35785.3732),
			DataType:              unifieddatalibrary.String("ECP"),
			Derived:               unifieddatalibrary.Bool(false),
			Description:           unifieddatalibrary.String("Observation Data Description Text"),
			Descriptor:            unifieddatalibrary.String("energetic charged particle"),
			ExternalID:            unifieddatalibrary.String("fe4ad5dc-0128-4ce8-b09c-0b404322025e"),
			Forecast:              unifieddatalibrary.Bool(true),
			GenSystem:             unifieddatalibrary.String("System Name"),
			GenTime:               unifieddatalibrary.Time(time.Now()),
			IDSensor:              unifieddatalibrary.String("ECP-1"),
			InstrumentType:        unifieddatalibrary.String("MAGNETOMETER"),
			Lat:                   unifieddatalibrary.Float(38.8339),
			Lon:                   unifieddatalibrary.Float(-104.8214),
			MeasType:              unifieddatalibrary.String("ENERGETIC PARTICLES"),
			MsgType:               unifieddatalibrary.String("SODM"),
			ObservatoryName:       unifieddatalibrary.String("GOES-16"),
			ObservatoryNotes:      unifieddatalibrary.String("Notes"),
			ObservatoryType:       unifieddatalibrary.String("ONORBIT"),
			ObSetID:               unifieddatalibrary.String("ECPOBSET-478125"),
			Origin:                unifieddatalibrary.String("OPS1"),
			OrigObjectID:          unifieddatalibrary.String("41866"),
			OrigSensorID:          unifieddatalibrary.String("ECP-1"),
			ParticleType:          unifieddatalibrary.String("PROTON"),
			Quality:               unifieddatalibrary.String("GOOD"),
			SatNo:                 unifieddatalibrary.Int(41866),
			SenEnergyLevel:        unifieddatalibrary.String("0500-700 keV"),
			SenPos:                []float64{4174.78541785946, -9969.69867853067, 40733.9284531208},
			SenReferenceFrame:     "J2000",
			SenVel:                []float64{0.727059797295872, 0.298037087322647, 0.00157064850994095},
			SeoList: []unifieddatalibrary.SpaceEnvObservationNewBulkParamsBodySeoList{{
				ObType:        "INTEGRAL FLUX",
				ObUoM:         "particles/cm^2/s/sr",
				ObArray:       []float64{1.7, 35.6, 21.2, 19.01},
				ObBool:        unifieddatalibrary.Bool(true),
				ObDescription: unifieddatalibrary.String("Observation Description Text"),
				ObQuality:     unifieddatalibrary.String("GOOD"),
				ObString:      unifieddatalibrary.String("C1.3"),
				ObValue:       unifieddatalibrary.Float(0.487687826),
			}},
			SrcIDs:  []string{"615236d7-d464-4b8c-9b0b-45994e017d80", "8a38e3d6-35fd-4bda-b883-e3724e6bc6b9"},
			SrcTyps: []string{"SPACEENVOB", "SGI"},
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

func TestSpaceEnvObservationQueryhelp(t *testing.T) {
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
	_, err := client.SpaceEnvObservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpaceEnvObservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SpaceEnvObservation.Tuple(context.TODO(), unifieddatalibrary.SpaceEnvObservationTupleParams{
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

func TestSpaceEnvObservationUnvalidatedPublish(t *testing.T) {
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
	err := client.SpaceEnvObservation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SpaceEnvObservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SpaceEnvObservationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("f13f82b8-5d2a-407a-b3f3-1fe30ca79eae"),
			Alt:                   unifieddatalibrary.Float(35785.3732),
			DataType:              unifieddatalibrary.String("ECP"),
			Derived:               unifieddatalibrary.Bool(false),
			Description:           unifieddatalibrary.String("Observation Data Description Text"),
			Descriptor:            unifieddatalibrary.String("energetic charged particle"),
			ExternalID:            unifieddatalibrary.String("fe4ad5dc-0128-4ce8-b09c-0b404322025e"),
			Forecast:              unifieddatalibrary.Bool(true),
			GenSystem:             unifieddatalibrary.String("System Name"),
			GenTime:               unifieddatalibrary.Time(time.Now()),
			IDSensor:              unifieddatalibrary.String("ECP-1"),
			InstrumentType:        unifieddatalibrary.String("MAGNETOMETER"),
			Lat:                   unifieddatalibrary.Float(38.8339),
			Lon:                   unifieddatalibrary.Float(-104.8214),
			MeasType:              unifieddatalibrary.String("ENERGETIC PARTICLES"),
			MsgType:               unifieddatalibrary.String("SODM"),
			ObservatoryName:       unifieddatalibrary.String("GOES-16"),
			ObservatoryNotes:      unifieddatalibrary.String("Notes"),
			ObservatoryType:       unifieddatalibrary.String("ONORBIT"),
			ObSetID:               unifieddatalibrary.String("ECPOBSET-478125"),
			Origin:                unifieddatalibrary.String("OPS1"),
			OrigObjectID:          unifieddatalibrary.String("41866"),
			OrigSensorID:          unifieddatalibrary.String("ECP-1"),
			ParticleType:          unifieddatalibrary.String("PROTON"),
			Quality:               unifieddatalibrary.String("GOOD"),
			SatNo:                 unifieddatalibrary.Int(41866),
			SenEnergyLevel:        unifieddatalibrary.String("0500-700 keV"),
			SenPos:                []float64{4174.78541785946, -9969.69867853067, 40733.9284531208},
			SenReferenceFrame:     "J2000",
			SenVel:                []float64{0.727059797295872, 0.298037087322647, 0.00157064850994095},
			SeoList: []unifieddatalibrary.SpaceEnvObservationUnvalidatedPublishParamsBodySeoList{{
				ObType:        "INTEGRAL FLUX",
				ObUoM:         "particles/cm^2/s/sr",
				ObArray:       []float64{1.7, 35.6, 21.2, 19.01},
				ObBool:        unifieddatalibrary.Bool(true),
				ObDescription: unifieddatalibrary.String("Observation Description Text"),
				ObQuality:     unifieddatalibrary.String("GOOD"),
				ObString:      unifieddatalibrary.String("C1.3"),
				ObValue:       unifieddatalibrary.Float(0.487687826),
			}},
			SrcIDs:  []string{"615236d7-d464-4b8c-9b0b-45994e017d80", "8a38e3d6-35fd-4bda-b883-e3724e6bc6b9"},
			SrcTyps: []string{"SPACEENVOB", "SGI"},
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
