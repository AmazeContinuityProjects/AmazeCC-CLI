// Package api is the HTTP layer for your FastAPI backend.
// Add one method per endpoint group (courses, grades, auth, …) as you grow.
// For now: a thin client + a health-check stub you can swap for a real route.
package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client talks to the academic-progress FastAPI service.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient builds a client pointed at baseURL (e.g. https://api.amazecc.com).
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// HealthResult is returned by Ping.
type HealthResult struct {
	OK     bool
	Status string
}

// Ping checks whether the API is reachable.
func (c *Client) Ping(ctx context.Context) (HealthResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return HealthResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return HealthResult{OK: false, Status: "unreachable"}, nil
	}
	defer resp.Body.Close()

	// Drain body so the connection can be reused.
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return HealthResult{OK: true, Status: "connected"}, nil
	}

	return HealthResult{
		OK:     false,
		Status: fmt.Sprintf("http %d", resp.StatusCode),
	}, nil
}
