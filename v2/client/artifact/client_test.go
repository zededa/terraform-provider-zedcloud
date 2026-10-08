package artifact

import (
	"bytes"
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
)

type recorded struct {
	method, path, contentType, contentRange string
	body                                    []byte
}

// newTestClient starts a TLS test server that records the request and answers
// with status/body, and returns a Client wired to it the same way the provider
// wires the real one (https scheme, /api base path).
func newTestClient(t *testing.T, status int, body string) (ClientService, *recorded) {
	t.Helper()
	rec := &recorded{}
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method = r.Method
		rec.path = r.URL.EscapedPath()
		rec.contentType = r.Header.Get("Content-Type")
		rec.contentRange = r.Header.Get("Content-Range")
		rec.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)

	host := strings.TrimPrefix(srv.URL, "https://")
	transport := httptransport.NewWithClient(host, "/api", []string{"https"}, srv.Client())
	return New(transport, strfmt.Default), rec
}

func TestContentRange(t *testing.T) {
	if got, want := ContentRange(70), "bytes 0-70/70"; got != want {
		t.Fatalf("ContentRange(70) = %q, want %q (exclusive end, as the UI sends)", got, want)
	}
}

func TestCreate(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"id":"0123_logo.png","name":"logo.png"}`)

	a, err := c.Create(context.Background(), "logo.png")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if a.ID != "0123_logo.png" {
		t.Errorf("id = %q", a.ID)
	}
	if rec.method != http.MethodPost || rec.path != "/api/v1/artifacts" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
	var sent map[string]string
	if err := json.Unmarshal(rec.body, &sent); err != nil || sent["name"] != "logo.png" || len(sent) != 1 {
		t.Errorf("body = %s (err %v), want {\"name\":\"logo.png\"}", rec.body, err)
	}
	if !strings.HasPrefix(rec.contentType, "application/json") {
		t.Errorf("content-type = %q", rec.contentType)
	}
}

func TestCreate_NoID(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"name":"logo.png"}`)
	if _, err := c.Create(context.Background(), "logo.png"); err == nil {
		t.Fatal("expected an error for a response without id")
	}
}

func TestUpload(t *testing.T) {
	c, rec := newTestClient(t, http.StatusAccepted, `{}`)
	data := []byte{0x89, 'P', 'N', 'G', 0x00}

	if err := c.Upload(context.Background(), "0123_logo.png", data); err != nil {
		t.Fatalf("Upload: %v", err)
	}
	if rec.method != http.MethodPut || rec.path != "/api/v1/artifacts/id/0123_logo.png/upload/chunked" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
	if rec.contentType != "application/octet-stream" {
		t.Errorf("content-type = %q, want application/octet-stream", rec.contentType)
	}
	if rec.contentRange != "bytes 0-5/5" {
		t.Errorf("content-range = %q, want bytes 0-5/5", rec.contentRange)
	}
	if !bytes.Equal(rec.body, data) {
		t.Errorf("body = %x, want raw bytes %x", rec.body, data)
	}
}

func TestUpload_Errors(t *testing.T) {
	c, _ := newTestClient(t, http.StatusInternalServerError, `{"httpStatusMsg":"boom"}`)
	err := c.Upload(context.Background(), "0123_logo.png", []byte("x"))
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Status != 500 || !strings.Contains(apiErr.Body, "boom") {
		t.Fatalf("err = %v, want APIError 500 carrying the body", err)
	}
	if err := c.Upload(context.Background(), "0123_logo.png", nil); err == nil {
		t.Error("expected an error for empty content")
	}
}

func TestGetSignedURL(t *testing.T) {
	cases := []struct {
		name     string
		status   int
		body     string
		notFound bool
		wantErr  bool
	}{
		{name: "ok", status: 200, body: `{"signedUrl":"https://s3/x","ttl":"10"}`},
		{name: "307 without Location is not found", status: 307, body: `{"httpStatusCode":307}`, notFound: true},
		{name: "404 is not found", status: 404, body: `{}`, notFound: true},
		{name: "200 without url is not found", status: 200, body: `{}`, notFound: true},
		{name: "500 is an error", status: 500, body: `{}`, wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, rec := newTestClient(t, tc.status, tc.body)
			a, err := c.GetSignedURL(context.Background(), "0123_logo.png")
			if rec.method != http.MethodGet || rec.path != "/api/v1/artifacts/id/0123_logo.png/url" {
				t.Errorf("request = %s %s", rec.method, rec.path)
			}
			switch {
			case tc.notFound:
				if !errors.Is(err, ErrNotFound) {
					t.Fatalf("err = %v, want ErrNotFound", err)
				}
			case tc.wantErr:
				if err == nil || errors.Is(err, ErrNotFound) {
					t.Fatalf("err = %v, want a non-ErrNotFound error", err)
				}
			default:
				if err != nil || a.SignedURL != "https://s3/x" {
					t.Fatalf("got %+v, %v", a, err)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	for _, status := range []int{200, 404} {
		c, rec := newTestClient(t, status, `{}`)
		if err := c.Delete(context.Background(), "0123_logo.png"); err != nil {
			t.Errorf("status %d: %v", status, err)
		}
		if rec.method != http.MethodDelete || rec.path != "/api/v1/artifacts/id/0123_logo.png" {
			t.Errorf("request = %s %s", rec.method, rec.path)
		}
	}
	c, _ := newTestClient(t, 500, `{}`)
	if err := c.Delete(context.Background(), "0123_logo.png"); err == nil {
		t.Error("expected an error for 500")
	}
}
