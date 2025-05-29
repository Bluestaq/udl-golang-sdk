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

func TestObservationRfObservationNewWithOptionalParams(t *testing.T) {
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
	err := client.Observations.RfObservation.New(context.TODO(), unifieddatalibrary.ObservationRfObservationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ObservationRfObservationNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		Type:                  "RF",
		ID:                    unifieddatalibrary.String("RFOBSERVATION-ID"),
		AntennaName:           unifieddatalibrary.String("Antenna1"),
		Azimuth:               unifieddatalibrary.Float(10.1),
		AzimuthMeasured:       unifieddatalibrary.Bool(true),
		AzimuthRate:           unifieddatalibrary.Float(1.1),
		AzimuthUnc:            unifieddatalibrary.Float(2.1),
		Bandwidth:             unifieddatalibrary.Float(10.1),
		BaudRate:              unifieddatalibrary.Float(10.1),
		BaudRates:             []float64{1.1, 2.2},
		BitErrorRate:          unifieddatalibrary.Float(10.1),
		CarrierStandard:       unifieddatalibrary.String("DVB-S2"),
		Channel:               unifieddatalibrary.Int(10),
		ChipRates:             []float64{1.1, 2.2},
		CodeFills:             []string{"TAG1", "TAG2"},
		CodeLengths:           []float64{1.1, 2.2},
		CodeTaps:              []string{"TAG1", "TAG2"},
		CollectionMode:        unifieddatalibrary.String("SURVEY"),
		Confidence:            unifieddatalibrary.Float(10.1),
		Confidences:           []float64{1.1, 2.2},
		ConstellationXPoints:  []float64{1.1, 2.2},
		ConstellationYPoints:  []float64{1.1, 2.2},
		Descriptor:            unifieddatalibrary.String("descriptor"),
		DetectionStatus:       unifieddatalibrary.String("DETECTED"),
		DetectionStatuses:     []string{"DETECTED"},
		Eirp:                  unifieddatalibrary.Float(10.1),
		Elevation:             unifieddatalibrary.Float(10.1),
		ElevationMeasured:     unifieddatalibrary.Bool(true),
		ElevationRate:         unifieddatalibrary.Float(10.1),
		ElevationUnc:          unifieddatalibrary.Float(10.1),
		Elnot:                 unifieddatalibrary.String("Ex. ELINT"),
		EndFrequency:          unifieddatalibrary.Float(10.1),
		Frequencies:           []float64{1.1, 2.2},
		Frequency:             unifieddatalibrary.Float(10.1),
		FrequencyShift:        unifieddatalibrary.Float(10.1),
		IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
		Incoming:              unifieddatalibrary.Bool(false),
		InnerCodingRate:       unifieddatalibrary.Int(7),
		MaxPsd:                unifieddatalibrary.Float(10.1),
		MinPsd:                unifieddatalibrary.Float(10.1),
		Modulation:            unifieddatalibrary.String("Auto"),
		NoisePwrDensity:       unifieddatalibrary.Float(10.1),
		NominalBandwidth:      unifieddatalibrary.Float(10.1),
		NominalEirp:           unifieddatalibrary.Float(10.1),
		NominalFrequency:      unifieddatalibrary.Float(10.1),
		NominalPowerOverNoise: unifieddatalibrary.Float(10.1),
		NominalSnr:            unifieddatalibrary.Float(10.1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:          unifieddatalibrary.String("ORIG-OBJECT-ID"),
		OrigSensorID:          unifieddatalibrary.String("ORIG-SENSOR-ID"),
		OuterCodingRate:       unifieddatalibrary.Int(4),
		Peak:                  unifieddatalibrary.Bool(false),
		Pgri:                  unifieddatalibrary.Float(10.1),
		PnOrders:              []int64{1, 2},
		Polarity:              unifieddatalibrary.Float(10.1),
		PolarityType:          unifieddatalibrary.ObservationRfObservationNewParamsPolarityTypeH,
		PowerOverNoise:        unifieddatalibrary.Float(10.1),
		Powers:                []float64{1.1, 2.2},
		Range:                 unifieddatalibrary.Float(10.1),
		RangeMeasured:         unifieddatalibrary.Bool(true),
		RangeRate:             unifieddatalibrary.Float(10.1),
		RangeRateMeasured:     unifieddatalibrary.Bool(true),
		RangeRateUnc:          unifieddatalibrary.Float(10.1),
		RangeUnc:              unifieddatalibrary.Float(10.1),
		RawFileUri:            unifieddatalibrary.String("Example URI"),
		ReferenceLevel:        unifieddatalibrary.Float(10.1),
		RelativeCarrierPower:  unifieddatalibrary.Float(10.1),
		RelativeNoiseFloor:    unifieddatalibrary.Float(10.1),
		ResolutionBandwidth:   unifieddatalibrary.Float(10.1),
		SatNo:                 unifieddatalibrary.Int(32258),
		Senalt:                unifieddatalibrary.Float(10.1),
		Senlat:                unifieddatalibrary.Float(45.2),
		Senlon:                unifieddatalibrary.Float(80.3),
		SignalIDs:             []string{"ID1", "ID2"},
		Snr:                   unifieddatalibrary.Float(10.1),
		Snrs:                  []float64{1.1, 2.2},
		SpectrumAnalyzerPower: unifieddatalibrary.Float(10.1),
		StartFrequency:        unifieddatalibrary.Float(10.1),
		SwitchPoint:           unifieddatalibrary.Int(10),
		SymbolToNoiseRatio:    unifieddatalibrary.Float(10.1),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		TaskID:                unifieddatalibrary.String("TASK-ID"),
		TelemetryIDs:          []string{"ID1", "ID2"},
		TrackID:               unifieddatalibrary.String("TRACK-ID"),
		TrackRange:            unifieddatalibrary.Float(10.1),
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		TransmitFilterRollOff: unifieddatalibrary.Float(10.1),
		TransmitFilterType:    unifieddatalibrary.String("RRC"),
		Transponder:           unifieddatalibrary.String("TRANSPONDER-A"),
		Uct:                   unifieddatalibrary.Bool(false),
		URL:                   unifieddatalibrary.String("https://some-url"),
		VideoBandwidth:        unifieddatalibrary.Float(10.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationRfObservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.RfObservation.List(context.TODO(), unifieddatalibrary.ObservationRfObservationListParams{
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

func TestObservationRfObservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.RfObservation.Count(context.TODO(), unifieddatalibrary.ObservationRfObservationCountParams{
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

func TestObservationRfObservationNewBulk(t *testing.T) {
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
	err := client.Observations.RfObservation.NewBulk(context.TODO(), unifieddatalibrary.ObservationRfObservationNewBulkParams{
		Body: []unifieddatalibrary.ObservationRfObservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			Type:                  "RF",
			ID:                    unifieddatalibrary.String("RFOBSERVATION-ID"),
			AntennaName:           unifieddatalibrary.String("Antenna1"),
			Azimuth:               unifieddatalibrary.Float(10.1),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.1),
			AzimuthUnc:            unifieddatalibrary.Float(2.1),
			Bandwidth:             unifieddatalibrary.Float(10.1),
			BaudRate:              unifieddatalibrary.Float(10.1),
			BaudRates:             []float64{1.1, 2.2},
			BitErrorRate:          unifieddatalibrary.Float(10.1),
			CarrierStandard:       unifieddatalibrary.String("DVB-S2"),
			Channel:               unifieddatalibrary.Int(10),
			ChipRates:             []float64{1.1, 2.2},
			CodeFills:             []string{"TAG1", "TAG2"},
			CodeLengths:           []float64{1.1, 2.2},
			CodeTaps:              []string{"TAG1", "TAG2"},
			CollectionMode:        unifieddatalibrary.String("SURVEY"),
			Confidence:            unifieddatalibrary.Float(10.1),
			Confidences:           []float64{1.1, 2.2},
			ConstellationXPoints:  []float64{1.1, 2.2},
			ConstellationYPoints:  []float64{1.1, 2.2},
			Descriptor:            unifieddatalibrary.String("descriptor"),
			DetectionStatus:       unifieddatalibrary.String("DETECTED"),
			DetectionStatuses:     []string{"DETECTED"},
			Eirp:                  unifieddatalibrary.Float(10.1),
			Elevation:             unifieddatalibrary.Float(10.1),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(10.1),
			ElevationUnc:          unifieddatalibrary.Float(10.1),
			Elnot:                 unifieddatalibrary.String("Ex. ELINT"),
			EndFrequency:          unifieddatalibrary.Float(10.1),
			Frequencies:           []float64{1.1, 2.2},
			Frequency:             unifieddatalibrary.Float(10.1),
			FrequencyShift:        unifieddatalibrary.Float(10.1),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			Incoming:              unifieddatalibrary.Bool(false),
			InnerCodingRate:       unifieddatalibrary.Int(7),
			MaxPsd:                unifieddatalibrary.Float(10.1),
			MinPsd:                unifieddatalibrary.Float(10.1),
			Modulation:            unifieddatalibrary.String("Auto"),
			NoisePwrDensity:       unifieddatalibrary.Float(10.1),
			NominalBandwidth:      unifieddatalibrary.Float(10.1),
			NominalEirp:           unifieddatalibrary.Float(10.1),
			NominalFrequency:      unifieddatalibrary.Float(10.1),
			NominalPowerOverNoise: unifieddatalibrary.Float(10.1),
			NominalSnr:            unifieddatalibrary.Float(10.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIG-OBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIG-SENSOR-ID"),
			OuterCodingRate:       unifieddatalibrary.Int(4),
			Peak:                  unifieddatalibrary.Bool(false),
			Pgri:                  unifieddatalibrary.Float(10.1),
			PnOrders:              []int64{1, 2},
			Polarity:              unifieddatalibrary.Float(10.1),
			PolarityType:          "H",
			PowerOverNoise:        unifieddatalibrary.Float(10.1),
			Powers:                []float64{1.1, 2.2},
			Range:                 unifieddatalibrary.Float(10.1),
			RangeMeasured:         unifieddatalibrary.Bool(true),
			RangeRate:             unifieddatalibrary.Float(10.1),
			RangeRateMeasured:     unifieddatalibrary.Bool(true),
			RangeRateUnc:          unifieddatalibrary.Float(10.1),
			RangeUnc:              unifieddatalibrary.Float(10.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			ReferenceLevel:        unifieddatalibrary.Float(10.1),
			RelativeCarrierPower:  unifieddatalibrary.Float(10.1),
			RelativeNoiseFloor:    unifieddatalibrary.Float(10.1),
			ResolutionBandwidth:   unifieddatalibrary.Float(10.1),
			SatNo:                 unifieddatalibrary.Int(32258),
			Senalt:                unifieddatalibrary.Float(10.1),
			Senlat:                unifieddatalibrary.Float(45.2),
			Senlon:                unifieddatalibrary.Float(80.3),
			SignalIDs:             []string{"ID1", "ID2"},
			Snr:                   unifieddatalibrary.Float(10.1),
			Snrs:                  []float64{1.1, 2.2},
			SpectrumAnalyzerPower: unifieddatalibrary.Float(10.1),
			StartFrequency:        unifieddatalibrary.Float(10.1),
			SwitchPoint:           unifieddatalibrary.Int(10),
			SymbolToNoiseRatio:    unifieddatalibrary.Float(10.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TelemetryIDs:          []string{"ID1", "ID2"},
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TrackRange:            unifieddatalibrary.Float(10.1),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			TransmitFilterRollOff: unifieddatalibrary.Float(10.1),
			TransmitFilterType:    unifieddatalibrary.String("RRC"),
			Transponder:           unifieddatalibrary.String("TRANSPONDER-A"),
			Uct:                   unifieddatalibrary.Bool(false),
			URL:                   unifieddatalibrary.String("https://some-url"),
			VideoBandwidth:        unifieddatalibrary.Float(10.1),
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

func TestObservationRfObservationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.RfObservation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ObservationRfObservationGetParams{
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

func TestObservationRfObservationQueryhelp(t *testing.T) {
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
	_, err := client.Observations.RfObservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservationRfObservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Observations.RfObservation.Tuple(context.TODO(), unifieddatalibrary.ObservationRfObservationTupleParams{
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

func TestObservationRfObservationUnvalidatedPublish(t *testing.T) {
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
	err := client.Observations.RfObservation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ObservationRfObservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ObservationRfObservationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			Type:                  "RF",
			ID:                    unifieddatalibrary.String("RFOBSERVATION-ID"),
			AntennaName:           unifieddatalibrary.String("Antenna1"),
			Azimuth:               unifieddatalibrary.Float(10.1),
			AzimuthMeasured:       unifieddatalibrary.Bool(true),
			AzimuthRate:           unifieddatalibrary.Float(1.1),
			AzimuthUnc:            unifieddatalibrary.Float(2.1),
			Bandwidth:             unifieddatalibrary.Float(10.1),
			BaudRate:              unifieddatalibrary.Float(10.1),
			BaudRates:             []float64{1.1, 2.2},
			BitErrorRate:          unifieddatalibrary.Float(10.1),
			CarrierStandard:       unifieddatalibrary.String("DVB-S2"),
			Channel:               unifieddatalibrary.Int(10),
			ChipRates:             []float64{1.1, 2.2},
			CodeFills:             []string{"TAG1", "TAG2"},
			CodeLengths:           []float64{1.1, 2.2},
			CodeTaps:              []string{"TAG1", "TAG2"},
			CollectionMode:        unifieddatalibrary.String("SURVEY"),
			Confidence:            unifieddatalibrary.Float(10.1),
			Confidences:           []float64{1.1, 2.2},
			ConstellationXPoints:  []float64{1.1, 2.2},
			ConstellationYPoints:  []float64{1.1, 2.2},
			Descriptor:            unifieddatalibrary.String("descriptor"),
			DetectionStatus:       unifieddatalibrary.String("DETECTED"),
			DetectionStatuses:     []string{"DETECTED"},
			Eirp:                  unifieddatalibrary.Float(10.1),
			Elevation:             unifieddatalibrary.Float(10.1),
			ElevationMeasured:     unifieddatalibrary.Bool(true),
			ElevationRate:         unifieddatalibrary.Float(10.1),
			ElevationUnc:          unifieddatalibrary.Float(10.1),
			Elnot:                 unifieddatalibrary.String("Ex. ELINT"),
			EndFrequency:          unifieddatalibrary.Float(10.1),
			Frequencies:           []float64{1.1, 2.2},
			Frequency:             unifieddatalibrary.Float(10.1),
			FrequencyShift:        unifieddatalibrary.Float(10.1),
			IDSensor:              unifieddatalibrary.String("SENSOR-ID"),
			Incoming:              unifieddatalibrary.Bool(false),
			InnerCodingRate:       unifieddatalibrary.Int(7),
			MaxPsd:                unifieddatalibrary.Float(10.1),
			MinPsd:                unifieddatalibrary.Float(10.1),
			Modulation:            unifieddatalibrary.String("Auto"),
			NoisePwrDensity:       unifieddatalibrary.Float(10.1),
			NominalBandwidth:      unifieddatalibrary.Float(10.1),
			NominalEirp:           unifieddatalibrary.Float(10.1),
			NominalFrequency:      unifieddatalibrary.Float(10.1),
			NominalPowerOverNoise: unifieddatalibrary.Float(10.1),
			NominalSnr:            unifieddatalibrary.Float(10.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIG-OBJECT-ID"),
			OrigSensorID:          unifieddatalibrary.String("ORIG-SENSOR-ID"),
			OuterCodingRate:       unifieddatalibrary.Int(4),
			Peak:                  unifieddatalibrary.Bool(false),
			Pgri:                  unifieddatalibrary.Float(10.1),
			PnOrders:              []int64{1, 2},
			Polarity:              unifieddatalibrary.Float(10.1),
			PolarityType:          "H",
			PowerOverNoise:        unifieddatalibrary.Float(10.1),
			Powers:                []float64{1.1, 2.2},
			Range:                 unifieddatalibrary.Float(10.1),
			RangeMeasured:         unifieddatalibrary.Bool(true),
			RangeRate:             unifieddatalibrary.Float(10.1),
			RangeRateMeasured:     unifieddatalibrary.Bool(true),
			RangeRateUnc:          unifieddatalibrary.Float(10.1),
			RangeUnc:              unifieddatalibrary.Float(10.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			ReferenceLevel:        unifieddatalibrary.Float(10.1),
			RelativeCarrierPower:  unifieddatalibrary.Float(10.1),
			RelativeNoiseFloor:    unifieddatalibrary.Float(10.1),
			ResolutionBandwidth:   unifieddatalibrary.Float(10.1),
			SatNo:                 unifieddatalibrary.Int(32258),
			Senalt:                unifieddatalibrary.Float(10.1),
			Senlat:                unifieddatalibrary.Float(45.2),
			Senlon:                unifieddatalibrary.Float(80.3),
			SignalIDs:             []string{"ID1", "ID2"},
			Snr:                   unifieddatalibrary.Float(10.1),
			Snrs:                  []float64{1.1, 2.2},
			SpectrumAnalyzerPower: unifieddatalibrary.Float(10.1),
			StartFrequency:        unifieddatalibrary.Float(10.1),
			SwitchPoint:           unifieddatalibrary.Int(10),
			SymbolToNoiseRatio:    unifieddatalibrary.Float(10.1),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TaskID:                unifieddatalibrary.String("TASK-ID"),
			TelemetryIDs:          []string{"ID1", "ID2"},
			TrackID:               unifieddatalibrary.String("TRACK-ID"),
			TrackRange:            unifieddatalibrary.Float(10.1),
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			TransmitFilterRollOff: unifieddatalibrary.Float(10.1),
			TransmitFilterType:    unifieddatalibrary.String("RRC"),
			Transponder:           unifieddatalibrary.String("TRANSPONDER-A"),
			Uct:                   unifieddatalibrary.Bool(false),
			URL:                   unifieddatalibrary.String("https://some-url"),
			VideoBandwidth:        unifieddatalibrary.Float(10.1),
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
