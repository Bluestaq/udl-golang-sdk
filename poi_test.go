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

func TestPoiNewWithOptionalParams(t *testing.T) {
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
	err := client.Poi.New(context.TODO(), unifieddatalibrary.PoiNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.PoiNewParamsDataModeTest,
		Name:                  "POI_NAME",
		Poiid:                 "POI-ID",
		Source:                "Bluestaq",
		Ts:                    time.Now(),
		ID:                    unifieddatalibrary.String("POI-ID"),
		Activity:              unifieddatalibrary.String("TRAINING"),
		Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
		Alt:                   unifieddatalibrary.Float(5.23),
		Andims:                unifieddatalibrary.Int(3),
		Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Asrid:                 unifieddatalibrary.Int(3),
		Asset:                 unifieddatalibrary.String("PLATFORM_NAME"),
		Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
		Atype:                 unifieddatalibrary.String("Type1"),
		Az:                    unifieddatalibrary.Float(45.23),
		BeNumber:              unifieddatalibrary.String("0427RT1030"),
		Ce:                    unifieddatalibrary.Float(10.23),
		Cntct:                 unifieddatalibrary.String("Contact Info"),
		Conf:                  unifieddatalibrary.Float(0.5),
		Desc:                  unifieddatalibrary.String("Description of the object"),
		El:                    unifieddatalibrary.Float(45.23),
		Elle:                  []float64{125.5, 85.1, 125.75},
		Env:                   unifieddatalibrary.String("SURFACE"),
		Groups:                []string{"GROUP1", "GROUP2"},
		How:                   unifieddatalibrary.String("h-g-i-g-o"),
		Ident:                 unifieddatalibrary.String("FRIEND"),
		IDWeatherReport:       []string{"WEATHER-EVENT-ID1", "WEATHER-EVENT-ID2"},
		Lat:                   unifieddatalibrary.Float(45.23),
		Le:                    unifieddatalibrary.Float(10.23),
		Lon:                   unifieddatalibrary.Float(45.23),
		Msnid:                 unifieddatalibrary.String("MSN-ID"),
		Orientation:           unifieddatalibrary.Float(45.23),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Plat:                  unifieddatalibrary.String("COMBAT_VEHICLE"),
		Pps:                   unifieddatalibrary.String("BDA"),
		Pri:                   unifieddatalibrary.Int(2),
		Spec:                  unifieddatalibrary.String("LIGHT_TANK"),
		SrcIDs:                []string{"ID1", "ID2"},
		SrcTyps:               []string{"TYPE1", "TYPE2"},
		Stale:                 unifieddatalibrary.Time(time.Now()),
		Start:                 unifieddatalibrary.Time(time.Now()),
		Tags:                  []string{"TAG1", "TAG2"},
		TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
		Trkid:                 unifieddatalibrary.String("TRK-ID"),
		Type:                  unifieddatalibrary.String("a-h-G"),
		URLs:                  []string{"URL1", "URL2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPoiListWithOptionalParams(t *testing.T) {
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
	_, err := client.Poi.List(context.TODO(), unifieddatalibrary.PoiListParams{
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

func TestPoiCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Poi.Count(context.TODO(), unifieddatalibrary.PoiCountParams{
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

func TestPoiNewBulk(t *testing.T) {
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
	err := client.Poi.NewBulk(context.TODO(), unifieddatalibrary.PoiNewBulkParams{
		Body: []unifieddatalibrary.PoiNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Name:                  "POI_NAME",
			Poiid:                 "POI-ID",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("POI-ID"),
			Activity:              unifieddatalibrary.String("TRAINING"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Alt:                   unifieddatalibrary.Float(5.23),
			Andims:                unifieddatalibrary.Int(3),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(3),
			Asset:                 unifieddatalibrary.String("PLATFORM_NAME"),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("Type1"),
			Az:                    unifieddatalibrary.Float(45.23),
			BeNumber:              unifieddatalibrary.String("0427RT1030"),
			Ce:                    unifieddatalibrary.Float(10.23),
			Cntct:                 unifieddatalibrary.String("Contact Info"),
			Conf:                  unifieddatalibrary.Float(0.5),
			Desc:                  unifieddatalibrary.String("Description of the object"),
			El:                    unifieddatalibrary.Float(45.23),
			Elle:                  []float64{125.5, 85.1, 125.75},
			Env:                   unifieddatalibrary.String("SURFACE"),
			Groups:                []string{"GROUP1", "GROUP2"},
			How:                   unifieddatalibrary.String("h-g-i-g-o"),
			Ident:                 unifieddatalibrary.String("FRIEND"),
			IDWeatherReport:       []string{"WEATHER-EVENT-ID1", "WEATHER-EVENT-ID2"},
			Lat:                   unifieddatalibrary.Float(45.23),
			Le:                    unifieddatalibrary.Float(10.23),
			Lon:                   unifieddatalibrary.Float(45.23),
			Msnid:                 unifieddatalibrary.String("MSN-ID"),
			Orientation:           unifieddatalibrary.Float(45.23),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Plat:                  unifieddatalibrary.String("COMBAT_VEHICLE"),
			Pps:                   unifieddatalibrary.String("BDA"),
			Pri:                   unifieddatalibrary.Int(2),
			Spec:                  unifieddatalibrary.String("LIGHT_TANK"),
			SrcIDs:                []string{"ID1", "ID2"},
			SrcTyps:               []string{"TYPE1", "TYPE2"},
			Stale:                 unifieddatalibrary.Time(time.Now()),
			Start:                 unifieddatalibrary.Time(time.Now()),
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Trkid:                 unifieddatalibrary.String("TRK-ID"),
			Type:                  unifieddatalibrary.String("a-h-G"),
			URLs:                  []string{"URL1", "URL2"},
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

func TestPoiGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Poi.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.PoiGetParams{
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

func TestPoiQueryhelp(t *testing.T) {
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
	_, err := client.Poi.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPoiTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Poi.Tuple(context.TODO(), unifieddatalibrary.PoiTupleParams{
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

func TestPoiUnvalidatedPublish(t *testing.T) {
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
	err := client.Poi.UnvalidatedPublish(context.TODO(), unifieddatalibrary.PoiUnvalidatedPublishParams{
		Body: []unifieddatalibrary.PoiUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Name:                  "POI_NAME",
			Poiid:                 "POI-ID",
			Source:                "Bluestaq",
			Ts:                    time.Now(),
			ID:                    unifieddatalibrary.String("POI-ID"),
			Activity:              unifieddatalibrary.String("TRAINING"),
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Alt:                   unifieddatalibrary.Float(5.23),
			Andims:                unifieddatalibrary.Int(3),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(3),
			Asset:                 unifieddatalibrary.String("PLATFORM_NAME"),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("Type1"),
			Az:                    unifieddatalibrary.Float(45.23),
			BeNumber:              unifieddatalibrary.String("0427RT1030"),
			Ce:                    unifieddatalibrary.Float(10.23),
			Cntct:                 unifieddatalibrary.String("Contact Info"),
			Conf:                  unifieddatalibrary.Float(0.5),
			Desc:                  unifieddatalibrary.String("Description of the object"),
			El:                    unifieddatalibrary.Float(45.23),
			Elle:                  []float64{125.5, 85.1, 125.75},
			Env:                   unifieddatalibrary.String("SURFACE"),
			Groups:                []string{"GROUP1", "GROUP2"},
			How:                   unifieddatalibrary.String("h-g-i-g-o"),
			Ident:                 unifieddatalibrary.String("FRIEND"),
			IDWeatherReport:       []string{"WEATHER-EVENT-ID1", "WEATHER-EVENT-ID2"},
			Lat:                   unifieddatalibrary.Float(45.23),
			Le:                    unifieddatalibrary.Float(10.23),
			Lon:                   unifieddatalibrary.Float(45.23),
			Msnid:                 unifieddatalibrary.String("MSN-ID"),
			Orientation:           unifieddatalibrary.Float(45.23),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Plat:                  unifieddatalibrary.String("COMBAT_VEHICLE"),
			Pps:                   unifieddatalibrary.String("BDA"),
			Pri:                   unifieddatalibrary.Int(2),
			Spec:                  unifieddatalibrary.String("LIGHT_TANK"),
			SrcIDs:                []string{"ID1", "ID2"},
			SrcTyps:               []string{"TYPE1", "TYPE2"},
			Stale:                 unifieddatalibrary.Time(time.Now()),
			Start:                 unifieddatalibrary.Time(time.Now()),
			Tags:                  []string{"TAG1", "TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Trkid:                 unifieddatalibrary.String("TRK-ID"),
			Type:                  unifieddatalibrary.String("a-h-G"),
			URLs:                  []string{"URL1", "URL2"},
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
