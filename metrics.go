package main

import (
	"context"
	"fmt"
	"time"

	"github.com/castai/promwrite"
)

type Metric struct{}

func (c *Metric) Push(host, chain, provider string, port, blocknumber int64) {
	server := fmt.Sprintf("http://%s:%d/api/v1/write", host, port)

	client := promwrite.NewClient(server)
	_, _ = client.Write(context.Background(), &promwrite.WriteRequest{
		TimeSeries: []promwrite.TimeSeries{
			{
				Labels: []promwrite.Label{
					{
						Name:  "__name__",
						Value: chain,
					},
					{
						Name:  "provider",
						Value: provider,
					},
				},
				Sample: promwrite.Sample{
					Time:  time.Now(),
					Value: float64(blocknumber),
				},
			},
		},
	})
}
