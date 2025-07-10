// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/shared"
)

func TestScFileGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Scs.File.Get(context.TODO(), unifieddatalibrary.ScFileGetParams{
		ID:          "id",
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

func TestScFileUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Scs.File.Update(context.TODO(), unifieddatalibrary.ScFileUpdateParams{
		FileDataList: []shared.FileDataParam{{
			ID: unifieddatalibrary.String("/example/folder/example_file.txt"),
			Attributes: shared.FileDataAttributesParam{
				ID:                    unifieddatalibrary.String("id"),
				Classification:        unifieddatalibrary.String("U"),
				ClassificationMarking: unifieddatalibrary.String("classificationMarking"),
				CreatedBy:             unifieddatalibrary.String("createdBy"),
				CreatedDate:           unifieddatalibrary.String("createdDate"),
				DeleteOn:              unifieddatalibrary.Int(0),
				Description:           unifieddatalibrary.String("A new Example description"),
				DocTitle:              unifieddatalibrary.String("docTitle"),
				DocType:               unifieddatalibrary.String("docType"),
				Doi:                   []string{"string"},
				EllipseLat:            unifieddatalibrary.Float(0),
				EllipseLon:            unifieddatalibrary.Float(0),
				FileName:              unifieddatalibrary.String("fileName"),
				IntrinsicTitle:        unifieddatalibrary.String("intrinsicTitle"),
				Keywords:              unifieddatalibrary.String("keywords"),
				MediaTitle:            unifieddatalibrary.String("mediaTitle"),
				MetaInfo:              unifieddatalibrary.String("A new Example metaInfo"),
				Milgrid:               unifieddatalibrary.String("milgrid"),
				MilgridLat:            unifieddatalibrary.Float(0),
				MilgridLon:            unifieddatalibrary.Float(0),
				ModifiedBy:            unifieddatalibrary.String("modifiedBy"),
				ModifiedDate:          unifieddatalibrary.String("modifiedDate"),
				Name:                  unifieddatalibrary.String("name"),
				Path:                  unifieddatalibrary.String("path"),
				Read:                  unifieddatalibrary.String("read"),
				Searchable:            unifieddatalibrary.Bool(true),
				SearchAfter:           unifieddatalibrary.String("searchAfter"),
				SerialNumber:          unifieddatalibrary.String("serialNumber"),
				Size:                  unifieddatalibrary.Int(0),
				Tags:                  []string{"string"},
				Write:                 unifieddatalibrary.String("write"),
			},
			ContentAction: shared.FileDataContentActionUpdate,
			TargetName:    unifieddatalibrary.String("targetName"),
			TargetPath:    unifieddatalibrary.String("targetPath"),
			Type:          shared.FileDataTypeFile,
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

func TestScFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Scs.File.List(context.TODO(), unifieddatalibrary.ScFileListParams{
		Path:        "path",
		Count:       unifieddatalibrary.Int(0),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		Offset:      unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
