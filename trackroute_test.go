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

func TestTrackrouteNewWithOptionalParams(t *testing.T) {
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
	err := client.Trackroute.New(context.TODO(), unifieddatalibrary.TrackrouteNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.TrackrouteNewParamsDataModeTest,
		LastUpdateDate:        time.Now(),
		Source:                "Bluestaq",
		Type:                  "AIR REFUELING",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AltitudeBlocks: []unifieddatalibrary.TrackrouteNewParamsAltitudeBlock{{
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
		Poc: []unifieddatalibrary.TrackrouteNewParamsPoc{{
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
		RoutePoints: []unifieddatalibrary.TrackrouteNewParamsRoutePoint{{
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

func TestTrackrouteUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Trackroute.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.TrackrouteUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.TrackrouteUpdateParamsDataModeTest,
			LastUpdateDate:        time.Now(),
			Source:                "Bluestaq",
			Type:                  "AIR REFUELING",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AltitudeBlocks: []unifieddatalibrary.TrackrouteUpdateParamsAltitudeBlock{{
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
			Poc: []unifieddatalibrary.TrackrouteUpdateParamsPoc{{
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
			RoutePoints: []unifieddatalibrary.TrackrouteUpdateParamsRoutePoint{{
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

func TestTrackrouteListWithOptionalParams(t *testing.T) {
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
	_, err := client.Trackroute.List(context.TODO(), unifieddatalibrary.TrackrouteListParams{
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

func TestTrackrouteDelete(t *testing.T) {
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
	err := client.Trackroute.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackrouteCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Trackroute.Count(context.TODO(), unifieddatalibrary.TrackrouteCountParams{
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

func TestTrackrouteNewBulk(t *testing.T) {
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
	err := client.Trackroute.NewBulk(context.TODO(), unifieddatalibrary.TrackrouteNewBulkParams{
		Body: []unifieddatalibrary.TrackrouteNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			LastUpdateDate:        time.Now(),
			Source:                "Bluestaq",
			Type:                  "AIR REFUELING",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AltitudeBlocks: []unifieddatalibrary.TrackrouteNewBulkParamsBodyAltitudeBlock{{
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
			Poc: []unifieddatalibrary.TrackrouteNewBulkParamsBodyPoc{{
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
			RoutePoints: []unifieddatalibrary.TrackrouteNewBulkParamsBodyRoutePoint{{
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

func TestTrackrouteGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Trackroute.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.TrackrouteGetParams{
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

func TestTrackrouteQueryhelp(t *testing.T) {
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
	err := client.Trackroute.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrackrouteTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Trackroute.Tuple(context.TODO(), unifieddatalibrary.TrackrouteTupleParams{
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

func TestTrackrouteUnvalidatedPublishWithOptionalParams(t *testing.T) {
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
	err := client.Trackroute.UnvalidatedPublish(context.TODO(), unifieddatalibrary.TrackrouteUnvalidatedPublishParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.TrackrouteUnvalidatedPublishParamsDataModeTest,
		LastUpdateDate:        time.Now(),
		Source:                "Bluestaq",
		Type:                  "AIR REFUELING",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AltitudeBlocks: []unifieddatalibrary.TrackrouteUnvalidatedPublishParamsAltitudeBlock{{
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
		Poc: []unifieddatalibrary.TrackrouteUnvalidatedPublishParamsPoc{{
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
		RoutePoints: []unifieddatalibrary.TrackrouteUnvalidatedPublishParamsRoutePoint{{
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
