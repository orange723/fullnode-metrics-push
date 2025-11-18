package main

import (
	"log"
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

var (
	once sync.Once
)

type RpcConfig struct{}

type ChainConfig struct {
	Server struct {
		Host string `toml:"host"`
		Port int64  `toml:"port"`
	} `toml:"server"`
	RPC []struct {
		Chain string   `toml:"chain"`
		List  []string `toml:"list"`
	} `toml:"rpc"`
}

func (r *RpcConfig) GetConfig(config string) *ChainConfig {
	conf := new(ChainConfig)
	once.Do(func() {
		data, err := os.ReadFile(config)
		if err != nil {
			log.Default().Fatal(err.Error())
		}

		err = toml.Unmarshal(data, &conf)
		if err != nil {
			log.Default().Fatal(err.Error())
		}
	})

	return conf
}
