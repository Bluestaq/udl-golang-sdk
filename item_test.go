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

func TestItemNewWithOptionalParams(t *testing.T) {
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
	err := client.Item.New(context.TODO(), unifieddatalibrary.ItemNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ItemNewParamsDataModeTest,
		ScanCode:              "12345ABCD",
		Source:                "Bluestaq",
		Type:                  "CARGO",
		ID:                    unifieddatalibrary.String("22f1f6da-a568-655a-ea37-76d013d04853"),
		AccSysKeys:            []string{"System key1", "System key2"},
		AccSysNotes:           unifieddatalibrary.String("Accepting System Notes"),
		AccSystem:             unifieddatalibrary.String("Accepting System"),
		AccSysValues:          []string{"System value1", "System value2"},
		Airdrop:               unifieddatalibrary.Bool(true),
		AltDataFormat:         unifieddatalibrary.String("Alt Data Format"),
		CargoType:             unifieddatalibrary.String("PALLET"),
		CenterlineOffset:      unifieddatalibrary.Float(3.1),
		Cg:                    unifieddatalibrary.Float(112.014),
		CommodityCode:         unifieddatalibrary.String("2304116"),
		CommoditySys:          unifieddatalibrary.String("STCC"),
		Container:             unifieddatalibrary.Bool(true),
		Departure:             unifieddatalibrary.String("CHS"),
		Destination:           unifieddatalibrary.String("RMS"),
		DvCode:                unifieddatalibrary.String("DV-2"),
		Fs:                    unifieddatalibrary.Float(412.1),
		HazCodes:              []float64{1.1, 1.2},
		Height:                unifieddatalibrary.Float(1.1),
		IDAirLoadPlan:         unifieddatalibrary.String("1038c389-d38e-270f-51cc-6a12e905abe8"),
		ItemContains:          []string{"2UJ8843K", "745YV1T65"},
		Keys:                  []string{"key1", "key2"},
		LastArrDate:           unifieddatalibrary.Time(time.Now()),
		Length:                unifieddatalibrary.Float(1.1),
		Moment:                unifieddatalibrary.Float(4000.1),
		Name:                  unifieddatalibrary.String("Product Name"),
		NetExpWt:              unifieddatalibrary.Float(51.437),
		Notes:                 unifieddatalibrary.String("Example notes"),
		NumPalletPos:          unifieddatalibrary.Int(2),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ProductCode:           unifieddatalibrary.String("530500234"),
		ProductSys:            unifieddatalibrary.String("NSN"),
		ReceivingBranch:       unifieddatalibrary.String("Air Force"),
		ReceivingUnit:         unifieddatalibrary.String("50 SBN"),
		ScGenTool:             unifieddatalibrary.String("bID"),
		Tcn:                   unifieddatalibrary.String("M1358232245912XXX"),
		Uln:                   unifieddatalibrary.String("T01ME01"),
		Values:                []string{"value1", "value2"},
		Volume:                unifieddatalibrary.Float(7.8902),
		Weight:                unifieddatalibrary.Float(5443.335),
		WeightTs:              unifieddatalibrary.Time(time.Now()),
		Width:                 unifieddatalibrary.Float(1.1),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Item.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.ItemUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.ItemUpdateParamsDataModeTest,
			ScanCode:              "12345ABCD",
			Source:                "Bluestaq",
			Type:                  "CARGO",
			ID:                    unifieddatalibrary.String("22f1f6da-a568-655a-ea37-76d013d04853"),
			AccSysKeys:            []string{"System key1", "System key2"},
			AccSysNotes:           unifieddatalibrary.String("Accepting System Notes"),
			AccSystem:             unifieddatalibrary.String("Accepting System"),
			AccSysValues:          []string{"System value1", "System value2"},
			Airdrop:               unifieddatalibrary.Bool(true),
			AltDataFormat:         unifieddatalibrary.String("Alt Data Format"),
			CargoType:             unifieddatalibrary.String("PALLET"),
			CenterlineOffset:      unifieddatalibrary.Float(3.1),
			Cg:                    unifieddatalibrary.Float(112.014),
			CommodityCode:         unifieddatalibrary.String("2304116"),
			CommoditySys:          unifieddatalibrary.String("STCC"),
			Container:             unifieddatalibrary.Bool(true),
			Departure:             unifieddatalibrary.String("CHS"),
			Destination:           unifieddatalibrary.String("RMS"),
			DvCode:                unifieddatalibrary.String("DV-2"),
			Fs:                    unifieddatalibrary.Float(412.1),
			HazCodes:              []float64{1.1, 1.2},
			Height:                unifieddatalibrary.Float(1.1),
			IDAirLoadPlan:         unifieddatalibrary.String("1038c389-d38e-270f-51cc-6a12e905abe8"),
			ItemContains:          []string{"2UJ8843K", "745YV1T65"},
			Keys:                  []string{"key1", "key2"},
			LastArrDate:           unifieddatalibrary.Time(time.Now()),
			Length:                unifieddatalibrary.Float(1.1),
			Moment:                unifieddatalibrary.Float(4000.1),
			Name:                  unifieddatalibrary.String("Product Name"),
			NetExpWt:              unifieddatalibrary.Float(51.437),
			Notes:                 unifieddatalibrary.String("Example notes"),
			NumPalletPos:          unifieddatalibrary.Int(2),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ProductCode:           unifieddatalibrary.String("530500234"),
			ProductSys:            unifieddatalibrary.String("NSN"),
			ReceivingBranch:       unifieddatalibrary.String("Air Force"),
			ReceivingUnit:         unifieddatalibrary.String("50 SBN"),
			ScGenTool:             unifieddatalibrary.String("bID"),
			Tcn:                   unifieddatalibrary.String("M1358232245912XXX"),
			Uln:                   unifieddatalibrary.String("T01ME01"),
			Values:                []string{"value1", "value2"},
			Volume:                unifieddatalibrary.Float(7.8902),
			Weight:                unifieddatalibrary.Float(5443.335),
			WeightTs:              unifieddatalibrary.Time(time.Now()),
			Width:                 unifieddatalibrary.Float(1.1),
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

func TestItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Item.List(context.TODO(), unifieddatalibrary.ItemListParams{
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

func TestItemDelete(t *testing.T) {
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
	err := client.Item.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Item.Count(context.TODO(), unifieddatalibrary.ItemCountParams{
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

func TestItemGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Item.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ItemGetParams{
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

func TestItemQueryhelp(t *testing.T) {
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
	err := client.Item.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Item.Tuple(context.TODO(), unifieddatalibrary.ItemTupleParams{
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

func TestItemUnvalidatedPublish(t *testing.T) {
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
	err := client.Item.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ItemUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ItemUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ScanCode:              "12345ABCD",
			Source:                "Bluestaq",
			Type:                  "CARGO",
			ID:                    unifieddatalibrary.String("22f1f6da-a568-655a-ea37-76d013d04853"),
			AccSysKeys:            []string{"System key1", "System key2"},
			AccSysNotes:           unifieddatalibrary.String("Accepting System Notes"),
			AccSystem:             unifieddatalibrary.String("Accepting System"),
			AccSysValues:          []string{"System value1", "System value2"},
			Airdrop:               unifieddatalibrary.Bool(true),
			AltDataFormat:         unifieddatalibrary.String("Alt Data Format"),
			CargoType:             unifieddatalibrary.String("PALLET"),
			CenterlineOffset:      unifieddatalibrary.Float(3.1),
			Cg:                    unifieddatalibrary.Float(112.014),
			CommodityCode:         unifieddatalibrary.String("2304116"),
			CommoditySys:          unifieddatalibrary.String("STCC"),
			Container:             unifieddatalibrary.Bool(true),
			Departure:             unifieddatalibrary.String("CHS"),
			Destination:           unifieddatalibrary.String("RMS"),
			DvCode:                unifieddatalibrary.String("DV-2"),
			Fs:                    unifieddatalibrary.Float(412.1),
			HazCodes:              []float64{1.1, 1.2},
			Height:                unifieddatalibrary.Float(1.1),
			IDAirLoadPlan:         unifieddatalibrary.String("1038c389-d38e-270f-51cc-6a12e905abe8"),
			ItemContains:          []string{"2UJ8843K", "745YV1T65"},
			Keys:                  []string{"key1", "key2"},
			LastArrDate:           unifieddatalibrary.Time(time.Now()),
			Length:                unifieddatalibrary.Float(1.1),
			Moment:                unifieddatalibrary.Float(4000.1),
			Name:                  unifieddatalibrary.String("Product Name"),
			NetExpWt:              unifieddatalibrary.Float(51.437),
			Notes:                 unifieddatalibrary.String("Example notes"),
			NumPalletPos:          unifieddatalibrary.Int(2),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ProductCode:           unifieddatalibrary.String("530500234"),
			ProductSys:            unifieddatalibrary.String("NSN"),
			ReceivingBranch:       unifieddatalibrary.String("Air Force"),
			ReceivingUnit:         unifieddatalibrary.String("50 SBN"),
			ScGenTool:             unifieddatalibrary.String("bID"),
			Tcn:                   unifieddatalibrary.String("M1358232245912XXX"),
			Uln:                   unifieddatalibrary.String("T01ME01"),
			Values:                []string{"value1", "value2"},
			Volume:                unifieddatalibrary.Float(7.8902),
			Weight:                unifieddatalibrary.Float(5443.335),
			WeightTs:              unifieddatalibrary.Time(time.Now()),
			Width:                 unifieddatalibrary.Float(1.1),
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
