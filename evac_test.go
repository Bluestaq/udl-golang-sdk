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

func TestEvacNewWithOptionalParams(t *testing.T) {
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
	err := client.Evac.New(context.TODO(), unifieddatalibrary.EvacNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EvacNewParamsDataModeTest,
		PickupLat:             75.1234,
		PickupLon:             175.1234,
		ReqTime:               time.Now(),
		Source:                "Bluestaq",
		Type:                  unifieddatalibrary.EvacNewParamsTypeRequest,
		ID:                    unifieddatalibrary.String("MEDEVACEVENT-ID"),
		CasualtyInfo: []unifieddatalibrary.EvacNewParamsCasualtyInfo{{
			Age: unifieddatalibrary.Int(35),
			Allergy: []unifieddatalibrary.EvacNewParamsCasualtyInfoAllergy{{
				Comments: unifieddatalibrary.String("Comments on the patient's allergies."),
				Type:     unifieddatalibrary.String("PENICILLIN"),
			}},
			BloodType:       unifieddatalibrary.String("O NEG"),
			BodyPart:        unifieddatalibrary.String("FACE"),
			BurialLocation:  []float64{-33.123, 150.33, 0.24},
			CallSign:        unifieddatalibrary.String("SHARK"),
			CareProviderUrn: unifieddatalibrary.String("CARE_PROVIDER-1"),
			CasualtyKey:     unifieddatalibrary.String("casualty-007"),
			CasualtyType:    unifieddatalibrary.String("DENTAL"),
			CollectionPoint: []float64{12.44, 122.55, 0.98},
			Comments:        unifieddatalibrary.String("Comments relating to this casualty info."),
			Condition: []unifieddatalibrary.EvacNewParamsCasualtyInfoCondition{{
				BodyPart: unifieddatalibrary.String("ANKLE LEFT FRONT"),
				Comments: unifieddatalibrary.String("Comments on the patient's condition."),
				Time:     unifieddatalibrary.Time(time.Now()),
				Type:     unifieddatalibrary.String("ACTIVITY LOW"),
			}},
			ContamType:      unifieddatalibrary.String("NONE"),
			Disposition:     unifieddatalibrary.String("EVACUATE WOUNDED"),
			DispositionType: unifieddatalibrary.String("EVACUATE"),
			Etiology: []unifieddatalibrary.EvacNewParamsCasualtyInfoEtiology{{
				BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
				Comments: unifieddatalibrary.String("Comments on the etiology info."),
				Time:     unifieddatalibrary.Time(time.Now()),
				Type:     unifieddatalibrary.String("BURN"),
			}},
			EvacType: unifieddatalibrary.String("GROUND"),
			Gender:   unifieddatalibrary.String("MALE"),
			HealthState: []unifieddatalibrary.EvacNewParamsCasualtyInfoHealthState{{
				HealthStateCode: unifieddatalibrary.String("BLUE"),
				MedConfFactor:   unifieddatalibrary.Int(1),
				Time:            unifieddatalibrary.Time(time.Now()),
				Type:            unifieddatalibrary.String("COGNITIVE"),
			}},
			Injury: []unifieddatalibrary.EvacNewParamsCasualtyInfoInjury{{
				BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
				Comments: unifieddatalibrary.String("Comments on the patient's injury."),
				Time:     unifieddatalibrary.Time(time.Now()),
				Type:     unifieddatalibrary.String("ABRASION"),
			}},
			Last4Ssn: unifieddatalibrary.String("1234"),
			Medication: []unifieddatalibrary.EvacNewParamsCasualtyInfoMedication{{
				AdminRoute: unifieddatalibrary.String("ORAL"),
				BodyPart:   unifieddatalibrary.String("ARM LEFT BACK"),
				Comments:   unifieddatalibrary.String("Comments on the patient's medication information."),
				Dose:       unifieddatalibrary.String("800mg"),
				Time:       unifieddatalibrary.Time(time.Now()),
				Type:       unifieddatalibrary.String("TYLENOL"),
			}},
			Name:            unifieddatalibrary.String("John Smith"),
			Nationality:     unifieddatalibrary.String("US"),
			OccSpeciality:   unifieddatalibrary.String("Healthcare"),
			PatientIdentity: unifieddatalibrary.String("FRIEND CIVILIAN"),
			PatientStatus:   unifieddatalibrary.String("US CIVILIAN"),
			PayGrade:        unifieddatalibrary.String("CIVILIAN"),
			Priority:        unifieddatalibrary.String("ROUTINE"),
			ReportGen:       unifieddatalibrary.String("DEVICE"),
			ReportTime:      unifieddatalibrary.Time(time.Now()),
			Service:         unifieddatalibrary.String("CIV"),
			SpecMedEquip:    []string{"OXYGEN", "HOIST"},
			Treatment: []unifieddatalibrary.EvacNewParamsCasualtyInfoTreatment{{
				BodyPart: unifieddatalibrary.String("CHEST"),
				Comments: unifieddatalibrary.String("Comments on the treatment info."),
				Time:     unifieddatalibrary.Time(time.Now()),
				Type:     unifieddatalibrary.String("BREATHING CHEST TUBE"),
			}},
			VitalSignData: []unifieddatalibrary.EvacNewParamsCasualtyInfoVitalSignData{{
				MedConfFactor: unifieddatalibrary.Int(1),
				Time:          unifieddatalibrary.Time(time.Now()),
				VitalSign:     unifieddatalibrary.String("HEART RATE"),
				VitalSign1:    unifieddatalibrary.Float(120),
				VitalSign2:    unifieddatalibrary.Float(80),
			}},
		}},
		Ce:        unifieddatalibrary.Float(10.1234),
		CntctFreq: unifieddatalibrary.Float(3.11),
		Comments:  unifieddatalibrary.String("Comments concerning mission"),
		EnemyData: []unifieddatalibrary.EvacNewParamsEnemyData{{
			DirToEnemy:        unifieddatalibrary.String("NORTH"),
			FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
			HlzRemarks:        unifieddatalibrary.String("Remarks about hot landing zone."),
			HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
		}},
		IDWeatherReport:   unifieddatalibrary.String("WeatherReport-ID"),
		Le:                unifieddatalibrary.Float(5.1234),
		MedevacID:         unifieddatalibrary.String("MedEvac-ID"),
		MedicReq:          unifieddatalibrary.Bool(true),
		MissionType:       unifieddatalibrary.String("GROUND"),
		NumAmbulatory:     unifieddatalibrary.Int(5),
		NumCasualties:     unifieddatalibrary.Int(5),
		NumKia:            unifieddatalibrary.Int(0),
		NumLitter:         unifieddatalibrary.Int(0),
		NumWia:            unifieddatalibrary.Int(3),
		ObstaclesRemarks:  unifieddatalibrary.String("N/A"),
		Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		PickupAlt:         unifieddatalibrary.Float(30.1234),
		PickupTime:        unifieddatalibrary.Time(time.Now()),
		ReqCallSign:       unifieddatalibrary.String("Bravo"),
		ReqNum:            unifieddatalibrary.String("MED.1.223908"),
		Terrain:           unifieddatalibrary.String("ROCKY"),
		TerrainRemarks:    unifieddatalibrary.String("N/A"),
		ZoneContrCallSign: unifieddatalibrary.String("Tango"),
		ZoneHot:           unifieddatalibrary.Bool(false),
		ZoneMarking:       unifieddatalibrary.String("ILLUMINATION"),
		ZoneMarkingColor:  unifieddatalibrary.String("RED"),
		ZoneName:          unifieddatalibrary.String("example-zone"),
		ZoneSecurity:      unifieddatalibrary.String("NO ENEMY"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvacGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Evac.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EvacGetParams{
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

func TestEvacListWithOptionalParams(t *testing.T) {
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
	_, err := client.Evac.List(context.TODO(), unifieddatalibrary.EvacListParams{
		ReqTime:     time.Now(),
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

func TestEvacCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Evac.Count(context.TODO(), unifieddatalibrary.EvacCountParams{
		ReqTime:     time.Now(),
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

func TestEvacNewBulk(t *testing.T) {
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
	err := client.Evac.NewBulk(context.TODO(), unifieddatalibrary.EvacNewBulkParams{
		Body: []unifieddatalibrary.EvacNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			PickupLat:             75.1234,
			PickupLon:             175.1234,
			ReqTime:               time.Now(),
			Source:                "Bluestaq",
			Type:                  "REQUEST",
			ID:                    unifieddatalibrary.String("MEDEVACEVENT-ID"),
			CasualtyInfo: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfo{{
				Age: unifieddatalibrary.Int(35),
				Allergy: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoAllergy{{
					Comments: unifieddatalibrary.String("Comments on the patient's allergies."),
					Type:     unifieddatalibrary.String("PENICILLIN"),
				}},
				BloodType:       unifieddatalibrary.String("O NEG"),
				BodyPart:        unifieddatalibrary.String("FACE"),
				BurialLocation:  []float64{-33.123, 150.33, 0.24},
				CallSign:        unifieddatalibrary.String("SHARK"),
				CareProviderUrn: unifieddatalibrary.String("CARE_PROVIDER-1"),
				CasualtyKey:     unifieddatalibrary.String("casualty-007"),
				CasualtyType:    unifieddatalibrary.String("DENTAL"),
				CollectionPoint: []float64{12.44, 122.55, 0.98},
				Comments:        unifieddatalibrary.String("Comments relating to this casualty info."),
				Condition: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoCondition{{
					BodyPart: unifieddatalibrary.String("ANKLE LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the patient's condition."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("ACTIVITY LOW"),
				}},
				ContamType:      unifieddatalibrary.String("NONE"),
				Disposition:     unifieddatalibrary.String("EVACUATE WOUNDED"),
				DispositionType: unifieddatalibrary.String("EVACUATE"),
				Etiology: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoEtiology{{
					BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the etiology info."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("BURN"),
				}},
				EvacType: unifieddatalibrary.String("GROUND"),
				Gender:   unifieddatalibrary.String("MALE"),
				HealthState: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoHealthState{{
					HealthStateCode: unifieddatalibrary.String("BLUE"),
					MedConfFactor:   unifieddatalibrary.Int(1),
					Time:            unifieddatalibrary.Time(time.Now()),
					Type:            unifieddatalibrary.String("COGNITIVE"),
				}},
				Injury: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoInjury{{
					BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the patient's injury."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("ABRASION"),
				}},
				Last4Ssn: unifieddatalibrary.String("1234"),
				Medication: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoMedication{{
					AdminRoute: unifieddatalibrary.String("ORAL"),
					BodyPart:   unifieddatalibrary.String("ARM LEFT BACK"),
					Comments:   unifieddatalibrary.String("Comments on the patient's medication information."),
					Dose:       unifieddatalibrary.String("800mg"),
					Time:       unifieddatalibrary.Time(time.Now()),
					Type:       unifieddatalibrary.String("TYLENOL"),
				}},
				Name:            unifieddatalibrary.String("John Smith"),
				Nationality:     unifieddatalibrary.String("US"),
				OccSpeciality:   unifieddatalibrary.String("Healthcare"),
				PatientIdentity: unifieddatalibrary.String("FRIEND CIVILIAN"),
				PatientStatus:   unifieddatalibrary.String("US CIVILIAN"),
				PayGrade:        unifieddatalibrary.String("CIVILIAN"),
				Priority:        unifieddatalibrary.String("ROUTINE"),
				ReportGen:       unifieddatalibrary.String("DEVICE"),
				ReportTime:      unifieddatalibrary.Time(time.Now()),
				Service:         unifieddatalibrary.String("CIV"),
				SpecMedEquip:    []string{"OXYGEN", "HOIST"},
				Treatment: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoTreatment{{
					BodyPart: unifieddatalibrary.String("CHEST"),
					Comments: unifieddatalibrary.String("Comments on the treatment info."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("BREATHING CHEST TUBE"),
				}},
				VitalSignData: []unifieddatalibrary.EvacNewBulkParamsBodyCasualtyInfoVitalSignData{{
					MedConfFactor: unifieddatalibrary.Int(1),
					Time:          unifieddatalibrary.Time(time.Now()),
					VitalSign:     unifieddatalibrary.String("HEART RATE"),
					VitalSign1:    unifieddatalibrary.Float(120),
					VitalSign2:    unifieddatalibrary.Float(80),
				}},
			}},
			Ce:        unifieddatalibrary.Float(10.1234),
			CntctFreq: unifieddatalibrary.Float(3.11),
			Comments:  unifieddatalibrary.String("Comments concerning mission"),
			EnemyData: []unifieddatalibrary.EvacNewBulkParamsBodyEnemyData{{
				DirToEnemy:        unifieddatalibrary.String("NORTH"),
				FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
				HlzRemarks:        unifieddatalibrary.String("Remarks about hot landing zone."),
				HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
			}},
			IDWeatherReport:   unifieddatalibrary.String("WeatherReport-ID"),
			Le:                unifieddatalibrary.Float(5.1234),
			MedevacID:         unifieddatalibrary.String("MedEvac-ID"),
			MedicReq:          unifieddatalibrary.Bool(true),
			MissionType:       unifieddatalibrary.String("GROUND"),
			NumAmbulatory:     unifieddatalibrary.Int(5),
			NumCasualties:     unifieddatalibrary.Int(5),
			NumKia:            unifieddatalibrary.Int(0),
			NumLitter:         unifieddatalibrary.Int(0),
			NumWia:            unifieddatalibrary.Int(3),
			ObstaclesRemarks:  unifieddatalibrary.String("N/A"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PickupAlt:         unifieddatalibrary.Float(30.1234),
			PickupTime:        unifieddatalibrary.Time(time.Now()),
			ReqCallSign:       unifieddatalibrary.String("Bravo"),
			ReqNum:            unifieddatalibrary.String("MED.1.223908"),
			Terrain:           unifieddatalibrary.String("ROCKY"),
			TerrainRemarks:    unifieddatalibrary.String("N/A"),
			ZoneContrCallSign: unifieddatalibrary.String("Tango"),
			ZoneHot:           unifieddatalibrary.Bool(false),
			ZoneMarking:       unifieddatalibrary.String("ILLUMINATION"),
			ZoneMarkingColor:  unifieddatalibrary.String("RED"),
			ZoneName:          unifieddatalibrary.String("example-zone"),
			ZoneSecurity:      unifieddatalibrary.String("NO ENEMY"),
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

func TestEvacQueryHelp(t *testing.T) {
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
	_, err := client.Evac.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvacUnvalidatedPublish(t *testing.T) {
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
	err := client.Evac.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EvacUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EvacUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			PickupLat:             75.1234,
			PickupLon:             175.1234,
			ReqTime:               time.Now(),
			Source:                "Bluestaq",
			Type:                  "REQUEST",
			ID:                    unifieddatalibrary.String("MEDEVACEVENT-ID"),
			CasualtyInfo: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfo{{
				Age: unifieddatalibrary.Int(35),
				Allergy: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoAllergy{{
					Comments: unifieddatalibrary.String("Comments on the patient's allergies."),
					Type:     unifieddatalibrary.String("PENICILLIN"),
				}},
				BloodType:       unifieddatalibrary.String("O NEG"),
				BodyPart:        unifieddatalibrary.String("FACE"),
				BurialLocation:  []float64{-33.123, 150.33, 0.24},
				CallSign:        unifieddatalibrary.String("SHARK"),
				CareProviderUrn: unifieddatalibrary.String("CARE_PROVIDER-1"),
				CasualtyKey:     unifieddatalibrary.String("casualty-007"),
				CasualtyType:    unifieddatalibrary.String("DENTAL"),
				CollectionPoint: []float64{12.44, 122.55, 0.98},
				Comments:        unifieddatalibrary.String("Comments relating to this casualty info."),
				Condition: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoCondition{{
					BodyPart: unifieddatalibrary.String("ANKLE LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the patient's condition."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("ACTIVITY LOW"),
				}},
				ContamType:      unifieddatalibrary.String("NONE"),
				Disposition:     unifieddatalibrary.String("EVACUATE WOUNDED"),
				DispositionType: unifieddatalibrary.String("EVACUATE"),
				Etiology: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoEtiology{{
					BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the etiology info."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("BURN"),
				}},
				EvacType: unifieddatalibrary.String("GROUND"),
				Gender:   unifieddatalibrary.String("MALE"),
				HealthState: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoHealthState{{
					HealthStateCode: unifieddatalibrary.String("BLUE"),
					MedConfFactor:   unifieddatalibrary.Int(1),
					Time:            unifieddatalibrary.Time(time.Now()),
					Type:            unifieddatalibrary.String("COGNITIVE"),
				}},
				Injury: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoInjury{{
					BodyPart: unifieddatalibrary.String("ARM LEFT FRONT"),
					Comments: unifieddatalibrary.String("Comments on the patient's injury."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("ABRASION"),
				}},
				Last4Ssn: unifieddatalibrary.String("1234"),
				Medication: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoMedication{{
					AdminRoute: unifieddatalibrary.String("ORAL"),
					BodyPart:   unifieddatalibrary.String("ARM LEFT BACK"),
					Comments:   unifieddatalibrary.String("Comments on the patient's medication information."),
					Dose:       unifieddatalibrary.String("800mg"),
					Time:       unifieddatalibrary.Time(time.Now()),
					Type:       unifieddatalibrary.String("TYLENOL"),
				}},
				Name:            unifieddatalibrary.String("John Smith"),
				Nationality:     unifieddatalibrary.String("US"),
				OccSpeciality:   unifieddatalibrary.String("Healthcare"),
				PatientIdentity: unifieddatalibrary.String("FRIEND CIVILIAN"),
				PatientStatus:   unifieddatalibrary.String("US CIVILIAN"),
				PayGrade:        unifieddatalibrary.String("CIVILIAN"),
				Priority:        unifieddatalibrary.String("ROUTINE"),
				ReportGen:       unifieddatalibrary.String("DEVICE"),
				ReportTime:      unifieddatalibrary.Time(time.Now()),
				Service:         unifieddatalibrary.String("CIV"),
				SpecMedEquip:    []string{"OXYGEN", "HOIST"},
				Treatment: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoTreatment{{
					BodyPart: unifieddatalibrary.String("CHEST"),
					Comments: unifieddatalibrary.String("Comments on the treatment info."),
					Time:     unifieddatalibrary.Time(time.Now()),
					Type:     unifieddatalibrary.String("BREATHING CHEST TUBE"),
				}},
				VitalSignData: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyCasualtyInfoVitalSignData{{
					MedConfFactor: unifieddatalibrary.Int(1),
					Time:          unifieddatalibrary.Time(time.Now()),
					VitalSign:     unifieddatalibrary.String("HEART RATE"),
					VitalSign1:    unifieddatalibrary.Float(120),
					VitalSign2:    unifieddatalibrary.Float(80),
				}},
			}},
			Ce:        unifieddatalibrary.Float(10.1234),
			CntctFreq: unifieddatalibrary.Float(3.11),
			Comments:  unifieddatalibrary.String("Comments concerning mission"),
			EnemyData: []unifieddatalibrary.EvacUnvalidatedPublishParamsBodyEnemyData{{
				DirToEnemy:        unifieddatalibrary.String("NORTH"),
				FriendliesRemarks: unifieddatalibrary.String("Comments from friendlies."),
				HlzRemarks:        unifieddatalibrary.String("Remarks about hot landing zone."),
				HostileFireType:   unifieddatalibrary.String("SMALL ARMS"),
			}},
			IDWeatherReport:   unifieddatalibrary.String("WeatherReport-ID"),
			Le:                unifieddatalibrary.Float(5.1234),
			MedevacID:         unifieddatalibrary.String("MedEvac-ID"),
			MedicReq:          unifieddatalibrary.Bool(true),
			MissionType:       unifieddatalibrary.String("GROUND"),
			NumAmbulatory:     unifieddatalibrary.Int(5),
			NumCasualties:     unifieddatalibrary.Int(5),
			NumKia:            unifieddatalibrary.Int(0),
			NumLitter:         unifieddatalibrary.Int(0),
			NumWia:            unifieddatalibrary.Int(3),
			ObstaclesRemarks:  unifieddatalibrary.String("N/A"),
			Origin:            unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PickupAlt:         unifieddatalibrary.Float(30.1234),
			PickupTime:        unifieddatalibrary.Time(time.Now()),
			ReqCallSign:       unifieddatalibrary.String("Bravo"),
			ReqNum:            unifieddatalibrary.String("MED.1.223908"),
			Terrain:           unifieddatalibrary.String("ROCKY"),
			TerrainRemarks:    unifieddatalibrary.String("N/A"),
			ZoneContrCallSign: unifieddatalibrary.String("Tango"),
			ZoneHot:           unifieddatalibrary.Bool(false),
			ZoneMarking:       unifieddatalibrary.String("ILLUMINATION"),
			ZoneMarkingColor:  unifieddatalibrary.String("RED"),
			ZoneName:          unifieddatalibrary.String("example-zone"),
			ZoneSecurity:      unifieddatalibrary.String("NO ENEMY"),
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
