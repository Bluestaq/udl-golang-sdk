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

func TestSensorMaintenanceNewWithOptionalParams(t *testing.T) {
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
	err := client.SensorMaintenance.New(context.TODO(), unifieddatalibrary.SensorMaintenanceNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SensorMaintenanceNewParamsDataModeTest,
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
		ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
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

func TestSensorMaintenanceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SensorMaintenance.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorMaintenanceUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SensorMaintenanceUpdateParamsDataModeTest,
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
			ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
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

func TestSensorMaintenanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorMaintenance.List(context.TODO(), unifieddatalibrary.SensorMaintenanceListParams{
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

func TestSensorMaintenanceDelete(t *testing.T) {
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
	err := client.SensorMaintenance.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorMaintenanceCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorMaintenance.Count(context.TODO(), unifieddatalibrary.SensorMaintenanceCountParams{
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

func TestSensorMaintenanceNewBulkWithOptionalParams(t *testing.T) {
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
	err := client.SensorMaintenance.NewBulk(context.TODO(), unifieddatalibrary.SensorMaintenanceNewBulkParams{
		Body: []unifieddatalibrary.SensorMaintenanceNewBulkParamsBody{{
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
			ExternalID:            unifieddatalibrary.String("EXTERNAL-ID"),
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

func TestSensorMaintenanceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorMaintenance.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorMaintenanceGetParams{
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

func TestSensorMaintenanceListCurrentWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorMaintenance.ListCurrent(context.TODO(), unifieddatalibrary.SensorMaintenanceListCurrentParams{
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

func TestSensorMaintenanceQueryHelp(t *testing.T) {
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
	_, err := client.SensorMaintenance.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorMaintenanceTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SensorMaintenance.Tuple(context.TODO(), unifieddatalibrary.SensorMaintenanceTupleParams{
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
