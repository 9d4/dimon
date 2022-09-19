package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"

	v "github.com/spf13/viper"
)

// Client is the API client that performs all operations
// against dimon server.
type Client struct {
	host string

	// client used to send and receive http requests.
	client *http.Client

	customHTTPHeaders map[string]string
}

func NewClient() *Client {
	var host = v.GetString("socketpath")

	return &Client{
		host: host,
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, _ string) (net.Conn, error) {
					return net.Dial("unix", host)
				},
			},
		},
	}
}

func (cli *Client) head(ctx context.Context, path string, query url.Values, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest(ctx, http.MethodHead, path, query, nil, headers)
}

func (cli *Client) get(ctx context.Context, path string, query url.Values, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest(ctx, http.MethodGet, path, query, nil, headers)
}

func (cli *Client) post(ctx context.Context, path string, query url.Values, obj interface{}, headers map[string][]string) (serverResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return serverResponse{}, err
	}

	return cli.sendRequest(ctx, http.MethodPost, path, query, body, headers)
}

type serverResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
	reqURL     *url.URL
}

func (cli *Client) addHeaders(req *http.Request, headers headers) *http.Request {
	for k, v := range cli.customHTTPHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range headers {
		req.Header[http.CanonicalHeaderKey(k)] = v
	}
	return req
}

func (cli *Client) buildRequest(method, path string, body io.Reader, headers headers) (*http.Request, error) {
	expectedPayload := (method == http.MethodPost || method == http.MethodPut)
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req = cli.addHeaders(req, headers)

	req.Host = "dimon"
	req.URL.Host = "dimon"
	req.URL.Scheme = "http"

	if expectedPayload && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "text/plain")
	}

	return req, nil
}

func (cli *Client) sendRequest(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	body io.Reader,
	headers headers,
) (serverResponse, error) {
	req, err := cli.buildRequest(method, cli.getAPIPath(ctx, path, query), body, headers)
	if err != nil {
		return serverResponse{}, err
	}

	return cli.doRequest(ctx, req)
}

func (cli *Client) doRequest(ctx context.Context, req *http.Request) (serverResponse, error) {
	serverResp := serverResponse{statusCode: -1, reqURL: req.URL}

	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		// Don't decorate context sentinel errors; users may be comparing to
		// them directly.
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return serverResp, err
		}

		if nErr, ok := err.(*url.Error); ok {
			if nErr, ok := nErr.Err.(*net.OpError); ok {
				if os.IsPermission(nErr.Err) {
					return serverResp, fmt.Errorf("%v, permission denied while trying to connect to the Dimon daemon socket at %v", err, cli.host)
				}
			}
		}
	}

	if resp != nil {
		serverResp.statusCode = resp.StatusCode
		serverResp.body = resp.Body
		serverResp.header = resp.Header
	}

	return serverResp, err
}

func ensureReaderClosed(response serverResponse) {
	if response.body != nil {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(io.Discard, response.body, 512)
		response.body.Close()
	}
}
