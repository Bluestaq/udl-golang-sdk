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

func TestSensorNewWithOptionalParams(t *testing.T) {
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
	err := client.Sensor.New(context.TODO(), unifieddatalibrary.SensorNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SensorNewParamsDataModeTest,
		SensorName:            "SENSOR_NAME",
		Source:                "some.user",
		Active:                unifieddatalibrary.Bool(true),
		AfID:                  unifieddatalibrary.String("AF-ID"),
		AsrType:               unifieddatalibrary.String("SENSOR_TYPE"),
		DataControl:           unifieddatalibrary.String("observations"),
		Entity: unifieddatalibrary.SensorNewParamsEntity{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Name:                  "Example name",
			Source:                "Bluestaq",
			Type:                  "ONORBIT",
			CountryCode:           unifieddatalibrary.String("US"),
			IDEntity:              unifieddatalibrary.String("ENTITY-ID"),
			IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
			IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
			IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
			Location: unifieddatalibrary.LocationIngestParam{
				ClassificationMarking: "U",
				DataMode:              unifieddatalibrary.LocationIngestDataModeTest,
				Name:                  "Example location",
				Source:                "Bluestaq",
				Altitude:              unifieddatalibrary.Float(10.23),
				CountryCode:           unifieddatalibrary.String("US"),
				IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
				Lat:                   unifieddatalibrary.Float(45.23),
				Lon:                   unifieddatalibrary.Float(179.1),
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			OnOrbit: unifieddatalibrary.SensorNewParamsEntityOnOrbit{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				SatNo:                 1,
				Source:                "Bluestaq",
				AltName:               unifieddatalibrary.String("Alternate Name"),
				Category:              "Lunar",
				CommonName:            unifieddatalibrary.String("Example common name"),
				Constellation:         unifieddatalibrary.String("Big Dipper"),
				CountryCode:           unifieddatalibrary.String("US"),
				DecayDate:             unifieddatalibrary.Time(time.Now()),
				IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
				IntlDes:               unifieddatalibrary.String("2021123ABC"),
				LaunchDate:            unifieddatalibrary.Time(time.Now()),
				LaunchSiteID:          unifieddatalibrary.String("LAUNCHSITE-ID"),
				LifetimeYears:         unifieddatalibrary.Int(10),
				MissionNumber:         unifieddatalibrary.String("Expedition 1"),
				ObjectType:            "PAYLOAD",
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			},
			Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OwnerType:     "Commercial",
			Taskable:      unifieddatalibrary.Bool(false),
			TerrestrialID: unifieddatalibrary.String("TERRESTRIAL-ID"),
			URLs:          []string{"URL1", "URL2"},
		},
		IDEntity: unifieddatalibrary.String("ENTITY-ID"),
		IDSensor: unifieddatalibrary.String("SENSOR-ID"),
		Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Sensorcharacteristics: []unifieddatalibrary.SensorNewParamsSensorcharacteristic{{
			ClassificationMarking:           "U",
			DataMode:                        "TEST",
			IDSensor:                        "SENSOR-ID",
			Source:                          "Bluestaq",
			ID:                              unifieddatalibrary.String("SENSORCHARACTERISTICS-ID"),
			AcceptSampleRanges:              []float64{3.01, 3.02},
			AnalogToDigitalBitSize:          unifieddatalibrary.Int(2),
			Aperture:                        unifieddatalibrary.Float(2.23),
			AsrScanRate:                     unifieddatalibrary.Float(20.23),
			AtmosReceiverLoss:               unifieddatalibrary.Float(0.5),
			AtmosTransmissionLoss:           unifieddatalibrary.Float(0.5),
			AvgAtmosSeeingConditions:        unifieddatalibrary.Float(10.23),
			AzAngs:                          []float64{135.1, 45.2},
			AzimuthRate:                     unifieddatalibrary.Float(0.3334),
			BackgroundSkyRadiance:           unifieddatalibrary.Float(10.23),
			BackgroundSkyVisMag:             unifieddatalibrary.Float(10.23),
			Band:                            unifieddatalibrary.String("BAND"),
			Bandwidth:                       unifieddatalibrary.Float(100.23),
			BeamOrder:                       []string{"vb1", "ob1"},
			BeamQty:                         unifieddatalibrary.Int(2),
			Boresight:                       unifieddatalibrary.Float(20.23),
			BoresightOffAngle:               unifieddatalibrary.Float(20.23),
			CenterWavelength:                unifieddatalibrary.Float(4.56),
			CollapsingLoss:                  unifieddatalibrary.Float(1.23),
			CritShear:                       unifieddatalibrary.Float(47.1),
			DarkCurrent:                     unifieddatalibrary.Float(12.3),
			DelayGates:                      []float64{690.2, 690.3},
			Description:                     unifieddatalibrary.String("PROFILER DATA - PROFILE/SOUNDER DATA FROM PRIMARY WINDS SOURCE"),
			DetectSnr:                       unifieddatalibrary.Float(2.1),
			DutyCycle:                       unifieddatalibrary.Float(0.5),
			EarthLimbExclHgt:                unifieddatalibrary.Float(20.23),
			ElAngs:                          []float64{75.3, 75.4},
			ElevationRateGeolm:              unifieddatalibrary.Float(0.9555),
			EquipmentType:                   unifieddatalibrary.String("PS"),
			FanBeamWidth:                    unifieddatalibrary.Float(3.1),
			Fft:                             unifieddatalibrary.Int(4096),
			FgpCrit:                         unifieddatalibrary.Int(5),
			FilterMismatchFactor:            unifieddatalibrary.Float(10.23),
			FNum:                            unifieddatalibrary.Float(1.23),
			FocalPoint:                      unifieddatalibrary.Float(20.23),
			HFov:                            unifieddatalibrary.Float(20.23),
			HResPixels:                      unifieddatalibrary.Int(1000),
			K:                               unifieddatalibrary.Float(1.4),
			LeftClockAngle:                  unifieddatalibrary.Float(20.23),
			LeftGeoBeltLimit:                unifieddatalibrary.Float(20.23),
			Location:                        unifieddatalibrary.String("KENNEDY SPACE CENTER, FL"),
			LoopGain:                        unifieddatalibrary.Float(150.1),
			LunarExclAngle:                  unifieddatalibrary.Float(45.2),
			MagDec:                          unifieddatalibrary.Float(45.23),
			MagnitudeLimit:                  unifieddatalibrary.Float(23.5),
			MaxDeviationAngle:               unifieddatalibrary.Float(20.23),
			MaxIntegrationTime:              unifieddatalibrary.Float(1.1),
			MaxObservableRange:              unifieddatalibrary.Float(20.23),
			MaxRangeLimit:                   unifieddatalibrary.Float(4972.1),
			MaxWavelength:                   unifieddatalibrary.Float(7.89),
			MinIntegrationTime:              unifieddatalibrary.Float(0.3),
			MinRangeLimit:                   unifieddatalibrary.Float(165.1),
			MinSignalNoiseRatio:             unifieddatalibrary.Float(31.5),
			MinWavelength:                   unifieddatalibrary.Float(1.23),
			NegativeRangeRateLimit:          unifieddatalibrary.Float(-19.25),
			NoiseFigure:                     unifieddatalibrary.Float(10.23),
			NonCoherentIntegratedPulses:     unifieddatalibrary.Int(2),
			NumIntegratedPulses:             unifieddatalibrary.Int(10),
			NumIntegrationFrames:            unifieddatalibrary.Int(2),
			NumOpticalIntegrationModes:      unifieddatalibrary.Int(2),
			NumWaveforms:                    unifieddatalibrary.Int(2),
			OpticalIntegrationAngularRates:  []float64{15.1, 0.1},
			OpticalIntegrationFrames:        []float64{2.1, 3.1},
			OpticalIntegrationPixelBinnings: []float64{2.1, 1.1},
			OpticalIntegrationSnRs:          []float64{6.1, 6.1},
			OpticalIntegrationTimes:         []float64{0.3, 1},
			OpticalTransmission:             unifieddatalibrary.Float(0.5),
			PatternAbsorptionLoss:           unifieddatalibrary.Float(1.23),
			PatternScanLoss:                 unifieddatalibrary.Float(1.23),
			PeakPower:                       unifieddatalibrary.Float(5000000.1),
			PixelInstantaneousFov:           unifieddatalibrary.Float(10.23),
			PixelWellDepth:                  unifieddatalibrary.Int(12),
			PositiveRangeRateLimit:          unifieddatalibrary.Float(19.25),
			Prf:                             unifieddatalibrary.Float(20.23),
			ProbDetectSnr:                   unifieddatalibrary.Float(0.5),
			ProbFalseAlarm:                  unifieddatalibrary.Float(0.5),
			PulseRepPeriods:                 []float64{153.8, 153.9},
			QuantumEff:                      unifieddatalibrary.Float(0.5),
			RadarFrequency:                  unifieddatalibrary.Float(45300000000.1),
			RadarMessageFormat:              unifieddatalibrary.String("DATA_FORMAT"),
			RadarMur:                        unifieddatalibrary.Float(20.23),
			RadarPulseWidths:                []float64{20.23, 20.33},
			RadioFrequency:                  unifieddatalibrary.Float(20.23),
			RadomeLoss:                      unifieddatalibrary.Float(1.23),
			RangeGates:                      []int64{51, 52},
			RangeSpacings:                   []float64{690.2, 690.3},
			ReadNoise:                       unifieddatalibrary.Int(12),
			ReceiveGain:                     unifieddatalibrary.Float(10.2),
			ReceiveHorizBeamWidth:           unifieddatalibrary.Float(75.3),
			ReceiveLoss:                     unifieddatalibrary.Float(1.23),
			ReceiveVertBeamWidth:            unifieddatalibrary.Float(75.4),
			RefTemp:                         unifieddatalibrary.Float(3.5),
			ReqRecords:                      []int64{0, 1},
			RightClockAngle:                 unifieddatalibrary.Float(20.23),
			RightGeoBeltLimit:               unifieddatalibrary.Float(20.23),
			RunMeanCodes:                    []int64{0, 5},
			SignalProcessingLoss:            unifieddatalibrary.Float(1.23),
			SiteCode:                        unifieddatalibrary.String("07"),
			SolarExclAngle:                  unifieddatalibrary.Float(50.5),
			SpecAvgSpectraNums:              []int64{3, 4},
			SystemNoiseTemperature:          unifieddatalibrary.Float(3.5),
			TaskableRange:                   unifieddatalibrary.Float(20.23),
			TempMedFiltCodes:                []int64{3, 4},
			TestNumber:                      unifieddatalibrary.String("02022"),
			TotRecNums:                      []int64{5, 2},
			TowerHeight:                     unifieddatalibrary.Float(20.23),
			TrackAngle:                      unifieddatalibrary.Float(0.043),
			TrackSnr:                        unifieddatalibrary.Float(15.1),
			TransmitGain:                    unifieddatalibrary.Float(32.1),
			TransmitHorizBeamWidth:          unifieddatalibrary.Float(135.1),
			TransmitLoss:                    unifieddatalibrary.Float(7.1),
			TransmitPower:                   unifieddatalibrary.Float(190000.1),
			TransmitVertBeamWidth:           unifieddatalibrary.Float(45.2),
			TrueNorthCorrector:              unifieddatalibrary.Int(10),
			TrueTilt:                        unifieddatalibrary.Float(20.23),
			TwilightAngle:                   unifieddatalibrary.Float(7.5),
			VertBeamFlag:                    unifieddatalibrary.Bool(false),
			VertGateSpacings:                []float64{149.1, 149.2},
			VertGateWidths:                  []float64{149.1, 149.2},
			VFov:                            unifieddatalibrary.Float(20.23),
			VResPixels:                      unifieddatalibrary.Int(1000),
			WaveformBandwidths:              []float64{0.25, 0.25},
			WaveformLoopGains:               []float64{150.1, 155.1},
			WaveformMaxRanges:               []float64{2000.1, 2000.1},
			WaveformMinRanges:               []float64{150.1, 150.1},
			WaveformPulseWidths:             []float64{100.1, 200.1},
			Z1MaxRange:                      unifieddatalibrary.Float(50.23),
			Z1MinRange:                      unifieddatalibrary.Float(20.23),
			Z2MaxRange:                      unifieddatalibrary.Float(50.23),
			Z2MinRange:                      unifieddatalibrary.Float(20.23),
		}},
		SensorlimitsCollection: []unifieddatalibrary.SensorNewParamsSensorlimitsCollection{{
			ClassificationMarking:    "U",
			DataMode:                 "TEST",
			Source:                   "Bluestaq",
			IDSensor:                 unifieddatalibrary.String("SENSORLIMITS-ID"),
			IDSensorLimits:           unifieddatalibrary.String("SENSORLIMITS-ID"),
			LowerLeftAzimuthLimit:    unifieddatalibrary.Float(1.23),
			LowerLeftElevationLimit:  unifieddatalibrary.Float(1.23),
			LowerRightAzimuthLimit:   unifieddatalibrary.Float(1.23),
			LowerRightElevationLimit: unifieddatalibrary.Float(1.23),
			UpperLeftAzimuthLimit:    unifieddatalibrary.Float(1.23),
			UpperLeftElevationLimit:  unifieddatalibrary.Float(1.23),
			UpperRightAzimuthLimit:   unifieddatalibrary.Float(1.23),
			UpperRightElevationLimit: unifieddatalibrary.Float(1.23),
		}},
		SensorNumber: unifieddatalibrary.Int(1234),
		SensorObservationType: unifieddatalibrary.SensorNewParamsSensorObservationType{
			ID:   unifieddatalibrary.String("3"),
			Type: unifieddatalibrary.String("5"),
		},
		SensorStats: []unifieddatalibrary.SensorNewParamsSensorStat{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSensor:              "idSensor",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SENSORSTATS-ID"),
			LastObTime:            unifieddatalibrary.Time(time.Now()),
		}},
		SensorType: unifieddatalibrary.SensorNewParamsSensorType{
			ID:   unifieddatalibrary.Int(12344411),
			Type: unifieddatalibrary.String("Space Borne"),
		},
		ShortName: unifieddatalibrary.String("SNR-1"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sensor.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SensorUpdateParamsDataModeTest,
			SensorName:            "SENSOR_NAME",
			Source:                "some.user",
			Active:                unifieddatalibrary.Bool(true),
			AfID:                  unifieddatalibrary.String("AF-ID"),
			AsrType:               unifieddatalibrary.String("SENSOR_TYPE"),
			DataControl:           unifieddatalibrary.String("observations"),
			Entity: unifieddatalibrary.SensorUpdateParamsEntity{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Name:                  "Example name",
				Source:                "Bluestaq",
				Type:                  "ONORBIT",
				CountryCode:           unifieddatalibrary.String("US"),
				IDEntity:              unifieddatalibrary.String("ENTITY-ID"),
				IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
				IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
				IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
				Location: unifieddatalibrary.LocationIngestParam{
					ClassificationMarking: "U",
					DataMode:              unifieddatalibrary.LocationIngestDataModeTest,
					Name:                  "Example location",
					Source:                "Bluestaq",
					Altitude:              unifieddatalibrary.Float(10.23),
					CountryCode:           unifieddatalibrary.String("US"),
					IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
					Lat:                   unifieddatalibrary.Float(45.23),
					Lon:                   unifieddatalibrary.Float(179.1),
					Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				},
				OnOrbit: unifieddatalibrary.SensorUpdateParamsEntityOnOrbit{
					ClassificationMarking: "U",
					DataMode:              "TEST",
					SatNo:                 1,
					Source:                "Bluestaq",
					AltName:               unifieddatalibrary.String("Alternate Name"),
					Category:              "Lunar",
					CommonName:            unifieddatalibrary.String("Example common name"),
					Constellation:         unifieddatalibrary.String("Big Dipper"),
					CountryCode:           unifieddatalibrary.String("US"),
					DecayDate:             unifieddatalibrary.Time(time.Now()),
					IDOnOrbit:             unifieddatalibrary.String("ONORBIT-ID"),
					IntlDes:               unifieddatalibrary.String("2021123ABC"),
					LaunchDate:            unifieddatalibrary.Time(time.Now()),
					LaunchSiteID:          unifieddatalibrary.String("LAUNCHSITE-ID"),
					LifetimeYears:         unifieddatalibrary.Int(10),
					MissionNumber:         unifieddatalibrary.String("Expedition 1"),
					ObjectType:            "PAYLOAD",
					Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				},
				Origin:        unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				OwnerType:     "Commercial",
				Taskable:      unifieddatalibrary.Bool(false),
				TerrestrialID: unifieddatalibrary.String("TERRESTRIAL-ID"),
				URLs:          []string{"URL1", "URL2"},
			},
			IDEntity: unifieddatalibrary.String("ENTITY-ID"),
			IDSensor: unifieddatalibrary.String("SENSOR-ID"),
			Origin:   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Sensorcharacteristics: []unifieddatalibrary.SensorUpdateParamsSensorcharacteristic{{
				ClassificationMarking:           "U",
				DataMode:                        "TEST",
				IDSensor:                        "SENSOR-ID",
				Source:                          "Bluestaq",
				ID:                              unifieddatalibrary.String("SENSORCHARACTERISTICS-ID"),
				AcceptSampleRanges:              []float64{3.01, 3.02},
				AnalogToDigitalBitSize:          unifieddatalibrary.Int(2),
				Aperture:                        unifieddatalibrary.Float(2.23),
				AsrScanRate:                     unifieddatalibrary.Float(20.23),
				AtmosReceiverLoss:               unifieddatalibrary.Float(0.5),
				AtmosTransmissionLoss:           unifieddatalibrary.Float(0.5),
				AvgAtmosSeeingConditions:        unifieddatalibrary.Float(10.23),
				AzAngs:                          []float64{135.1, 45.2},
				AzimuthRate:                     unifieddatalibrary.Float(0.3334),
				BackgroundSkyRadiance:           unifieddatalibrary.Float(10.23),
				BackgroundSkyVisMag:             unifieddatalibrary.Float(10.23),
				Band:                            unifieddatalibrary.String("BAND"),
				Bandwidth:                       unifieddatalibrary.Float(100.23),
				BeamOrder:                       []string{"vb1", "ob1"},
				BeamQty:                         unifieddatalibrary.Int(2),
				Boresight:                       unifieddatalibrary.Float(20.23),
				BoresightOffAngle:               unifieddatalibrary.Float(20.23),
				CenterWavelength:                unifieddatalibrary.Float(4.56),
				CollapsingLoss:                  unifieddatalibrary.Float(1.23),
				CritShear:                       unifieddatalibrary.Float(47.1),
				DarkCurrent:                     unifieddatalibrary.Float(12.3),
				DelayGates:                      []float64{690.2, 690.3},
				Description:                     unifieddatalibrary.String("PROFILER DATA - PROFILE/SOUNDER DATA FROM PRIMARY WINDS SOURCE"),
				DetectSnr:                       unifieddatalibrary.Float(2.1),
				DutyCycle:                       unifieddatalibrary.Float(0.5),
				EarthLimbExclHgt:                unifieddatalibrary.Float(20.23),
				ElAngs:                          []float64{75.3, 75.4},
				ElevationRateGeolm:              unifieddatalibrary.Float(0.9555),
				EquipmentType:                   unifieddatalibrary.String("PS"),
				FanBeamWidth:                    unifieddatalibrary.Float(3.1),
				Fft:                             unifieddatalibrary.Int(4096),
				FgpCrit:                         unifieddatalibrary.Int(5),
				FilterMismatchFactor:            unifieddatalibrary.Float(10.23),
				FNum:                            unifieddatalibrary.Float(1.23),
				FocalPoint:                      unifieddatalibrary.Float(20.23),
				HFov:                            unifieddatalibrary.Float(20.23),
				HResPixels:                      unifieddatalibrary.Int(1000),
				K:                               unifieddatalibrary.Float(1.4),
				LeftClockAngle:                  unifieddatalibrary.Float(20.23),
				LeftGeoBeltLimit:                unifieddatalibrary.Float(20.23),
				Location:                        unifieddatalibrary.String("KENNEDY SPACE CENTER, FL"),
				LoopGain:                        unifieddatalibrary.Float(150.1),
				LunarExclAngle:                  unifieddatalibrary.Float(45.2),
				MagDec:                          unifieddatalibrary.Float(45.23),
				MagnitudeLimit:                  unifieddatalibrary.Float(23.5),
				MaxDeviationAngle:               unifieddatalibrary.Float(20.23),
				MaxIntegrationTime:              unifieddatalibrary.Float(1.1),
				MaxObservableRange:              unifieddatalibrary.Float(20.23),
				MaxRangeLimit:                   unifieddatalibrary.Float(4972.1),
				MaxWavelength:                   unifieddatalibrary.Float(7.89),
				MinIntegrationTime:              unifieddatalibrary.Float(0.3),
				MinRangeLimit:                   unifieddatalibrary.Float(165.1),
				MinSignalNoiseRatio:             unifieddatalibrary.Float(31.5),
				MinWavelength:                   unifieddatalibrary.Float(1.23),
				NegativeRangeRateLimit:          unifieddatalibrary.Float(-19.25),
				NoiseFigure:                     unifieddatalibrary.Float(10.23),
				NonCoherentIntegratedPulses:     unifieddatalibrary.Int(2),
				NumIntegratedPulses:             unifieddatalibrary.Int(10),
				NumIntegrationFrames:            unifieddatalibrary.Int(2),
				NumOpticalIntegrationModes:      unifieddatalibrary.Int(2),
				NumWaveforms:                    unifieddatalibrary.Int(2),
				OpticalIntegrationAngularRates:  []float64{15.1, 0.1},
				OpticalIntegrationFrames:        []float64{2.1, 3.1},
				OpticalIntegrationPixelBinnings: []float64{2.1, 1.1},
				OpticalIntegrationSnRs:          []float64{6.1, 6.1},
				OpticalIntegrationTimes:         []float64{0.3, 1},
				OpticalTransmission:             unifieddatalibrary.Float(0.5),
				PatternAbsorptionLoss:           unifieddatalibrary.Float(1.23),
				PatternScanLoss:                 unifieddatalibrary.Float(1.23),
				PeakPower:                       unifieddatalibrary.Float(5000000.1),
				PixelInstantaneousFov:           unifieddatalibrary.Float(10.23),
				PixelWellDepth:                  unifieddatalibrary.Int(12),
				PositiveRangeRateLimit:          unifieddatalibrary.Float(19.25),
				Prf:                             unifieddatalibrary.Float(20.23),
				ProbDetectSnr:                   unifieddatalibrary.Float(0.5),
				ProbFalseAlarm:                  unifieddatalibrary.Float(0.5),
				PulseRepPeriods:                 []float64{153.8, 153.9},
				QuantumEff:                      unifieddatalibrary.Float(0.5),
				RadarFrequency:                  unifieddatalibrary.Float(45300000000.1),
				RadarMessageFormat:              unifieddatalibrary.String("DATA_FORMAT"),
				RadarMur:                        unifieddatalibrary.Float(20.23),
				RadarPulseWidths:                []float64{20.23, 20.33},
				RadioFrequency:                  unifieddatalibrary.Float(20.23),
				RadomeLoss:                      unifieddatalibrary.Float(1.23),
				RangeGates:                      []int64{51, 52},
				RangeSpacings:                   []float64{690.2, 690.3},
				ReadNoise:                       unifieddatalibrary.Int(12),
				ReceiveGain:                     unifieddatalibrary.Float(10.2),
				ReceiveHorizBeamWidth:           unifieddatalibrary.Float(75.3),
				ReceiveLoss:                     unifieddatalibrary.Float(1.23),
				ReceiveVertBeamWidth:            unifieddatalibrary.Float(75.4),
				RefTemp:                         unifieddatalibrary.Float(3.5),
				ReqRecords:                      []int64{0, 1},
				RightClockAngle:                 unifieddatalibrary.Float(20.23),
				RightGeoBeltLimit:               unifieddatalibrary.Float(20.23),
				RunMeanCodes:                    []int64{0, 5},
				SignalProcessingLoss:            unifieddatalibrary.Float(1.23),
				SiteCode:                        unifieddatalibrary.String("07"),
				SolarExclAngle:                  unifieddatalibrary.Float(50.5),
				SpecAvgSpectraNums:              []int64{3, 4},
				SystemNoiseTemperature:          unifieddatalibrary.Float(3.5),
				TaskableRange:                   unifieddatalibrary.Float(20.23),
				TempMedFiltCodes:                []int64{3, 4},
				TestNumber:                      unifieddatalibrary.String("02022"),
				TotRecNums:                      []int64{5, 2},
				TowerHeight:                     unifieddatalibrary.Float(20.23),
				TrackAngle:                      unifieddatalibrary.Float(0.043),
				TrackSnr:                        unifieddatalibrary.Float(15.1),
				TransmitGain:                    unifieddatalibrary.Float(32.1),
				TransmitHorizBeamWidth:          unifieddatalibrary.Float(135.1),
				TransmitLoss:                    unifieddatalibrary.Float(7.1),
				TransmitPower:                   unifieddatalibrary.Float(190000.1),
				TransmitVertBeamWidth:           unifieddatalibrary.Float(45.2),
				TrueNorthCorrector:              unifieddatalibrary.Int(10),
				TrueTilt:                        unifieddatalibrary.Float(20.23),
				TwilightAngle:                   unifieddatalibrary.Float(7.5),
				VertBeamFlag:                    unifieddatalibrary.Bool(false),
				VertGateSpacings:                []float64{149.1, 149.2},
				VertGateWidths:                  []float64{149.1, 149.2},
				VFov:                            unifieddatalibrary.Float(20.23),
				VResPixels:                      unifieddatalibrary.Int(1000),
				WaveformBandwidths:              []float64{0.25, 0.25},
				WaveformLoopGains:               []float64{150.1, 155.1},
				WaveformMaxRanges:               []float64{2000.1, 2000.1},
				WaveformMinRanges:               []float64{150.1, 150.1},
				WaveformPulseWidths:             []float64{100.1, 200.1},
				Z1MaxRange:                      unifieddatalibrary.Float(50.23),
				Z1MinRange:                      unifieddatalibrary.Float(20.23),
				Z2MaxRange:                      unifieddatalibrary.Float(50.23),
				Z2MinRange:                      unifieddatalibrary.Float(20.23),
			}},
			SensorlimitsCollection: []unifieddatalibrary.SensorUpdateParamsSensorlimitsCollection{{
				ClassificationMarking:    "U",
				DataMode:                 "TEST",
				Source:                   "Bluestaq",
				IDSensor:                 unifieddatalibrary.String("SENSORLIMITS-ID"),
				IDSensorLimits:           unifieddatalibrary.String("SENSORLIMITS-ID"),
				LowerLeftAzimuthLimit:    unifieddatalibrary.Float(1.23),
				LowerLeftElevationLimit:  unifieddatalibrary.Float(1.23),
				LowerRightAzimuthLimit:   unifieddatalibrary.Float(1.23),
				LowerRightElevationLimit: unifieddatalibrary.Float(1.23),
				UpperLeftAzimuthLimit:    unifieddatalibrary.Float(1.23),
				UpperLeftElevationLimit:  unifieddatalibrary.Float(1.23),
				UpperRightAzimuthLimit:   unifieddatalibrary.Float(1.23),
				UpperRightElevationLimit: unifieddatalibrary.Float(1.23),
			}},
			SensorNumber: unifieddatalibrary.Int(1234),
			SensorObservationType: unifieddatalibrary.SensorUpdateParamsSensorObservationType{
				ID:   unifieddatalibrary.String("3"),
				Type: unifieddatalibrary.String("5"),
			},
			SensorStats: []unifieddatalibrary.SensorUpdateParamsSensorStat{{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				IDSensor:              "idSensor",
				Source:                "Bluestaq",
				ID:                    unifieddatalibrary.String("SENSORSTATS-ID"),
				LastObTime:            unifieddatalibrary.Time(time.Now()),
			}},
			SensorType: unifieddatalibrary.SensorUpdateParamsSensorType{
				ID:   unifieddatalibrary.Int(12344411),
				Type: unifieddatalibrary.String("Space Borne"),
			},
			ShortName: unifieddatalibrary.String("SNR-1"),
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

func TestSensorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.List(context.TODO(), unifieddatalibrary.SensorListParams{
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

func TestSensorDelete(t *testing.T) {
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
	err := client.Sensor.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Count(context.TODO(), unifieddatalibrary.SensorCountParams{
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

func TestSensorGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorGetParams{
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

func TestSensorQueryhelp(t *testing.T) {
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
	_, err := client.Sensor.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensor.Tuple(context.TODO(), unifieddatalibrary.SensorTupleParams{
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
