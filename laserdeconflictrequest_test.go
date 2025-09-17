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

func TestLaserdeconflictrequestNewWithOptionalParams(t *testing.T) {
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
	err := client.Laserdeconflictrequest.New(context.TODO(), unifieddatalibrary.LaserdeconflictrequestNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LaserdeconflictrequestNewParamsDataModeTest,
		EndDate:               time.Now(),
		IDLaserEmitters:       []string{"2346c0a0-585f-4232-af5d-93bad320fdc0", "4446c0a0-585f-4232-af5d-93bad320fbb1"},
		NumTargets:            25,
		RequestID:             "3856c0a0-585f-4232-af5d-93bad320fac6",
		RequestTs:             time.Now(),
		Source:                "Bluestaq",
		StartDate:             time.Now(),
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		CenterlineAzimuth:     unifieddatalibrary.Float(20.3),
		CenterlineElevation:   unifieddatalibrary.Float(19.434),
		DefaultCha:            unifieddatalibrary.Float(2.5),
		EnableDss:             unifieddatalibrary.Bool(true),
		FixedPoints: []unifieddatalibrary.LaserdeconflictrequestNewParamsFixedPoint{{
			Latitude:  -10.18,
			Longitude: -179.98,
			Height:    unifieddatalibrary.Float(-18.13),
		}},
		GeopotentialModel: unifieddatalibrary.String("WGS84"),
		LaserDeconflictTargets: []unifieddatalibrary.LaserdeconflictrequestNewParamsLaserDeconflictTarget{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Azimuth:               unifieddatalibrary.Float(27.91),
			AzimuthEnd:            unifieddatalibrary.Float(90.5),
			AzimuthIncrement:      unifieddatalibrary.Float(1.5),
			AzimuthStart:          unifieddatalibrary.Float(60.5),
			CenterlineAzimuth:     unifieddatalibrary.Float(11.02),
			CenterlineElevation:   unifieddatalibrary.Float(1.68),
			Declination:           unifieddatalibrary.Float(10.23),
			Elevation:             unifieddatalibrary.Float(17.09),
			ElevationEnd:          unifieddatalibrary.Float(88.05),
			ElevationIncrement:    unifieddatalibrary.Float(0.5),
			ElevationStart:        unifieddatalibrary.Float(67.05),
			FixedPoints: []unifieddatalibrary.LaserdeconflictrequestNewParamsLaserDeconflictTargetFixedPoint{{
				Latitude:  -10.18,
				Longitude: -179.98,
				Height:    unifieddatalibrary.Float(-18.13),
			}},
			IDLaserDeconflictRequest: unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			LengthCenterline:         unifieddatalibrary.Float(369.79),
			LengthLeftRight:          unifieddatalibrary.Float(20.23),
			LengthUpDown:             unifieddatalibrary.Float(28.67),
			MaximumHeight:            unifieddatalibrary.Float(0.5),
			MinimumHeight:            unifieddatalibrary.Float(0.25),
			Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Ra:                       unifieddatalibrary.Float(11.93),
			SolarSystemBody:          unifieddatalibrary.String("MARS"),
			StarNumber:               unifieddatalibrary.Int(3791),
			TargetNumber:             unifieddatalibrary.Int(100),
			TargetObjectNo:           unifieddatalibrary.Int(46852),
		}},
		LaserSystemName:      unifieddatalibrary.String("HEL-1"),
		LengthCenterline:     unifieddatalibrary.Float(79.35),
		LengthLeftRight:      unifieddatalibrary.Float(56.23),
		LengthUpDown:         unifieddatalibrary.Float(22.6),
		MaximumHeight:        unifieddatalibrary.Float(440.3),
		MinimumHeight:        unifieddatalibrary.Float(0.5),
		MissionName:          unifieddatalibrary.String("USSF LP 18-1 Test Laser"),
		Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:         unifieddatalibrary.String("ORIGOBJECT-ID"),
		PlatformLocationName: unifieddatalibrary.String("Vandenberg"),
		PlatformLocationType: unifieddatalibrary.String("FIXED_POINT"),
		ProgramID:            unifieddatalibrary.String("performance_test_llh-sat"),
		Propagator:           unifieddatalibrary.String("GP"),
		ProtectList:          []int64{1234, 5678},
		SatNo:                unifieddatalibrary.Int(46852),
		SourceEnabled:        unifieddatalibrary.Bool(false),
		Status:               unifieddatalibrary.String("REQUESTED"),
		Tags:                 []string{"TAG1", "TAG2"},
		TargetEnabled:        unifieddatalibrary.Bool(true),
		TargetType:           unifieddatalibrary.String("BOX_CENTERPOINT_LINE"),
		TransactionID:        unifieddatalibrary.String("TRANSACTION-ID"),
		TreatEarthAs:         unifieddatalibrary.String("VICTIM"),
		UseFieldOfRegard:     unifieddatalibrary.Bool(true),
		VictimEnabled:        unifieddatalibrary.Bool(true),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaserdeconflictrequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Laserdeconflictrequest.List(context.TODO(), unifieddatalibrary.LaserdeconflictrequestListParams{
		StartDate:   time.Now(),
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

func TestLaserdeconflictrequestCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Laserdeconflictrequest.Count(context.TODO(), unifieddatalibrary.LaserdeconflictrequestCountParams{
		StartDate:   time.Now(),
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

func TestLaserdeconflictrequestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Laserdeconflictrequest.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.LaserdeconflictrequestGetParams{
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

func TestLaserdeconflictrequestQueryhelp(t *testing.T) {
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
	_, err := client.Laserdeconflictrequest.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaserdeconflictrequestTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Laserdeconflictrequest.Tuple(context.TODO(), unifieddatalibrary.LaserdeconflictrequestTupleParams{
		Columns:     "columns",
		StartDate:   time.Now(),
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

func TestLaserdeconflictrequestUnvalidatedPublishWithOptionalParams(t *testing.T) {
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
	err := client.Laserdeconflictrequest.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LaserdeconflictrequestUnvalidatedPublishParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.LaserdeconflictrequestUnvalidatedPublishParamsDataModeTest,
		EndDate:               time.Now(),
		IDLaserEmitters:       []string{"2346c0a0-585f-4232-af5d-93bad320fdc0", "4446c0a0-585f-4232-af5d-93bad320fbb1"},
		NumTargets:            25,
		RequestID:             "3856c0a0-585f-4232-af5d-93bad320fac6",
		RequestTs:             time.Now(),
		Source:                "Bluestaq",
		StartDate:             time.Now(),
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		CenterlineAzimuth:     unifieddatalibrary.Float(20.3),
		CenterlineElevation:   unifieddatalibrary.Float(19.434),
		DefaultCha:            unifieddatalibrary.Float(2.5),
		EnableDss:             unifieddatalibrary.Bool(true),
		FixedPoints: []unifieddatalibrary.LaserdeconflictrequestUnvalidatedPublishParamsFixedPoint{{
			Latitude:  -10.18,
			Longitude: -179.98,
			Height:    unifieddatalibrary.Float(-18.13),
		}},
		GeopotentialModel: unifieddatalibrary.String("WGS84"),
		LaserDeconflictTargets: []unifieddatalibrary.LaserdeconflictrequestUnvalidatedPublishParamsLaserDeconflictTarget{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Azimuth:               unifieddatalibrary.Float(27.91),
			AzimuthEnd:            unifieddatalibrary.Float(90.5),
			AzimuthIncrement:      unifieddatalibrary.Float(1.5),
			AzimuthStart:          unifieddatalibrary.Float(60.5),
			CenterlineAzimuth:     unifieddatalibrary.Float(11.02),
			CenterlineElevation:   unifieddatalibrary.Float(1.68),
			Declination:           unifieddatalibrary.Float(10.23),
			Elevation:             unifieddatalibrary.Float(17.09),
			ElevationEnd:          unifieddatalibrary.Float(88.05),
			ElevationIncrement:    unifieddatalibrary.Float(0.5),
			ElevationStart:        unifieddatalibrary.Float(67.05),
			FixedPoints: []unifieddatalibrary.LaserdeconflictrequestUnvalidatedPublishParamsLaserDeconflictTargetFixedPoint{{
				Latitude:  -10.18,
				Longitude: -179.98,
				Height:    unifieddatalibrary.Float(-18.13),
			}},
			IDLaserDeconflictRequest: unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			LengthCenterline:         unifieddatalibrary.Float(369.79),
			LengthLeftRight:          unifieddatalibrary.Float(20.23),
			LengthUpDown:             unifieddatalibrary.Float(28.67),
			MaximumHeight:            unifieddatalibrary.Float(0.5),
			MinimumHeight:            unifieddatalibrary.Float(0.25),
			Origin:                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Ra:                       unifieddatalibrary.Float(11.93),
			SolarSystemBody:          unifieddatalibrary.String("MARS"),
			StarNumber:               unifieddatalibrary.Int(3791),
			TargetNumber:             unifieddatalibrary.Int(100),
			TargetObjectNo:           unifieddatalibrary.Int(46852),
		}},
		LaserSystemName:      unifieddatalibrary.String("HEL-1"),
		LengthCenterline:     unifieddatalibrary.Float(79.35),
		LengthLeftRight:      unifieddatalibrary.Float(56.23),
		LengthUpDown:         unifieddatalibrary.Float(22.6),
		MaximumHeight:        unifieddatalibrary.Float(440.3),
		MinimumHeight:        unifieddatalibrary.Float(0.5),
		MissionName:          unifieddatalibrary.String("USSF LP 18-1 Test Laser"),
		Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:         unifieddatalibrary.String("ORIGOBJECT-ID"),
		PlatformLocationName: unifieddatalibrary.String("Vandenberg"),
		PlatformLocationType: unifieddatalibrary.String("FIXED_POINT"),
		ProgramID:            unifieddatalibrary.String("performance_test_llh-sat"),
		Propagator:           unifieddatalibrary.String("GP"),
		ProtectList:          []int64{1234, 5678},
		SatNo:                unifieddatalibrary.Int(46852),
		SourceEnabled:        unifieddatalibrary.Bool(false),
		Status:               unifieddatalibrary.String("REQUESTED"),
		Tags:                 []string{"TAG1", "TAG2"},
		TargetEnabled:        unifieddatalibrary.Bool(true),
		TargetType:           unifieddatalibrary.String("BOX_CENTERPOINT_LINE"),
		TransactionID:        unifieddatalibrary.String("TRANSACTION-ID"),
		TreatEarthAs:         unifieddatalibrary.String("VICTIM"),
		UseFieldOfRegard:     unifieddatalibrary.Bool(true),
		VictimEnabled:        unifieddatalibrary.Bool(true),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
