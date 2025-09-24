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

func TestLinkstatusUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Linkstatus.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.LinkstatusUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.LinkstatusUpdateParamsDataModeTest,
			EndPoint1Lat:          45.23,
			EndPoint1Lon:          80.23,
			EndPoint1Name:         "Example endpoint",
			EndPoint2Lat:          45.23,
			EndPoint2Lon:          80.23,
			EndPoint2Name:         "Example description",
			LinkName:              "Example description",
			LinkStartTime:         time.Now(),
			LinkStopTime:          time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("LINKSTATUS-ID"),
			Band:                  unifieddatalibrary.String("MIL-KA"),
			Constellation:         unifieddatalibrary.String("Fornax"),
			DataRate1To2:          unifieddatalibrary.Float(10.23),
			DataRate2To1:          unifieddatalibrary.Float(10.23),
			IDBeam1:               unifieddatalibrary.String("REF-BEAM1-ID"),
			IDBeam2:               unifieddatalibrary.String("REF-BEAM2-ID"),
			LinkState:             unifieddatalibrary.String("DEGRADED-WEATHER"),
			LinkType:              unifieddatalibrary.String("Example link"),
			OpsCap:                unifieddatalibrary.String("Example status"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			SatNo1:                unifieddatalibrary.Int(1),
			SatNo2:                unifieddatalibrary.Int(2),
			Snr:                   unifieddatalibrary.Float(10.1),
			SysCap:                unifieddatalibrary.String("Example status"),
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

func TestLinkstatusDelete(t *testing.T) {
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
	err := client.Linkstatus.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
