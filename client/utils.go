package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/url"
	"path"
)

type headers map[string][]string

func encodeBody(obj interface{}, headers headers) (io.Reader, headers, error) {
	if obj == nil {
		return nil, headers, nil
	}

	body, err := encodeData(obj)
	if err != nil {
		return nil, headers, err
	}
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Type"] = []string{"application/json"}
	return body, headers, nil
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}

func (cli *Client) getAPIPath(ctx context.Context, p string, query url.Values) string {
	var apiPath string
	apiPath = path.Join(p)
	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}
