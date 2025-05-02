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

func TestAircraftstatusremarkUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Aircraftstatusremark.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.AircraftstatusremarkUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.AircraftstatusremarkUpdateParamsDataModeTest,
			IDAircraftStatus:      "388b1f64-ccff-4113-b049-3cf5542c2a42",
			Source:                "Bluestaq",
			Text:                  "Remark text",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AltRmkID:              unifieddatalibrary.String("GDSSBL022307131714250077"),
			LastUpdatedAt:         unifieddatalibrary.Time(time.Now()),
			LastUpdatedBy:         unifieddatalibrary.String("JOHN SMITH"),
			Name:                  unifieddatalibrary.String("DISCREPANCY - 202297501"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Timestamp:             unifieddatalibrary.Time(time.Now()),
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

func TestAircraftstatusremarkDelete(t *testing.T) {
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
	err := client.Aircraftstatusremark.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
