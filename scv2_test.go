// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestScV2UpdateWithOptionalParams(t *testing.T) {
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
	err := client.Scs.V2.Update(context.TODO(), unifieddatalibrary.ScV2UpdateParams{
		Path:             "path",
		SendNotification: unifieddatalibrary.Bool(true),
		ID:               unifieddatalibrary.String("/my-folder/"),
		Attachment: unifieddatalibrary.ScV2UpdateParamsAttachment{
			Author:        unifieddatalibrary.String("John.Doe"),
			ContentLength: unifieddatalibrary.Int(0),
			ContentType:   unifieddatalibrary.String("text/plain"),
			Date:          unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
			Keywords:      unifieddatalibrary.String("keywords"),
			Language:      unifieddatalibrary.String("en"),
			Title:         unifieddatalibrary.String("title"),
		},
		ClassificationMarking: unifieddatalibrary.String("U"),
		CreatedAt:             unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
		CreatedBy:             unifieddatalibrary.String("John.Doe"),
		DeleteOn:              unifieddatalibrary.Int(0),
		Description:           unifieddatalibrary.String("A description of the updated folder."),
		Filename:              unifieddatalibrary.String("my-folder"),
		FilePath:              unifieddatalibrary.String("/my-folder/sub-folder/"),
		Keywords:              unifieddatalibrary.String("keywords"),
		ParentPath:            unifieddatalibrary.String("/"),
		PathType:              unifieddatalibrary.ScV2UpdateParamsPathTypeFile,
		ReadACL:               unifieddatalibrary.String("user.id1,group.id1"),
		Size:                  unifieddatalibrary.Int(0),
		Tags:                  []string{"TAG1", "TAG2"},
		UpdatedAt:             unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
		UpdatedBy:             unifieddatalibrary.String("John.Doe"),
		WriteACL:              unifieddatalibrary.String("user.id1,group.id1"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScV2ListWithOptionalParams(t *testing.T) {
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
	_, err := client.Scs.V2.List(context.TODO(), unifieddatalibrary.ScV2ListParams{
		Path:        "path",
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
		Order:       unifieddatalibrary.String("order"),
		SearchAfter: unifieddatalibrary.String("searchAfter"),
		Size:        unifieddatalibrary.Int(0),
		Sort:        unifieddatalibrary.String("sort"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScV2Delete(t *testing.T) {
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
	err := client.Scs.V2.Delete(context.TODO(), unifieddatalibrary.ScV2DeleteParams{
		Path: "path",
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScV2Copy(t *testing.T) {
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
	err := client.Scs.V2.Copy(context.TODO(), unifieddatalibrary.ScV2CopyParams{
		FromPath: "fromPath",
		ToPath:   "toPath",
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScV2FileUploadWithOptionalParams(t *testing.T) {
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
	err := client.Scs.V2.FileUpload(
		context.TODO(),
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		unifieddatalibrary.ScV2FileUploadParams{
			ClassificationMarking: "classificationMarking",
			Path:                  "path",
			DeleteAfter:           unifieddatalibrary.String("deleteAfter"),
			Description:           unifieddatalibrary.String("description"),
			Overwrite:             unifieddatalibrary.Bool(true),
			SendNotification:      unifieddatalibrary.Bool(true),
			Tags:                  unifieddatalibrary.String("tags"),
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

func TestScV2FolderNewWithOptionalParams(t *testing.T) {
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
	err := client.Scs.V2.FolderNew(context.TODO(), unifieddatalibrary.ScV2FolderNewParams{
		Path:             "path",
		SendNotification: unifieddatalibrary.Bool(true),
		ID:               unifieddatalibrary.String("/my-folder/"),
		Attachment: unifieddatalibrary.ScV2FolderNewParamsAttachment{
			Author:        unifieddatalibrary.String("John.Doe"),
			ContentLength: unifieddatalibrary.Int(0),
			ContentType:   unifieddatalibrary.String("text/plain"),
			Date:          unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
			Keywords:      unifieddatalibrary.String("keywords"),
			Language:      unifieddatalibrary.String("en"),
			Title:         unifieddatalibrary.String("title"),
		},
		ClassificationMarking: unifieddatalibrary.String("U"),
		CreatedAt:             unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
		CreatedBy:             unifieddatalibrary.String("John.Doe"),
		DeleteOn:              unifieddatalibrary.Int(0),
		Description:           unifieddatalibrary.String("My first folder"),
		Filename:              unifieddatalibrary.String("my-folder"),
		FilePath:              unifieddatalibrary.String("/my-folder/sub-folder/"),
		Keywords:              unifieddatalibrary.String("keywords"),
		ParentPath:            unifieddatalibrary.String("/"),
		PathType:              unifieddatalibrary.ScV2FolderNewParamsPathTypeFile,
		ReadACL:               unifieddatalibrary.String("user.id1,group.id1"),
		Size:                  unifieddatalibrary.Int(0),
		Tags:                  []string{"TAG1", "TAG2"},
		UpdatedAt:             unifieddatalibrary.String("2025-07-03T16:27:57.970Z"),
		UpdatedBy:             unifieddatalibrary.String("John.Doe"),
		WriteACL:              unifieddatalibrary.String("user.id1,group.id1"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScV2Move(t *testing.T) {
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
	err := client.Scs.V2.Move(context.TODO(), unifieddatalibrary.ScV2MoveParams{
		FromPath: "fromPath",
		ToPath:   "toPath",
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
