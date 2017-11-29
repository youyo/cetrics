package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	cAdvisor "github.com/google/cadvisor/client/v2"
	v2 "github.com/google/cadvisor/info/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var filesystemCmd = &cobra.Command{
	Use:   "filesystem",
	Short: "filesystem",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cAdvisor.NewClient(cAdvisorUrl)
		if err != nil {
			log.Fatal(err)
		}
		request := v2.RequestOptions{
			IdType:    "docker",
			Count:     1,
			Recursive: false,
		}
		containers, err := client.Stats(containerID, &request)
		if err != nil {
			log.Fatal(err)
		}
		for _, container := range containers {
			switch metrics {
			case "All":
				allFileSystemMetrics(container.Stats[0].Filesystem)
			case "TotalUsageBytes":
				fmt.Println(*container.Stats[0].Filesystem.TotalUsageBytes)
			case "BaseUsageBytes":
				fmt.Println(*container.Stats[0].Filesystem.BaseUsageBytes)
			case "InodeUsage":
				fmt.Println(*container.Stats[0].Filesystem.InodeUsage)
			default:
				allFileSystemMetrics(container.Stats[0].Filesystem)
			}
		}
	},
}

func allFileSystemMetrics(f *v2.FilesystemStats) {
	data := [][]string{
		[]string{"TotalUsageBytes", strconv.FormatUint(*f.TotalUsageBytes, 10)},
		[]string{"BaseUsageBytes", strconv.FormatUint(*f.BaseUsageBytes, 10)},
		[]string{"InodeUsage", strconv.FormatUint(*f.InodeUsage, 10)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Bytes"})
	table.SetBorders(
		tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false},
	)
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
}

func init() {
	resourceCmd.AddCommand(filesystemCmd)
	filesystemCmd.Flags().StringVarP(&containerID, "container-id", "c", "", "Container ID")
	filesystemCmd.Flags().StringVarP(&metrics, "metrics", "m", "all", "Metrics")
}
