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

func TestSeradataspacecraftdetailNewWithOptionalParams(t *testing.T) {
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
	err := client.Seradataspacecraftdetails.New(context.TODO(), unifieddatalibrary.SeradataspacecraftdetailNewParams{
		ClassificationMarking:                    "U",
		DataMode:                                 unifieddatalibrary.SeradataspacecraftdetailNewParamsDataModeTest,
		Name:                                     "name",
		Source:                                   "Bluestaq",
		ID:                                       unifieddatalibrary.String("SERADATASPACECRAFTDETAILS-ID"),
		AdditionalMissionsGroups:                 unifieddatalibrary.String("additionalMissionsGroups"),
		Altitude:                                 unifieddatalibrary.Float(36036.6330576414),
		AnnualInsuredDepreciationFactor:          unifieddatalibrary.Float(1.23),
		AnnualInsuredDepreciationFactorEstimated: unifieddatalibrary.Bool(true),
		Apogee:                                   unifieddatalibrary.Float(1.23),
		BusID:                                    unifieddatalibrary.String("BUS-ID"),
		CapabilityLost:                           unifieddatalibrary.Float(1.23),
		CapacityLost:                             unifieddatalibrary.Float(1.23),
		CatalogNumber:                            unifieddatalibrary.Int(1),
		CollisionRiskCm:                          unifieddatalibrary.Float(1.43),
		CollisionRiskMm:                          unifieddatalibrary.Float(1.33),
		CombinedCostEstimated:                    unifieddatalibrary.Bool(true),
		CombinedNewCost:                          unifieddatalibrary.Float(1.23),
		CommercialLaunch:                         unifieddatalibrary.Bool(true),
		Constellation:                            unifieddatalibrary.String("GPS"),
		CostEstimated:                            unifieddatalibrary.Bool(true),
		CubesatDispenserType:                     unifieddatalibrary.String("cubesatDispenserType"),
		CurrentAge:                               unifieddatalibrary.Float(5.898630136986301),
		DateOfObservation:                        unifieddatalibrary.Time(time.Now()),
		Description:                              unifieddatalibrary.String("description"),
		DesignLife:                               unifieddatalibrary.Int(231),
		DryMass:                                  unifieddatalibrary.Float(1.23),
		ExpectedLife:                             unifieddatalibrary.Int(231),
		GeoPosition:                              unifieddatalibrary.Float(-8.23),
		IDOnOrbit:                                unifieddatalibrary.String("503"),
		Inclination:                              unifieddatalibrary.Float(1.23),
		InsuranceLossesTotal:                     unifieddatalibrary.Float(0.393),
		InsuranceNotes:                           unifieddatalibrary.String("Sample Notes"),
		InsurancePremiumAtLaunch:                 unifieddatalibrary.Float(1.23),
		InsurancePremiumAtLaunchEstimated:        unifieddatalibrary.Bool(true),
		InsuredAtLaunch:                          unifieddatalibrary.Bool(true),
		InsuredValueAtLaunch:                     unifieddatalibrary.Float(1.23),
		InsuredValueLaunchEstimated:              unifieddatalibrary.Bool(true),
		IntlNumber:                               unifieddatalibrary.String("number"),
		Lat:                                      unifieddatalibrary.Float(1.23),
		LaunchArranger:                           unifieddatalibrary.String("launchArranger"),
		LaunchArrangerCountry:                    unifieddatalibrary.String("USA"),
		LaunchCharacteristic:                     unifieddatalibrary.String("Expendable"),
		LaunchCost:                               unifieddatalibrary.Float(1.23),
		LaunchCostEstimated:                      unifieddatalibrary.Bool(true),
		LaunchCountry:                            unifieddatalibrary.String("USA"),
		LaunchDate:                               unifieddatalibrary.Time(time.Now()),
		LaunchDateRemarks:                        unifieddatalibrary.String("launchDateRemarks"),
		LaunchID:                                 unifieddatalibrary.String("11573"),
		LaunchMass:                               unifieddatalibrary.Float(1.23),
		LaunchNotes:                              unifieddatalibrary.String("Sample Notes"),
		LaunchNumber:                             unifieddatalibrary.String("FN040"),
		LaunchProvider:                           unifieddatalibrary.String("launchProvider"),
		LaunchProviderCountry:                    unifieddatalibrary.String("USA"),
		LaunchProviderFlightNumber:               unifieddatalibrary.String("launchProviderFlightNumber"),
		LaunchSiteID:                             unifieddatalibrary.String("28"),
		LaunchSiteName:                           unifieddatalibrary.String("launchSiteName"),
		LaunchType:                               unifieddatalibrary.String("Future"),
		LaunchVehicleID:                          unifieddatalibrary.String("123"),
		Leased:                                   unifieddatalibrary.Bool(true),
		LifeLost:                                 unifieddatalibrary.Float(1.23),
		Lon:                                      unifieddatalibrary.Float(1.23),
		MassCategory:                             unifieddatalibrary.String("2500 - 3500kg  - Large Satellite"),
		NameAtLaunch:                             unifieddatalibrary.String("nameAtLaunch"),
		NewCost:                                  unifieddatalibrary.Float(1.23),
		Notes:                                    unifieddatalibrary.String("Sample Notes"),
		NumHumans:                                unifieddatalibrary.Int(1),
		Operator:                                 unifieddatalibrary.String("operator"),
		OperatorCountry:                          unifieddatalibrary.String("USA"),
		OrbitCategory:                            unifieddatalibrary.String("GEO"),
		OrbitSubCategory:                         unifieddatalibrary.String("Geostationary"),
		OrderDate:                                unifieddatalibrary.Time(time.Now()),
		Origin:                                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Owner:                                    unifieddatalibrary.String("owner"),
		OwnerCountry:                             unifieddatalibrary.String("USA"),
		Perigee:                                  unifieddatalibrary.Float(1.23),
		Period:                                   unifieddatalibrary.Float(1.23),
		PrimaryMissionGroup:                      unifieddatalibrary.String("primaryMissionGroup"),
		PrimeManufacturerOrgID:                   unifieddatalibrary.String("05c43360-382e-4aa2-b875-ed28945ff2e5"),
		ProgramName:                              unifieddatalibrary.String("programName"),
		Quantity:                                 unifieddatalibrary.Int(1),
		ReusableFlights:                          unifieddatalibrary.String("reusableFlights"),
		ReusedHullName:                           unifieddatalibrary.String("reusedHullName"),
		Sector:                                   unifieddatalibrary.String("Commercial"),
		SerialNumber:                             unifieddatalibrary.String("serialNumber"),
		Stabilizer:                               unifieddatalibrary.String("3-Axis"),
		Status:                                   unifieddatalibrary.String("Inactive - Retired"),
		TotalClaims:                              unifieddatalibrary.Int(1),
		TotalFatalities:                          unifieddatalibrary.Int(1),
		TotalInjuries:                            unifieddatalibrary.Int(1),
		TotalPayloadPower:                        unifieddatalibrary.Float(1.23),
		YoutubeLaunchLink:                        unifieddatalibrary.String("youtubeLaunchLink"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataspacecraftdetailUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Seradataspacecraftdetails.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataspacecraftdetailUpdateParams{
			ClassificationMarking:                    "U",
			DataMode:                                 unifieddatalibrary.SeradataspacecraftdetailUpdateParamsDataModeTest,
			Name:                                     "name",
			Source:                                   "Bluestaq",
			ID:                                       unifieddatalibrary.String("SERADATASPACECRAFTDETAILS-ID"),
			AdditionalMissionsGroups:                 unifieddatalibrary.String("additionalMissionsGroups"),
			Altitude:                                 unifieddatalibrary.Float(36036.6330576414),
			AnnualInsuredDepreciationFactor:          unifieddatalibrary.Float(1.23),
			AnnualInsuredDepreciationFactorEstimated: unifieddatalibrary.Bool(true),
			Apogee:                                   unifieddatalibrary.Float(1.23),
			BusID:                                    unifieddatalibrary.String("BUS-ID"),
			CapabilityLost:                           unifieddatalibrary.Float(1.23),
			CapacityLost:                             unifieddatalibrary.Float(1.23),
			CatalogNumber:                            unifieddatalibrary.Int(1),
			CollisionRiskCm:                          unifieddatalibrary.Float(1.43),
			CollisionRiskMm:                          unifieddatalibrary.Float(1.33),
			CombinedCostEstimated:                    unifieddatalibrary.Bool(true),
			CombinedNewCost:                          unifieddatalibrary.Float(1.23),
			CommercialLaunch:                         unifieddatalibrary.Bool(true),
			Constellation:                            unifieddatalibrary.String("GPS"),
			CostEstimated:                            unifieddatalibrary.Bool(true),
			CubesatDispenserType:                     unifieddatalibrary.String("cubesatDispenserType"),
			CurrentAge:                               unifieddatalibrary.Float(5.898630136986301),
			DateOfObservation:                        unifieddatalibrary.Time(time.Now()),
			Description:                              unifieddatalibrary.String("description"),
			DesignLife:                               unifieddatalibrary.Int(231),
			DryMass:                                  unifieddatalibrary.Float(1.23),
			ExpectedLife:                             unifieddatalibrary.Int(231),
			GeoPosition:                              unifieddatalibrary.Float(-8.23),
			IDOnOrbit:                                unifieddatalibrary.String("503"),
			Inclination:                              unifieddatalibrary.Float(1.23),
			InsuranceLossesTotal:                     unifieddatalibrary.Float(0.393),
			InsuranceNotes:                           unifieddatalibrary.String("Sample Notes"),
			InsurancePremiumAtLaunch:                 unifieddatalibrary.Float(1.23),
			InsurancePremiumAtLaunchEstimated:        unifieddatalibrary.Bool(true),
			InsuredAtLaunch:                          unifieddatalibrary.Bool(true),
			InsuredValueAtLaunch:                     unifieddatalibrary.Float(1.23),
			InsuredValueLaunchEstimated:              unifieddatalibrary.Bool(true),
			IntlNumber:                               unifieddatalibrary.String("number"),
			Lat:                                      unifieddatalibrary.Float(1.23),
			LaunchArranger:                           unifieddatalibrary.String("launchArranger"),
			LaunchArrangerCountry:                    unifieddatalibrary.String("USA"),
			LaunchCharacteristic:                     unifieddatalibrary.String("Expendable"),
			LaunchCost:                               unifieddatalibrary.Float(1.23),
			LaunchCostEstimated:                      unifieddatalibrary.Bool(true),
			LaunchCountry:                            unifieddatalibrary.String("USA"),
			LaunchDate:                               unifieddatalibrary.Time(time.Now()),
			LaunchDateRemarks:                        unifieddatalibrary.String("launchDateRemarks"),
			LaunchID:                                 unifieddatalibrary.String("11573"),
			LaunchMass:                               unifieddatalibrary.Float(1.23),
			LaunchNotes:                              unifieddatalibrary.String("Sample Notes"),
			LaunchNumber:                             unifieddatalibrary.String("FN040"),
			LaunchProvider:                           unifieddatalibrary.String("launchProvider"),
			LaunchProviderCountry:                    unifieddatalibrary.String("USA"),
			LaunchProviderFlightNumber:               unifieddatalibrary.String("launchProviderFlightNumber"),
			LaunchSiteID:                             unifieddatalibrary.String("28"),
			LaunchSiteName:                           unifieddatalibrary.String("launchSiteName"),
			LaunchType:                               unifieddatalibrary.String("Future"),
			LaunchVehicleID:                          unifieddatalibrary.String("123"),
			Leased:                                   unifieddatalibrary.Bool(true),
			LifeLost:                                 unifieddatalibrary.Float(1.23),
			Lon:                                      unifieddatalibrary.Float(1.23),
			MassCategory:                             unifieddatalibrary.String("2500 - 3500kg  - Large Satellite"),
			NameAtLaunch:                             unifieddatalibrary.String("nameAtLaunch"),
			NewCost:                                  unifieddatalibrary.Float(1.23),
			Notes:                                    unifieddatalibrary.String("Sample Notes"),
			NumHumans:                                unifieddatalibrary.Int(1),
			Operator:                                 unifieddatalibrary.String("operator"),
			OperatorCountry:                          unifieddatalibrary.String("USA"),
			OrbitCategory:                            unifieddatalibrary.String("GEO"),
			OrbitSubCategory:                         unifieddatalibrary.String("Geostationary"),
			OrderDate:                                unifieddatalibrary.Time(time.Now()),
			Origin:                                   unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Owner:                                    unifieddatalibrary.String("owner"),
			OwnerCountry:                             unifieddatalibrary.String("USA"),
			Perigee:                                  unifieddatalibrary.Float(1.23),
			Period:                                   unifieddatalibrary.Float(1.23),
			PrimaryMissionGroup:                      unifieddatalibrary.String("primaryMissionGroup"),
			PrimeManufacturerOrgID:                   unifieddatalibrary.String("05c43360-382e-4aa2-b875-ed28945ff2e5"),
			ProgramName:                              unifieddatalibrary.String("programName"),
			Quantity:                                 unifieddatalibrary.Int(1),
			ReusableFlights:                          unifieddatalibrary.String("reusableFlights"),
			ReusedHullName:                           unifieddatalibrary.String("reusedHullName"),
			Sector:                                   unifieddatalibrary.String("Commercial"),
			SerialNumber:                             unifieddatalibrary.String("serialNumber"),
			Stabilizer:                               unifieddatalibrary.String("3-Axis"),
			Status:                                   unifieddatalibrary.String("Inactive - Retired"),
			TotalClaims:                              unifieddatalibrary.Int(1),
			TotalFatalities:                          unifieddatalibrary.Int(1),
			TotalInjuries:                            unifieddatalibrary.Int(1),
			TotalPayloadPower:                        unifieddatalibrary.Float(1.23),
			YoutubeLaunchLink:                        unifieddatalibrary.String("youtubeLaunchLink"),
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

func TestSeradataspacecraftdetailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataspacecraftdetails.List(context.TODO(), unifieddatalibrary.SeradataspacecraftdetailListParams{
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

func TestSeradataspacecraftdetailDelete(t *testing.T) {
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
	err := client.Seradataspacecraftdetails.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataspacecraftdetailCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataspacecraftdetails.Count(context.TODO(), unifieddatalibrary.SeradataspacecraftdetailCountParams{
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

func TestSeradataspacecraftdetailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataspacecraftdetails.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SeradataspacecraftdetailGetParams{
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

func TestSeradataspacecraftdetailQueryhelp(t *testing.T) {
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
	err := client.Seradataspacecraftdetails.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSeradataspacecraftdetailTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Seradataspacecraftdetails.Tuple(context.TODO(), unifieddatalibrary.SeradataspacecraftdetailTupleParams{
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
