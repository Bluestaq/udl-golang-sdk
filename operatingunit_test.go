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

func TestOperatingunitNewWithOptionalParams(t *testing.T) {
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
	err := client.Operatingunit.New(context.TODO(), unifieddatalibrary.OperatingunitNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OperatingunitNewParamsDataModeTest,
		Name:                  "SOME_NAME",
		Source:                "some.user",
		AirDefArea:            unifieddatalibrary.String("AL006"),
		Allegiance:            unifieddatalibrary.String("OTHR"),
		AltAllegiance:         unifieddatalibrary.String("HL"),
		AltCountryCode:        unifieddatalibrary.String("IZ"),
		AltOperatingUnitID:    unifieddatalibrary.String("32100000000021"),
		ClassRating:           unifieddatalibrary.String("1"),
		Condition:             unifieddatalibrary.String("RDY"),
		ConditionAvail:        unifieddatalibrary.String("A"),
		Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
		CoordDatum:            unifieddatalibrary.String("WGS"),
		CoordDerivAcc:         unifieddatalibrary.Float(12.345),
		CountryCode:           unifieddatalibrary.String("IQ"),
		DeployStatus:          unifieddatalibrary.String("ND"),
		Description:           unifieddatalibrary.String("Description of unit"),
		DivCat:                unifieddatalibrary.String("5"),
		Echelon:               unifieddatalibrary.String("SHIP"),
		EchelonTier:           unifieddatalibrary.String("68"),
		ElevMsl:               unifieddatalibrary.Float(123.45),
		ElevMslConfLvl:        unifieddatalibrary.Int(50),
		ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
		Eval:                  unifieddatalibrary.Int(7),
		FlagFlown:             unifieddatalibrary.String("IZ"),
		FleetID:               unifieddatalibrary.String("A"),
		Force:                 unifieddatalibrary.String("NV"),
		ForceName:             unifieddatalibrary.String("FORCE-NAME"),
		Fpa:                   unifieddatalibrary.String("EOB"),
		FunctRole:             unifieddatalibrary.String("MIL"),
		GeoidalMslSep:         unifieddatalibrary.Float(12.34),
		IDContact:             unifieddatalibrary.String("CONTACT-ID"),
		Ident:                 unifieddatalibrary.String("FRIEND"),
		IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
		IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
		IDOrganization:        unifieddatalibrary.String("ORGANIZATION-ID"),
		Lat:                   unifieddatalibrary.Float(45.23),
		LocName:               unifieddatalibrary.String("LOCATION_NAME"),
		LocReason:             unifieddatalibrary.String("GR"),
		Lon:                   unifieddatalibrary.Float(179.1),
		MasterUnit:            unifieddatalibrary.Bool(true),
		MilGrid:               unifieddatalibrary.String("4QFJ12345678"),
		MilGridSys:            unifieddatalibrary.String("UTM"),
		MsnPrimary:            unifieddatalibrary.String("W6"),
		MsnPrimarySpecialty:   unifieddatalibrary.String("QK"),
		OperStatus:            unifieddatalibrary.String("OPR"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PolSubdiv:             unifieddatalibrary.String("IZ07"),
		RecStatus:             unifieddatalibrary.String("A"),
		ReferenceDoc:          unifieddatalibrary.String("Provider Reference Documentation"),
		ResProd:               unifieddatalibrary.String("RT"),
		ReviewDate:            unifieddatalibrary.Time(time.Now()),
		StylizedUnit:          unifieddatalibrary.Bool(true),
		SymCode:               unifieddatalibrary.String("SOGPU----------"),
		UnitIdentifier:        unifieddatalibrary.String("AZXAZ12345"),
		Utm:                   unifieddatalibrary.String("19P4390691376966"),
		Wac:                   unifieddatalibrary.String("0427"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperatingunitUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Operatingunit.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OperatingunitUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OperatingunitUpdateParamsDataModeTest,
			Name:                  "SOME_NAME",
			Source:                "some.user",
			AirDefArea:            unifieddatalibrary.String("AL006"),
			Allegiance:            unifieddatalibrary.String("OTHR"),
			AltAllegiance:         unifieddatalibrary.String("HL"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			AltOperatingUnitID:    unifieddatalibrary.String("32100000000021"),
			ClassRating:           unifieddatalibrary.String("1"),
			Condition:             unifieddatalibrary.String("RDY"),
			ConditionAvail:        unifieddatalibrary.String("A"),
			Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
			CoordDatum:            unifieddatalibrary.String("WGS"),
			CoordDerivAcc:         unifieddatalibrary.Float(12.345),
			CountryCode:           unifieddatalibrary.String("IQ"),
			DeployStatus:          unifieddatalibrary.String("ND"),
			Description:           unifieddatalibrary.String("Description of unit"),
			DivCat:                unifieddatalibrary.String("5"),
			Echelon:               unifieddatalibrary.String("SHIP"),
			EchelonTier:           unifieddatalibrary.String("68"),
			ElevMsl:               unifieddatalibrary.Float(123.45),
			ElevMslConfLvl:        unifieddatalibrary.Int(50),
			ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
			Eval:                  unifieddatalibrary.Int(7),
			FlagFlown:             unifieddatalibrary.String("IZ"),
			FleetID:               unifieddatalibrary.String("A"),
			Force:                 unifieddatalibrary.String("NV"),
			ForceName:             unifieddatalibrary.String("FORCE-NAME"),
			Fpa:                   unifieddatalibrary.String("EOB"),
			FunctRole:             unifieddatalibrary.String("MIL"),
			GeoidalMslSep:         unifieddatalibrary.Float(12.34),
			IDContact:             unifieddatalibrary.String("CONTACT-ID"),
			Ident:                 unifieddatalibrary.String("FRIEND"),
			IDLocation:            unifieddatalibrary.String("LOCATION-ID"),
			IDOperatingUnit:       unifieddatalibrary.String("OPERATINGUNIT-ID"),
			IDOrganization:        unifieddatalibrary.String("ORGANIZATION-ID"),
			Lat:                   unifieddatalibrary.Float(45.23),
			LocName:               unifieddatalibrary.String("LOCATION_NAME"),
			LocReason:             unifieddatalibrary.String("GR"),
			Lon:                   unifieddatalibrary.Float(179.1),
			MasterUnit:            unifieddatalibrary.Bool(true),
			MilGrid:               unifieddatalibrary.String("4QFJ12345678"),
			MilGridSys:            unifieddatalibrary.String("UTM"),
			MsnPrimary:            unifieddatalibrary.String("W6"),
			MsnPrimarySpecialty:   unifieddatalibrary.String("QK"),
			OperStatus:            unifieddatalibrary.String("OPR"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PolSubdiv:             unifieddatalibrary.String("IZ07"),
			RecStatus:             unifieddatalibrary.String("A"),
			ReferenceDoc:          unifieddatalibrary.String("Provider Reference Documentation"),
			ResProd:               unifieddatalibrary.String("RT"),
			ReviewDate:            unifieddatalibrary.Time(time.Now()),
			StylizedUnit:          unifieddatalibrary.Bool(true),
			SymCode:               unifieddatalibrary.String("SOGPU----------"),
			UnitIdentifier:        unifieddatalibrary.String("AZXAZ12345"),
			Utm:                   unifieddatalibrary.String("19P4390691376966"),
			Wac:                   unifieddatalibrary.String("0427"),
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

func TestOperatingunitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operatingunit.List(context.TODO(), unifieddatalibrary.OperatingunitListParams{
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

func TestOperatingunitDelete(t *testing.T) {
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
	err := client.Operatingunit.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperatingunitCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Operatingunit.Count(context.TODO(), unifieddatalibrary.OperatingunitCountParams{
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

func TestOperatingunitGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operatingunit.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OperatingunitGetParams{
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

func TestOperatingunitQueryhelp(t *testing.T) {
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
	_, err := client.Operatingunit.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperatingunitTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Operatingunit.Tuple(context.TODO(), unifieddatalibrary.OperatingunitTupleParams{
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
