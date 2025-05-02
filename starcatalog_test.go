// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestStarcatalogNewWithOptionalParams(t *testing.T) {
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
	err := client.Starcatalog.New(context.TODO(), unifieddatalibrary.StarcatalogNewParams{
		AstrometryOrigin:      unifieddatalibrary.StarcatalogNewParamsAstrometryOriginGaiadr3,
		ClassificationMarking: "U",
		CsID:                  12345,
		DataMode:              unifieddatalibrary.StarcatalogNewParamsDataModeTest,
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

func TestStarcatalogUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Starcatalog.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.StarcatalogUpdateParams{
			AstrometryOrigin:      unifieddatalibrary.StarcatalogUpdateParamsAstrometryOriginGaiadr3,
			ClassificationMarking: "U",
			CsID:                  12345,
			DataMode:              unifieddatalibrary.StarcatalogUpdateParamsDataModeTest,
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

func TestStarcatalogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Starcatalog.List(context.TODO(), unifieddatalibrary.StarcatalogListParams{
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

func TestStarcatalogDelete(t *testing.T) {
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
	err := client.Starcatalog.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarcatalogCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Starcatalog.Count(context.TODO(), unifieddatalibrary.StarcatalogCountParams{
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

func TestStarcatalogNewBulk(t *testing.T) {
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
	err := client.Starcatalog.NewBulk(context.TODO(), unifieddatalibrary.StarcatalogNewBulkParams{
		Body: []unifieddatalibrary.StarcatalogNewBulkParamsBody{{
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

func TestStarcatalogGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Starcatalog.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.StarcatalogGetParams{
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

func TestStarcatalogQueryhelp(t *testing.T) {
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
	err := client.Starcatalog.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStarcatalogTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Starcatalog.Tuple(context.TODO(), unifieddatalibrary.StarcatalogTupleParams{
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

func TestStarcatalogUnvalidatedPublish(t *testing.T) {
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
	err := client.Starcatalog.UnvalidatedPublish(context.TODO(), unifieddatalibrary.StarcatalogUnvalidatedPublishParams{
		Body: []unifieddatalibrary.StarcatalogUnvalidatedPublishParamsBody{{
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
