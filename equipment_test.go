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

func TestEquipmentNewWithOptionalParams(t *testing.T) {
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
	err := client.Equipment.New(context.TODO(), unifieddatalibrary.EquipmentNewParams{
		ClassificationMarking: "U",
		CountryCode:           "IQ",
		DataMode:              unifieddatalibrary.EquipmentNewParamsDataModeTest,
		Lat:                   39.019242,
		Lon:                   -104.251659,
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
		AirDefArea:            unifieddatalibrary.String("AL006"),
		Allegiance:            unifieddatalibrary.String("OTHR"),
		AltAllegiance:         unifieddatalibrary.String("HL"),
		AltCountryCode:        unifieddatalibrary.String("IZ"),
		AltEqpID:              unifieddatalibrary.String("ORIG-EQP-ID"),
		ClassRating:           unifieddatalibrary.String("1"),
		Condition:             unifieddatalibrary.String("RDY"),
		ConditionAvail:        unifieddatalibrary.String("A"),
		Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
		CoordDatum:            unifieddatalibrary.String("WGS"),
		CoordDerivAcc:         unifieddatalibrary.Float(12.345),
		ElevMsl:               unifieddatalibrary.Float(123.45),
		ElevMslConfLvl:        unifieddatalibrary.Int(50),
		ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
		EqpCode:               unifieddatalibrary.String("X12345"),
		EqpIDNum:              unifieddatalibrary.String("001"),
		Eval:                  unifieddatalibrary.Int(7),
		Fpa:                   unifieddatalibrary.String("NOB"),
		Function:              unifieddatalibrary.String("OCC"),
		FunctPrimary:          unifieddatalibrary.String("JG"),
		GeoidalMslSep:         unifieddatalibrary.Float(12.34),
		Ident:                 unifieddatalibrary.String("FRIEND"),
		IDOperatingUnit:       unifieddatalibrary.String("UNIT-ID"),
		IDParentEquipment:     unifieddatalibrary.String("PARENT-EQUIPMENT-ID"),
		IDSite:                unifieddatalibrary.String("SITE-ID"),
		LocReason:             unifieddatalibrary.String("GR"),
		MilGrid:               unifieddatalibrary.String("4QFJ12345678"),
		MilGridSys:            unifieddatalibrary.String("UTM"),
		Nomen:                 unifieddatalibrary.String("AMPHIBIOUS WARFARE SHIP"),
		OperAreaPrimary:       unifieddatalibrary.String("Territorial Sea"),
		OperStatus:            unifieddatalibrary.String("OPR"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PolSubdiv:             unifieddatalibrary.String("IZ07"),
		QtyOh:                 unifieddatalibrary.Int(7),
		RecStatus:             unifieddatalibrary.String("A"),
		ReferenceDoc:          unifieddatalibrary.String("Provider Reference Documentation"),
		ResProd:               unifieddatalibrary.String("RT"),
		ReviewDate:            unifieddatalibrary.Time(time.Now()),
		SeqNum:                unifieddatalibrary.Int(5),
		SrcIDs:                []string{"SRC_ID_1"},
		SrcTyps:               []string{"AIRCRAFT"},
		SymCode:               unifieddatalibrary.String("SOGPU----------"),
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

func TestEquipmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Equipment.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EquipmentGetParams{
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

func TestEquipmentUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Equipment.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.EquipmentUpdateParams{
			ClassificationMarking: "U",
			CountryCode:           "IQ",
			DataMode:              unifieddatalibrary.EquipmentUpdateParamsDataModeTest,
			Lat:                   39.019242,
			Lon:                   -104.251659,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AirDefArea:            unifieddatalibrary.String("AL006"),
			Allegiance:            unifieddatalibrary.String("OTHR"),
			AltAllegiance:         unifieddatalibrary.String("HL"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			AltEqpID:              unifieddatalibrary.String("ORIG-EQP-ID"),
			ClassRating:           unifieddatalibrary.String("1"),
			Condition:             unifieddatalibrary.String("RDY"),
			ConditionAvail:        unifieddatalibrary.String("A"),
			Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
			CoordDatum:            unifieddatalibrary.String("WGS"),
			CoordDerivAcc:         unifieddatalibrary.Float(12.345),
			ElevMsl:               unifieddatalibrary.Float(123.45),
			ElevMslConfLvl:        unifieddatalibrary.Int(50),
			ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
			EqpCode:               unifieddatalibrary.String("X12345"),
			EqpIDNum:              unifieddatalibrary.String("001"),
			Eval:                  unifieddatalibrary.Int(7),
			Fpa:                   unifieddatalibrary.String("NOB"),
			Function:              unifieddatalibrary.String("OCC"),
			FunctPrimary:          unifieddatalibrary.String("JG"),
			GeoidalMslSep:         unifieddatalibrary.Float(12.34),
			Ident:                 unifieddatalibrary.String("FRIEND"),
			IDOperatingUnit:       unifieddatalibrary.String("UNIT-ID"),
			IDParentEquipment:     unifieddatalibrary.String("PARENT-EQUIPMENT-ID"),
			IDSite:                unifieddatalibrary.String("SITE-ID"),
			LocReason:             unifieddatalibrary.String("GR"),
			MilGrid:               unifieddatalibrary.String("4QFJ12345678"),
			MilGridSys:            unifieddatalibrary.String("UTM"),
			Nomen:                 unifieddatalibrary.String("AMPHIBIOUS WARFARE SHIP"),
			OperAreaPrimary:       unifieddatalibrary.String("Territorial Sea"),
			OperStatus:            unifieddatalibrary.String("OPR"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PolSubdiv:             unifieddatalibrary.String("IZ07"),
			QtyOh:                 unifieddatalibrary.Int(7),
			RecStatus:             unifieddatalibrary.String("A"),
			ReferenceDoc:          unifieddatalibrary.String("Provider Reference Documentation"),
			ResProd:               unifieddatalibrary.String("RT"),
			ReviewDate:            unifieddatalibrary.Time(time.Now()),
			SeqNum:                unifieddatalibrary.Int(5),
			SrcIDs:                []string{"SRC_ID_1"},
			SrcTyps:               []string{"AIRCRAFT"},
			SymCode:               unifieddatalibrary.String("SOGPU----------"),
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

func TestEquipmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Equipment.List(context.TODO(), unifieddatalibrary.EquipmentListParams{
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

func TestEquipmentDelete(t *testing.T) {
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
	err := client.Equipment.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEquipmentCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Equipment.Count(context.TODO(), unifieddatalibrary.EquipmentCountParams{
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

func TestEquipmentNewBulk(t *testing.T) {
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
	err := client.Equipment.NewBulk(context.TODO(), unifieddatalibrary.EquipmentNewBulkParams{
		Body: []unifieddatalibrary.EquipmentNewBulkParamsBody{{
			ClassificationMarking: "U",
			CountryCode:           "IQ",
			DataMode:              "TEST",
			Lat:                   39.019242,
			Lon:                   -104.251659,
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("0167f577-e06c-358e-85aa-0a07a730bdd0"),
			AirDefArea:            unifieddatalibrary.String("AL006"),
			Allegiance:            unifieddatalibrary.String("OTHR"),
			AltAllegiance:         unifieddatalibrary.String("HL"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			AltEqpID:              unifieddatalibrary.String("ORIG-EQP-ID"),
			ClassRating:           unifieddatalibrary.String("1"),
			Condition:             unifieddatalibrary.String("RDY"),
			ConditionAvail:        unifieddatalibrary.String("A"),
			Coord:                 unifieddatalibrary.String("340000000N0430000000E"),
			CoordDatum:            unifieddatalibrary.String("WGS"),
			CoordDerivAcc:         unifieddatalibrary.Float(12.345),
			ElevMsl:               unifieddatalibrary.Float(123.45),
			ElevMslConfLvl:        unifieddatalibrary.Int(50),
			ElevMslDerivAcc:       unifieddatalibrary.Float(12.34),
			EqpCode:               unifieddatalibrary.String("X12345"),
			EqpIDNum:              unifieddatalibrary.String("001"),
			Eval:                  unifieddatalibrary.Int(7),
			Fpa:                   unifieddatalibrary.String("NOB"),
			Function:              unifieddatalibrary.String("OCC"),
			FunctPrimary:          unifieddatalibrary.String("JG"),
			GeoidalMslSep:         unifieddatalibrary.Float(12.34),
			Ident:                 unifieddatalibrary.String("FRIEND"),
			IDOperatingUnit:       unifieddatalibrary.String("UNIT-ID"),
			IDParentEquipment:     unifieddatalibrary.String("PARENT-EQUIPMENT-ID"),
			IDSite:                unifieddatalibrary.String("SITE-ID"),
			LocReason:             unifieddatalibrary.String("GR"),
			MilGrid:               unifieddatalibrary.String("4QFJ12345678"),
			MilGridSys:            unifieddatalibrary.String("UTM"),
			Nomen:                 unifieddatalibrary.String("AMPHIBIOUS WARFARE SHIP"),
			OperAreaPrimary:       unifieddatalibrary.String("Territorial Sea"),
			OperStatus:            unifieddatalibrary.String("OPR"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PolSubdiv:             unifieddatalibrary.String("IZ07"),
			QtyOh:                 unifieddatalibrary.Int(7),
			RecStatus:             unifieddatalibrary.String("A"),
			ReferenceDoc:          unifieddatalibrary.String("Provider Reference Documentation"),
			ResProd:               unifieddatalibrary.String("RT"),
			ReviewDate:            unifieddatalibrary.Time(time.Now()),
			SeqNum:                unifieddatalibrary.Int(5),
			SrcIDs:                []string{"SRC_ID_1"},
			SrcTyps:               []string{"AIRCRAFT"},
			SymCode:               unifieddatalibrary.String("SOGPU----------"),
			Utm:                   unifieddatalibrary.String("19P4390691376966"),
			Wac:                   unifieddatalibrary.String("0427"),
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

func TestEquipmentQueryHelp(t *testing.T) {
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
	err := client.Equipment.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEquipmentTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Equipment.Tuple(context.TODO(), unifieddatalibrary.EquipmentTupleParams{
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
