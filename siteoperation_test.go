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

func TestSiteOperationNewWithOptionalParams(t *testing.T) {
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
	err := client.Site.Operations.New(context.TODO(), unifieddatalibrary.SiteOperationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SiteOperationNewParamsDataModeTest,
		IDSite:                "a150b3ee-884b-b9ac-60a0-6408b4b16088",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
		DailyOperations: []unifieddatalibrary.SiteOperationNewParamsDailyOperation{{
			DayOfWeek: "MONDAY",
			OperatingHours: []unifieddatalibrary.SiteOperationNewParamsDailyOperationOperatingHour{{
				OpStartTime: unifieddatalibrary.String("12:00"),
				OpStopTime:  unifieddatalibrary.String("22:00"),
			}},
			OperationName:        unifieddatalibrary.String("Arrivals"),
			OphrsLastChangedBy:   unifieddatalibrary.String("John Smith"),
			OphrsLastChangedDate: unifieddatalibrary.Time(time.Now()),
		}},
		DopsLastChangedBy:     unifieddatalibrary.String("John Smith"),
		DopsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
		DopsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
		IDLaunchSite:          unifieddatalibrary.String("b150b3ee-884b-b9ac-60a0-6408b4b16088"),
		MaximumOnGrounds: []unifieddatalibrary.SiteOperationNewParamsMaximumOnGround{{
			AircraftMds:        unifieddatalibrary.String("C017A"),
			ContingencyMog:     unifieddatalibrary.Int(3),
			MogLastChangedBy:   unifieddatalibrary.String("John Smith"),
			MogLastChangedDate: unifieddatalibrary.Time(time.Now()),
			WideParkingMog:     unifieddatalibrary.Int(1),
			WideWorkingMog:     unifieddatalibrary.Int(1),
		}},
		MogsLastChangedBy:     unifieddatalibrary.String("Jane Doe"),
		MogsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
		MogsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
		OperationalDeviations: []unifieddatalibrary.SiteOperationNewParamsOperationalDeviation{{
			AffectedAircraftMds:  unifieddatalibrary.String("C017A"),
			AffectedMog:          unifieddatalibrary.Int(1),
			AircraftOnGroundTime: unifieddatalibrary.String("14:00"),
			CrewRestTime:         unifieddatalibrary.String("14:00"),
			OdLastChangedBy:      unifieddatalibrary.String("John Smith"),
			OdLastChangedDate:    unifieddatalibrary.Time(time.Now()),
			OdRemark:             unifieddatalibrary.String("Example remark about this operational deviation."),
		}},
		OperationalPlannings: []unifieddatalibrary.SiteOperationNewParamsOperationalPlanning{{
			OpEndDate:         unifieddatalibrary.Time(time.Now()),
			OpLastChangedBy:   unifieddatalibrary.String("John Smith"),
			OpLastChangedDate: unifieddatalibrary.Time(time.Now()),
			OpRemark:          unifieddatalibrary.String("Example planning remark"),
			OpSource:          unifieddatalibrary.String("a3"),
			OpStartDate:       unifieddatalibrary.Time(time.Now()),
			OpStatus:          unifieddatalibrary.String("Verified"),
		}},
		Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Pathways: []unifieddatalibrary.SiteOperationNewParamsPathway{{
			PwDefinition:      unifieddatalibrary.String("AGP: 14L, K6, K, G (ANG APRN TO TWY K), GUARD (MAIN)"),
			PwLastChangedBy:   unifieddatalibrary.String("John Smith"),
			PwLastChangedDate: unifieddatalibrary.Time(time.Now()),
			PwType:            unifieddatalibrary.String("Taxiway"),
			PwUsage:           unifieddatalibrary.String("Arrival"),
		}},
		Waivers: []unifieddatalibrary.SiteOperationNewParamsWaiver{{
			ExpirationDate:        unifieddatalibrary.Time(time.Now()),
			HasExpired:            unifieddatalibrary.Bool(false),
			IssueDate:             unifieddatalibrary.Time(time.Now()),
			IssuerName:            unifieddatalibrary.String("John Smith"),
			RequesterName:         unifieddatalibrary.String("Jane Doe"),
			RequesterPhoneNumber:  unifieddatalibrary.String("808-123-4567"),
			RequestingUnit:        unifieddatalibrary.String("2A1"),
			WaiverAppliesTo:       unifieddatalibrary.String("C017A"),
			WaiverDescription:     unifieddatalibrary.String("Example waiver description"),
			WaiverLastChangedBy:   unifieddatalibrary.String("J. Appleseed"),
			WaiverLastChangedDate: unifieddatalibrary.Time(time.Now()),
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

func TestSiteOperationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Operations.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteOperationGetParams{
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

func TestSiteOperationUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Site.Operations.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SiteOperationUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SiteOperationUpdateParamsDataModeTest,
			IDSite:                "a150b3ee-884b-b9ac-60a0-6408b4b16088",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			DailyOperations: []unifieddatalibrary.SiteOperationUpdateParamsDailyOperation{{
				DayOfWeek: "MONDAY",
				OperatingHours: []unifieddatalibrary.SiteOperationUpdateParamsDailyOperationOperatingHour{{
					OpStartTime: unifieddatalibrary.String("12:00"),
					OpStopTime:  unifieddatalibrary.String("22:00"),
				}},
				OperationName:        unifieddatalibrary.String("Arrivals"),
				OphrsLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OphrsLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
			DopsLastChangedBy:     unifieddatalibrary.String("John Smith"),
			DopsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			DopsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			IDLaunchSite:          unifieddatalibrary.String("b150b3ee-884b-b9ac-60a0-6408b4b16088"),
			MaximumOnGrounds: []unifieddatalibrary.SiteOperationUpdateParamsMaximumOnGround{{
				AircraftMds:        unifieddatalibrary.String("C017A"),
				ContingencyMog:     unifieddatalibrary.Int(3),
				MogLastChangedBy:   unifieddatalibrary.String("John Smith"),
				MogLastChangedDate: unifieddatalibrary.Time(time.Now()),
				WideParkingMog:     unifieddatalibrary.Int(1),
				WideWorkingMog:     unifieddatalibrary.Int(1),
			}},
			MogsLastChangedBy:     unifieddatalibrary.String("Jane Doe"),
			MogsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			MogsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			OperationalDeviations: []unifieddatalibrary.SiteOperationUpdateParamsOperationalDeviation{{
				AffectedAircraftMds:  unifieddatalibrary.String("C017A"),
				AffectedMog:          unifieddatalibrary.Int(1),
				AircraftOnGroundTime: unifieddatalibrary.String("14:00"),
				CrewRestTime:         unifieddatalibrary.String("14:00"),
				OdLastChangedBy:      unifieddatalibrary.String("John Smith"),
				OdLastChangedDate:    unifieddatalibrary.Time(time.Now()),
				OdRemark:             unifieddatalibrary.String("Example remark about this operational deviation."),
			}},
			OperationalPlannings: []unifieddatalibrary.SiteOperationUpdateParamsOperationalPlanning{{
				OpEndDate:         unifieddatalibrary.Time(time.Now()),
				OpLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OpLastChangedDate: unifieddatalibrary.Time(time.Now()),
				OpRemark:          unifieddatalibrary.String("Example planning remark"),
				OpSource:          unifieddatalibrary.String("a3"),
				OpStartDate:       unifieddatalibrary.Time(time.Now()),
				OpStatus:          unifieddatalibrary.String("Verified"),
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Pathways: []unifieddatalibrary.SiteOperationUpdateParamsPathway{{
				PwDefinition:      unifieddatalibrary.String("AGP: 14L, K6, K, G (ANG APRN TO TWY K), GUARD (MAIN)"),
				PwLastChangedBy:   unifieddatalibrary.String("John Smith"),
				PwLastChangedDate: unifieddatalibrary.Time(time.Now()),
				PwType:            unifieddatalibrary.String("Taxiway"),
				PwUsage:           unifieddatalibrary.String("Arrival"),
			}},
			Waivers: []unifieddatalibrary.SiteOperationUpdateParamsWaiver{{
				ExpirationDate:        unifieddatalibrary.Time(time.Now()),
				HasExpired:            unifieddatalibrary.Bool(false),
				IssueDate:             unifieddatalibrary.Time(time.Now()),
				IssuerName:            unifieddatalibrary.String("John Smith"),
				RequesterName:         unifieddatalibrary.String("Jane Doe"),
				RequesterPhoneNumber:  unifieddatalibrary.String("808-123-4567"),
				RequestingUnit:        unifieddatalibrary.String("2A1"),
				WaiverAppliesTo:       unifieddatalibrary.String("C017A"),
				WaiverDescription:     unifieddatalibrary.String("Example waiver description"),
				WaiverLastChangedBy:   unifieddatalibrary.String("J. Appleseed"),
				WaiverLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
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

func TestSiteOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Operations.List(context.TODO(), unifieddatalibrary.SiteOperationListParams{
		IDSite:      "idSite",
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

func TestSiteOperationDelete(t *testing.T) {
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
	err := client.Site.Operations.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteOperationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Operations.Count(context.TODO(), unifieddatalibrary.SiteOperationCountParams{
		IDSite:      "idSite",
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

func TestSiteOperationNewBulk(t *testing.T) {
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
	err := client.Site.Operations.NewBulk(context.TODO(), unifieddatalibrary.SiteOperationNewBulkParams{
		Body: []unifieddatalibrary.SiteOperationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSite:                "a150b3ee-884b-b9ac-60a0-6408b4b16088",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			DailyOperations: []unifieddatalibrary.SiteOperationNewBulkParamsBodyDailyOperation{{
				DayOfWeek: "MONDAY",
				OperatingHours: []unifieddatalibrary.SiteOperationNewBulkParamsBodyDailyOperationOperatingHour{{
					OpStartTime: unifieddatalibrary.String("12:00"),
					OpStopTime:  unifieddatalibrary.String("22:00"),
				}},
				OperationName:        unifieddatalibrary.String("Arrivals"),
				OphrsLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OphrsLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
			DopsLastChangedBy:     unifieddatalibrary.String("John Smith"),
			DopsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			DopsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			IDLaunchSite:          unifieddatalibrary.String("b150b3ee-884b-b9ac-60a0-6408b4b16088"),
			MaximumOnGrounds: []unifieddatalibrary.SiteOperationNewBulkParamsBodyMaximumOnGround{{
				AircraftMds:        unifieddatalibrary.String("C017A"),
				ContingencyMog:     unifieddatalibrary.Int(3),
				MogLastChangedBy:   unifieddatalibrary.String("John Smith"),
				MogLastChangedDate: unifieddatalibrary.Time(time.Now()),
				WideParkingMog:     unifieddatalibrary.Int(1),
				WideWorkingMog:     unifieddatalibrary.Int(1),
			}},
			MogsLastChangedBy:     unifieddatalibrary.String("Jane Doe"),
			MogsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			MogsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			OperationalDeviations: []unifieddatalibrary.SiteOperationNewBulkParamsBodyOperationalDeviation{{
				AffectedAircraftMds:  unifieddatalibrary.String("C017A"),
				AffectedMog:          unifieddatalibrary.Int(1),
				AircraftOnGroundTime: unifieddatalibrary.String("14:00"),
				CrewRestTime:         unifieddatalibrary.String("14:00"),
				OdLastChangedBy:      unifieddatalibrary.String("John Smith"),
				OdLastChangedDate:    unifieddatalibrary.Time(time.Now()),
				OdRemark:             unifieddatalibrary.String("Example remark about this operational deviation."),
			}},
			OperationalPlannings: []unifieddatalibrary.SiteOperationNewBulkParamsBodyOperationalPlanning{{
				OpEndDate:         unifieddatalibrary.Time(time.Now()),
				OpLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OpLastChangedDate: unifieddatalibrary.Time(time.Now()),
				OpRemark:          unifieddatalibrary.String("Example planning remark"),
				OpSource:          unifieddatalibrary.String("a3"),
				OpStartDate:       unifieddatalibrary.Time(time.Now()),
				OpStatus:          unifieddatalibrary.String("Verified"),
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Pathways: []unifieddatalibrary.SiteOperationNewBulkParamsBodyPathway{{
				PwDefinition:      unifieddatalibrary.String("AGP: 14L, K6, K, G (ANG APRN TO TWY K), GUARD (MAIN)"),
				PwLastChangedBy:   unifieddatalibrary.String("John Smith"),
				PwLastChangedDate: unifieddatalibrary.Time(time.Now()),
				PwType:            unifieddatalibrary.String("Taxiway"),
				PwUsage:           unifieddatalibrary.String("Arrival"),
			}},
			Waivers: []unifieddatalibrary.SiteOperationNewBulkParamsBodyWaiver{{
				ExpirationDate:        unifieddatalibrary.Time(time.Now()),
				HasExpired:            unifieddatalibrary.Bool(false),
				IssueDate:             unifieddatalibrary.Time(time.Now()),
				IssuerName:            unifieddatalibrary.String("John Smith"),
				RequesterName:         unifieddatalibrary.String("Jane Doe"),
				RequesterPhoneNumber:  unifieddatalibrary.String("808-123-4567"),
				RequestingUnit:        unifieddatalibrary.String("2A1"),
				WaiverAppliesTo:       unifieddatalibrary.String("C017A"),
				WaiverDescription:     unifieddatalibrary.String("Example waiver description"),
				WaiverLastChangedBy:   unifieddatalibrary.String("J. Appleseed"),
				WaiverLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
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

func TestSiteOperationQueryHelp(t *testing.T) {
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
	err := client.Site.Operations.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteOperationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Operations.Tuple(context.TODO(), unifieddatalibrary.SiteOperationTupleParams{
		Columns:     "columns",
		IDSite:      "idSite",
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

func TestSiteOperationUnvalidatedPublish(t *testing.T) {
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
	err := client.Site.Operations.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SiteOperationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSite:                "a150b3ee-884b-b9ac-60a0-6408b4b16088",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			DailyOperations: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyDailyOperation{{
				DayOfWeek: "MONDAY",
				OperatingHours: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyDailyOperationOperatingHour{{
					OpStartTime: unifieddatalibrary.String("12:00"),
					OpStopTime:  unifieddatalibrary.String("22:00"),
				}},
				OperationName:        unifieddatalibrary.String("Arrivals"),
				OphrsLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OphrsLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
			DopsLastChangedBy:     unifieddatalibrary.String("John Smith"),
			DopsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			DopsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			IDLaunchSite:          unifieddatalibrary.String("b150b3ee-884b-b9ac-60a0-6408b4b16088"),
			MaximumOnGrounds: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyMaximumOnGround{{
				AircraftMds:        unifieddatalibrary.String("C017A"),
				ContingencyMog:     unifieddatalibrary.Int(3),
				MogLastChangedBy:   unifieddatalibrary.String("John Smith"),
				MogLastChangedDate: unifieddatalibrary.Time(time.Now()),
				WideParkingMog:     unifieddatalibrary.Int(1),
				WideWorkingMog:     unifieddatalibrary.Int(1),
			}},
			MogsLastChangedBy:     unifieddatalibrary.String("Jane Doe"),
			MogsLastChangedDate:   unifieddatalibrary.Time(time.Now()),
			MogsLastChangedReason: unifieddatalibrary.String("Example reason for change."),
			OperationalDeviations: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyOperationalDeviation{{
				AffectedAircraftMds:  unifieddatalibrary.String("C017A"),
				AffectedMog:          unifieddatalibrary.Int(1),
				AircraftOnGroundTime: unifieddatalibrary.String("14:00"),
				CrewRestTime:         unifieddatalibrary.String("14:00"),
				OdLastChangedBy:      unifieddatalibrary.String("John Smith"),
				OdLastChangedDate:    unifieddatalibrary.Time(time.Now()),
				OdRemark:             unifieddatalibrary.String("Example remark about this operational deviation."),
			}},
			OperationalPlannings: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyOperationalPlanning{{
				OpEndDate:         unifieddatalibrary.Time(time.Now()),
				OpLastChangedBy:   unifieddatalibrary.String("John Smith"),
				OpLastChangedDate: unifieddatalibrary.Time(time.Now()),
				OpRemark:          unifieddatalibrary.String("Example planning remark"),
				OpSource:          unifieddatalibrary.String("a3"),
				OpStartDate:       unifieddatalibrary.Time(time.Now()),
				OpStatus:          unifieddatalibrary.String("Verified"),
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Pathways: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyPathway{{
				PwDefinition:      unifieddatalibrary.String("AGP: 14L, K6, K, G (ANG APRN TO TWY K), GUARD (MAIN)"),
				PwLastChangedBy:   unifieddatalibrary.String("John Smith"),
				PwLastChangedDate: unifieddatalibrary.Time(time.Now()),
				PwType:            unifieddatalibrary.String("Taxiway"),
				PwUsage:           unifieddatalibrary.String("Arrival"),
			}},
			Waivers: []unifieddatalibrary.SiteOperationUnvalidatedPublishParamsBodyWaiver{{
				ExpirationDate:        unifieddatalibrary.Time(time.Now()),
				HasExpired:            unifieddatalibrary.Bool(false),
				IssueDate:             unifieddatalibrary.Time(time.Now()),
				IssuerName:            unifieddatalibrary.String("John Smith"),
				RequesterName:         unifieddatalibrary.String("Jane Doe"),
				RequesterPhoneNumber:  unifieddatalibrary.String("808-123-4567"),
				RequestingUnit:        unifieddatalibrary.String("2A1"),
				WaiverAppliesTo:       unifieddatalibrary.String("C017A"),
				WaiverDescription:     unifieddatalibrary.String("Example waiver description"),
				WaiverLastChangedBy:   unifieddatalibrary.String("J. Appleseed"),
				WaiverLastChangedDate: unifieddatalibrary.Time(time.Now()),
			}},
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
