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

func TestOrbittrackListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbittrack.List(context.TODO(), unifieddatalibrary.OrbittrackListParams{
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

func TestOrbittrackCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbittrack.Count(context.TODO(), unifieddatalibrary.OrbittrackCountParams{
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

func TestOrbittrackNewBulk(t *testing.T) {
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
	err := client.Orbittrack.NewBulk(context.TODO(), unifieddatalibrary.OrbittrackNewBulkParams{
		Body: []unifieddatalibrary.OrbittrackNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Lat:                   19.88550102,
			Lon:                   46.74596844,
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("ORBIT_TRACK_ID"),
			Alt:                   unifieddatalibrary.Float(585.71),
			Amplification:         unifieddatalibrary.String("A note regarding this spacecraft"),
			AngElev:               unifieddatalibrary.Float(15.2),
			AouData:               []float64{34.3, 26.5, 1.2},
			AouType:               unifieddatalibrary.String("ELLIPSE"),
			CallSign:              unifieddatalibrary.String("Charlie"),
			CharlieLine:           unifieddatalibrary.String("323751332255940400010000003635829600010200072500098205001150"),
			ChXRef:                unifieddatalibrary.String("FHKX"),
			Cntnmnt:               unifieddatalibrary.Float(90),
			CountryCode:           unifieddatalibrary.String("US"),
			Decay:                 unifieddatalibrary.Float(0.5868),
			Dummy:                 unifieddatalibrary.Bool(false),
			Feint:                 unifieddatalibrary.Bool(false),
			Hq:                    unifieddatalibrary.Bool(false),
			IDElset:               unifieddatalibrary.String("c715a619-8695-44d2-9e7d-effd257b4843"),
			IdentAmp:              unifieddatalibrary.String("JOKER"),
			IDOnOrbit:             unifieddatalibrary.String("32375"),
			Iff:                   unifieddatalibrary.String("ID Mode"),
			Installation:          unifieddatalibrary.Bool(false),
			ObjectType:            "PAYLOAD",
			ObjIdent:              "FRIEND",
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("L2045"),
			RdfRf:                 unifieddatalibrary.Float(1.5273),
			Reduced:               unifieddatalibrary.Bool(false),
			Reinforced:            unifieddatalibrary.Bool(false),
			RptNum:                unifieddatalibrary.String("123"),
			SatNo:                 unifieddatalibrary.Int(37375),
			SatStatus:             unifieddatalibrary.String("INACTIVE"),
			Spd:                   unifieddatalibrary.Float(15.03443),
			TaskForce:             unifieddatalibrary.Bool(false),
			TrackSensors: []unifieddatalibrary.OrbittrackNewBulkParamsBodyTrackSensor{{
				Az:            90,
				Range:         4023.95,
				MinRangeLimit: unifieddatalibrary.Float(20.23),
				MissionNumber: unifieddatalibrary.String("Example Mission"),
				SensorFovType: "UNKNOWN",
				SensorName:    unifieddatalibrary.String("SENSOR_NAME"),
				SensorNumber:  unifieddatalibrary.Int(1234),
			}},
			TrkID:   unifieddatalibrary.String("3668f135-fcba-4630-a43d-e7782e11d988"),
			VehType: unifieddatalibrary.String("SPACE"),
			Xref:    unifieddatalibrary.String("XE"),
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

func TestOrbittrackQueryhelp(t *testing.T) {
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
	_, err := client.Orbittrack.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbittrackTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbittrack.Tuple(context.TODO(), unifieddatalibrary.OrbittrackTupleParams{
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

func TestOrbittrackUnvalidatedPublish(t *testing.T) {
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
	err := client.Orbittrack.UnvalidatedPublish(context.TODO(), unifieddatalibrary.OrbittrackUnvalidatedPublishParams{
		Body: []unifieddatalibrary.OrbittrackUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Lat:                   19.88550102,
			Lon:                   46.74596844,
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("ORBIT_TRACK_ID"),
			Alt:                   unifieddatalibrary.Float(585.71),
			Amplification:         unifieddatalibrary.String("A note regarding this spacecraft"),
			AngElev:               unifieddatalibrary.Float(15.2),
			AouData:               []float64{34.3, 26.5, 1.2},
			AouType:               unifieddatalibrary.String("ELLIPSE"),
			CallSign:              unifieddatalibrary.String("Charlie"),
			CharlieLine:           unifieddatalibrary.String("323751332255940400010000003635829600010200072500098205001150"),
			ChXRef:                unifieddatalibrary.String("FHKX"),
			Cntnmnt:               unifieddatalibrary.Float(90),
			CountryCode:           unifieddatalibrary.String("US"),
			Decay:                 unifieddatalibrary.Float(0.5868),
			Dummy:                 unifieddatalibrary.Bool(false),
			Feint:                 unifieddatalibrary.Bool(false),
			Hq:                    unifieddatalibrary.Bool(false),
			IDElset:               unifieddatalibrary.String("c715a619-8695-44d2-9e7d-effd257b4843"),
			IdentAmp:              unifieddatalibrary.String("JOKER"),
			IDOnOrbit:             unifieddatalibrary.String("32375"),
			Iff:                   unifieddatalibrary.String("ID Mode"),
			Installation:          unifieddatalibrary.Bool(false),
			ObjectType:            "PAYLOAD",
			ObjIdent:              "FRIEND",
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("L2045"),
			RdfRf:                 unifieddatalibrary.Float(1.5273),
			Reduced:               unifieddatalibrary.Bool(false),
			Reinforced:            unifieddatalibrary.Bool(false),
			RptNum:                unifieddatalibrary.String("123"),
			SatNo:                 unifieddatalibrary.Int(37375),
			SatStatus:             unifieddatalibrary.String("INACTIVE"),
			Spd:                   unifieddatalibrary.Float(15.03443),
			TaskForce:             unifieddatalibrary.Bool(false),
			TrackSensors: []unifieddatalibrary.OrbittrackUnvalidatedPublishParamsBodyTrackSensor{{
				Az:            90,
				Range:         4023.95,
				MinRangeLimit: unifieddatalibrary.Float(20.23),
				MissionNumber: unifieddatalibrary.String("Example Mission"),
				SensorFovType: "UNKNOWN",
				SensorName:    unifieddatalibrary.String("SENSOR_NAME"),
				SensorNumber:  unifieddatalibrary.Int(1234),
			}},
			TrkID:   unifieddatalibrary.String("3668f135-fcba-4630-a43d-e7782e11d988"),
			VehType: unifieddatalibrary.String("SPACE"),
			Xref:    unifieddatalibrary.String("XE"),
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
