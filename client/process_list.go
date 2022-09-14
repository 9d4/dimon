package client

import (
	"context"
	"io"
	"log"
)

func (cli *Client) ProcessList(ctx context.Context) {
	resp, err := cli.get(ctx, "/processes", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.body)
	if err == nil {
		log.Println(string(body))
	}
}
