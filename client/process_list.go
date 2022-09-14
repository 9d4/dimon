package client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/9d4/dimon/process"
)

func (cli *Client) ProcessList(ctx context.Context) ([]process.Process, error) {
	resp, err := cli.get(ctx, "/processes", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		log.Fatal(err)
	}

	var process []process.Process
	err = json.NewDecoder(resp.body).Decode(&process)
	return process, err
}
