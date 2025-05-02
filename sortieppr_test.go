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

func TestSortiepprNewWithOptionalParams(t *testing.T) {
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
	err := client.Sortieppr.New(context.TODO(), unifieddatalibrary.SortiepprNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SortiepprNewParamsDataModeTest,
		IDSortie:              "4ef3d1e8-ab08-ab70-498f-edc479734e5c",
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("SORTIEPPR-ID"),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		ExternalID:            unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
		Grantor:               unifieddatalibrary.String("SMITH"),
		Number:                unifieddatalibrary.String("07-21-07W"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Remarks:               unifieddatalibrary.String("PPR remark"),
		Requestor:             unifieddatalibrary.String("jsmith1"),
		StartTime:             unifieddatalibrary.Time(time.Now()),
		Type:                  unifieddatalibrary.SortiepprNewParamsTypeM,
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiepprUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sortieppr.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SortiepprUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SortiepprUpdateParamsDataModeTest,
			IDSortie:              "4ef3d1e8-ab08-ab70-498f-edc479734e5c",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SORTIEPPR-ID"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ExternalID:            unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
			Grantor:               unifieddatalibrary.String("SMITH"),
			Number:                unifieddatalibrary.String("07-21-07W"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Remarks:               unifieddatalibrary.String("PPR remark"),
			Requestor:             unifieddatalibrary.String("jsmith1"),
			StartTime:             unifieddatalibrary.Time(time.Now()),
			Type:                  unifieddatalibrary.SortiepprUpdateParamsTypeM,
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

func TestSortiepprListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sortieppr.List(context.TODO(), unifieddatalibrary.SortiepprListParams{
		IDSortie:    "idSortie",
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

func TestSortiepprDelete(t *testing.T) {
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
	err := client.Sortieppr.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiepprCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sortieppr.Count(context.TODO(), unifieddatalibrary.SortiepprCountParams{
		IDSortie:    "idSortie",
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

func TestSortiepprNewBulk(t *testing.T) {
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
	err := client.Sortieppr.NewBulk(context.TODO(), unifieddatalibrary.SortiepprNewBulkParams{
		Body: []unifieddatalibrary.SortiepprNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSortie:              "4ef3d1e8-ab08-ab70-498f-edc479734e5c",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SORTIEPPR-ID"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ExternalID:            unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
			Grantor:               unifieddatalibrary.String("SMITH"),
			Number:                unifieddatalibrary.String("07-21-07W"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Remarks:               unifieddatalibrary.String("PPR remark"),
			Requestor:             unifieddatalibrary.String("jsmith1"),
			StartTime:             unifieddatalibrary.Time(time.Now()),
			Type:                  "M",
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

func TestSortiepprGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sortieppr.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SortiepprGetParams{
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

func TestSortiepprQueryhelp(t *testing.T) {
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
	err := client.Sortieppr.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiepprTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sortieppr.Tuple(context.TODO(), unifieddatalibrary.SortiepprTupleParams{
		Columns:     "columns",
		IDSortie:    "idSortie",
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

func TestSortiepprUnvalidatedPublish(t *testing.T) {
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
	err := client.Sortieppr.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SortiepprUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SortiepprUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			IDSortie:              "4ef3d1e8-ab08-ab70-498f-edc479734e5c",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SORTIEPPR-ID"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			ExternalID:            unifieddatalibrary.String("aa714f4d52a37ab1a00b21af9566e379"),
			Grantor:               unifieddatalibrary.String("SMITH"),
			Number:                unifieddatalibrary.String("07-21-07W"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Remarks:               unifieddatalibrary.String("PPR remark"),
			Requestor:             unifieddatalibrary.String("jsmith1"),
			StartTime:             unifieddatalibrary.Time(time.Now()),
			Type:                  "M",
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
