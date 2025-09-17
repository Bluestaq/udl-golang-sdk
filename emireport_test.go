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

func TestEmireportNewWithOptionalParams(t *testing.T) {
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
	err := client.Emireport.New(context.TODO(), unifieddatalibrary.EmireportNewParams{
		ClassificationMarking:   "U",
		DataMode:                unifieddatalibrary.EmireportNewParamsDataModeTest,
		Isr:                     true,
		ReportID:                "REPORT-ID",
		ReportTime:              time.Now(),
		ReportType:              "SATCOM",
		RequestAssist:           true,
		Source:                  "Bluestaq",
		StartTime:               time.Now(),
		ID:                      unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		ActionsTaken:            unifieddatalibrary.String("verified connections, cables and antenna pointing angles"),
		AffActivity:             unifieddatalibrary.String("UPLINK"),
		Alt:                     unifieddatalibrary.Float(1750),
		Aor:                     unifieddatalibrary.String("NORTHCOM"),
		Band:                    unifieddatalibrary.String("SHF"),
		BeamPattern:             unifieddatalibrary.String("MAIN LOBE"),
		Channel:                 unifieddatalibrary.String("10C-10CU"),
		ChanPirate:              unifieddatalibrary.Bool(false),
		Description:             unifieddatalibrary.String("Interference on channel"),
		DneImpact:               unifieddatalibrary.String("Text description of the duration, nature and extent (DNE) of the impact."),
		EmiType:                 unifieddatalibrary.String("BARRAGE"),
		EndTime:                 unifieddatalibrary.Time(time.Now()),
		Frequency:               unifieddatalibrary.Float(1575.42),
		GeoLocErrEllp:           []float64{1300, 700, 35},
		GpsEncrypted:            unifieddatalibrary.Bool(false),
		GpsFreq:                 unifieddatalibrary.String("L1"),
		HighAffectedFrequency:   unifieddatalibrary.Float(1725),
		Intercept:               unifieddatalibrary.Bool(false),
		InterceptLang:           unifieddatalibrary.String("ENGLISH"),
		InterceptType:           unifieddatalibrary.String("VOICE"),
		IntSrcAmplitude:         unifieddatalibrary.Float(0.275),
		IntSrcBandwidth:         unifieddatalibrary.Float(30),
		IntSrcCentFreq:          unifieddatalibrary.Float(485.7),
		IntSrcEncrypted:         unifieddatalibrary.Bool(false),
		IntSrcModulation:        unifieddatalibrary.String("FSK"),
		IsrCollectionImpact:     unifieddatalibrary.Bool(false),
		KillBox:                 unifieddatalibrary.String("7F9SW"),
		Lat:                     unifieddatalibrary.Float(38.7375),
		Link:                    unifieddatalibrary.String("SPOT-21"),
		Lon:                     unifieddatalibrary.Float(-104.7889),
		MilGrid:                 unifieddatalibrary.String("4QFJ12345678"),
		Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:            unifieddatalibrary.String("25724"),
		Persistence:             unifieddatalibrary.String("CONTINUOUS"),
		Platform:                unifieddatalibrary.String("CVN-78"),
		RcvrDemod:               unifieddatalibrary.String("FSK"),
		RcvrGain:                unifieddatalibrary.Float(23.7),
		RcvrLocation:            unifieddatalibrary.String("FORT CARSON GARAGE"),
		RcvrType:                unifieddatalibrary.String("OMNI"),
		RespService:             unifieddatalibrary.String("ARMY"),
		SatcomPriority:          unifieddatalibrary.String("HIGH"),
		SatDownlinkFrequency:    unifieddatalibrary.Float(47432.5),
		SatDownlinkPolarization: unifieddatalibrary.String("V"),
		SatName:                 unifieddatalibrary.String("MILSTAR DFS-3"),
		SatNo:                   unifieddatalibrary.Int(25724),
		SatTransponderID:        unifieddatalibrary.String("36097-8433-10C"),
		SatUplinkFrequency:      unifieddatalibrary.Float(44532.1),
		SatUplinkPolarization:   unifieddatalibrary.String("H"),
		Status:                  unifieddatalibrary.String("INITIAL"),
		SupportedIsrRole:        unifieddatalibrary.String("IMAGERY"),
		System:                  unifieddatalibrary.String("RADIO"),
		Tags:                    []string{"TAG1", "TAG2"},
		TransactionID:           unifieddatalibrary.String("TRANSACTION-ID"),
		VictimAltCountry:        unifieddatalibrary.String("US"),
		VictimCountryCode:       unifieddatalibrary.String("US"),
		VictimFuncImpacts:       unifieddatalibrary.String("C2"),
		VictimPocMail:           unifieddatalibrary.String("bob@jammer.com"),
		VictimPocName:           unifieddatalibrary.String("Robert Smith"),
		VictimPocPhone:          unifieddatalibrary.String("7198675309"),
		VictimPocUnit:           unifieddatalibrary.String("4th Engineering Battalion"),
		VictimReaction:          unifieddatalibrary.String("TROUBLESHOOT"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmireportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Emireport.List(context.TODO(), unifieddatalibrary.EmireportListParams{
		ReportTime:  time.Now(),
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

func TestEmireportCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Emireport.Count(context.TODO(), unifieddatalibrary.EmireportCountParams{
		ReportTime:  time.Now(),
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

func TestEmireportNewBulk(t *testing.T) {
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
	err := client.Emireport.NewBulk(context.TODO(), unifieddatalibrary.EmireportNewBulkParams{
		Body: []unifieddatalibrary.EmireportNewBulkParamsBody{{
			ClassificationMarking:   "U",
			DataMode:                "TEST",
			Isr:                     true,
			ReportID:                "REPORT-ID",
			ReportTime:              time.Now(),
			ReportType:              "SATCOM",
			RequestAssist:           true,
			Source:                  "Bluestaq",
			StartTime:               time.Now(),
			ID:                      unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActionsTaken:            unifieddatalibrary.String("verified connections, cables and antenna pointing angles"),
			AffActivity:             unifieddatalibrary.String("UPLINK"),
			Alt:                     unifieddatalibrary.Float(1750),
			Aor:                     unifieddatalibrary.String("NORTHCOM"),
			Band:                    unifieddatalibrary.String("SHF"),
			BeamPattern:             unifieddatalibrary.String("MAIN LOBE"),
			Channel:                 unifieddatalibrary.String("10C-10CU"),
			ChanPirate:              unifieddatalibrary.Bool(false),
			Description:             unifieddatalibrary.String("Interference on channel"),
			DneImpact:               unifieddatalibrary.String("Text description of the duration, nature and extent (DNE) of the impact."),
			EmiType:                 unifieddatalibrary.String("BARRAGE"),
			EndTime:                 unifieddatalibrary.Time(time.Now()),
			Frequency:               unifieddatalibrary.Float(1575.42),
			GeoLocErrEllp:           []float64{1300, 700, 35},
			GpsEncrypted:            unifieddatalibrary.Bool(false),
			GpsFreq:                 unifieddatalibrary.String("L1"),
			HighAffectedFrequency:   unifieddatalibrary.Float(1725),
			Intercept:               unifieddatalibrary.Bool(false),
			InterceptLang:           unifieddatalibrary.String("ENGLISH"),
			InterceptType:           unifieddatalibrary.String("VOICE"),
			IntSrcAmplitude:         unifieddatalibrary.Float(0.275),
			IntSrcBandwidth:         unifieddatalibrary.Float(30),
			IntSrcCentFreq:          unifieddatalibrary.Float(485.7),
			IntSrcEncrypted:         unifieddatalibrary.Bool(false),
			IntSrcModulation:        unifieddatalibrary.String("FSK"),
			IsrCollectionImpact:     unifieddatalibrary.Bool(false),
			KillBox:                 unifieddatalibrary.String("7F9SW"),
			Lat:                     unifieddatalibrary.Float(38.7375),
			Link:                    unifieddatalibrary.String("SPOT-21"),
			Lon:                     unifieddatalibrary.Float(-104.7889),
			MilGrid:                 unifieddatalibrary.String("4QFJ12345678"),
			Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:            unifieddatalibrary.String("25724"),
			Persistence:             unifieddatalibrary.String("CONTINUOUS"),
			Platform:                unifieddatalibrary.String("CVN-78"),
			RcvrDemod:               unifieddatalibrary.String("FSK"),
			RcvrGain:                unifieddatalibrary.Float(23.7),
			RcvrLocation:            unifieddatalibrary.String("FORT CARSON GARAGE"),
			RcvrType:                unifieddatalibrary.String("OMNI"),
			RespService:             unifieddatalibrary.String("ARMY"),
			SatcomPriority:          unifieddatalibrary.String("HIGH"),
			SatDownlinkFrequency:    unifieddatalibrary.Float(47432.5),
			SatDownlinkPolarization: unifieddatalibrary.String("V"),
			SatName:                 unifieddatalibrary.String("MILSTAR DFS-3"),
			SatNo:                   unifieddatalibrary.Int(25724),
			SatTransponderID:        unifieddatalibrary.String("36097-8433-10C"),
			SatUplinkFrequency:      unifieddatalibrary.Float(44532.1),
			SatUplinkPolarization:   unifieddatalibrary.String("H"),
			Status:                  unifieddatalibrary.String("INITIAL"),
			SupportedIsrRole:        unifieddatalibrary.String("IMAGERY"),
			System:                  unifieddatalibrary.String("RADIO"),
			Tags:                    []string{"TAG1", "TAG2"},
			TransactionID:           unifieddatalibrary.String("TRANSACTION-ID"),
			VictimAltCountry:        unifieddatalibrary.String("US"),
			VictimCountryCode:       unifieddatalibrary.String("US"),
			VictimFuncImpacts:       unifieddatalibrary.String("C2"),
			VictimPocMail:           unifieddatalibrary.String("bob@jammer.com"),
			VictimPocName:           unifieddatalibrary.String("Robert Smith"),
			VictimPocPhone:          unifieddatalibrary.String("7198675309"),
			VictimPocUnit:           unifieddatalibrary.String("4th Engineering Battalion"),
			VictimReaction:          unifieddatalibrary.String("TROUBLESHOOT"),
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

func TestEmireportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Emireport.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EmireportGetParams{
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

func TestEmireportQueryhelp(t *testing.T) {
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
	_, err := client.Emireport.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmireportTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Emireport.Tuple(context.TODO(), unifieddatalibrary.EmireportTupleParams{
		Columns:     "columns",
		ReportTime:  time.Now(),
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

func TestEmireportUnvalidatedPublish(t *testing.T) {
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
	err := client.Emireport.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EmireportUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EmireportUnvalidatedPublishParamsBody{{
			ClassificationMarking:   "U",
			DataMode:                "TEST",
			Isr:                     true,
			ReportID:                "REPORT-ID",
			ReportTime:              time.Now(),
			ReportType:              "SATCOM",
			RequestAssist:           true,
			Source:                  "Bluestaq",
			StartTime:               time.Now(),
			ID:                      unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			ActionsTaken:            unifieddatalibrary.String("verified connections, cables and antenna pointing angles"),
			AffActivity:             unifieddatalibrary.String("UPLINK"),
			Alt:                     unifieddatalibrary.Float(1750),
			Aor:                     unifieddatalibrary.String("NORTHCOM"),
			Band:                    unifieddatalibrary.String("SHF"),
			BeamPattern:             unifieddatalibrary.String("MAIN LOBE"),
			Channel:                 unifieddatalibrary.String("10C-10CU"),
			ChanPirate:              unifieddatalibrary.Bool(false),
			Description:             unifieddatalibrary.String("Interference on channel"),
			DneImpact:               unifieddatalibrary.String("Text description of the duration, nature and extent (DNE) of the impact."),
			EmiType:                 unifieddatalibrary.String("BARRAGE"),
			EndTime:                 unifieddatalibrary.Time(time.Now()),
			Frequency:               unifieddatalibrary.Float(1575.42),
			GeoLocErrEllp:           []float64{1300, 700, 35},
			GpsEncrypted:            unifieddatalibrary.Bool(false),
			GpsFreq:                 unifieddatalibrary.String("L1"),
			HighAffectedFrequency:   unifieddatalibrary.Float(1725),
			Intercept:               unifieddatalibrary.Bool(false),
			InterceptLang:           unifieddatalibrary.String("ENGLISH"),
			InterceptType:           unifieddatalibrary.String("VOICE"),
			IntSrcAmplitude:         unifieddatalibrary.Float(0.275),
			IntSrcBandwidth:         unifieddatalibrary.Float(30),
			IntSrcCentFreq:          unifieddatalibrary.Float(485.7),
			IntSrcEncrypted:         unifieddatalibrary.Bool(false),
			IntSrcModulation:        unifieddatalibrary.String("FSK"),
			IsrCollectionImpact:     unifieddatalibrary.Bool(false),
			KillBox:                 unifieddatalibrary.String("7F9SW"),
			Lat:                     unifieddatalibrary.Float(38.7375),
			Link:                    unifieddatalibrary.String("SPOT-21"),
			Lon:                     unifieddatalibrary.Float(-104.7889),
			MilGrid:                 unifieddatalibrary.String("4QFJ12345678"),
			Origin:                  unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:            unifieddatalibrary.String("25724"),
			Persistence:             unifieddatalibrary.String("CONTINUOUS"),
			Platform:                unifieddatalibrary.String("CVN-78"),
			RcvrDemod:               unifieddatalibrary.String("FSK"),
			RcvrGain:                unifieddatalibrary.Float(23.7),
			RcvrLocation:            unifieddatalibrary.String("FORT CARSON GARAGE"),
			RcvrType:                unifieddatalibrary.String("OMNI"),
			RespService:             unifieddatalibrary.String("ARMY"),
			SatcomPriority:          unifieddatalibrary.String("HIGH"),
			SatDownlinkFrequency:    unifieddatalibrary.Float(47432.5),
			SatDownlinkPolarization: unifieddatalibrary.String("V"),
			SatName:                 unifieddatalibrary.String("MILSTAR DFS-3"),
			SatNo:                   unifieddatalibrary.Int(25724),
			SatTransponderID:        unifieddatalibrary.String("36097-8433-10C"),
			SatUplinkFrequency:      unifieddatalibrary.Float(44532.1),
			SatUplinkPolarization:   unifieddatalibrary.String("H"),
			Status:                  unifieddatalibrary.String("INITIAL"),
			SupportedIsrRole:        unifieddatalibrary.String("IMAGERY"),
			System:                  unifieddatalibrary.String("RADIO"),
			Tags:                    []string{"TAG1", "TAG2"},
			TransactionID:           unifieddatalibrary.String("TRANSACTION-ID"),
			VictimAltCountry:        unifieddatalibrary.String("US"),
			VictimCountryCode:       unifieddatalibrary.String("US"),
			VictimFuncImpacts:       unifieddatalibrary.String("C2"),
			VictimPocMail:           unifieddatalibrary.String("bob@jammer.com"),
			VictimPocName:           unifieddatalibrary.String("Robert Smith"),
			VictimPocPhone:          unifieddatalibrary.String("7198675309"),
			VictimPocUnit:           unifieddatalibrary.String("4th Engineering Battalion"),
			VictimReaction:          unifieddatalibrary.String("TROUBLESHOOT"),
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
