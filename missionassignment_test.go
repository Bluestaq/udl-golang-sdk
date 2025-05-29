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

func TestMissionAssignmentNewWithOptionalParams(t *testing.T) {
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
	err := client.MissionAssignment.New(context.TODO(), unifieddatalibrary.MissionAssignmentNewParams{
		ClassificationMarking:      "U",
		DataMode:                   unifieddatalibrary.MissionAssignmentNewParamsDataModeTest,
		Mad:                        "MAD",
		Source:                     "Bluestaq",
		Ts:                         time.Now(),
		ID:                         unifieddatalibrary.String("MISSIONASSIGNMENT-ID"),
		C1associateddmpis:          unifieddatalibrary.Int(3),
		C2air:                      unifieddatalibrary.String("C2AIR"),
		C2alt:                      unifieddatalibrary.Int(3),
		C2crs:                      unifieddatalibrary.Int(3),
		C2exerciseindicator:        unifieddatalibrary.String("C2EXERCISE"),
		C2exercisemof:              unifieddatalibrary.String("MOF"),
		C2id:                       unifieddatalibrary.String("C2ID"),
		C2idamplifyingdescriptor:   unifieddatalibrary.String("C2IDAMP"),
		C2lnd:                      unifieddatalibrary.String("C2LND"),
		C2spc:                      unifieddatalibrary.String("C2SPC"),
		C2spd:                      unifieddatalibrary.Int(3),
		C2specialinterestindicator: unifieddatalibrary.String("C2SPECIAL"),
		C2sur:                      unifieddatalibrary.String("C2SUR"),
		C3elv:                      unifieddatalibrary.Float(10.23),
		C3lat:                      unifieddatalibrary.Float(10.23),
		C3lon:                      unifieddatalibrary.Float(10.23),
		C3ptl:                      unifieddatalibrary.String("C3PTL"),
		C3ptnum:                    unifieddatalibrary.String("C3PTNUM"),
		C4colon:                    unifieddatalibrary.Int(5),
		C4def:                      unifieddatalibrary.String("C4DEF"),
		C4egress:                   unifieddatalibrary.Int(4),
		C4mod:                      unifieddatalibrary.Int(5),
		C4numberofstores:           unifieddatalibrary.Int(3),
		C4runin:                    unifieddatalibrary.Int(5),
		C4tgt:                      unifieddatalibrary.String("C4TGT"),
		C4timediscrete:             unifieddatalibrary.String("C4TIMED"),
		C4tm:                       unifieddatalibrary.Int(4),
		C4typeofstores:             unifieddatalibrary.Int(2),
		C5colon:                    unifieddatalibrary.Int(5),
		C5elevationlsbs:            unifieddatalibrary.Int(5),
		C5haeadj:                   unifieddatalibrary.Int(5),
		C5latlsb:                   unifieddatalibrary.Int(5),
		C5lonlsb:                   unifieddatalibrary.Int(5),
		C5tgtbrng:                  unifieddatalibrary.Int(5),
		C5tw:                       unifieddatalibrary.Int(5),
		C6dspc:                     unifieddatalibrary.String("C6DSPC"),
		C6dspct:                    unifieddatalibrary.String("C6DSPCT"),
		C6fplpm:                    unifieddatalibrary.String("C6FPLPM"),
		C6intel:                    unifieddatalibrary.Int(5),
		C6laser:                    unifieddatalibrary.Int(5),
		C6longpm:                   unifieddatalibrary.String("C6LONGPM"),
		C6tnr3:                     unifieddatalibrary.Int(5),
		C7elang2:                   unifieddatalibrary.Float(5.23),
		C7in3p:                     unifieddatalibrary.Int(3),
		C7tnor:                     unifieddatalibrary.String("C7TNOR"),
		Env:                        unifieddatalibrary.String("ENV"),
		Index:                      unifieddatalibrary.Int(5),
		Lat:                        unifieddatalibrary.Float(45.23),
		Lon:                        unifieddatalibrary.Float(45.23),
		Orginx:                     unifieddatalibrary.String("ORIGIN"),
		Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Rc:                         unifieddatalibrary.String("RC-123"),
		Rr:                         unifieddatalibrary.Int(2),
		Sz:                         unifieddatalibrary.String("STRENGTH"),
		Tno:                        unifieddatalibrary.String("TRACK_NUMBER"),
		TrkID:                      unifieddatalibrary.String("TRK-ID"),
		Twenv:                      unifieddatalibrary.String("THREAT_WARNING"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMissionAssignmentUpdateWithOptionalParams(t *testing.T) {
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
	err := client.MissionAssignment.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.MissionAssignmentUpdateParams{
			ClassificationMarking:      "U",
			DataMode:                   unifieddatalibrary.MissionAssignmentUpdateParamsDataModeTest,
			Mad:                        "MAD",
			Source:                     "Bluestaq",
			Ts:                         time.Now(),
			ID:                         unifieddatalibrary.String("MISSIONASSIGNMENT-ID"),
			C1associateddmpis:          unifieddatalibrary.Int(3),
			C2air:                      unifieddatalibrary.String("C2AIR"),
			C2alt:                      unifieddatalibrary.Int(3),
			C2crs:                      unifieddatalibrary.Int(3),
			C2exerciseindicator:        unifieddatalibrary.String("C2EXERCISE"),
			C2exercisemof:              unifieddatalibrary.String("MOF"),
			C2id:                       unifieddatalibrary.String("C2ID"),
			C2idamplifyingdescriptor:   unifieddatalibrary.String("C2IDAMP"),
			C2lnd:                      unifieddatalibrary.String("C2LND"),
			C2spc:                      unifieddatalibrary.String("C2SPC"),
			C2spd:                      unifieddatalibrary.Int(3),
			C2specialinterestindicator: unifieddatalibrary.String("C2SPECIAL"),
			C2sur:                      unifieddatalibrary.String("C2SUR"),
			C3elv:                      unifieddatalibrary.Float(10.23),
			C3lat:                      unifieddatalibrary.Float(10.23),
			C3lon:                      unifieddatalibrary.Float(10.23),
			C3ptl:                      unifieddatalibrary.String("C3PTL"),
			C3ptnum:                    unifieddatalibrary.String("C3PTNUM"),
			C4colon:                    unifieddatalibrary.Int(5),
			C4def:                      unifieddatalibrary.String("C4DEF"),
			C4egress:                   unifieddatalibrary.Int(4),
			C4mod:                      unifieddatalibrary.Int(5),
			C4numberofstores:           unifieddatalibrary.Int(3),
			C4runin:                    unifieddatalibrary.Int(5),
			C4tgt:                      unifieddatalibrary.String("C4TGT"),
			C4timediscrete:             unifieddatalibrary.String("C4TIMED"),
			C4tm:                       unifieddatalibrary.Int(4),
			C4typeofstores:             unifieddatalibrary.Int(2),
			C5colon:                    unifieddatalibrary.Int(5),
			C5elevationlsbs:            unifieddatalibrary.Int(5),
			C5haeadj:                   unifieddatalibrary.Int(5),
			C5latlsb:                   unifieddatalibrary.Int(5),
			C5lonlsb:                   unifieddatalibrary.Int(5),
			C5tgtbrng:                  unifieddatalibrary.Int(5),
			C5tw:                       unifieddatalibrary.Int(5),
			C6dspc:                     unifieddatalibrary.String("C6DSPC"),
			C6dspct:                    unifieddatalibrary.String("C6DSPCT"),
			C6fplpm:                    unifieddatalibrary.String("C6FPLPM"),
			C6intel:                    unifieddatalibrary.Int(5),
			C6laser:                    unifieddatalibrary.Int(5),
			C6longpm:                   unifieddatalibrary.String("C6LONGPM"),
			C6tnr3:                     unifieddatalibrary.Int(5),
			C7elang2:                   unifieddatalibrary.Float(5.23),
			C7in3p:                     unifieddatalibrary.Int(3),
			C7tnor:                     unifieddatalibrary.String("C7TNOR"),
			Env:                        unifieddatalibrary.String("ENV"),
			Index:                      unifieddatalibrary.Int(5),
			Lat:                        unifieddatalibrary.Float(45.23),
			Lon:                        unifieddatalibrary.Float(45.23),
			Orginx:                     unifieddatalibrary.String("ORIGIN"),
			Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Rc:                         unifieddatalibrary.String("RC-123"),
			Rr:                         unifieddatalibrary.Int(2),
			Sz:                         unifieddatalibrary.String("STRENGTH"),
			Tno:                        unifieddatalibrary.String("TRACK_NUMBER"),
			TrkID:                      unifieddatalibrary.String("TRK-ID"),
			Twenv:                      unifieddatalibrary.String("THREAT_WARNING"),
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

func TestMissionAssignmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.MissionAssignment.List(context.TODO(), unifieddatalibrary.MissionAssignmentListParams{
		Ts:          time.Now(),
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

func TestMissionAssignmentDelete(t *testing.T) {
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
	err := client.MissionAssignment.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMissionAssignmentCountWithOptionalParams(t *testing.T) {
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
	_, err := client.MissionAssignment.Count(context.TODO(), unifieddatalibrary.MissionAssignmentCountParams{
		Ts:          time.Now(),
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

func TestMissionAssignmentNewBulk(t *testing.T) {
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
	err := client.MissionAssignment.NewBulk(context.TODO(), unifieddatalibrary.MissionAssignmentNewBulkParams{
		Body: []unifieddatalibrary.MissionAssignmentNewBulkParamsBody{{
			ClassificationMarking:      "U",
			DataMode:                   "TEST",
			Mad:                        "MAD",
			Source:                     "Bluestaq",
			Ts:                         time.Now(),
			ID:                         unifieddatalibrary.String("MISSIONASSIGNMENT-ID"),
			C1associateddmpis:          unifieddatalibrary.Int(3),
			C2air:                      unifieddatalibrary.String("C2AIR"),
			C2alt:                      unifieddatalibrary.Int(3),
			C2crs:                      unifieddatalibrary.Int(3),
			C2exerciseindicator:        unifieddatalibrary.String("C2EXERCISE"),
			C2exercisemof:              unifieddatalibrary.String("MOF"),
			C2id:                       unifieddatalibrary.String("C2ID"),
			C2idamplifyingdescriptor:   unifieddatalibrary.String("C2IDAMP"),
			C2lnd:                      unifieddatalibrary.String("C2LND"),
			C2spc:                      unifieddatalibrary.String("C2SPC"),
			C2spd:                      unifieddatalibrary.Int(3),
			C2specialinterestindicator: unifieddatalibrary.String("C2SPECIAL"),
			C2sur:                      unifieddatalibrary.String("C2SUR"),
			C3elv:                      unifieddatalibrary.Float(10.23),
			C3lat:                      unifieddatalibrary.Float(10.23),
			C3lon:                      unifieddatalibrary.Float(10.23),
			C3ptl:                      unifieddatalibrary.String("C3PTL"),
			C3ptnum:                    unifieddatalibrary.String("C3PTNUM"),
			C4colon:                    unifieddatalibrary.Int(5),
			C4def:                      unifieddatalibrary.String("C4DEF"),
			C4egress:                   unifieddatalibrary.Int(4),
			C4mod:                      unifieddatalibrary.Int(5),
			C4numberofstores:           unifieddatalibrary.Int(3),
			C4runin:                    unifieddatalibrary.Int(5),
			C4tgt:                      unifieddatalibrary.String("C4TGT"),
			C4timediscrete:             unifieddatalibrary.String("C4TIMED"),
			C4tm:                       unifieddatalibrary.Int(4),
			C4typeofstores:             unifieddatalibrary.Int(2),
			C5colon:                    unifieddatalibrary.Int(5),
			C5elevationlsbs:            unifieddatalibrary.Int(5),
			C5haeadj:                   unifieddatalibrary.Int(5),
			C5latlsb:                   unifieddatalibrary.Int(5),
			C5lonlsb:                   unifieddatalibrary.Int(5),
			C5tgtbrng:                  unifieddatalibrary.Int(5),
			C5tw:                       unifieddatalibrary.Int(5),
			C6dspc:                     unifieddatalibrary.String("C6DSPC"),
			C6dspct:                    unifieddatalibrary.String("C6DSPCT"),
			C6fplpm:                    unifieddatalibrary.String("C6FPLPM"),
			C6intel:                    unifieddatalibrary.Int(5),
			C6laser:                    unifieddatalibrary.Int(5),
			C6longpm:                   unifieddatalibrary.String("C6LONGPM"),
			C6tnr3:                     unifieddatalibrary.Int(5),
			C7elang2:                   unifieddatalibrary.Float(5.23),
			C7in3p:                     unifieddatalibrary.Int(3),
			C7tnor:                     unifieddatalibrary.String("C7TNOR"),
			Env:                        unifieddatalibrary.String("ENV"),
			Index:                      unifieddatalibrary.Int(5),
			Lat:                        unifieddatalibrary.Float(45.23),
			Lon:                        unifieddatalibrary.Float(45.23),
			Orginx:                     unifieddatalibrary.String("ORIGIN"),
			Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Rc:                         unifieddatalibrary.String("RC-123"),
			Rr:                         unifieddatalibrary.Int(2),
			Sz:                         unifieddatalibrary.String("STRENGTH"),
			Tno:                        unifieddatalibrary.String("TRACK_NUMBER"),
			TrkID:                      unifieddatalibrary.String("TRK-ID"),
			Twenv:                      unifieddatalibrary.String("THREAT_WARNING"),
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

func TestMissionAssignmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.MissionAssignment.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.MissionAssignmentGetParams{
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

func TestMissionAssignmentQueryhelp(t *testing.T) {
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
	_, err := client.MissionAssignment.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMissionAssignmentTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.MissionAssignment.Tuple(context.TODO(), unifieddatalibrary.MissionAssignmentTupleParams{
		Columns:     "columns",
		Ts:          time.Now(),
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
