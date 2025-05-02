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

func TestMtiListWithOptionalParams(t *testing.T) {
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
	_, err := client.Mti.List(context.TODO(), unifieddatalibrary.MtiListParams{
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

func TestMtiCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Mti.Count(context.TODO(), unifieddatalibrary.MtiCountParams{
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

func TestMtiNewBulk(t *testing.T) {
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
	err := client.Mti.NewBulk(context.TODO(), unifieddatalibrary.MtiNewBulkParams{
		Body: []unifieddatalibrary.MtiNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("MTI-ID"),
			Dwells: []unifieddatalibrary.MtiNewBulkParamsBodyDwell{{
				D10: unifieddatalibrary.Float(1.2),
				D11: unifieddatalibrary.Float(1.2),
				D12: unifieddatalibrary.Int(12),
				D13: unifieddatalibrary.Int(12),
				D14: unifieddatalibrary.Int(12),
				D15: unifieddatalibrary.Float(1.2),
				D16: unifieddatalibrary.Int(12),
				D17: unifieddatalibrary.Int(2),
				D18: unifieddatalibrary.Int(2),
				D19: unifieddatalibrary.Int(12),
				D2:  unifieddatalibrary.Int(12),
				D20: unifieddatalibrary.Int(12),
				D21: unifieddatalibrary.Float(1.2),
				D22: unifieddatalibrary.Float(12.23),
				D23: unifieddatalibrary.Float(12.23),
				D24: unifieddatalibrary.Float(12.23),
				D25: unifieddatalibrary.Float(12.23),
				D26: unifieddatalibrary.Float(12.23),
				D27: unifieddatalibrary.Float(12.23),
				D28: unifieddatalibrary.Float(12.23),
				D29: unifieddatalibrary.Float(12.23),
				D3:  unifieddatalibrary.Int(12),
				D30: unifieddatalibrary.Float(12.23),
				D31: unifieddatalibrary.Int(1),
				D32: []unifieddatalibrary.MtiNewBulkParamsBodyDwellD32{{
					D32_1:  unifieddatalibrary.Int(2),
					D32_10: unifieddatalibrary.String("vehicle"),
					D32_11: unifieddatalibrary.Int(90),
					D32_12: unifieddatalibrary.Int(2),
					D32_13: unifieddatalibrary.Int(2),
					D32_14: unifieddatalibrary.Int(2),
					D32_15: unifieddatalibrary.Int(2),
					D32_16: unifieddatalibrary.Int(2),
					D32_17: unifieddatalibrary.Int(1234567890),
					D32_18: unifieddatalibrary.Int(98),
					D32_2:  unifieddatalibrary.Float(20.23),
					D32_3:  unifieddatalibrary.Float(20.23),
					D32_4:  unifieddatalibrary.Int(2),
					D32_5:  unifieddatalibrary.Int(2),
					D32_6:  unifieddatalibrary.Int(2),
					D32_7:  unifieddatalibrary.Int(2),
					D32_8:  unifieddatalibrary.Int(2),
					D32_9:  unifieddatalibrary.Int(17),
				}},
				D4:      unifieddatalibrary.Bool(false),
				D5:      unifieddatalibrary.Int(12),
				D6:      unifieddatalibrary.Int(1234567890),
				D7:      unifieddatalibrary.Float(1.2),
				D8:      unifieddatalibrary.Float(12),
				D9:      unifieddatalibrary.Int(12),
				Dwellts: unifieddatalibrary.Time(time.Now()),
			}},
			FreeTexts: []unifieddatalibrary.MtiNewBulkParamsBodyFreeText{{
				F1: unifieddatalibrary.String("ORIGINATOR"),
				F2: unifieddatalibrary.String("RECIPIENT"),
				F3: unifieddatalibrary.String("TEXT"),
			}},
			Hrrs: []unifieddatalibrary.MtiNewBulkParamsBodyHrr{{
				H10: unifieddatalibrary.Int(1),
				H11: unifieddatalibrary.Float(12.23),
				H12: unifieddatalibrary.Float(12.23),
				H13: unifieddatalibrary.Float(12.23),
				H14: unifieddatalibrary.Float(12.23),
				H15: unifieddatalibrary.Float(12.23),
				H16: unifieddatalibrary.String("TABLE"),
				H17: unifieddatalibrary.String("TABLE"),
				H18: unifieddatalibrary.String("TABLE"),
				H19: unifieddatalibrary.Float(12.23),
				H2:  unifieddatalibrary.Int(12),
				H20: unifieddatalibrary.Int(117),
				H21: unifieddatalibrary.Int(1),
				H22: unifieddatalibrary.Int(1),
				H23: unifieddatalibrary.String("FIELD"),
				H24: unifieddatalibrary.String("FLAG"),
				H27: unifieddatalibrary.Int(1),
				H28: unifieddatalibrary.Int(1234567890),
				H29: unifieddatalibrary.Int(1),
				H3:  unifieddatalibrary.Int(12),
				H30: unifieddatalibrary.Int(22),
				H31: unifieddatalibrary.Int(55),
				H32: []unifieddatalibrary.MtiNewBulkParamsBodyHrrH32{{
					H32_1: unifieddatalibrary.Int(1),
					H32_2: unifieddatalibrary.Int(1),
					H32_3: unifieddatalibrary.Int(1),
					H32_4: unifieddatalibrary.Int(1),
				}},
				H4: unifieddatalibrary.Bool(true),
				H5: unifieddatalibrary.Int(12),
				H6: unifieddatalibrary.Int(12),
				H7: unifieddatalibrary.Int(12),
				H8: unifieddatalibrary.Int(12),
				H9: unifieddatalibrary.Int(1),
			}},
			JobDefs: []unifieddatalibrary.MtiNewBulkParamsBodyJobDef{{
				J1:  unifieddatalibrary.Int(1234567890),
				J10: unifieddatalibrary.Float(10.23),
				J11: unifieddatalibrary.Float(10.23),
				J12: unifieddatalibrary.Float(10.23),
				J13: unifieddatalibrary.Float(10.23),
				J14: unifieddatalibrary.String("MODE"),
				J15: unifieddatalibrary.Int(100),
				J16: unifieddatalibrary.Int(100),
				J17: unifieddatalibrary.Int(100),
				J18: unifieddatalibrary.Int(100),
				J19: unifieddatalibrary.Int(10),
				J2:  unifieddatalibrary.String("TYPE"),
				J20: unifieddatalibrary.Int(10),
				J21: unifieddatalibrary.Int(10),
				J22: unifieddatalibrary.Float(10.23),
				J23: unifieddatalibrary.Int(10),
				J24: unifieddatalibrary.Int(10),
				J25: unifieddatalibrary.Int(10),
				J26: unifieddatalibrary.Int(10),
				J27: unifieddatalibrary.String("MODEL"),
				J28: unifieddatalibrary.String("MODEL"),
				J3:  unifieddatalibrary.String("J3-ID"),
				J4:  unifieddatalibrary.Int(3),
				J5:  unifieddatalibrary.Int(1),
				J6:  unifieddatalibrary.Float(10.23),
				J7:  unifieddatalibrary.Float(10.23),
				J8:  unifieddatalibrary.Float(10.23),
				J9:  unifieddatalibrary.Float(10.23),
			}},
			JobRequests: []unifieddatalibrary.MtiNewBulkParamsBodyJobRequest{{
				JobReqEst: unifieddatalibrary.Time(time.Now()),
				R1:        unifieddatalibrary.String("REQUESTER"),
				R10:       unifieddatalibrary.Float(10.23),
				R11:       unifieddatalibrary.Float(10.23),
				R12:       unifieddatalibrary.String("MODE"),
				R13:       unifieddatalibrary.Int(10),
				R14:       unifieddatalibrary.Int(100),
				R2:        unifieddatalibrary.String("IDENTIFIER"),
				R21:       unifieddatalibrary.Int(10),
				R22:       unifieddatalibrary.Int(10),
				R23:       unifieddatalibrary.Int(100),
				R24:       unifieddatalibrary.String("TYPE"),
				R25:       unifieddatalibrary.String("VARIANT"),
				R26:       unifieddatalibrary.Bool(true),
				R3:        unifieddatalibrary.Int(15),
				R4:        unifieddatalibrary.Float(10.23),
				R5:        unifieddatalibrary.Float(10.23),
				R6:        unifieddatalibrary.Float(10.23),
				R7:        unifieddatalibrary.Float(10.23),
				R8:        unifieddatalibrary.Float(10.23),
				R9:        unifieddatalibrary.Float(10.23),
			}},
			Missions: []unifieddatalibrary.MtiNewBulkParamsBodyMission{{
				M1:       unifieddatalibrary.String("M1-ID"),
				M2:       unifieddatalibrary.String("M2-ID"),
				M3:       unifieddatalibrary.String("PLATFORM"),
				M4:       unifieddatalibrary.String("IDENT"),
				MsnRefTs: unifieddatalibrary.Time(time.Now()),
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			P10:    unifieddatalibrary.Int(45),
			P3:     unifieddatalibrary.String("NATIONALITY"),
			P6:     unifieddatalibrary.String("MARKING"),
			P7:     unifieddatalibrary.String("INDICATOR"),
			P8:     unifieddatalibrary.String("P8-ID"),
			P9:     unifieddatalibrary.Int(45),
			PlatformLocs: []unifieddatalibrary.MtiNewBulkParamsBodyPlatformLoc{{
				L1:        unifieddatalibrary.Int(1234567890),
				L2:        unifieddatalibrary.Float(45.23),
				L3:        unifieddatalibrary.Float(45.23),
				L4:        unifieddatalibrary.Int(45),
				L5:        unifieddatalibrary.Float(45.23),
				L6:        unifieddatalibrary.Int(50),
				L7:        unifieddatalibrary.Int(82),
				Platlocts: unifieddatalibrary.Time(time.Now()),
			}},
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

func TestMtiQueryhelp(t *testing.T) {
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
	err := client.Mti.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMtiTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Mti.Tuple(context.TODO(), unifieddatalibrary.MtiTupleParams{
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

func TestMtiUnvalidatedPublish(t *testing.T) {
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
	err := client.Mti.UnvalidatedPublish(context.TODO(), unifieddatalibrary.MtiUnvalidatedPublishParams{
		Body: []unifieddatalibrary.MtiUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("MTI-ID"),
			Dwells: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyDwell{{
				D10: unifieddatalibrary.Float(1.2),
				D11: unifieddatalibrary.Float(1.2),
				D12: unifieddatalibrary.Int(12),
				D13: unifieddatalibrary.Int(12),
				D14: unifieddatalibrary.Int(12),
				D15: unifieddatalibrary.Float(1.2),
				D16: unifieddatalibrary.Int(12),
				D17: unifieddatalibrary.Int(2),
				D18: unifieddatalibrary.Int(2),
				D19: unifieddatalibrary.Int(12),
				D2:  unifieddatalibrary.Int(12),
				D20: unifieddatalibrary.Int(12),
				D21: unifieddatalibrary.Float(1.2),
				D22: unifieddatalibrary.Float(12.23),
				D23: unifieddatalibrary.Float(12.23),
				D24: unifieddatalibrary.Float(12.23),
				D25: unifieddatalibrary.Float(12.23),
				D26: unifieddatalibrary.Float(12.23),
				D27: unifieddatalibrary.Float(12.23),
				D28: unifieddatalibrary.Float(12.23),
				D29: unifieddatalibrary.Float(12.23),
				D3:  unifieddatalibrary.Int(12),
				D30: unifieddatalibrary.Float(12.23),
				D31: unifieddatalibrary.Int(1),
				D32: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyDwellD32{{
					D32_1:  unifieddatalibrary.Int(2),
					D32_10: unifieddatalibrary.String("vehicle"),
					D32_11: unifieddatalibrary.Int(90),
					D32_12: unifieddatalibrary.Int(2),
					D32_13: unifieddatalibrary.Int(2),
					D32_14: unifieddatalibrary.Int(2),
					D32_15: unifieddatalibrary.Int(2),
					D32_16: unifieddatalibrary.Int(2),
					D32_17: unifieddatalibrary.Int(1234567890),
					D32_18: unifieddatalibrary.Int(98),
					D32_2:  unifieddatalibrary.Float(20.23),
					D32_3:  unifieddatalibrary.Float(20.23),
					D32_4:  unifieddatalibrary.Int(2),
					D32_5:  unifieddatalibrary.Int(2),
					D32_6:  unifieddatalibrary.Int(2),
					D32_7:  unifieddatalibrary.Int(2),
					D32_8:  unifieddatalibrary.Int(2),
					D32_9:  unifieddatalibrary.Int(17),
				}},
				D4:      unifieddatalibrary.Bool(false),
				D5:      unifieddatalibrary.Int(12),
				D6:      unifieddatalibrary.Int(1234567890),
				D7:      unifieddatalibrary.Float(1.2),
				D8:      unifieddatalibrary.Float(12),
				D9:      unifieddatalibrary.Int(12),
				Dwellts: unifieddatalibrary.Time(time.Now()),
			}},
			FreeTexts: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyFreeText{{
				F1: unifieddatalibrary.String("ORIGINATOR"),
				F2: unifieddatalibrary.String("RECIPIENT"),
				F3: unifieddatalibrary.String("TEXT"),
			}},
			Hrrs: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyHrr{{
				H10: unifieddatalibrary.Int(1),
				H11: unifieddatalibrary.Float(12.23),
				H12: unifieddatalibrary.Float(12.23),
				H13: unifieddatalibrary.Float(12.23),
				H14: unifieddatalibrary.Float(12.23),
				H15: unifieddatalibrary.Float(12.23),
				H16: unifieddatalibrary.String("TABLE"),
				H17: unifieddatalibrary.String("TABLE"),
				H18: unifieddatalibrary.String("TABLE"),
				H19: unifieddatalibrary.Float(12.23),
				H2:  unifieddatalibrary.Int(12),
				H20: unifieddatalibrary.Int(117),
				H21: unifieddatalibrary.Int(1),
				H22: unifieddatalibrary.Int(1),
				H23: unifieddatalibrary.String("FIELD"),
				H24: unifieddatalibrary.String("FLAG"),
				H27: unifieddatalibrary.Int(1),
				H28: unifieddatalibrary.Int(1234567890),
				H29: unifieddatalibrary.Int(1),
				H3:  unifieddatalibrary.Int(12),
				H30: unifieddatalibrary.Int(22),
				H31: unifieddatalibrary.Int(55),
				H32: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyHrrH32{{
					H32_1: unifieddatalibrary.Int(1),
					H32_2: unifieddatalibrary.Int(1),
					H32_3: unifieddatalibrary.Int(1),
					H32_4: unifieddatalibrary.Int(1),
				}},
				H4: unifieddatalibrary.Bool(true),
				H5: unifieddatalibrary.Int(12),
				H6: unifieddatalibrary.Int(12),
				H7: unifieddatalibrary.Int(12),
				H8: unifieddatalibrary.Int(12),
				H9: unifieddatalibrary.Int(1),
			}},
			JobDefs: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyJobDef{{
				J1:  unifieddatalibrary.Int(1234567890),
				J10: unifieddatalibrary.Float(10.23),
				J11: unifieddatalibrary.Float(10.23),
				J12: unifieddatalibrary.Float(10.23),
				J13: unifieddatalibrary.Float(10.23),
				J14: unifieddatalibrary.String("MODE"),
				J15: unifieddatalibrary.Int(100),
				J16: unifieddatalibrary.Int(100),
				J17: unifieddatalibrary.Int(100),
				J18: unifieddatalibrary.Int(100),
				J19: unifieddatalibrary.Int(10),
				J2:  unifieddatalibrary.String("TYPE"),
				J20: unifieddatalibrary.Int(10),
				J21: unifieddatalibrary.Int(10),
				J22: unifieddatalibrary.Float(10.23),
				J23: unifieddatalibrary.Int(10),
				J24: unifieddatalibrary.Int(10),
				J25: unifieddatalibrary.Int(10),
				J26: unifieddatalibrary.Int(10),
				J27: unifieddatalibrary.String("MODEL"),
				J28: unifieddatalibrary.String("MODEL"),
				J3:  unifieddatalibrary.String("J3-ID"),
				J4:  unifieddatalibrary.Int(3),
				J5:  unifieddatalibrary.Int(1),
				J6:  unifieddatalibrary.Float(10.23),
				J7:  unifieddatalibrary.Float(10.23),
				J8:  unifieddatalibrary.Float(10.23),
				J9:  unifieddatalibrary.Float(10.23),
			}},
			JobRequests: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyJobRequest{{
				JobReqEst: unifieddatalibrary.Time(time.Now()),
				R1:        unifieddatalibrary.String("REQUESTER"),
				R10:       unifieddatalibrary.Float(10.23),
				R11:       unifieddatalibrary.Float(10.23),
				R12:       unifieddatalibrary.String("MODE"),
				R13:       unifieddatalibrary.Int(10),
				R14:       unifieddatalibrary.Int(100),
				R2:        unifieddatalibrary.String("IDENTIFIER"),
				R21:       unifieddatalibrary.Int(10),
				R22:       unifieddatalibrary.Int(10),
				R23:       unifieddatalibrary.Int(100),
				R24:       unifieddatalibrary.String("TYPE"),
				R25:       unifieddatalibrary.String("VARIANT"),
				R26:       unifieddatalibrary.Bool(true),
				R3:        unifieddatalibrary.Int(15),
				R4:        unifieddatalibrary.Float(10.23),
				R5:        unifieddatalibrary.Float(10.23),
				R6:        unifieddatalibrary.Float(10.23),
				R7:        unifieddatalibrary.Float(10.23),
				R8:        unifieddatalibrary.Float(10.23),
				R9:        unifieddatalibrary.Float(10.23),
			}},
			Missions: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyMission{{
				M1:       unifieddatalibrary.String("M1-ID"),
				M2:       unifieddatalibrary.String("M2-ID"),
				M3:       unifieddatalibrary.String("PLATFORM"),
				M4:       unifieddatalibrary.String("IDENT"),
				MsnRefTs: unifieddatalibrary.Time(time.Now()),
			}},
			Origin: unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			P10:    unifieddatalibrary.Int(45),
			P3:     unifieddatalibrary.String("NATIONALITY"),
			P6:     unifieddatalibrary.String("MARKING"),
			P7:     unifieddatalibrary.String("INDICATOR"),
			P8:     unifieddatalibrary.String("P8-ID"),
			P9:     unifieddatalibrary.Int(45),
			PlatformLocs: []unifieddatalibrary.MtiUnvalidatedPublishParamsBodyPlatformLoc{{
				L1:        unifieddatalibrary.Int(1234567890),
				L2:        unifieddatalibrary.Float(45.23),
				L3:        unifieddatalibrary.Float(45.23),
				L4:        unifieddatalibrary.Int(45),
				L5:        unifieddatalibrary.Float(45.23),
				L6:        unifieddatalibrary.Int(50),
				L7:        unifieddatalibrary.Int(82),
				Platlocts: unifieddatalibrary.Time(time.Now()),
			}},
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
