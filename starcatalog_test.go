// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk/v2"
	"github.com/Bluestaq/udl-golang-sdk/v2/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/v2/option"
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
		AstrometryOrigin:            "GA",
		ClassificationMarking:       "U",
		CsID:                        12345,
		DataMode:                    unifieddatalibrary.StarCatalogNewParamsDataModeTest,
		Dec:                         21.8,
		Ra:                          14.43,
		Source:                      "Bluestaq",
		StarEpoch:                   2018.864,
		AavsoVsxID:                  unifieddatalibrary.Int(2387137),
		Abgmag:                      unifieddatalibrary.Float(11.24),
		AbgmagOrigin:                unifieddatalibrary.String("PS"),
		AbgmagUnc:                   unifieddatalibrary.Float(0.023),
		Abimag:                      unifieddatalibrary.Float(11.24),
		AbimagOrigin:                unifieddatalibrary.String("GA"),
		AbimagUnc:                   unifieddatalibrary.Float(0.023),
		Abrmag:                      unifieddatalibrary.Float(11.24),
		AbrmagOrigin:                unifieddatalibrary.String("SK"),
		AbrmagUnc:                   unifieddatalibrary.Float(0.023),
		Abymag:                      unifieddatalibrary.Float(11.24),
		AbymagOrigin:                unifieddatalibrary.String("AP"),
		AbymagUnc:                   unifieddatalibrary.Float(0.023),
		Abzmag:                      unifieddatalibrary.Float(11.24),
		AbzmagOrigin:                unifieddatalibrary.String("AP"),
		AbzmagUnc:                   unifieddatalibrary.Float(0.023),
		AllWisEccInd:                unifieddatalibrary.String("0000"),
		AllWiseID:                   unifieddatalibrary.String("WISEA J152743.04+624823.6"),
		AllWisEnaInd:                unifieddatalibrary.Int(1),
		AllWisEphQualInd:            unifieddatalibrary.String("AAAA"),
		ApassID:                     unifieddatalibrary.String("175-0082419"),
		AstrometricExcessNoise:      unifieddatalibrary.Float(921.2857),
		AstrometricExcessNoiseSig:   unifieddatalibrary.Float(194326670.1),
		Bmag:                        unifieddatalibrary.Float(27.004),
		BmagOrigin:                  unifieddatalibrary.String("AP"),
		BmagUnc:                     unifieddatalibrary.Float(9.999),
		Bpmag:                       unifieddatalibrary.Float(0.04559),
		BpmagUnc:                    unifieddatalibrary.Float(0.2227),
		CarrascoCatID:               unifieddatalibrary.Int(1244),
		CatVersion:                  unifieddatalibrary.String("1.23 DR3"),
		CatWise2020ID:               unifieddatalibrary.String("3584p196_b0-084425"),
		DecUnc:                      unifieddatalibrary.Float(40.996),
		DucatiCatID:                 unifieddatalibrary.String("AB ORI"),
		Gaiadr3CatID:                unifieddatalibrary.Int(89012345678901),
		Gmag:                        unifieddatalibrary.Float(0.0046),
		GmagUnc:                     unifieddatalibrary.Float(0.00292),
		GncCatID:                    unifieddatalibrary.Int(12345),
		HealpixIndex:                unifieddatalibrary.Int(196607),
		HipCatID:                    unifieddatalibrary.Int(12345),
		Hmag:                        unifieddatalibrary.Float(12.126),
		HmagOrigin:                  unifieddatalibrary.String("UL"),
		HmagUnc:                     unifieddatalibrary.Float(5.722),
		Imag:                        unifieddatalibrary.Float(22.46249),
		ImagOrigin:                  unifieddatalibrary.String("HI"),
		ImagUnc:                     unifieddatalibrary.Float(1.2000417),
		Jmag:                        unifieddatalibrary.Float(9.515),
		JmagOrigin:                  unifieddatalibrary.String("TP"),
		JmagUnc:                     unifieddatalibrary.Float(7.559),
		Kmag:                        unifieddatalibrary.Float(13.545),
		KmagOrigin:                  unifieddatalibrary.String("UC"),
		KmagUnc:                     unifieddatalibrary.Float(0.052),
		MorphologyInd:               unifieddatalibrary.Int(5),
		MultFlag:                    unifieddatalibrary.Bool(true),
		Multiplicity:                unifieddatalibrary.String("2"),
		NeighborDec:                 unifieddatalibrary.Float(89.99),
		NeighborDistance:            unifieddatalibrary.Float(201.406),
		NeighborFlag:                unifieddatalibrary.Bool(false),
		NeighborID:                  unifieddatalibrary.Int(2456),
		NeighborRa:                  unifieddatalibrary.Float(359.99),
		NonSingleStar:               unifieddatalibrary.String("7"),
		NumNeighbors:                unifieddatalibrary.Int(519),
		Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PanStarrsID:                 unifieddatalibrary.Int(215993386231483000),
		Parallax:                    unifieddatalibrary.Float(-6.8),
		ParallaxUnc:                 unifieddatalibrary.Float(82.35),
		Pmdec:                       unifieddatalibrary.Float(-970.1003),
		PmdecUnc:                    unifieddatalibrary.Float(1.22),
		Pmra:                        unifieddatalibrary.Float(1000.45),
		PmraUnc:                     unifieddatalibrary.Float(5.6),
		PmUncFlag:                   unifieddatalibrary.Bool(false),
		PosUncFlag:                  unifieddatalibrary.Bool(false),
		Ps1astrometryCorrectionFlag: unifieddatalibrary.Int(7),
		Ps1ObjInfoFlag:              unifieddatalibrary.Int(2005196800),
		Ps1QualityFlag:              unifieddatalibrary.Int(239),
		RaUnc:                       unifieddatalibrary.Float(509.466),
		Rmag:                        unifieddatalibrary.Float(22.657284),
		RmagOrigin:                  unifieddatalibrary.String("GA"),
		RmagUnc:                     unifieddatalibrary.Float(0.053),
		Rpmag:                       unifieddatalibrary.Float(8.0047),
		RpmagUnc:                    unifieddatalibrary.Float(1.233),
		Ruwe:                        unifieddatalibrary.Float(116.016365),
		SdaCatID:                    unifieddatalibrary.Int(3015023687),
		Sgmag:                       unifieddatalibrary.Float(28.663515),
		SgmagUnc:                    unifieddatalibrary.Float(2.3097522),
		Shift:                       unifieddatalibrary.Float(4.548),
		ShiftFlag:                   unifieddatalibrary.Bool(false),
		ShiftFwhm1:                  unifieddatalibrary.Float(0.157),
		ShiftFwhm6:                  unifieddatalibrary.Float(1.065),
		SkyMapperID:                 unifieddatalibrary.Int(505176683),
		TwoMassID:                   unifieddatalibrary.String("A1B2C3D4E5"),
		TwoMassPhQualInd:            unifieddatalibrary.String("AAAA"),
		TwoMassReadFlag:             unifieddatalibrary.String("111"),
		TwoMassXscID:                unifieddatalibrary.String("5000540"),
		TychoDscID:                  unifieddatalibrary.Int(9537000661),
		UhsID:                       unifieddatalibrary.Int(460074663768),
		UkidssGcsID:                 unifieddatalibrary.Int(442466709194),
		UkidssGpsID:                 unifieddatalibrary.Int(439491265503),
		UkidssLasID:                 unifieddatalibrary.Int(433883403451),
		VarFlag:                     unifieddatalibrary.Bool(true),
		Variability:                 unifieddatalibrary.String("1"),
		VhsID:                       unifieddatalibrary.Int(473820608583),
		Vmag:                        unifieddatalibrary.Float(25.829414),
		VmagOrigin:                  unifieddatalibrary.String("CR"),
		VmagUnc:                     unifieddatalibrary.Float(2.055),
		W1mag:                       unifieddatalibrary.Float(15.782),
		W1magOrigin:                 unifieddatalibrary.String("CA"),
		W1magUnc:                    unifieddatalibrary.Float(0.042),
		W1sat:                       unifieddatalibrary.Float(0.993),
		W2mag:                       unifieddatalibrary.Float(16.523),
		W2magOrigin:                 unifieddatalibrary.String("CA"),
		W2magUnc:                    unifieddatalibrary.Float(0.021),
		W2sat:                       unifieddatalibrary.Float(0.962),
		W3mag:                       unifieddatalibrary.Float(11.541),
		W3magOrigin:                 unifieddatalibrary.String("AL"),
		W3magUnc:                    unifieddatalibrary.Float(0.159),
		W3sat:                       unifieddatalibrary.Float(0.999),
		W4mag:                       unifieddatalibrary.Float(9.007),
		W4magOrigin:                 unifieddatalibrary.String("AL"),
		W4magUnc:                    unifieddatalibrary.Float(0.468),
		W4sat:                       unifieddatalibrary.Float(0.523),
		WdsCatID:                    unifieddatalibrary.String("155506"),
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
			AstrometryOrigin:            "GA",
			ClassificationMarking:       "U",
			CsID:                        12345,
			DataMode:                    unifieddatalibrary.StarCatalogUpdateParamsDataModeTest,
			Dec:                         21.8,
			Ra:                          14.43,
			Source:                      "Bluestaq",
			StarEpoch:                   2018.864,
			AavsoVsxID:                  unifieddatalibrary.Int(2387137),
			Abgmag:                      unifieddatalibrary.Float(11.24),
			AbgmagOrigin:                unifieddatalibrary.String("PS"),
			AbgmagUnc:                   unifieddatalibrary.Float(0.023),
			Abimag:                      unifieddatalibrary.Float(11.24),
			AbimagOrigin:                unifieddatalibrary.String("GA"),
			AbimagUnc:                   unifieddatalibrary.Float(0.023),
			Abrmag:                      unifieddatalibrary.Float(11.24),
			AbrmagOrigin:                unifieddatalibrary.String("SK"),
			AbrmagUnc:                   unifieddatalibrary.Float(0.023),
			Abymag:                      unifieddatalibrary.Float(11.24),
			AbymagOrigin:                unifieddatalibrary.String("AP"),
			AbymagUnc:                   unifieddatalibrary.Float(0.023),
			Abzmag:                      unifieddatalibrary.Float(11.24),
			AbzmagOrigin:                unifieddatalibrary.String("AP"),
			AbzmagUnc:                   unifieddatalibrary.Float(0.023),
			AllWisEccInd:                unifieddatalibrary.String("0000"),
			AllWiseID:                   unifieddatalibrary.String("WISEA J152743.04+624823.6"),
			AllWisEnaInd:                unifieddatalibrary.Int(1),
			AllWisEphQualInd:            unifieddatalibrary.String("AAAA"),
			ApassID:                     unifieddatalibrary.String("175-0082419"),
			AstrometricExcessNoise:      unifieddatalibrary.Float(921.2857),
			AstrometricExcessNoiseSig:   unifieddatalibrary.Float(194326670.1),
			Bmag:                        unifieddatalibrary.Float(27.004),
			BmagOrigin:                  unifieddatalibrary.String("AP"),
			BmagUnc:                     unifieddatalibrary.Float(9.999),
			Bpmag:                       unifieddatalibrary.Float(0.04559),
			BpmagUnc:                    unifieddatalibrary.Float(0.2227),
			CarrascoCatID:               unifieddatalibrary.Int(1244),
			CatVersion:                  unifieddatalibrary.String("1.23 DR3"),
			CatWise2020ID:               unifieddatalibrary.String("3584p196_b0-084425"),
			DecUnc:                      unifieddatalibrary.Float(40.996),
			DucatiCatID:                 unifieddatalibrary.String("AB ORI"),
			Gaiadr3CatID:                unifieddatalibrary.Int(89012345678901),
			Gmag:                        unifieddatalibrary.Float(0.0046),
			GmagUnc:                     unifieddatalibrary.Float(0.00292),
			GncCatID:                    unifieddatalibrary.Int(12345),
			HealpixIndex:                unifieddatalibrary.Int(196607),
			HipCatID:                    unifieddatalibrary.Int(12345),
			Hmag:                        unifieddatalibrary.Float(12.126),
			HmagOrigin:                  unifieddatalibrary.String("UL"),
			HmagUnc:                     unifieddatalibrary.Float(5.722),
			Imag:                        unifieddatalibrary.Float(22.46249),
			ImagOrigin:                  unifieddatalibrary.String("HI"),
			ImagUnc:                     unifieddatalibrary.Float(1.2000417),
			Jmag:                        unifieddatalibrary.Float(9.515),
			JmagOrigin:                  unifieddatalibrary.String("TP"),
			JmagUnc:                     unifieddatalibrary.Float(7.559),
			Kmag:                        unifieddatalibrary.Float(13.545),
			KmagOrigin:                  unifieddatalibrary.String("UC"),
			KmagUnc:                     unifieddatalibrary.Float(0.052),
			MorphologyInd:               unifieddatalibrary.Int(5),
			MultFlag:                    unifieddatalibrary.Bool(true),
			Multiplicity:                unifieddatalibrary.String("2"),
			NeighborDec:                 unifieddatalibrary.Float(89.99),
			NeighborDistance:            unifieddatalibrary.Float(201.406),
			NeighborFlag:                unifieddatalibrary.Bool(false),
			NeighborID:                  unifieddatalibrary.Int(2456),
			NeighborRa:                  unifieddatalibrary.Float(359.99),
			NonSingleStar:               unifieddatalibrary.String("7"),
			NumNeighbors:                unifieddatalibrary.Int(519),
			Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PanStarrsID:                 unifieddatalibrary.Int(215993386231483000),
			Parallax:                    unifieddatalibrary.Float(-6.8),
			ParallaxUnc:                 unifieddatalibrary.Float(82.35),
			Pmdec:                       unifieddatalibrary.Float(-970.1003),
			PmdecUnc:                    unifieddatalibrary.Float(1.22),
			Pmra:                        unifieddatalibrary.Float(1000.45),
			PmraUnc:                     unifieddatalibrary.Float(5.6),
			PmUncFlag:                   unifieddatalibrary.Bool(false),
			PosUncFlag:                  unifieddatalibrary.Bool(false),
			Ps1astrometryCorrectionFlag: unifieddatalibrary.Int(7),
			Ps1ObjInfoFlag:              unifieddatalibrary.Int(2005196800),
			Ps1QualityFlag:              unifieddatalibrary.Int(239),
			RaUnc:                       unifieddatalibrary.Float(509.466),
			Rmag:                        unifieddatalibrary.Float(22.657284),
			RmagOrigin:                  unifieddatalibrary.String("GA"),
			RmagUnc:                     unifieddatalibrary.Float(0.053),
			Rpmag:                       unifieddatalibrary.Float(8.0047),
			RpmagUnc:                    unifieddatalibrary.Float(1.233),
			Ruwe:                        unifieddatalibrary.Float(116.016365),
			SdaCatID:                    unifieddatalibrary.Int(3015023687),
			Sgmag:                       unifieddatalibrary.Float(28.663515),
			SgmagUnc:                    unifieddatalibrary.Float(2.3097522),
			Shift:                       unifieddatalibrary.Float(4.548),
			ShiftFlag:                   unifieddatalibrary.Bool(false),
			ShiftFwhm1:                  unifieddatalibrary.Float(0.157),
			ShiftFwhm6:                  unifieddatalibrary.Float(1.065),
			SkyMapperID:                 unifieddatalibrary.Int(505176683),
			TwoMassID:                   unifieddatalibrary.String("A1B2C3D4E5"),
			TwoMassPhQualInd:            unifieddatalibrary.String("AAAA"),
			TwoMassReadFlag:             unifieddatalibrary.String("111"),
			TwoMassXscID:                unifieddatalibrary.String("5000540"),
			TychoDscID:                  unifieddatalibrary.Int(9537000661),
			UhsID:                       unifieddatalibrary.Int(460074663768),
			UkidssGcsID:                 unifieddatalibrary.Int(442466709194),
			UkidssGpsID:                 unifieddatalibrary.Int(439491265503),
			UkidssLasID:                 unifieddatalibrary.Int(433883403451),
			VarFlag:                     unifieddatalibrary.Bool(true),
			Variability:                 unifieddatalibrary.String("1"),
			VhsID:                       unifieddatalibrary.Int(473820608583),
			Vmag:                        unifieddatalibrary.Float(25.829414),
			VmagOrigin:                  unifieddatalibrary.String("CR"),
			VmagUnc:                     unifieddatalibrary.Float(2.055),
			W1mag:                       unifieddatalibrary.Float(15.782),
			W1magOrigin:                 unifieddatalibrary.String("CA"),
			W1magUnc:                    unifieddatalibrary.Float(0.042),
			W1sat:                       unifieddatalibrary.Float(0.993),
			W2mag:                       unifieddatalibrary.Float(16.523),
			W2magOrigin:                 unifieddatalibrary.String("CA"),
			W2magUnc:                    unifieddatalibrary.Float(0.021),
			W2sat:                       unifieddatalibrary.Float(0.962),
			W3mag:                       unifieddatalibrary.Float(11.541),
			W3magOrigin:                 unifieddatalibrary.String("AL"),
			W3magUnc:                    unifieddatalibrary.Float(0.159),
			W3sat:                       unifieddatalibrary.Float(0.999),
			W4mag:                       unifieddatalibrary.Float(9.007),
			W4magOrigin:                 unifieddatalibrary.String("AL"),
			W4magUnc:                    unifieddatalibrary.Float(0.468),
			W4sat:                       unifieddatalibrary.Float(0.523),
			WdsCatID:                    unifieddatalibrary.String("155506"),
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
			AstrometryOrigin:            "GA",
			ClassificationMarking:       "U",
			CsID:                        12345,
			DataMode:                    "TEST",
			Dec:                         21.8,
			Ra:                          14.43,
			Source:                      "Bluestaq",
			StarEpoch:                   2018.864,
			AavsoVsxID:                  unifieddatalibrary.Int(2387137),
			Abgmag:                      unifieddatalibrary.Float(11.24),
			AbgmagOrigin:                unifieddatalibrary.String("PS"),
			AbgmagUnc:                   unifieddatalibrary.Float(0.023),
			Abimag:                      unifieddatalibrary.Float(11.24),
			AbimagOrigin:                unifieddatalibrary.String("GA"),
			AbimagUnc:                   unifieddatalibrary.Float(0.023),
			Abrmag:                      unifieddatalibrary.Float(11.24),
			AbrmagOrigin:                unifieddatalibrary.String("SK"),
			AbrmagUnc:                   unifieddatalibrary.Float(0.023),
			Abymag:                      unifieddatalibrary.Float(11.24),
			AbymagOrigin:                unifieddatalibrary.String("AP"),
			AbymagUnc:                   unifieddatalibrary.Float(0.023),
			Abzmag:                      unifieddatalibrary.Float(11.24),
			AbzmagOrigin:                unifieddatalibrary.String("AP"),
			AbzmagUnc:                   unifieddatalibrary.Float(0.023),
			AllWisEccInd:                unifieddatalibrary.String("0000"),
			AllWiseID:                   unifieddatalibrary.String("WISEA J152743.04+624823.6"),
			AllWisEnaInd:                unifieddatalibrary.Int(1),
			AllWisEphQualInd:            unifieddatalibrary.String("AAAA"),
			ApassID:                     unifieddatalibrary.String("175-0082419"),
			AstrometricExcessNoise:      unifieddatalibrary.Float(921.2857),
			AstrometricExcessNoiseSig:   unifieddatalibrary.Float(194326670.1),
			Bmag:                        unifieddatalibrary.Float(27.004),
			BmagOrigin:                  unifieddatalibrary.String("AP"),
			BmagUnc:                     unifieddatalibrary.Float(9.999),
			Bpmag:                       unifieddatalibrary.Float(0.04559),
			BpmagUnc:                    unifieddatalibrary.Float(0.2227),
			CarrascoCatID:               unifieddatalibrary.Int(1244),
			CatVersion:                  unifieddatalibrary.String("1.23 DR3"),
			CatWise2020ID:               unifieddatalibrary.String("3584p196_b0-084425"),
			DecUnc:                      unifieddatalibrary.Float(40.996),
			DucatiCatID:                 unifieddatalibrary.String("AB ORI"),
			Gaiadr3CatID:                unifieddatalibrary.Int(89012345678901),
			Gmag:                        unifieddatalibrary.Float(0.0046),
			GmagUnc:                     unifieddatalibrary.Float(0.00292),
			GncCatID:                    unifieddatalibrary.Int(12345),
			HealpixIndex:                unifieddatalibrary.Int(196607),
			HipCatID:                    unifieddatalibrary.Int(12345),
			Hmag:                        unifieddatalibrary.Float(12.126),
			HmagOrigin:                  unifieddatalibrary.String("UL"),
			HmagUnc:                     unifieddatalibrary.Float(5.722),
			Imag:                        unifieddatalibrary.Float(22.46249),
			ImagOrigin:                  unifieddatalibrary.String("HI"),
			ImagUnc:                     unifieddatalibrary.Float(1.2000417),
			Jmag:                        unifieddatalibrary.Float(9.515),
			JmagOrigin:                  unifieddatalibrary.String("TP"),
			JmagUnc:                     unifieddatalibrary.Float(7.559),
			Kmag:                        unifieddatalibrary.Float(13.545),
			KmagOrigin:                  unifieddatalibrary.String("UC"),
			KmagUnc:                     unifieddatalibrary.Float(0.052),
			MorphologyInd:               unifieddatalibrary.Int(5),
			MultFlag:                    unifieddatalibrary.Bool(true),
			Multiplicity:                unifieddatalibrary.String("2"),
			NeighborDec:                 unifieddatalibrary.Float(89.99),
			NeighborDistance:            unifieddatalibrary.Float(201.406),
			NeighborFlag:                unifieddatalibrary.Bool(false),
			NeighborID:                  unifieddatalibrary.Int(2456),
			NeighborRa:                  unifieddatalibrary.Float(359.99),
			NonSingleStar:               unifieddatalibrary.String("7"),
			NumNeighbors:                unifieddatalibrary.Int(519),
			Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PanStarrsID:                 unifieddatalibrary.Int(215993386231483000),
			Parallax:                    unifieddatalibrary.Float(-6.8),
			ParallaxUnc:                 unifieddatalibrary.Float(82.35),
			Pmdec:                       unifieddatalibrary.Float(-970.1003),
			PmdecUnc:                    unifieddatalibrary.Float(1.22),
			Pmra:                        unifieddatalibrary.Float(1000.45),
			PmraUnc:                     unifieddatalibrary.Float(5.6),
			PmUncFlag:                   unifieddatalibrary.Bool(false),
			PosUncFlag:                  unifieddatalibrary.Bool(false),
			Ps1astrometryCorrectionFlag: unifieddatalibrary.Int(7),
			Ps1ObjInfoFlag:              unifieddatalibrary.Int(2005196800),
			Ps1QualityFlag:              unifieddatalibrary.Int(239),
			RaUnc:                       unifieddatalibrary.Float(509.466),
			Rmag:                        unifieddatalibrary.Float(22.657284),
			RmagOrigin:                  unifieddatalibrary.String("GA"),
			RmagUnc:                     unifieddatalibrary.Float(0.053),
			Rpmag:                       unifieddatalibrary.Float(8.0047),
			RpmagUnc:                    unifieddatalibrary.Float(1.233),
			Ruwe:                        unifieddatalibrary.Float(116.016365),
			SdaCatID:                    unifieddatalibrary.Int(3015023687),
			Sgmag:                       unifieddatalibrary.Float(28.663515),
			SgmagUnc:                    unifieddatalibrary.Float(2.3097522),
			Shift:                       unifieddatalibrary.Float(4.548),
			ShiftFlag:                   unifieddatalibrary.Bool(false),
			ShiftFwhm1:                  unifieddatalibrary.Float(0.157),
			ShiftFwhm6:                  unifieddatalibrary.Float(1.065),
			SkyMapperID:                 unifieddatalibrary.Int(505176683),
			TwoMassID:                   unifieddatalibrary.String("A1B2C3D4E5"),
			TwoMassPhQualInd:            unifieddatalibrary.String("AAAA"),
			TwoMassReadFlag:             unifieddatalibrary.String("111"),
			TwoMassXscID:                unifieddatalibrary.String("5000540"),
			TychoDscID:                  unifieddatalibrary.Int(9537000661),
			UhsID:                       unifieddatalibrary.Int(460074663768),
			UkidssGcsID:                 unifieddatalibrary.Int(442466709194),
			UkidssGpsID:                 unifieddatalibrary.Int(439491265503),
			UkidssLasID:                 unifieddatalibrary.Int(433883403451),
			VarFlag:                     unifieddatalibrary.Bool(true),
			Variability:                 unifieddatalibrary.String("1"),
			VhsID:                       unifieddatalibrary.Int(473820608583),
			Vmag:                        unifieddatalibrary.Float(25.829414),
			VmagOrigin:                  unifieddatalibrary.String("CR"),
			VmagUnc:                     unifieddatalibrary.Float(2.055),
			W1mag:                       unifieddatalibrary.Float(15.782),
			W1magOrigin:                 unifieddatalibrary.String("CA"),
			W1magUnc:                    unifieddatalibrary.Float(0.042),
			W1sat:                       unifieddatalibrary.Float(0.993),
			W2mag:                       unifieddatalibrary.Float(16.523),
			W2magOrigin:                 unifieddatalibrary.String("CA"),
			W2magUnc:                    unifieddatalibrary.Float(0.021),
			W2sat:                       unifieddatalibrary.Float(0.962),
			W3mag:                       unifieddatalibrary.Float(11.541),
			W3magOrigin:                 unifieddatalibrary.String("AL"),
			W3magUnc:                    unifieddatalibrary.Float(0.159),
			W3sat:                       unifieddatalibrary.Float(0.999),
			W4mag:                       unifieddatalibrary.Float(9.007),
			W4magOrigin:                 unifieddatalibrary.String("AL"),
			W4magUnc:                    unifieddatalibrary.Float(0.468),
			W4sat:                       unifieddatalibrary.Float(0.523),
			WdsCatID:                    unifieddatalibrary.String("155506"),
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
			AstrometryOrigin:            "GA",
			ClassificationMarking:       "U",
			CsID:                        12345,
			DataMode:                    "TEST",
			Dec:                         21.8,
			Ra:                          14.43,
			Source:                      "Bluestaq",
			StarEpoch:                   2018.864,
			AavsoVsxID:                  unifieddatalibrary.Int(2387137),
			Abgmag:                      unifieddatalibrary.Float(11.24),
			AbgmagOrigin:                unifieddatalibrary.String("PS"),
			AbgmagUnc:                   unifieddatalibrary.Float(0.023),
			Abimag:                      unifieddatalibrary.Float(11.24),
			AbimagOrigin:                unifieddatalibrary.String("GA"),
			AbimagUnc:                   unifieddatalibrary.Float(0.023),
			Abrmag:                      unifieddatalibrary.Float(11.24),
			AbrmagOrigin:                unifieddatalibrary.String("SK"),
			AbrmagUnc:                   unifieddatalibrary.Float(0.023),
			Abymag:                      unifieddatalibrary.Float(11.24),
			AbymagOrigin:                unifieddatalibrary.String("AP"),
			AbymagUnc:                   unifieddatalibrary.Float(0.023),
			Abzmag:                      unifieddatalibrary.Float(11.24),
			AbzmagOrigin:                unifieddatalibrary.String("AP"),
			AbzmagUnc:                   unifieddatalibrary.Float(0.023),
			AllWisEccInd:                unifieddatalibrary.String("0000"),
			AllWiseID:                   unifieddatalibrary.String("WISEA J152743.04+624823.6"),
			AllWisEnaInd:                unifieddatalibrary.Int(1),
			AllWisEphQualInd:            unifieddatalibrary.String("AAAA"),
			ApassID:                     unifieddatalibrary.String("175-0082419"),
			AstrometricExcessNoise:      unifieddatalibrary.Float(921.2857),
			AstrometricExcessNoiseSig:   unifieddatalibrary.Float(194326670.1),
			Bmag:                        unifieddatalibrary.Float(27.004),
			BmagOrigin:                  unifieddatalibrary.String("AP"),
			BmagUnc:                     unifieddatalibrary.Float(9.999),
			Bpmag:                       unifieddatalibrary.Float(0.04559),
			BpmagUnc:                    unifieddatalibrary.Float(0.2227),
			CarrascoCatID:               unifieddatalibrary.Int(1244),
			CatVersion:                  unifieddatalibrary.String("1.23 DR3"),
			CatWise2020ID:               unifieddatalibrary.String("3584p196_b0-084425"),
			DecUnc:                      unifieddatalibrary.Float(40.996),
			DucatiCatID:                 unifieddatalibrary.String("AB ORI"),
			Gaiadr3CatID:                unifieddatalibrary.Int(89012345678901),
			Gmag:                        unifieddatalibrary.Float(0.0046),
			GmagUnc:                     unifieddatalibrary.Float(0.00292),
			GncCatID:                    unifieddatalibrary.Int(12345),
			HealpixIndex:                unifieddatalibrary.Int(196607),
			HipCatID:                    unifieddatalibrary.Int(12345),
			Hmag:                        unifieddatalibrary.Float(12.126),
			HmagOrigin:                  unifieddatalibrary.String("UL"),
			HmagUnc:                     unifieddatalibrary.Float(5.722),
			Imag:                        unifieddatalibrary.Float(22.46249),
			ImagOrigin:                  unifieddatalibrary.String("HI"),
			ImagUnc:                     unifieddatalibrary.Float(1.2000417),
			Jmag:                        unifieddatalibrary.Float(9.515),
			JmagOrigin:                  unifieddatalibrary.String("TP"),
			JmagUnc:                     unifieddatalibrary.Float(7.559),
			Kmag:                        unifieddatalibrary.Float(13.545),
			KmagOrigin:                  unifieddatalibrary.String("UC"),
			KmagUnc:                     unifieddatalibrary.Float(0.052),
			MorphologyInd:               unifieddatalibrary.Int(5),
			MultFlag:                    unifieddatalibrary.Bool(true),
			Multiplicity:                unifieddatalibrary.String("2"),
			NeighborDec:                 unifieddatalibrary.Float(89.99),
			NeighborDistance:            unifieddatalibrary.Float(201.406),
			NeighborFlag:                unifieddatalibrary.Bool(false),
			NeighborID:                  unifieddatalibrary.Int(2456),
			NeighborRa:                  unifieddatalibrary.Float(359.99),
			NonSingleStar:               unifieddatalibrary.String("7"),
			NumNeighbors:                unifieddatalibrary.Int(519),
			Origin:                      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PanStarrsID:                 unifieddatalibrary.Int(215993386231483000),
			Parallax:                    unifieddatalibrary.Float(-6.8),
			ParallaxUnc:                 unifieddatalibrary.Float(82.35),
			Pmdec:                       unifieddatalibrary.Float(-970.1003),
			PmdecUnc:                    unifieddatalibrary.Float(1.22),
			Pmra:                        unifieddatalibrary.Float(1000.45),
			PmraUnc:                     unifieddatalibrary.Float(5.6),
			PmUncFlag:                   unifieddatalibrary.Bool(false),
			PosUncFlag:                  unifieddatalibrary.Bool(false),
			Ps1astrometryCorrectionFlag: unifieddatalibrary.Int(7),
			Ps1ObjInfoFlag:              unifieddatalibrary.Int(2005196800),
			Ps1QualityFlag:              unifieddatalibrary.Int(239),
			RaUnc:                       unifieddatalibrary.Float(509.466),
			Rmag:                        unifieddatalibrary.Float(22.657284),
			RmagOrigin:                  unifieddatalibrary.String("GA"),
			RmagUnc:                     unifieddatalibrary.Float(0.053),
			Rpmag:                       unifieddatalibrary.Float(8.0047),
			RpmagUnc:                    unifieddatalibrary.Float(1.233),
			Ruwe:                        unifieddatalibrary.Float(116.016365),
			SdaCatID:                    unifieddatalibrary.Int(3015023687),
			Sgmag:                       unifieddatalibrary.Float(28.663515),
			SgmagUnc:                    unifieddatalibrary.Float(2.3097522),
			Shift:                       unifieddatalibrary.Float(4.548),
			ShiftFlag:                   unifieddatalibrary.Bool(false),
			ShiftFwhm1:                  unifieddatalibrary.Float(0.157),
			ShiftFwhm6:                  unifieddatalibrary.Float(1.065),
			SkyMapperID:                 unifieddatalibrary.Int(505176683),
			TwoMassID:                   unifieddatalibrary.String("A1B2C3D4E5"),
			TwoMassPhQualInd:            unifieddatalibrary.String("AAAA"),
			TwoMassReadFlag:             unifieddatalibrary.String("111"),
			TwoMassXscID:                unifieddatalibrary.String("5000540"),
			TychoDscID:                  unifieddatalibrary.Int(9537000661),
			UhsID:                       unifieddatalibrary.Int(460074663768),
			UkidssGcsID:                 unifieddatalibrary.Int(442466709194),
			UkidssGpsID:                 unifieddatalibrary.Int(439491265503),
			UkidssLasID:                 unifieddatalibrary.Int(433883403451),
			VarFlag:                     unifieddatalibrary.Bool(true),
			Variability:                 unifieddatalibrary.String("1"),
			VhsID:                       unifieddatalibrary.Int(473820608583),
			Vmag:                        unifieddatalibrary.Float(25.829414),
			VmagOrigin:                  unifieddatalibrary.String("CR"),
			VmagUnc:                     unifieddatalibrary.Float(2.055),
			W1mag:                       unifieddatalibrary.Float(15.782),
			W1magOrigin:                 unifieddatalibrary.String("CA"),
			W1magUnc:                    unifieddatalibrary.Float(0.042),
			W1sat:                       unifieddatalibrary.Float(0.993),
			W2mag:                       unifieddatalibrary.Float(16.523),
			W2magOrigin:                 unifieddatalibrary.String("CA"),
			W2magUnc:                    unifieddatalibrary.Float(0.021),
			W2sat:                       unifieddatalibrary.Float(0.962),
			W3mag:                       unifieddatalibrary.Float(11.541),
			W3magOrigin:                 unifieddatalibrary.String("AL"),
			W3magUnc:                    unifieddatalibrary.Float(0.159),
			W3sat:                       unifieddatalibrary.Float(0.999),
			W4mag:                       unifieddatalibrary.Float(9.007),
			W4magOrigin:                 unifieddatalibrary.String("AL"),
			W4magUnc:                    unifieddatalibrary.Float(0.468),
			W4sat:                       unifieddatalibrary.Float(0.523),
			WdsCatID:                    unifieddatalibrary.String("155506"),
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
