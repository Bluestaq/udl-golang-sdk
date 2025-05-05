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

func TestTrackRouteNewWithOptionalParams(t *testing.T) {
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
	err := client.TrackRoute.New(context.TODO(), unifieddatalibrary.TrackRouteNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.TrackRouteNewParamsDataModeTest,
		LastUpdateDate:        time.Now(),
		Source:                "Bluestaq",
		Type:                  "AIR REFUELING",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AltitudeBlocks: []unifieddatalibrary.TrackRouteNewParamsAltitudeBlock{{
			AltitudeSequenceID: unifieddatalibrary.String("A1"),
			LowerAltitude:      unifieddatalibrary.Float(27000.1),
			UpperAltitude:      unifieddatalibrary.Float(27200.5),
		}},
		ApnSetting:      unifieddatalibrary.String("1-3-1"),
		ApxBeaconCode:   unifieddatalibrary.String("5/1"),
		ArtccMessage:    unifieddatalibrary.String("OAKLAND CTR/GUAM CERAP"),
		CreatingOrg:     unifieddatalibrary.String("HQPAC"),
		Direction:       unifieddatalibrary.String("NE"),
		EffectiveDate:   unifieddatalibrary.Time(time.Now()),
		ExternalID:      unifieddatalibrary.String("GDSSMH121004232315303094"),
		LastUsedDate:    unifieddatalibrary.Time(time.Now()),
		LocationTrackID: unifieddatalibrary.String("POACHR"),
		Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Poc: []unifieddatalibrary.TrackRouteNewParamsPoc{{
			Office:        unifieddatalibrary.String("A34"),
			Phone:         unifieddatalibrary.String("8675309"),
			PocName:       unifieddatalibrary.String("Fred Smith"),
			PocOrg:        unifieddatalibrary.String("HQAF"),
			PocSequenceID: unifieddatalibrary.Int(1),
			PocTypeName:   unifieddatalibrary.String("Originator"),
			Rank:          unifieddatalibrary.String("Capt"),
			Remark:        unifieddatalibrary.String("POC remark."),
			Username:      unifieddatalibrary.String("fgsmith"),
		}},
		PriFreq:              unifieddatalibrary.Float(357.5),
		ReceiverTankerChCode: unifieddatalibrary.String("31/094"),
		RegionCode:           unifieddatalibrary.String("5"),
		RegionName:           unifieddatalibrary.String("North America"),
		ReviewDate:           unifieddatalibrary.Time(time.Now()),
		RoutePoints: []unifieddatalibrary.TrackRouteNewParamsRoutePoint{{
			AltCountryCode: unifieddatalibrary.String("IZ"),
			CountryCode:    unifieddatalibrary.String("NL"),
			DafifPt:        unifieddatalibrary.Bool(true),
			MagDec:         unifieddatalibrary.Float(7.35),
			Navaid:         unifieddatalibrary.String("HTO"),
			NavaidLength:   unifieddatalibrary.Float(100.2),
			NavaidType:     unifieddatalibrary.String("VORTAC"),
			PtLat:          unifieddatalibrary.Float(45.23),
			PtLon:          unifieddatalibrary.Float(179.1),
			PtSequenceID:   unifieddatalibrary.Int(1),
			PtTypeCode:     unifieddatalibrary.String("EP"),
			PtTypeName:     unifieddatalibrary.String("ENTRY POINT"),
			WaypointName:   unifieddatalibrary.String("KCHS"),
		}},
		SchedulerOrgName: unifieddatalibrary.String("97 OSS/OSOS DSN 866-5555"),
		SchedulerOrgUnit: unifieddatalibrary.String("612 AOC"),
		SecFreq:          unifieddatalibrary.Float(319.7),
		ShortName:        unifieddatalibrary.String("CH61"),
		Sic:              unifieddatalibrary.String("N"),
		TrackID:          unifieddatalibrary.String("CH61A"),
		TrackName:        unifieddatalibrary.String("CH61 POST"),
		TypeCode:         unifieddatalibrary.String("V"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteUpdateWithOptionalParams(t *testing.T) {
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
	err := client.TrackRoute.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.TrackRouteUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.TrackRouteUpdateParamsDataModeTest,
			LastUpdateDate:        time.Now(),
			Source:                "Bluestaq",
			Type:                  "AIR REFUELING",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AltitudeBlocks: []unifieddatalibrary.TrackRouteUpdateParamsAltitudeBlock{{
				AltitudeSequenceID: unifieddatalibrary.String("A1"),
				LowerAltitude:      unifieddatalibrary.Float(27000.1),
				UpperAltitude:      unifieddatalibrary.Float(27200.5),
			}},
			ApnSetting:      unifieddatalibrary.String("1-3-1"),
			ApxBeaconCode:   unifieddatalibrary.String("5/1"),
			ArtccMessage:    unifieddatalibrary.String("OAKLAND CTR/GUAM CERAP"),
			CreatingOrg:     unifieddatalibrary.String("HQPAC"),
			Direction:       unifieddatalibrary.String("NE"),
			EffectiveDate:   unifieddatalibrary.Time(time.Now()),
			ExternalID:      unifieddatalibrary.String("GDSSMH121004232315303094"),
			LastUsedDate:    unifieddatalibrary.Time(time.Now()),
			LocationTrackID: unifieddatalibrary.String("POACHR"),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Poc: []unifieddatalibrary.TrackRouteUpdateParamsPoc{{
				Office:        unifieddatalibrary.String("A34"),
				Phone:         unifieddatalibrary.String("8675309"),
				PocName:       unifieddatalibrary.String("Fred Smith"),
				PocOrg:        unifieddatalibrary.String("HQAF"),
				PocSequenceID: unifieddatalibrary.Int(1),
				PocTypeName:   unifieddatalibrary.String("Originator"),
				Rank:          unifieddatalibrary.String("Capt"),
				Remark:        unifieddatalibrary.String("POC remark."),
				Username:      unifieddatalibrary.String("fgsmith"),
			}},
			PriFreq:              unifieddatalibrary.Float(357.5),
			ReceiverTankerChCode: unifieddatalibrary.String("31/094"),
			RegionCode:           unifieddatalibrary.String("5"),
			RegionName:           unifieddatalibrary.String("North America"),
			ReviewDate:           unifieddatalibrary.Time(time.Now()),
			RoutePoints: []unifieddatalibrary.TrackRouteUpdateParamsRoutePoint{{
				AltCountryCode: unifieddatalibrary.String("IZ"),
				CountryCode:    unifieddatalibrary.String("NL"),
				DafifPt:        unifieddatalibrary.Bool(true),
				MagDec:         unifieddatalibrary.Float(7.35),
				Navaid:         unifieddatalibrary.String("HTO"),
				NavaidLength:   unifieddatalibrary.Float(100.2),
				NavaidType:     unifieddatalibrary.String("VORTAC"),
				PtLat:          unifieddatalibrary.Float(45.23),
				PtLon:          unifieddatalibrary.Float(179.1),
				PtSequenceID:   unifieddatalibrary.Int(1),
				PtTypeCode:     unifieddatalibrary.String("EP"),
				PtTypeName:     unifieddatalibrary.String("ENTRY POINT"),
				WaypointName:   unifieddatalibrary.String("KCHS"),
			}},
			SchedulerOrgName: unifieddatalibrary.String("97 OSS/OSOS DSN 866-5555"),
			SchedulerOrgUnit: unifieddatalibrary.String("612 AOC"),
			SecFreq:          unifieddatalibrary.Float(319.7),
			ShortName:        unifieddatalibrary.String("CH61"),
			Sic:              unifieddatalibrary.String("N"),
			TrackID:          unifieddatalibrary.String("CH61A"),
			TrackName:        unifieddatalibrary.String("CH61 POST"),
			TypeCode:         unifieddatalibrary.String("V"),
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

func TestTrackRouteListWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackRoute.List(context.TODO(), unifieddatalibrary.TrackRouteListParams{
		LastUpdateDate: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteDelete(t *testing.T) {
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
	err := client.TrackRoute.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteCountWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackRoute.Count(context.TODO(), unifieddatalibrary.TrackRouteCountParams{
		LastUpdateDate: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteNewBulk(t *testing.T) {
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
	err := client.TrackRoute.NewBulk(context.TODO(), unifieddatalibrary.TrackRouteNewBulkParams{
		Body: []unifieddatalibrary.TrackRouteNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			LastUpdateDate:        time.Now(),
			Source:                "Bluestaq",
			Type:                  "AIR REFUELING",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AltitudeBlocks: []unifieddatalibrary.TrackRouteNewBulkParamsBodyAltitudeBlock{{
				AltitudeSequenceID: unifieddatalibrary.String("A1"),
				LowerAltitude:      unifieddatalibrary.Float(27000.1),
				UpperAltitude:      unifieddatalibrary.Float(27200.5),
			}},
			ApnSetting:      unifieddatalibrary.String("1-3-1"),
			ApxBeaconCode:   unifieddatalibrary.String("5/1"),
			ArtccMessage:    unifieddatalibrary.String("OAKLAND CTR/GUAM CERAP"),
			CreatingOrg:     unifieddatalibrary.String("HQPAC"),
			Direction:       unifieddatalibrary.String("NE"),
			EffectiveDate:   unifieddatalibrary.Time(time.Now()),
			ExternalID:      unifieddatalibrary.String("GDSSMH121004232315303094"),
			LastUsedDate:    unifieddatalibrary.Time(time.Now()),
			LocationTrackID: unifieddatalibrary.String("POACHR"),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Poc: []unifieddatalibrary.TrackRouteNewBulkParamsBodyPoc{{
				Office:        unifieddatalibrary.String("A34"),
				Phone:         unifieddatalibrary.String("8675309"),
				PocName:       unifieddatalibrary.String("Fred Smith"),
				PocOrg:        unifieddatalibrary.String("HQAF"),
				PocSequenceID: unifieddatalibrary.Int(1),
				PocTypeName:   unifieddatalibrary.String("Originator"),
				Rank:          unifieddatalibrary.String("Capt"),
				Remark:        unifieddatalibrary.String("POC remark."),
				Username:      unifieddatalibrary.String("fgsmith"),
			}},
			PriFreq:              unifieddatalibrary.Float(357.5),
			ReceiverTankerChCode: unifieddatalibrary.String("31/094"),
			RegionCode:           unifieddatalibrary.String("5"),
			RegionName:           unifieddatalibrary.String("North America"),
			ReviewDate:           unifieddatalibrary.Time(time.Now()),
			RoutePoints: []unifieddatalibrary.TrackRouteNewBulkParamsBodyRoutePoint{{
				AltCountryCode: unifieddatalibrary.String("IZ"),
				CountryCode:    unifieddatalibrary.String("NL"),
				DafifPt:        unifieddatalibrary.Bool(true),
				MagDec:         unifieddatalibrary.Float(7.35),
				Navaid:         unifieddatalibrary.String("HTO"),
				NavaidLength:   unifieddatalibrary.Float(100.2),
				NavaidType:     unifieddatalibrary.String("VORTAC"),
				PtLat:          unifieddatalibrary.Float(45.23),
				PtLon:          unifieddatalibrary.Float(179.1),
				PtSequenceID:   unifieddatalibrary.Int(1),
				PtTypeCode:     unifieddatalibrary.String("EP"),
				PtTypeName:     unifieddatalibrary.String("ENTRY POINT"),
				WaypointName:   unifieddatalibrary.String("KCHS"),
			}},
			SchedulerOrgName: unifieddatalibrary.String("97 OSS/OSOS DSN 866-5555"),
			SchedulerOrgUnit: unifieddatalibrary.String("612 AOC"),
			SecFreq:          unifieddatalibrary.Float(319.7),
			ShortName:        unifieddatalibrary.String("CH61"),
			Sic:              unifieddatalibrary.String("N"),
			TrackID:          unifieddatalibrary.String("CH61A"),
			TrackName:        unifieddatalibrary.String("CH61 POST"),
			TypeCode:         unifieddatalibrary.String("V"),
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

func TestTrackRouteGetWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackRoute.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.TrackRouteGetParams{
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

func TestTrackRouteQueryhelp(t *testing.T) {
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
	err := client.TrackRoute.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.TrackRoute.Tuple(context.TODO(), unifieddatalibrary.TrackRouteTupleParams{
		Columns:        "columns",
		LastUpdateDate: time.Now(),
		FirstResult:    unifieddatalibrary.Int(0),
		MaxResults:     unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackRouteUnvalidatedPublishWithOptionalParams(t *testing.T) {
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
	err := client.TrackRoute.UnvalidatedPublish(context.TODO(), unifieddatalibrary.TrackRouteUnvalidatedPublishParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.TrackRouteUnvalidatedPublishParamsDataModeTest,
		LastUpdateDate:        time.Now(),
		Source:                "Bluestaq",
		Type:                  "AIR REFUELING",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AltitudeBlocks: []unifieddatalibrary.TrackRouteUnvalidatedPublishParamsAltitudeBlock{{
			AltitudeSequenceID: unifieddatalibrary.String("A1"),
			LowerAltitude:      unifieddatalibrary.Float(27000.1),
			UpperAltitude:      unifieddatalibrary.Float(27200.5),
		}},
		ApnSetting:      unifieddatalibrary.String("1-3-1"),
		ApxBeaconCode:   unifieddatalibrary.String("5/1"),
		ArtccMessage:    unifieddatalibrary.String("OAKLAND CTR/GUAM CERAP"),
		CreatingOrg:     unifieddatalibrary.String("HQPAC"),
		Direction:       unifieddatalibrary.String("NE"),
		EffectiveDate:   unifieddatalibrary.Time(time.Now()),
		ExternalID:      unifieddatalibrary.String("GDSSMH121004232315303094"),
		LastUsedDate:    unifieddatalibrary.Time(time.Now()),
		LocationTrackID: unifieddatalibrary.String("POACHR"),
		Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Poc: []unifieddatalibrary.TrackRouteUnvalidatedPublishParamsPoc{{
			Office:        unifieddatalibrary.String("A34"),
			Phone:         unifieddatalibrary.String("8675309"),
			PocName:       unifieddatalibrary.String("Fred Smith"),
			PocOrg:        unifieddatalibrary.String("HQAF"),
			PocSequenceID: unifieddatalibrary.Int(1),
			PocTypeName:   unifieddatalibrary.String("Originator"),
			Rank:          unifieddatalibrary.String("Capt"),
			Remark:        unifieddatalibrary.String("POC remark."),
			Username:      unifieddatalibrary.String("fgsmith"),
		}},
		PriFreq:              unifieddatalibrary.Float(357.5),
		ReceiverTankerChCode: unifieddatalibrary.String("31/094"),
		RegionCode:           unifieddatalibrary.String("5"),
		RegionName:           unifieddatalibrary.String("North America"),
		ReviewDate:           unifieddatalibrary.Time(time.Now()),
		RoutePoints: []unifieddatalibrary.TrackRouteUnvalidatedPublishParamsRoutePoint{{
			AltCountryCode: unifieddatalibrary.String("IZ"),
			CountryCode:    unifieddatalibrary.String("NL"),
			DafifPt:        unifieddatalibrary.Bool(true),
			MagDec:         unifieddatalibrary.Float(7.35),
			Navaid:         unifieddatalibrary.String("HTO"),
			NavaidLength:   unifieddatalibrary.Float(100.2),
			NavaidType:     unifieddatalibrary.String("VORTAC"),
			PtLat:          unifieddatalibrary.Float(45.23),
			PtLon:          unifieddatalibrary.Float(179.1),
			PtSequenceID:   unifieddatalibrary.Int(1),
			PtTypeCode:     unifieddatalibrary.String("EP"),
			PtTypeName:     unifieddatalibrary.String("ENTRY POINT"),
			WaypointName:   unifieddatalibrary.String("KCHS"),
		}},
		SchedulerOrgName: unifieddatalibrary.String("97 OSS/OSOS DSN 866-5555"),
		SchedulerOrgUnit: unifieddatalibrary.String("612 AOC"),
		SecFreq:          unifieddatalibrary.Float(319.7),
		ShortName:        unifieddatalibrary.String("CH61"),
		Sic:              unifieddatalibrary.String("N"),
		TrackID:          unifieddatalibrary.String("CH61A"),
		TrackName:        unifieddatalibrary.String("CH61 POST"),
		TypeCode:         unifieddatalibrary.String("V"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
