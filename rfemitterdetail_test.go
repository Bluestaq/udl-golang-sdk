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

func TestRfEmitterDetailNewWithOptionalParams(t *testing.T) {
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
	err := client.RfEmitterDetails.New(context.TODO(), unifieddatalibrary.RfEmitterDetailNewParams{
		ClassificationMarking:        "U",
		DataMode:                     unifieddatalibrary.RfEmitterDetailNewParamsDataModeTest,
		IDRfEmitter:                  "RFEMITTER-ID",
		Source:                       "Bluestaq",
		ID:                           unifieddatalibrary.String("RFEMITTERDETAILS-ID"),
		AlternateFacilityName:        unifieddatalibrary.String("ALTERNATE_FACILITY_NAME"),
		AltName:                      unifieddatalibrary.String("ALTERNATE_NAME"),
		AntennaDiameter:              unifieddatalibrary.Float(20.23),
		AntennaSize:                  []float64{1.1, 2.2},
		BarrageNoiseBandwidth:        unifieddatalibrary.Float(5.23),
		Description:                  unifieddatalibrary.String("DESCRIPTION"),
		Designator:                   unifieddatalibrary.String("DESIGNATOR"),
		DopplerNoise:                 unifieddatalibrary.Float(10.23),
		DrfmInstantaneousBandwidth:   unifieddatalibrary.Float(20.23),
		Family:                       unifieddatalibrary.String("FAMILY"),
		ManufacturerOrgID:            unifieddatalibrary.String("MANUFACTURERORG-ID"),
		Notes:                        unifieddatalibrary.String("NOTES"),
		NumBits:                      unifieddatalibrary.Int(256),
		NumChannels:                  unifieddatalibrary.Int(10),
		Origin:                       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ProductionFacilityLocationID: unifieddatalibrary.String("PRODUCTIONFACILITYLOCATION-ID"),
		ProductionFacilityName:       unifieddatalibrary.String("PRODUCTION_FACILITY"),
		ReceiverBandwidth:            unifieddatalibrary.Float(15.23),
		ReceiverSensitivity:          unifieddatalibrary.Float(10.23),
		ReceiverType:                 unifieddatalibrary.String("RECEIVER_TYPE"),
		SecondaryNotes:               unifieddatalibrary.String("MORE_NOTES"),
		SystemSensitivityEnd:         unifieddatalibrary.Float(150.23),
		SystemSensitivityStart:       unifieddatalibrary.Float(50.23),
		TransmitPower:                unifieddatalibrary.Float(100.23),
		TransmitterBandwidth:         unifieddatalibrary.Float(0.125),
		TransmitterFrequency:         unifieddatalibrary.Float(105.9),
		URLs:                         []string{"TAG1", "TAG2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfEmitterDetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.RfEmitterDetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.RfEmitterDetailUpdateParams{
			ClassificationMarking:        "U",
			DataMode:                     unifieddatalibrary.RfEmitterDetailUpdateParamsDataModeTest,
			IDRfEmitter:                  "RFEMITTER-ID",
			Source:                       "Bluestaq",
			ID:                           unifieddatalibrary.String("RFEMITTERDETAILS-ID"),
			AlternateFacilityName:        unifieddatalibrary.String("ALTERNATE_FACILITY_NAME"),
			AltName:                      unifieddatalibrary.String("ALTERNATE_NAME"),
			AntennaDiameter:              unifieddatalibrary.Float(20.23),
			AntennaSize:                  []float64{1.1, 2.2},
			BarrageNoiseBandwidth:        unifieddatalibrary.Float(5.23),
			Description:                  unifieddatalibrary.String("DESCRIPTION"),
			Designator:                   unifieddatalibrary.String("DESIGNATOR"),
			DopplerNoise:                 unifieddatalibrary.Float(10.23),
			DrfmInstantaneousBandwidth:   unifieddatalibrary.Float(20.23),
			Family:                       unifieddatalibrary.String("FAMILY"),
			ManufacturerOrgID:            unifieddatalibrary.String("MANUFACTURERORG-ID"),
			Notes:                        unifieddatalibrary.String("NOTES"),
			NumBits:                      unifieddatalibrary.Int(256),
			NumChannels:                  unifieddatalibrary.Int(10),
			Origin:                       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ProductionFacilityLocationID: unifieddatalibrary.String("PRODUCTIONFACILITYLOCATION-ID"),
			ProductionFacilityName:       unifieddatalibrary.String("PRODUCTION_FACILITY"),
			ReceiverBandwidth:            unifieddatalibrary.Float(15.23),
			ReceiverSensitivity:          unifieddatalibrary.Float(10.23),
			ReceiverType:                 unifieddatalibrary.String("RECEIVER_TYPE"),
			SecondaryNotes:               unifieddatalibrary.String("MORE_NOTES"),
			SystemSensitivityEnd:         unifieddatalibrary.Float(150.23),
			SystemSensitivityStart:       unifieddatalibrary.Float(50.23),
			TransmitPower:                unifieddatalibrary.Float(100.23),
			TransmitterBandwidth:         unifieddatalibrary.Float(0.125),
			TransmitterFrequency:         unifieddatalibrary.Float(105.9),
			URLs:                         []string{"TAG1", "TAG2"},
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

func TestRfEmitterDetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitterDetails.List(context.TODO(), unifieddatalibrary.RfEmitterDetailListParams{
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

func TestRfEmitterDetailDelete(t *testing.T) {
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
	err := client.RfEmitterDetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfEmitterDetailCountWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitterDetails.Count(context.TODO(), unifieddatalibrary.RfEmitterDetailCountParams{
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

func TestRfEmitterDetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitterDetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.RfEmitterDetailGetParams{
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

func TestRfEmitterDetailQueryhelp(t *testing.T) {
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
	err := client.RfEmitterDetails.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRfEmitterDetailTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.RfEmitterDetails.Tuple(context.TODO(), unifieddatalibrary.RfEmitterDetailTupleParams{
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
