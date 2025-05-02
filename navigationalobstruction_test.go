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

func TestNavigationalobstructionNewWithOptionalParams(t *testing.T) {
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
	err := client.Navigationalobstruction.New(context.TODO(), unifieddatalibrary.NavigationalobstructionNewParams{
		ClassificationMarking: "U",
		CycleDate:             time.Now(),
		DataMode:              unifieddatalibrary.NavigationalobstructionNewParamsDataModeTest,
		ObstacleID:            "359655",
		ObstacleType:          "V",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		ActDelCode:            unifieddatalibrary.String("A"),
		AiracCycle:            unifieddatalibrary.Int(2406),
		BaseAiracCycle:        unifieddatalibrary.Int(2405),
		BaselineCutoffDate:    unifieddatalibrary.Time(time.Now()),
		BoundNeLat:            unifieddatalibrary.Float(29.1),
		BoundNeLon:            unifieddatalibrary.Float(99.1),
		BoundSwLat:            unifieddatalibrary.Float(-44.1),
		BoundSwLon:            unifieddatalibrary.Float(-144.1),
		CountryCode:           unifieddatalibrary.String("US"),
		CutoffDate:            unifieddatalibrary.Time(time.Now()),
		DataSetRemarks:        unifieddatalibrary.String("Data set remarks"),
		DeletingOrg:           unifieddatalibrary.String("ACME"),
		DerivingOrg:           unifieddatalibrary.String("ACME"),
		DirectivityCode:       unifieddatalibrary.Int(2),
		Elevation:             unifieddatalibrary.Float(840.1),
		ElevationAcc:          unifieddatalibrary.Float(17.1),
		ExternalID:            unifieddatalibrary.String("OU812"),
		Facc:                  unifieddatalibrary.String("AT040"),
		FeatureCode:           unifieddatalibrary.String("540"),
		FeatureDescription:    unifieddatalibrary.String("Powerline Pylon, General"),
		FeatureName:           unifieddatalibrary.String("PYLON"),
		FeatureType:           unifieddatalibrary.String("540"),
		HeightAgl:             unifieddatalibrary.Float(314.1),
		HeightAglAcc:          unifieddatalibrary.Float(30.1),
		HeightMsl:             unifieddatalibrary.Float(1154.1),
		HeightMslAcc:          unifieddatalibrary.Float(34.1),
		HorizAcc:              unifieddatalibrary.Float(8.1),
		HorizDatumCode:        unifieddatalibrary.String("WGS-84"),
		InitRecordDate:        unifieddatalibrary.Time(time.Now()),
		Keys:                  []string{"key1", "key2"},
		LightingCode:          unifieddatalibrary.String("U"),
		LineNeLat:             unifieddatalibrary.Float(49.000584),
		LineNeLon:             unifieddatalibrary.Float(-122.197891),
		LinesFilename:         unifieddatalibrary.String("lines.txt"),
		LineSwLat:             unifieddatalibrary.Float(48.507027),
		LineSwLon:             unifieddatalibrary.Float(-122.722946),
		MinHeightAgl:          unifieddatalibrary.Float(20.1),
		MultObs:               unifieddatalibrary.String("S"),
		NextCycleDate:         unifieddatalibrary.Time(time.Now()),
		NumLines:              unifieddatalibrary.Int(45993),
		NumObs:                unifieddatalibrary.Int(1),
		NumPoints:             unifieddatalibrary.Int(21830590),
		ObstacleRemarks:       unifieddatalibrary.String("Obstacle remarks"),
		OrigID:                unifieddatalibrary.String("L0000002289"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OwnerCountryCode:      unifieddatalibrary.String("US"),
		PointLat:              unifieddatalibrary.Float(46.757211),
		PointLon:              unifieddatalibrary.Float(-67.759494),
		PointsFilename:        unifieddatalibrary.String("points.txt"),
		ProcessCode:           unifieddatalibrary.String("OT"),
		Producer:              unifieddatalibrary.String("ACME"),
		ProvinceCode:          unifieddatalibrary.String("23"),
		Quality:               unifieddatalibrary.String("0"),
		RevDate:               unifieddatalibrary.Time(time.Now()),
		SegEndPoint:           unifieddatalibrary.Int(359655),
		SegNum:                unifieddatalibrary.Int(1),
		SegStartPoint:         unifieddatalibrary.Int(359655),
		SourceDate:            unifieddatalibrary.Time(time.Now()),
		SurfaceMatCode:        unifieddatalibrary.String("U"),
		TransactionCode:       unifieddatalibrary.String("V"),
		ValidationCode:        unifieddatalibrary.Int(3),
		Values:                []string{"value1", "value2"},
		VectorsFilename:       unifieddatalibrary.String("vectors.txt"),
		Wac:                   unifieddatalibrary.String("262"),
		WacInnr:               unifieddatalibrary.String("0409-00039"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNavigationalobstructionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Navigationalobstruction.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.NavigationalobstructionUpdateParams{
			ClassificationMarking: "U",
			CycleDate:             time.Now(),
			DataMode:              unifieddatalibrary.NavigationalobstructionUpdateParamsDataModeTest,
			ObstacleID:            "359655",
			ObstacleType:          "V",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActDelCode:            unifieddatalibrary.String("A"),
			AiracCycle:            unifieddatalibrary.Int(2406),
			BaseAiracCycle:        unifieddatalibrary.Int(2405),
			BaselineCutoffDate:    unifieddatalibrary.Time(time.Now()),
			BoundNeLat:            unifieddatalibrary.Float(29.1),
			BoundNeLon:            unifieddatalibrary.Float(99.1),
			BoundSwLat:            unifieddatalibrary.Float(-44.1),
			BoundSwLon:            unifieddatalibrary.Float(-144.1),
			CountryCode:           unifieddatalibrary.String("US"),
			CutoffDate:            unifieddatalibrary.Time(time.Now()),
			DataSetRemarks:        unifieddatalibrary.String("Data set remarks"),
			DeletingOrg:           unifieddatalibrary.String("ACME"),
			DerivingOrg:           unifieddatalibrary.String("ACME"),
			DirectivityCode:       unifieddatalibrary.Int(2),
			Elevation:             unifieddatalibrary.Float(840.1),
			ElevationAcc:          unifieddatalibrary.Float(17.1),
			ExternalID:            unifieddatalibrary.String("OU812"),
			Facc:                  unifieddatalibrary.String("AT040"),
			FeatureCode:           unifieddatalibrary.String("540"),
			FeatureDescription:    unifieddatalibrary.String("Powerline Pylon, General"),
			FeatureName:           unifieddatalibrary.String("PYLON"),
			FeatureType:           unifieddatalibrary.String("540"),
			HeightAgl:             unifieddatalibrary.Float(314.1),
			HeightAglAcc:          unifieddatalibrary.Float(30.1),
			HeightMsl:             unifieddatalibrary.Float(1154.1),
			HeightMslAcc:          unifieddatalibrary.Float(34.1),
			HorizAcc:              unifieddatalibrary.Float(8.1),
			HorizDatumCode:        unifieddatalibrary.String("WGS-84"),
			InitRecordDate:        unifieddatalibrary.Time(time.Now()),
			Keys:                  []string{"key1", "key2"},
			LightingCode:          unifieddatalibrary.String("U"),
			LineNeLat:             unifieddatalibrary.Float(49.000584),
			LineNeLon:             unifieddatalibrary.Float(-122.197891),
			LinesFilename:         unifieddatalibrary.String("lines.txt"),
			LineSwLat:             unifieddatalibrary.Float(48.507027),
			LineSwLon:             unifieddatalibrary.Float(-122.722946),
			MinHeightAgl:          unifieddatalibrary.Float(20.1),
			MultObs:               unifieddatalibrary.String("S"),
			NextCycleDate:         unifieddatalibrary.Time(time.Now()),
			NumLines:              unifieddatalibrary.Int(45993),
			NumObs:                unifieddatalibrary.Int(1),
			NumPoints:             unifieddatalibrary.Int(21830590),
			ObstacleRemarks:       unifieddatalibrary.String("Obstacle remarks"),
			OrigID:                unifieddatalibrary.String("L0000002289"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OwnerCountryCode:      unifieddatalibrary.String("US"),
			PointLat:              unifieddatalibrary.Float(46.757211),
			PointLon:              unifieddatalibrary.Float(-67.759494),
			PointsFilename:        unifieddatalibrary.String("points.txt"),
			ProcessCode:           unifieddatalibrary.String("OT"),
			Producer:              unifieddatalibrary.String("ACME"),
			ProvinceCode:          unifieddatalibrary.String("23"),
			Quality:               unifieddatalibrary.String("0"),
			RevDate:               unifieddatalibrary.Time(time.Now()),
			SegEndPoint:           unifieddatalibrary.Int(359655),
			SegNum:                unifieddatalibrary.Int(1),
			SegStartPoint:         unifieddatalibrary.Int(359655),
			SourceDate:            unifieddatalibrary.Time(time.Now()),
			SurfaceMatCode:        unifieddatalibrary.String("U"),
			TransactionCode:       unifieddatalibrary.String("V"),
			ValidationCode:        unifieddatalibrary.Int(3),
			Values:                []string{"value1", "value2"},
			VectorsFilename:       unifieddatalibrary.String("vectors.txt"),
			Wac:                   unifieddatalibrary.String("262"),
			WacInnr:               unifieddatalibrary.String("0409-00039"),
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

func TestNavigationalobstructionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigationalobstruction.List(context.TODO(), unifieddatalibrary.NavigationalobstructionListParams{
		CycleDate:   unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		ObstacleID:  unifieddatalibrary.String("obstacleId"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNavigationalobstructionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigationalobstruction.Count(context.TODO(), unifieddatalibrary.NavigationalobstructionCountParams{
		CycleDate:   unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		ObstacleID:  unifieddatalibrary.String("obstacleId"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNavigationalobstructionNewBulk(t *testing.T) {
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
	err := client.Navigationalobstruction.NewBulk(context.TODO(), unifieddatalibrary.NavigationalobstructionNewBulkParams{
		Body: []unifieddatalibrary.NavigationalobstructionNewBulkParamsBody{{
			ClassificationMarking: "U",
			CycleDate:             time.Now(),
			DataMode:              "TEST",
			ObstacleID:            "359655",
			ObstacleType:          "V",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActDelCode:            unifieddatalibrary.String("A"),
			AiracCycle:            unifieddatalibrary.Int(2406),
			BaseAiracCycle:        unifieddatalibrary.Int(2405),
			BaselineCutoffDate:    unifieddatalibrary.Time(time.Now()),
			BoundNeLat:            unifieddatalibrary.Float(29.1),
			BoundNeLon:            unifieddatalibrary.Float(99.1),
			BoundSwLat:            unifieddatalibrary.Float(-44.1),
			BoundSwLon:            unifieddatalibrary.Float(-144.1),
			CountryCode:           unifieddatalibrary.String("US"),
			CutoffDate:            unifieddatalibrary.Time(time.Now()),
			DataSetRemarks:        unifieddatalibrary.String("Data set remarks"),
			DeletingOrg:           unifieddatalibrary.String("ACME"),
			DerivingOrg:           unifieddatalibrary.String("ACME"),
			DirectivityCode:       unifieddatalibrary.Int(2),
			Elevation:             unifieddatalibrary.Float(840.1),
			ElevationAcc:          unifieddatalibrary.Float(17.1),
			ExternalID:            unifieddatalibrary.String("OU812"),
			Facc:                  unifieddatalibrary.String("AT040"),
			FeatureCode:           unifieddatalibrary.String("540"),
			FeatureDescription:    unifieddatalibrary.String("Powerline Pylon, General"),
			FeatureName:           unifieddatalibrary.String("PYLON"),
			FeatureType:           unifieddatalibrary.String("540"),
			HeightAgl:             unifieddatalibrary.Float(314.1),
			HeightAglAcc:          unifieddatalibrary.Float(30.1),
			HeightMsl:             unifieddatalibrary.Float(1154.1),
			HeightMslAcc:          unifieddatalibrary.Float(34.1),
			HorizAcc:              unifieddatalibrary.Float(8.1),
			HorizDatumCode:        unifieddatalibrary.String("WGS-84"),
			InitRecordDate:        unifieddatalibrary.Time(time.Now()),
			Keys:                  []string{"key1", "key2"},
			LightingCode:          unifieddatalibrary.String("U"),
			LineNeLat:             unifieddatalibrary.Float(49.000584),
			LineNeLon:             unifieddatalibrary.Float(-122.197891),
			LinesFilename:         unifieddatalibrary.String("lines.txt"),
			LineSwLat:             unifieddatalibrary.Float(48.507027),
			LineSwLon:             unifieddatalibrary.Float(-122.722946),
			MinHeightAgl:          unifieddatalibrary.Float(20.1),
			MultObs:               unifieddatalibrary.String("S"),
			NextCycleDate:         unifieddatalibrary.Time(time.Now()),
			NumLines:              unifieddatalibrary.Int(45993),
			NumObs:                unifieddatalibrary.Int(1),
			NumPoints:             unifieddatalibrary.Int(21830590),
			ObstacleRemarks:       unifieddatalibrary.String("Obstacle remarks"),
			OrigID:                unifieddatalibrary.String("L0000002289"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OwnerCountryCode:      unifieddatalibrary.String("US"),
			PointLat:              unifieddatalibrary.Float(46.757211),
			PointLon:              unifieddatalibrary.Float(-67.759494),
			PointsFilename:        unifieddatalibrary.String("points.txt"),
			ProcessCode:           unifieddatalibrary.String("OT"),
			Producer:              unifieddatalibrary.String("ACME"),
			ProvinceCode:          unifieddatalibrary.String("23"),
			Quality:               unifieddatalibrary.String("0"),
			RevDate:               unifieddatalibrary.Time(time.Now()),
			SegEndPoint:           unifieddatalibrary.Int(359655),
			SegNum:                unifieddatalibrary.Int(1),
			SegStartPoint:         unifieddatalibrary.Int(359655),
			SourceDate:            unifieddatalibrary.Time(time.Now()),
			SurfaceMatCode:        unifieddatalibrary.String("U"),
			TransactionCode:       unifieddatalibrary.String("V"),
			ValidationCode:        unifieddatalibrary.Int(3),
			Values:                []string{"value1", "value2"},
			VectorsFilename:       unifieddatalibrary.String("vectors.txt"),
			Wac:                   unifieddatalibrary.String("262"),
			WacInnr:               unifieddatalibrary.String("0409-00039"),
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

func TestNavigationalobstructionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigationalobstruction.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.NavigationalobstructionGetParams{
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

func TestNavigationalobstructionQueryhelp(t *testing.T) {
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
	err := client.Navigationalobstruction.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNavigationalobstructionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigationalobstruction.Tuple(context.TODO(), unifieddatalibrary.NavigationalobstructionTupleParams{
		Columns:     "columns",
		CycleDate:   unifieddatalibrary.Time(time.Now()),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		ObstacleID:  unifieddatalibrary.String("obstacleId"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
