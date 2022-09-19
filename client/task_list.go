package client

import (
	"context"
	"encoding/json"

	"github.com/9d4/dimon/server"
)

func (cli *Client) TaskList(ctx context.Context) ([]server.Task, error) {
	resp, err := cli.get(ctx, "/tasks", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var tasks []server.Task
	err = json.NewDecoder(resp.body).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
