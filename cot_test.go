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

func TestCotNewWithOptionalParams(t *testing.T) {
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
	err := client.Cots.New(context.TODO(), unifieddatalibrary.CotNewParams{
		Lat:       45.23,
		Lon:       45.23,
		Alt:       unifieddatalibrary.Float(5.23),
		CallSigns: []string{"string"},
		Ce:        unifieddatalibrary.Float(10.23),
		CotChatData: unifieddatalibrary.CotNewParamsCotChatData{
			ChatMsg:            unifieddatalibrary.String("Mission is go"),
			ChatRoom:           unifieddatalibrary.String("All Chat Rooms"),
			ChatSenderCallSign: unifieddatalibrary.String("Pebble"),
		},
		CotPositionData: unifieddatalibrary.CotNewParamsCotPositionData{
			CallSign: "POI_NAME",
			Team:     "Description of the object",
			TeamRole: "Team Member",
		},
		Groups:    []string{"string"},
		How:       unifieddatalibrary.String("h-e"),
		Le:        unifieddatalibrary.Float(10.23),
		SenderUid: unifieddatalibrary.String("POI-ID"),
		Stale:     unifieddatalibrary.Time(time.Now()),
		Start:     unifieddatalibrary.Time(time.Now()),
		Type:      unifieddatalibrary.String("a-h-G"),
		Uids:      []string{"string"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
