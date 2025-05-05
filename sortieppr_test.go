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

func TestSortiePprNewWithOptionalParams(t *testing.T) {
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
	err := client.SortiePpr.New(context.TODO(), unifieddatalibrary.SortiePprNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SortiePprNewParamsDataModeTest,
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
		Type:                  unifieddatalibrary.SortiePprNewParamsTypeM,
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiePprUpdateWithOptionalParams(t *testing.T) {
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
	err := client.SortiePpr.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SortiePprUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SortiePprUpdateParamsDataModeTest,
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
			Type:                  unifieddatalibrary.SortiePprUpdateParamsTypeM,
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

func TestSortiePprListWithOptionalParams(t *testing.T) {
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
	_, err := client.SortiePpr.List(context.TODO(), unifieddatalibrary.SortiePprListParams{
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

func TestSortiePprDelete(t *testing.T) {
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
	err := client.SortiePpr.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiePprCountWithOptionalParams(t *testing.T) {
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
	_, err := client.SortiePpr.Count(context.TODO(), unifieddatalibrary.SortiePprCountParams{
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

func TestSortiePprNewBulk(t *testing.T) {
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
	err := client.SortiePpr.NewBulk(context.TODO(), unifieddatalibrary.SortiePprNewBulkParams{
		Body: []unifieddatalibrary.SortiePprNewBulkParamsBody{{
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

func TestSortiePprGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SortiePpr.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SortiePprGetParams{
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

func TestSortiePprQueryhelp(t *testing.T) {
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
	err := client.SortiePpr.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSortiePprTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.SortiePpr.Tuple(context.TODO(), unifieddatalibrary.SortiePprTupleParams{
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

func TestSortiePprUnvalidatedPublish(t *testing.T) {
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
	err := client.SortiePpr.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SortiePprUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SortiePprUnvalidatedPublishParamsBody{{
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
