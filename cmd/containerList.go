package cmd

import (
	"fmt"
	"log"

	cAdvisor "github.com/google/cadvisor/client/v2"
	v2 "github.com/google/cadvisor/info/v2"
	"github.com/spf13/cobra"
)

var containerListCmd = &cobra.Command{
	Use:   "container-list",
	Short: "List containers",
	//Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cAdvisor.NewClient(cAdvisorUrl)
		if err != nil {
			log.Fatal(err)
		}
		request := v2.RequestOptions{
			IdType:    "docker",
			Count:     1,
			Recursive: true,
		}
		containers, err := client.Stats("", &request)
		if err != nil {
			log.Fatal(err)
		}
		switch format {
		case "string":
			for _, container := range containers {
				fmt.Println(container.Spec.Aliases[0])
			}
		case "zabbix-discovery":
			for _, container := range containers {
				discoveryData = append(discoveryData, zabbixDiscoveryItem{
					"#CONTAINER_NAME":  container.Spec.Aliases[0],
					"#CONTAINER_ID":    container.Spec.Aliases[1],
					"#CONTAINER_IMAGE": container.Spec.Image,
				})
			}
			fmt.Println(discoveryData.Json())
		default:
			fmt.Println("Format is not match. Use 'string' format.")
			for _, container := range containers {
				fmt.Println(container.Spec.Aliases[0])
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(containerListCmd)
	containerListCmd.Flags().StringVarP(&format, "format", "f", "string", "Output format [string, zabbix-discovery]")
}
