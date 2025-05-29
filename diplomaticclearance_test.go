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

func TestDiplomaticClearanceNewWithOptionalParams(t *testing.T) {
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
	err := client.DiplomaticClearance.New(context.TODO(), unifieddatalibrary.DiplomaticClearanceNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.DiplomaticClearanceNewParamsDataModeTest,
		FirstDepDate:          time.Now(),
		IDMission:             "0dba1363-2d09-49fa-a784-4bb4cbb1674a",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
		ApacsID:               unifieddatalibrary.String("1083034"),
		DiplomaticClearanceDetails: []unifieddatalibrary.DiplomaticClearanceNewParamsDiplomaticClearanceDetail{{
			Action:              unifieddatalibrary.String("O"),
			AltCountryCode:      unifieddatalibrary.String("IZ"),
			ClearanceID:         unifieddatalibrary.String("MFMW225662GHQ"),
			ClearanceRemark:     unifieddatalibrary.String("Clearance remarks"),
			ClearedCallSign:     unifieddatalibrary.String("FALCN09"),
			CountryCode:         unifieddatalibrary.String("NL"),
			CountryName:         unifieddatalibrary.String("NETHERLANDS"),
			EntryNet:            unifieddatalibrary.Time(time.Now()),
			EntryPoint:          unifieddatalibrary.String("LOMOS"),
			ExitNlt:             unifieddatalibrary.Time(time.Now()),
			ExitPoint:           unifieddatalibrary.String("BUDOP"),
			ExternalClearanceID: unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
			IDSortie:            unifieddatalibrary.String("207010e0-f97d-431c-8c00-7e46acfef0f5"),
			LegNum:              unifieddatalibrary.Int(825),
			Profile:             unifieddatalibrary.String("T LAND/OFLY IATA COMPLIANT CARGO 23"),
			ReqIcao:             unifieddatalibrary.Bool(true),
			ReqPoint:            unifieddatalibrary.Bool(true),
			RouteString:         unifieddatalibrary.String("DCT DOH P430 BAYAN/M062F150 P430 RAMKI"),
			SequenceNum:         unifieddatalibrary.Int(3),
			Status:              unifieddatalibrary.String("IN WORK"),
			ValidDesc:           unifieddatalibrary.String("CY2023"),
			ValidEndTime:        unifieddatalibrary.Time(time.Now()),
			ValidStartTime:      unifieddatalibrary.Time(time.Now()),
			WindowRemark:        unifieddatalibrary.String("Period remarks"),
		}},
		DiplomaticClearanceRemarks: []unifieddatalibrary.DiplomaticClearanceNewParamsDiplomaticClearanceRemark{{
			Date:         unifieddatalibrary.Time(time.Now()),
			GdssRemarkID: unifieddatalibrary.String("GDSSREMARK-ID"),
			Text:         unifieddatalibrary.String("Example mission remarks."),
			User:         unifieddatalibrary.String("John Doe"),
		}},
		DipWorksheetName:    unifieddatalibrary.String("G2-939911-AC"),
		DocDeadline:         unifieddatalibrary.Time(time.Now()),
		ExternalWorksheetID: unifieddatalibrary.String("990ae849089e3d6cad69655324176bb6"),
		Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.DiplomaticClearanceGetParams{
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

func TestDiplomaticClearanceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.DiplomaticClearance.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.DiplomaticClearanceUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.DiplomaticClearanceUpdateParamsDataModeTest,
			FirstDepDate:          time.Now(),
			IDMission:             "0dba1363-2d09-49fa-a784-4bb4cbb1674a",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
			ApacsID:               unifieddatalibrary.String("1083034"),
			DiplomaticClearanceDetails: []unifieddatalibrary.DiplomaticClearanceUpdateParamsDiplomaticClearanceDetail{{
				Action:              unifieddatalibrary.String("O"),
				AltCountryCode:      unifieddatalibrary.String("IZ"),
				ClearanceID:         unifieddatalibrary.String("MFMW225662GHQ"),
				ClearanceRemark:     unifieddatalibrary.String("Clearance remarks"),
				ClearedCallSign:     unifieddatalibrary.String("FALCN09"),
				CountryCode:         unifieddatalibrary.String("NL"),
				CountryName:         unifieddatalibrary.String("NETHERLANDS"),
				EntryNet:            unifieddatalibrary.Time(time.Now()),
				EntryPoint:          unifieddatalibrary.String("LOMOS"),
				ExitNlt:             unifieddatalibrary.Time(time.Now()),
				ExitPoint:           unifieddatalibrary.String("BUDOP"),
				ExternalClearanceID: unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
				IDSortie:            unifieddatalibrary.String("207010e0-f97d-431c-8c00-7e46acfef0f5"),
				LegNum:              unifieddatalibrary.Int(825),
				Profile:             unifieddatalibrary.String("T LAND/OFLY IATA COMPLIANT CARGO 23"),
				ReqIcao:             unifieddatalibrary.Bool(true),
				ReqPoint:            unifieddatalibrary.Bool(true),
				RouteString:         unifieddatalibrary.String("DCT DOH P430 BAYAN/M062F150 P430 RAMKI"),
				SequenceNum:         unifieddatalibrary.Int(3),
				Status:              unifieddatalibrary.String("IN WORK"),
				ValidDesc:           unifieddatalibrary.String("CY2023"),
				ValidEndTime:        unifieddatalibrary.Time(time.Now()),
				ValidStartTime:      unifieddatalibrary.Time(time.Now()),
				WindowRemark:        unifieddatalibrary.String("Period remarks"),
			}},
			DiplomaticClearanceRemarks: []unifieddatalibrary.DiplomaticClearanceUpdateParamsDiplomaticClearanceRemark{{
				Date:         unifieddatalibrary.Time(time.Now()),
				GdssRemarkID: unifieddatalibrary.String("GDSSREMARK-ID"),
				Text:         unifieddatalibrary.String("Example mission remarks."),
				User:         unifieddatalibrary.String("John Doe"),
			}},
			DipWorksheetName:    unifieddatalibrary.String("G2-939911-AC"),
			DocDeadline:         unifieddatalibrary.Time(time.Now()),
			ExternalWorksheetID: unifieddatalibrary.String("990ae849089e3d6cad69655324176bb6"),
			Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestDiplomaticClearanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.List(context.TODO(), unifieddatalibrary.DiplomaticClearanceListParams{
		FirstDepDate: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceDelete(t *testing.T) {
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
	err := client.DiplomaticClearance.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceCountWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Count(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountParams{
		FirstDepDate: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceNewBulk(t *testing.T) {
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
	err := client.DiplomaticClearance.NewBulk(context.TODO(), unifieddatalibrary.DiplomaticClearanceNewBulkParams{
		Body: []unifieddatalibrary.DiplomaticClearanceNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			FirstDepDate:          time.Now(),
			IDMission:             "0dba1363-2d09-49fa-a784-4bb4cbb1674a",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
			ApacsID:               unifieddatalibrary.String("1083034"),
			DiplomaticClearanceDetails: []unifieddatalibrary.DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceDetail{{
				Action:              unifieddatalibrary.String("O"),
				AltCountryCode:      unifieddatalibrary.String("IZ"),
				ClearanceID:         unifieddatalibrary.String("MFMW225662GHQ"),
				ClearanceRemark:     unifieddatalibrary.String("Clearance remarks"),
				ClearedCallSign:     unifieddatalibrary.String("FALCN09"),
				CountryCode:         unifieddatalibrary.String("NL"),
				CountryName:         unifieddatalibrary.String("NETHERLANDS"),
				EntryNet:            unifieddatalibrary.Time(time.Now()),
				EntryPoint:          unifieddatalibrary.String("LOMOS"),
				ExitNlt:             unifieddatalibrary.Time(time.Now()),
				ExitPoint:           unifieddatalibrary.String("BUDOP"),
				ExternalClearanceID: unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
				IDSortie:            unifieddatalibrary.String("207010e0-f97d-431c-8c00-7e46acfef0f5"),
				LegNum:              unifieddatalibrary.Int(825),
				Profile:             unifieddatalibrary.String("T LAND/OFLY IATA COMPLIANT CARGO 23"),
				ReqIcao:             unifieddatalibrary.Bool(true),
				ReqPoint:            unifieddatalibrary.Bool(true),
				RouteString:         unifieddatalibrary.String("DCT DOH P430 BAYAN/M062F150 P430 RAMKI"),
				SequenceNum:         unifieddatalibrary.Int(3),
				Status:              unifieddatalibrary.String("IN WORK"),
				ValidDesc:           unifieddatalibrary.String("CY2023"),
				ValidEndTime:        unifieddatalibrary.Time(time.Now()),
				ValidStartTime:      unifieddatalibrary.Time(time.Now()),
				WindowRemark:        unifieddatalibrary.String("Period remarks"),
			}},
			DiplomaticClearanceRemarks: []unifieddatalibrary.DiplomaticClearanceNewBulkParamsBodyDiplomaticClearanceRemark{{
				Date:         unifieddatalibrary.Time(time.Now()),
				GdssRemarkID: unifieddatalibrary.String("GDSSREMARK-ID"),
				Text:         unifieddatalibrary.String("Example mission remarks."),
				User:         unifieddatalibrary.String("John Doe"),
			}},
			DipWorksheetName:    unifieddatalibrary.String("G2-939911-AC"),
			DocDeadline:         unifieddatalibrary.Time(time.Now()),
			ExternalWorksheetID: unifieddatalibrary.String("990ae849089e3d6cad69655324176bb6"),
			Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestDiplomaticClearanceQueryhelp(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Tuple(context.TODO(), unifieddatalibrary.DiplomaticClearanceTupleParams{
		Columns:      "columns",
		FirstDepDate: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
