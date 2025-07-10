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

func TestSurfaceNewWithOptionalParams(t *testing.T) {
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
	err := client.Surface.New(context.TODO(), unifieddatalibrary.SurfaceNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SurfaceNewParamsDataModeTest,
		Name:                  "West lot",
		Source:                "Bluestaq",
		Type:                  "PARKING",
		ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
		AltSiteID:             unifieddatalibrary.String("ORIG-SITE-ID"),
		Condition:             unifieddatalibrary.String("GOOD"),
		DdtWtKip:              unifieddatalibrary.Float(833.5),
		DdtWtKipMod:           unifieddatalibrary.Float(625.125),
		DdtWtKipModNote:       unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		DdtWtKn:               unifieddatalibrary.Float(3705.44),
		DdWtKip:               unifieddatalibrary.Float(416.25),
		DdWtKipMod:            unifieddatalibrary.Float(312.1875),
		DdWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		DdWtKn:                unifieddatalibrary.Float(1850.1),
		EndLat:                unifieddatalibrary.Float(45.844),
		EndLon:                unifieddatalibrary.Float(-67.0094),
		IDSite:                unifieddatalibrary.String("SITE-ID"),
		Lcn:                   unifieddatalibrary.Int(50),
		LdaFt:                 unifieddatalibrary.Float(475.2),
		LdaM:                  unifieddatalibrary.Float(145.75),
		LengthFt:              unifieddatalibrary.Float(1500.75),
		LengthM:               unifieddatalibrary.Float(457.5),
		Lighting:              unifieddatalibrary.Bool(true),
		LightsAprch:           unifieddatalibrary.Bool(true),
		LightsCl:              unifieddatalibrary.Bool(true),
		LightsOls:             unifieddatalibrary.Bool(true),
		LightsPapi:            unifieddatalibrary.Bool(true),
		LightsReil:            unifieddatalibrary.Bool(true),
		LightsRwy:             unifieddatalibrary.Bool(true),
		LightsSeqfl:           unifieddatalibrary.Bool(true),
		LightsTdzl:            unifieddatalibrary.Bool(true),
		LightsUnkn:            unifieddatalibrary.Bool(false),
		LightsVasi:            unifieddatalibrary.Bool(true),
		Material:              unifieddatalibrary.String("Concrete"),
		Obstacle:              unifieddatalibrary.Bool(true),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Pcn:                   unifieddatalibrary.String("73RBWT"),
		PointReference:        unifieddatalibrary.String("Reference Point Example"),
		Primary:               unifieddatalibrary.Bool(true),
		RawWbc:                unifieddatalibrary.String("LCN 42"),
		SbttWtKip:             unifieddatalibrary.Float(603.025),
		SbttWtKipMod:          unifieddatalibrary.Float(452.269),
		SbttWtKipModNote:      unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		SbttWtKn:              unifieddatalibrary.Float(2682.05),
		StartLat:              unifieddatalibrary.Float(46.757211),
		StartLon:              unifieddatalibrary.Float(-67.759494),
		StWtKip:               unifieddatalibrary.Float(195.75),
		StWtKipMod:            unifieddatalibrary.Float(146.8125),
		StWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		StWtKn:                unifieddatalibrary.Float(867.6),
		SWtKip:                unifieddatalibrary.Float(143.5),
		SWtKipMod:             unifieddatalibrary.Float(107.625),
		SWtKipModNote:         unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		SWtKn:                 unifieddatalibrary.Float(636.4),
		TdtWtkip:              unifieddatalibrary.Float(870.2),
		TdtWtKipMod:           unifieddatalibrary.Float(652.65),
		TdtWtKipModNote:       unifieddatalibrary.String("Serious 25 percent reduction."),
		TdtWtKn:               unifieddatalibrary.Float(3870.05),
		TrtWtKip:              unifieddatalibrary.Float(622.85),
		TrtWtKipMod:           unifieddatalibrary.Float(467.1375),
		TrtWtKipModNote:       unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		TrtWtKn:               unifieddatalibrary.Float(2767.15),
		TtWtKip:               unifieddatalibrary.Float(414.9),
		TtWtKipMod:            unifieddatalibrary.Float(311.175),
		TtWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		TtWtKn:                unifieddatalibrary.Float(1842.55),
		TWtKip:                unifieddatalibrary.Float(188.2),
		TWtKipMod:             unifieddatalibrary.Float(141.15),
		TWtKipModNote:         unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
		TWtKn:                 unifieddatalibrary.Float(836.45),
		WidthFt:               unifieddatalibrary.Float(220.3),
		WidthM:                unifieddatalibrary.Float(67.25),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Surface.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SurfaceUpdateParamsDataModeTest,
			Name:                  "West lot",
			Source:                "Bluestaq",
			Type:                  "PARKING",
			ID:                    unifieddatalibrary.String("be831d39-1822-da9f-7ace-6cc5643397dc"),
			AltSiteID:             unifieddatalibrary.String("ORIG-SITE-ID"),
			Condition:             unifieddatalibrary.String("GOOD"),
			DdtWtKip:              unifieddatalibrary.Float(833.5),
			DdtWtKipMod:           unifieddatalibrary.Float(625.125),
			DdtWtKipModNote:       unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			DdtWtKn:               unifieddatalibrary.Float(3705.44),
			DdWtKip:               unifieddatalibrary.Float(416.25),
			DdWtKipMod:            unifieddatalibrary.Float(312.1875),
			DdWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			DdWtKn:                unifieddatalibrary.Float(1850.1),
			EndLat:                unifieddatalibrary.Float(45.844),
			EndLon:                unifieddatalibrary.Float(-67.0094),
			IDSite:                unifieddatalibrary.String("SITE-ID"),
			Lcn:                   unifieddatalibrary.Int(50),
			LdaFt:                 unifieddatalibrary.Float(475.2),
			LdaM:                  unifieddatalibrary.Float(145.75),
			LengthFt:              unifieddatalibrary.Float(1500.75),
			LengthM:               unifieddatalibrary.Float(457.5),
			Lighting:              unifieddatalibrary.Bool(true),
			LightsAprch:           unifieddatalibrary.Bool(true),
			LightsCl:              unifieddatalibrary.Bool(true),
			LightsOls:             unifieddatalibrary.Bool(true),
			LightsPapi:            unifieddatalibrary.Bool(true),
			LightsReil:            unifieddatalibrary.Bool(true),
			LightsRwy:             unifieddatalibrary.Bool(true),
			LightsSeqfl:           unifieddatalibrary.Bool(true),
			LightsTdzl:            unifieddatalibrary.Bool(true),
			LightsUnkn:            unifieddatalibrary.Bool(false),
			LightsVasi:            unifieddatalibrary.Bool(true),
			Material:              unifieddatalibrary.String("Concrete"),
			Obstacle:              unifieddatalibrary.Bool(true),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Pcn:                   unifieddatalibrary.String("73RBWT"),
			PointReference:        unifieddatalibrary.String("Reference Point Example"),
			Primary:               unifieddatalibrary.Bool(true),
			RawWbc:                unifieddatalibrary.String("LCN 42"),
			SbttWtKip:             unifieddatalibrary.Float(603.025),
			SbttWtKipMod:          unifieddatalibrary.Float(452.269),
			SbttWtKipModNote:      unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			SbttWtKn:              unifieddatalibrary.Float(2682.05),
			StartLat:              unifieddatalibrary.Float(46.757211),
			StartLon:              unifieddatalibrary.Float(-67.759494),
			StWtKip:               unifieddatalibrary.Float(195.75),
			StWtKipMod:            unifieddatalibrary.Float(146.8125),
			StWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			StWtKn:                unifieddatalibrary.Float(867.6),
			SWtKip:                unifieddatalibrary.Float(143.5),
			SWtKipMod:             unifieddatalibrary.Float(107.625),
			SWtKipModNote:         unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			SWtKn:                 unifieddatalibrary.Float(636.4),
			TdtWtkip:              unifieddatalibrary.Float(870.2),
			TdtWtKipMod:           unifieddatalibrary.Float(652.65),
			TdtWtKipModNote:       unifieddatalibrary.String("Serious 25 percent reduction."),
			TdtWtKn:               unifieddatalibrary.Float(3870.05),
			TrtWtKip:              unifieddatalibrary.Float(622.85),
			TrtWtKipMod:           unifieddatalibrary.Float(467.1375),
			TrtWtKipModNote:       unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			TrtWtKn:               unifieddatalibrary.Float(2767.15),
			TtWtKip:               unifieddatalibrary.Float(414.9),
			TtWtKipMod:            unifieddatalibrary.Float(311.175),
			TtWtKipModNote:        unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			TtWtKn:                unifieddatalibrary.Float(1842.55),
			TWtKip:                unifieddatalibrary.Float(188.2),
			TWtKipMod:             unifieddatalibrary.Float(141.15),
			TWtKipModNote:         unifieddatalibrary.String("Observed 25% reduction in surface integrity."),
			TWtKn:                 unifieddatalibrary.Float(836.45),
			WidthFt:               unifieddatalibrary.Float(220.3),
			WidthM:                unifieddatalibrary.Float(67.25),
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

func TestSurfaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Surface.List(context.TODO(), unifieddatalibrary.SurfaceListParams{
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

func TestSurfaceDelete(t *testing.T) {
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
	err := client.Surface.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Surface.Count(context.TODO(), unifieddatalibrary.SurfaceCountParams{
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

func TestSurfaceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Surface.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SurfaceGetParams{
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

func TestSurfaceQueryhelp(t *testing.T) {
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
	_, err := client.Surface.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSurfaceTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Surface.Tuple(context.TODO(), unifieddatalibrary.SurfaceTupleParams{
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
