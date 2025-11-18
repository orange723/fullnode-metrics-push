package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/config.example.toml", "config file (default is config/config.example.toml)")
}

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "fullnode-metrics-push",
	Short: "Get Full node data and push metrics to prometheus",
	Run: func(cmd *cobra.Command, args []string) {
		chainConfig := config.GetConfig(cfgFile)

		for _, v := range chainConfig.RPC {
			for _, url := range v.List {
				wg.Add(1)

				go chain.GetBlockNumber(url, v.Chain, &wg, *chainConfig)
			}
		}
		wg.Wait()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
