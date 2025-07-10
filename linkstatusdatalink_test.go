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

func TestLinkStatusDatalinkNewWithOptionalParams(t *testing.T) {
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
	err := client.LinkStatus.Datalink.New(context.TODO(), unifieddatalibrary.LinkStatusDatalinkNewParams{
		DatalinkIngest: unifieddatalibrary.DatalinkIngestParam{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.DatalinkIngestDataModeTest,
			OpExName:              "DESERT WIND",
			Originator:            "USCENTCOM",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AckInstUnits:          []string{"AOC EXT 2345", "317 AW"},
			AckReq:                unifieddatalibrary.Bool(true),
			AltDiff:               unifieddatalibrary.Int(20),
			CanxID:                unifieddatalibrary.String("ABSTAT"),
			CanxOriginator:        unifieddatalibrary.String("505 AOC"),
			CanxSerialNum:         unifieddatalibrary.String("ABC1234567"),
			CanxSiCs:              []string{"RDU", "X234BS"},
			CanxSpecialNotation:   unifieddatalibrary.String("PASEP"),
			CanxTs:                unifieddatalibrary.Time(time.Now()),
			ClassReasons:          []string{"15C", "15D"},
			ClassSource:           unifieddatalibrary.String("USJFCOM EXORD SOLID WASTE 98"),
			ConsecDecorr:          unifieddatalibrary.Int(3),
			CourseDiff:            unifieddatalibrary.Int(60),
			DecExemptCodes:        []string{"X1", "X2"},
			DecInstDates:          []string{"AT EXERCISE ENDEX", "DATE:25NOV1997"},
			DecorrWinMult:         unifieddatalibrary.Float(1.7),
			GeoDatum:              unifieddatalibrary.String("EUR-T"),
			JreCallSign:           unifieddatalibrary.String("CHARLIE ONE"),
			JreDetails:            unifieddatalibrary.String("JRE details"),
			JrePriAdd:             unifieddatalibrary.Int(71777),
			JreSecAdd:             unifieddatalibrary.Int(77771),
			JreUnitDes:            unifieddatalibrary.String("CVN-72"),
			MaxGeoPosQual:         unifieddatalibrary.Int(12),
			MaxTrackQual:          unifieddatalibrary.Int(12),
			MgmtCode:              unifieddatalibrary.String("VICTOR"),
			MgmtCodeMeaning:       unifieddatalibrary.String("ORBIT AT POINT BRAVO"),
			MinGeoPosQual:         unifieddatalibrary.Int(3),
			MinTrackQual:          unifieddatalibrary.Int(6),
			Month:                 unifieddatalibrary.String("OCT"),
			MultiDuty: []unifieddatalibrary.DatalinkIngestMultiDutyParam{{
				Duty:             unifieddatalibrary.String("SICO"),
				DutyTeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
				MultiDutyVoiceCoord: []unifieddatalibrary.DatalinkIngestMultiDutyMultiDutyVoiceCoordParam{{
					MultiCommPri:      unifieddatalibrary.String("P"),
					MultiFreqDes:      unifieddatalibrary.String("ST300A"),
					MultiTeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
					MultiVoiceNetDes:  unifieddatalibrary.String("VPN"),
				}},
				Name:    unifieddatalibrary.String("POPOVICH"),
				Rank:    unifieddatalibrary.String("LCDR"),
				UnitDes: unifieddatalibrary.String("SHIP:STENNIS"),
			}},
			NonLinkUnitDes: []string{"CS:GRAY GHOST", "CS:WHITE WHALE"},
			OpExInfo:       unifieddatalibrary.String("CONTROL"),
			OpExInfoAlt:    unifieddatalibrary.String("ORANGE"),
			Ops: []unifieddatalibrary.DatalinkIngestOpParam{{
				LinkDetails:     unifieddatalibrary.String("Link details"),
				LinkName:        unifieddatalibrary.String("Link-16"),
				LinkStartTime:   unifieddatalibrary.Time(time.Now()),
				LinkStopTime:    unifieddatalibrary.Time(time.Now()),
				LinkStopTimeMod: unifieddatalibrary.String("AFTER"),
			}},
			Origin:      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlanOrigNum: unifieddatalibrary.String("SACEUR 106"),
			PocCallSign: unifieddatalibrary.String("4077 MASH"),
			PocLat:      unifieddatalibrary.Float(45.23),
			PocLocName:  unifieddatalibrary.String("CAMP SWAMPY"),
			PocLon:      unifieddatalibrary.Float(179.1),
			PocName:     unifieddatalibrary.String("F. BURNS"),
			PocNums:     []string{"TEL:804-555-4142", "TEL:804-867-5309"},
			PocRank:     unifieddatalibrary.String("MAJ"),
			Qualifier:   unifieddatalibrary.String("CHG"),
			QualSn:      unifieddatalibrary.Int(1),
			References: []unifieddatalibrary.DatalinkIngestReferenceParam{{
				RefOriginator:      unifieddatalibrary.String("CENTCOM"),
				RefSerialID:        unifieddatalibrary.String("A"),
				RefSerialNum:       unifieddatalibrary.String("1402001"),
				RefSiCs:            []string{"RDU", "C-123-92"},
				RefSpecialNotation: unifieddatalibrary.String("NOTAL"),
				RefTs:              unifieddatalibrary.Time(time.Now()),
				RefType:            unifieddatalibrary.String("ABSTAT"),
			}},
			RefPoints: []unifieddatalibrary.DatalinkIngestRefPointParam{{
				EffEventTime: unifieddatalibrary.Time(time.Now()),
				RefDes:       unifieddatalibrary.String("L5"),
				RefLat:       unifieddatalibrary.Float(45.23),
				RefLocName:   unifieddatalibrary.String("FORT BRAGG"),
				RefLon:       unifieddatalibrary.Float(179.1),
				RefPointType: unifieddatalibrary.String("DLRP"),
			}},
			Remarks: []unifieddatalibrary.DatalinkIngestRemarkParam{{
				Text: unifieddatalibrary.String("Example data link remarks"),
				Type: unifieddatalibrary.String("CONTINGENCY PROCEDURES"),
			}},
			ResTrackQual: unifieddatalibrary.Int(3),
			SerialNum:    unifieddatalibrary.String("1201003"),
			SpecTracks: []unifieddatalibrary.DatalinkIngestSpecTrackParam{{
				SpecTrackNum:     unifieddatalibrary.String("12345"),
				SpecTrackNumDesc: unifieddatalibrary.String("SAM SITE CHARLIE"),
			}},
			SpeedDiff:        unifieddatalibrary.Int(50),
			StopTime:         unifieddatalibrary.Time(time.Now()),
			StopTimeMod:      unifieddatalibrary.String("AFTER"),
			SysDefaultCode:   unifieddatalibrary.String("MAN"),
			TrackNumBlockLLs: []int64{1234, 2345},
			TrackNumBlocks:   []string{"0200-0300", "0400-4412"},
			VoiceCoord: []unifieddatalibrary.DatalinkIngestVoiceCoordParam{{
				CommPri:      unifieddatalibrary.String("P"),
				FreqDes:      unifieddatalibrary.String("ST300A"),
				TeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
				VoiceNetDes:  unifieddatalibrary.String("VPN"),
			}},
			WinSizeMin:  unifieddatalibrary.Float(1.25),
			WinSizeMult: unifieddatalibrary.Float(2.1),
		},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLinkStatusDatalinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkStatus.Datalink.List(context.TODO(), unifieddatalibrary.LinkStatusDatalinkListParams{
		StartTime:   time.Now(),
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

func TestLinkStatusDatalinkCountWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkStatus.Datalink.Count(context.TODO(), unifieddatalibrary.LinkStatusDatalinkCountParams{
		StartTime:   time.Now(),
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

func TestLinkStatusDatalinkQueryhelp(t *testing.T) {
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
	_, err := client.LinkStatus.Datalink.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLinkStatusDatalinkTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkStatus.Datalink.Tuple(context.TODO(), unifieddatalibrary.LinkStatusDatalinkTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestLinkStatusDatalinkUnvalidatedPublish(t *testing.T) {
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
	err := client.LinkStatus.Datalink.UnvalidatedPublish(context.TODO(), unifieddatalibrary.LinkStatusDatalinkUnvalidatedPublishParams{
		Body: []unifieddatalibrary.DatalinkIngestParam{{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.DatalinkIngestDataModeTest,
			OpExName:              "DESERT WIND",
			Originator:            "USCENTCOM",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AckInstUnits:          []string{"AOC EXT 2345", "317 AW"},
			AckReq:                unifieddatalibrary.Bool(true),
			AltDiff:               unifieddatalibrary.Int(20),
			CanxID:                unifieddatalibrary.String("ABSTAT"),
			CanxOriginator:        unifieddatalibrary.String("505 AOC"),
			CanxSerialNum:         unifieddatalibrary.String("ABC1234567"),
			CanxSiCs:              []string{"RDU", "X234BS"},
			CanxSpecialNotation:   unifieddatalibrary.String("PASEP"),
			CanxTs:                unifieddatalibrary.Time(time.Now()),
			ClassReasons:          []string{"15C", "15D"},
			ClassSource:           unifieddatalibrary.String("USJFCOM EXORD SOLID WASTE 98"),
			ConsecDecorr:          unifieddatalibrary.Int(3),
			CourseDiff:            unifieddatalibrary.Int(60),
			DecExemptCodes:        []string{"X1", "X2"},
			DecInstDates:          []string{"AT EXERCISE ENDEX", "DATE:25NOV1997"},
			DecorrWinMult:         unifieddatalibrary.Float(1.7),
			GeoDatum:              unifieddatalibrary.String("EUR-T"),
			JreCallSign:           unifieddatalibrary.String("CHARLIE ONE"),
			JreDetails:            unifieddatalibrary.String("JRE details"),
			JrePriAdd:             unifieddatalibrary.Int(71777),
			JreSecAdd:             unifieddatalibrary.Int(77771),
			JreUnitDes:            unifieddatalibrary.String("CVN-72"),
			MaxGeoPosQual:         unifieddatalibrary.Int(12),
			MaxTrackQual:          unifieddatalibrary.Int(12),
			MgmtCode:              unifieddatalibrary.String("VICTOR"),
			MgmtCodeMeaning:       unifieddatalibrary.String("ORBIT AT POINT BRAVO"),
			MinGeoPosQual:         unifieddatalibrary.Int(3),
			MinTrackQual:          unifieddatalibrary.Int(6),
			Month:                 unifieddatalibrary.String("OCT"),
			MultiDuty: []unifieddatalibrary.DatalinkIngestMultiDutyParam{{
				Duty:             unifieddatalibrary.String("SICO"),
				DutyTeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
				MultiDutyVoiceCoord: []unifieddatalibrary.DatalinkIngestMultiDutyMultiDutyVoiceCoordParam{{
					MultiCommPri:      unifieddatalibrary.String("P"),
					MultiFreqDes:      unifieddatalibrary.String("ST300A"),
					MultiTeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
					MultiVoiceNetDes:  unifieddatalibrary.String("VPN"),
				}},
				Name:    unifieddatalibrary.String("POPOVICH"),
				Rank:    unifieddatalibrary.String("LCDR"),
				UnitDes: unifieddatalibrary.String("SHIP:STENNIS"),
			}},
			NonLinkUnitDes: []string{"CS:GRAY GHOST", "CS:WHITE WHALE"},
			OpExInfo:       unifieddatalibrary.String("CONTROL"),
			OpExInfoAlt:    unifieddatalibrary.String("ORANGE"),
			Ops: []unifieddatalibrary.DatalinkIngestOpParam{{
				LinkDetails:     unifieddatalibrary.String("Link details"),
				LinkName:        unifieddatalibrary.String("Link-16"),
				LinkStartTime:   unifieddatalibrary.Time(time.Now()),
				LinkStopTime:    unifieddatalibrary.Time(time.Now()),
				LinkStopTimeMod: unifieddatalibrary.String("AFTER"),
			}},
			Origin:      unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			PlanOrigNum: unifieddatalibrary.String("SACEUR 106"),
			PocCallSign: unifieddatalibrary.String("4077 MASH"),
			PocLat:      unifieddatalibrary.Float(45.23),
			PocLocName:  unifieddatalibrary.String("CAMP SWAMPY"),
			PocLon:      unifieddatalibrary.Float(179.1),
			PocName:     unifieddatalibrary.String("F. BURNS"),
			PocNums:     []string{"TEL:804-555-4142", "TEL:804-867-5309"},
			PocRank:     unifieddatalibrary.String("MAJ"),
			Qualifier:   unifieddatalibrary.String("CHG"),
			QualSn:      unifieddatalibrary.Int(1),
			References: []unifieddatalibrary.DatalinkIngestReferenceParam{{
				RefOriginator:      unifieddatalibrary.String("CENTCOM"),
				RefSerialID:        unifieddatalibrary.String("A"),
				RefSerialNum:       unifieddatalibrary.String("1402001"),
				RefSiCs:            []string{"RDU", "C-123-92"},
				RefSpecialNotation: unifieddatalibrary.String("NOTAL"),
				RefTs:              unifieddatalibrary.Time(time.Now()),
				RefType:            unifieddatalibrary.String("ABSTAT"),
			}},
			RefPoints: []unifieddatalibrary.DatalinkIngestRefPointParam{{
				EffEventTime: unifieddatalibrary.Time(time.Now()),
				RefDes:       unifieddatalibrary.String("L5"),
				RefLat:       unifieddatalibrary.Float(45.23),
				RefLocName:   unifieddatalibrary.String("FORT BRAGG"),
				RefLon:       unifieddatalibrary.Float(179.1),
				RefPointType: unifieddatalibrary.String("DLRP"),
			}},
			Remarks: []unifieddatalibrary.DatalinkIngestRemarkParam{{
				Text: unifieddatalibrary.String("Example data link remarks"),
				Type: unifieddatalibrary.String("CONTINGENCY PROCEDURES"),
			}},
			ResTrackQual: unifieddatalibrary.Int(3),
			SerialNum:    unifieddatalibrary.String("1201003"),
			SpecTracks: []unifieddatalibrary.DatalinkIngestSpecTrackParam{{
				SpecTrackNum:     unifieddatalibrary.String("12345"),
				SpecTrackNumDesc: unifieddatalibrary.String("SAM SITE CHARLIE"),
			}},
			SpeedDiff:        unifieddatalibrary.Int(50),
			StopTime:         unifieddatalibrary.Time(time.Now()),
			StopTimeMod:      unifieddatalibrary.String("AFTER"),
			SysDefaultCode:   unifieddatalibrary.String("MAN"),
			TrackNumBlockLLs: []int64{1234, 2345},
			TrackNumBlocks:   []string{"0200-0300", "0400-4412"},
			VoiceCoord: []unifieddatalibrary.DatalinkIngestVoiceCoordParam{{
				CommPri:      unifieddatalibrary.String("P"),
				FreqDes:      unifieddatalibrary.String("ST300A"),
				TeleFreqNums: []string{"TEL:804-555-4142", "TEL:804-867-5309"},
				VoiceNetDes:  unifieddatalibrary.String("VPN"),
			}},
			WinSizeMin:  unifieddatalibrary.Float(1.25),
			WinSizeMult: unifieddatalibrary.Float(2.1),
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
