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

func TestEffectResponseNewWithOptionalParams(t *testing.T) {
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
	err := client.EffectResponses.New(context.TODO(), unifieddatalibrary.EffectResponseNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EffectResponseNewParamsDataModeTest,
		Source:                "Bluestaq",
		Type:                  "COA",
		ID:                    unifieddatalibrary.String("EFFECTRESPONSE-ID"),
		ActionsList: []unifieddatalibrary.EffectResponseNewParamsActionsList{{
			ActionActorSrcID:   unifieddatalibrary.String("ACTIONACTORSRC-ID"),
			ActionActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
			ActionEndTime:      unifieddatalibrary.Time(time.Now()),
			ActionID:           unifieddatalibrary.String("ACTION-ID"),
			ActionMetrics: []unifieddatalibrary.EffectResponseNewParamsActionsListActionMetric{{
				DomainValue:   unifieddatalibrary.Float(10.1),
				MetricType:    unifieddatalibrary.String("GoalAchievement"),
				Provenance:    unifieddatalibrary.String("Example metric"),
				RelativeValue: unifieddatalibrary.Float(10.1),
			}},
			ActionStartTime:    unifieddatalibrary.Time(time.Now()),
			ActorInterceptAlt:  unifieddatalibrary.Float(1.1),
			ActorInterceptLat:  unifieddatalibrary.Float(45.1),
			ActorInterceptLon:  unifieddatalibrary.Float(180.1),
			Effector:           unifieddatalibrary.String("SENSOR1"),
			Summary:            unifieddatalibrary.String("Example summary"),
			TargetSrcID:        unifieddatalibrary.String("TARGETSRC-ID"),
			TargetSrcType:      unifieddatalibrary.String("POI"),
			TotEndTime:         unifieddatalibrary.Time(time.Now()),
			TotStartTime:       unifieddatalibrary.Time(time.Now()),
			WeaponInterceptAlt: unifieddatalibrary.Float(1.1),
			WeaponInterceptLat: unifieddatalibrary.Float(45.1),
			WeaponInterceptLon: unifieddatalibrary.Float(180.1),
		}},
		ActorSrcID:   unifieddatalibrary.String("RC-ID"),
		ActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
		CoaMetrics: []unifieddatalibrary.EffectResponseNewParamsCoaMetric{{
			DomainValue:   unifieddatalibrary.Float(10.1),
			MetricType:    unifieddatalibrary.String("GoalAchievement"),
			Provenance:    unifieddatalibrary.String("Example metric"),
			RelativeValue: unifieddatalibrary.Float(10.1),
		}},
		CollateralDamageEst: unifieddatalibrary.Float(0.5),
		DecisionDeadline:    unifieddatalibrary.Time(time.Now()),
		ExternalActions:     []string{"ACTION1", "ACTION2"},
		ExternalRequestID:   unifieddatalibrary.String("EXTERNALREQUEST-ID"),
		IDEffectRequest:     unifieddatalibrary.String("EFFECTREQUEST-ID"),
		MunitionID:          unifieddatalibrary.String("MUNITION-ID"),
		MunitionType:        unifieddatalibrary.String("Dummy"),
		Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		ProbabilityOfKill:   unifieddatalibrary.Float(0.7),
		RedTargetSrcID:      unifieddatalibrary.String("REDTARGETSRC-ID"),
		RedTargetSrcType:    unifieddatalibrary.String("POI"),
		RedTimeToOverhead:   unifieddatalibrary.Time(time.Now()),
		ShotsRequired:       unifieddatalibrary.Int(10),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEffectResponseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectResponses.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EffectResponseGetParams{
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

func TestEffectResponseListWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectResponses.List(context.TODO(), unifieddatalibrary.EffectResponseListParams{
		CreatedAt:   time.Now(),
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

func TestEffectResponseCountWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectResponses.Count(context.TODO(), unifieddatalibrary.EffectResponseCountParams{
		CreatedAt:   time.Now(),
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

func TestEffectResponseNewBulk(t *testing.T) {
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
	err := client.EffectResponses.NewBulk(context.TODO(), unifieddatalibrary.EffectResponseNewBulkParams{
		Body: []unifieddatalibrary.EffectResponseNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Type:                  "COA",
			ID:                    unifieddatalibrary.String("EFFECTRESPONSE-ID"),
			ActionsList: []unifieddatalibrary.EffectResponseNewBulkParamsBodyActionsList{{
				ActionActorSrcID:   unifieddatalibrary.String("ACTIONACTORSRC-ID"),
				ActionActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
				ActionEndTime:      unifieddatalibrary.Time(time.Now()),
				ActionID:           unifieddatalibrary.String("ACTION-ID"),
				ActionMetrics: []unifieddatalibrary.EffectResponseNewBulkParamsBodyActionsListActionMetric{{
					DomainValue:   unifieddatalibrary.Float(10.1),
					MetricType:    unifieddatalibrary.String("GoalAchievement"),
					Provenance:    unifieddatalibrary.String("Example metric"),
					RelativeValue: unifieddatalibrary.Float(10.1),
				}},
				ActionStartTime:    unifieddatalibrary.Time(time.Now()),
				ActorInterceptAlt:  unifieddatalibrary.Float(1.1),
				ActorInterceptLat:  unifieddatalibrary.Float(45.1),
				ActorInterceptLon:  unifieddatalibrary.Float(180.1),
				Effector:           unifieddatalibrary.String("SENSOR1"),
				Summary:            unifieddatalibrary.String("Example summary"),
				TargetSrcID:        unifieddatalibrary.String("TARGETSRC-ID"),
				TargetSrcType:      unifieddatalibrary.String("POI"),
				TotEndTime:         unifieddatalibrary.Time(time.Now()),
				TotStartTime:       unifieddatalibrary.Time(time.Now()),
				WeaponInterceptAlt: unifieddatalibrary.Float(1.1),
				WeaponInterceptLat: unifieddatalibrary.Float(45.1),
				WeaponInterceptLon: unifieddatalibrary.Float(180.1),
			}},
			ActorSrcID:   unifieddatalibrary.String("RC-ID"),
			ActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
			CoaMetrics: []unifieddatalibrary.EffectResponseNewBulkParamsBodyCoaMetric{{
				DomainValue:   unifieddatalibrary.Float(10.1),
				MetricType:    unifieddatalibrary.String("GoalAchievement"),
				Provenance:    unifieddatalibrary.String("Example metric"),
				RelativeValue: unifieddatalibrary.Float(10.1),
			}},
			CollateralDamageEst: unifieddatalibrary.Float(0.5),
			DecisionDeadline:    unifieddatalibrary.Time(time.Now()),
			ExternalActions:     []string{"ACTION1", "ACTION2"},
			ExternalRequestID:   unifieddatalibrary.String("EXTERNALREQUEST-ID"),
			IDEffectRequest:     unifieddatalibrary.String("EFFECTREQUEST-ID"),
			MunitionID:          unifieddatalibrary.String("MUNITION-ID"),
			MunitionType:        unifieddatalibrary.String("Dummy"),
			Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ProbabilityOfKill:   unifieddatalibrary.Float(0.7),
			RedTargetSrcID:      unifieddatalibrary.String("REDTARGETSRC-ID"),
			RedTargetSrcType:    unifieddatalibrary.String("POI"),
			RedTimeToOverhead:   unifieddatalibrary.Time(time.Now()),
			ShotsRequired:       unifieddatalibrary.Int(10),
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

func TestEffectResponseQueryHelp(t *testing.T) {
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
	_, err := client.EffectResponses.QueryHelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEffectResponseTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.EffectResponses.Tuple(context.TODO(), unifieddatalibrary.EffectResponseTupleParams{
		Columns:     "columns",
		CreatedAt:   time.Now(),
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

func TestEffectResponseUnvalidatedPublish(t *testing.T) {
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
	err := client.EffectResponses.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EffectResponseUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EffectResponseUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Type:                  "COA",
			ID:                    unifieddatalibrary.String("EFFECTRESPONSE-ID"),
			ActionsList: []unifieddatalibrary.EffectResponseUnvalidatedPublishParamsBodyActionsList{{
				ActionActorSrcID:   unifieddatalibrary.String("ACTIONACTORSRC-ID"),
				ActionActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
				ActionEndTime:      unifieddatalibrary.Time(time.Now()),
				ActionID:           unifieddatalibrary.String("ACTION-ID"),
				ActionMetrics: []unifieddatalibrary.EffectResponseUnvalidatedPublishParamsBodyActionsListActionMetric{{
					DomainValue:   unifieddatalibrary.Float(10.1),
					MetricType:    unifieddatalibrary.String("GoalAchievement"),
					Provenance:    unifieddatalibrary.String("Example metric"),
					RelativeValue: unifieddatalibrary.Float(10.1),
				}},
				ActionStartTime:    unifieddatalibrary.Time(time.Now()),
				ActorInterceptAlt:  unifieddatalibrary.Float(1.1),
				ActorInterceptLat:  unifieddatalibrary.Float(45.1),
				ActorInterceptLon:  unifieddatalibrary.Float(180.1),
				Effector:           unifieddatalibrary.String("SENSOR1"),
				Summary:            unifieddatalibrary.String("Example summary"),
				TargetSrcID:        unifieddatalibrary.String("TARGETSRC-ID"),
				TargetSrcType:      unifieddatalibrary.String("POI"),
				TotEndTime:         unifieddatalibrary.Time(time.Now()),
				TotStartTime:       unifieddatalibrary.Time(time.Now()),
				WeaponInterceptAlt: unifieddatalibrary.Float(1.1),
				WeaponInterceptLat: unifieddatalibrary.Float(45.1),
				WeaponInterceptLon: unifieddatalibrary.Float(180.1),
			}},
			ActorSrcID:   unifieddatalibrary.String("RC-ID"),
			ActorSrcType: unifieddatalibrary.String("AIRCRAFT"),
			CoaMetrics: []unifieddatalibrary.EffectResponseUnvalidatedPublishParamsBodyCoaMetric{{
				DomainValue:   unifieddatalibrary.Float(10.1),
				MetricType:    unifieddatalibrary.String("GoalAchievement"),
				Provenance:    unifieddatalibrary.String("Example metric"),
				RelativeValue: unifieddatalibrary.Float(10.1),
			}},
			CollateralDamageEst: unifieddatalibrary.Float(0.5),
			DecisionDeadline:    unifieddatalibrary.Time(time.Now()),
			ExternalActions:     []string{"ACTION1", "ACTION2"},
			ExternalRequestID:   unifieddatalibrary.String("EXTERNALREQUEST-ID"),
			IDEffectRequest:     unifieddatalibrary.String("EFFECTREQUEST-ID"),
			MunitionID:          unifieddatalibrary.String("MUNITION-ID"),
			MunitionType:        unifieddatalibrary.String("Dummy"),
			Origin:              unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			ProbabilityOfKill:   unifieddatalibrary.Float(0.7),
			RedTargetSrcID:      unifieddatalibrary.String("REDTARGETSRC-ID"),
			RedTargetSrcType:    unifieddatalibrary.String("POI"),
			RedTimeToOverhead:   unifieddatalibrary.Time(time.Now()),
			ShotsRequired:       unifieddatalibrary.Int(10),
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
