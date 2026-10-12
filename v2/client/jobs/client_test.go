package jobs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func newTestClient(t *testing.T, status int, body string) (ClientService, *string) {
	t.Helper()
	var path string
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.Method + " " + r.URL.EscapedPath()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	host := strings.TrimPrefix(srv.URL, "https://")
	return New(httptransport.NewWithClient(host, "/api", []string{"https"}, srv.Client()), strfmt.Default), &path
}

func TestGet(t *testing.T) {
	c, path := newTestClient(t, http.StatusOK, `{"id":"j1","name":"x_create_j1","status":"JOB_STATUS_FAILED",
		"errorDetailOfJob":"boom","individualOpStatus":{"x":"name already exists","y":"OBJECT_INDIVIDUAL_OP_STATUS_SUCCESS"}}`)
	j, err := c.Get(context.Background(), "j1", ObjectTypeEdgeAppInstance)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if *path != "GET /api/v1/jobs/id/j1/objectType/OBJECT_TYPE_EDGE_APP_INSTANCE" {
		t.Errorf("request = %s", *path)
	}
	if j.Status != StatusFailed || !j.Status.Final() {
		t.Errorf("status = %s", j.Status)
	}
	if got := j.ErrorText(); !strings.Contains(got, "boom") || !strings.Contains(got, "x: name already exists") ||
		strings.Contains(got, "y:") {
		t.Errorf("ErrorText = %q", got)
	}
}

func TestGet_UnspecifiedStatusIsOmitted(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"id":"j1"}`)
	j, err := c.Get(context.Background(), "j1", ObjectTypeEdgeAppInstance)
	if err != nil || j.Status != "" || j.Status.Final() {
		t.Fatalf("job = %+v, err = %v", j, err)
	}
}

func TestGet_NoJobInResponse(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"httpStatusCode":200}`)
	if _, err := c.Get(context.Background(), "j1", ObjectTypeEdgeAppInstance); err == nil {
		t.Fatal("expected an error when the response carries no job")
	}
}

func TestGet_Forbidden(t *testing.T) {
	c, _ := newTestClient(t, http.StatusForbidden, `{"httpStatusCode":403}`)
	_, err := c.Get(context.Background(), "j1", ObjectTypeEdgeAppInstance)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code() != http.StatusForbidden {
		t.Fatalf("err = %v", err)
	}
}
