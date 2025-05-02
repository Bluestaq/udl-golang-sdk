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

func TestDiplomaticClearanceCountryNewWithOptionalParams(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.New(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryNewParams{
		ClassificationMarking: "U",
		CountryCode:           "NL",
		DataMode:              unifieddatalibrary.DiplomaticClearanceCountryNewParamsDataModeTest,
		LastChangedDate:       time.Now(),
		Source:                "Bluestaq",
		ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
		AcceptsDms:            unifieddatalibrary.Bool(true),
		AcceptsEmail:          unifieddatalibrary.Bool(true),
		AcceptsFax:            unifieddatalibrary.Bool(true),
		AcceptsSiprNet:        unifieddatalibrary.Bool(false),
		Agency:                unifieddatalibrary.String("TACC"),
		AltCountryCode:        unifieddatalibrary.String("IZ"),
		CloseTime:             unifieddatalibrary.String("16:00"),
		CountryID:             unifieddatalibrary.String("GDSSBL010412140742262246"),
		CountryName:           unifieddatalibrary.String("NETHERLANDS"),
		CountryRemark:         unifieddatalibrary.String("Amsterdam airport EHAM will not accept hazardous cargo."),
		DiplomaticClearanceCountryContacts: []unifieddatalibrary.DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryContact{{
			AhNum:           unifieddatalibrary.String("256039858"),
			AhSpdDialCode:   unifieddatalibrary.String("75"),
			CommNum:         unifieddatalibrary.String("904716104"),
			CommSpdDialCode: unifieddatalibrary.String("74"),
			ContactID:       unifieddatalibrary.String("GDSSMC112108191329534522"),
			ContactName:     unifieddatalibrary.String("John Smith"),
			ContactRemark:   unifieddatalibrary.String("Contact remark"),
			DsnNum:          unifieddatalibrary.String("513827215"),
			DsnSpdDialCode:  unifieddatalibrary.String("94"),
			FaxNum:          unifieddatalibrary.String("571654897"),
			NiprNum:         unifieddatalibrary.String("525574441"),
			SiprNum:         unifieddatalibrary.String("546144352"),
		}},
		DiplomaticClearanceCountryEntryExitPoints: []unifieddatalibrary.DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryEntryExitPoint{{
			IsEntry:   unifieddatalibrary.Bool(true),
			IsExit:    unifieddatalibrary.Bool(true),
			PointName: unifieddatalibrary.String("BATEL"),
		}},
		DiplomaticClearanceCountryProfiles: []unifieddatalibrary.DiplomaticClearanceCountryNewParamsDiplomaticClearanceCountryProfile{{
			CargoPaxRemark:           unifieddatalibrary.String("Cargo passenger remark"),
			ClearanceID:              unifieddatalibrary.String("MDCNPER231360050AAR"),
			CrewInfoRemark:           unifieddatalibrary.String("Crew info remark"),
			DefClearanceStatus:       unifieddatalibrary.String("R"),
			DefEntryRemark:           unifieddatalibrary.String("Default entry remark"),
			DefEntryTime:             unifieddatalibrary.String("15:00"),
			DefExitRemark:            unifieddatalibrary.String("Default exit remark"),
			DefExitTime:              unifieddatalibrary.String("17:00"),
			FltInfoRemark:            unifieddatalibrary.String("Flight info remark"),
			HazInfoRemark:            unifieddatalibrary.String("Hazmat remark"),
			LandDefProf:              unifieddatalibrary.Bool(true),
			LandLeadTime:             unifieddatalibrary.Int(7),
			LandLeadTimeRemark:       unifieddatalibrary.String("Landing lead time remark"),
			LandLeadTimeUnit:         unifieddatalibrary.String("Day"),
			LandValidPeriodMinus:     unifieddatalibrary.Int(0),
			LandValidPeriodPlus:      unifieddatalibrary.Int(72),
			LandValidPeriodRemark:    unifieddatalibrary.String("Landing valid period remark"),
			LandValidPeriodUnit:      unifieddatalibrary.String("Hour"),
			OverflyDefProf:           unifieddatalibrary.Bool(true),
			OverflyLeadTime:          unifieddatalibrary.Int(7),
			OverflyLeadTimeRemark:    unifieddatalibrary.String("Overfly remark"),
			OverflyLeadTimeUnit:      unifieddatalibrary.String("Day"),
			OverflyValidPeriodMinus:  unifieddatalibrary.Int(0),
			OverflyValidPeriodPlus:   unifieddatalibrary.Int(72),
			OverflyValidPeriodRemark: unifieddatalibrary.String("Overfly valid period remark"),
			OverflyValidPeriodUnit:   unifieddatalibrary.String("Hour"),
			Profile:                  unifieddatalibrary.String("Netherlands Non Haz Landing"),
			ProfileAgency:            unifieddatalibrary.String("USAFE"),
			ProfileID:                unifieddatalibrary.String("GDSSBL010412140742262247"),
			ProfileRemark:            unifieddatalibrary.String("Profile remark"),
			ReqAcAltName:             unifieddatalibrary.Bool(false),
			ReqAllHazInfo:            unifieddatalibrary.Bool(false),
			ReqAmcStdInfo:            unifieddatalibrary.Bool(false),
			ReqCargoList:             unifieddatalibrary.Bool(false),
			ReqCargoPax:              unifieddatalibrary.Bool(false),
			ReqClass1Info:            unifieddatalibrary.Bool(false),
			ReqClass9Info:            unifieddatalibrary.Bool(false),
			ReqCrewComp:              unifieddatalibrary.Bool(false),
			ReqCrewDetail:            unifieddatalibrary.Bool(false),
			ReqCrewInfo:              unifieddatalibrary.Bool(false),
			ReqDiv1Info:              unifieddatalibrary.Bool(false),
			ReqDv:                    unifieddatalibrary.Bool(false),
			ReqEntryExitCoord:        unifieddatalibrary.Bool(false),
			ReqFltInfo:               unifieddatalibrary.Bool(false),
			ReqFltPlanRoute:          unifieddatalibrary.Bool(false),
			ReqFundSource:            unifieddatalibrary.Bool(false),
			ReqHazInfo:               unifieddatalibrary.Bool(false),
			ReqIcao:                  unifieddatalibrary.Bool(false),
			ReqPassportInfo:          unifieddatalibrary.Bool(false),
			ReqRaven:                 unifieddatalibrary.Bool(false),
			ReqRepChange:             unifieddatalibrary.Bool(false),
			ReqTailNum:               unifieddatalibrary.Bool(false),
			ReqWeaponsInfo:           unifieddatalibrary.Bool(false),
			UndefinedCrewReporting:   unifieddatalibrary.Bool(false),
		}},
		ExistingProfile: unifieddatalibrary.Bool(true),
		GmtOffset:       unifieddatalibrary.String("-04:30"),
		OfficeName:      unifieddatalibrary.String("DAO.EU"),
		OfficePoc:       unifieddatalibrary.String("John Smith"),
		OfficeRemark:    unifieddatalibrary.String("Diplomatic clearance office remark"),
		OpenFri:         unifieddatalibrary.Bool(true),
		OpenMon:         unifieddatalibrary.Bool(true),
		OpenSat:         unifieddatalibrary.Bool(false),
		OpenSun:         unifieddatalibrary.Bool(false),
		OpenThu:         unifieddatalibrary.Bool(true),
		OpenTime:        unifieddatalibrary.String("07:00"),
		OpenTue:         unifieddatalibrary.Bool(true),
		OpenWed:         unifieddatalibrary.Bool(true),
		Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceCountryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Country.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.DiplomaticClearanceCountryGetParams{
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

func TestDiplomaticClearanceCountryUpdateWithOptionalParams(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.DiplomaticClearanceCountryUpdateParams{
			ClassificationMarking: "U",
			CountryCode:           "NL",
			DataMode:              unifieddatalibrary.DiplomaticClearanceCountryUpdateParamsDataModeTest,
			LastChangedDate:       time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
			AcceptsDms:            unifieddatalibrary.Bool(true),
			AcceptsEmail:          unifieddatalibrary.Bool(true),
			AcceptsFax:            unifieddatalibrary.Bool(true),
			AcceptsSiprNet:        unifieddatalibrary.Bool(false),
			Agency:                unifieddatalibrary.String("TACC"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			CloseTime:             unifieddatalibrary.String("16:00"),
			CountryID:             unifieddatalibrary.String("GDSSBL010412140742262246"),
			CountryName:           unifieddatalibrary.String("NETHERLANDS"),
			CountryRemark:         unifieddatalibrary.String("Amsterdam airport EHAM will not accept hazardous cargo."),
			DiplomaticClearanceCountryContacts: []unifieddatalibrary.DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryContact{{
				AhNum:           unifieddatalibrary.String("256039858"),
				AhSpdDialCode:   unifieddatalibrary.String("75"),
				CommNum:         unifieddatalibrary.String("904716104"),
				CommSpdDialCode: unifieddatalibrary.String("74"),
				ContactID:       unifieddatalibrary.String("GDSSMC112108191329534522"),
				ContactName:     unifieddatalibrary.String("John Smith"),
				ContactRemark:   unifieddatalibrary.String("Contact remark"),
				DsnNum:          unifieddatalibrary.String("513827215"),
				DsnSpdDialCode:  unifieddatalibrary.String("94"),
				FaxNum:          unifieddatalibrary.String("571654897"),
				NiprNum:         unifieddatalibrary.String("525574441"),
				SiprNum:         unifieddatalibrary.String("546144352"),
			}},
			DiplomaticClearanceCountryEntryExitPoints: []unifieddatalibrary.DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryEntryExitPoint{{
				IsEntry:   unifieddatalibrary.Bool(true),
				IsExit:    unifieddatalibrary.Bool(true),
				PointName: unifieddatalibrary.String("BATEL"),
			}},
			DiplomaticClearanceCountryProfiles: []unifieddatalibrary.DiplomaticClearanceCountryUpdateParamsDiplomaticClearanceCountryProfile{{
				CargoPaxRemark:           unifieddatalibrary.String("Cargo passenger remark"),
				ClearanceID:              unifieddatalibrary.String("MDCNPER231360050AAR"),
				CrewInfoRemark:           unifieddatalibrary.String("Crew info remark"),
				DefClearanceStatus:       unifieddatalibrary.String("R"),
				DefEntryRemark:           unifieddatalibrary.String("Default entry remark"),
				DefEntryTime:             unifieddatalibrary.String("15:00"),
				DefExitRemark:            unifieddatalibrary.String("Default exit remark"),
				DefExitTime:              unifieddatalibrary.String("17:00"),
				FltInfoRemark:            unifieddatalibrary.String("Flight info remark"),
				HazInfoRemark:            unifieddatalibrary.String("Hazmat remark"),
				LandDefProf:              unifieddatalibrary.Bool(true),
				LandLeadTime:             unifieddatalibrary.Int(7),
				LandLeadTimeRemark:       unifieddatalibrary.String("Landing lead time remark"),
				LandLeadTimeUnit:         unifieddatalibrary.String("Day"),
				LandValidPeriodMinus:     unifieddatalibrary.Int(0),
				LandValidPeriodPlus:      unifieddatalibrary.Int(72),
				LandValidPeriodRemark:    unifieddatalibrary.String("Landing valid period remark"),
				LandValidPeriodUnit:      unifieddatalibrary.String("Hour"),
				OverflyDefProf:           unifieddatalibrary.Bool(true),
				OverflyLeadTime:          unifieddatalibrary.Int(7),
				OverflyLeadTimeRemark:    unifieddatalibrary.String("Overfly remark"),
				OverflyLeadTimeUnit:      unifieddatalibrary.String("Day"),
				OverflyValidPeriodMinus:  unifieddatalibrary.Int(0),
				OverflyValidPeriodPlus:   unifieddatalibrary.Int(72),
				OverflyValidPeriodRemark: unifieddatalibrary.String("Overfly valid period remark"),
				OverflyValidPeriodUnit:   unifieddatalibrary.String("Hour"),
				Profile:                  unifieddatalibrary.String("Netherlands Non Haz Landing"),
				ProfileAgency:            unifieddatalibrary.String("USAFE"),
				ProfileID:                unifieddatalibrary.String("GDSSBL010412140742262247"),
				ProfileRemark:            unifieddatalibrary.String("Profile remark"),
				ReqAcAltName:             unifieddatalibrary.Bool(false),
				ReqAllHazInfo:            unifieddatalibrary.Bool(false),
				ReqAmcStdInfo:            unifieddatalibrary.Bool(false),
				ReqCargoList:             unifieddatalibrary.Bool(false),
				ReqCargoPax:              unifieddatalibrary.Bool(false),
				ReqClass1Info:            unifieddatalibrary.Bool(false),
				ReqClass9Info:            unifieddatalibrary.Bool(false),
				ReqCrewComp:              unifieddatalibrary.Bool(false),
				ReqCrewDetail:            unifieddatalibrary.Bool(false),
				ReqCrewInfo:              unifieddatalibrary.Bool(false),
				ReqDiv1Info:              unifieddatalibrary.Bool(false),
				ReqDv:                    unifieddatalibrary.Bool(false),
				ReqEntryExitCoord:        unifieddatalibrary.Bool(false),
				ReqFltInfo:               unifieddatalibrary.Bool(false),
				ReqFltPlanRoute:          unifieddatalibrary.Bool(false),
				ReqFundSource:            unifieddatalibrary.Bool(false),
				ReqHazInfo:               unifieddatalibrary.Bool(false),
				ReqIcao:                  unifieddatalibrary.Bool(false),
				ReqPassportInfo:          unifieddatalibrary.Bool(false),
				ReqRaven:                 unifieddatalibrary.Bool(false),
				ReqRepChange:             unifieddatalibrary.Bool(false),
				ReqTailNum:               unifieddatalibrary.Bool(false),
				ReqWeaponsInfo:           unifieddatalibrary.Bool(false),
				UndefinedCrewReporting:   unifieddatalibrary.Bool(false),
			}},
			ExistingProfile: unifieddatalibrary.Bool(true),
			GmtOffset:       unifieddatalibrary.String("-04:30"),
			OfficeName:      unifieddatalibrary.String("DAO.EU"),
			OfficePoc:       unifieddatalibrary.String("John Smith"),
			OfficeRemark:    unifieddatalibrary.String("Diplomatic clearance office remark"),
			OpenFri:         unifieddatalibrary.Bool(true),
			OpenMon:         unifieddatalibrary.Bool(true),
			OpenSat:         unifieddatalibrary.Bool(false),
			OpenSun:         unifieddatalibrary.Bool(false),
			OpenThu:         unifieddatalibrary.Bool(true),
			OpenTime:        unifieddatalibrary.String("07:00"),
			OpenTue:         unifieddatalibrary.Bool(true),
			OpenWed:         unifieddatalibrary.Bool(true),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestDiplomaticClearanceCountryListWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Country.List(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryListParams{
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

func TestDiplomaticClearanceCountryDelete(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceCountryCountWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Country.Count(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryCountParams{
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

func TestDiplomaticClearanceCountryNewBulk(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.NewBulk(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryNewBulkParams{
		Body: []unifieddatalibrary.DiplomaticClearanceCountryNewBulkParamsBody{{
			ClassificationMarking: "U",
			CountryCode:           "NL",
			DataMode:              "TEST",
			LastChangedDate:       time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
			AcceptsDms:            unifieddatalibrary.Bool(true),
			AcceptsEmail:          unifieddatalibrary.Bool(true),
			AcceptsFax:            unifieddatalibrary.Bool(true),
			AcceptsSiprNet:        unifieddatalibrary.Bool(false),
			Agency:                unifieddatalibrary.String("TACC"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			CloseTime:             unifieddatalibrary.String("16:00"),
			CountryID:             unifieddatalibrary.String("GDSSBL010412140742262246"),
			CountryName:           unifieddatalibrary.String("NETHERLANDS"),
			CountryRemark:         unifieddatalibrary.String("Amsterdam airport EHAM will not accept hazardous cargo."),
			DiplomaticClearanceCountryContacts: []unifieddatalibrary.DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryContact{{
				AhNum:           unifieddatalibrary.String("256039858"),
				AhSpdDialCode:   unifieddatalibrary.String("75"),
				CommNum:         unifieddatalibrary.String("904716104"),
				CommSpdDialCode: unifieddatalibrary.String("74"),
				ContactID:       unifieddatalibrary.String("GDSSMC112108191329534522"),
				ContactName:     unifieddatalibrary.String("John Smith"),
				ContactRemark:   unifieddatalibrary.String("Contact remark"),
				DsnNum:          unifieddatalibrary.String("513827215"),
				DsnSpdDialCode:  unifieddatalibrary.String("94"),
				FaxNum:          unifieddatalibrary.String("571654897"),
				NiprNum:         unifieddatalibrary.String("525574441"),
				SiprNum:         unifieddatalibrary.String("546144352"),
			}},
			DiplomaticClearanceCountryEntryExitPoints: []unifieddatalibrary.DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryEntryExitPoint{{
				IsEntry:   unifieddatalibrary.Bool(true),
				IsExit:    unifieddatalibrary.Bool(true),
				PointName: unifieddatalibrary.String("BATEL"),
			}},
			DiplomaticClearanceCountryProfiles: []unifieddatalibrary.DiplomaticClearanceCountryNewBulkParamsBodyDiplomaticClearanceCountryProfile{{
				CargoPaxRemark:           unifieddatalibrary.String("Cargo passenger remark"),
				ClearanceID:              unifieddatalibrary.String("MDCNPER231360050AAR"),
				CrewInfoRemark:           unifieddatalibrary.String("Crew info remark"),
				DefClearanceStatus:       unifieddatalibrary.String("R"),
				DefEntryRemark:           unifieddatalibrary.String("Default entry remark"),
				DefEntryTime:             unifieddatalibrary.String("15:00"),
				DefExitRemark:            unifieddatalibrary.String("Default exit remark"),
				DefExitTime:              unifieddatalibrary.String("17:00"),
				FltInfoRemark:            unifieddatalibrary.String("Flight info remark"),
				HazInfoRemark:            unifieddatalibrary.String("Hazmat remark"),
				LandDefProf:              unifieddatalibrary.Bool(true),
				LandLeadTime:             unifieddatalibrary.Int(7),
				LandLeadTimeRemark:       unifieddatalibrary.String("Landing lead time remark"),
				LandLeadTimeUnit:         unifieddatalibrary.String("Day"),
				LandValidPeriodMinus:     unifieddatalibrary.Int(0),
				LandValidPeriodPlus:      unifieddatalibrary.Int(72),
				LandValidPeriodRemark:    unifieddatalibrary.String("Landing valid period remark"),
				LandValidPeriodUnit:      unifieddatalibrary.String("Hour"),
				OverflyDefProf:           unifieddatalibrary.Bool(true),
				OverflyLeadTime:          unifieddatalibrary.Int(7),
				OverflyLeadTimeRemark:    unifieddatalibrary.String("Overfly remark"),
				OverflyLeadTimeUnit:      unifieddatalibrary.String("Day"),
				OverflyValidPeriodMinus:  unifieddatalibrary.Int(0),
				OverflyValidPeriodPlus:   unifieddatalibrary.Int(72),
				OverflyValidPeriodRemark: unifieddatalibrary.String("Overfly valid period remark"),
				OverflyValidPeriodUnit:   unifieddatalibrary.String("Hour"),
				Profile:                  unifieddatalibrary.String("Netherlands Non Haz Landing"),
				ProfileAgency:            unifieddatalibrary.String("USAFE"),
				ProfileID:                unifieddatalibrary.String("GDSSBL010412140742262247"),
				ProfileRemark:            unifieddatalibrary.String("Profile remark"),
				ReqAcAltName:             unifieddatalibrary.Bool(false),
				ReqAllHazInfo:            unifieddatalibrary.Bool(false),
				ReqAmcStdInfo:            unifieddatalibrary.Bool(false),
				ReqCargoList:             unifieddatalibrary.Bool(false),
				ReqCargoPax:              unifieddatalibrary.Bool(false),
				ReqClass1Info:            unifieddatalibrary.Bool(false),
				ReqClass9Info:            unifieddatalibrary.Bool(false),
				ReqCrewComp:              unifieddatalibrary.Bool(false),
				ReqCrewDetail:            unifieddatalibrary.Bool(false),
				ReqCrewInfo:              unifieddatalibrary.Bool(false),
				ReqDiv1Info:              unifieddatalibrary.Bool(false),
				ReqDv:                    unifieddatalibrary.Bool(false),
				ReqEntryExitCoord:        unifieddatalibrary.Bool(false),
				ReqFltInfo:               unifieddatalibrary.Bool(false),
				ReqFltPlanRoute:          unifieddatalibrary.Bool(false),
				ReqFundSource:            unifieddatalibrary.Bool(false),
				ReqHazInfo:               unifieddatalibrary.Bool(false),
				ReqIcao:                  unifieddatalibrary.Bool(false),
				ReqPassportInfo:          unifieddatalibrary.Bool(false),
				ReqRaven:                 unifieddatalibrary.Bool(false),
				ReqRepChange:             unifieddatalibrary.Bool(false),
				ReqTailNum:               unifieddatalibrary.Bool(false),
				ReqWeaponsInfo:           unifieddatalibrary.Bool(false),
				UndefinedCrewReporting:   unifieddatalibrary.Bool(false),
			}},
			ExistingProfile: unifieddatalibrary.Bool(true),
			GmtOffset:       unifieddatalibrary.String("-04:30"),
			OfficeName:      unifieddatalibrary.String("DAO.EU"),
			OfficePoc:       unifieddatalibrary.String("John Smith"),
			OfficeRemark:    unifieddatalibrary.String("Diplomatic clearance office remark"),
			OpenFri:         unifieddatalibrary.Bool(true),
			OpenMon:         unifieddatalibrary.Bool(true),
			OpenSat:         unifieddatalibrary.Bool(false),
			OpenSun:         unifieddatalibrary.Bool(false),
			OpenThu:         unifieddatalibrary.Bool(true),
			OpenTime:        unifieddatalibrary.String("07:00"),
			OpenTue:         unifieddatalibrary.Bool(true),
			OpenWed:         unifieddatalibrary.Bool(true),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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

func TestDiplomaticClearanceCountryQueryHelp(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiplomaticClearanceCountryTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.DiplomaticClearance.Country.Tuple(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryTupleParams{
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

func TestDiplomaticClearanceCountryUnvalidatedPublish(t *testing.T) {
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
	err := client.DiplomaticClearance.Country.UnvalidatedPublish(context.TODO(), unifieddatalibrary.DiplomaticClearanceCountryUnvalidatedPublishParams{
		Body: []unifieddatalibrary.DiplomaticClearanceCountryUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			CountryCode:           "NL",
			DataMode:              "TEST",
			LastChangedDate:       time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("25059135-4afc-45c2-b78b-d6e843dbd96d"),
			AcceptsDms:            unifieddatalibrary.Bool(true),
			AcceptsEmail:          unifieddatalibrary.Bool(true),
			AcceptsFax:            unifieddatalibrary.Bool(true),
			AcceptsSiprNet:        unifieddatalibrary.Bool(false),
			Agency:                unifieddatalibrary.String("TACC"),
			AltCountryCode:        unifieddatalibrary.String("IZ"),
			CloseTime:             unifieddatalibrary.String("16:00"),
			CountryID:             unifieddatalibrary.String("GDSSBL010412140742262246"),
			CountryName:           unifieddatalibrary.String("NETHERLANDS"),
			CountryRemark:         unifieddatalibrary.String("Amsterdam airport EHAM will not accept hazardous cargo."),
			DiplomaticClearanceCountryContacts: []unifieddatalibrary.DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryContact{{
				AhNum:           unifieddatalibrary.String("256039858"),
				AhSpdDialCode:   unifieddatalibrary.String("75"),
				CommNum:         unifieddatalibrary.String("904716104"),
				CommSpdDialCode: unifieddatalibrary.String("74"),
				ContactID:       unifieddatalibrary.String("GDSSMC112108191329534522"),
				ContactName:     unifieddatalibrary.String("John Smith"),
				ContactRemark:   unifieddatalibrary.String("Contact remark"),
				DsnNum:          unifieddatalibrary.String("513827215"),
				DsnSpdDialCode:  unifieddatalibrary.String("94"),
				FaxNum:          unifieddatalibrary.String("571654897"),
				NiprNum:         unifieddatalibrary.String("525574441"),
				SiprNum:         unifieddatalibrary.String("546144352"),
			}},
			DiplomaticClearanceCountryEntryExitPoints: []unifieddatalibrary.DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryEntryExitPoint{{
				IsEntry:   unifieddatalibrary.Bool(true),
				IsExit:    unifieddatalibrary.Bool(true),
				PointName: unifieddatalibrary.String("BATEL"),
			}},
			DiplomaticClearanceCountryProfiles: []unifieddatalibrary.DiplomaticClearanceCountryUnvalidatedPublishParamsBodyDiplomaticClearanceCountryProfile{{
				CargoPaxRemark:           unifieddatalibrary.String("Cargo passenger remark"),
				ClearanceID:              unifieddatalibrary.String("MDCNPER231360050AAR"),
				CrewInfoRemark:           unifieddatalibrary.String("Crew info remark"),
				DefClearanceStatus:       unifieddatalibrary.String("R"),
				DefEntryRemark:           unifieddatalibrary.String("Default entry remark"),
				DefEntryTime:             unifieddatalibrary.String("15:00"),
				DefExitRemark:            unifieddatalibrary.String("Default exit remark"),
				DefExitTime:              unifieddatalibrary.String("17:00"),
				FltInfoRemark:            unifieddatalibrary.String("Flight info remark"),
				HazInfoRemark:            unifieddatalibrary.String("Hazmat remark"),
				LandDefProf:              unifieddatalibrary.Bool(true),
				LandLeadTime:             unifieddatalibrary.Int(7),
				LandLeadTimeRemark:       unifieddatalibrary.String("Landing lead time remark"),
				LandLeadTimeUnit:         unifieddatalibrary.String("Day"),
				LandValidPeriodMinus:     unifieddatalibrary.Int(0),
				LandValidPeriodPlus:      unifieddatalibrary.Int(72),
				LandValidPeriodRemark:    unifieddatalibrary.String("Landing valid period remark"),
				LandValidPeriodUnit:      unifieddatalibrary.String("Hour"),
				OverflyDefProf:           unifieddatalibrary.Bool(true),
				OverflyLeadTime:          unifieddatalibrary.Int(7),
				OverflyLeadTimeRemark:    unifieddatalibrary.String("Overfly remark"),
				OverflyLeadTimeUnit:      unifieddatalibrary.String("Day"),
				OverflyValidPeriodMinus:  unifieddatalibrary.Int(0),
				OverflyValidPeriodPlus:   unifieddatalibrary.Int(72),
				OverflyValidPeriodRemark: unifieddatalibrary.String("Overfly valid period remark"),
				OverflyValidPeriodUnit:   unifieddatalibrary.String("Hour"),
				Profile:                  unifieddatalibrary.String("Netherlands Non Haz Landing"),
				ProfileAgency:            unifieddatalibrary.String("USAFE"),
				ProfileID:                unifieddatalibrary.String("GDSSBL010412140742262247"),
				ProfileRemark:            unifieddatalibrary.String("Profile remark"),
				ReqAcAltName:             unifieddatalibrary.Bool(false),
				ReqAllHazInfo:            unifieddatalibrary.Bool(false),
				ReqAmcStdInfo:            unifieddatalibrary.Bool(false),
				ReqCargoList:             unifieddatalibrary.Bool(false),
				ReqCargoPax:              unifieddatalibrary.Bool(false),
				ReqClass1Info:            unifieddatalibrary.Bool(false),
				ReqClass9Info:            unifieddatalibrary.Bool(false),
				ReqCrewComp:              unifieddatalibrary.Bool(false),
				ReqCrewDetail:            unifieddatalibrary.Bool(false),
				ReqCrewInfo:              unifieddatalibrary.Bool(false),
				ReqDiv1Info:              unifieddatalibrary.Bool(false),
				ReqDv:                    unifieddatalibrary.Bool(false),
				ReqEntryExitCoord:        unifieddatalibrary.Bool(false),
				ReqFltInfo:               unifieddatalibrary.Bool(false),
				ReqFltPlanRoute:          unifieddatalibrary.Bool(false),
				ReqFundSource:            unifieddatalibrary.Bool(false),
				ReqHazInfo:               unifieddatalibrary.Bool(false),
				ReqIcao:                  unifieddatalibrary.Bool(false),
				ReqPassportInfo:          unifieddatalibrary.Bool(false),
				ReqRaven:                 unifieddatalibrary.Bool(false),
				ReqRepChange:             unifieddatalibrary.Bool(false),
				ReqTailNum:               unifieddatalibrary.Bool(false),
				ReqWeaponsInfo:           unifieddatalibrary.Bool(false),
				UndefinedCrewReporting:   unifieddatalibrary.Bool(false),
			}},
			ExistingProfile: unifieddatalibrary.Bool(true),
			GmtOffset:       unifieddatalibrary.String("-04:30"),
			OfficeName:      unifieddatalibrary.String("DAO.EU"),
			OfficePoc:       unifieddatalibrary.String("John Smith"),
			OfficeRemark:    unifieddatalibrary.String("Diplomatic clearance office remark"),
			OpenFri:         unifieddatalibrary.Bool(true),
			OpenMon:         unifieddatalibrary.Bool(true),
			OpenSat:         unifieddatalibrary.Bool(false),
			OpenSun:         unifieddatalibrary.Bool(false),
			OpenThu:         unifieddatalibrary.Bool(true),
			OpenTime:        unifieddatalibrary.String("07:00"),
			OpenTue:         unifieddatalibrary.Bool(true),
			OpenWed:         unifieddatalibrary.Bool(true),
			Origin:          unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
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
