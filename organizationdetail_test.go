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

func TestOrganizationdetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Organizationdetails.New(context.TODO(), unifieddatalibrary.OrganizationdetailNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OrganizationdetailNewParamsDataModeTest,
		IDOrganization:        "ORGANIZATION-ID",
		Name:                  "some.user",
		Source:                "some.user",
		ID:                    unifieddatalibrary.String("ORGANIZATIONDETAILS-ID"),
		Address1:              unifieddatalibrary.String("123 Main Street"),
		Address2:              unifieddatalibrary.String("Apt 4B"),
		Address3:              unifieddatalibrary.String("Colorado Springs CO, 80903"),
		Broker:                unifieddatalibrary.String("some.user"),
		Ceo:                   unifieddatalibrary.String("some.user"),
		Cfo:                   unifieddatalibrary.String("some.user"),
		Cto:                   unifieddatalibrary.String("some.user"),
		Description:           unifieddatalibrary.String("Example description"),
		Ebitda:                unifieddatalibrary.Float(123.4),
		Email:                 unifieddatalibrary.String("some_organization@organization.com"),
		FinancialNotes:        unifieddatalibrary.String("Example notes"),
		FinancialYearEndDate:  unifieddatalibrary.Time(time.Now()),
		FleetPlanNotes:        unifieddatalibrary.String("Example notes"),
		FormerOrgID:           unifieddatalibrary.String("FORMERORG-ID"),
		Ftes:                  unifieddatalibrary.Int(123),
		GeoAdminLevel1:        unifieddatalibrary.String("Colorado"),
		GeoAdminLevel2:        unifieddatalibrary.String("El Paso County"),
		GeoAdminLevel3:        unifieddatalibrary.String("Colorado Springs"),
		MassRanking:           unifieddatalibrary.Int(123),
		Origin:                unifieddatalibrary.String("some.user"),
		ParentOrgID:           unifieddatalibrary.String("PARENTORG-ID"),
		PostalCode:            unifieddatalibrary.String("80903"),
		Profit:                unifieddatalibrary.Float(123.4),
		Revenue:               unifieddatalibrary.Float(123.4),
		RevenueRanking:        unifieddatalibrary.Int(123),
		RiskManager:           unifieddatalibrary.String("some.user"),
		ServicesNotes:         unifieddatalibrary.String("Example notes"),
		Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationdetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Organizationdetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.OrganizationdetailUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.OrganizationdetailUpdateParamsDataModeTest,
			IDOrganization:        "ORGANIZATION-ID",
			Name:                  "some.user",
			Source:                "some.user",
			ID:                    unifieddatalibrary.String("ORGANIZATIONDETAILS-ID"),
			Address1:              unifieddatalibrary.String("123 Main Street"),
			Address2:              unifieddatalibrary.String("Apt 4B"),
			Address3:              unifieddatalibrary.String("Colorado Springs CO, 80903"),
			Broker:                unifieddatalibrary.String("some.user"),
			Ceo:                   unifieddatalibrary.String("some.user"),
			Cfo:                   unifieddatalibrary.String("some.user"),
			Cto:                   unifieddatalibrary.String("some.user"),
			Description:           unifieddatalibrary.String("Example description"),
			Ebitda:                unifieddatalibrary.Float(123.4),
			Email:                 unifieddatalibrary.String("some_organization@organization.com"),
			FinancialNotes:        unifieddatalibrary.String("Example notes"),
			FinancialYearEndDate:  unifieddatalibrary.Time(time.Now()),
			FleetPlanNotes:        unifieddatalibrary.String("Example notes"),
			FormerOrgID:           unifieddatalibrary.String("FORMERORG-ID"),
			Ftes:                  unifieddatalibrary.Int(123),
			GeoAdminLevel1:        unifieddatalibrary.String("Colorado"),
			GeoAdminLevel2:        unifieddatalibrary.String("El Paso County"),
			GeoAdminLevel3:        unifieddatalibrary.String("Colorado Springs"),
			MassRanking:           unifieddatalibrary.Int(123),
			Origin:                unifieddatalibrary.String("some.user"),
			ParentOrgID:           unifieddatalibrary.String("PARENTORG-ID"),
			PostalCode:            unifieddatalibrary.String("80903"),
			Profit:                unifieddatalibrary.Float(123.4),
			Revenue:               unifieddatalibrary.Float(123.4),
			RevenueRanking:        unifieddatalibrary.Int(123),
			RiskManager:           unifieddatalibrary.String("some.user"),
			ServicesNotes:         unifieddatalibrary.String("Example notes"),
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
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

func TestOrganizationdetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizationdetails.List(context.TODO(), unifieddatalibrary.OrganizationdetailListParams{
		Name:        "name",
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

func TestOrganizationdetailDelete(t *testing.T) {
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
	err := client.Organizationdetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationdetailFindBySourceWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizationdetails.FindBySource(context.TODO(), unifieddatalibrary.OrganizationdetailFindBySourceParams{
		Name:        "name",
		Source:      "source",
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

func TestOrganizationdetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizationdetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OrganizationdetailGetParams{
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
