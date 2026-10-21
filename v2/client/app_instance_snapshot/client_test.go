package app_instance_snapshot

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
)

type recorded struct {
	method, path, query string
	body                []byte
}

func newTestClient(t *testing.T, status int, body string) (ClientService, *recorded) {
	t.Helper()
	rec := &recorded{}
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method = r.Method
		rec.path = r.URL.EscapedPath()
		rec.query = r.URL.RawQuery
		rec.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	host := strings.TrimPrefix(srv.URL, "https://")
	return New(httptransport.NewWithClient(host, "/api", []string{"https"}, srv.Client()), strfmt.Default), rec
}

func TestCreate(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK,
		`{"httpStatusCode":200,"objectId":"snap-1","objectName":"snapshot-20261021120000-abcdef12"}`)
	res, err := c.Create(context.Background(), "app-1")
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "snap-1" || res.Name != "snapshot-20261021120000-abcdef12" {
		t.Errorf("result = %+v", res)
	}
	if rec.method != http.MethodPost || rec.path != "/api/v1/apps/instances/id/app-1/snapshot" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
	if strings.TrimSpace(string(rec.body)) != "{}" {
		t.Errorf("body = %q, want {}", rec.body)
	}
}

func TestCreate_OldControllerWithoutObjectID(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"httpStatusCode":200}`)
	res, err := c.Create(context.Background(), "app-1")
	if err != nil || res.ID != "" || res.Name != "" {
		t.Fatalf("res = %+v, err = %v", res, err)
	}
}

func TestCreate_ErrorCarriesDetails(t *testing.T) {
	c, _ := newTestClient(t, http.StatusPreconditionFailed,
		`{"httpStatusCode":412,"error":[{"details":"snapshots are not allowed for the devices with the kubevirt"}]}`)
	_, err := c.Create(context.Background(), "app-1")
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code() != 412 || apiErr.GetPayload() == nil ||
		!strings.Contains(err.Error(), "kubevirt") {
		t.Fatalf("err = %v", err)
	}
}

func TestGet(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK,
		`{"id":"snap-1","name":"s1","config":{"id":"app-1","name":"inst","encryptedSecrets":{"k":"v"}}}`)
	s, err := c.Get(context.Background(), "snap-1")
	if err != nil {
		t.Fatal(err)
	}
	if rec.method != http.MethodGet || rec.path != "/api/v1/apps/instances/snapshot/id/snap-1" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
	if s.ID != "snap-1" || s.Name != "s1" || s.AppInstanceID() != "app-1" {
		t.Errorf("snapshot = %+v, app instance %q", s, s.AppInstanceID())
	}
	cfg, err := s.ConfigJSON()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(cfg, "encryptedSecrets") {
		t.Errorf("config_json must not carry encryptedSecrets: %s", cfg)
	}
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(cfg), &m); err != nil || m["name"] != "inst" {
		t.Errorf("config_json = %s", cfg)
	}
}

func TestGet_NotFound(t *testing.T) {
	c, _ := newTestClient(t, http.StatusNotFound, `{"httpStatusCode":404}`)
	if _, err := c.Get(context.Background(), "snap-1"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestGetByName(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"id":"snap-1","name":"s1"}`)
	if _, err := c.GetByName(context.Background(), "app-1", "s1"); err != nil {
		t.Fatal(err)
	}
	if rec.path != "/api/v1/apps/instances/id/app-1/snapshot/name/s1" {
		t.Errorf("path = %s", rec.path)
	}
}

func TestList(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"snapshots":[{"id":"a","name":"n1"},{"id":"b","name":"n2"}]}`)
	l, err := c.List(context.Background(), "app-1")
	if err != nil {
		t.Fatal(err)
	}
	if rec.path != "/api/v1/apps/instances/id/app-1/snapshots" || len(l) != 2 || l[1].Name != "n2" {
		t.Errorf("path %s, list %+v", rec.path, l)
	}
}

func TestList_Empty(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{}`)
	l, err := c.List(context.Background(), "app-1")
	if err != nil || len(l) != 0 {
		t.Fatalf("list = %v, err = %v", l, err)
	}
}

