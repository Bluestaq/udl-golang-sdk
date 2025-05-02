// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestReportAndActivityUdlSigactFileGetWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	resp, err := client.ReportAndActivity.UdlSigact.FileGet(
		context.TODO(),
		"id",
		unifieddatalibrary.ReportAndActivityUdlSigactFileGetParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestReportAndActivityUdlSigactUnvalidatedPublish(t *testing.T) {
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
	err := client.ReportAndActivity.UdlSigact.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ReportAndActivityUdlSigactUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ReportAndActivityUdlSigactUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			ReportDate:            time.Now(),
			Source:                "Bluestaq",
			ID:                    unifieddatalibrary.String("SIGACT-ID"),
			Accuracy:              unifieddatalibrary.Int(13),
			Actors:                []string{"US", "CAN"},
			Agjson:                unifieddatalibrary.String(`{"type":"Polygon","coordinates":[ [ [ 67.3291113966927, 26.156175339112 ], [ 67.2580009640721, 26.091022064271 ], [ 67.1795862381682, 26.6637992964562 ], [ 67.2501237475598, 26.730115808233 ], [ 67.3291113966927, 26.156175339112 ] ] ] }`),
			Andims:                unifieddatalibrary.Int(3),
			Area:                  unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Asrid:                 unifieddatalibrary.Int(3),
			Atext:                 unifieddatalibrary.String("POLYGON((67.3291113966927 26.156175339112,67.2580009640721 26.091022064271,67.1795862381682 26.6637992964562,67.2501237475598 26.730115808233,67.3291113966927 26.156175339112))"),
			Atype:                 unifieddatalibrary.String("Type1"),
			AvgTone:               unifieddatalibrary.Float(8.23),
			CameoBaseCode:         unifieddatalibrary.String("Example_cameoBaseCode"),
			CameoCode:             unifieddatalibrary.String("CAMEO_CODE"),
			CameoRootCode:         unifieddatalibrary.String("Example_cameoRootCode"),
			ChecksumValue:         unifieddatalibrary.String("120EA8A25E5D487BF68B5F7096440019"),
			City:                  unifieddatalibrary.String("Austin"),
			CivAbd:                unifieddatalibrary.Int(423),
			CivDet:                unifieddatalibrary.Int(234),
			CivKia:                unifieddatalibrary.Int(2),
			CivWound:              unifieddatalibrary.Int(123),
			Clarity:               unifieddatalibrary.Int(1),
			CoalAbd:               unifieddatalibrary.Int(123),
			CoalDet:               unifieddatalibrary.Int(123),
			CoalKia:               unifieddatalibrary.Int(123),
			CoalWound:             unifieddatalibrary.Int(123),
			ComplexAttack:         unifieddatalibrary.Bool(false),
			Confidence:            unifieddatalibrary.Int(13),
			CountryCode:           unifieddatalibrary.String("US"),
			District:              unifieddatalibrary.String("district 12"),
			DocumentFilename:      unifieddatalibrary.String("Example_documentFilename"),
			DocumentSource:        unifieddatalibrary.String("Example_documentSource"),
			EnemyAbd:              unifieddatalibrary.Int(123),
			EnemyDet:              unifieddatalibrary.Int(123),
			EnemyKia:              unifieddatalibrary.Int(123),
			EventDescription:      unifieddatalibrary.String("Example_Description"),
			EventEnd:              unifieddatalibrary.Time(time.Now()),
			EventStart:            unifieddatalibrary.Time(time.Now()),
			EventType:             unifieddatalibrary.String("Military"),
			Filesize:              unifieddatalibrary.Int(0),
			FriendlyAbd:           unifieddatalibrary.Int(123),
			FriendlyDet:           unifieddatalibrary.Int(123),
			FriendlyKia:           unifieddatalibrary.Int(123),
			FriendlyWound:         unifieddatalibrary.Int(123),
			Goldstein:             unifieddatalibrary.Float(9.32),
			HasAttachment:         unifieddatalibrary.Bool(true),
			HostNatAbd:            unifieddatalibrary.Int(123),
			HostNatDet:            unifieddatalibrary.Int(123),
			HostNatKia:            unifieddatalibrary.Int(123),
			HostNatWound:          unifieddatalibrary.Int(123),
			IDNumber:              unifieddatalibrary.String("NUMBER-ID"),
			Lat:                   unifieddatalibrary.Float(45.31),
			Lon:                   unifieddatalibrary.Float(90.23),
			Milgrid:               unifieddatalibrary.String("4QFJ12345678"),
			Notes:                 unifieddatalibrary.String("Example_Notes"),
			NumArticles:           unifieddatalibrary.Int(8),
			NumMentions:           unifieddatalibrary.Int(8),
			NumSources:            unifieddatalibrary.Int(8),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Province:              unifieddatalibrary.String("Province_Example"),
			RelatedDocs: []unifieddatalibrary.ReportAndActivityUdlSigactUnvalidatedPublishParamsBodyRelatedDoc{{
				DataSourceRefs: []unifieddatalibrary.ReportAndActivityUdlSigactUnvalidatedPublishParamsBodyRelatedDocDataSourceRef{{
					DataSourceID:    unifieddatalibrary.String("dataSourceId"),
					EndPosition:     unifieddatalibrary.String("endPosition"),
					ParagraphNumber: unifieddatalibrary.String("paragraphNumber"),
					SentenceNumber:  unifieddatalibrary.String("sentenceNumber"),
					StartPosition:   unifieddatalibrary.String("startPosition"),
				}},
				DocumentID: unifieddatalibrary.String("documentId"),
			}},
			RepUnit:         unifieddatalibrary.String("Unit_1"),
			RepUnitActivity: unifieddatalibrary.String("Example_Activity"),
			RepUnitType:     unifieddatalibrary.String("Example_repUnitType"),
			SideAAbd:        unifieddatalibrary.Int(123),
			SideADet:        unifieddatalibrary.Int(123),
			SideAkia:        unifieddatalibrary.Int(123),
			SideAWound:      unifieddatalibrary.Int(123),
			SideBAbd:        unifieddatalibrary.Int(123),
			SideBDet:        unifieddatalibrary.Int(123),
			SideBkia:        unifieddatalibrary.Int(123),
			SideBWound:      unifieddatalibrary.Int(123),
			SourceLanguage:  unifieddatalibrary.String("eng"),
			SourceURL:       unifieddatalibrary.String("Example_URL"),
			Summary:         unifieddatalibrary.String("Example_Summary"),
			Target:          unifieddatalibrary.String("US"),
			Theater:         unifieddatalibrary.String("Kabul"),
			TypeOfAttack:    unifieddatalibrary.String("IED Explosion"),
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
