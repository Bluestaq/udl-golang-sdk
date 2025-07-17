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

func TestObservationEoObservationNewWithOptionalParams(t *testing.T) {
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
	err := client.Observations.EoObservations.New(context.TODO(), unifieddatalibrary.ObservationEoObservationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ObservationEoObservationNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("EOOBSERVATION-ID"),
		Azimuth:               unifieddatalibrary.Float(1.1),
		AzimuthBias:           unifieddatalibrary.Float(1.1),
		AzimuthMeasured:       unifieddatalibrary.Bool(true),
		AzimuthRate:           unifieddatalibrary.Float(1.1),
		AzimuthUnc:            unifieddatalibrary.Float(1.1),
		BgIntensity:           unifieddatalibrary.Float(1.1),
		CollectMethod:         unifieddatalibrary.String("AUTOTRACK"),
		CorrQuality:           unifieddatalibrary.Float(1.1),
		Declination:           unifieddatalibrary.Float(1.1),
		DeclinationBias:       unifieddatalibrary.Float(1.1),
		DeclinationMeasured:   unifieddatalibrary.Bool(true),
		DeclinationRate:       unifieddatalibrary.Float(1.1),
		DeclinationUnc:        unifieddatalibrary.Float(1.1),
		Descriptor:            unifieddatalibrary.String("PROVIDED_DATA1"),
		Elevation:             unifieddatalibrary.Float(1.1),
		ElevationBias:         unifieddatalibrary.Float(1.1),
		ElevationMeasured:     unifieddatalibrary.Bool(true),
		ElevationRate:         unifieddatalibrary.Float(1.1),
		ElevationUnc:          unifieddatalibrary.Float(1.1),
		EoobservationDetails: unifieddatalibrary.ObservationEoObservationNewParamsEoobservationDetails{
			AcalCrPixX:                           unifieddatalibrary.Float(123.2),
			AcalCrPixY:                           unifieddatalibrary.Float(123.2),
			AcalCrValX:                           unifieddatalibrary.Float(123.2),
			AcalCrValY:                           unifieddatalibrary.Float(123.2),
			AcalNumStars:                         unifieddatalibrary.Int(123),
			BackgroundSignal:                     unifieddatalibrary.Float(4134.1),
			BackgroundSignalUnc:                  unifieddatalibrary.Float(123.2),
			BinningHoriz:                         unifieddatalibrary.Int(12),
			BinningVert:                          unifieddatalibrary.Int(14),
			CcdObjPosX:                           unifieddatalibrary.Float(123.3),
			CcdObjPosY:                           unifieddatalibrary.Float(321.4),
			CcdObjWidth:                          unifieddatalibrary.Float(133.2),
			CcdTemp:                              unifieddatalibrary.Float(123.4),
			CentroidColumn:                       unifieddatalibrary.Float(0.5),
			CentroidRow:                          unifieddatalibrary.Float(0.1),
			ClassificationMarking:                unifieddatalibrary.String("U"),
			ColorCoeffs:                          []float64{1.1, 2.1, 3.1},
			ColumnVariance:                       unifieddatalibrary.Float(0.1),
			CurrentNeutralDensityFilterNum:       unifieddatalibrary.Int(3),
			CurrentSpectralFilterNum:             unifieddatalibrary.Int(23),
			DataMode:                             "TEST",
			DeclinationCov:                       unifieddatalibrary.Float(123.2),
			DistFromStreakCenter:                 []float64{-127.153, -126.153, -125.153},
			Does:                                 unifieddatalibrary.Float(123.2),
			ExtinctionCoeffs:                     []float64{1.1, 2.1, 3.1},
			ExtinctionCoeffsUnc:                  []float64{1.1, 2.1, 3.1},
			Gain:                                 unifieddatalibrary.Float(234.2),
			IDEoObservation:                      unifieddatalibrary.String("EOOBSERVATION-ID"),
			Ifov:                                 unifieddatalibrary.Float(0.2),
			MagInstrumental:                      unifieddatalibrary.Float(123.3),
			MagInstrumentalUnc:                   unifieddatalibrary.Float(123.3),
			NeutralDensityFilterNames:            []string{"numNeutralDensityFilters1", "numNeutralDensityFilters2"},
			NeutralDensityFilterTransmissions:    []float64{1.1, 2.1, 3.1},
			NeutralDensityFilterTransmissionsUnc: []float64{1.1, 2.1, 3.1},
			NumCatalogStars:                      unifieddatalibrary.Int(123),
			NumCorrelatedStars:                   unifieddatalibrary.Int(123),
			NumDetectedStars:                     unifieddatalibrary.Int(123),
			NumNeutralDensityFilters:             unifieddatalibrary.Int(12),
			NumSpectralFilters:                   unifieddatalibrary.Int(10),
			ObjSunRange:                          unifieddatalibrary.Float(123.2),
			ObTime:                               unifieddatalibrary.Time(time.Now()),
			OpticalCrossSection:                  unifieddatalibrary.Float(123.3),
			OpticalCrossSectionUnc:               unifieddatalibrary.Float(123.3),
			PcalNumStars:                         unifieddatalibrary.Int(23),
			PeakApertureCount:                    unifieddatalibrary.Float(123.2),
			PeakBackgroundCount:                  unifieddatalibrary.Int(321),
			PhaseAngBisect:                       unifieddatalibrary.Float(123.2),
			PixelArrayHeight:                     unifieddatalibrary.Int(23),
			PixelArrayWidth:                      unifieddatalibrary.Int(12),
			PixelMax:                             unifieddatalibrary.Int(256),
			PixelMin:                             unifieddatalibrary.Int(12),
			PredictedAzimuth:                     unifieddatalibrary.Float(10.1),
			PredictedDeclination:                 unifieddatalibrary.Float(10.1),
			PredictedDeclinationUnc:              unifieddatalibrary.Float(123.2),
			PredictedElevation:                   unifieddatalibrary.Float(10.1),
			PredictedRa:                          unifieddatalibrary.Float(10.1),
			PredictedRaUnc:                       unifieddatalibrary.Float(123.2),
			RaCov:                                unifieddatalibrary.Float(123.2),
			RaDeclinationCov:                     unifieddatalibrary.Float(123.2),
			RowColCov:                            unifieddatalibrary.Float(0.01),
			RowVariance:                          unifieddatalibrary.Float(0.1),
			SnrEst:                               unifieddatalibrary.Float(13.4),
			SolarDiskFrac:                        unifieddatalibrary.Float(123.2),
			Source:                               unifieddatalibrary.String("Bluestaq"),
			SpectralFilters:                      []string{"Keyword1", "Keyword2"},
			SpectralFilterSolarMag:               []float64{1.1, 2.1, 3.1},
			SpectralZmfl:                         []float64{1.1, 2.1, 3.1},
			SunAzimuth:                           unifieddatalibrary.Float(10.1),
			SunElevation:                         unifieddatalibrary.Float(10.1),
			SunStatePosX:                         unifieddatalibrary.Float(123.3),
			SunStatePosY:                         unifieddatalibrary.Float(123.3),
			SunStatePosZ:                         unifieddatalibrary.Float(123.3),
			SunStateVelX:                         unifieddatalibrary.Float(123.3),
			SunStateVelY:                         unifieddatalibrary.Float(123.3),
			SunStateVelZ:                         unifieddatalibrary.Float(123.3),
			SurfBrightness:                       []float64{21.01, 21.382, 21.725},
			SurfBrightnessUnc:                    []float64{0.165, 0.165, 0.165},
			TimesUnc:                             unifieddatalibrary.Float(13.1),
			Toes:                                 unifieddatalibrary.Float(123.2),
			ZeroPoints:                           []float64{1.1, 2.1, 3.1},
			ZeroPointsUnc:                        []float64{1.1, 2.1, 3.1},
		},
		ExpDuration:          unifieddatalibrary.Float(1.1),
		FovCount:             unifieddatalibrary.Int(1),
		FovCountUct:          unifieddatalibrary.Int(2),
		Geoalt:               unifieddatalibrary.Float(1.1),
		Geolat:               unifieddatalibrary.Float(1.1),
		Geolon:               unifieddatalibrary.Float(1.1),
		Georange:             unifieddatalibrary.Float(1.1),
		IDSensor:             unifieddatalibrary.String("SENSOR-ID"),
		IDSkyImagery:         unifieddatalibrary.String("SKYIMAGERY-ID"),
		Intensity:            unifieddatalibrary.Float(1.1),
		LosUnc:               unifieddatalibrary.Float(1.1),
		Losx:                 unifieddatalibrary.Float(1.1),
		Losxvel:              unifieddatalibrary.Float(1.1),
		Losy:                 unifieddatalibrary.Float(1.1),
		Losyvel:              unifieddatalibrary.Float(1.1),
		Losz:                 unifieddatalibrary.Float(1.1),
		Loszvel:              unifieddatalibrary.Float(1.1),
		Mag:                  unifieddatalibrary.Float(1.1),
		MagNormRange:         unifieddatalibrary.Float(1.1),
		MagUnc:               unifieddatalibrary.Float(1.1),
		NetObjSig:            unifieddatalibrary.Float(1.1),
		NetObjSigUnc:         unifieddatalibrary.Float(1.1),
		ObPosition:           unifieddatalibrary.String("FIRST"),
		Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:         unifieddatalibrary.String("ORIGOBJECT-ID"),
		OrigSensorID:         unifieddatalibrary.String("ORIGSENSOR-ID"),
		Penumbra:             unifieddatalibrary.Bool(false),
		PrimaryExtinction:    unifieddatalibrary.Float(1.1),
		PrimaryExtinctionUnc: unifieddatalibrary.Float(1.1),
		Ra:                   unifieddatalibrary.Float(1.1),
		RaBias:               unifieddatalibrary.Float(1.1),
		RaMeasured:           unifieddatalibrary.Bool(true),
		Range:                unifieddatalibrary.Float(1.1),
		RangeBias:            unifieddatalibrary.Float(1.1),
		RangeMeasured:        unifieddatalibrary.Bool(true),
		RangeRate:            unifieddatalibrary.Float(1.1),
		RangeRateMeasured:    unifieddatalibrary.Bool(true),
		RangeRateUnc:         unifieddatalibrary.Float(1.1),
		RangeUnc:             unifieddatalibrary.Float(1.1),
		RaRate:               unifieddatalibrary.Float(1.1),
		RaUnc:                unifieddatalibrary.Float(1.1),
		RawFileUri:           unifieddatalibrary.String("Example URI"),
		ReferenceFrame:       unifieddatalibrary.ObservationEoObservationNewParamsReferenceFrameJ2000,
		SatNo:                unifieddatalibrary.Int(5),
		Senalt:               unifieddatalibrary.Float(1.1),
		Senlat:               unifieddatalibrary.Float(45.1),
		Senlon:               unifieddatalibrary.Float(179.1),
		SenQuat:              []float64{0.4492, 0.02, 0.8765, 0.2213},
		SenReferenceFrame:    unifieddatalibrary.ObservationEoObservationNewParamsSenReferenceFrameJ2000,
		Senvelx:              unifieddatalibrary.Float(1.1),
		Senvely:              unifieddatalibrary.Float(1.1),
		Senvelz:              unifieddatalibrary.Float(1.1),
		Senx:                 unifieddatalibrary.Float(1.1),
		Seny:                 unifieddatalibrary.Float(1.1),
		Senz:                 unifieddatalibrary.Float(1.1),
		ShutterDelay:         unifieddatalibrary.Float(1.1),
		SkyBkgrnd:            unifieddatalibrary.Float(1.1),
		SolarDecAngle:        unifieddatalibrary.Float(1.1),
		SolarEqPhaseAngle:    unifieddatalibrary.Float(1.1),
		SolarPhaseAngle:      unifieddatalibrary.Float(1.1),
		Tags:                 []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TaskID:               unifieddatalibrary.String("TASK-ID"),
		TimingBias:           unifieddatalibrary.Float(1.1),
		TrackID:              unifieddatalibrary.String("TRACK-ID"),
		TransactionID:        unifieddatalibrary.String("TRANSACTION-ID"),
		Uct:                  unifieddatalibrary.Bool(false),
		Umbra:                unifieddatalibrary.Bool(false),
		Zeroptd:              unifieddatalibrary.Float(1.1),
		ZeroPtdUnc:           unifieddatalibrary.Float(1.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationEoObservationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.EoObservations.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObservationEoObservationGetParams{
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

func TestObservationEoObservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.EoObservations.List(context.TODO(), unifieddatalibrary.ObservationEoObservationListParams{
		ObTime:      time.Now(),
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

func TestObservationEoObservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.EoObservations.Count(context.TODO(), unifieddatalibrary.ObservationEoObservationCountParams{
		ObTime:      time.Now(),
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

func TestObservationEoObservationNewBulkWithOptionalParams(t *testing.T) {
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
	err := client.Observations.EoObservations.NewBulk(context.TODO(), unifieddatalibrary.ObservationEoObservationNewBulkParams{
		Body: []unifieddatalibrary.ObservationEoObservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("EOOBSERVATION-ID"),
			Azimuth:               unifieddatalibrary.Float(1.1),
			AzimuthBias:           unifieddatalibrary.Float(1.1),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.1),
			AzimuthUnc:            unifieddatalibrary.Float(1.1),
			BgIntensity:           unifieddatalibrary.Float(1.1),
			CollectMethod:         unifieddatalibrary.String("AUTOTRACK"),
			CorrQuality:           unifieddatalibrary.Float(1.1),
			Declination:           unifieddatalibrary.Float(1.1),
			DeclinationBias:       unifieddatalibrary.Float(1.1),
			DeclinationMeasured:   unifieddatalibrary.Bool(true),
			DeclinationRate:       unifieddatalibrary.Float(1.1),
			DeclinationUnc:        unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("PROVIDED_DATA1"),
			Elevation:             unifieddatalibrary.Float(1.1),
			ElevationBias:         unifieddatalibrary.Float(1.1),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(1.1),
			ElevationUnc:          unifieddatalibrary.Float(1.1),
			EoobservationDetails: unifieddatalibrary.ObservationEoObservationNewBulkParamsBodyEoobservationDetails{
				AcalCrPixX:                           unifieddatalibrary.Float(123.2),
				AcalCrPixY:                           unifieddatalibrary.Float(123.2),
				AcalCrValX:                           unifieddatalibrary.Float(123.2),
				AcalCrValY:                           unifieddatalibrary.Float(123.2),
				AcalNumStars:                         unifieddatalibrary.Int(123),
				BackgroundSignal:                     unifieddatalibrary.Float(4134.1),
				BackgroundSignalUnc:                  unifieddatalibrary.Float(123.2),
				BinningHoriz:                         unifieddatalibrary.Int(12),
				BinningVert:                          unifieddatalibrary.Int(14),
				CcdObjPosX:                           unifieddatalibrary.Float(123.3),
				CcdObjPosY:                           unifieddatalibrary.Float(321.4),
				CcdObjWidth:                          unifieddatalibrary.Float(133.2),
				CcdTemp:                              unifieddatalibrary.Float(123.4),
				CentroidColumn:                       unifieddatalibrary.Float(0.5),
				CentroidRow:                          unifieddatalibrary.Float(0.1),
				ClassificationMarking:                unifieddatalibrary.String("U"),
				ColorCoeffs:                          []float64{1.1, 2.1, 3.1},
				ColumnVariance:                       unifieddatalibrary.Float(0.1),
				CurrentNeutralDensityFilterNum:       unifieddatalibrary.Int(3),
				CurrentSpectralFilterNum:             unifieddatalibrary.Int(23),
				DataMode:                             "TEST",
				DeclinationCov:                       unifieddatalibrary.Float(123.2),
				DistFromStreakCenter:                 []float64{-127.153, -126.153, -125.153},
				Does:                                 unifieddatalibrary.Float(123.2),
				ExtinctionCoeffs:                     []float64{1.1, 2.1, 3.1},
				ExtinctionCoeffsUnc:                  []float64{1.1, 2.1, 3.1},
				Gain:                                 unifieddatalibrary.Float(234.2),
				IDEoObservation:                      unifieddatalibrary.String("EOOBSERVATION-ID"),
				Ifov:                                 unifieddatalibrary.Float(0.2),
				MagInstrumental:                      unifieddatalibrary.Float(123.3),
				MagInstrumentalUnc:                   unifieddatalibrary.Float(123.3),
				NeutralDensityFilterNames:            []string{"numNeutralDensityFilters1", "numNeutralDensityFilters2"},
				NeutralDensityFilterTransmissions:    []float64{1.1, 2.1, 3.1},
				NeutralDensityFilterTransmissionsUnc: []float64{1.1, 2.1, 3.1},
				NumCatalogStars:                      unifieddatalibrary.Int(123),
				NumCorrelatedStars:                   unifieddatalibrary.Int(123),
				NumDetectedStars:                     unifieddatalibrary.Int(123),
				NumNeutralDensityFilters:             unifieddatalibrary.Int(12),
				NumSpectralFilters:                   unifieddatalibrary.Int(10),
				ObjSunRange:                          unifieddatalibrary.Float(123.2),
				ObTime:                               unifieddatalibrary.Time(time.Now()),
				OpticalCrossSection:                  unifieddatalibrary.Float(123.3),
				OpticalCrossSectionUnc:               unifieddatalibrary.Float(123.3),
				PcalNumStars:                         unifieddatalibrary.Int(23),
				PeakApertureCount:                    unifieddatalibrary.Float(123.2),
				PeakBackgroundCount:                  unifieddatalibrary.Int(321),
				PhaseAngBisect:                       unifieddatalibrary.Float(123.2),
				PixelArrayHeight:                     unifieddatalibrary.Int(23),
				PixelArrayWidth:                      unifieddatalibrary.Int(12),
				PixelMax:                             unifieddatalibrary.Int(256),
				PixelMin:                             unifieddatalibrary.Int(12),
				PredictedAzimuth:                     unifieddatalibrary.Float(10.1),
				PredictedDeclination:                 unifieddatalibrary.Float(10.1),
				PredictedDeclinationUnc:              unifieddatalibrary.Float(123.2),
				PredictedElevation:                   unifieddatalibrary.Float(10.1),
				PredictedRa:                          unifieddatalibrary.Float(10.1),
				PredictedRaUnc:                       unifieddatalibrary.Float(123.2),
				RaCov:                                unifieddatalibrary.Float(123.2),
				RaDeclinationCov:                     unifieddatalibrary.Float(123.2),
				RowColCov:                            unifieddatalibrary.Float(0.01),
				RowVariance:                          unifieddatalibrary.Float(0.1),
				SnrEst:                               unifieddatalibrary.Float(13.4),
				SolarDiskFrac:                        unifieddatalibrary.Float(123.2),
				Source:                               unifieddatalibrary.String("Bluestaq"),
				SpectralFilters:                      []string{"Keyword1", "Keyword2"},
				SpectralFilterSolarMag:               []float64{1.1, 2.1, 3.1},
				SpectralZmfl:                         []float64{1.1, 2.1, 3.1},
				SunAzimuth:                           unifieddatalibrary.Float(10.1),
				SunElevation:                         unifieddatalibrary.Float(10.1),
				SunStatePosX:                         unifieddatalibrary.Float(123.3),
				SunStatePosY:                         unifieddatalibrary.Float(123.3),
				SunStatePosZ:                         unifieddatalibrary.Float(123.3),
				SunStateVelX:                         unifieddatalibrary.Float(123.3),
				SunStateVelY:                         unifieddatalibrary.Float(123.3),
				SunStateVelZ:                         unifieddatalibrary.Float(123.3),
				SurfBrightness:                       []float64{21.01, 21.382, 21.725},
				SurfBrightnessUnc:                    []float64{0.165, 0.165, 0.165},
				TimesUnc:                             unifieddatalibrary.Float(13.1),
				Toes:                                 unifieddatalibrary.Float(123.2),
				ZeroPoints:                           []float64{1.1, 2.1, 3.1},
				ZeroPointsUnc:                        []float64{1.1, 2.1, 3.1},
			},
			ExpDuration:          unifieddatalibrary.Float(1.1),
			FovCount:             unifieddatalibrary.Int(1),
			FovCountUct:          unifieddatalibrary.Int(2),
			Geoalt:               unifieddatalibrary.Float(1.1),
			Geolat:               unifieddatalibrary.Float(1.1),
			Geolon:               unifieddatalibrary.Float(1.1),
			Georange:             unifieddatalibrary.Float(1.1),
			IDSensor:             unifieddatalibrary.String("SENSOR-ID"),
			IDSkyImagery:         unifieddatalibrary.String("SKYIMAGERY-ID"),
			Intensity:            unifieddatalibrary.Float(1.1),
			LosUnc:               unifieddatalibrary.Float(1.1),
			Losx:                 unifieddatalibrary.Float(1.1),
			Losxvel:              unifieddatalibrary.Float(1.1),
			Losy:                 unifieddatalibrary.Float(1.1),
			Losyvel:              unifieddatalibrary.Float(1.1),
			Losz:                 unifieddatalibrary.Float(1.1),
			Loszvel:              unifieddatalibrary.Float(1.1),
			Mag:                  unifieddatalibrary.Float(1.1),
			MagNormRange:         unifieddatalibrary.Float(1.1),
			MagUnc:               unifieddatalibrary.Float(1.1),
			NetObjSig:            unifieddatalibrary.Float(1.1),
			NetObjSigUnc:         unifieddatalibrary.Float(1.1),
			ObPosition:           unifieddatalibrary.String("FIRST"),
			Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:         unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:         unifieddatalibrary.String("ORIGSENSOR-ID"),
			Penumbra:             unifieddatalibrary.Bool(false),
			PrimaryExtinction:    unifieddatalibrary.Float(1.1),
			PrimaryExtinctionUnc: unifieddatalibrary.Float(1.1),
			Ra:                   unifieddatalibrary.Float(1.1),
			RaBias:               unifieddatalibrary.Float(1.1),
			RaMeasured:           unifieddatalibrary.Bool(true),
			Range:                unifieddatalibrary.Float(1.1),
			RangeBias:            unifieddatalibrary.Float(1.1),
			RangeMeasured:        unifieddatalibrary.Bool(true),
			RangeRate:            unifieddatalibrary.Float(1.1),
			RangeRateMeasured:    unifieddatalibrary.Bool(true),
			RangeRateUnc:         unifieddatalibrary.Float(1.1),
			RangeUnc:             unifieddatalibrary.Float(1.1),
			RaRate:               unifieddatalibrary.Float(1.1),
			RaUnc:                unifieddatalibrary.Float(1.1),
			RawFileUri:           unifieddatalibrary.String("Example URI"),
			ReferenceFrame:       "J2000",
			SatNo:                unifieddatalibrary.Int(5),
			Senalt:               unifieddatalibrary.Float(1.1),
			Senlat:               unifieddatalibrary.Float(45.1),
			Senlon:               unifieddatalibrary.Float(179.1),
			SenQuat:              []float64{0.4492, 0.02, 0.8765, 0.2213},
			SenReferenceFrame:    "J2000",
			Senvelx:              unifieddatalibrary.Float(1.1),
			Senvely:              unifieddatalibrary.Float(1.1),
			Senvelz:              unifieddatalibrary.Float(1.1),
			Senx:                 unifieddatalibrary.Float(1.1),
			Seny:                 unifieddatalibrary.Float(1.1),
			Senz:                 unifieddatalibrary.Float(1.1),
			ShutterDelay:         unifieddatalibrary.Float(1.1),
			SkyBkgrnd:            unifieddatalibrary.Float(1.1),
			SolarDecAngle:        unifieddatalibrary.Float(1.1),
			SolarEqPhaseAngle:    unifieddatalibrary.Float(1.1),
			SolarPhaseAngle:      unifieddatalibrary.Float(1.1),
			Tags:                 []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:               unifieddatalibrary.String("TASK-ID"),
			TimingBias:           unifieddatalibrary.Float(1.1),
			TrackID:              unifieddatalibrary.String("TRACK-ID"),
			TransactionID:        unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                  unifieddatalibrary.Bool(false),
			Umbra:                unifieddatalibrary.Bool(false),
			Zeroptd:              unifieddatalibrary.Float(1.1),
			ZeroPtdUnc:           unifieddatalibrary.Float(1.1),
		}},
		ConvertToJ2K: unifieddatalibrary.Bool(true),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationEoObservationQueryhelp(t *testing.T) {
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
	_, err := client.Observations.EoObservations.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationEoObservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.EoObservations.Tuple(context.TODO(), unifieddatalibrary.ObservationEoObservationTupleParams{
		Columns:     "columns",
		ObTime:      time.Now(),
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

func TestObservationEoObservationUnvalidatedPublish(t *testing.T) {
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
	err := client.Observations.EoObservations.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ObservationEoObservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ObservationEoObservationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("EOOBSERVATION-ID"),
			Azimuth:               unifieddatalibrary.Float(1.1),
			AzimuthBias:           unifieddatalibrary.Float(1.1),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.1),
			AzimuthUnc:            unifieddatalibrary.Float(1.1),
			BgIntensity:           unifieddatalibrary.Float(1.1),
			CollectMethod:         unifieddatalibrary.String("AUTOTRACK"),
			CorrQuality:           unifieddatalibrary.Float(1.1),
			Declination:           unifieddatalibrary.Float(1.1),
			DeclinationBias:       unifieddatalibrary.Float(1.1),
			DeclinationMeasured:   unifieddatalibrary.Bool(true),
			DeclinationRate:       unifieddatalibrary.Float(1.1),
			DeclinationUnc:        unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("PROVIDED_DATA1"),
			Elevation:             unifieddatalibrary.Float(1.1),
			ElevationBias:         unifieddatalibrary.Float(1.1),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(1.1),
			ElevationUnc:          unifieddatalibrary.Float(1.1),
			EoobservationDetails: unifieddatalibrary.ObservationEoObservationUnvalidatedPublishParamsBodyEoobservationDetails{
				AcalCrPixX:                           unifieddatalibrary.Float(123.2),
				AcalCrPixY:                           unifieddatalibrary.Float(123.2),
				AcalCrValX:                           unifieddatalibrary.Float(123.2),
				AcalCrValY:                           unifieddatalibrary.Float(123.2),
				AcalNumStars:                         unifieddatalibrary.Int(123),
				BackgroundSignal:                     unifieddatalibrary.Float(4134.1),
				BackgroundSignalUnc:                  unifieddatalibrary.Float(123.2),
				BinningHoriz:                         unifieddatalibrary.Int(12),
				BinningVert:                          unifieddatalibrary.Int(14),
				CcdObjPosX:                           unifieddatalibrary.Float(123.3),
				CcdObjPosY:                           unifieddatalibrary.Float(321.4),
				CcdObjWidth:                          unifieddatalibrary.Float(133.2),
				CcdTemp:                              unifieddatalibrary.Float(123.4),
				CentroidColumn:                       unifieddatalibrary.Float(0.5),
				CentroidRow:                          unifieddatalibrary.Float(0.1),
				ClassificationMarking:                unifieddatalibrary.String("U"),
				ColorCoeffs:                          []float64{1.1, 2.1, 3.1},
				ColumnVariance:                       unifieddatalibrary.Float(0.1),
				CurrentNeutralDensityFilterNum:       unifieddatalibrary.Int(3),
				CurrentSpectralFilterNum:             unifieddatalibrary.Int(23),
				DataMode:                             "TEST",
				DeclinationCov:                       unifieddatalibrary.Float(123.2),
				DistFromStreakCenter:                 []float64{-127.153, -126.153, -125.153},
				Does:                                 unifieddatalibrary.Float(123.2),
				ExtinctionCoeffs:                     []float64{1.1, 2.1, 3.1},
				ExtinctionCoeffsUnc:                  []float64{1.1, 2.1, 3.1},
				Gain:                                 unifieddatalibrary.Float(234.2),
				IDEoObservation:                      unifieddatalibrary.String("EOOBSERVATION-ID"),
				Ifov:                                 unifieddatalibrary.Float(0.2),
				MagInstrumental:                      unifieddatalibrary.Float(123.3),
				MagInstrumentalUnc:                   unifieddatalibrary.Float(123.3),
				NeutralDensityFilterNames:            []string{"numNeutralDensityFilters1", "numNeutralDensityFilters2"},
				NeutralDensityFilterTransmissions:    []float64{1.1, 2.1, 3.1},
				NeutralDensityFilterTransmissionsUnc: []float64{1.1, 2.1, 3.1},
				NumCatalogStars:                      unifieddatalibrary.Int(123),
				NumCorrelatedStars:                   unifieddatalibrary.Int(123),
				NumDetectedStars:                     unifieddatalibrary.Int(123),
				NumNeutralDensityFilters:             unifieddatalibrary.Int(12),
				NumSpectralFilters:                   unifieddatalibrary.Int(10),
				ObjSunRange:                          unifieddatalibrary.Float(123.2),
				ObTime:                               unifieddatalibrary.Time(time.Now()),
				OpticalCrossSection:                  unifieddatalibrary.Float(123.3),
				OpticalCrossSectionUnc:               unifieddatalibrary.Float(123.3),
				PcalNumStars:                         unifieddatalibrary.Int(23),
				PeakApertureCount:                    unifieddatalibrary.Float(123.2),
				PeakBackgroundCount:                  unifieddatalibrary.Int(321),
				PhaseAngBisect:                       unifieddatalibrary.Float(123.2),
				PixelArrayHeight:                     unifieddatalibrary.Int(23),
				PixelArrayWidth:                      unifieddatalibrary.Int(12),
				PixelMax:                             unifieddatalibrary.Int(256),
				PixelMin:                             unifieddatalibrary.Int(12),
				PredictedAzimuth:                     unifieddatalibrary.Float(10.1),
				PredictedDeclination:                 unifieddatalibrary.Float(10.1),
				PredictedDeclinationUnc:              unifieddatalibrary.Float(123.2),
				PredictedElevation:                   unifieddatalibrary.Float(10.1),
				PredictedRa:                          unifieddatalibrary.Float(10.1),
				PredictedRaUnc:                       unifieddatalibrary.Float(123.2),
				RaCov:                                unifieddatalibrary.Float(123.2),
				RaDeclinationCov:                     unifieddatalibrary.Float(123.2),
				RowColCov:                            unifieddatalibrary.Float(0.01),
				RowVariance:                          unifieddatalibrary.Float(0.1),
				SnrEst:                               unifieddatalibrary.Float(13.4),
				SolarDiskFrac:                        unifieddatalibrary.Float(123.2),
				Source:                               unifieddatalibrary.String("Bluestaq"),
				SpectralFilters:                      []string{"Keyword1", "Keyword2"},
				SpectralFilterSolarMag:               []float64{1.1, 2.1, 3.1},
				SpectralZmfl:                         []float64{1.1, 2.1, 3.1},
				SunAzimuth:                           unifieddatalibrary.Float(10.1),
				SunElevation:                         unifieddatalibrary.Float(10.1),
				SunStatePosX:                         unifieddatalibrary.Float(123.3),
				SunStatePosY:                         unifieddatalibrary.Float(123.3),
				SunStatePosZ:                         unifieddatalibrary.Float(123.3),
				SunStateVelX:                         unifieddatalibrary.Float(123.3),
				SunStateVelY:                         unifieddatalibrary.Float(123.3),
				SunStateVelZ:                         unifieddatalibrary.Float(123.3),
				SurfBrightness:                       []float64{21.01, 21.382, 21.725},
				SurfBrightnessUnc:                    []float64{0.165, 0.165, 0.165},
				TimesUnc:                             unifieddatalibrary.Float(13.1),
				Toes:                                 unifieddatalibrary.Float(123.2),
				ZeroPoints:                           []float64{1.1, 2.1, 3.1},
				ZeroPointsUnc:                        []float64{1.1, 2.1, 3.1},
			},
			ExpDuration:          unifieddatalibrary.Float(1.1),
			FovCount:             unifieddatalibrary.Int(1),
			FovCountUct:          unifieddatalibrary.Int(2),
			Geoalt:               unifieddatalibrary.Float(1.1),
			Geolat:               unifieddatalibrary.Float(1.1),
			Geolon:               unifieddatalibrary.Float(1.1),
			Georange:             unifieddatalibrary.Float(1.1),
			IDSensor:             unifieddatalibrary.String("SENSOR-ID"),
			IDSkyImagery:         unifieddatalibrary.String("SKYIMAGERY-ID"),
			Intensity:            unifieddatalibrary.Float(1.1),
			LosUnc:               unifieddatalibrary.Float(1.1),
			Losx:                 unifieddatalibrary.Float(1.1),
			Losxvel:              unifieddatalibrary.Float(1.1),
			Losy:                 unifieddatalibrary.Float(1.1),
			Losyvel:              unifieddatalibrary.Float(1.1),
			Losz:                 unifieddatalibrary.Float(1.1),
			Loszvel:              unifieddatalibrary.Float(1.1),
			Mag:                  unifieddatalibrary.Float(1.1),
			MagNormRange:         unifieddatalibrary.Float(1.1),
			MagUnc:               unifieddatalibrary.Float(1.1),
			NetObjSig:            unifieddatalibrary.Float(1.1),
			NetObjSigUnc:         unifieddatalibrary.Float(1.1),
			ObPosition:           unifieddatalibrary.String("FIRST"),
			Origin:               unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:         unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:         unifieddatalibrary.String("ORIGSENSOR-ID"),
			Penumbra:             unifieddatalibrary.Bool(false),
			PrimaryExtinction:    unifieddatalibrary.Float(1.1),
			PrimaryExtinctionUnc: unifieddatalibrary.Float(1.1),
			Ra:                   unifieddatalibrary.Float(1.1),
			RaBias:               unifieddatalibrary.Float(1.1),
			RaMeasured:           unifieddatalibrary.Bool(true),
			Range:                unifieddatalibrary.Float(1.1),
			RangeBias:            unifieddatalibrary.Float(1.1),
			RangeMeasured:        unifieddatalibrary.Bool(true),
			RangeRate:            unifieddatalibrary.Float(1.1),
			RangeRateMeasured:    unifieddatalibrary.Bool(true),
			RangeRateUnc:         unifieddatalibrary.Float(1.1),
			RangeUnc:             unifieddatalibrary.Float(1.1),
			RaRate:               unifieddatalibrary.Float(1.1),
			RaUnc:                unifieddatalibrary.Float(1.1),
			RawFileUri:           unifieddatalibrary.String("Example URI"),
			ReferenceFrame:       "J2000",
			SatNo:                unifieddatalibrary.Int(5),
			Senalt:               unifieddatalibrary.Float(1.1),
			Senlat:               unifieddatalibrary.Float(45.1),
			Senlon:               unifieddatalibrary.Float(179.1),
			SenQuat:              []float64{0.4492, 0.02, 0.8765, 0.2213},
			SenReferenceFrame:    "J2000",
			Senvelx:              unifieddatalibrary.Float(1.1),
			Senvely:              unifieddatalibrary.Float(1.1),
			Senvelz:              unifieddatalibrary.Float(1.1),
			Senx:                 unifieddatalibrary.Float(1.1),
			Seny:                 unifieddatalibrary.Float(1.1),
			Senz:                 unifieddatalibrary.Float(1.1),
			ShutterDelay:         unifieddatalibrary.Float(1.1),
			SkyBkgrnd:            unifieddatalibrary.Float(1.1),
			SolarDecAngle:        unifieddatalibrary.Float(1.1),
			SolarEqPhaseAngle:    unifieddatalibrary.Float(1.1),
			SolarPhaseAngle:      unifieddatalibrary.Float(1.1),
			Tags:                 []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:               unifieddatalibrary.String("TASK-ID"),
			TimingBias:           unifieddatalibrary.Float(1.1),
			TrackID:              unifieddatalibrary.String("TRACK-ID"),
			TransactionID:        unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                  unifieddatalibrary.Bool(false),
			Umbra:                unifieddatalibrary.Bool(false),
			Zeroptd:              unifieddatalibrary.Float(1.1),
			ZeroPtdUnc:           unifieddatalibrary.Float(1.1),
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
