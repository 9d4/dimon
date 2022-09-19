package client

import (
	"context"
	"encoding/json"

	"github.com/9d4/dimon/server"
)

func (cli *Client) ProcessList(ctx context.Context) ([]server.Process, error) {
	resp, err := cli.get(ctx, "/processes", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var processes []server.Process
	err = json.NewDecoder(resp.body).Decode(&processes)
	if err != nil {
		return nil, err
	}

	return processes, err
}
