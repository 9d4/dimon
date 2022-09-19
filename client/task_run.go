package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/9d4/dimon/task"
)

func (cli *Client) TaskRun(ctx context.Context, taskID int) (task.Task, error) {
	resp, err := cli.post(ctx, fmt.Sprintf("/tasks/%d", taskID), nil, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		log.Fatal(err)
	}

	var task task.Task
	if err := json.NewDecoder(resp.body).Decode(&task); err != nil {
		fmt.Printf("%+v\n", task)
		return task, err
	}

	return task, nil
}
