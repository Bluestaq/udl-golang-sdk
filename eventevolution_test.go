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

func TestEventEvolutionNewWithOptionalParams(t *testing.T) {
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
	err := client.EventEvolution.New(context.TODO(), unifieddatalibrary.EventEvolutionNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.EventEvolutionNewParamsDataModeTest,
		EventID:               "EVENT_ID",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		Summary:               "Example summary of the event.",
		ID:                    unifieddatalibrary.String("EVENT_EVOL_ID"),
		Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
		Andims:                unifieddatalibrary.Int(2),
		Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Asrid:                 unifieddatalibrary.Int(4326),
		Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Atype:                 unifieddatalibrary.String("POLYGON"),
		Category:              unifieddatalibrary.String("PROTEST"),
		CountryCode:           unifieddatalibrary.String("US"),
		DataDescription:       unifieddatalibrary.String("Description of relationship between srcTyps and srcIds"),
		EndTime:               unifieddatalibrary.Time(time.Now()),
		GeoAdminLevel1:        unifieddatalibrary.String("Colorado"),
		GeoAdminLevel2:        unifieddatalibrary.String("El Paso County"),
		GeoAdminLevel3:        unifieddatalibrary.String("Colorado Springs"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Redact:                unifieddatalibrary.Bool(false),
		SrcIDs:                []string{"SRC_ID_1", "SRC_ID_2"},
		SrcTyps:               []string{"AIS", "CONJUNCTION"},
		Status:                unifieddatalibrary.String("UNKNOWN"),
		Tags:                  []string{"TAG1", "TAG2"},
		URL:                   []string{"URL1", "URL2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventEvolutionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EventEvolution.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.EventEvolutionGetParams{
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

func TestEventEvolutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.EventEvolution.List(context.TODO(), unifieddatalibrary.EventEvolutionListParams{
		EventID:     unifieddatalibrary.String("eventId"),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventEvolutionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.EventEvolution.Count(context.TODO(), unifieddatalibrary.EventEvolutionCountParams{
		EventID:     unifieddatalibrary.String("eventId"),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventEvolutionNewBulk(t *testing.T) {
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
	err := client.EventEvolution.NewBulk(context.TODO(), unifieddatalibrary.EventEvolutionNewBulkParams{
		Body: []unifieddatalibrary.EventEvolutionNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EventID:               "EVENT_ID",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			Summary:               "Example summary of the event.",
			ID:                    unifieddatalibrary.String("EVENT_EVOL_ID"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Andims:                unifieddatalibrary.Int(2),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(4326),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("POLYGON"),
			Category:              unifieddatalibrary.String("PROTEST"),
			CountryCode:           unifieddatalibrary.String("US"),
			DataDescription:       unifieddatalibrary.String("Description of relationship between srcTyps and srcIds"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			GeoAdminLevel1:        unifieddatalibrary.String("Colorado"),
			GeoAdminLevel2:        unifieddatalibrary.String("El Paso County"),
			GeoAdminLevel3:        unifieddatalibrary.String("Colorado Springs"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Redact:                unifieddatalibrary.Bool(false),
			SrcIDs:                []string{"SRC_ID_1", "SRC_ID_2"},
			SrcTyps:               []string{"AIS", "CONJUNCTION"},
			Status:                unifieddatalibrary.String("UNKNOWN"),
			Tags:                  []string{"TAG1", "TAG2"},
			URL:                   []string{"URL1", "URL2"},
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

func TestEventEvolutionQueryhelp(t *testing.T) {
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
	_, err := client.EventEvolution.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventEvolutionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.EventEvolution.Tuple(context.TODO(), unifieddatalibrary.EventEvolutionTupleParams{
		Columns:     "columns",
		EventID:     unifieddatalibrary.String("eventId"),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventEvolutionUnvalidatedPublish(t *testing.T) {
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
	err := client.EventEvolution.UnvalidatedPublish(context.TODO(), unifieddatalibrary.EventEvolutionUnvalidatedPublishParams{
		Body: []unifieddatalibrary.EventEvolutionUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EventID:               "EVENT_ID",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			Summary:               "Example summary of the event.",
			ID:                    unifieddatalibrary.String("EVENT_EVOL_ID"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Andims:                unifieddatalibrary.Int(2),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(4326),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("POLYGON"),
			Category:              unifieddatalibrary.String("PROTEST"),
			CountryCode:           unifieddatalibrary.String("US"),
			DataDescription:       unifieddatalibrary.String("Description of relationship between srcTyps and srcIds"),
			EndTime:               unifieddatalibrary.Time(time.Now()),
			GeoAdminLevel1:        unifieddatalibrary.String("Colorado"),
			GeoAdminLevel2:        unifieddatalibrary.String("El Paso County"),
			GeoAdminLevel3:        unifieddatalibrary.String("Colorado Springs"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Redact:                unifieddatalibrary.Bool(false),
			SrcIDs:                []string{"SRC_ID_1", "SRC_ID_2"},
			SrcTyps:               []string{"AIS", "CONJUNCTION"},
			Status:                unifieddatalibrary.String("UNKNOWN"),
			Tags:                  []string{"TAG1", "TAG2"},
			URL:                   []string{"URL1", "URL2"},
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
