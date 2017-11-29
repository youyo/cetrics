package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cAdvisorUrl   string
	containerID   string
	metrics       string
	format        string
	discoveryData zabbixDiscoveryData
)

var RootCmd = &cobra.Command{
	Use:   "cetrics",
	Short: "cAdvisor metrics fetcher",
	//Long:  `A brief description of your application`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cAdvisorUrl, "cadvisor-url", "u", "http://127.0.0.1:8080/", "cAdvisor URL")
}

func initConfig() {
	//viper.SetEnvPrefix("cetrics")
	//viper.AutomaticEnv()
}
