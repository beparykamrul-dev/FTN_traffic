package dataplane

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrRemoteWriteEndpoint = errors.New("invalid remote write endpoint")
	ErrRemoteWriteResponse = errors.New("remote write rejected")
)

type RemoteWriteBackend struct {
	Endpoint string
	Token    string
	Client   *http.Client
}

func (b RemoteWriteBackend) Name() string { return "remote-write" }

func (b RemoteWriteBackend) validate() (*url.URL, error) {
	u, err := url.Parse(strings.TrimSpace(b.Endpoint))
	if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
		return nil, ErrRemoteWriteEndpoint
	}
	if strings.ContainsAny(b.Endpoint, "\r\n\x00") || strings.TrimSpace(b.Token) == "" {
		return nil, ErrRemoteWriteEndpoint
	}
	return u, nil
}

func (b RemoteWriteBackend) Health(ctx context.Context) error {
	u, err := b.validate()
	if err != nil { return err }
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u.String(), nil)
	if err != nil { return err }
	client := b.Client
	if client == nil { client = &http.Client{Timeout: 10 * time.Second} }
	resp, err := client.Do(req)
	if err != nil { return err }
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 400 { return fmt.Errorf("%w: status %d", ErrRemoteWriteResponse, resp.StatusCode) }
	return nil
}

func (b RemoteWriteBackend) Publish(ctx context.Context, m Metric) error {
	u, err := b.validate()
	if err != nil { return err }
	if err := ValidateMetric(m); err != nil { return err }
	// This is a compact FTN transport envelope. A Prometheus remote-write
	// adapter can translate it at the boundary without exposing credentials
	// or customer payloads to the application model.
	body, err := json.Marshal(struct { Metrics []Metric `json:"metrics"` }{[]Metric{m}})
	if err != nil { return err }
	var compressed bytes.Buffer
	zw := gzip.NewWriter(&compressed)
	if _, err := zw.Write(body); err != nil { return err }
	if err := zw.Close(); err != nil { return err }
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), &compressed)
	if err != nil { return err }
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Authorization", "Bearer "+b.Token)
	client := b.Client
	if client == nil { client = &http.Client{Timeout: 10 * time.Second} }
	resp, err := client.Do(req)
	if err != nil { return err }
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 { return fmt.Errorf("%w: status %d", ErrRemoteWriteResponse, resp.StatusCode) }
	return nil
}
