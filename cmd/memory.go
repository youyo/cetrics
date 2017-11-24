package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	cAdvisor "github.com/google/cadvisor/client/v2"
	v1 "github.com/google/cadvisor/info/v1"
	v2 "github.com/google/cadvisor/info/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "memory",
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
			case "all":
				allMemoryMetrics(container.Stats[0].Memory)
			case "usage":
				fmt.Println(container.Stats[0].Memory.Usage)
			case "max-usage":
				fmt.Println(container.Stats[0].Memory.MaxUsage)
			case "cache":
				fmt.Println(container.Stats[0].Memory.Cache)
			case "rss":
				fmt.Println(container.Stats[0].Memory.RSS)
			case "swap":
				fmt.Println(container.Stats[0].Memory.Swap)
			default:
				allMemoryMetrics(container.Stats[0].Memory)
			}
		}
	},
}

func allMemoryMetrics(m *v1.MemoryStats) {
	data := [][]string{
		[]string{"usage", strconv.FormatUint(m.Usage, 10)},
		[]string{"max-usage", strconv.FormatUint(m.MaxUsage, 10)},
		[]string{"cache", strconv.FormatUint(m.Cache, 10)},
		[]string{"rss", strconv.FormatUint(m.RSS, 10)},
		[]string{"swap", strconv.FormatUint(m.Swap, 10)},
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
	resourceCmd.AddCommand(memoryCmd)
	memoryCmd.Flags().StringVarP(&containerID, "container-id", "c", "", "Container ID")
	memoryCmd.Flags().StringVarP(&metrics, "metrics", "m", "all", "Metrics [all, usage, max-usage, cache, rss, swap]")
}
