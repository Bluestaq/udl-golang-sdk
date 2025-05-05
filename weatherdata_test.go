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

func TestWeatherDataNewWithOptionalParams(t *testing.T) {
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
	err := client.WeatherData.New(context.TODO(), unifieddatalibrary.WeatherDataNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.WeatherDataNewParamsDataModeTest,
		ObTime:                time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("WEATHER-DATA-ID"),
		AngleOrientation:      unifieddatalibrary.Float(75.7),
		AvgRefPwr:             unifieddatalibrary.Float(714.9),
		AvgTxPwr:              unifieddatalibrary.Float(20.23),
		Checksum:              unifieddatalibrary.Int(133),
		CoIntegs:              []int64{4, 3},
		ConsRecs:              []int64{5, 2},
		DoppVels:              []float64{44.4, 467.3},
		FileCreation:          unifieddatalibrary.Time(time.Now()),
		FirstGuessAvgs:        []int64{16, 1},
		IDSensor:              unifieddatalibrary.String("0129f577-e04c-441e-65ca-0a04a750bed9"),
		InterpulsePeriods:     []float64{1000.3, 1000.2},
		LightDetSensors:       []int64{11, 28, 190},
		LightEventNum:         unifieddatalibrary.Int(9),
		NoiseLvls:             []float64{58.2, 58.3},
		NumElements:           unifieddatalibrary.Int(640),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		PosConfidence:         unifieddatalibrary.Float(0.1),
		QcValue:               unifieddatalibrary.Int(4),
		SectorNum:             unifieddatalibrary.Int(20),
		SemiMajorAxis:         unifieddatalibrary.Float(3.4),
		SemiMinorAxis:         unifieddatalibrary.Float(0.3),
		SigPwrs:               []float64{116.5, 121.6},
		SigStrength:           unifieddatalibrary.Float(163.7),
		Snrs:                  []float64{14.5, -16.2},
		SpecAvgs:              []int64{4, 3},
		SpecWidths:            []float64{0.3, 0.6},
		SrcIDs:                []string{"1b23ba93-0957-4654-b5ca-8c3703f3ec57", "32944ee4-0437-4d94-95ce-2f2823ffa001"},
		SrcTyps:               []string{"SENSOR", "WEATHERREPORT"},
		TdAvgSampleNums:       []int64{32, 30},
		TermAlt:               unifieddatalibrary.Float(19505.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWeatherDataListWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherData.List(context.TODO(), unifieddatalibrary.WeatherDataListParams{
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

func TestWeatherDataCountWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherData.Count(context.TODO(), unifieddatalibrary.WeatherDataCountParams{
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

func TestWeatherDataNewBulk(t *testing.T) {
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
	err := client.WeatherData.NewBulk(context.TODO(), unifieddatalibrary.WeatherDataNewBulkParams{
		Body: []unifieddatalibrary.WeatherDataNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("WEATHER-DATA-ID"),
			AngleOrientation:      unifieddatalibrary.Float(75.7),
			AvgRefPwr:             unifieddatalibrary.Float(714.9),
			AvgTxPwr:              unifieddatalibrary.Float(20.23),
			Checksum:              unifieddatalibrary.Int(133),
			CoIntegs:              []int64{4, 3},
			ConsRecs:              []int64{5, 2},
			DoppVels:              []float64{44.4, 467.3},
			FileCreation:          unifieddatalibrary.Time(time.Now()),
			FirstGuessAvgs:        []int64{16, 1},
			IDSensor:              unifieddatalibrary.String("0129f577-e04c-441e-65ca-0a04a750bed9"),
			InterpulsePeriods:     []float64{1000.3, 1000.2},
			LightDetSensors:       []int64{11, 28, 190},
			LightEventNum:         unifieddatalibrary.Int(9),
			NoiseLvls:             []float64{58.2, 58.3},
			NumElements:           unifieddatalibrary.Int(640),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			PosConfidence:         unifieddatalibrary.Float(0.1),
			QcValue:               unifieddatalibrary.Int(4),
			SectorNum:             unifieddatalibrary.Int(20),
			SemiMajorAxis:         unifieddatalibrary.Float(3.4),
			SemiMinorAxis:         unifieddatalibrary.Float(0.3),
			SigPwrs:               []float64{116.5, 121.6},
			SigStrength:           unifieddatalibrary.Float(163.7),
			Snrs:                  []float64{14.5, -16.2},
			SpecAvgs:              []int64{4, 3},
			SpecWidths:            []float64{0.3, 0.6},
			SrcIDs:                []string{"1b23ba93-0957-4654-b5ca-8c3703f3ec57", "32944ee4-0437-4d94-95ce-2f2823ffa001"},
			SrcTyps:               []string{"SENSOR", "WEATHERREPORT"},
			TdAvgSampleNums:       []int64{32, 30},
			TermAlt:               unifieddatalibrary.Float(19505.1),
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

func TestWeatherDataGetWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherData.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.WeatherDataGetParams{
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

func TestWeatherDataQueryhelp(t *testing.T) {
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
	err := client.WeatherData.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWeatherDataTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherData.Tuple(context.TODO(), unifieddatalibrary.WeatherDataTupleParams{
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

func TestWeatherDataUnvalidatedPublish(t *testing.T) {
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
	err := client.WeatherData.UnvalidatedPublish(context.TODO(), unifieddatalibrary.WeatherDataUnvalidatedPublishParams{
		Body: []unifieddatalibrary.WeatherDataUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ObTime:                time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("WEATHER-DATA-ID"),
			AngleOrientation:      unifieddatalibrary.Float(75.7),
			AvgRefPwr:             unifieddatalibrary.Float(714.9),
			AvgTxPwr:              unifieddatalibrary.Float(20.23),
			Checksum:              unifieddatalibrary.Int(133),
			CoIntegs:              []int64{4, 3},
			ConsRecs:              []int64{5, 2},
			DoppVels:              []float64{44.4, 467.3},
			FileCreation:          unifieddatalibrary.Time(time.Now()),
			FirstGuessAvgs:        []int64{16, 1},
			IDSensor:              unifieddatalibrary.String("0129f577-e04c-441e-65ca-0a04a750bed9"),
			InterpulsePeriods:     []float64{1000.3, 1000.2},
			LightDetSensors:       []int64{11, 28, 190},
			LightEventNum:         unifieddatalibrary.Int(9),
			NoiseLvls:             []float64{58.2, 58.3},
			NumElements:           unifieddatalibrary.Int(640),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			PosConfidence:         unifieddatalibrary.Float(0.1),
			QcValue:               unifieddatalibrary.Int(4),
			SectorNum:             unifieddatalibrary.Int(20),
			SemiMajorAxis:         unifieddatalibrary.Float(3.4),
			SemiMinorAxis:         unifieddatalibrary.Float(0.3),
			SigPwrs:               []float64{116.5, 121.6},
			SigStrength:           unifieddatalibrary.Float(163.7),
			Snrs:                  []float64{14.5, -16.2},
			SpecAvgs:              []int64{4, 3},
			SpecWidths:            []float64{0.3, 0.6},
			SrcIDs:                []string{"1b23ba93-0957-4654-b5ca-8c3703f3ec57", "32944ee4-0437-4d94-95ce-2f2823ffa001"},
			SrcTyps:               []string{"SENSOR", "WEATHERREPORT"},
			TdAvgSampleNums:       []int64{32, 30},
			TermAlt:               unifieddatalibrary.Float(19505.1),
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
