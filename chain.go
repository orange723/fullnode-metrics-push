package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
)

type Chain struct{}

type BlockNumber struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

func (c *Chain) GetBlockNumber(url, chain string, wg *sync.WaitGroup, chainConfig ChainConfig) {
	defer wg.Done()

	var number int64

	client := req.SetTimeout(1 * time.Second)

	resp, err := client.R().
		SetBody(&BlockNumber{
			Jsonrpc: "2.0",
			Method:  "eth_blockNumber",
			Params:  []any{},
			ID:      1,
		}).
		Post(url)
	if err != nil {
		number = 0
		s := fmt.Sprintf("error: %s %d", url, number)
		fmt.Println(s)
		monitor.Push(chainConfig.Server.Host, chain, url, chainConfig.Server.Port, number)
	}

	value := gjson.Get(resp.String(), "result")

	number, _ = c.ConvertHexBlockNumber(value.String())

	monitor.Push(chainConfig.Server.Host, chain, url, chainConfig.Server.Port, number)
}

func (c *Chain) ConvertHexBlockNumber(hex string) (int64, error) {
	if len(hex) > 2 && hex[:2] == "0x" {
		hex = hex[2:]
	}

	blockNumber, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, fmt.Errorf(" %s error: %v", hex, err)
	}

	return blockNumber, nil
}
