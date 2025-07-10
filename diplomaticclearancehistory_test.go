// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestDiplomaticClearanceHistoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.History.List(context.TODO(), unifieddatalibrary.DiplomaticClearanceHistoryListParams{
		FirstDepDate: time.Now(),
		Columns:      unifieddatalibrary.String("columns"),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceHistoryAodrWithOptionalParams(t *testing.T) {
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
	err := client.DiplomaticClearance.History.Aodr(context.TODO(), unifieddatalibrary.DiplomaticClearanceHistoryAodrParams{
		FirstDepDate:    time.Now(),
		Columns:         unifieddatalibrary.String("columns"),
		FirstResult:     unifieddatalibrary.Int(0),
		MaxResults:      unifieddatalibrary.Int(0),
		Notification:    unifieddatalibrary.String("notification"),
		OutputDelimiter: unifieddatalibrary.String("outputDelimiter"),
		OutputFormat:    unifieddatalibrary.String("outputFormat"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceHistoryCountWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.History.Count(context.TODO(), unifieddatalibrary.DiplomaticClearanceHistoryCountParams{
		FirstDepDate: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
