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

func TestObservationSwirUnvalidatedPublish(t *testing.T) {
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
	err := client.Observations.Swir.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ObservationSwirUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ObservationSwirUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("SWIR-ID"),
			AbsFluxes:             []float64{1.23, 4.56},
			BadWave:               unifieddatalibrary.String("Example Comments"),
			FluxRatios:            []float64{1.23, 4.56},
			Lat:                   unifieddatalibrary.Float(70.55208),
			LocationName:          unifieddatalibrary.String("AeroTel"),
			Lon:                   unifieddatalibrary.Float(81.18191),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("WildBlue-1"),
			RatioWavelengths:      []float64{1.23, 4.56},
			SatNo:                 unifieddatalibrary.Int(25544),
			SolarPhaseAngle:       unifieddatalibrary.Float(1.23),
			Wavelengths:           []float64{1.23, 4.56},
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
