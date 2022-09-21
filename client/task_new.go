package client

import (
	"context"
	"encoding/json"

	"github.com/9d4/dimon/task"
)

func (cli *Client) TaskNew(ctx context.Context, name string, command string, args ...string) (task.Task, error) {
	data := map[string]interface{}{
		"name":    name,
		"command": command,
		"args":    args,
	}

	var task task.Task

	resp, err := cli.post(ctx, "/tasks", nil, data, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return task, err
	}

	err = json.NewDecoder(resp.body).Decode(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}
