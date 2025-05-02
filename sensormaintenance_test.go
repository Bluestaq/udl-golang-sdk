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

func TestSensormaintenanceNewWithOptionalParams(t *testing.T) {
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
	err := client.Sensormaintenance.New(context.TODO(), unifieddatalibrary.SensormaintenanceNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SensormaintenanceNewParamsDataModeTest,
		EndTime:               time.Now(),
		SiteCode:              "site01",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("SENSORMAINTENANCE-ID"),
		Activity:              unifieddatalibrary.String("Activity Description"),
		Approver:              unifieddatalibrary.String("approver"),
		Changer:               unifieddatalibrary.String("changer"),
		Duration:              unifieddatalibrary.String("128:16:52"),
		EowID:                 unifieddatalibrary.String("eowId"),
		EquipStatus:           unifieddatalibrary.String("FMC"),
		IDSensor:              unifieddatalibrary.String("idSensor"),
		ImpactedFaces:         unifieddatalibrary.String("impactedFaces"),
		LineNumber:            unifieddatalibrary.String("lineNumber"),
		MdOpsCap:              unifieddatalibrary.String("R"),
		MwOpsCap:              unifieddatalibrary.String("G"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Priority:              unifieddatalibrary.String("low"),
		Recall:                unifieddatalibrary.String("128:16:52"),
		Rel:                   unifieddatalibrary.String("rel"),
		Remark:                unifieddatalibrary.String("Remarks"),
		Requestor:             unifieddatalibrary.String("requestor"),
		Resource:              unifieddatalibrary.String("resource"),
		Rev:                   unifieddatalibrary.String("rev"),
		SSOpsCap:              unifieddatalibrary.String("Y"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sensormaintenance.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SensormaintenanceUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SensormaintenanceUpdateParamsDataModeTest,
			EndTime:               time.Now(),
			SiteCode:              "site01",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("SENSORMAINTENANCE-ID"),
			Activity:              unifieddatalibrary.String("Activity Description"),
			Approver:              unifieddatalibrary.String("approver"),
			Changer:               unifieddatalibrary.String("changer"),
			Duration:              unifieddatalibrary.String("128:16:52"),
			EowID:                 unifieddatalibrary.String("eowId"),
			EquipStatus:           unifieddatalibrary.String("FMC"),
			IDSensor:              unifieddatalibrary.String("idSensor"),
			ImpactedFaces:         unifieddatalibrary.String("impactedFaces"),
			LineNumber:            unifieddatalibrary.String("lineNumber"),
			MdOpsCap:              unifieddatalibrary.String("R"),
			MwOpsCap:              unifieddatalibrary.String("G"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Priority:              unifieddatalibrary.String("low"),
			Recall:                unifieddatalibrary.String("128:16:52"),
			Rel:                   unifieddatalibrary.String("rel"),
			Remark:                unifieddatalibrary.String("Remarks"),
			Requestor:             unifieddatalibrary.String("requestor"),
			Resource:              unifieddatalibrary.String("resource"),
			Rev:                   unifieddatalibrary.String("rev"),
			SSOpsCap:              unifieddatalibrary.String("Y"),
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

func TestSensormaintenanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensormaintenance.List(context.TODO(), unifieddatalibrary.SensormaintenanceListParams{
		EndTime:     unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceDelete(t *testing.T) {
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
	err := client.Sensormaintenance.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensormaintenance.Count(context.TODO(), unifieddatalibrary.SensormaintenanceCountParams{
		EndTime:     unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceNewBulkWithOptionalParams(t *testing.T) {
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
	err := client.Sensormaintenance.NewBulk(context.TODO(), unifieddatalibrary.SensormaintenanceNewBulkParams{
		Body: []unifieddatalibrary.SensormaintenanceNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EndTime:               time.Now(),
			SiteCode:              "site01",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("SENSORMAINTENANCE-ID"),
			Activity:              unifieddatalibrary.String("Activity Description"),
			Approver:              unifieddatalibrary.String("approver"),
			Changer:               unifieddatalibrary.String("changer"),
			Duration:              unifieddatalibrary.String("128:16:52"),
			EowID:                 unifieddatalibrary.String("eowId"),
			EquipStatus:           unifieddatalibrary.String("FMC"),
			IDSensor:              unifieddatalibrary.String("idSensor"),
			ImpactedFaces:         unifieddatalibrary.String("impactedFaces"),
			LineNumber:            unifieddatalibrary.String("lineNumber"),
			MdOpsCap:              unifieddatalibrary.String("R"),
			MwOpsCap:              unifieddatalibrary.String("G"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Priority:              unifieddatalibrary.String("low"),
			Recall:                unifieddatalibrary.String("128:16:52"),
			Rel:                   unifieddatalibrary.String("rel"),
			Remark:                unifieddatalibrary.String("Remarks"),
			Requestor:             unifieddatalibrary.String("requestor"),
			Resource:              unifieddatalibrary.String("resource"),
			Rev:                   unifieddatalibrary.String("rev"),
			SSOpsCap:              unifieddatalibrary.String("Y"),
		}},
		Origin: unifieddatalibrary.String("origin"),
		Source: unifieddatalibrary.String("source"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceCurrentWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensormaintenance.Current(context.TODO(), unifieddatalibrary.SensormaintenanceCurrentParams{
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

func TestSensormaintenanceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensormaintenance.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensormaintenanceGetParams{
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

func TestSensormaintenanceQueryhelp(t *testing.T) {
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
	err := client.Sensormaintenance.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensormaintenanceTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensormaintenance.Tuple(context.TODO(), unifieddatalibrary.SensormaintenanceTupleParams{
		Columns:     "columns",
		EndTime:     unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
