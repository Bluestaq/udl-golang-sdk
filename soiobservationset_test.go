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

func TestSoiobservationsetNewWithOptionalParams(t *testing.T) {
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
	err := client.Soiobservationset.New(context.TODO(), unifieddatalibrary.SoiobservationsetNewParams{
		ClassificationMarking:            "U",
		DataMode:                         unifieddatalibrary.SoiobservationsetNewParamsDataModeTest,
		NumObs:                           1,
		Source:                           "Bluestaq",
		StartTime:                        time.Now(),
		Type:                             unifieddatalibrary.SoiobservationsetNewParamsTypeOptical,
		ID:                               unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		BinningHoriz:                     unifieddatalibrary.Int(2),
		BinningVert:                      unifieddatalibrary.Int(2),
		BrightnessVarianceChangeDetected: unifieddatalibrary.Bool(true),
		Calibrations: []unifieddatalibrary.SoiobservationsetNewParamsCalibration{{
			CalBgIntensity:            unifieddatalibrary.Float(1.1),
			CalExtinctionCoeff:        unifieddatalibrary.Float(0.2),
			CalExtinctionCoeffMaxUnc:  unifieddatalibrary.Float(0.19708838),
			CalExtinctionCoeffUnc:     unifieddatalibrary.Float(0.06474939),
			CalNumCorrelatedStars:     unifieddatalibrary.Int(1),
			CalNumDetectedStars:       unifieddatalibrary.Int(1),
			CalSkyBg:                  unifieddatalibrary.Float(30086.25),
			CalSpectralFilterSolarMag: unifieddatalibrary.Float(19.23664587),
			CalTime:                   unifieddatalibrary.Time(time.Now()),
			CalType:                   unifieddatalibrary.String("PRE"),
			CalZeroPoint:              unifieddatalibrary.Float(25.15682157),
		}},
		CalibrationType:       unifieddatalibrary.String("ALL SKY"),
		ChangeConf:            unifieddatalibrary.String("MEDIUM"),
		ChangeDetected:        unifieddatalibrary.Bool(true),
		CollectionDensityConf: unifieddatalibrary.String("MEDIUM"),
		CollectionID:          unifieddatalibrary.String("b5133288-ab63-4b15-81f6-c7eec0cdb0c0"),
		CollectionMode:        unifieddatalibrary.String("RATE TRACK"),
		CorrQuality:           unifieddatalibrary.Float(0.327),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		Gain:                  unifieddatalibrary.Float(234.2),
		IDElset:               unifieddatalibrary.String("REF-ELSET-ID"),
		IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
		LosDeclinationEnd:     unifieddatalibrary.Float(1.1),
		LosDeclinationStart:   unifieddatalibrary.Float(1.1),
		MsgCreateDate:         unifieddatalibrary.Time(time.Now()),
		NumSpectralFilters:    unifieddatalibrary.Int(10),
		OpticalSoiObservationList: []unifieddatalibrary.SoiobservationsetNewParamsOpticalSoiObservationList{{
			ObStartTime:              time.Now(),
			CurrentSpectralFilterNum: unifieddatalibrary.Int(0),
			Declinations:             []float64{-0.45, -0.45, -0.45},
			ExpDuration:              unifieddatalibrary.Float(0.455),
			ExtinctionCoeffs:         []float64{0.32, 0.32, 0.32},
			ExtinctionCoeffsUnc:      []float64{0.06, 0.06, 0.06},
			Intensities:              []float64{1.1, 1.1, 1.1},
			IntensityTimes:           []time.Time{time.Now(), time.Now(), time.Now()},
			LocalSkyBgs:              []float64{100625.375, 100625.375, 100625.375},
			LocalSkyBgsUnc:           []float64{0.065, 0.065, 0.065},
			NumCorrelatedStars:       []int64{3, 3, 3},
			NumDetectedStars:         []int64{6, 6, 6},
			PercentSats:              []float64{0.1, 0.2, 1},
			RaRates:                  []float64{0, 0, 0},
			Ras:                      []float64{107.4, 107.4, 107.4},
			SkyBgs:                   []float64{100625.375, 100625.375, 100625.375},
			ZeroPoints:               []float64{24.711, 24.711, 24.711},
		}},
		Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:              unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:              unifieddatalibrary.String("ORIGSENSOR-ID"),
		PercentSatThreshold:       unifieddatalibrary.Float(0.1),
		PeriodicityChangeDetected: unifieddatalibrary.Bool(true),
		PeriodicityDetectionConf:  unifieddatalibrary.String("MEDIUM"),
		PeriodicitySamplingConf:   unifieddatalibrary.String("MEDIUM"),
		PixelArrayHeight:          unifieddatalibrary.Int(32),
		PixelArrayWidth:           unifieddatalibrary.Int(32),
		PixelMax:                  unifieddatalibrary.Int(16383),
		PixelMin:                  unifieddatalibrary.Int(0),
		PointingAngleAzEnd:        unifieddatalibrary.Float(1.1),
		PointingAngleAzStart:      unifieddatalibrary.Float(1.1),
		PointingAngleElEnd:        unifieddatalibrary.Float(1.1),
		PointingAngleElStart:      unifieddatalibrary.Float(1.1),
		PolarAngleEnd:             unifieddatalibrary.Float(1.1),
		PolarAngleStart:           unifieddatalibrary.Float(1.1),
		RadarSoiObservationList: []unifieddatalibrary.SoiobservationsetNewParamsRadarSoiObservationList{{
			ObStartTime:       time.Now(),
			AspectAngles:      []float64{4.278, 4.278, 4.278},
			AzimuthBiases:     []float64{45.23, 45.23, 45.23},
			AzimuthRates:      []float64{-1.481, -1.481, -1.481},
			Azimuths:          []float64{278.27, 278.27, 278.27},
			Beta:              unifieddatalibrary.Float(-89.97),
			CenterFrequency:   unifieddatalibrary.Float(160047.0625),
			CrossRangeRes:     []float64{11.301, 11.301, 11.301},
			DeltaTimes:        []float64{0.005, 0.005, 0.005},
			Doppler2XRs:       []float64{5644.27, 5644.27, 5644.27},
			ElevationBiases:   []float64{1.23, 1.23, 1.23},
			ElevationRates:    []float64{-0.074, -0.074, -0.074},
			Elevations:        []float64{70.85, 70.85, 70.85},
			IDAttitudeSet:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
			IDStateVector:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
			IntegrationAngles: []float64{8.594, 8.594, 8.594},
			Kappa:             unifieddatalibrary.Float(103.04),
			PeakAmplitudes:    []float64{33.1, 33.1, 33.1},
			Polarizations:     []string{"H", "L", "V"},
			ProjAngVels:       []float64{0.166, 0.166, 0.166},
			PulseBandwidth:    unifieddatalibrary.Float(24094.12),
			RangeAccels:       []float64{0.12, 0.01, 0.2},
			RangeBiases:       []float64{1.23, 1.23, 1.23},
			RangeRates:        []float64{0.317, 0.317, 0.317},
			Ranges:            []float64{877.938, 877.938, 877.938},
			RcsErrorEsts:      []float64{0.01, 0.003, 0.001},
			RcsValues:         []float64{12.34, 26.11, 43.21},
			Rspaces:           []float64{0.006, 0.006, 0.006},
			SpectralWidths:    []float64{23.45, 20.57, 12.21},
			Tovs:              []time.Time{time.Now(), time.Now(), time.Now()},
			Xaccel:            []float64{-0.075, -0.74, -0.4},
			Xpos:              []float64{-1118.577381, -1118.577381, -1118.577381},
			Xspaces:           []float64{0.006, 0.006, 0.006},
			Xvel:              []float64{-4.25242784, -4.25242784, -4.25242784},
			Yaccel:            []float64{-0.007, 0.003, 0.1},
			Ypos:              []float64{3026.231084, 3026.231084, 3026.231084},
			Yvel:              []float64{5.291107434, 5.291107434, 5.291107434},
			Zaccel:            []float64{0.1, 0.2, 0.3},
			Zpos:              []float64{6167.831808, 6167.831808, 6167.831808},
			Zvel:              []float64{-3.356493869, -3.356493869, -3.356493869},
		}},
		ReferenceFrame:                          unifieddatalibrary.SoiobservationsetNewParamsReferenceFrameJ2000,
		SatelliteName:                           unifieddatalibrary.String("TITAN 3C TRANSTAGE R/B"),
		SatNo:                                   unifieddatalibrary.Int(101),
		Senalt:                                  unifieddatalibrary.Float(1.1),
		Senlat:                                  unifieddatalibrary.Float(45.1),
		Senlon:                                  unifieddatalibrary.Float(179.1),
		SenReferenceFrame:                       unifieddatalibrary.SoiobservationsetNewParamsSenReferenceFrameJ2000,
		SensorAsID:                              unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		Senvelx:                                 unifieddatalibrary.Float(1.1),
		Senvely:                                 unifieddatalibrary.Float(1.1),
		Senvelz:                                 unifieddatalibrary.Float(1.1),
		Senx:                                    unifieddatalibrary.Float(1.1),
		Seny:                                    unifieddatalibrary.Float(1.1),
		Senz:                                    unifieddatalibrary.Float(1.1),
		SoftwareVersion:                         unifieddatalibrary.String("GSV99/17-1"),
		SolarMag:                                unifieddatalibrary.Float(-26.91),
		SolarPhaseAngleBrightnessChangeDetected: unifieddatalibrary.Bool(true),
		SpectralFilters:                         []string{"Keyword1", "Keyword2"},
		StarCatName:                             unifieddatalibrary.String("SSTRC5"),
		Tags:                                    []string{"TAG1", "TAG2"},
		TransactionID:                           unifieddatalibrary.String("TRANSACTION-ID"),
		Uct:                                     unifieddatalibrary.Bool(true),
		ValidCalibrations:                       unifieddatalibrary.String("BOTH"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSoiobservationsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Soiobservationset.List(context.TODO(), unifieddatalibrary.SoiobservationsetListParams{
		StartTime:   time.Now(),
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

func TestSoiobservationsetCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Soiobservationset.Count(context.TODO(), unifieddatalibrary.SoiobservationsetCountParams{
		StartTime:   time.Now(),
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

func TestSoiobservationsetNewBulk(t *testing.T) {
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
	err := client.Soiobservationset.NewBulk(context.TODO(), unifieddatalibrary.SoiobservationsetNewBulkParams{
		Body: []unifieddatalibrary.SoiobservationsetNewBulkParamsBody{{
			ClassificationMarking:            "U",
			DataMode:                         "TEST",
			NumObs:                           1,
			Source:                           "Bluestaq",
			StartTime:                        time.Now(),
			Type:                             "OPTICAL",
			ID:                               unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			BinningHoriz:                     unifieddatalibrary.Int(2),
			BinningVert:                      unifieddatalibrary.Int(2),
			BrightnessVarianceChangeDetected: unifieddatalibrary.Bool(true),
			Calibrations: []unifieddatalibrary.SoiobservationsetNewBulkParamsBodyCalibration{{
				CalBgIntensity:            unifieddatalibrary.Float(1.1),
				CalExtinctionCoeff:        unifieddatalibrary.Float(0.2),
				CalExtinctionCoeffMaxUnc:  unifieddatalibrary.Float(0.19708838),
				CalExtinctionCoeffUnc:     unifieddatalibrary.Float(0.06474939),
				CalNumCorrelatedStars:     unifieddatalibrary.Int(1),
				CalNumDetectedStars:       unifieddatalibrary.Int(1),
				CalSkyBg:                  unifieddatalibrary.Float(30086.25),
				CalSpectralFilterSolarMag: unifieddatalibrary.Float(19.23664587),
				CalTime:                   unifieddatalibrary.Time(time.Now()),
				CalType:                   unifieddatalibrary.String("PRE"),
				CalZeroPoint:              unifieddatalibrary.Float(25.15682157),
			}},
			CalibrationType:       unifieddatalibrary.String("ALL SKY"),
			ChangeConf:            unifieddatalibrary.String("MEDIUM"),
			ChangeDetected:        unifieddatalibrary.Bool(true),
			CollectionDensityConf: unifieddatalibrary.String("MEDIUM"),
			CollectionID:          unifieddatalibrary.String("b5133288-ab63-4b15-81f6-c7eec0cdb0c0"),
			CollectionMode:        unifieddatalibrary.String("RATE TRACK"),
			CorrQuality:           unifieddatalibrary.Float(0.327),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			Gain:                  unifieddatalibrary.Float(234.2),
			IDElset:               unifieddatalibrary.String("REF-ELSET-ID"),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			LosDeclinationEnd:     unifieddatalibrary.Float(1.1),
			LosDeclinationStart:   unifieddatalibrary.Float(1.1),
			MsgCreateDate:         unifieddatalibrary.Time(time.Now()),
			NumSpectralFilters:    unifieddatalibrary.Int(10),
			OpticalSoiObservationList: []unifieddatalibrary.SoiobservationsetNewBulkParamsBodyOpticalSoiObservationList{{
				ObStartTime:              time.Now(),
				CurrentSpectralFilterNum: unifieddatalibrary.Int(0),
				Declinations:             []float64{-0.45, -0.45, -0.45},
				ExpDuration:              unifieddatalibrary.Float(0.455),
				ExtinctionCoeffs:         []float64{0.32, 0.32, 0.32},
				ExtinctionCoeffsUnc:      []float64{0.06, 0.06, 0.06},
				Intensities:              []float64{1.1, 1.1, 1.1},
				IntensityTimes:           []time.Time{time.Now(), time.Now(), time.Now()},
				LocalSkyBgs:              []float64{100625.375, 100625.375, 100625.375},
				LocalSkyBgsUnc:           []float64{0.065, 0.065, 0.065},
				NumCorrelatedStars:       []int64{3, 3, 3},
				NumDetectedStars:         []int64{6, 6, 6},
				PercentSats:              []float64{0.1, 0.2, 1},
				RaRates:                  []float64{0, 0, 0},
				Ras:                      []float64{107.4, 107.4, 107.4},
				SkyBgs:                   []float64{100625.375, 100625.375, 100625.375},
				ZeroPoints:               []float64{24.711, 24.711, 24.711},
			}},
			Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:              unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:              unifieddatalibrary.String("ORIGSENSOR-ID"),
			PercentSatThreshold:       unifieddatalibrary.Float(0.1),
			PeriodicityChangeDetected: unifieddatalibrary.Bool(true),
			PeriodicityDetectionConf:  unifieddatalibrary.String("MEDIUM"),
			PeriodicitySamplingConf:   unifieddatalibrary.String("MEDIUM"),
			PixelArrayHeight:          unifieddatalibrary.Int(32),
			PixelArrayWidth:           unifieddatalibrary.Int(32),
			PixelMax:                  unifieddatalibrary.Int(16383),
			PixelMin:                  unifieddatalibrary.Int(0),
			PointingAngleAzEnd:        unifieddatalibrary.Float(1.1),
			PointingAngleAzStart:      unifieddatalibrary.Float(1.1),
			PointingAngleElEnd:        unifieddatalibrary.Float(1.1),
			PointingAngleElStart:      unifieddatalibrary.Float(1.1),
			PolarAngleEnd:             unifieddatalibrary.Float(1.1),
			PolarAngleStart:           unifieddatalibrary.Float(1.1),
			RadarSoiObservationList: []unifieddatalibrary.SoiobservationsetNewBulkParamsBodyRadarSoiObservationList{{
				ObStartTime:       time.Now(),
				AspectAngles:      []float64{4.278, 4.278, 4.278},
				AzimuthBiases:     []float64{45.23, 45.23, 45.23},
				AzimuthRates:      []float64{-1.481, -1.481, -1.481},
				Azimuths:          []float64{278.27, 278.27, 278.27},
				Beta:              unifieddatalibrary.Float(-89.97),
				CenterFrequency:   unifieddatalibrary.Float(160047.0625),
				CrossRangeRes:     []float64{11.301, 11.301, 11.301},
				DeltaTimes:        []float64{0.005, 0.005, 0.005},
				Doppler2XRs:       []float64{5644.27, 5644.27, 5644.27},
				ElevationBiases:   []float64{1.23, 1.23, 1.23},
				ElevationRates:    []float64{-0.074, -0.074, -0.074},
				Elevations:        []float64{70.85, 70.85, 70.85},
				IDAttitudeSet:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
				IDStateVector:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
				IntegrationAngles: []float64{8.594, 8.594, 8.594},
				Kappa:             unifieddatalibrary.Float(103.04),
				PeakAmplitudes:    []float64{33.1, 33.1, 33.1},
				Polarizations:     []string{"H", "L", "V"},
				ProjAngVels:       []float64{0.166, 0.166, 0.166},
				PulseBandwidth:    unifieddatalibrary.Float(24094.12),
				RangeAccels:       []float64{0.12, 0.01, 0.2},
				RangeBiases:       []float64{1.23, 1.23, 1.23},
				RangeRates:        []float64{0.317, 0.317, 0.317},
				Ranges:            []float64{877.938, 877.938, 877.938},
				RcsErrorEsts:      []float64{0.01, 0.003, 0.001},
				RcsValues:         []float64{12.34, 26.11, 43.21},
				Rspaces:           []float64{0.006, 0.006, 0.006},
				SpectralWidths:    []float64{23.45, 20.57, 12.21},
				Tovs:              []time.Time{time.Now(), time.Now(), time.Now()},
				Xaccel:            []float64{-0.075, -0.74, -0.4},
				Xpos:              []float64{-1118.577381, -1118.577381, -1118.577381},
				Xspaces:           []float64{0.006, 0.006, 0.006},
				Xvel:              []float64{-4.25242784, -4.25242784, -4.25242784},
				Yaccel:            []float64{-0.007, 0.003, 0.1},
				Ypos:              []float64{3026.231084, 3026.231084, 3026.231084},
				Yvel:              []float64{5.291107434, 5.291107434, 5.291107434},
				Zaccel:            []float64{0.1, 0.2, 0.3},
				Zpos:              []float64{6167.831808, 6167.831808, 6167.831808},
				Zvel:              []float64{-3.356493869, -3.356493869, -3.356493869},
			}},
			ReferenceFrame:                          "J2000",
			SatelliteName:                           unifieddatalibrary.String("TITAN 3C TRANSTAGE R/B"),
			SatNo:                                   unifieddatalibrary.Int(101),
			Senalt:                                  unifieddatalibrary.Float(1.1),
			Senlat:                                  unifieddatalibrary.Float(45.1),
			Senlon:                                  unifieddatalibrary.Float(179.1),
			SenReferenceFrame:                       "J2000",
			SensorAsID:                              unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Senvelx:                                 unifieddatalibrary.Float(1.1),
			Senvely:                                 unifieddatalibrary.Float(1.1),
			Senvelz:                                 unifieddatalibrary.Float(1.1),
			Senx:                                    unifieddatalibrary.Float(1.1),
			Seny:                                    unifieddatalibrary.Float(1.1),
			Senz:                                    unifieddatalibrary.Float(1.1),
			SoftwareVersion:                         unifieddatalibrary.String("GSV99/17-1"),
			SolarMag:                                unifieddatalibrary.Float(-26.91),
			SolarPhaseAngleBrightnessChangeDetected: unifieddatalibrary.Bool(true),
			SpectralFilters:                         []string{"Keyword1", "Keyword2"},
			StarCatName:                             unifieddatalibrary.String("SSTRC5"),
			Tags:                                    []string{"TAG1", "TAG2"},
			TransactionID:                           unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                                     unifieddatalibrary.Bool(true),
			ValidCalibrations:                       unifieddatalibrary.String("BOTH"),
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

func TestSoiobservationsetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Soiobservationset.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SoiobservationsetGetParams{
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

func TestSoiobservationsetQueryhelp(t *testing.T) {
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
	err := client.Soiobservationset.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSoiobservationsetTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Soiobservationset.Tuple(context.TODO(), unifieddatalibrary.SoiobservationsetTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestSoiobservationsetUnvalidatedPublish(t *testing.T) {
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
	err := client.Soiobservationset.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SoiobservationsetUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SoiobservationsetUnvalidatedPublishParamsBody{{
			ClassificationMarking:            "U",
			DataMode:                         "TEST",
			NumObs:                           1,
			Source:                           "Bluestaq",
			StartTime:                        time.Now(),
			Type:                             "OPTICAL",
			ID:                               unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			BinningHoriz:                     unifieddatalibrary.Int(2),
			BinningVert:                      unifieddatalibrary.Int(2),
			BrightnessVarianceChangeDetected: unifieddatalibrary.Bool(true),
			Calibrations: []unifieddatalibrary.SoiobservationsetUnvalidatedPublishParamsBodyCalibration{{
				CalBgIntensity:            unifieddatalibrary.Float(1.1),
				CalExtinctionCoeff:        unifieddatalibrary.Float(0.2),
				CalExtinctionCoeffMaxUnc:  unifieddatalibrary.Float(0.19708838),
				CalExtinctionCoeffUnc:     unifieddatalibrary.Float(0.06474939),
				CalNumCorrelatedStars:     unifieddatalibrary.Int(1),
				CalNumDetectedStars:       unifieddatalibrary.Int(1),
				CalSkyBg:                  unifieddatalibrary.Float(30086.25),
				CalSpectralFilterSolarMag: unifieddatalibrary.Float(19.23664587),
				CalTime:                   unifieddatalibrary.Time(time.Now()),
				CalType:                   unifieddatalibrary.String("PRE"),
				CalZeroPoint:              unifieddatalibrary.Float(25.15682157),
			}},
			CalibrationType:       unifieddatalibrary.String("ALL SKY"),
			ChangeConf:            unifieddatalibrary.String("MEDIUM"),
			ChangeDetected:        unifieddatalibrary.Bool(true),
			CollectionDensityConf: unifieddatalibrary.String("MEDIUM"),
			CollectionID:          unifieddatalibrary.String("b5133288-ab63-4b15-81f6-c7eec0cdb0c0"),
			CollectionMode:        unifieddatalibrary.String("RATE TRACK"),
			CorrQuality:           unifieddatalibrary.Float(0.327),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			Gain:                  unifieddatalibrary.Float(234.2),
			IDElset:               unifieddatalibrary.String("REF-ELSET-ID"),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			LosDeclinationEnd:     unifieddatalibrary.Float(1.1),
			LosDeclinationStart:   unifieddatalibrary.Float(1.1),
			MsgCreateDate:         unifieddatalibrary.Time(time.Now()),
			NumSpectralFilters:    unifieddatalibrary.Int(10),
			OpticalSoiObservationList: []unifieddatalibrary.SoiobservationsetUnvalidatedPublishParamsBodyOpticalSoiObservationList{{
				ObStartTime:              time.Now(),
				CurrentSpectralFilterNum: unifieddatalibrary.Int(0),
				Declinations:             []float64{-0.45, -0.45, -0.45},
				ExpDuration:              unifieddatalibrary.Float(0.455),
				ExtinctionCoeffs:         []float64{0.32, 0.32, 0.32},
				ExtinctionCoeffsUnc:      []float64{0.06, 0.06, 0.06},
				Intensities:              []float64{1.1, 1.1, 1.1},
				IntensityTimes:           []time.Time{time.Now(), time.Now(), time.Now()},
				LocalSkyBgs:              []float64{100625.375, 100625.375, 100625.375},
				LocalSkyBgsUnc:           []float64{0.065, 0.065, 0.065},
				NumCorrelatedStars:       []int64{3, 3, 3},
				NumDetectedStars:         []int64{6, 6, 6},
				PercentSats:              []float64{0.1, 0.2, 1},
				RaRates:                  []float64{0, 0, 0},
				Ras:                      []float64{107.4, 107.4, 107.4},
				SkyBgs:                   []float64{100625.375, 100625.375, 100625.375},
				ZeroPoints:               []float64{24.711, 24.711, 24.711},
			}},
			Origin:                    unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:              unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:              unifieddatalibrary.String("ORIGSENSOR-ID"),
			PercentSatThreshold:       unifieddatalibrary.Float(0.1),
			PeriodicityChangeDetected: unifieddatalibrary.Bool(true),
			PeriodicityDetectionConf:  unifieddatalibrary.String("MEDIUM"),
			PeriodicitySamplingConf:   unifieddatalibrary.String("MEDIUM"),
			PixelArrayHeight:          unifieddatalibrary.Int(32),
			PixelArrayWidth:           unifieddatalibrary.Int(32),
			PixelMax:                  unifieddatalibrary.Int(16383),
			PixelMin:                  unifieddatalibrary.Int(0),
			PointingAngleAzEnd:        unifieddatalibrary.Float(1.1),
			PointingAngleAzStart:      unifieddatalibrary.Float(1.1),
			PointingAngleElEnd:        unifieddatalibrary.Float(1.1),
			PointingAngleElStart:      unifieddatalibrary.Float(1.1),
			PolarAngleEnd:             unifieddatalibrary.Float(1.1),
			PolarAngleStart:           unifieddatalibrary.Float(1.1),
			RadarSoiObservationList: []unifieddatalibrary.SoiobservationsetUnvalidatedPublishParamsBodyRadarSoiObservationList{{
				ObStartTime:       time.Now(),
				AspectAngles:      []float64{4.278, 4.278, 4.278},
				AzimuthBiases:     []float64{45.23, 45.23, 45.23},
				AzimuthRates:      []float64{-1.481, -1.481, -1.481},
				Azimuths:          []float64{278.27, 278.27, 278.27},
				Beta:              unifieddatalibrary.Float(-89.97),
				CenterFrequency:   unifieddatalibrary.Float(160047.0625),
				CrossRangeRes:     []float64{11.301, 11.301, 11.301},
				DeltaTimes:        []float64{0.005, 0.005, 0.005},
				Doppler2XRs:       []float64{5644.27, 5644.27, 5644.27},
				ElevationBiases:   []float64{1.23, 1.23, 1.23},
				ElevationRates:    []float64{-0.074, -0.074, -0.074},
				Elevations:        []float64{70.85, 70.85, 70.85},
				IDAttitudeSet:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
				IDStateVector:     unifieddatalibrary.String("99a0de63-b38f-4d81-b057"),
				IntegrationAngles: []float64{8.594, 8.594, 8.594},
				Kappa:             unifieddatalibrary.Float(103.04),
				PeakAmplitudes:    []float64{33.1, 33.1, 33.1},
				Polarizations:     []string{"H", "L", "V"},
				ProjAngVels:       []float64{0.166, 0.166, 0.166},
				PulseBandwidth:    unifieddatalibrary.Float(24094.12),
				RangeAccels:       []float64{0.12, 0.01, 0.2},
				RangeBiases:       []float64{1.23, 1.23, 1.23},
				RangeRates:        []float64{0.317, 0.317, 0.317},
				Ranges:            []float64{877.938, 877.938, 877.938},
				RcsErrorEsts:      []float64{0.01, 0.003, 0.001},
				RcsValues:         []float64{12.34, 26.11, 43.21},
				Rspaces:           []float64{0.006, 0.006, 0.006},
				SpectralWidths:    []float64{23.45, 20.57, 12.21},
				Tovs:              []time.Time{time.Now(), time.Now(), time.Now()},
				Xaccel:            []float64{-0.075, -0.74, -0.4},
				Xpos:              []float64{-1118.577381, -1118.577381, -1118.577381},
				Xspaces:           []float64{0.006, 0.006, 0.006},
				Xvel:              []float64{-4.25242784, -4.25242784, -4.25242784},
				Yaccel:            []float64{-0.007, 0.003, 0.1},
				Ypos:              []float64{3026.231084, 3026.231084, 3026.231084},
				Yvel:              []float64{5.291107434, 5.291107434, 5.291107434},
				Zaccel:            []float64{0.1, 0.2, 0.3},
				Zpos:              []float64{6167.831808, 6167.831808, 6167.831808},
				Zvel:              []float64{-3.356493869, -3.356493869, -3.356493869},
			}},
			ReferenceFrame:                          "J2000",
			SatelliteName:                           unifieddatalibrary.String("TITAN 3C TRANSTAGE R/B"),
			SatNo:                                   unifieddatalibrary.Int(101),
			Senalt:                                  unifieddatalibrary.Float(1.1),
			Senlat:                                  unifieddatalibrary.Float(45.1),
			Senlon:                                  unifieddatalibrary.Float(179.1),
			SenReferenceFrame:                       "J2000",
			SensorAsID:                              unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Senvelx:                                 unifieddatalibrary.Float(1.1),
			Senvely:                                 unifieddatalibrary.Float(1.1),
			Senvelz:                                 unifieddatalibrary.Float(1.1),
			Senx:                                    unifieddatalibrary.Float(1.1),
			Seny:                                    unifieddatalibrary.Float(1.1),
			Senz:                                    unifieddatalibrary.Float(1.1),
			SoftwareVersion:                         unifieddatalibrary.String("GSV99/17-1"),
			SolarMag:                                unifieddatalibrary.Float(-26.91),
			SolarPhaseAngleBrightnessChangeDetected: unifieddatalibrary.Bool(true),
			SpectralFilters:                         []string{"Keyword1", "Keyword2"},
			StarCatName:                             unifieddatalibrary.String("SSTRC5"),
			Tags:                                    []string{"TAG1", "TAG2"},
			TransactionID:                           unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                                     unifieddatalibrary.Bool(true),
			ValidCalibrations:                       unifieddatalibrary.String("BOTH"),
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
