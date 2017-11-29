package cmd

import (
	"fmt"
	"log"

	cAdvisor "github.com/google/cadvisor/client/v2"
	"github.com/spf13/cobra"
)

var numCoresCmd = &cobra.Command{
	Use:   "num-cores",
	Short: "Number of cpu cores",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cAdvisor.NewClient(cAdvisorUrl)
		if err != nil {
			log.Fatal(err)
		}
		machineInfo, err := client.MachineInfo()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(machineInfo.NumCores)
	},
}

func init() {
	machineCmd.AddCommand(numCoresCmd)
}
