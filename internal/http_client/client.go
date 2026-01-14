package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    http.Header
}

type Option func(*Client)

type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

func New(baseURL string, opts ...Option) *Client {
	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(http.Header),
	}

	client.headers.Set("Content-Type", "application/json")

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(hc *Client) {
		if httpClient != nil {
			hc.httpClient = httpClient
		}
	}
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		if key != "" && value != "" {
			c.headers.Set(key, value)
		}
	}
}

func WithBearerToken(token string) Option {
	return func(c *Client) {
		if token != "" {
			c.headers.Set("Authorization", "Bearer "+token)
		}
	}
}

func (hc *Client) DoJSON(ctx context.Context, method, path string, reqBody any, extraHeaders map[string]string) (Response, error) {
	var body io.Reader

	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return Response{}, fmt.Errorf("HttpClient: marshale json: %w", err)
		}
		body = bytes.NewReader(b)
	}

	request, err := http.NewRequestWithContext(ctx, method, hc.baseURL+path, body)

	if err != nil {
		return Response{}, fmt.Errorf("HttpClient: request: %w", err)
	}

	for key, value := range hc.headers {
		request.Header[key] = value
	}

	for key, value := range extraHeaders {
		request.Header.Set(key, value)
	}

	response, err := hc.httpClient.Do(request)

	if err != nil {
		return Response{}, fmt.Errorf("HttpClient: do request: %w", err)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)

	if err != nil {
		return Response{}, fmt.Errorf("HttpClient: read body: %w", err)
	}

	return Response{
		StatusCode: response.StatusCode,
		Body:       b,
		Headers:    response.Header,
	}, nil

}
