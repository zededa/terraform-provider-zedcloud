package application_instance_v2

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

type recorded struct {
	method, path string
	body         []byte
}

func newTestClient(t *testing.T, status int, body string) (ClientService, *recorded) {
	t.Helper()
	rec := &recorded{}
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method = r.Method
		rec.path = r.URL.EscapedPath()
		rec.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	host := strings.TrimPrefix(srv.URL, "https://")
	return New(httptransport.NewWithClient(host, "/api", []string{"https"}, srv.Client()), strfmt.Default), rec
}

func strp(s string) *string { return &s }

func TestCreate_JobIDInObjectID(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"objectId":"job-1","objectName":"inst_create_job-1"}`)
	jobID, err := c.Create(context.Background(), &models.AppInstance{ID: "id-1", Name: strp("inst")})
	if err != nil || jobID != "job-1" {
		t.Fatalf("jobID = %q, err = %v", jobID, err)
	}
	if rec.method != http.MethodPost || rec.path != "/api/v2/apps/instances" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
	var sent map[string]interface{}
	if err := json.Unmarshal(rec.body, &sent); err != nil || sent["id"] != "id-1" {
		t.Errorf("body = %s, want the caller-supplied id", rec.body)
	}
}

func TestCreate_JobIDInAppInstanceServiceResp(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"jobID":"job-2","result":{"objectId":"ignored"}}`)
	jobID, err := c.Create(context.Background(), &models.AppInstance{ID: "id-1"})
	if err != nil || jobID != "job-2" {
		t.Fatalf("jobID = %q, err = %v", jobID, err)
	}
}

func TestCreate_NoJobID(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{}`)
	if _, err := c.Create(context.Background(), &models.AppInstance{ID: "id-1"}); err == nil {
		t.Fatal("expected an error for a response without a job id")
	}
}

func TestCreate_BadRequestCarriesDetails(t *testing.T) {
	c, _ := newTestClient(t, http.StatusBadRequest,
		`{"httpStatusCode":400,"error":[{"details":"project id is empty"}]}`)
	_, err := c.Create(context.Background(), &models.AppInstance{ID: "id-1"})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code() != 400 || apiErr.GetPayload() == nil ||
		!strings.Contains(err.Error(), "project id is empty") {
		t.Fatalf("err = %v", err)
	}
}

func TestUpdate(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"objectId":"job-3"}`)
	jobID, err := c.Update(context.Background(), "id-1", &models.AppInstance{ID: "id-1"})
	if err != nil || jobID != "job-3" {
		t.Fatalf("jobID = %q, err = %v", jobID, err)
	}
	if rec.method != http.MethodPut || rec.path != "/api/v2/apps/instances/id/id-1" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
}

func TestUpdate_IDMismatch(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"objectId":"job-3"}`)
	if _, err := c.Update(context.Background(), "id-1", &models.AppInstance{ID: "id-2"}); err == nil {
		t.Fatal("expected an error: the controller rejects a body id that differs from the path id")
	}
	if rec.method != "" {
		t.Errorf("no request should be sent, got %s %s", rec.method, rec.path)
	}
}

func TestDelete_NotFound(t *testing.T) {
	c, rec := newTestClient(t, http.StatusNotFound, `{"httpStatusCode":404}`)
	_, err := c.Delete(context.Background(), "id-1")
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code() != http.StatusNotFound {
		t.Fatalf("err = %v, want a 404 APIError", err)
	}
	if rec.method != http.MethodDelete || rec.path != "/api/v2/apps/instances/id/id-1" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
}
