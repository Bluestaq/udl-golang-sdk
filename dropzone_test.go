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

func TestDropzoneNewWithOptionalParams(t *testing.T) {
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
	err := client.Dropzone.New(context.TODO(), unifieddatalibrary.DropzoneNewParams{
		ClassificationMarking:   "U",
		DataMode:                unifieddatalibrary.DropzoneNewParamsDataModeTest,
		Lat:                     33.54,
		Lon:                     -117.162,
		Name:                    "Viper DZ",
		Source:                  "Bluestaq",
		ID:                      unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
		AltCountryCode:          unifieddatalibrary.String("USA"),
		AltCountryName:          unifieddatalibrary.String("United States of America"),
		ApprovalDate:            unifieddatalibrary.Time(time.Now()),
		Code:                    unifieddatalibrary.String("DZ"),
		CountryCode:             unifieddatalibrary.String("US"),
		CountryName:             unifieddatalibrary.String("United States"),
		ExpirationDate:          unifieddatalibrary.Time(time.Now()),
		ExtIdentifier:           unifieddatalibrary.String("1001"),
		IDSite:                  unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
		LastUpdate:              unifieddatalibrary.Time(time.Now()),
		Length:                  unifieddatalibrary.Float(549.1),
		Majcom:                  unifieddatalibrary.String("United States Northern Command"),
		NearestLoc:              unifieddatalibrary.String("March AFB"),
		OperationalApprovalDate: unifieddatalibrary.Time(time.Now()),
		Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PointName:               unifieddatalibrary.String("CENTER POINT"),
		Radius:                  unifieddatalibrary.Float(495.1),
		RecertDate:              unifieddatalibrary.Time(time.Now()),
		Remark:                  unifieddatalibrary.String("The text of the remark."),
		StateAbbr:               unifieddatalibrary.String("CA"),
		StateName:               unifieddatalibrary.String("CALIFORNIA"),
		SurveyDate:              unifieddatalibrary.Time(time.Now()),
		Width:                   unifieddatalibrary.Float(549.1),
		ZarID:                   unifieddatalibrary.String("1001"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDropzoneGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Dropzone.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.DropzoneGetParams{
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

func TestDropzoneUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Dropzone.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.DropzoneUpdateParams{
			ClassificationMarking:   "U",
			DataMode:                unifieddatalibrary.DropzoneUpdateParamsDataModeTest,
			Lat:                     33.54,
			Lon:                     -117.162,
			Name:                    "Viper DZ",
			Source:                  "Bluestaq",
			ID:                      unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			AltCountryCode:          unifieddatalibrary.String("USA"),
			AltCountryName:          unifieddatalibrary.String("United States of America"),
			ApprovalDate:            unifieddatalibrary.Time(time.Now()),
			Code:                    unifieddatalibrary.String("DZ"),
			CountryCode:             unifieddatalibrary.String("US"),
			CountryName:             unifieddatalibrary.String("United States"),
			ExpirationDate:          unifieddatalibrary.Time(time.Now()),
			ExtIdentifier:           unifieddatalibrary.String("1001"),
			IDSite:                  unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			LastUpdate:              unifieddatalibrary.Time(time.Now()),
			Length:                  unifieddatalibrary.Float(549.1),
			Majcom:                  unifieddatalibrary.String("United States Northern Command"),
			NearestLoc:              unifieddatalibrary.String("March AFB"),
			OperationalApprovalDate: unifieddatalibrary.Time(time.Now()),
			Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PointName:               unifieddatalibrary.String("CENTER POINT"),
			Radius:                  unifieddatalibrary.Float(495.1),
			RecertDate:              unifieddatalibrary.Time(time.Now()),
			Remark:                  unifieddatalibrary.String("The text of the remark."),
			StateAbbr:               unifieddatalibrary.String("CA"),
			StateName:               unifieddatalibrary.String("CALIFORNIA"),
			SurveyDate:              unifieddatalibrary.Time(time.Now()),
			Width:                   unifieddatalibrary.Float(549.1),
			ZarID:                   unifieddatalibrary.String("1001"),
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

func TestDropzoneDelete(t *testing.T) {
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
	err := client.Dropzone.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDropzoneCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Dropzone.Count(context.TODO(), unifieddatalibrary.DropzoneCountParams{
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

func TestDropzoneNewBulk(t *testing.T) {
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
	err := client.Dropzone.NewBulk(context.TODO(), unifieddatalibrary.DropzoneNewBulkParams{
		Body: []unifieddatalibrary.DropzoneNewBulkParamsBody{{
			ClassificationMarking:   "U",
			DataMode:                "TEST",
			Lat:                     33.54,
			Lon:                     -117.162,
			Name:                    "Viper DZ",
			Source:                  "Bluestaq",
			ID:                      unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			AltCountryCode:          unifieddatalibrary.String("USA"),
			AltCountryName:          unifieddatalibrary.String("United States of America"),
			ApprovalDate:            unifieddatalibrary.Time(time.Now()),
			Code:                    unifieddatalibrary.String("DZ"),
			CountryCode:             unifieddatalibrary.String("US"),
			CountryName:             unifieddatalibrary.String("United States"),
			ExpirationDate:          unifieddatalibrary.Time(time.Now()),
			ExtIdentifier:           unifieddatalibrary.String("1001"),
			IDSite:                  unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			LastUpdate:              unifieddatalibrary.Time(time.Now()),
			Length:                  unifieddatalibrary.Float(549.1),
			Majcom:                  unifieddatalibrary.String("United States Northern Command"),
			NearestLoc:              unifieddatalibrary.String("March AFB"),
			OperationalApprovalDate: unifieddatalibrary.Time(time.Now()),
			Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PointName:               unifieddatalibrary.String("CENTER POINT"),
			Radius:                  unifieddatalibrary.Float(495.1),
			RecertDate:              unifieddatalibrary.Time(time.Now()),
			Remark:                  unifieddatalibrary.String("The text of the remark."),
			StateAbbr:               unifieddatalibrary.String("CA"),
			StateName:               unifieddatalibrary.String("CALIFORNIA"),
			SurveyDate:              unifieddatalibrary.Time(time.Now()),
			Width:                   unifieddatalibrary.Float(549.1),
			ZarID:                   unifieddatalibrary.String("1001"),
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

func TestDropzoneQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Dropzone.Query(context.TODO(), unifieddatalibrary.DropzoneQueryParams{
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

func TestDropzoneQueryHelp(t *testing.T) {
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
	_, err := client.Dropzone.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDropzoneTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Dropzone.Tuple(context.TODO(), unifieddatalibrary.DropzoneTupleParams{
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

func TestDropzoneUnvalidatedPublish(t *testing.T) {
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
	err := client.Dropzone.UnvalidatedPublish(context.TODO(), unifieddatalibrary.DropzoneUnvalidatedPublishParams{
		Body: []unifieddatalibrary.DropzoneUnvalidatedPublishParamsBody{{
			ClassificationMarking:   "U",
			DataMode:                "TEST",
			Lat:                     33.54,
			Lon:                     -117.162,
			Name:                    "Viper DZ",
			Source:                  "Bluestaq",
			ID:                      unifieddatalibrary.String("3f28f60b-3a50-2aef-ac88-8e9d0e39912b"),
			AltCountryCode:          unifieddatalibrary.String("USA"),
			AltCountryName:          unifieddatalibrary.String("United States of America"),
			ApprovalDate:            unifieddatalibrary.Time(time.Now()),
			Code:                    unifieddatalibrary.String("DZ"),
			CountryCode:             unifieddatalibrary.String("US"),
			CountryName:             unifieddatalibrary.String("United States"),
			ExpirationDate:          unifieddatalibrary.Time(time.Now()),
			ExtIdentifier:           unifieddatalibrary.String("1001"),
			IDSite:                  unifieddatalibrary.String("a150b3ee-884b-b9ac-60a0-6408b4b16088"),
			LastUpdate:              unifieddatalibrary.Time(time.Now()),
			Length:                  unifieddatalibrary.Float(549.1),
			Majcom:                  unifieddatalibrary.String("United States Northern Command"),
			NearestLoc:              unifieddatalibrary.String("March AFB"),
			OperationalApprovalDate: unifieddatalibrary.Time(time.Now()),
			Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PointName:               unifieddatalibrary.String("CENTER POINT"),
			Radius:                  unifieddatalibrary.Float(495.1),
			RecertDate:              unifieddatalibrary.Time(time.Now()),
			Remark:                  unifieddatalibrary.String("The text of the remark."),
			StateAbbr:               unifieddatalibrary.String("CA"),
			StateName:               unifieddatalibrary.String("CALIFORNIA"),
			SurveyDate:              unifieddatalibrary.Time(time.Now()),
			Width:                   unifieddatalibrary.Float(549.1),
			ZarID:                   unifieddatalibrary.String("1001"),
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
