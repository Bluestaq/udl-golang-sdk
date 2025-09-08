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
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.RfEmitterDetailNewParamsDataModeTest,
		IDRfEmitter:           "RFEMITTER-ID",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
		AlternateFacilityName: unifieddatalibrary.String("ALTERNATE_FACILITY_NAME"),
		AltName:               unifieddatalibrary.String("ALTERNATE_NAME"),
		Amplifier: unifieddatalibrary.RfEmitterDetailNewParamsAmplifier{
			DeviceIdentifier: unifieddatalibrary.String("1200W_Amplifier_GE"),
			Manufacturer:     unifieddatalibrary.String("General Electric"),
			ModelName:        unifieddatalibrary.String("Model 2600"),
			Power:            unifieddatalibrary.Float(1200.1),
		},
		Antennas: []unifieddatalibrary.RfEmitterDetailNewParamsAntenna{{
			AntennaDiameter: unifieddatalibrary.Float(20.23),
			AntennaSize:     []float64{1.1, 2.2},
			AzElFixed:       unifieddatalibrary.Bool(true),
			Feeds: []unifieddatalibrary.RfEmitterDetailNewParamsAntennaFeed{{
				FreqMax:      unifieddatalibrary.Float(43500.1),
				FreqMin:      unifieddatalibrary.Float(39500.1),
				Name:         unifieddatalibrary.String("Feed A"),
				Polarization: unifieddatalibrary.String("HORIZONTAL"),
			}},
			FixedAzimuth:   unifieddatalibrary.Float(5.1),
			FixedElevation: unifieddatalibrary.Float(10.1),
			MaxAzimuths:    []float64{359.1, 359.1},
			MaxElevation:   unifieddatalibrary.Float(88.1),
			MinAzimuths:    []float64{5.1, 7.5},
			MinElevation:   unifieddatalibrary.Float(10.1),
			Name:           unifieddatalibrary.String("ds1Rotor"),
			ReceiverChannels: []unifieddatalibrary.RfEmitterDetailNewParamsAntennaReceiverChannel{{
				Bandwidth:        unifieddatalibrary.Float(0.51),
				ChannelNum:       unifieddatalibrary.String("315"),
				DeviceIdentifier: unifieddatalibrary.String("MMS1"),
				FreqMax:          unifieddatalibrary.Float(43500.1),
				FreqMin:          unifieddatalibrary.Float(39500.1),
				MaxBandwidth:     unifieddatalibrary.Float(100.1),
				MinBandwidth:     unifieddatalibrary.Float(0.05),
				Sensitivity:      unifieddatalibrary.Float(10.23),
			}},
			TransmitChannels: []unifieddatalibrary.RfEmitterDetailNewParamsAntennaTransmitChannel{{
				Power:              100.23,
				Bandwidth:          unifieddatalibrary.Float(0.125),
				ChannelNum:         unifieddatalibrary.String("12"),
				DeviceIdentifier:   unifieddatalibrary.String("TX1-B4-778"),
				Freq:               unifieddatalibrary.Float(42000.1),
				FreqMax:            unifieddatalibrary.Float(43500.1),
				FreqMin:            unifieddatalibrary.Float(39500.1),
				HardwareSampleRate: unifieddatalibrary.Int(192000),
				MaxBandwidth:       unifieddatalibrary.Float(100.1),
				MaxGain:            unifieddatalibrary.Float(20.1),
				MinBandwidth:       unifieddatalibrary.Float(0.05),
				MinGain:            unifieddatalibrary.Float(0.1),
				SampleRates:        []float64{192000, 9216000},
			}},
		}},
		BarrageNoiseBandwidth:        unifieddatalibrary.Float(5.23),
		BitRunTime:                   unifieddatalibrary.Float(120.1),
		Description:                  unifieddatalibrary.String("DESCRIPTION"),
		Designator:                   unifieddatalibrary.String("DESIGNATOR"),
		DopplerNoise:                 unifieddatalibrary.Float(10.23),
		DrfmInstantaneousBandwidth:   unifieddatalibrary.Float(20.23),
		Family:                       unifieddatalibrary.String("FAMILY"),
		FixedAttenuation:             unifieddatalibrary.Float(5.1),
		IDManufacturerOrg:            unifieddatalibrary.String("MANUFACTURERORG-ID"),
		IDProductionFacilityLocation: unifieddatalibrary.String("PRODUCTIONFACILITYLOCATION-ID"),
		LoanedToCocom:                unifieddatalibrary.String("SPACEFOR-STRATNORTH"),
		Notes:                        unifieddatalibrary.String("NOTES"),
		NumBits:                      unifieddatalibrary.Int(256),
		NumChannels:                  unifieddatalibrary.Int(10),
		Origin:                       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PowerOffsets: []unifieddatalibrary.RfEmitterDetailNewParamsPowerOffset{{
			FrequencyBand: unifieddatalibrary.String("KU"),
			PowerOffset:   unifieddatalibrary.Float(-5.1),
		}},
		PrepTime:               unifieddatalibrary.Float(60.1),
		PrimaryCocom:           unifieddatalibrary.String("SPACESOC"),
		ProductionFacilityName: unifieddatalibrary.String("PRODUCTION_FACILITY"),
		ReceiverType:           unifieddatalibrary.String("RECEIVER_TYPE"),
		SecondaryNotes:         unifieddatalibrary.String("MORE_NOTES"),
		Services: []unifieddatalibrary.RfEmitterDetailNewParamsService{{
			Name:    unifieddatalibrary.String("tlc-freqcontrol"),
			Version: unifieddatalibrary.String("17.2.5_build72199"),
		}},
		SystemSensitivityEnd:   unifieddatalibrary.Float(150.23),
		SystemSensitivityStart: unifieddatalibrary.Float(50.32),
		Ttps: []unifieddatalibrary.RfEmitterDetailNewParamsTtp{{
			OutputSignalName: unifieddatalibrary.String("SIGNAL_A"),
			TechniqueDefinitions: []unifieddatalibrary.RfEmitterDetailNewParamsTtpTechniqueDefinition{{
				Name: unifieddatalibrary.String("SIGNAL TUNE"),
				ParamDefinitions: []unifieddatalibrary.RfEmitterDetailNewParamsTtpTechniqueDefinitionParamDefinition{{
					DefaultValue: unifieddatalibrary.String("444.0"),
					Max:          unifieddatalibrary.Float(1000.1),
					Min:          unifieddatalibrary.Float(0.1),
					Name:         unifieddatalibrary.String("CENTER_FREQ"),
					Optional:     unifieddatalibrary.Bool(false),
					Type:         unifieddatalibrary.String("DOUBLE"),
					Units:        unifieddatalibrary.String("hertz"),
					ValidValues:  []string{"100.1", "444.1", "1000.1"},
				}},
			}},
		}},
		URLs: []string{"TAG1", "TAG2"},
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
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.RfEmitterDetailUpdateParamsDataModeTest,
			IDRfEmitter:           "RFEMITTER-ID",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("ad88770b-d824-443f-bdce-5f9e3fa500a9"),
			AlternateFacilityName: unifieddatalibrary.String("ALTERNATE_FACILITY_NAME"),
			AltName:               unifieddatalibrary.String("ALTERNATE_NAME"),
			Amplifier: unifieddatalibrary.RfEmitterDetailUpdateParamsAmplifier{
				DeviceIdentifier: unifieddatalibrary.String("1200W_Amplifier_GE"),
				Manufacturer:     unifieddatalibrary.String("General Electric"),
				ModelName:        unifieddatalibrary.String("Model 2600"),
				Power:            unifieddatalibrary.Float(1200.1),
			},
			Antennas: []unifieddatalibrary.RfEmitterDetailUpdateParamsAntenna{{
				AntennaDiameter: unifieddatalibrary.Float(20.23),
				AntennaSize:     []float64{1.1, 2.2},
				AzElFixed:       unifieddatalibrary.Bool(true),
				Feeds: []unifieddatalibrary.RfEmitterDetailUpdateParamsAntennaFeed{{
					FreqMax:      unifieddatalibrary.Float(43500.1),
					FreqMin:      unifieddatalibrary.Float(39500.1),
					Name:         unifieddatalibrary.String("Feed A"),
					Polarization: unifieddatalibrary.String("HORIZONTAL"),
				}},
				FixedAzimuth:   unifieddatalibrary.Float(5.1),
				FixedElevation: unifieddatalibrary.Float(10.1),
				MaxAzimuths:    []float64{359.1, 359.1},
				MaxElevation:   unifieddatalibrary.Float(88.1),
				MinAzimuths:    []float64{5.1, 7.5},
				MinElevation:   unifieddatalibrary.Float(10.1),
				Name:           unifieddatalibrary.String("ds1Rotor"),
				ReceiverChannels: []unifieddatalibrary.RfEmitterDetailUpdateParamsAntennaReceiverChannel{{
					Bandwidth:        unifieddatalibrary.Float(0.51),
					ChannelNum:       unifieddatalibrary.String("315"),
					DeviceIdentifier: unifieddatalibrary.String("MMS1"),
					FreqMax:          unifieddatalibrary.Float(43500.1),
					FreqMin:          unifieddatalibrary.Float(39500.1),
					MaxBandwidth:     unifieddatalibrary.Float(100.1),
					MinBandwidth:     unifieddatalibrary.Float(0.05),
					Sensitivity:      unifieddatalibrary.Float(10.23),
				}},
				TransmitChannels: []unifieddatalibrary.RfEmitterDetailUpdateParamsAntennaTransmitChannel{{
					Power:              100.23,
					Bandwidth:          unifieddatalibrary.Float(0.125),
					ChannelNum:         unifieddatalibrary.String("12"),
					DeviceIdentifier:   unifieddatalibrary.String("TX1-B4-778"),
					Freq:               unifieddatalibrary.Float(42000.1),
					FreqMax:            unifieddatalibrary.Float(43500.1),
					FreqMin:            unifieddatalibrary.Float(39500.1),
					HardwareSampleRate: unifieddatalibrary.Int(192000),
					MaxBandwidth:       unifieddatalibrary.Float(100.1),
					MaxGain:            unifieddatalibrary.Float(20.1),
					MinBandwidth:       unifieddatalibrary.Float(0.05),
					MinGain:            unifieddatalibrary.Float(0.1),
					SampleRates:        []float64{192000, 9216000},
				}},
			}},
			BarrageNoiseBandwidth:        unifieddatalibrary.Float(5.23),
			BitRunTime:                   unifieddatalibrary.Float(120.1),
			Description:                  unifieddatalibrary.String("DESCRIPTION"),
			Designator:                   unifieddatalibrary.String("DESIGNATOR"),
			DopplerNoise:                 unifieddatalibrary.Float(10.23),
			DrfmInstantaneousBandwidth:   unifieddatalibrary.Float(20.23),
			Family:                       unifieddatalibrary.String("FAMILY"),
			FixedAttenuation:             unifieddatalibrary.Float(5.1),
			IDManufacturerOrg:            unifieddatalibrary.String("MANUFACTURERORG-ID"),
			IDProductionFacilityLocation: unifieddatalibrary.String("PRODUCTIONFACILITYLOCATION-ID"),
			LoanedToCocom:                unifieddatalibrary.String("SPACEFOR-STRATNORTH"),
			Notes:                        unifieddatalibrary.String("NOTES"),
			NumBits:                      unifieddatalibrary.Int(256),
			NumChannels:                  unifieddatalibrary.Int(10),
			Origin:                       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PowerOffsets: []unifieddatalibrary.RfEmitterDetailUpdateParamsPowerOffset{{
				FrequencyBand: unifieddatalibrary.String("KU"),
				PowerOffset:   unifieddatalibrary.Float(-5.1),
			}},
			PrepTime:               unifieddatalibrary.Float(60.1),
			PrimaryCocom:           unifieddatalibrary.String("SPACESOC"),
			ProductionFacilityName: unifieddatalibrary.String("PRODUCTION_FACILITY"),
			ReceiverType:           unifieddatalibrary.String("RECEIVER_TYPE"),
			SecondaryNotes:         unifieddatalibrary.String("MORE_NOTES"),
			Services: []unifieddatalibrary.RfEmitterDetailUpdateParamsService{{
				Name:    unifieddatalibrary.String("tlc-freqcontrol"),
				Version: unifieddatalibrary.String("17.2.5_build72199"),
			}},
			SystemSensitivityEnd:   unifieddatalibrary.Float(150.23),
			SystemSensitivityStart: unifieddatalibrary.Float(50.32),
			Ttps: []unifieddatalibrary.RfEmitterDetailUpdateParamsTtp{{
				OutputSignalName: unifieddatalibrary.String("SIGNAL_A"),
				TechniqueDefinitions: []unifieddatalibrary.RfEmitterDetailUpdateParamsTtpTechniqueDefinition{{
					Name: unifieddatalibrary.String("SIGNAL TUNE"),
					ParamDefinitions: []unifieddatalibrary.RfEmitterDetailUpdateParamsTtpTechniqueDefinitionParamDefinition{{
						DefaultValue: unifieddatalibrary.String("444.0"),
						Max:          unifieddatalibrary.Float(1000.1),
						Min:          unifieddatalibrary.Float(0.1),
						Name:         unifieddatalibrary.String("CENTER_FREQ"),
						Optional:     unifieddatalibrary.Bool(false),
						Type:         unifieddatalibrary.String("DOUBLE"),
						Units:        unifieddatalibrary.String("hertz"),
						ValidValues:  []string{"100.1", "444.1", "1000.1"},
					}},
				}},
			}},
			URLs: []string{"TAG1", "TAG2"},
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
	_, err := client.RfEmitterDetails.Queryhelp(context.TODO())
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
