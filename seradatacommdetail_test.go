// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestSeraDataCommDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataCommDetails.New(context.TODO(), unifieddatalibrary.SeraDataCommDetailNewParams{
		ClassificationMarking:                   "U",
		DataMode:                                unifieddatalibrary.SeraDataCommDetailNewParamsDataModeTest,
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

func TestSeraDataCommDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SeraDataCommDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataCommDetailUpdateParams{
			ClassificationMarking:                   "U",
			DataMode:                                unifieddatalibrary.SeraDataCommDetailUpdateParamsDataModeTest,
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

func TestSeraDataCommDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataCommDetails.List(context.TODO(), unifieddatalibrary.SeraDataCommDetailListParams{
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

func TestSeraDataCommDetailDelete(t *testing.T) {
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
	err := client.SeraDataCommDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataCommDetailCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataCommDetails.Count(context.TODO(), unifieddatalibrary.SeraDataCommDetailCountParams{
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

func TestSeraDataCommDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataCommDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeraDataCommDetailGetParams{
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

func TestSeraDataCommDetailQueryhelp(t *testing.T) {
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
	_, err := client.SeraDataCommDetails.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeraDataCommDetailTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SeraDataCommDetails.Tuple(context.TODO(), unifieddatalibrary.SeraDataCommDetailTupleParams{
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
