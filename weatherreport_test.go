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

func TestWeatherReportNewWithOptionalParams(t *testing.T) {
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
	err := client.WeatherReport.New(context.TODO(), unifieddatalibrary.WeatherReportNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.WeatherReportNewParamsDataModeTest,
		Lat:                   56.12,
		Lon:                   -156.6,
		ObTime:                time.Now(),
		ReportType:            "FORECAST",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("WEATHER-REPORT-ID"),
		ActWeather:            unifieddatalibrary.String("NO STATEMENT"),
		Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
		Alt:                   unifieddatalibrary.Float(123.12),
		Andims:                unifieddatalibrary.Int(2),
		Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Asrid:                 unifieddatalibrary.Int(4326),
		Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Atype:                 unifieddatalibrary.String("ST_Polygon"),
		BarPress:              unifieddatalibrary.Float(101.2),
		CcEvent:               unifieddatalibrary.Bool(true),
		CloudCover:            []string{"OVERCAST", "BROKEN"},
		CloudHght:             []float64{1.2, 2.2},
		ContrailHghtLower:     unifieddatalibrary.Float(123.123),
		ContrailHghtUpper:     unifieddatalibrary.Float(123.123),
		DataLevel:             unifieddatalibrary.String("MANDATORY"),
		DewPoint:              unifieddatalibrary.Float(15.6),
		DifRad:                unifieddatalibrary.Float(234.5),
		DirDev:                unifieddatalibrary.Float(9.1),
		EnRouteWeather:        unifieddatalibrary.String("THUNDERSTORMS"),
		ExternalID:            unifieddatalibrary.String("GDSSMB022408301601304517"),
		ExternalLocationID:    unifieddatalibrary.String("TMDS060AD4OG03CC"),
		ForecastEndTime:       unifieddatalibrary.Time(time.Now()),
		ForecastStartTime:     unifieddatalibrary.Time(time.Now()),
		GeoPotentialAlt:       unifieddatalibrary.Float(1000),
		Hshear:                unifieddatalibrary.Float(3.8),
		Icao:                  unifieddatalibrary.String("KAFF"),
		IcingLowerLimit:       unifieddatalibrary.Float(123.123),
		IcingUpperLimit:       unifieddatalibrary.Float(123.123),
		IDAirfield:            unifieddatalibrary.String("8fb38d6d-a3de-45dd-8974-4e3ed73e9449"),
		IDGroundImagery:       unifieddatalibrary.String("GROUND-IMAGERY-ID"),
		IDSensor:              unifieddatalibrary.String("0129f577-e04c-441e-65ca-0a04a750bed9"),
		IDSite:                unifieddatalibrary.String("AIRFIELD-ID"),
		IndexRefraction:       unifieddatalibrary.Float(1.1),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
		PrecipRate:            unifieddatalibrary.Float(3.4),
		Qnh:                   unifieddatalibrary.Float(1234.456),
		RadVel:                unifieddatalibrary.Float(-0.04),
		RadVelBeam1:           unifieddatalibrary.Float(4.4),
		RadVelBeam2:           unifieddatalibrary.Float(-0.2),
		RadVelBeam3:           unifieddatalibrary.Float(-0.2),
		RadVelBeam4:           unifieddatalibrary.Float(11.4),
		RadVelBeam5:           unifieddatalibrary.Float(4.1),
		RainHour:              unifieddatalibrary.Float(1.2),
		RawMetar:              unifieddatalibrary.String("KXYZ 241456Z 19012G20KT 10SM FEW120 SCT200 BKN250 26/M04 A2981 RMK AO2 PK WND 19026/1420 SLP068 T02611039 51015"),
		RawTaf:                unifieddatalibrary.String("KXYZ 051730Z 0518/0624 31008KT 3SM -SHRA BKN020 FM052300 30006KT 5SM -SHRA OVC030 PROB30 0604/0606 VRB20G35KT 1SM TSRA BKN015CB FM060600 25010KT 4SM -SHRA OVC050 TEMPO 0608/0611 2SM -SHRA OVC030 RMK NXT FCST BY 00Z="),
		RefRad:                unifieddatalibrary.Float(56.7),
		RelHumidity:           unifieddatalibrary.Float(34.456),
		Senalt:                unifieddatalibrary.Float(1.23),
		Senlat:                unifieddatalibrary.Float(12.456),
		Senlon:                unifieddatalibrary.Float(123.456),
		SoilMoisture:          unifieddatalibrary.Float(3.5),
		SoilTemp:              unifieddatalibrary.Float(22.4),
		SolarRad:              unifieddatalibrary.Float(1234.456),
		SrcIDs:                []string{"e609a90d-4059-4043-9f1a-fd7b49a3e1d0", "c739fcdb-c0c9-43c0-97b6-bfc80d0ffd52"},
		SrcTyps:               []string{"SENSOR", "WEATHERDATA"},
		SurroundingWeather:    unifieddatalibrary.String("NO STATEMENT"),
		Temperature:           unifieddatalibrary.Float(23.45),
		Visibility:            unifieddatalibrary.Float(1234.456),
		Vshear:                unifieddatalibrary.Float(3.8),
		WeatherAmp:            unifieddatalibrary.String("NO STATEMENT"),
		WeatherDesc:           unifieddatalibrary.String("NO STATEMENT"),
		WeatherID:             unifieddatalibrary.String("WEATHER-ID"),
		WeatherInt:            unifieddatalibrary.String("NO STATEMENT"),
		WindChill:             unifieddatalibrary.Float(15.6),
		WindCov:               []float64{1.1, 2.2},
		WindDir:               unifieddatalibrary.Float(75.1234),
		WindDirAvg:            unifieddatalibrary.Float(57.1),
		WindDirPeak:           unifieddatalibrary.Float(78.4),
		WindDirPeak10:         unifieddatalibrary.Float(44.5),
		WindGust:              unifieddatalibrary.Float(10.23),
		WindGust10:            unifieddatalibrary.Float(13.2),
		WindSpd:               unifieddatalibrary.Float(1.23),
		WindSpdAvg:            unifieddatalibrary.Float(12.1),
		WindVar:               unifieddatalibrary.Bool(false),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWeatherReportListWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherReport.List(context.TODO(), unifieddatalibrary.WeatherReportListParams{
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

func TestWeatherReportCountWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherReport.Count(context.TODO(), unifieddatalibrary.WeatherReportCountParams{
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

func TestWeatherReportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherReport.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.WeatherReportGetParams{
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

func TestWeatherReportQueryhelp(t *testing.T) {
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
	_, err := client.WeatherReport.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWeatherReportTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.WeatherReport.Tuple(context.TODO(), unifieddatalibrary.WeatherReportTupleParams{
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

func TestWeatherReportUnvalidatedPublish(t *testing.T) {
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
	err := client.WeatherReport.UnvalidatedPublish(context.TODO(), unifieddatalibrary.WeatherReportUnvalidatedPublishParams{
		Body: []unifieddatalibrary.WeatherReportUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Lat:                   56.12,
			Lon:                   -156.6,
			ObTime:                time.Now(),
			ReportType:            "FORECAST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("WEATHER-REPORT-ID"),
			ActWeather:            unifieddatalibrary.String("NO STATEMENT"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Alt:                   unifieddatalibrary.Float(123.12),
			Andims:                unifieddatalibrary.Int(2),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(4326),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("ST_Polygon"),
			BarPress:              unifieddatalibrary.Float(101.2),
			CcEvent:               unifieddatalibrary.Bool(true),
			CloudCover:            []string{"OVERCAST", "BROKEN"},
			CloudHght:             []float64{1.2, 2.2},
			ContrailHghtLower:     unifieddatalibrary.Float(123.123),
			ContrailHghtUpper:     unifieddatalibrary.Float(123.123),
			DataLevel:             unifieddatalibrary.String("MANDATORY"),
			DewPoint:              unifieddatalibrary.Float(15.6),
			DifRad:                unifieddatalibrary.Float(234.5),
			DirDev:                unifieddatalibrary.Float(9.1),
			EnRouteWeather:        unifieddatalibrary.String("THUNDERSTORMS"),
			ExternalID:            unifieddatalibrary.String("GDSSMB022408301601304517"),
			ExternalLocationID:    unifieddatalibrary.String("TMDS060AD4OG03CC"),
			ForecastEndTime:       unifieddatalibrary.Time(time.Now()),
			ForecastStartTime:     unifieddatalibrary.Time(time.Now()),
			GeoPotentialAlt:       unifieddatalibrary.Float(1000),
			Hshear:                unifieddatalibrary.Float(3.8),
			Icao:                  unifieddatalibrary.String("KAFF"),
			IcingLowerLimit:       unifieddatalibrary.Float(123.123),
			IcingUpperLimit:       unifieddatalibrary.Float(123.123),
			IDAirfield:            unifieddatalibrary.String("8fb38d6d-a3de-45dd-8974-4e3ed73e9449"),
			IDGroundImagery:       unifieddatalibrary.String("GROUND-IMAGERY-ID"),
			IDSensor:              unifieddatalibrary.String("0129f577-e04c-441e-65ca-0a04a750bed9"),
			IDSite:                unifieddatalibrary.String("AIRFIELD-ID"),
			IndexRefraction:       unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:          unifieddatalibrary.String("ORIGSENSOR-ID"),
			PrecipRate:            unifieddatalibrary.Float(3.4),
			Qnh:                   unifieddatalibrary.Float(1234.456),
			RadVel:                unifieddatalibrary.Float(-0.04),
			RadVelBeam1:           unifieddatalibrary.Float(4.4),
			RadVelBeam2:           unifieddatalibrary.Float(-0.2),
			RadVelBeam3:           unifieddatalibrary.Float(-0.2),
			RadVelBeam4:           unifieddatalibrary.Float(11.4),
			RadVelBeam5:           unifieddatalibrary.Float(4.1),
			RainHour:              unifieddatalibrary.Float(1.2),
			RawMetar:              unifieddatalibrary.String("KXYZ 241456Z 19012G20KT 10SM FEW120 SCT200 BKN250 26/M04 A2981 RMK AO2 PK WND 19026/1420 SLP068 T02611039 51015"),
			RawTaf:                unifieddatalibrary.String("KXYZ 051730Z 0518/0624 31008KT 3SM -SHRA BKN020 FM052300 30006KT 5SM -SHRA OVC030 PROB30 0604/0606 VRB20G35KT 1SM TSRA BKN015CB FM060600 25010KT 4SM -SHRA OVC050 TEMPO 0608/0611 2SM -SHRA OVC030 RMK NXT FCST BY 00Z="),
			RefRad:                unifieddatalibrary.Float(56.7),
			RelHumidity:           unifieddatalibrary.Float(34.456),
			Senalt:                unifieddatalibrary.Float(1.23),
			Senlat:                unifieddatalibrary.Float(12.456),
			Senlon:                unifieddatalibrary.Float(123.456),
			SoilMoisture:          unifieddatalibrary.Float(3.5),
			SoilTemp:              unifieddatalibrary.Float(22.4),
			SolarRad:              unifieddatalibrary.Float(1234.456),
			SrcIDs:                []string{"e609a90d-4059-4043-9f1a-fd7b49a3e1d0", "c739fcdb-c0c9-43c0-97b6-bfc80d0ffd52"},
			SrcTyps:               []string{"SENSOR", "WEATHERDATA"},
			SurroundingWeather:    unifieddatalibrary.String("NO STATEMENT"),
			Temperature:           unifieddatalibrary.Float(23.45),
			Visibility:            unifieddatalibrary.Float(1234.456),
			Vshear:                unifieddatalibrary.Float(3.8),
			WeatherAmp:            unifieddatalibrary.String("NO STATEMENT"),
			WeatherDesc:           unifieddatalibrary.String("NO STATEMENT"),
			WeatherID:             unifieddatalibrary.String("WEATHER-ID"),
			WeatherInt:            unifieddatalibrary.String("NO STATEMENT"),
			WindChill:             unifieddatalibrary.Float(15.6),
			WindCov:               []float64{1.1, 2.2},
			WindDir:               unifieddatalibrary.Float(75.1234),
			WindDirAvg:            unifieddatalibrary.Float(57.1),
			WindDirPeak:           unifieddatalibrary.Float(78.4),
			WindDirPeak10:         unifieddatalibrary.Float(44.5),
			WindGust:              unifieddatalibrary.Float(10.23),
			WindGust10:            unifieddatalibrary.Float(13.2),
			WindSpd:               unifieddatalibrary.Float(1.23),
			WindSpdAvg:            unifieddatalibrary.Float(12.1),
			WindVar:               unifieddatalibrary.Bool(false),
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
