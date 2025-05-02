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

func TestSarobservationNewWithOptionalParams(t *testing.T) {
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
	err := client.Sarobservation.New(context.TODO(), unifieddatalibrary.SarobservationNewParams{
		ClassificationMarking:      "U",
		CollectionEnd:              time.Now(),
		CollectionStart:            time.Now(),
		DataMode:                   unifieddatalibrary.SarobservationNewParamsDataModeTest,
		SarMode:                    "SPOTLIGHT",
		Source:                     "Bluestaq",
		ID:                         unifieddatalibrary.String("SAROBSERVATION-ID"),
		Agjson:                     unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
		Andims:                     unifieddatalibrary.Int(3),
		Area:                       unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Asrid:                      unifieddatalibrary.Int(3),
		Atext:                      unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Atype:                      unifieddatalibrary.String("POLYGON"),
		AzimuthAngle:               unifieddatalibrary.Float(285.4481793),
		CenterTime:                 unifieddatalibrary.Time(time.Now()),
		CollectionID:               unifieddatalibrary.String("COLLECTION-ID"),
		ContinuousSpotAngle:        unifieddatalibrary.Float(45.1),
		CoordSys:                   unifieddatalibrary.String("ECEF"),
		DetectionEnd:               unifieddatalibrary.Time(time.Now()),
		DetectionID:                unifieddatalibrary.String("DETECTION-ID"),
		DetectionStart:             unifieddatalibrary.Time(time.Now()),
		DwellTime:                  unifieddatalibrary.Float(79.156794),
		ExternalID:                 unifieddatalibrary.String("EXTERNAL-ID"),
		FarRange:                   unifieddatalibrary.Float(34.1),
		GrazeAngle:                 unifieddatalibrary.Float(45.1),
		GroundResolutionProjection: unifieddatalibrary.Float(0.5),
		IDSensor:                   unifieddatalibrary.String("36036-1L"),
		IncidenceAngle:             unifieddatalibrary.Float(45.1),
		LooksAzimuth:               unifieddatalibrary.Int(2),
		LooksRange:                 unifieddatalibrary.Int(1),
		MultilookNumber:            unifieddatalibrary.Float(5),
		NearRange:                  unifieddatalibrary.Float(12.1),
		ObDirection:                unifieddatalibrary.String("RIGHT"),
		OperatingBand:              unifieddatalibrary.String("L"),
		OperatingFreq:              unifieddatalibrary.Float(2345.6),
		OrbitState:                 unifieddatalibrary.String("ASCENDING"),
		Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:               unifieddatalibrary.String("36036"),
		OrigSensorID:               unifieddatalibrary.String("SMOS-1L"),
		PulseBandwidth:             unifieddatalibrary.Float(500.1),
		PulseDuration:              unifieddatalibrary.Float(0.000011),
		ResolutionAzimuth:          unifieddatalibrary.Float(0.123),
		ResolutionRange:            unifieddatalibrary.Float(0.123),
		RxPolarization:             unifieddatalibrary.String("H"),
		SatNo:                      unifieddatalibrary.Int(36036),
		Senalt:                     unifieddatalibrary.Float(1.1),
		SenlatEnd:                  unifieddatalibrary.Float(45.1),
		SenlatStart:                unifieddatalibrary.Float(45.1),
		SenlonEnd:                  unifieddatalibrary.Float(179.1),
		SenlonStart:                unifieddatalibrary.Float(179.1),
		Senvelx:                    unifieddatalibrary.Float(1.1),
		Senvely:                    unifieddatalibrary.Float(1.1),
		Senvelz:                    unifieddatalibrary.Float(1.1),
		SlantRange:                 unifieddatalibrary.Float(60.1),
		Snr:                        unifieddatalibrary.Float(10.1),
		SpacingAzimuth:             unifieddatalibrary.Float(0.123),
		SpacingRange:               unifieddatalibrary.Float(0.123),
		SquintAngle:                unifieddatalibrary.Float(1.2),
		SrcIDs:                     []string{"f7e01cd4-626b-441f-a423-17b160eb78ba", "223833c4-be0d-4fdb-a2e4-325a48eccced"},
		SrcTyps:                    []string{"ESID", "GROUNDIMAGE"},
		SwathLength:                unifieddatalibrary.Float(12.1),
		Tags:                       []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
		Targetposx:                 unifieddatalibrary.Float(50.23),
		Targetposy:                 unifieddatalibrary.Float(50.23),
		Targetposz:                 unifieddatalibrary.Float(50.23),
		TransactionID:              unifieddatalibrary.String("TRANSACTION-ID"),
		TxPolarization:             unifieddatalibrary.String("H"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSarobservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sarobservation.List(context.TODO(), unifieddatalibrary.SarobservationListParams{
		CollectionStart: time.Now(),
		FirstResult:     unifieddatalibrary.Int(0),
		MaxResults:      unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSarobservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sarobservation.Count(context.TODO(), unifieddatalibrary.SarobservationCountParams{
		CollectionStart: time.Now(),
		FirstResult:     unifieddatalibrary.Int(0),
		MaxResults:      unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSarobservationNewBulk(t *testing.T) {
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
	err := client.Sarobservation.NewBulk(context.TODO(), unifieddatalibrary.SarobservationNewBulkParams{
		Body: []unifieddatalibrary.SarobservationNewBulkParamsBody{{
			ClassificationMarking:      "U",
			CollectionEnd:              time.Now(),
			CollectionStart:            time.Now(),
			DataMode:                   "TEST",
			SarMode:                    "SPOTLIGHT",
			Source:                     "Bluestaq",
			ID:                         unifieddatalibrary.String("SAROBSERVATION-ID"),
			Agjson:                     unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Andims:                     unifieddatalibrary.Int(3),
			Area:                       unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                      unifieddatalibrary.Int(3),
			Atext:                      unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                      unifieddatalibrary.String("POLYGON"),
			AzimuthAngle:               unifieddatalibrary.Float(285.4481793),
			CenterTime:                 unifieddatalibrary.Time(time.Now()),
			CollectionID:               unifieddatalibrary.String("COLLECTION-ID"),
			ContinuousSpotAngle:        unifieddatalibrary.Float(45.1),
			CoordSys:                   unifieddatalibrary.String("ECEF"),
			DetectionEnd:               unifieddatalibrary.Time(time.Now()),
			DetectionID:                unifieddatalibrary.String("DETECTION-ID"),
			DetectionStart:             unifieddatalibrary.Time(time.Now()),
			DwellTime:                  unifieddatalibrary.Float(79.156794),
			ExternalID:                 unifieddatalibrary.String("EXTERNAL-ID"),
			FarRange:                   unifieddatalibrary.Float(34.1),
			GrazeAngle:                 unifieddatalibrary.Float(45.1),
			GroundResolutionProjection: unifieddatalibrary.Float(0.5),
			IDSensor:                   unifieddatalibrary.String("36036-1L"),
			IncidenceAngle:             unifieddatalibrary.Float(45.1),
			LooksAzimuth:               unifieddatalibrary.Int(2),
			LooksRange:                 unifieddatalibrary.Int(1),
			MultilookNumber:            unifieddatalibrary.Float(5),
			NearRange:                  unifieddatalibrary.Float(12.1),
			ObDirection:                unifieddatalibrary.String("RIGHT"),
			OperatingBand:              unifieddatalibrary.String("L"),
			OperatingFreq:              unifieddatalibrary.Float(2345.6),
			OrbitState:                 unifieddatalibrary.String("ASCENDING"),
			Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:               unifieddatalibrary.String("36036"),
			OrigSensorID:               unifieddatalibrary.String("SMOS-1L"),
			PulseBandwidth:             unifieddatalibrary.Float(500.1),
			PulseDuration:              unifieddatalibrary.Float(0.000011),
			ResolutionAzimuth:          unifieddatalibrary.Float(0.123),
			ResolutionRange:            unifieddatalibrary.Float(0.123),
			RxPolarization:             unifieddatalibrary.String("H"),
			SatNo:                      unifieddatalibrary.Int(36036),
			Senalt:                     unifieddatalibrary.Float(1.1),
			SenlatEnd:                  unifieddatalibrary.Float(45.1),
			SenlatStart:                unifieddatalibrary.Float(45.1),
			SenlonEnd:                  unifieddatalibrary.Float(179.1),
			SenlonStart:                unifieddatalibrary.Float(179.1),
			Senvelx:                    unifieddatalibrary.Float(1.1),
			Senvely:                    unifieddatalibrary.Float(1.1),
			Senvelz:                    unifieddatalibrary.Float(1.1),
			SlantRange:                 unifieddatalibrary.Float(60.1),
			Snr:                        unifieddatalibrary.Float(10.1),
			SpacingAzimuth:             unifieddatalibrary.Float(0.123),
			SpacingRange:               unifieddatalibrary.Float(0.123),
			SquintAngle:                unifieddatalibrary.Float(1.2),
			SrcIDs:                     []string{"f7e01cd4-626b-441f-a423-17b160eb78ba", "223833c4-be0d-4fdb-a2e4-325a48eccced"},
			SrcTyps:                    []string{"ESID", "GROUNDIMAGE"},
			SwathLength:                unifieddatalibrary.Float(12.1),
			Tags:                       []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			Targetposx:                 unifieddatalibrary.Float(50.23),
			Targetposy:                 unifieddatalibrary.Float(50.23),
			Targetposz:                 unifieddatalibrary.Float(50.23),
			TransactionID:              unifieddatalibrary.String("TRANSACTION-ID"),
			TxPolarization:             unifieddatalibrary.String("H"),
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

func TestSarobservationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sarobservation.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SarobservationGetParams{
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

func TestSarobservationQueryhelp(t *testing.T) {
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
	err := client.Sarobservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSarobservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sarobservation.Tuple(context.TODO(), unifieddatalibrary.SarobservationTupleParams{
		CollectionStart: time.Now(),
		Columns:         "columns",
		FirstResult:     unifieddatalibrary.Int(0),
		MaxResults:      unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSarobservationUnvalidatedPublish(t *testing.T) {
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
	err := client.Sarobservation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SarobservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SarobservationUnvalidatedPublishParamsBody{{
			ClassificationMarking:      "U",
			CollectionEnd:              time.Now(),
			CollectionStart:            time.Now(),
			DataMode:                   "TEST",
			SarMode:                    "SPOTLIGHT",
			Source:                     "Bluestaq",
			ID:                         unifieddatalibrary.String("SAROBSERVATION-ID"),
			Agjson:                     unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Andims:                     unifieddatalibrary.Int(3),
			Area:                       unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                      unifieddatalibrary.Int(3),
			Atext:                      unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                      unifieddatalibrary.String("POLYGON"),
			AzimuthAngle:               unifieddatalibrary.Float(285.4481793),
			CenterTime:                 unifieddatalibrary.Time(time.Now()),
			CollectionID:               unifieddatalibrary.String("COLLECTION-ID"),
			ContinuousSpotAngle:        unifieddatalibrary.Float(45.1),
			CoordSys:                   unifieddatalibrary.String("ECEF"),
			DetectionEnd:               unifieddatalibrary.Time(time.Now()),
			DetectionID:                unifieddatalibrary.String("DETECTION-ID"),
			DetectionStart:             unifieddatalibrary.Time(time.Now()),
			DwellTime:                  unifieddatalibrary.Float(79.156794),
			ExternalID:                 unifieddatalibrary.String("EXTERNAL-ID"),
			FarRange:                   unifieddatalibrary.Float(34.1),
			GrazeAngle:                 unifieddatalibrary.Float(45.1),
			GroundResolutionProjection: unifieddatalibrary.Float(0.5),
			IDSensor:                   unifieddatalibrary.String("36036-1L"),
			IncidenceAngle:             unifieddatalibrary.Float(45.1),
			LooksAzimuth:               unifieddatalibrary.Int(2),
			LooksRange:                 unifieddatalibrary.Int(1),
			MultilookNumber:            unifieddatalibrary.Float(5),
			NearRange:                  unifieddatalibrary.Float(12.1),
			ObDirection:                unifieddatalibrary.String("RIGHT"),
			OperatingBand:              unifieddatalibrary.String("L"),
			OperatingFreq:              unifieddatalibrary.Float(2345.6),
			OrbitState:                 unifieddatalibrary.String("ASCENDING"),
			Origin:                     unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:               unifieddatalibrary.String("36036"),
			OrigSensorID:               unifieddatalibrary.String("SMOS-1L"),
			PulseBandwidth:             unifieddatalibrary.Float(500.1),
			PulseDuration:              unifieddatalibrary.Float(0.000011),
			ResolutionAzimuth:          unifieddatalibrary.Float(0.123),
			ResolutionRange:            unifieddatalibrary.Float(0.123),
			RxPolarization:             unifieddatalibrary.String("H"),
			SatNo:                      unifieddatalibrary.Int(36036),
			Senalt:                     unifieddatalibrary.Float(1.1),
			SenlatEnd:                  unifieddatalibrary.Float(45.1),
			SenlatStart:                unifieddatalibrary.Float(45.1),
			SenlonEnd:                  unifieddatalibrary.Float(179.1),
			SenlonStart:                unifieddatalibrary.Float(179.1),
			Senvelx:                    unifieddatalibrary.Float(1.1),
			Senvely:                    unifieddatalibrary.Float(1.1),
			Senvelz:                    unifieddatalibrary.Float(1.1),
			SlantRange:                 unifieddatalibrary.Float(60.1),
			Snr:                        unifieddatalibrary.Float(10.1),
			SpacingAzimuth:             unifieddatalibrary.Float(0.123),
			SpacingRange:               unifieddatalibrary.Float(0.123),
			SquintAngle:                unifieddatalibrary.Float(1.2),
			SrcIDs:                     []string{"f7e01cd4-626b-441f-a423-17b160eb78ba", "223833c4-be0d-4fdb-a2e4-325a48eccced"},
			SrcTyps:                    []string{"ESID", "GROUNDIMAGE"},
			SwathLength:                unifieddatalibrary.Float(12.1),
			Tags:                       []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			Targetposx:                 unifieddatalibrary.Float(50.23),
			Targetposy:                 unifieddatalibrary.Float(50.23),
			Targetposz:                 unifieddatalibrary.Float(50.23),
			TransactionID:              unifieddatalibrary.String("TRANSACTION-ID"),
			TxPolarization:             unifieddatalibrary.String("H"),
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
