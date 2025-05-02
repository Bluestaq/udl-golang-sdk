// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestSeradatacommdetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradatacommdetails.New(context.TODO(), unifieddatalibrary.SeradatacommdetailNewParams{
		ClassificationMarking:                   "U",
		DataMode:                                unifieddatalibrary.SeradatacommdetailNewParamsDataModeTest,
		Source:                                  "Bluestaq",
		ID:                                      unifieddatalibrary.String("SERADATACOMMDETAILS-ID"),
		Band:                                    unifieddatalibrary.String("X"),
		Bandwidth:                               unifieddatalibrary.Float(1.23),
		Eirp:                                    unifieddatalibrary.Float(1.23),
		EstHtsTotalCapacity:                     unifieddatalibrary.Float(1.23),
		EstHtsTotalUserDownlinkBandwidthPerBeam: unifieddatalibrary.Float(1.23),
		EstHtsTotalUserUplinkBandwidthPerBeam:   unifieddatalibrary.Float(1.23),
		GatewayDownlinkFrom:                     unifieddatalibrary.Float(1.23),
		GatewayDownlinkTo:                       unifieddatalibrary.Float(1.23),
		GatewayUplinkFrom:                       unifieddatalibrary.Float(1.23),
		GatewayUplinkTo:                         unifieddatalibrary.Float(1.23),
		HostedForCompanyOrgID:                   unifieddatalibrary.String("hostedForCompanyOrgId"),
		HtsNumUserSpotBeams:                     unifieddatalibrary.Int(1),
		HtsUserDownlinkBandwidthPerBeam:         unifieddatalibrary.Float(1.23),
		HtsUserUplinkBandwidthPerBeam:           unifieddatalibrary.Float(1.23),
		IDComm:                                  unifieddatalibrary.String("idComm"),
		ManufacturerOrgID:                       unifieddatalibrary.String("manufacturerOrgId"),
		Num36MhzEquivalentTransponders:          unifieddatalibrary.Int(1),
		NumOperationalTransponders:              unifieddatalibrary.Int(1),
		NumSpareTransponders:                    unifieddatalibrary.Int(1),
		Origin:                                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PayloadNotes:                            unifieddatalibrary.String("Sample Notes"),
		Polarization:                            unifieddatalibrary.String("polarization"),
		SolidStatePowerAmp:                      unifieddatalibrary.Float(1.23),
		SpacecraftID:                            unifieddatalibrary.String("spacecraftId"),
		TradeLeaseOrgID:                         unifieddatalibrary.String("tradeLeaseOrgId"),
		TravelingWaveTubeAmplifier:              unifieddatalibrary.Float(1.23),
		UserDownlinkFrom:                        unifieddatalibrary.Float(1.23),
		UserDownlinkTo:                          unifieddatalibrary.Float(1.23),
		UserUplinkFrom:                          unifieddatalibrary.Float(1.23),
		UserUplinkTo:                            unifieddatalibrary.Float(1.23),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatacommdetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradatacommdetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatacommdetailUpdateParams{
			ClassificationMarking:                   "U",
			DataMode:                                unifieddatalibrary.SeradatacommdetailUpdateParamsDataModeTest,
			Source:                                  "Bluestaq",
			ID:                                      unifieddatalibrary.String("SERADATACOMMDETAILS-ID"),
			Band:                                    unifieddatalibrary.String("X"),
			Bandwidth:                               unifieddatalibrary.Float(1.23),
			Eirp:                                    unifieddatalibrary.Float(1.23),
			EstHtsTotalCapacity:                     unifieddatalibrary.Float(1.23),
			EstHtsTotalUserDownlinkBandwidthPerBeam: unifieddatalibrary.Float(1.23),
			EstHtsTotalUserUplinkBandwidthPerBeam:   unifieddatalibrary.Float(1.23),
			GatewayDownlinkFrom:                     unifieddatalibrary.Float(1.23),
			GatewayDownlinkTo:                       unifieddatalibrary.Float(1.23),
			GatewayUplinkFrom:                       unifieddatalibrary.Float(1.23),
			GatewayUplinkTo:                         unifieddatalibrary.Float(1.23),
			HostedForCompanyOrgID:                   unifieddatalibrary.String("hostedForCompanyOrgId"),
			HtsNumUserSpotBeams:                     unifieddatalibrary.Int(1),
			HtsUserDownlinkBandwidthPerBeam:         unifieddatalibrary.Float(1.23),
			HtsUserUplinkBandwidthPerBeam:           unifieddatalibrary.Float(1.23),
			IDComm:                                  unifieddatalibrary.String("idComm"),
			ManufacturerOrgID:                       unifieddatalibrary.String("manufacturerOrgId"),
			Num36MhzEquivalentTransponders:          unifieddatalibrary.Int(1),
			NumOperationalTransponders:              unifieddatalibrary.Int(1),
			NumSpareTransponders:                    unifieddatalibrary.Int(1),
			Origin:                                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PayloadNotes:                            unifieddatalibrary.String("Sample Notes"),
			Polarization:                            unifieddatalibrary.String("polarization"),
			SolidStatePowerAmp:                      unifieddatalibrary.Float(1.23),
			SpacecraftID:                            unifieddatalibrary.String("spacecraftId"),
			TradeLeaseOrgID:                         unifieddatalibrary.String("tradeLeaseOrgId"),
			TravelingWaveTubeAmplifier:              unifieddatalibrary.Float(1.23),
			UserDownlinkFrom:                        unifieddatalibrary.Float(1.23),
			UserDownlinkTo:                          unifieddatalibrary.Float(1.23),
			UserUplinkFrom:                          unifieddatalibrary.Float(1.23),
			UserUplinkTo:                            unifieddatalibrary.Float(1.23),
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

func TestSeradatacommdetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatacommdetails.List(context.TODO(), unifieddatalibrary.SeradatacommdetailListParams{
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

func TestSeradatacommdetailDelete(t *testing.T) {
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
	err := client.Seradatacommdetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatacommdetailCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatacommdetails.Count(context.TODO(), unifieddatalibrary.SeradatacommdetailCountParams{
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

func TestSeradatacommdetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatacommdetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradatacommdetailGetParams{
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

func TestSeradatacommdetailQueryhelp(t *testing.T) {
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
	err := client.Seradatacommdetails.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradatacommdetailTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradatacommdetails.Tuple(context.TODO(), unifieddatalibrary.SeradatacommdetailTupleParams{
		Columns:     "columns",
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
