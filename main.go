package main

import (
	"sync"
)

var (
	config  RpcConfig
	monitor Metric
	chain   Chain
	wg      sync.WaitGroup
)

func main() {
	Execute()
}