func TestState(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"state":"CREATED","timestamp":"2026-10-21T12:00:00Z"}`)
	s, err := c.State(context.Background(), "snap-1")
	if err != nil {
		t.Fatal(err)
	}
	if rec.path != "/api/v1/apps/instances/snapshot/id/snap-1/state" || s.Name() != StateCreated || s.Timestamp == "" {
		t.Errorf("path %s, state %+v", rec.path, s)
	}
}

func TestState_ZeroValueIsInvalid(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{"timestamp":"2026-10-21T12:00:00Z"}`)
	s, err := c.State(context.Background(), "snap-1")
	if err != nil || s.Name() != StateInvalid {
		t.Fatalf("state = %+v, err = %v", s, err)
	}
}

func TestActivate(t *testing.T) {
	for _, tc := range []struct {
		force bool
		query string
	}{{false, ""}, {true, "force=true"}} {
		c, rec := newTestClient(t, http.StatusOK, `{"httpStatusCode":200}`)
		if err := c.Activate(context.Background(), "app-1", "s1", tc.force); err != nil {
			t.Fatal(err)
		}
		if rec.method != http.MethodPut || rec.path != "/api/v1/apps/instances/id/app-1/snapshot/name/s1/activate" ||
			rec.query != tc.query {
			t.Errorf("force=%v: request = %s %s?%s", tc.force, rec.method, rec.path, rec.query)
		}
	}
}

func TestDelete(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK, `{"httpStatusCode":200}`)
	if err := c.Delete(context.Background(), "app-1", "s1"); err != nil {
		t.Fatal(err)
	}
	if rec.method != http.MethodDelete || rec.path != "/api/v1/apps/instances/id/app-1/snapshot/name/s1" {
		t.Errorf("request = %s %s", rec.method, rec.path)
	}
}

func TestDelete_MissingIsSuccess(t *testing.T) {
	c, _ := newTestClient(t, http.StatusNotFound, `{"httpStatusCode":404}`)
	if err := c.Delete(context.Background(), "app-1", "s1"); err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestDeviceRawStatus(t *testing.T) {
	c, rec := newTestClient(t, http.StatusOK,
		`{"id":"dev-1","rawStatus":"{\"apiCapability\":\"API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE\"}"}`)
	raw, err := c.DeviceRawStatus(context.Background(), "dev-1")
	if err != nil {
		t.Fatal(err)
	}
	if rec.path != "/api/v1/devices/id/dev-1/status" || !strings.Contains(raw, ImmediateSnapshotCapability) {
		t.Errorf("path %s, raw %q", rec.path, raw)
	}
}

func TestEmptyArguments(t *testing.T) {
	c, _ := newTestClient(t, http.StatusOK, `{}`)
	ctx := context.Background()
	if _, err := c.Create(ctx, ""); err == nil {
		t.Error("Create with an empty id must fail")
	}
	if _, err := c.Get(ctx, ""); err == nil {
		t.Error("Get with an empty id must fail")
	}
	if err := c.Activate(ctx, "app-1", "", false); err == nil {
		t.Error("Activate with an empty name must fail")
	}
	if err := c.Delete(ctx, "", "s1"); err == nil {
		t.Error("Delete with an empty app instance id must fail")
	}
	if _, err := c.DeviceRawStatus(ctx, ""); err == nil {
		t.Error("DeviceRawStatus with an empty id must fail")
	}
}

func TestSupportsImmediateSnapshots(t *testing.T) {
	for raw, want := range map[string]bool{
		`{"apiCapability":"API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE"}`: true,
		`{"apiCapability":"API_CAPABILITY_REPORT_TPM_EVENTLOG"}`:        true, // 21: levels are cumulative
		`{"apiCapability":"API_CAPABILITY_SOMETHING_FROM_THE_FUTURE"}`:  true, // newer than the table
		`{"apiCapability":"API_CAPABILITY_WIN_LIC_PASSTHROUGH"}`:        false,
		`{"apiCapability":13}`: true,
		`{"apiCapability":5}`:  false,
		`{"api_capability":"API_CAPABILITY_ENCRYPTED_PATCH_ENVELOPE"}`: true,
		`{}`: false,
	} {
		got, reported, err := SupportsImmediateSnapshots(raw)
		if err != nil || got != want {
			t.Errorf("%s: got %v (%q), err %v, want %v", raw, got, reported, err, want)
		}
	}
	if _, _, err := SupportsImmediateSnapshots("not json"); err == nil {
		t.Error("invalid JSON must be an error")
	}
}
