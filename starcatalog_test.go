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

func TestStarCatalogNewWithOptionalParams(t *testing.T) {
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
	err := client.StarCatalog.New(context.TODO(), unifieddatalibrary.StarCatalogNewParams{
		AstrometryOrigin:      unifieddatalibrary.StarCatalogNewParamsAstrometryOriginGaiadr3,
		ClassificationMarking: "U",
		CsID:                  12345,
		DataMode:              unifieddatalibrary.StarCatalogNewParamsDataModeTest,
		Dec:                   21.8,
		Ra:                    14.43,
		Source:                "Bluestaq",
		StarEpoch:             2016,
		ID:                    unifieddatalibrary.String("STAR-CAT-DATASET-ID"),
		Bpmag:                 unifieddatalibrary.Float(0.04559),
		BpmagUnc:              unifieddatalibrary.Float(0.2227),
		CatVersion:            unifieddatalibrary.String("1.23"),
		DecUnc:                unifieddatalibrary.Float(40.996),
		Gaiadr3CatID:          unifieddatalibrary.Int(89012345678901),
		Gmag:                  unifieddatalibrary.Float(0.0046),
		GmagUnc:               unifieddatalibrary.Float(0.00292),
		GncCatID:              unifieddatalibrary.Int(12345),
		HipCatID:              unifieddatalibrary.Int(12345),
		Hmag:                  unifieddatalibrary.Float(12.126),
		HmagUnc:               unifieddatalibrary.Float(5.722),
		Jmag:                  unifieddatalibrary.Float(9.515),
		JmagUnc:               unifieddatalibrary.Float(7.559),
		Kmag:                  unifieddatalibrary.Float(13.545),
		KmagUnc:               unifieddatalibrary.Float(0.052),
		MultFlag:              unifieddatalibrary.Bool(true),
		NeighborDistance:      unifieddatalibrary.Float(201.406),
		NeighborFlag:          unifieddatalibrary.Bool(false),
		NeighborID:            unifieddatalibrary.Int(2456),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Parallax:              unifieddatalibrary.Float(-6.8),
		ParallaxUnc:           unifieddatalibrary.Float(82.35),
		Pmdec:                 unifieddatalibrary.Float(-970.1003),
		PmdecUnc:              unifieddatalibrary.Float(1.22),
		Pmra:                  unifieddatalibrary.Float(1000.45),
		PmraUnc:               unifieddatalibrary.Float(5.6),
		PmUncFlag:             unifieddatalibrary.Bool(false),
		PosUncFlag:            unifieddatalibrary.Bool(false),
		RaUnc:                 unifieddatalibrary.Float(509.466),
		Rpmag:                 unifieddatalibrary.Float(8.0047),
		RpmagUnc:              unifieddatalibrary.Float(1.233),
		Shift:                 unifieddatalibrary.Float(4.548),
		ShiftFlag:             unifieddatalibrary.Bool(false),
		VarFlag:               unifieddatalibrary.Bool(true),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogUpdateWithOptionalParams(t *testing.T) {
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
	err := client.StarCatalog.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.StarCatalogUpdateParams{
			AstrometryOrigin:      unifieddatalibrary.StarCatalogUpdateParamsAstrometryOriginGaiadr3,
			ClassificationMarking: "U",
			CsID:                  12345,
			DataMode:              unifieddatalibrary.StarCatalogUpdateParamsDataModeTest,
			Dec:                   21.8,
			Ra:                    14.43,
			Source:                "Bluestaq",
			StarEpoch:             2016,
			ID:                    unifieddatalibrary.String("STAR-CAT-DATASET-ID"),
			Bpmag:                 unifieddatalibrary.Float(0.04559),
			BpmagUnc:              unifieddatalibrary.Float(0.2227),
			CatVersion:            unifieddatalibrary.String("1.23"),
			DecUnc:                unifieddatalibrary.Float(40.996),
			Gaiadr3CatID:          unifieddatalibrary.Int(89012345678901),
			Gmag:                  unifieddatalibrary.Float(0.0046),
			GmagUnc:               unifieddatalibrary.Float(0.00292),
			GncCatID:              unifieddatalibrary.Int(12345),
			HipCatID:              unifieddatalibrary.Int(12345),
			Hmag:                  unifieddatalibrary.Float(12.126),
			HmagUnc:               unifieddatalibrary.Float(5.722),
			Jmag:                  unifieddatalibrary.Float(9.515),
			JmagUnc:               unifieddatalibrary.Float(7.559),
			Kmag:                  unifieddatalibrary.Float(13.545),
			KmagUnc:               unifieddatalibrary.Float(0.052),
			MultFlag:              unifieddatalibrary.Bool(true),
			NeighborDistance:      unifieddatalibrary.Float(201.406),
			NeighborFlag:          unifieddatalibrary.Bool(false),
			NeighborID:            unifieddatalibrary.Int(2456),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Parallax:              unifieddatalibrary.Float(-6.8),
			ParallaxUnc:           unifieddatalibrary.Float(82.35),
			Pmdec:                 unifieddatalibrary.Float(-970.1003),
			PmdecUnc:              unifieddatalibrary.Float(1.22),
			Pmra:                  unifieddatalibrary.Float(1000.45),
			PmraUnc:               unifieddatalibrary.Float(5.6),
			PmUncFlag:             unifieddatalibrary.Bool(false),
			PosUncFlag:            unifieddatalibrary.Bool(false),
			RaUnc:                 unifieddatalibrary.Float(509.466),
			Rpmag:                 unifieddatalibrary.Float(8.0047),
			RpmagUnc:              unifieddatalibrary.Float(1.233),
			Shift:                 unifieddatalibrary.Float(4.548),
			ShiftFlag:             unifieddatalibrary.Bool(false),
			VarFlag:               unifieddatalibrary.Bool(true),
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

func TestStarCatalogListWithOptionalParams(t *testing.T) {
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
	_, err := client.StarCatalog.List(context.TODO(), unifieddatalibrary.StarCatalogListParams{
		Dec:         unifieddatalibrary.Float(0),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		Ra:          unifieddatalibrary.Float(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogDelete(t *testing.T) {
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
	err := client.StarCatalog.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogCountWithOptionalParams(t *testing.T) {
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
	_, err := client.StarCatalog.Count(context.TODO(), unifieddatalibrary.StarCatalogCountParams{
		Dec:         unifieddatalibrary.Float(0),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		Ra:          unifieddatalibrary.Float(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogNewBulk(t *testing.T) {
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
	err := client.StarCatalog.NewBulk(context.TODO(), unifieddatalibrary.StarCatalogNewBulkParams{
		Body: []unifieddatalibrary.StarCatalogNewBulkParamsBody{{
			AstrometryOrigin:      "GAIADR3",
			ClassificationMarking: "U",
			CsID:                  12345,
			DataMode:              "TEST",
			Dec:                   21.8,
			Ra:                    14.43,
			Source:                "Bluestaq",
			StarEpoch:             2016,
			ID:                    unifieddatalibrary.String("STAR-CAT-DATASET-ID"),
			Bpmag:                 unifieddatalibrary.Float(0.04559),
			BpmagUnc:              unifieddatalibrary.Float(0.2227),
			CatVersion:            unifieddatalibrary.String("1.23"),
			DecUnc:                unifieddatalibrary.Float(40.996),
			Gaiadr3CatID:          unifieddatalibrary.Int(89012345678901),
			Gmag:                  unifieddatalibrary.Float(0.0046),
			GmagUnc:               unifieddatalibrary.Float(0.00292),
			GncCatID:              unifieddatalibrary.Int(12345),
			HipCatID:              unifieddatalibrary.Int(12345),
			Hmag:                  unifieddatalibrary.Float(12.126),
			HmagUnc:               unifieddatalibrary.Float(5.722),
			Jmag:                  unifieddatalibrary.Float(9.515),
			JmagUnc:               unifieddatalibrary.Float(7.559),
			Kmag:                  unifieddatalibrary.Float(13.545),
			KmagUnc:               unifieddatalibrary.Float(0.052),
			MultFlag:              unifieddatalibrary.Bool(true),
			NeighborDistance:      unifieddatalibrary.Float(201.406),
			NeighborFlag:          unifieddatalibrary.Bool(false),
			NeighborID:            unifieddatalibrary.Int(2456),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Parallax:              unifieddatalibrary.Float(-6.8),
			ParallaxUnc:           unifieddatalibrary.Float(82.35),
			Pmdec:                 unifieddatalibrary.Float(-970.1003),
			PmdecUnc:              unifieddatalibrary.Float(1.22),
			Pmra:                  unifieddatalibrary.Float(1000.45),
			PmraUnc:               unifieddatalibrary.Float(5.6),
			PmUncFlag:             unifieddatalibrary.Bool(false),
			PosUncFlag:            unifieddatalibrary.Bool(false),
			RaUnc:                 unifieddatalibrary.Float(509.466),
			Rpmag:                 unifieddatalibrary.Float(8.0047),
			RpmagUnc:              unifieddatalibrary.Float(1.233),
			Shift:                 unifieddatalibrary.Float(4.548),
			ShiftFlag:             unifieddatalibrary.Bool(false),
			VarFlag:               unifieddatalibrary.Bool(true),
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

func TestStarCatalogGetWithOptionalParams(t *testing.T) {
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
	_, err := client.StarCatalog.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.StarCatalogGetParams{
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

func TestStarCatalogQueryhelp(t *testing.T) {
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
	_, err := client.StarCatalog.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.StarCatalog.Tuple(context.TODO(), unifieddatalibrary.StarCatalogTupleParams{
		Columns:     "columns",
		Dec:         unifieddatalibrary.Float(0),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		Ra:          unifieddatalibrary.Float(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarCatalogUnvalidatedPublish(t *testing.T) {
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
	err := client.StarCatalog.UnvalidatedPublish(context.TODO(), unifieddatalibrary.StarCatalogUnvalidatedPublishParams{
		Body: []unifieddatalibrary.StarCatalogUnvalidatedPublishParamsBody{{
			AstrometryOrigin:      "GAIADR3",
			ClassificationMarking: "U",
			CsID:                  12345,
			DataMode:              "TEST",
			Dec:                   21.8,
			Ra:                    14.43,
			Source:                "Bluestaq",
			StarEpoch:             2016,
			ID:                    unifieddatalibrary.String("STAR-CAT-DATASET-ID"),
			Bpmag:                 unifieddatalibrary.Float(0.04559),
			BpmagUnc:              unifieddatalibrary.Float(0.2227),
			CatVersion:            unifieddatalibrary.String("1.23"),
			DecUnc:                unifieddatalibrary.Float(40.996),
			Gaiadr3CatID:          unifieddatalibrary.Int(89012345678901),
			Gmag:                  unifieddatalibrary.Float(0.0046),
			GmagUnc:               unifieddatalibrary.Float(0.00292),
			GncCatID:              unifieddatalibrary.Int(12345),
			HipCatID:              unifieddatalibrary.Int(12345),
			Hmag:                  unifieddatalibrary.Float(12.126),
			HmagUnc:               unifieddatalibrary.Float(5.722),
			Jmag:                  unifieddatalibrary.Float(9.515),
			JmagUnc:               unifieddatalibrary.Float(7.559),
			Kmag:                  unifieddatalibrary.Float(13.545),
			KmagUnc:               unifieddatalibrary.Float(0.052),
			MultFlag:              unifieddatalibrary.Bool(true),
			NeighborDistance:      unifieddatalibrary.Float(201.406),
			NeighborFlag:          unifieddatalibrary.Bool(false),
			NeighborID:            unifieddatalibrary.Int(2456),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Parallax:              unifieddatalibrary.Float(-6.8),
			ParallaxUnc:           unifieddatalibrary.Float(82.35),
			Pmdec:                 unifieddatalibrary.Float(-970.1003),
			PmdecUnc:              unifieddatalibrary.Float(1.22),
			Pmra:                  unifieddatalibrary.Float(1000.45),
			PmraUnc:               unifieddatalibrary.Float(5.6),
			PmUncFlag:             unifieddatalibrary.Bool(false),
			PosUncFlag:            unifieddatalibrary.Bool(false),
			RaUnc:                 unifieddatalibrary.Float(509.466),
			Rpmag:                 unifieddatalibrary.Float(8.0047),
			RpmagUnc:              unifieddatalibrary.Float(1.233),
			Shift:                 unifieddatalibrary.Float(4.548),
			ShiftFlag:             unifieddatalibrary.Bool(false),
			VarFlag:               unifieddatalibrary.Bool(true),
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
